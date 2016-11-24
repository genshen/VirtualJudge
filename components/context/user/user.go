package user

import (
	"time"
	"gensh.me/VirtualJudge/models"
	"gensh.me/VirtualJudge/components/auth"
	"github.com/astaxie/beego/orm"
	"gensh.me/VirtualJudge/models/database"
)

func CreateUser(user *auth.User, thirdPart int) {
	u := models.User{}
	err := database.O.QueryTable(models.UserTableName).Filter("third_part_id", user.ThirdPartId).Filter("third_part",thirdPart).One(&u,"id")
	if err == orm.ErrNoRows {
		id, _ := database.O.Insert(&models.User{Name:user.Name, ThirdPartId:user.ThirdPartId,
			Avatar:user.Avatar, Email:user.Email, ThirdPart:thirdPart, Url:user.Url,CreatedAt:time.Now()})
		user.Id = int(id)
	} else if err ==nil {
		user.Id = int(u.Id)
	}
}
