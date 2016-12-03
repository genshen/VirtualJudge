package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(Problem), new(ProblemDetail), new(Submission))
}

const (
	UserTableName = "user"
)
//one table is enough
type User struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Avatar       string        `json:"avatar"`
	Email        string        `json:"email"`
	School       string        `json:"school"`
	Url          string        `json:"url"`
	ThirdPartId  string        `json:"third_part_id"`
	ThirdPart    int           `json:"third_part"`
	PasswordHash string        `json:"_"`
	CreatedAt    time.Time     `json:"created_at"`
}

func (u *User) TableName() string {
	return UserTableName
}