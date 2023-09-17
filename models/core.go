package models

import (
	"dictService/global"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDb() {
	DB, err = gorm.Open(sqlite.Open(global.Global.Config.Sqlite.Path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
	}
}
