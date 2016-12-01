package submitter

import (
	"gensh.me/VirtualJudge/components/crawler/utils"
	"net/http"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type POJSubmitInterface struct {

}

func (r POJSubmitInterface)RemoteSubmit(session string, problemId string, language int8, code string) error {
	client := &http.Client{}
	values := url.Values{}
	values.Set("problem_id", problemId)
	values.Set("language", strconv.Itoa(int(language)))
	values.Set("source", code)
	values.Set("encoded", "1")

	req, err := http.NewRequest("POST", "http://poj.org/submit", strings.NewReader(values.Encode()))
	if err != nil {
		return errors.New("error request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", session)
	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		return errors.New("error request")
	}
	return nil
}

//convert language type to poj language type
//from
//to 0:G++,1:GCC,2:JAVA,3:Pascal,4:C++,5:C,6:Fortran
func (r POJSubmitInterface) GetLanguageType(language int8) int8 {
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
