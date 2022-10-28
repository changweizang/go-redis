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

// 判断用户是否存在
func IfUserExist(phone string) bool{
	var count int64
	db.Model(&User{}).Where("phone = ?", phone).Count(&count)
	if count == 1 {
		return true
	}
	return false
}

// 添加用户
func AddUser(phone string) error {
	user := User{
		Phone: phone,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return db.Create(&user).Error
}
