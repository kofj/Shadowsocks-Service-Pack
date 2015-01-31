package models

import (
	"fmt"
	_ "github.com/astaxie/beego/session/mysql"
)

// User Model Struct
type User struct {
	Id            int64
	Username      string `xorm:"varchar(10) notnull unique 'username' "`
	Password      string `xorm:"varchar(40)"`
	Invite        string `xorm:"varchar(9) unique 'invite' "`
	Email         string `xorm:"varchar(25) unique 'email' "`
	Status        int    `xorm:"default 0"`
	Lastlogintime int64  `xorm:"int64(10)"`
	Createtime    int64  `xorm:"int64(10)"`
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

// add new user
func AddUser(user User) (User, error) {
	_, err := Orm.InsertOne(&user)
	return user, err
}

// update login time
func LastLoginTime(id, value int64) {
	log := User{Lastlogintime: value}
	_, err := Orm.Id(id).Update(&log)
	if err != nil {
		fmt.Println(err)
	}
}

// Update user status
func UpdateUserStatus(id int64, status int) {
	user := User{Status: status}
	Orm.Id(id).Update(&user)
}
