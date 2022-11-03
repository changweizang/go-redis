package utils

import "go-redis/models"

type PhoneBody struct {
	Phone string `form:"phone" json:"phone"`
}

type Login struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type ShopReq struct {
	Shop models.Shop `json:"shop"`
}
