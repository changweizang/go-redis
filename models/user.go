package models

import "time"

type User struct {
	Id int
	Phone string
	Password string
	NickName string
	Icon string
	CreateTime time.Time
	UpdateTime time.Time
}

func (User) TableName() string {
	return "tb_user"
}
