package routers

import (
	"gensh.me/VirtualJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Home")
	//auth
	beego.Router("/auth/user_status", &controllers.AuthController{}, "get:GetUserStatus")  //JSON
	beego.Router("/auth/callback/github", &controllers.AuthController{}, "get:GithubCallback")

	beego.Router("/problems", &controllers.ProblemController{}, "get:Problems")
	beego.Router("/problem/detail/:id([0-9]+)", &controllers.ProblemController{}, "get:Detail")
	beego.Router("/problem/summary/:id([0-9]+)", &controllers.ProblemController{}, "get:Summary")
	beego.Router("/problem/add", &controllers.ProblemController{}, "get,post:AddProblem")

	beego.Router("/submit", &controllers.SubmitController{}, "post:Submit")

	beego.Router("/status", &controllers.StatusController{}, "get:Status")
	beego.Router("/status/ws", &controllers.StatusController{}, "get:WebSocket")
	beego.Router("/status/test", &controllers.StatusController{}, "get:Test")
}
