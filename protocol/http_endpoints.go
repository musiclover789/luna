package protocol

import (
	"fmt"
	"github.com/musiclover789/luna/log"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

type EndpointResult struct {
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	URL                  string `json:"url"`
	Type                 string `json:"type"`
	Title                string `json:"title"`
	ID                   string `json:"id"`
}

func GetDefaultWebSocketDebuggerUrl(port int) *string {
	for _, item := range *parseEndpoints(httpEndpoints(port, "/json/version")) {
		return &item.WebSocketDebuggerURL
	}
	return nil
}

func GetPageEndpoints(port int) *[]EndpointResult {
	res := make([]EndpointResult, 0)
	for _, obj := range *parseEndpoints(httpEndpoints(port, "/json/list")) {
		if obj.Type == "page" {
			res = append(res, obj)
		}
	}
	return &res
}

func GetPageEndpointByID(port int, id string) *EndpointResult {
	for _, obj := range *parseEndpoints(httpEndpoints(port, "/json/list")) {
		if obj.ID == id {
			return &obj
		}
	}
	return nil
}

func ClosePageEndpoint(port int, targetId string) {
	httpEndpoints(port, "/json/close/"+targetId)
}

func httpEndpoints(port int, path string) string {
	maxTries := 30
	tryInterval := time.Second
	var err error
	var result string

	for try := 0; try < maxTries; try++ {
		var url string
		if try%2 == 0 {
			url = fmt.Sprintf("http://localhost:%d%s", port, path)
		} else {
			url = fmt.Sprintf("http://127.0.0.1:%d%s", port, path)
		}
		result, err = httpGet(url)
		if err != nil {
			time.Sleep(tryInterval)
			continue
		}
		// 成功获取到结果，退出循环
		break
	}
	if err != nil {
		luna_log.Logf("获取WebSocketDebuggerUrl失败，错误信息:%v", err)
		luna_log.Logf("Failed to retrieve WebSocketDebuggerUrl. Error message:%v", err)
	}
	return result
}

func parseEndpoints(jsonStr string) *[]EndpointResult {
	// 解析JSON
	result := gjson.Parse(jsonStr)
	// 获取数组
	array := result.Array()
	// 创建返回结果的数组
	res := make([]EndpointResult, len(array))
	// 遍历每个对象，提取所需的属性
	for i, obj := range array {
		res[i] = EndpointResult{
			WebSocketDebuggerURL: obj.Get("webSocketDebuggerUrl").String(),
			URL:                  obj.Get("url").String(),
			Type:                 obj.Get("type").String(),
			Title:                obj.Get("title").String(),
			ID:                   obj.Get("id").String(),
		}
	}
	return &res
}

func httpGet(url string) (string, error) {
	// 创建一个http.Client，设置请求头
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	// 发送 GET 请求
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// 读取响应的内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
