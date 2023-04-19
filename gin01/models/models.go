package models

import "time"

type Model struct {
	ID         uint      `gorm:"primarykey"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	Name     string `gorm:"size: 230"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"size: 230;not null"`
	Age      int    `gorm:""`
	Menus    []Menu `gorm:"many2many:user_menu;"`
	Model
}

func (u *User) TableName() string { // 数据库表名
	return "user"
}

type Menu struct {
	Name string `gorm:"size: 230"`
	Model
}

func (u *Menu) TableName() string { // 数据库表名
	return "menu"
}
