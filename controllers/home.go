package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Home() {
	c.TplName = "home.html"
	//c.SetSession()
}

//todo
func (c *MainController) About() {
	c.TplName = "about.html"
	//c.SetSession()
}