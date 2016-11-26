package controllers

import "gensh.me/VirtualJudge/components/context/problem"

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.TplName = "home.html"
	//c.SetSession()
}

func (c *MainController)AddProblem() {
	if c.Ctx.Input.Method() == "POST" {
		pro := c.Input().Get("problem_id")
		ty := c.Input().Get("oj_type")
		pam := problem.ProblemAddMeta{Type:ty, ProblemId:pro}
		result := pam.ValidAndSave()
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		c.TplName = "problem_add.html"
	}
}