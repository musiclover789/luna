package network

import (
	"fmt"
	"luna/luna_utils"
	"luna/protocol"
	"strconv"
	"time"
)

/*
**
Network.dataReceived：接收到网络数据时触发的事件。
Network.eventSourceMessageReceived：接收到EventSource消息时触发的事件。
Network.loadingFailed：加载失败时触发的事件，比如网络错误或资源无法访问等。
Network.loadingFinished：加载完成时触发的事件。
Network.requestServedFromCache：从缓存中提供请求资源时触发的事件。
Network.requestWillBeSent：发送请求之前触发的事件。
Network.responseReceived：接收到响应时触发的事件。
Network.webSocketClosed：WebSocket连接关闭时触发的事件。
Network.webSocketCreated：WebSocket连接创建时触发的事件。
Network.webSocketFrameError：WebSocket帧错误时触发的事件。
Network.webSocketFrameReceived：接收到WebSocket帧时触发的事件。
Network.webSocketFrameSent：发送WebSocket帧时触发的事件。
Network.webSocketHandshakeResponseReceived：WebSocket握手响应接收时触发的事件。
Network.webSocketWillSendHandshakeRequest：发送WebSocket握手请求之前触发的事件。
Network.webTransportClosed：WebTransport连接关闭时触发的事件。
Network.webTransportConnectionEstablished：建立WebTransport连接时触发的事件。
Network.webTransportCreated：创建WebTransport时触发的事件。
Network.reportingApiEndpointsChangedForOrigin：Reporting API端点变化时触发的事件（实验性的）。
Network.reportingApiReportAdded：添加Reporting API报告时触发的事件（实验性的）。
Network.reportingApiReportUpdated：更新Reporting API报告时触发的事件（实验性的）。
Network.requestWillBeSentExtraInfo：发送请求之前附加信息时触发的事件（实验性的）。
Network.resourceChangedPriority：资源优先级发生变化时触发的事件（实验性的）。
Network.responseReceivedExtraInfo：接收到响应附加信息时触发的事件（实验性的）。
Network.signedExchangeReceived：接收到Signed Exchange时触发的事件（实验性的）。
Network.subresourceWebBundleInnerResponseError：子资源Web Bundle内部响应错误时触发的事件（实验性的）。
Network.subresourceWebBundleInnerResponseParsed：子资源Web Bundle内部响应解析完成时触发的事件（实验性的）。
Network.subresourceWebBundleMetadataError：子资源Web Bundle元数据错误时触发的事件（实验性的）。
Network.subresourceWebBundleMetadataReceived：接收到子资源Web Bundle元数据时触发的事件（实验性的）。
Network.trustTokenOperationDone：Trust Token操作完成时触发的事件（实验性的）。
Network.requestIntercepted：拦截请求时触发的事件
*/
func EnableNetwork(conn *protocol.DevToolsConn) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Network.enable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func NetworkDisable(conn *protocol.DevToolsConn) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Network.disable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

// 设置cookie 百度举例 https://www.baidu.com/  baidu.com作为domain 即可
func SetCookie(conn *protocol.DevToolsConn, key, value, domain string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Network.setCookie",
		"params": map[string]interface{}{
			"name":   key,
			"value":  value,
			"domain": domain,
		},
	}
	conn.WriteMessage(req)
}

// 通过url设置cookie 百度举例https://www.baidu.com/, https://www.baidu.com作为url即可
func SetCookieByURL(conn *protocol.DevToolsConn, key, value, url string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Network.setCookie",
		"params": map[string]interface{}{
			"name":  key,
			"value": value,
			"url":   url,
		},
	}
	conn.WriteMessage(req)
}

func ClearBrowserCookies(conn *protocol.DevToolsConn) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Network.clearBrowserCookies",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func GetCookies(conn *protocol.DevToolsConn, urls []string) (map[string]interface{}, error) {
	//---
	id := luna_utils.IdGen.NextID()
	resultChan := make(chan map[string]interface{}, 1)
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		if paramMap, ok := param.(map[string]interface{}); ok {
			resultChan <- paramMap
		}
	})
	req := map[string]interface{}{
		"id":     id,
		"method": "Network.getCookies",
		"params": map[string]interface{}{
			"urls": urls,
		},
	}
	conn.WriteMessage(req)
	for {
		select {
		case result := <-resultChan:
			return result, nil
		case <-time.After(time.Second * 1):
			return nil, fmt.Errorf("getCookie timeout")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func GetResponseBody(conn *protocol.DevToolsConn, requestId string, timeout time.Duration) (map[string]interface{}, error) {
	id := luna_utils.IdGen.NextID()
	resultChan := make(chan map[string]interface{}, 1)
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		if paramMap, ok := param.(map[string]interface{}); ok {
			resultChan <- paramMap
		}
	})
	req := map[string]interface{}{
		"id":     id,
		"method": "Network.getResponseBody",
		"params": map[string]interface{}{
			"requestId": requestId,
		},
	}
	conn.WriteMessage(req)
	for {
		select {
		case result := <-resultChan:
			return result, nil
		case <-time.After(timeout):
			return nil, fmt.Errorf("PageCaptureScreenshot timeout")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func RequestResponseAsync(conn *protocol.DevToolsConn, handler func(requestId string, request, response map[string]interface{})) {
	go func() {
		processor := NewDataProcessor()
		conn.SubscribePersistentEvent("Network.requestWillBeSent", func(param interface{}) {
			if paramMap, ok := param.(map[string]interface{}); ok {
				requestID := paramMap["params"].(map[string]interface{})["requestId"].(string)
				request, response := processor.ProcessRequest(requestID, paramMap)
				if request != nil && response != nil {
					handler(requestID, request, response)
				}
			}
		})
		conn.SubscribePersistentEvent("Network.responseReceived", func(param interface{}) {
			if paramMap, ok := param.(map[string]interface{}); ok {
				requestID := paramMap["params"].(map[string]interface{})["requestId"].(string)
				request, response := processor.ProcessResponse(requestID, paramMap)
				if request != nil && response != nil {
					handler(requestID, request, response)
				}
			}
		})
		//conn.DeleteEternalEventRegistration("Network.requestWillBeSent")
		//conn.DeleteEternalEventRegistration("Network.responseReceived")
	}()
}
