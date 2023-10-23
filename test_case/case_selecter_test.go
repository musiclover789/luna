package test_case

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"sync"
	"testing"
	"time"
)

func TestSelector(t *testing.T) {
	luna_utils.KillProcess()
	/********************************/
	chromiumPath := "/Users/xxx/Documents/workspace/golang/Chromium.app/Contents/MacOS/Chromium"
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/xxx/Documents/workspace/golang/cache"),
		Headless:  false,
	})
	//browserObj.DevToolsConn.ShowLogJson(true)
	//browserObj.OpenPage("https://www.baidu.com")
	var wg sync.WaitGroup // 同步等待
	wg.Add(1)             // 增加等待的数量
	err, p1 := browserObj.OpenPageAndListen("https://www.baidu.com/", func(devToolsConn *protocol.DevToolsConn) {
		//第一个处理
		devToolsConn.ShowLog(true)
		page.PageEnable(devToolsConn)
		devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
			wg.Done() // 标记回调函数执行完成
			page.PageDisable(devToolsConn)
		})
	})
	wg.Wait() // 等待回调函数执行完成
	time.Sleep(3 * time.Second)
	err, x, y := p1.GetElementPositionByCssOnPage(`#browser-new-page`)
	/********************************/
	fmt.Println("测试一下", err, x, y)
	p1.SimulateMouseClickOnPage(x, y)
	for _, pi := range browserObj.GetPages() {
		browserObj.SwitchPage(pi)
		fmt.Println(pi.CurrentURL, pi.Title)
		fmt.Println(">>>>>>>>>>>>")
	}
	time.Sleep(1 * time.Hour)
}
