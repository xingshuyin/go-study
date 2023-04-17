package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 定义全局变量

func Init() *gorm.DB {
	DB, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}
