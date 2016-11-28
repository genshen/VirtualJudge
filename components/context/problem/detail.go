package problem

import (
	"gensh.me/VirtualJudge/models/database"
	"github.com/astaxie/beego/orm"
)

const SQL = "select * from problem " +
	"inner join problem_detail " +
	"on problem_detail.problem_id = problem.id" +
	" where problem.id = ?" +
	" limit 1"

func LoadProblemDetail(id int) *orm.Params {
	//problem := make(map[string]interface{})

	var problems []orm.Params
	num, err := database.O.Raw(SQL,id).Values(&problems)
	if err == nil && num > 0 {
		return &problems[0]
	}
	//problem := models.ProblemDetail{}
	//err := database.O.Raw(SQL,id).QueryRow(&problem)
	//err := database.O.QueryTable(models.ProblemDetailTableName).Filter("id", id).One(&problem)
	return nil
}
