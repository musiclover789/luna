package test_case

/***
官方文档
//https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/

import (
	"fmt"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	//启动前先杀死其他的chromium进程;为了防止程序以及停止但是依然在内存中贮存;
	//他会根据你的系统不同,使用命令行的命令进行杀死进程
	luna_utils.KillProcess()

	//初始化浏览器对象
	chromiumPath := "/Users/hongyuji/Documents/workspace/golang/Chromium.app/Contents/MacOS/Chromium"
	_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		//设置缓存目录,
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/hongyuji/Documents/workspace/golang/cache"),
		//设置你的代理IP、他支持所有主流的种类https http socks5 有无密码均支持、白名的模式也支持、
		//"http://46.19.160.60:39349"
		//"http://API1M5T:9BFF4922D11@48.19.160.60:39349"
		//"https://46.19.160.60:39349"
		//"https://API1M5T:9BFF4922D11@48.19.160.60:39349"
		//"socks5://API1M5V:9BF49220D11@42.179.160.60:39349"
		ProxyStr: "http://API1M5TV:9BFF49220D11@42.179.160.60:39349",

		//设置非隐身模式
		Headless: false,
	})

	//打开一个tap
	browserObj.OpenPage("https://www.baidu.com")

	fmt.Println("恭喜你、基本上已经学会了如何设置proxy")

	time.Sleep(1 * time.Hour)
}
