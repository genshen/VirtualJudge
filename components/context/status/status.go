package status

import (
	"log"
	"time"
	"gensh.me/VirtualJudge/models/database"
	"gensh.me/VirtualJudge/components/crawler/utils"
	"gensh.me/VirtualJudge/components/crawler/submitter"
	s "gensh.me/VirtualJudge/components/crawler/status"
	u "gensh.me/VirtualJudge/components/crawler/utils"
)

const DATETIME_LAYOUT = "2006-01-02 15:04:05"

type StatusData struct {
	Id          int
	StatusCode  int8
	ExecuteTime string
	Memory      string
	UpdatedTime time.Time
}

func onStatusChanged(result *s.TaskResult) {
	now := time.Now()
	messages <- StatusData{Id:result.LocalRunId, Memory:result.Memory, ExecuteTime:result.ExecuteTime,
		StatusCode:result.StatusCode, UpdatedTime:now}
}

//todo normal all errors
func onSubmittedStatusChangedListener(localSubmissionId int, ojType int8, accountIndex uint, accountUsername string, status *submitter.SubmitStatus, err error) {
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
			status.StatusCode, status.RunId, accountUsername, 1, status.SubmitTime.Format(DATETIME_LAYOUT),
			time.Now().Format(DATETIME_LAYOUT), localSubmissionId).Exec()
		if err != nil {
			log.Println("error write to database in onSubmitResult local RunId:", localSubmissionId)
			return
		}
		log.Println("successfully submit to remote OJ (will query remote status later),local RunId:", localSubmissionId)
		s.AddTaskToQueue(localSubmissionId, ojType, accountIndex, status.StatusCode, status.RunId)

	case status.StatusCode < u.STATUS_DIV_END: //task finished and add to database
		_, err := database.O.Raw("UPDATE submission SET status_code = ? , execute_time = ? , memory = ? , origin_run_id = ? , " +
			"origin_account_id = ? , query_count = ? , origin_submit_time = ? ,updated_at = ?  WHERE id = ?",
			status.StatusCode, status.ExecuteTime, status.Memory, status.RunId, accountUsername, 1,
			status.SubmitTime.Format(DATETIME_LAYOUT), time.Now().Format(DATETIME_LAYOUT), localSubmissionId).Exec()
		if err != nil {
			log.Println("error write to database in onSubmitResult,local RunId:", localSubmissionId)
			return
		}
		log.Println("successfully submit to remote OJ,local RunId:", localSubmissionId)
	}
	//todo update query times

	now := time.Now()
	messages <- StatusData{Id:localSubmissionId, Memory:status.Memory, ExecuteTime:status.ExecuteTime,
		StatusCode:status.StatusCode, UpdatedTime:now}
}

func Test() {
	s.AddTaskToQueue(1, utils.POJ, 0, 9, "16359867")
}