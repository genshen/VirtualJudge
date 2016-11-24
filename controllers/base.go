package controllers

import (
	"github.com/astaxie/beego"
	"gensh.me/blog/components/auth"
	"encoding/gob"
)

const (
	UserData = "user_data"
	IsAuth = "is_auth"
)

type BaseController struct {
	beego.Controller
}

func init(){
	gob.Register(auth.User{})
}

func (this *BaseController)HasAuthenticated() bool {
	is_auth := this.GetSession(IsAuth)
	if is_auth == nil {
		return false
	}
	return is_auth.(bool)
}

func (this *BaseController)GetUserData() auth.User {
	user, ok := this.GetSession(UserData).(auth.User)
	if !ok {
		return auth.User{}
	}
	return user
}