package controllers

import (
	"github.com/astaxie/beego"
	"encoding/gob"
	"gensh.me/VirtualJudge/components/auth"
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/components/context/user"
)

const (
	UserData = "user_data"
	IsAuth = "is_auth"
)

var Login_Json_Err = utils.SimpleJsonResponse{Status:3, Error:"Unauthorized"}

type Rules interface {
	getRules(string) int
}

type BaseController struct {
	beego.Controller
}

func init() {
	gob.Register(auth.User{})
}

// Prepare implemented Prepare method for baseRouter.
func (base *BaseController) Prepare() {
	if app, ok := base.AppController.(Rules); ok {
		var _, action = base.GetControllerAndAction()
		rule := app.getRules(action)
		is_login := base.isAuthenticated()
		if ((rule & user.Login) == user.Login) && !is_login {
			base.Redirect("/user/auth", 302)
		} else if ((rule & user.LoginJSON) == user.LoginJSON) && !is_login {
			base.Data["json"] = &Login_Json_Err
			base.Ctx.Output.Status = 401
			base.ServeJSON()
			base.StopRun()
		}
	}
}

func (base *BaseController)isAuthenticated() bool {
	is_auth := base.GetSession(IsAuth)
	if is_auth == nil {
		return false
	}
	return is_auth.(bool)
}

func (this *BaseController)GetUserData() auth.User {
	u, ok := this.GetSession(UserData).(auth.User)
	if !ok {
		return auth.User{}
	}
	return u
}