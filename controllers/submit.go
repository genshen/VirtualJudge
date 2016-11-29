package controllers

import (
	"gensh.me/VirtualJudge/components/utils"
	"gensh.me/VirtualJudge/components/context/submit"
)

type SubmitController struct {
	BaseController
}

/**
#include "stdio.h"
int main(){
int a,b;
scanf("%d%d",&a,&b);
printf("%d",a+b);
return 0;
}
*/
func (s *SubmitController)Submit() {
	language, err := s.GetInt8("language")
	problemId, err_ := s.GetInt("problem_id")
	if err == nil && err_ == nil {
		submitForm := submit.SubmitForm{Language:language, ProblemId:problemId, Code:s.GetString("code")}
		response := submitForm.Valid()
		s.Data["json"] = response
	} else {
		s.Data["json"] = utils.SimpleJsonResponse{Status:0, Error:"error parsing submitted data"}
	}
	s.ServeJSON()
}