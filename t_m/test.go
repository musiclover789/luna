package main

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/emulation"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"os"
	"time"
)

/*
**
官方文档
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func main() {
	luna_utils.KillProcess()
	/********************************/
	chromiumPath := "/Users/Documents/workspace/ios/chromedev/chromium/src/out/Default-test/Chromium.app/Contents/MacOS/Chromium"
	_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		ProxyStr:    "",
		CachePath:   "",
		Fingerprint: []string{},
		Headless:    false,
		WindowSize: &devtools.WindowSize{
			Width:  1496,
			Height: 967,
		},
	})

	//&{2 1496 967 1496 858}
	browserObj.Session.ShowLog(true)

	_, p1 := browserObj.OpenPageAndListen("https://www.baidu.com", func(session *protocol.Session) {

	})
	browserObj.Close()
	time.Sleep(time.Second * 3)
	luna_utils.KillProcessByPid(browserObj.Pid)
	os.Exit(0)
	time.Sleep(1000000)
	emulation.SetTouchEmulationEnabled(p1.Session, 5)
	p1.SetMaxTouchPoints(5)
	page.PageEnable(p1.Session)
	p1.Session.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
		//页面加载完成
		fmt.Println("页面打开")
	})
	time.Sleep(50 * time.Second)
	_, x, y := p1.GetElementPositionByXpathOnPage("//*[@id=\"kw\"]")
	p1.SimulateMouseClickOnPage(x, y)
	p1.SimulateKeyboardInputOnPage("1k张来24")
	//模拟回车键
	p1.SimulateEnterKey()
	//模拟退格键
	p1.SimulateBackspaceKey()
	browserObj.Close()
}
