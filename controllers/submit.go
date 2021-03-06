package controllers

import (
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/components/context/submit"
	"gensh.me/VirtualJudge/components/context/user"
)

type SubmitController struct {
	BaseController
}

var profile_rules = map[string]int{
	"Submit": user.LoginJSON,
}

func (c *SubmitController) getRules(action string) int {
	return profile_rules[action]
}

func (s *SubmitController)Submit() {
	language, err := s.GetInt8("language")
	problemId, err_ := s.GetInt("problem_id")
	if err == nil && err_ == nil {
		u := s.GetUserData()
		submitForm := submit.SubmitForm{Language:language, ProblemId:problemId, Code:s.GetString("code")}
		response := submitForm.Valid(u.Id, u.Name)
		s.Data["json"] = response
	} else {
		s.Data["json"] = utils.SimpleJsonResponse{Status:0, Error:"error parsing submitted data"}
	}
	s.ServeJSON()
}