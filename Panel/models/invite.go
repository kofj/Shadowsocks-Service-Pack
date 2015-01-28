package models

import (
	"fmt"
	_ "github.com/astaxie/beego/session/mysql"
)

// Invite code Model Struct
type Invite struct {
	Id    int
	Code  string `orm:"size(9)"`
	Owner int    `orm:"size(6)"`
	Used  int    `orm:"size(1)"`
}

// find user by name
func FindCode(code string) (Invite, error) {
	invite := Invite{Code: code}
	_, err := Orm.Get(&invite)
	if err != nil {
		fmt.Println(err)
	}

	return invite, err
}
