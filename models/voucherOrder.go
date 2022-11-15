package models

import (
	"go-redis/utils"
	"time"
)

type VoucherOrder struct {
	Id         int              `json:"id"`
	UserId     int              `json:"user_id"`
	VoucherId  int              `json:"voucher_id"`
	PayType    int              `json:"pay_type"`
	Status     int              `json:"status"`
	CreateTime time.Time        `json:"create_time"`
	PayTime    *utils.LocalTime `json:"pay_time"`
	UseTime    *utils.LocalTime `json:"use_time"`
	RefundTime *utils.LocalTime `json:"refund_time"`
	UpdateTime time.Time        `json:"update_time"`
}

func (VoucherOrder) TableName() string {
	return "tb_voucher_order"
}

func SaveVoucherOrder(voucherOrder VoucherOrder) error {
	return db.Create(&voucherOrder).Error

}
