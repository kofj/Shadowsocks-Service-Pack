package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

// User Model Struct
type User struct {
	Id         int
	Username   string `orm:size"size(10)" form:"Username"`
	Password   string `orm:"size(255)" form:"Password"`
	Invite     string `orm:"size(9)" form:"Invite"`
	Email      string `orm:size"size(25)" form:"Email"`
	Usertype   int    `orm:"size(2)"`
	Group      int    `orm:"size(2)"`
	Status     int    `orm:default(1);"size(1)"`
	Lastlogin  int64  `orm:"size(10)"`
	Createtime int64  `orm:"size(10)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:654321@/kofj?charset=utf8", 30)
	orm.RegisterDriver("mysql", orm.DR_MySQL)
}
