package luna_request

import (
	"fmt"
	"net/http"
	"strings"
)

/****
目的是管理cookie
但是cookie分两种;
我就先实现第一种，后面的再说，因为我估计我暂时也用不到
*/

type Cookie struct {
	cookie HttpMap
}

func NewCookie() *Cookie {
	return &Cookie{}
}

func (receiver *Cookie) ReadCookie(resp *http.Response) {
	if resp.Close{
		fmt.Println("连接已经关闭;")
	}
	for k, v := range resp.Header {
		if strings.EqualFold(k, "Set-Cookie") {
			for _, item := range v {
				cookieValue := ParserOneByGroup("[\\S][^=]*?=[^;]*?;", item, 0)
				//进一步过滤key，为了解决同样的key只取值第一个;
				cookieKey := ParserOneByGroup("([\\S][^=]*?)=", cookieValue, 1)

				//把cookie 放入map中
				receiver.cookie.Push(cookieKey, cookieValue)
			}
		}
	}
}

func (receiver *Cookie) PrintHeader(resp *http.Response) {
	for k, v := range resp.Header {
		for _, item := range v {
			println(k, ">>>", item)
		}
	}
}

func (receiver *Cookie) PrintCookie() {
	for _, v := range receiver.cookie.keys {
		println()
		fmt.Println(v, receiver.cookie.Get(v))
		println()
	}
}

func (receiver *Cookie) ToString() string {
	sb := strings.Builder{}
	for _, v := range receiver.cookie.keys {
		sb.WriteString(receiver.cookie.Get(v).(string))
	}
	return sb.String()
}

func (receiver *Cookie) SetCookie(key interface{}, value interface{}) {
	receiver.cookie.Push(key, value)
}

func (receiver *Cookie) SetCookieJustValue(valueV interface{}) {
	value := ParserOneByGroup("[\\S][^=]*?=[^;]*?;", valueV.(string), 0)
	key := ParserOneByGroup("([\\S][^=]*?)=", valueV.(string), 1)
	if !strings.EqualFold(key, "") && strings.EqualFold(value, "") {
		value = ParserOneByGroup("[\\S][^=]*?=[^;]*", valueV.(string), 0)
	}
	receiver.cookie.Push(key, value)
}

func (receiver *Cookie) SetCookieAll(str string) {
	for _, item := range strings.Split(str, ";") {
		cookieValue := ParserOneByGroup("[\\S][^=]*?=[^;]*", item, 0)
		//进一步过滤key，为了解决同样的key只取值第一个;
		cookieKey := ParserOneByGroup("([\\S][^=]*)=", item, 1)
		cookieKey=strings.TrimSpace(cookieKey)
		//把cookie 放入map中
		//println(">"+cookieKey+"<")
		if len(cookieKey)>0{
			receiver.cookie.Push(cookieKey, cookieValue+`;`)
		}
	}
}
func (receiver *Cookie) SetCookieAllAndReturn(str string) *Cookie {
	for _, item := range strings.Split(str, ";") {
		cookieValue := ParserOneByGroup("[\\S][^=]*?=[^;]*", item, 0)
		//进一步过滤key，为了解决同样的key只取值第一个;
		cookieKey := ParserOneByGroup("([\\S][^=]*)=", item, 1)
		cookieKey=strings.TrimSpace(cookieKey)
		//把cookie 放入map中
		//println(">"+cookieKey+"<")
		if len(cookieKey)>0{
			receiver.cookie.Push(cookieKey, cookieValue+`;`)
		}
	}
	return receiver
}

func (receiver *Cookie) WrriteCookie(resp *http.Request) {
	_, values := receiver.cookie.List()
	for _, value := range values {
		resp.Header.Add("cookie", value.(string))
	}
}



func (receiver *Cookie) RemoveCookie(key interface{}) {
	for _, v := range receiver.cookie.keys {
		if ISABSMatch(key.(string), v.(string)) {
			receiver.cookie.Remove(v)
		}
	}
}

func (receiver *Cookie) ClearCookie() {
	for _, v := range receiver.cookie.keys {
		receiver.cookie.Remove(v)
	}
}

func (receiver *Cookie) JustNeedCookie(keys ...interface{}) {
	for _, v := range receiver.cookie.keys {
		b := true
		for _, item := range keys {
			if strings.EqualFold(item.(string), v.(string)) {
				b = false
				break
			}
		}
		if b {
			receiver.cookie.Remove(v)
		}
	}
	keysT, _ := receiver.cookie.List()
	for _, v := range keysT {
		b := true
		for _, item := range keys {
			if strings.EqualFold(item.(string), v.(string)) {
				b = false
				break
			}
		}
		if b {
			receiver.cookie.Remove(v)
		}
	}
}

func (receiver *Cookie) GetCookie(key interface{}) string {
	if receiver.cookie.Get(key)==nil{
		return ""
	}
	return receiver.cookie.Get(key).(string)
}

