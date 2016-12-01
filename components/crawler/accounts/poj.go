package accounts

import (
	"gensh.me/VirtualJudge/components/crawler/utils"
	"net/http"
	"net/url"
	"errors"
)

type PojAccountInterface struct {

}

func (pi PojAccountInterface)LoginAccount(account *Account) error {
	client := &http.Client{CheckRedirect:func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	values := url.Values{}
	values.Set("user_id1", account.Username)
	values.Set("password1", account.Password)
	values.Set("B1", "login")
	values.Set("url", "/")
	response, err := client.PostForm("http://poj.org/login", values)

	if err == nil && response.StatusCode == 302 {
		for _, cookie := range response.Cookies() {
			if cookie.Name == "JSESSIONID" {
				UpdateSessionByIndex(account.Index, "JSESSIONID=" + cookie.Value)
				return nil
			}
		}
		return errors.New("no SESSION_ID in Cookies")
	}
	return errors.New("error request")
}

//make sure accountIndex is safe!
func (pi PojAccountInterface)LoginAccountByIndex(accountIndex uint) error {
	return pi.LoginAccount(&OJs[utils.POJ - 1].Accounts[accountIndex])
}

//todo select the minimal task account
func (pi PojAccountInterface)GetAvailableAccount() (uint, *Account) {

	return 0, &OJs[utils.POJ - 1].Accounts[0]
}