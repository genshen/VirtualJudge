package controllers

import (
	"gensh.me/VirtualJudge/components/context/problem"
	"strconv"
)

type ProblemController struct {
	BaseController
}

//todo add Pagination
func (c *ProblemController)Problems() {
	c.Data["json"] = problem.FindProblems()
	c.ServeJSON()
}

func (c *ProblemController)Detail() {
	id,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.Data["json"] = problem.LoadProblemDetail(id)
	c.ServeJSON()
}

func (c *ProblemController)AddProblem() {
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