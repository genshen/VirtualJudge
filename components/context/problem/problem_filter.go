package problem

import (
	"gensh.me/VirtualJudge/models/database"
	"gensh.me/VirtualJudge/models"
)

func FindProblems() (p []*models.Problem ){
	database.O.QueryTable(models.ProblemTableName).All(&p)
	return
}
