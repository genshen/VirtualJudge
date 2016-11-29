package submit

import (
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/models"
	"gensh.me/VirtualJudge/models/database"
	"github.com/astaxie/beego/orm"
	"gensh.me/VirtualJudge/components/crawler/submitter"
)

type SubmitForm struct {
	Code      string
	Language  int
	ProblemId int
}

const (
	LANG_C = iota
	LANG_CPP
	LANG_JAVA
	LANG_GCC
	LANG_GPP
	LANG_COUNT
)

func (s *SubmitForm)Valid() (response *utils.SimpleJsonResponse) {
	response.Status = 0
	if s.Code == "" {
		response.Error = "source code can not be blank"
		return
	} else if (s.Language < 0 || s.Language >= LANG_COUNT) {
		response.Error = "language is not in that range"
		return
	}

	//get problem by problem id
	problem := models.Problem{Id:s.ProblemId}
	err := database.O.Read(&problem, "id", "oj", "origin_id")
	if err == orm.ErrNoRows {
		response.Error = "problem not exists"
	} else if err == orm.ErrMissPK {
		//can remove this case.
		response.Error = "inner serve error"
	} else {
		if err := submitter.SubmitProblem(problem.Oj, problem.OriginId); err != nil {
			//has summit error,eg: network,error password
			response.Error = err.Error()
		} else {
			response.Status = 1
		}
	}
	return
}