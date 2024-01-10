package protocol

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"luna/log"
	"luna/luna_utils"
	"math"
	"strconv"
	"time"
)

type DevToolsConn struct {
	conn            *websocket.Conn
	inputChan       chan []byte       //往inputChan 写入命令
	outputChan      chan []byte       //从websocket 读取内容
	funQueue        *DevToolsQueueMap //里面存事件和函数的
	funEternalQueue *DevToolsQueueMap //里面存事件和函数的
	queue           *Queue
	showLog         bool
	showLogJson     bool
	done            chan bool //代表结束了
	doneNum         chan int  //代表结束了
	ID              string
	Speed           int64
}

func CreteDevToolsConn(devtoolsURL string) (error, *DevToolsConn) {
	dialer := websocket.DefaultDialer
	dialer.WriteBufferSize =  math.MaxInt32 // 设置写入缓冲区大小为 1024 字节
	//dialer.ReadBufferSize = 512
	conn, _, err := dialer.Dial(devtoolsURL, nil)

	//conn, _, err := websocket.DefaultDialer.Dial(devtoolsURL, nil)
	if err != nil {
		luna_log.LogError("->failed to connect to WebSocket debugger:", err)
		return err, nil
	}
	d := &DevToolsConn{
		conn:            conn,
		inputChan:       make(chan []byte, 3),
		outputChan:      make(chan []byte, 3),
		funQueue:        NewDevToolsQueueMap(),
		funEternalQueue: NewDevToolsQueueMap(),
		queue:           NewQueue(),
		showLog:         false,
		done:            make(chan bool, 3),
		doneNum:         make(chan int, 3),
	}
	//创建后默认开启.
	d.start()
	return nil, d
}

/***
1、我必须是异步的;
2、你们随便可以往我这边写数据不受任何限制
3、我读取出来的数据都是通知你,你需要有个函数接受，至于你处理不处理那个是你自己的事儿。
4、如果多个进程来调用多个浏览器,我这边要保证互相不影响.
5、你生成对象的时候我就自动开启了.
*/

func (d *DevToolsConn) WriteMessage(req map[string]interface{}) {
	if d.Speed != 0 {
		time.Sleep(time.Millisecond * time.Duration(d.Speed))
	}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		luna_log.LogError("failed to marshal JSON-RPC request:", err)
		luna_log.LogError("无法序列化JSON-RPC请求:", err)
	}
	d.inputChan <- jsonReq
}

func (d *DevToolsConn) readMessageFromChin() {
	for {
		select {
		case item, _ := <-d.done:
			if item {
				d.done <- true
				d.doneNum <- 1
				return
			}
		case message, _ := <-d.outputChan:
			if message != nil {
				// 解析 JSON RPC 响应
				var resp map[string]interface{}
				if err := json.Unmarshal(message, &resp); err != nil {
					luna_log.LogError("failed to unmarshal JSON-RPC response: ", err)
					luna_log.LogError("解析 JSON-RPC 响应失败：", err)
				} else {
					if d.showLog {
						fmt.Println(resp)
					}
					if d.showLogJson {
						fmt.Println(luna_utils.FormatJSONAsString(resp))
					}
					if id := resp["id"]; id != nil {
						switch id := id.(type) {
						case float64:
							d.handleResponse(strconv.FormatFloat(id, 'f', -1, 64), resp)
						case int:
							d.handleResponse(strconv.Itoa(id), resp)
						case string:
							d.handleResponse(id, resp)
						default:
							d.handleResponse(fmt.Sprintf("%v", id), resp)
						}
					} else if method, ok := resp["method"].(string); ok {
						d.handleResponse(method, resp)
					}
				}
			}
		}
	}
}

func (d *DevToolsConn) handleResponse(method string, resp map[string]interface{}) {
	// 执行取出的项中的函数
	for _, fn := range d.funQueue.Dequeue(method) {
		go fn(resp)
		break
	}
	for _, fn := range d.funEternalQueue.Peek(method) {
		go fn(resp)
		break
	}
}

func (d *DevToolsConn) ShowLog(isShow bool) {
	d.showLog = isShow
}
func (d *DevToolsConn) ShowLogJson(isShow bool) {
	d.showLogJson = isShow
}

func (d *DevToolsConn) SubscribeOneTimeEvent(eventName string, handle func(param interface{})) {
	d.funQueue.Enqueue(eventName, map[string]func(param interface{}){eventName: handle})
}

func (d *DevToolsConn) SubscribePersistentEvent(eventName string, handle func(param interface{})) {
	d.funEternalQueue.Enqueue(eventName, map[string]func(param interface{}){eventName: handle})
}
func (d *DevToolsConn) UnsubscribePersistentEvent(eventName string) {
	d.funEternalQueue.Remove(eventName)
}

func (d *DevToolsConn) ReduceSpeed(speed int64) {
	d.Speed = speed
}
func (d *DevToolsConn) GetSpeed() int64 {
	return d.Speed
}
func (d *DevToolsConn) ResetSpeed() {
	d.Speed = 0
}

func (d *DevToolsConn) handleReadMessages() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "-handleReadMessages")
			return
		}
	}()
	for {
		_, message, err := d.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("发现退出信号~")
				d.doneNum <- 1
				d.done <- true
				return
			} else {
				luna_log.LogError("failed to read JSON-RPC response: ", err)
				luna_log.LogError("读取 JSON-RPC 响应失败：", err)
			}
		} else {

			// 检查是否是需要监听的事件
			if message != nil {
				d.outputChan <- message
			}
		}
	}
}

func (d *DevToolsConn) handleWriteMessages() {
	for {
		select {
		case item, _ := <-d.done:
			if item {
				d.done <- true
				d.doneNum <- 1
				return
			}
		case message, _ := <-d.inputChan:
			err := d.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				luna_log.LogError("failed to send JSON-RPC request:", err)
				luna_log.LogError("发送JSON-RPC请求失败:", err)
			}
		}
	}

}

func (d *DevToolsConn) start() {
	go d.readMessageFromChin() //需要释放
	go d.handleWriteMessages() //需要写释放
	go d.handleReadMessages()  //释放

}

func (d *DevToolsConn) Close() error {
	err := d.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		fmt.Println("Close:", err)
	}
	var sum int = 0
	for {
		sum += <-d.doneNum
		if sum == 3 {
			close(d.done)
			close(d.doneNum)
			break
		}
	}
	err = d.conn.Close()
	return err
}
