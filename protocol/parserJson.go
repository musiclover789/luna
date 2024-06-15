package protocol

import (
	"github.com/tidwall/gjson"
)

var ParserWebSocketDebuggerUrl = func(jsonStr string) []map[string]string {
	// 解析JSON
	result := gjson.Parse(jsonStr)
	// 获取数组
	array := result.Array()
	// 创建返回结果的数组
	res := make([]map[string]string, len(array))
	// 遍历每个对象，提取所需的属性
	for i, obj := range array {
		res[i] = map[string]string{
			"webSocketDebuggerUrl": obj.Get("webSocketDebuggerUrl").String(),
			"url":                  obj.Get("url").String(),
			"type":                 obj.Get("type").String(),
			"title":                obj.Get("title").String(),
			"id":                   obj.Get("id").String(),
		}
	}

	return res
}
