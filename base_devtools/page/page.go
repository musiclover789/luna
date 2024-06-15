package page

import (
	"encoding/base64"
	"fmt"
	"github.com/musiclover789/luna/base_devtools/runtime"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"github.com/musiclover789/luna/script"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var PageNavigate = func(conn *protocol.Session, url string) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.navigate",
		"params": map[string]interface{}{
			"url": url,
		},
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

var GetFrameTree = func(conn *protocol.Session, id int) {
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.getFrameTree",
	}
	conn.WriteMessage(req)
}

var PageClose = func(conn *protocol.Session) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.close",
	}
	conn.WriteMessage(req)
}

func BringToFront(conn *protocol.Session) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.bringToFront",
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func GetResourceContent(conn *protocol.Session, frameID string, url string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.getResourceContent",
		"params": map[string]interface{}{
			"frameId": frameID,
			"url":     url,
		},
	}
	conn.WriteMessage(req)
}

func GetNavigationHistory(conn *protocol.Session) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.getNavigationHistory",
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

/*
计算缩放因子
*/
func CalculateScalingFactorSync(conn *protocol.Session, timeout time.Duration) (error, float64) {
	id := luna_utils.IdGen.NextID()
	resultChan := make(chan float64, 1)
	eventHandler := func(param interface{}) {
		value := gjson.Get(luna_utils.FormatJSONAsString(param.(map[string]interface{})), "result.result.value").Float()
		resultChan <- value
	}
	//先订阅后执行
	conn.SubscribePersistentEvent(strconv.Itoa(id), eventHandler)
	runtime.EvaluateById(conn, script.GetScalingFactor(), id)
	defer conn.UnsubscribePersistentEvent(strconv.Itoa(id))
	for {
		select {
		case result := <-resultChan:
			return nil, result
		case <-time.After(timeout * time.Second):
			return fmt.Errorf("CalculateScalingFactor timeout"), 0
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func PageCaptureScreenshotSync(conn *protocol.Session, timeout time.Duration) ([]byte, error) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.captureScreenshot",
		"params": map[string]interface{}{
			"format":      "jpeg",
			"quality":     30,
			"fromSurface": false,
		},
	}
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	eventHandler := func(param interface{}) {
		if dataMap, ok := param.(map[string]interface{}); ok {
			if result, ok := dataMap["result"].(map[string]interface{}); ok {
				if data, ok := result["data"].(string); ok {
					decodedData, err := base64.StdEncoding.DecodeString(data)
					if err != nil {
						errChan <- err
						return
					}
					resultChan <- decodedData
					return
				}
			}
		}
	}
	conn.SubscribePersistentEvent(strconv.Itoa(id), eventHandler)
	conn.WriteMessage(req)
	defer conn.UnsubscribePersistentEvent(strconv.Itoa(id))
	for {
		select {
		case result := <-resultChan:
			return result, nil
		case err := <-errChan:
			return nil, err
		case <-time.After(timeout):
			return nil, fmt.Errorf("PageCaptureScreenshot timeout")
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func pageCaptureScreenshotRetrySync(conn *protocol.Session, timeout time.Duration) ([]byte, error) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Page.captureScreenshot",
		"params": map[string]interface{}{
			"format":  "jpeg",
			"quality": 30,
		},
	}
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	eventHandler := func(param interface{}) {
		if dataMap, ok := param.(map[string]interface{}); ok {
			if result, ok := dataMap["result"].(map[string]interface{}); ok {
				if data, ok := result["data"].(string); ok {
					decodedData, err := base64.StdEncoding.DecodeString(data)
					if err != nil {
						errChan <- err
						return
					}
					resultChan <- decodedData
					return
				}
			}
		}
	}
	conn.SubscribePersistentEvent(strconv.Itoa(id), eventHandler)
	conn.WriteMessage(req)
	defer conn.UnsubscribePersistentEvent(strconv.Itoa(id))
	for {
		select {
		case result := <-resultChan:
			return result, nil
		case err := <-errChan:
			return nil, err
		case <-time.After(timeout):
			return nil, fmt.Errorf("PageCaptureScreenshot timeout")
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func SaveScreenshot(filepath string, data []byte) error {
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		// 处理保存文件错误
		return err
	}
	return nil
}

func CaptureTestRetry(conn *protocol.Session, path string, timeout time.Duration) error {
	start := time.Now()
	b, err := pageCaptureScreenshotRetrySync(conn, timeout)
	if err != nil {
		elapsed := time.Since(start).Milliseconds()
		fmt.Printf("截图失败 执行耗时：%d 毫秒\n", elapsed)
		return err
	}
	err = SaveScreenshot(path, b)
	if err != nil {
		elapsed := time.Since(start).Milliseconds()
		fmt.Printf("截图保存失败 执行耗时：%d 毫秒\n", elapsed)
		return err
	}
	//elapsed := time.Since(start).Milliseconds()
	//fmt.Printf("截图成功 执行耗时：%d 毫秒\n", elapsed)
	return nil
}

func CaptureTest(conn *protocol.Session, path string, timeout time.Duration) error {
	start := time.Now()
	b, err := PageCaptureScreenshotSync(conn, timeout)
	if err != nil {
		elapsed := time.Since(start).Milliseconds()
		fmt.Printf("截图失败 执行耗时：%d 毫秒\n", elapsed)
		return err
	}
	err = SaveScreenshot(path, b)
	if err != nil {
		elapsed := time.Since(start).Milliseconds()
		fmt.Printf("截图保存失败 执行耗时：%d 毫秒\n", elapsed)
		return err
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("截图成功 执行耗时：%d 毫秒\n", elapsed)
	return nil
}

func ClosePage(conn *protocol.Session) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.close",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func ReloadPage(conn *protocol.Session, ignoreCache bool, scriptToEvaluateOnLoad string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.reload",
		"params": map[string]interface{}{
			"ignoreCache":            ignoreCache,
			"scriptToEvaluateOnLoad": scriptToEvaluateOnLoad,
		},
	}
	conn.WriteMessage(req)
}

func PageEnable(conn *protocol.Session) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.enable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func PageDisable(conn *protocol.Session) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.disable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func GetLayoutMetrics(conn *protocol.Session) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Page.getLayoutMetrics",
		"params": map[string]interface{}{},
	}

	conn.WriteMessage(req)
}

func DecodeHTMLString(input string) string {
	// 查找字符串中的Unicode转义序列
	start := strings.Index(input, `\u`)
	for start != -1 {
		end := start + 6
		if end <= len(input) {
			// 解析Unicode转义序列
			unicodeStr := input[start+2 : end]
			unicodeInt, err := strconv.ParseInt(unicodeStr, 16, 32)
			if err == nil {
				unicodeChar := string(rune(unicodeInt))
				// 替换转义序列为对应的Unicode字符
				input = strings.Replace(input, `\u`+unicodeStr, unicodeChar, 1)
			}
		}
		// 继续查找下一个Unicode转义序列
		start = strings.Index(input, `\u`)
	}
	input = strings.ReplaceAll(input, `\`, ``)
	return input
}
