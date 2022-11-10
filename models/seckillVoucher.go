package models

import "time"

type SeckillVoucher struct {
	VoucherId  int       `json:"voucher_id"`
	Stock      int       `json:"stock"`
	CreateTime time.Time `json:"create_time"`
	BeginTime  time.Time `json:"begin_time"`
	EndTime    time.Time `json:"end_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (SeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
