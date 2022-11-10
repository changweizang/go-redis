package utils

import "time"

type PhoneBody struct {
	Phone string `form:"phone" json:"phone"`
}

type Login struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type ReqSkillVoucher struct {
	ShopId      int       `json:"shop_id"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"sub_title"`
	Rules       string    `json:"rules"`
	PayValue    int       `json:"pay_value"`
	ActualValue int       `json:"actual_value"`
	Type        int       `json:"type"`
	Stock       int       `json:"stock"`
	BeginTime   time.Time `json:"begin_time"`
	EndTime     time.Time `json:"end_time"`
}
