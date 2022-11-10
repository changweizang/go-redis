package models

import "time"

type Voucher struct {
	Id          int       `json:"id"`
	ShopId      int       `json:"shop_id"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"sub_title"`
	Rules       string    `json:"rules"`
	PayValue    int       `json:"pay_value"`
	ActualValue int       `json:"actual_value"`
	Type        int       `json:"type"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (Voucher)TableName() string {
	return "tb_voucher"
}
