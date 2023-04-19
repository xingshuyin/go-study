package models

import "time"

type Model struct {
	ID         uint      `gorm:"primarykey"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	DeptBelong uint
	Creator    uint
}

type User struct {
	Name     string `gorm:"size:230;"`
	Username string `gorm:"size:230;index;unique;not null;comment:用户名;"`
	Password string `gorm:"size:230;not null;comment:密码;"`
	Age      int    `gorm:"type:int;"`
	Menus    []Menu `gorm:"many2many:user_menu;"`
	Model
}

func (u *User) TableName() string { // 数据库表名
	return "user"
}

type Menu struct {
	Name string `gorm:"size:230;comment:菜单名称;"`
	Model
}

func (u *Menu) TableName() string { // 数据库表名
	return "menu"
}

type Interface struct {
}
