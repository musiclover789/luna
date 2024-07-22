package main

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"strconv"
	"sync"
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
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)
		go func() {
			chromiumPath := "/Users/Documents/workspace/ios/chromedev/chromium/src/out/Default-test/Chromium.app/Contents/MacOS/Chromium"
			_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
				CachePath: luna_utils.CreateCacheDirInSubDir("/Users/hongyuji/Documents/workspace/golang/cache"),

				Fingerprint: []string{
					"--luna_webrtc_public_ip=101.29.120." + strconv.Itoa(i),
				},
				Headless: false,
				WindowSize: &devtools.WindowSize{
					Width:  1496,
					Height: 967,
				},
			})

			var wg sync.WaitGroup
			wg.Add(1)
			browserObj.OpenPageAndListen("https://abrahamjuliot.github.io/creepjs/", func(devToolsConn *protocol.Session) {
				devToolsConn.ShowLog(false)
				page.PageEnable(devToolsConn)
				devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
					fmt.Println("load ok")
					wg.Done()
				})
			})
			wg.Wait()
			browserObj.OpenPage("https://www.browserscan.net/")
			time.Sleep(time.Hour)
		}()
	}

}
