package test_case

import (
	"fmt"
	"luna/base_devtools/input"
	"luna/base_devtools/page"
	"luna/base_devtools/runtime"
	"luna/devtools"
	"luna/luna_utils"
	"luna/protocol"
	"luna/script"
	"testing"
	"time"
)

/***
官方文档
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/

func TestBaidu(t *testing.T) {
	luna_utils.KillProcess()
	/********************************/
	//chromiumPath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	chromiumPath := "C:\\src\\chromedev\\chromium\\src\\out\\Default/chrome.exe"
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("C:\\workspace\\tempcatch"),
		ImgPath:   "C:\\workspace\\v2\\v2\\luna/test_img/baidu_img",
		Headless:  false,
		//ProxyStr:  "https://API1M5TV:9BFF49220D11@42.179.160.60:39349",
		//Fingerprint: []string{"fingerprint1", "fingerprint2"},
		WindowSize: &devtools.WindowSize{
			Width:  1496,
			Height: 967,
		},
	})
	//&{2 1496 967 1496 858}
	fmt.Println()
	browserObj.DevToolsConn.ShowLog(false)

	err, itemPage := browserObj.OpenPageAndListen("https://www.baidu.com", func(devToolsConn *protocol.DevToolsConn) {
		//第一个处理
		devToolsConn.ShowLog(true)
		page.PageEnable(devToolsConn)
		devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
			runtime.Evaluate(devToolsConn, script.ShowMousePosition())
		})
		devToolsConn.SubscribeOneTimeEvent("Page.windowOpen", func(param interface{}) {
			runtime.Evaluate(devToolsConn, script.ShowMousePosition())
		})
	})
	fmt.Println(itemPage.RunJSSync(" function aa(){ return navigator.languages;} aa();", time.Minute))
	if err == nil {
		itemPage.RunJS(script.ShowMousePosition())
		itemPage.DevToolsConn.ReduceSpeed(10)
		err, ok := itemPage.WaitForMatchOnPageSync("home.png", 0.5, time.Hour)
		if err == nil && ok {
			fmt.Println("说明、页面已经成功打开")
			//time.Sleep(time.Hour)
			err, imageCoordinates := itemPage.SimilarityWithMargin("home.png", 50, 200, 10, 10, time.Minute)
			if err == nil && imageCoordinates.MatchScore > 0.5 {
				targetX, targetY := imageCoordinates.RandomX, imageCoordinates.RandomY
				itemPage.SimulateMouseMoveOnPage(luna_utils.RandomInRange(-1, devtools.BrowserGlobal.ScreenAvailWidth), -1, targetX, targetY)
				itemPage.SimulateMouseClickOnPage(targetX, targetY)
				itemPage.SimulateKeyboardInputOnPage("随便")

				//点击按钮
				err, imageCoordinates := itemPage.ImageSimilarity("button_01.png", time.Minute)
				if err == nil && imageCoordinates.MatchScore > 0.5 {
					itemPage.SimulateMouseMoveOnPage(targetX, targetY, imageCoordinates.RandomX, imageCoordinates.RandomY)
					itemPage.SimulateMouseClickOnPage(imageCoordinates.RandomX, imageCoordinates.RandomY)
				}
				time.Sleep(5 * time.Second)
				//滚动到页面底部
				fmt.Println("~~~~~~~~~~>==============<~~~~~~~~~~~~~~")
				err, ok := itemPage.ScrollToTargetImagePosition(imageCoordinates.RandomX, imageCoordinates.RandomY, input.DOWN, "bottom.png", 0.5, time.Minute)
				fmt.Println(ok, err)

				itemPage.Close()
				browserObj.Close()
			}

		}
		fmt.Println("~~~~~~~~~~>==============<~~~~~~~~~~~~~~")
	} else {
		fmt.Println(err, "错误了")
	}
	/********************************/
	time.Sleep(1 * time.Minute)
}
