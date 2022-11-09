package redis

import (
	"time"
)

// 编辑此处时的时间戳
var BEGIN_TIME_STAMP int64 = 1667978585

// 生成全局唯一ID
// keyPrefix：任务前缀
func NextId(keyPrefix string) int64 {
	// 1.生成时间戳
	timeStamp := time.Now().Unix() - BEGIN_TIME_STAMP
	// 2.生成序列号
	// 2.1.获取当前日期，精确到天
	date := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	// 2.2.自增长
	count := IncreId(keyPrefix, date)
	// 3.拼接并返回
	return timeStamp << 32 | count
}
