package protocol

import (
	"fmt"
	"testing"
)

func Test_http_endpoints_test(t *testing.T) {
	/**
	http://127.0.0.1:63057/json/new?http://www.baidu.com
	http://127.0.0.1:63057/json/version 这个和控制台启动的那个一摸一样
	/json/activate/
	json/close
	*/
	fmt.Println("Test_http_endpoints_test")
	//fmt.Println(httpEndpoints(50625,"/json/version"))
	//fmt.Println(httpEndpoints(50625,"/json/list"))
	//
	//fmt.Println(*GetDefaultWebSocketDebuggerUrl(50625))
	//fmt.Println(*GetPageEndpoints(50625))

}
