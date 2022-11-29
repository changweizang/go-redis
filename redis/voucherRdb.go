package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

// lua脚本
var luaScript = redis.NewScript(`
if (ARGV[1] == redis.call('get', KEYS[1])) 
then 
    -- 释放锁
    return redis.call('del', KEYS[1])
end
return 0
`)

// 获取分布式锁
func TryLockOfOrder(userId, uuid string) bool {
	flag, err := rdb.SetNX("order:"+userId, uuid, time.Second*120).Result()
	if err != nil {
		log.Println("setnx order failed err:", err)
		return false
	}
	return flag
}

// 删除分布式锁
func DeleteLockOfOrder(userId, uuid string) {
	// 获取锁中标示
	//result, _ := rdb.Get("order:" + userId).Result()
	//// 判断标示是否一致
	//if result == uuid {
	//	err := rdb.Del("order:" + userId).Err()
	//	if err != nil {
	//		log.Println("delete key of  order failed err:", err)
	//	}
	//}
	key := "order:" + userId
	err := luaScript.Run(rdb, []string{key}, uuid).Err()
	if err != nil {
		log.Println("exec lua script failed err:", err)
	}
}
