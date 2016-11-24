package auth

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"errors"
)

var GitHubKey  struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
}

type GithubAuthUser struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

func loadGithubKeys() {
	GitHubKey.AuthUrl = beego.AppConfig.String("github_auth_url")
	GitHubKey.ClientId = beego.AppConfig.String("github_client_id")
	GitHubKey.ClientSecret = beego.AppConfig.String("github_client_secret")
}

func (g GithubAuthUser) GetAccessToken(code string) string {
	req := httplib.Post(GitHubKey.AuthUrl)
	req.Param("client_id", GitHubKey.ClientId)
	req.Param("client_secret", GitHubKey.ClientSecret)
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
	if err != nil || g.Id == 0 {
		return &User{}, errors.New("error to parse json data")
	}

	u := User{}
	u.ThirdPartId = strconv.Itoa(g.Id)
	u.Name = g.Name
	u.Url = g.HtmlUrl
	u.Email = g.Email
	u.Avatar = g.AvatarUrl
	return &u, err
}
