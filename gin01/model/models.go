package models

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) TableName() string { // 数据库表名
	return "user"
}
