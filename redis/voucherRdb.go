package redis

import (
	"log"
	"time"
)

// 获取分布式锁
func TryLockOfOrder(userId string) bool {
	flag, err := rdb.SetNX("order:"+userId, userId, time.Second*120).Result()
	if err != nil {
		log.Println("setnx order failed err:", err)
		return false
	}
	return flag
}

// 删除分布式锁
func DeleteLockOfOrder(userId string) {
	err := rdb.Del("order:" + userId).Err()
	if err != nil {
		log.Println("delete key of  order failed err:", err)
	}
}
