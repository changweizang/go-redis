package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成验证码
func GetRandomCode() string {
	return strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)))
}
