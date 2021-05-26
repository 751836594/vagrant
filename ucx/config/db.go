package config

import "fmt"

type Database struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

var DatabaseSetting string = ""

func SetUp() {
	db := Database{"52.139.199.133", "readonly", "123456/&", "ucx", "3306"}

	DatabaseSetting = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DbName)
}
