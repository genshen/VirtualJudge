package controllers

import (
	"gensh.me/VirtualJudge/components/auth"
	users "gensh.me/VirtualJudge/components/context/user"
	"encoding/json"
)

type AuthController struct {
	BaseController
}

type UserStatus struct {
	IsLogin bool        `json:"is_login"`
	User    *auth.User  `json:"user"`
}

func (c *AuthController) GetUserStatus() {
	if c.isAuthenticated() {
		user := c.GetUserData()
		c.Data["json"] = &UserStatus{IsLogin:true, User:&user}
	} else {
		c.Data["json"] = &UserStatus{IsLogin:false, User:&auth.User{}}
	}
	c.ServeJSON()
}

func (c *AuthController) GithubCallback() {
	user := &auth.User{}
	if c.isAuthenticated() {
		u := c.GetUserData()
		user = &u
		user.Status = auth.UserStatusAlreadyAuth
	} else {
		code := c.GetString("code")
		github := auth.GithubAuthUser{}
		if len(code) > 0 {
			u, err := auth.StartAuth(&github, code)
			if err == nil {
				user = u
				user.Status = auth.UserStatusAuthOK
				c.loginUser(user)
				users.CreateUser(user,auth.Github)
			}
		}
	}
	b, _ := json.Marshal(user)
	c.Data["json"] = string(b)
	c.TplName = "auth_callback.html"
}

func (this *AuthController)loginUser(u *auth.User) {
	this.SetSession(UserData, *u)
	this.SetSession(IsAuth, true)
}