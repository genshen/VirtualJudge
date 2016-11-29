package submitter

import (
	"errors"
	"gensh.me/VirtualJudge/components/crawler/accounts"
)

type SubmitInterface interface {
	RemoteSubmit(account *accounts.Account, problemId string, language  int8, code  string) error
	GetLanguageType(int8) int8
}

//remember to follow the order in crawler/utils/values.go->const
var submitInterfaces = []SubmitInterface{new(POJSubmitInterface)}

func SubmitProblem(ojType int, problemId string, language int8, code string, callback func()) error {
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
			go submit(si, accountInterface, problemId, l, code, callback)
			return nil
		}
	} else {
		return errors.New("no OJ matched")
	}
}

//get available account and submit problem
//we just try to login the account once as default now
func submit(si SubmitInterface, accountInterface accounts.AccountInterface, problemId string, language int8, code string, callback func()) {
	_, account := accountInterface.GetAvailableAccount()
	if account.Session == "" {
		if err := accountInterface.LoginAccount(account); err != nil {
			print("submit faild") //todo counts fail time
			return //todo callback login error
		}
	}
	if err := si.RemoteSubmit(account, problemId, language, code); err != nil {
		//todo login and try again
		if err := accountInterface.LoginAccount(account); err != nil {
			print("submit faild") //todo counts fail time
			return //todo callback
		}
		if err := si.RemoteSubmit(account, problemId, language, code); err != nil {
			return //todo callback
		}
		return //todo callback success
	}
	return //todo callback
	//account.Tasks++;
}