package main

import (
	"fmt"
	"time"
)

func getTimeZoneOffset(timeZoneStr string) (string, error) {
	// 获取指定时区的信息
	loc, err := time.LoadLocation(timeZoneStr)
	if err != nil {
		return "", err
	}

	// 获取当前时间
	currentTime := time.Now().In(loc)

	// 计算时间偏移
	_, offset := currentTime.Zone()
	offsetHours := offset / 3600

	// 将偏移转换为字符串
	offsetStr := fmt.Sprintf("%+03d", offsetHours*60*60*1000)

	return offsetStr, nil
}

func main() {
	timeZoneStr := "Europe/London" // 替换为您的时区字符串
	offset, err := getTimeZoneOffset(timeZoneStr)
	if err != nil {
		fmt.Println("无法加载时区:", err)
	} else {
		fmt.Printf("时区偏移: %s 小时\n", offset)
	}
}
