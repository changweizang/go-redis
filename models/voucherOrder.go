package models

import (
	"go-redis/utils"
	"gorm.io/gorm"
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

func SaveVoucherOrder(voucherOrder VoucherOrder, tx *gorm.DB) error {
	return tx.Create(&voucherOrder).Error
}

func QueryCountByUserId(userId int) (int64, error) {
	var count int64
	err := db.Model(&VoucherOrder{}).Where("user_id = ?", userId).Count(&count).Error
	return count, err
}
