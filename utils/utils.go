package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成验证码
func GetRandomCode() string {
	return fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
