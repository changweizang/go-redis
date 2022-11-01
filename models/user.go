package models

import (
	"log"
	"time"
)

type User struct {
	Id         int
	Phone      string
	Password   string
	NickName   string
	Icon       string
	CreateTime time.Time
	UpdateTime time.Time
}

func (User) TableName() string {
	return "tb_user"
}

// 判断用户是否存在
func IfUserExist(phone string) bool {
	var count int64
	db.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count == 1
}

// 根据phone查询用户
func SearchUserByPhone(phone string) User {
	user := User{}
	err := db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		log.Println("select user failed err:", err)
	}
	return user
}

// 添加用户
func AddUser(phone string) error {
	user := User{
		Phone:      phone,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return db.Create(&user).Error
}
