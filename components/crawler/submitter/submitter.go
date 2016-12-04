package submitter

import (
	"errors"
	"gensh.me/VirtualJudge/components/crawler/accounts"
	"time"
	"log"
	"gensh.me/VirtualJudge/components/crawler/utils"
)

type SubmitInterface interface {
	RemoteSubmit(session string, username, problemId string, language  int8, code  string) (*SubmitStatus, error)
	GetLanguageType(int8) int8
}

type SubmitStatus struct {
	RunId       string
	StatusCode  int8
	Memory      string
	ExecuteTime string
	SubmitTime  *time.Time
}

//remember to follow the order in crawler/utils/values.go->const
var submitInterfaces = []SubmitInterface{new(POJSubmitInterface)}

func SubmitProblem(localSubmissionId int, ojType int, problemId string, language int8, code string,
callback func(int, int, uint, string, *SubmitStatus, error)) error {
	accountInterface, err := accounts.GetInterfaceByOjType(ojType)
	if err != nil {
		return err
	}
	//get submitter interface
	if index := ojType - 1; index < len(submitInterfaces) && index >= 0 {
		si := submitInterfaces[index]
		if l := si.GetLanguageType(language); l == -1 {
			return errors.New("no such language")
		} else {
			go submit(si, accountInterface, localSubmissionId, ojType, problemId, l, code, callback)
			return nil
		}
	} else {
		return errors.New("no OJ matched")
	}
}

//get available account and submit problem
//we just try to login the account once as default now
func submit(si SubmitInterface, accountInterface accounts.AccountInterface, localSubmissionId int, ojType int, problemId string, language int8, code string,
callback func(int, int, uint, string, *SubmitStatus, error)) {
	_, account := accountInterface.GetAvailableAccount()
	if accounts.GetSessionByIndex(account.Index) == "" {
		//login and update sessions
		if err := accountInterface.LoginAccount(account); err != nil {
			log.Println("login account faild while submitting solution") //todo counts fail time
			callback(localSubmissionId, ojType, account.Index, account.Username, &SubmitStatus{StatusCode:utils.STATUS_KNOWN_ERROR}, err)
			return
		}
	}

	if status, err := si.RemoteSubmit(accounts.GetSessionByIndex(account.Index), account.Username, problemId, language, code); err != nil {
		//todo login and try again
		//todo if err == notLogin ->try again
		log.Println(err.Error())
		if err := accountInterface.LoginAccount(account); err != nil {
			log.Println("login account faild while submitting solution") //todo counts fail time
			callback(localSubmissionId, ojType, account.Index, account.Username, &SubmitStatus{StatusCode:utils.STATUS_KNOWN_ERROR}, err)
			return
		}
		if status, err := si.RemoteSubmit(accounts.GetSessionByIndex(account.Index), account.Username, problemId, language, code); err != nil {
			log.Println(err.Error())
			callback(localSubmissionId, ojType, account.Index, account.Username, status, nil)
		}
	} else {
		callback(localSubmissionId, ojType, account.Index, account.Username, status, nil)
	}
	//todo account.Tasks++;
}