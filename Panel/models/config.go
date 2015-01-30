package models

import (
	"fmt"
	_ "github.com/astaxie/beego/session/mysql"
)

// Key-Value Model Struct
type Config struct {
	Key     string `xorm:"size(9) pk"`
	Value   string `xorm:"size(255)"`
	Comment string `xorm:"size(255)"`
}

// find value by key
func GetConfig(key string) (Config, error) {
	config := Config{Key: key}
	_, err := Orm.Get(&config)
	if err != nil {
		fmt.Println(err)
	}

	return config, err
}

// find value by key
func SetConfig(key, value string, comment ...string) (ok int64, err error) {
	keyMap := Config{Key: key}
	config := Config{Key: key, Value: value}
	if len(comment) != 0 {
		config.Comment = comment[0]
	}
	haskey, _ := Orm.Get(&keyMap)
	if haskey {
		ok, err = Orm.Update(&config)
	} else {
		ok, err = Orm.Insert(&config)
	}
	if err != nil {
		fmt.Println(ok, err)
	}

	return ok, err
}
