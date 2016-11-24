package controllers

import (

)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.TplName = "home.html"
	//c.SetSession()
}