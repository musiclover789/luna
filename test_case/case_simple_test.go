package test_case

import (
	"fmt"
	"github.com/musiclover789/luna/devtools"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	//初始化浏览器对象
	chromiumPath := "C:\\src\\chromedev\\chromium\\src\\out\\Default/chrome.exe"
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		//设置非隐身模式
		Headless: false,
	})
	//打开一个tap
	browserObj.OpenPage("https://www.baidu.com")
	fmt.Println("恭喜你、非常nice的第一个案例")
	time.Sleep(1 * time.Hour)
}
