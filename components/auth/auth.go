package auth

import (
	"github.com/astaxie/beego/httplib"
	"strings"
	"errors"
	"gensh.me/blog/components/keys"
)

const LenAccessTokenName = 13 //access_token=
const (
	UserStatusUnAuth = iota
	UserStatusHasAuth
)

type GithubAuthUser struct {
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

type User struct {
	Status int    `json:"status"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Email  string `json:"-"`
	Avatar string `json:"avatar"`
}

type AuthInterface interface {
	GetAccessToken(code string) string
	GetAuthUserInfo(access_token string) (*User, error)
}

func (g GithubAuthUser) GetAccessToken(code string) string {
	req := httplib.Post(keys.GitHubKey.AuthUrl)
	req.Param("client_id", keys.GitHubKey.ClientId)
	req.Param("client_secret", keys.GitHubKey.ClientSecret)
	req.Param("code", code)
	req.Param("accept", "json")

	response, err := req.String()
	if err == nil {
		var index = strings.Index(response, "access_token=")
		if index >= 0 && len(response) > index + LenAccessTokenName {
			var after = response[index + LenAccessTokenName:]
			index = strings.Index(after, "&")
			if index < 0 {
				return after
			}
			return after[:index]
		}
	}
	return ""
}

func (g GithubAuthUser) GetAuthUserInfo(access_token string) (*User, error) {
	req := httplib.Get("https://api.github.com/user?access_token=" + access_token)
	err := req.ToJSON(&g)

	u := User{}
	u.Name = g.Name
	u.Url = g.HtmlUrl
	u.Email = g.Email
	u.Avatar = g.AvatarUrl
	return &u, err
}

func StartAuth(a AuthInterface, code string) (*User, error) {
	if access_token := a.GetAccessToken(code); access_token == "" {
		return &User{}, errors.New("can't fetch access_token")
	} else {
		return a.GetAuthUserInfo(access_token)
	}
}
