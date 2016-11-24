package main

import (
	_ "gensh.me/VirtualJudge/routers"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	//orm.RunCommand()
	// 打印SQL： main orm sqlall
	//自动建表 main orm syncdb -force=1 -v
	//
}
