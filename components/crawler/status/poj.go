package status

import (
	"gensh.me/VirtualJudge/components/crawler/accounts"
	"log"
	"qiniupkg.com/x/errors.v7"
	"github.com/astaxie/beego/httplib"
	"strings"
	"gensh.me/VirtualJudge/components/crawler/utils"
)

type POJStatusInterface struct {

}

func (p POJStatusInterface) FetchStatus(ai accounts.AccountInterface, si StatusInterface, accountIndex uint, runId string) (*TaskResult, error) {
	var session string
	if session = accounts.GetSessionByIndex(accountIndex); session == "" {
		//login and update sessions
		if err := ai.LoginAccountByIndex(accountIndex); err != nil {
			//todo counts fail time
			log.Println(err)
			return nil, errors.New("login account faild while fetch problem status")
		}
	}

	if result, err := p.requestStatus(accounts.GetSessionByIndex(accountIndex), runId); err != nil {
		//todo login and try again
		//todo if err == notLogin ->try again
		log.Println(err.Error())
		//login and update sessions
		if err := ai.LoginAccountByIndex(accountIndex); err != nil {
			//todo counts fail time
			log.Println(err)
			return nil, errors.New("login account faild while fetch problem status")
		}

		if result, err := p.requestStatus(accounts.GetSessionByIndex(accountIndex), runId); err != nil {
			log.Println(err)
			return nil, err
		} else {
			return result, nil
		}
	} else {
		return result, nil
	}
}

//todo not login handle
func (p POJStatusInterface)requestStatus(session string, runId string) (*TaskResult, error) {
	req := httplib.Get("http://poj.org/showsource?solution_id=" + runId)
	req.Header("Cookie", session)
	body, err := req.String()
	if err != nil {
		return nil, err
	}
	if len(body) >= 50 {
		if (strings.Contains(body[:50], `"Pragma"`)) {
			return nil, errors.New("not login")
		} else {
			_, po := utils.FindMatchString(body, "<tr><td><b>", "</td>", "</tr>")
			if po != -1 && po < len(body) {
				body = body[po:]
			}
			mem, po := utils.FindMatchString(body, "<tr><td><b>", "</b>", "</td>")
			if po != -1 && po < len(body) {
				body = body[po:]
			}

			time, po := utils.FindMatchString(body, "<td><b>", "</b>", "</td></tr>")
			if po != -1 && po < len(body) {
				body = body[po:]
			}

			statusName, po := utils.FindMatchString(body, "color=", ">", "</font></td></tr>")

			return &TaskResult{Memory:mem, ExecuteTime:time, StatusCode:GetStatusByOJType(utils.POJ - 1,statusName)}, nil
		}
	}
	return nil, errors.New("unkonwn reason")
}