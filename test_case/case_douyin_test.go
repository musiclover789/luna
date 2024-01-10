package test_case

import (
	"fmt"
	"luna/base_devtools/network"
	"luna/base_devtools/page"
	"luna/devtools"
	"luna/luna_utils"
	"luna/protocol"
	"sync"
	"testing"
	"time"
)

/***
官方文档
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/

func TestDouyin(t *testing.T) {
	//启动前先杀死其他的chromium进程;为了防止程序以及停止但是依然在内存中贮存;
	//他会根据你的系统不同,使用命令行的命令进行杀死进程

	luna_utils.KillProcess()
	/********************************/
	chromiumPath := "/Users/hongyuji/Documents/workspace/golang/Chromium.app/Contents/MacOS/Chromium"
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/hongyuji/Documents/workspace/golang/cache"),
		Headless:  false,
	})

	//第一组获取方式
	var wg sync.WaitGroup // 同步等待
	wg.Add(1)             // 增加等待的数量
	err, p1 := browserObj.OpenPageAndListen("https://live.douyin.com/865250281495", func(devToolsConn *protocol.DevToolsConn) {
		//第一个处理
		devToolsConn.ShowLogJson(false)
		page.PageEnable(devToolsConn)
		network.EnableNetwork(devToolsConn)
		network.RequestResponseAsync(devToolsConn, func(requestId string, request, response map[string]interface{}) {
			fmt.Println(luna_utils.FormatJSONAsString(request),luna_utils.FormatJSONAsString(request))
			//network.GetResponseBody(devToolsConn,requestId,time.Minute)
		})
	})
	wg.Wait() // 等待回调函数执行完成

	fmt.Println(err, p1)
	time.Sleep(1 * time.Hour)
}
