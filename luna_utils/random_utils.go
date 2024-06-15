package luna_utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomInRange(min, max float64) float64 {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int) int {
	// 使用当前时间的纳秒作为随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机整数
	return rand.Intn(max-min+1) + min
}

// CalculateMD5 计算字符串的 MD5 值
func CalculateMD5(input string) string {
	// 将字符串转换为字节数组
	data := []byte(input)

	// 计算 MD5 值
	hash := md5.Sum(data)

	// 将 MD5 值转换为十六进制字符串
	md5Str := hex.EncodeToString(hash[:])

	return md5Str
}

func RandomString(length int) string {
	// 定义字符集
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机字符串
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
