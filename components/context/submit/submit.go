package submit

import (
	"time"
	"encoding/base64"
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/models"
	"gensh.me/VirtualJudge/models/database"
	"github.com/astaxie/beego/orm"
	"gensh.me/VirtualJudge/components/crawler/submitter"
	u "gensh.me/VirtualJudge/components/crawler/utils"
	"log"
)

const DATETIME_LAYOUT = "2006-01-02 15:04:05"

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
			if err := submitter.SubmitProblem(int(id), int(problem.OjType), problem.OriginId, s.Language, s.Code, onSubmitResult); err != nil {
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

//todo normal all errors
func onSubmitResult(localSubmissionId int, ojType int, accountIndex uint, accountUsername string, status *submitter.SubmitStatus, err error) {
	var errText = ""
	if err != nil {
		errText = err.Error()
	}

	switch {
	case status.StatusCode == 0 || status.StatusCode < u.STATUS_DIV_ERROR:
		fallthrough
	case status.StatusCode < u.STATUS_DIV_LOCAL:
		log.Println("error:", errText)
		log.Println("error submit sulotion to remote OJ,local RunId:", localSubmissionId)
		_, err := database.O.Raw("UPDATE submission SET status_code = ? , updated_at = ? , error_detail = ? WHERE id = ?",
			status.StatusCode, time.Now().Format(DATETIME_LAYOUT), errText, localSubmissionId).Exec()
		if err != nil {
			log.Println("error write to database in onSubmitResult local RunId:", localSubmissionId)
			return
		}
	case status.StatusCode < u.STATUS_DIV_REMOTE_PENDING:
		_, err := database.O.Raw("UPDATE submission SET status_code = ? , origin_run_id = ? , origin_account_id = ? , " +
			"query_count = ? , origin_submit_time = ?, updated_at = ?  WHERE id = ?",
			status.StatusCode, status.RunId, accountUsername,1, status.SubmitTime.Format(DATETIME_LAYOUT),
			time.Now().Format(DATETIME_LAYOUT), localSubmissionId).Exec()
		if err != nil {
			log.Println("error write to database in onSubmitResult local RunId:", localSubmissionId)
			return
		}
		log.Println("successfully submit to remote OJ (will query remote status later),local RunId:", localSubmissionId)
	//todo add to queue

	case status.StatusCode < u.STATUS_DIV_END: //task finished and add to database
		_, err := database.O.Raw("UPDATE submission SET status_code = ? , execute_time = ? , memory = ? , origin_run_id = ? , " +
			"origin_account_id = ? , query_count = ? , origin_submit_time = ? ,updated_at = ?  WHERE id = ?",
			status.StatusCode, status.ExecuteTime, status.Memory, status.RunId, accountUsername,1,
			status.SubmitTime.Format(DATETIME_LAYOUT), time.Now().Format(DATETIME_LAYOUT), localSubmissionId).Exec()
		if err != nil {
			log.Println("error write to database in onSubmitResult,local RunId:", localSubmissionId)
			return
		}
		log.Println("successfully submit to remote OJ,local RunId:", localSubmissionId)
	}
	//todo update query time
}