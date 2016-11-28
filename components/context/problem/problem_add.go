package problem

import (
	"strconv"
	"gensh.me/VirtualJudge/components/crawler/problem"
	"gensh.me/VirtualJudge/models"
	"gensh.me/VirtualJudge/components/utils"
	"time"
	"github.com/astaxie/beego/orm"
	"gensh.me/VirtualJudge/models/database"
)

type ProblemAddMeta struct {
	Type      string
	ProblemId string
}

func (p *ProblemAddMeta)ValidAndSave() (*utils.SimpleJsonResponse) {
	ty, err := strconv.ParseInt(p.Type, 10, 8)
	_, err_ := strconv.ParseInt(p.ProblemId, 10, 64)
	if err != nil && err_ != nil {
		return &utils.SimpleJsonResponse{Status:0, Error:"error oj type or error problem id"}
	}
	pm, err_c := problem.CrawlerProblem(p.ProblemId, int8(ty))

	if err_c != nil {
		return &utils.SimpleJsonResponse{Status:0, Error:err_c.Error()}
	}

	//todo check problem exist
	err = database.O.QueryTable(models.ProblemTableName).Filter("oj", p.Type).Filter("origin_id", p.ProblemId).One(&models.Problem{}, "id")
	if err == orm.ErrNoRows {
		//then  add this problem
		o := orm.NewOrm()
		o.Begin()
		timeNow := time.Now()
		prob := models.Problem{Title:pm.Title, Source:pm.Source, SourceUrl:pm.SourceUrl,
			Oj:int8(ty), OriginUrl:pm.OriginUrl, OriginId:p.ProblemId,
			MemLimit:pm.MemLimit, TimeLimit:pm.TimeLimit, CreatedAt:timeNow, UpdatedAt:timeNow}
		_, err = o.Insert(&prob)
		if err != nil {
			o.Rollback()
			return &utils.SimpleJsonResponse{Status:0, Error:"error while saving problem data"}
		}
		problemDetail := models.ProblemDetail{Problem:&prob, Describe:pm.Describe, Input:pm.Input, Output:pm.Output,
			InputSample:pm.InputSample, OutputSample:pm.OutputSample, Hint:pm.Hint, UpdatedAt:timeNow}
		_, err = o.Insert(&problemDetail)
		if err != nil {
			o.Rollback()
			return &utils.SimpleJsonResponse{Status:0, Error:"error while saving problem data"}
		} else {
			o.Commit()
			return &utils.SimpleJsonResponse{Status:1, Addition:&problemDetail}
		}
	} else {
		return &utils.SimpleJsonResponse{Status:0, Error:"the problem already exists"}
	}
}