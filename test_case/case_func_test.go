package test_case

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/network"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/base_devtools/runtime"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"sync"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
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
	err, p1 := browserObj.OpenPageAndListen("https://www.baidu.com/", func(devToolsConn *protocol.DevToolsConn) {
		//第一个处理
		devToolsConn.ShowLogJson(true)
		page.PageEnable(devToolsConn)
		network.EnableNetwork(devToolsConn)
		network.RequestResponseAsync(devToolsConn, func(requestId string, request, response map[string]interface{}) {
			fmt.Println(luna_utils.FormatJSONAsString(request), luna_utils.FormatJSONAsString(request))
			network.GetResponseBody(devToolsConn, requestId, time.Minute)
		})
		devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
			wg.Done() // 标记回调函数执行完成
			runtime.Evaluate(devToolsConn, " your js ")
			runtime.EvaluateWithResultSync(devToolsConn, " your js ", time.Minute)
			page.PageDisable(devToolsConn)
		})
	})
	p1.RunJS(" your js ")
	p1.RunJSSync(" your js ", time.Minute)
	wg.Wait() // 等待回调函数执行完成
	fmt.Println(p1.GetHtml())
	fmt.Println(page.DecodeHTMLString(p1.GetHtml().Get("result.outerHTML").String()))

	//我们准备等待一组信号的出现、但是过程中有可能出现其他的信号,所以我们需要单独的处理
	//p1.WaitForMatchOnPageSync()
	//---------
	//browserObj.Close()
	fmt.Println(err, p1)
	time.Sleep(1 * time.Hour)
}
