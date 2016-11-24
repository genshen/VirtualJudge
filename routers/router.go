package routers

import (
	"gensh.me/VirtualJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//auth
	beego.Router("/auth/user_status", &controllers.AuthController{}, "get:GetUserStatus")  //JSON
	beego.Router("/auth/github/callback", &controllers.AuthController{}, "get:GetUserStatus")  //JSON
}
