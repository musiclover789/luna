package luna_utils

import (
	"math/rand"
	"time"
)

func RandomInRange(min, max float64) float64 {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

