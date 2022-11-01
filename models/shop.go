package models

import (
	"log"
	"time"
)

type Shop struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	TypeId     string    `json:"type_id"`
	Images     string    `json:"images"`
	Area       string    `json:"area"`
	Address    string    `json:"address"`
	X          float64   `json:"x"`
	Y          float64   `json:"y"`
	AvgPrice   int       `json:"avg_price"`
	Sold       int       `json:"sold"`
	Comments   int       `json:"comments"`
	Score      int       `json:"score"`
	OpenHours  string    `json:"open_hours"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (Shop) TableName() string {
	return "tb_shop"
}

func SearchShopById(id string) Shop {
	shop := Shop{}
	err := db.Where("id = ?", id).First(&shop).Error
	if err != nil {
		log.Println("select shop failed err:", err)
	}
	return shop
}
