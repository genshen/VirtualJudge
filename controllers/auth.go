package controllers

import "gensh.me/VirtualJudge/components/auth"

type AuthController struct {
	BaseController
}

func (c *AuthController) GetUserStatus() {
	if c.HasAuthenticated(){
		c.Data["json"] = c.GetUserData()
	}else{
		c.Data["json"] = auth.User{}
	}
	c.ServeJSON()
}