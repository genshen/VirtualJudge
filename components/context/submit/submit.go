package submit

import (
	"time"
	"encoding/base64"
	"github.com/astaxie/beego/orm"
	"gensh.me/VirtualJudge/models"
	"gensh.me/VirtualJudge/models/database"
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/components/crawler/submitter"
	u "gensh.me/VirtualJudge/components/crawler/utils"
)

type SubmitForm struct {
	Code      string
	Language  int8
	ProblemId int
}

func (s *SubmitForm)Valid(userId int, username string) (response *utils.SimpleJsonResponse) {
	response = &utils.SimpleJsonResponse{}
	response.Status = 0
	if s.Code == "" {
		response.Error = "source code can not be blank"
		return
	} else if (s.Language < 0 || s.Language >= u.LANG_COUNT) {
		response.Error = "language is not in that range"
		return
	}

	//get problem by problem id
	problem := models.Problem{}
	err := database.O.QueryTable(models.ProblemTableName).Filter("id", s.ProblemId).One(&problem, "id", "oj_type", "origin_id")
	if err == orm.ErrNoRows {
		response.Error = "problem not exists"
	} else if err == orm.ErrMissPK {
		//can remove this case.
		response.Error = "inner serve error"
	} else {
		temp_code_data, err := base64.StdEncoding.DecodeString(s.Code)
		if err != nil {
			response.Error = "code source decode error"
			return
		}
		codeData := string(temp_code_data)

		o := orm.NewOrm()
		o.Begin()
		//todo public?

		now := time.Now()

		submission := models.Submission{Contest:models.ContextIdDefault, ProblemId:s.ProblemId,
			OjType:problem.OjType, OriginId:problem.OriginId, Language:s.Language, SourceCode:codeData,
			CodeLength:len(codeData), Public:false, UserId:userId, UserName:username, CreatedAt:now, UpdatedAt:now}
		id, err := o.Insert(&submission)
		if err != nil {
			o.Rollback()
			response.Error = "inner serve error"
			return
		} else {
			if err := submitter.SubmitProblem(int(id), int(problem.OjType), problem.OriginId, s.Language, s.Code); err != nil {
				//has summit error,eg: network,error password
				o.Rollback()
				response.Error = err.Error()
				return
			} else {
				response.Status = 1
			}
		}
		o.Commit()
	}
	return
}