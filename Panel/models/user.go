package models

import (
	"fmt"
	_ "github.com/astaxie/beego/session/mysql"
)

// User Model Struct
type User struct {
	Id            int
	Username      string `orm:size"size(10)" form:"Username"`
	Password      string `orm:"size(255)" form:"Password"`
	Invite        string `orm:"size(9)" form:"Invite"`
	Email         string `orm:size"size(25)" form:"Email"`
	Usertype      int    `orm:"size(2)"`
	Group         int    `orm:"size(2)"`
	Status        int    `orm:default(1);"size(1)"`
	Lastlogintime int64  `orm:"size(10)"`
	Createtime    int64  `orm:"size(10)"`
}

// find user by name
func FindUser(username string) (User, error) {
	user := User{Username: username}
	_, err := Orm.Get(&user)
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

// find user by mail
func FindUserByMail(mail string) (User, error) {
	user := User{Email: mail}
	_, err := Orm.Get(&user)
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}
