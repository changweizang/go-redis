package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 生成验证码
func GetRandomCode() string {
	return fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// 生成uuid
func GetUUid() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}
