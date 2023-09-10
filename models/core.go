package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(sqlite.Open("~/Downloads/word_database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
	}
}
