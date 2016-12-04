package submitter

import (
	"net/http"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
	"gensh.me/VirtualJudge/components/crawler/utils"
	"github.com/astaxie/beego/httplib"
	"gensh.me/VirtualJudge/components/crawler/status"
)

type POJSubmitInterface struct {

}

func (p POJSubmitInterface)RemoteSubmit(session, username, problemId string, language int8, code string, ) (*SubmitStatus, error) {
	client := &http.Client{}
	values := url.Values{}
	lang := strconv.Itoa(int(language))
	values.Set("problem_id", problemId)
	values.Set("language", lang)
	values.Set("source", code)
	values.Set("encoded", "1")

	req, err := http.NewRequest("POST", "http://poj.org/submit", strings.NewReader(values.Encode()))
	if err != nil {
		return &SubmitStatus{StatusCode:utils.STATUS_UNKNOWN_ERROR}, errors.New("error submit request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", session)
	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		return &SubmitStatus{StatusCode:utils.STATUS_UNKNOWN_ERROR}, errors.New("error submit request")
	}

	queryUrl := "http://poj.org/status?problem_id=" + problemId + "&user_id=" + username + "&result=&language=" + lang
	return p.queryStatus(queryUrl, time.Now()) //server submit time in [now-2sec,now-28sec] is legal
}

func (p *POJSubmitInterface) queryStatus(queryUrl string, timeBeforeSubmit time.Time) (*SubmitStatus, error) {
	req := httplib.Get(queryUrl)
	body, err := req.Bytes()
	if err != nil {
		return &SubmitStatus{StatusCode:utils.STATUS_UNKNOWN_ERROR}, errors.New("error status request")
	}
	start, length := 0, len(body)
	var line string
	for k, v := range body {
		if v == '\n' {
			if length - start >= 16 {
				temp := string(body[start:start + 16])
				if strings.HasPrefix(temp, "<tr align=center") {
					line = string(body[start:k])
					run_id, po := utils.FindMatchString(line, "", "<td>", "</td>")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					s, po := utils.FindMatchString(line, "color=", ">", "<")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					mem, po := utils.FindMatchString(line, "", "<td>", "<")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					execute_time, po := utils.FindMatchString(line, "", "<td>", "<")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					//language
					_, po = utils.FindMatchString(line, "", "<td>", "<")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					_, po = utils.FindMatchString(line, "", "<td>", "<")
					if po != -1 && po < len(line) {
						line = line[po:]
					}

					date, po := utils.FindMatchString(line, "", "<td>", "<")
					serverSubmitDate, err := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
					if po != -1 && err == nil {
						tm2 := serverSubmitDate.Add(2 * time.Second)
						if tm2.After(timeBeforeSubmit) &&tm2.Before(timeBeforeSubmit.Add(30 * time.Second)) {
							return &SubmitStatus{RunId:run_id, StatusCode:p.convertStatus(s), Memory:mem,
								ExecuteTime:execute_time, SubmitTime:&serverSubmitDate}, nil
						} else if (tm2.Before(timeBeforeSubmit) || tm2.Equal(timeBeforeSubmit)) {
							return &SubmitStatus{StatusCode:utils.STATUS_KNOWN_ERROR}, errors.New("status not found")
						}
					}
				}
			}
			start = k + 1
		}
	}
	return &SubmitStatus{StatusCode:utils.STATUS_UNKNOWN_ERROR}, errors.New("status not found")
}

func (p *POJSubmitInterface)convertStatus(s string) int8 {
	return status.GetStatusByOJType(utils.POJ - 1, s)
}

//convert language type to poj language type
//from
//to 0:G++,1:GCC,2:JAVA,3:Pascal,4:C++,5:C,6:Fortran
func (p POJSubmitInterface) GetLanguageType(language int8) int8 {
	switch language {
	case utils.LANG_C:
		return 5
	case utils.LANG_CPP:
		return 4
	case utils.LANG_JAVA:
		return 2
	case utils.LANG_GCC:
		return 1
	case utils.LANG_GPP:
		return 0
	default:
		return 5 //Clang as default language
	}
}
