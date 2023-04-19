package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 定义全局变量, 下面的初始化函数不能用  :=  不然会无法赋值
var dberr error

func Init() {
	DB, dberr = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if dberr != nil {
		panic(dberr)
	}
	DB.AutoMigrate(&User{}, &Menu{})

}
