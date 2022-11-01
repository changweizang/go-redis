package models

import (
	"log"
	"time"
)

type Shop struct {
	Id         int
	Name       string
	TypeId     string
	Images     string
	Area       string
	Address    string
	X          float64
	Y          float64
	AvgPrice   int
	Sold       int
	Comments   int
	Score      int
	OpenHours  string
	CreateTime time.Time
	UpdateTime time.Time
}

func SearchShopById(id int) Shop {
	shop := Shop{}
	err := db.Where("id = ?", id).First(&shop).Error
	if err != nil {
		log.Println("select shop failed err:", err)
	}
	return shop
}
