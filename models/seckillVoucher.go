package models

import (
	"go-redis/utils"
	"time"

	"gorm.io/gorm"
)

type SeckillVoucher struct {
	VoucherId  int              `json:"voucher_id"`
	Stock      int              `json:"stock"`
	CreateTime time.Time        `json:"create_time"`
	BeginTime  *utils.LocalTime `json:"begin_time"`
	EndTime    *utils.LocalTime `json:"end_time"`
	UpdateTime time.Time        `json:"update_time"`
}

func (SeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}

func AddSkillVoucher(skillVoucher SeckillVoucher) error {
	return db.Create(&skillVoucher).Error
}

func QuerySeckillVoucherById(id string) (SeckillVoucher, error) {
	seckillVoucher := SeckillVoucher{}
	err := db.Where("voucher_id = ?", id).First(&seckillVoucher).Error
	return seckillVoucher, err
}

func DecVoucherSock(seckillVoucher SeckillVoucher, tx *gorm.DB) (int64, error) {
	result := tx.Model(&seckillVoucher).
		Where("voucher_id = ? and stock > ?", seckillVoucher.VoucherId, 0).
		UpdateColumn("stock", gorm.Expr("stock - ?", 1))
	return result.RowsAffected, result.Error
}
