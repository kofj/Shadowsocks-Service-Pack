package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Orm *xorm.Engine

func init() {
	var err error
	Orm, err = xorm.NewEngine("mysql", beego.AppConfig.String("mysql"))
	if err != nil {
		fmt.Println(err)
	}
	//Orm.ShowWarn = true
	//Orm.ShowSQL = true
	err = Orm.Sync2(new(User), new(Invite))
	if err != nil && beego.RunMode == "dev" {
		fmt.Println(err)
	}
}
