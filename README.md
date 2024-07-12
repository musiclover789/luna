# Luna - 基于视觉的抗指纹爬虫第三方库



		Luna是专为抗指纹自动化爬虫设计的工具，包含抗指纹浏览器和自动化框架，让您能够自由实现所需功能。



作者QQ: 80258153

email:80258153@qq.com


python 版-https://github.com/musiclover789/luna_python

完整文档:https://github.com/musiclover789/luna-browser/tree/main/%E5%A6%82%E4%BD%95%E4%BF%AE%E6%94%B9%E6%8C%87%E7%BA%B9

## Luna 是什么？

- luna是一个自动化框架，类似Selenium Pyppeteer Playwright。满足专业的自动化测试需求。
- luna浏览器、支持浏览器指纹、防关联相关功能的浏览器。





## Luna 有什么不同？

- 你可以使用luna框架，结合  **luna浏览器**  实现模拟浏览器指纹的能力，从而达到防关联测试的效果。





## 普通框架拥有的能力，luna框架也有吗？

- 基本上都有、包括不限于

- 打开浏览器、访问页面、获取页面网页内容

- css、xpath选择器、视觉选择器、鼠标点击、鼠标移动轨迹移动、键盘输入、等

- http、https、socks5、百名单、或者用户名密码方式代理IP 所有格式、所有类型均支持。

- 执行js

- cookie

- 数据包采集等



## 效果演示



![效果展示-加载可能有些慢](https://i.ibb.co/yPkZLd0/mnggiflab-compressed-20231026-215253-min.gif)

## 抗指纹效果演示

你不但可以模拟pc、还可以模拟手机。

![效果展示-加载可能有些慢](https://i.ibb.co/nftHyHW/511714127971-pic.jpg)

![效果展示-加载可能有些慢](https://i.ibb.co/hCXrxn2/BEE68478001-EBDF49-A93-FA7-CBC7-C60-FD.png)


经过大量测试，目前基本可以过掉主流抗指纹识别,提供一些指纹检测网址、仅供参考;



```
https://www.browserscan.net/
https://uutool.cn/browser/
https://abrahamjuliot.github.io/creepjs/
https://browserleaks.com/
https://bot.sannysoft.com/
```



## 使用限制

1、目前仅支持 Windows   x86-64 ，其他平台测试尚不充分。



##  Luna文档部分



详细的使用说明和示例代码，请查看本项目的[文档](https://github.com/musiclover789/luna-browser)。

示例代码部分也可以查看源码的test_case包下内容。




## Luna浏览器部分

目前，我们已经将浏览器文件上传到 百度 网盘，并提供了下载链接：

116版本

链接: https://pan.baidu.com/s/1Cy3g6BcAgzSEFT-ATUtq3g 提取码: 8a9o

125版本

链接: https://pan.baidu.com/s/1DH0Thiq42UTFCFZnxoE39Q 提取码: jime 



<没有授权文件的用户,仅可以测试useragent指纹部分,其他指纹不会生效,付费获取授权文件，联系作者获取>
<仅框架部分是免费、开源的，浏览器部分不开源也不免费，如果您介意收费，请使用其他项目>




## 目前支持指纹项：



**注意**：您必须需要下载和使用luna浏览器，才能使在框架中设置的指纹生效。
如果你即便知道可以修改，但是不知道改成什么样子的指纹，建议直接咨询作者本人。



|      | 指纹项                                                   |
| ---- | -------------------------------------------------------- |
|      | user_agent指纹                                           |
|      | canvas指纹                                               |
|      | webgl指纹\webgpu                                                |
|      | platform平台                                             |
|      | timezone时区                                             |
|      | timezone_offset时区偏移量                                |
|      | languages语言                                            |
|      | userAgentData、全版本号、内核类型等                      |
|      | platform                                                 |
|      | header 修改                                              |
|      | deviceMemory                                             |
|      | hardwareConcurrency                                      |
|      | UNMASKED_VENDOR_WEBGL                                    |
|      | UNMASKED_RENDERER_WEBGL                                  |
|      | GL_VERSION                                               |
|      | GL_SupportedExtensions                                   |
|      | GL_VENDOR                                                |
|      | GL_RENDERER                                              |
|      | GL_SHADING_LANGUAGE_VERSION                              |
|      | 是否webdriver                                            |
|      | 是否brave                                                |
|      | 是否selenium                                             |
|      | 是否来自于真实键盘                                       |
|      | 是否来自于真实鼠标                                       |
|      | 鼠标移动轨迹                                             |
|      | 键盘拼音输入法模拟输入                                   |
|      | cdp检测                                                  |
|      | webRTC 公网ip4、局域网ip6                                |
|      | screen、屏幕尺寸、分辨率、色彩深度、devicePixelRatio等。 |
|      | 声卡指纹                                                 |
|      | 字体列表                                                 |
|      | 触控支持                                                 |
|      | 电池电量等                                               |








## 快速入门

引入包:     go get -u github.com/musiclover789/luna

###### 如果您执行到这一步遇到依赖包问题，那么请执行 go get -v -d ./... 命令即可。

```bash
package main

import (
    "fmt"
    "github.com/musiclover789/luna/devtools"
    "time"
)

func main() {
    // 初始化浏览器对象
    //你浏览器的地址
    //chromiumPath := "/Users/你自己的浏览器的地址/Chromium.app/Contents/MacOS/Chromium"
    chromiumPath := "C:\\src\\chromedev\\chromium\\src\\out\\Default/chrome.exe"
    _, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
        // 设置非隐身模式
        Headless: false,
    })
    // 打开一个页面  如果你想看更多示例 请参考文档 https://github.com/musiclover789/luna-browser
    browserObj.OpenPage("https://www.baidu.com")
    fmt.Println("恭喜你，非常nice的第一个案例")
    time.Sleep(1 * time.Hour)

}

```



**增加难度-等待页面加载-选择器示例**

```
package test_case

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/input"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"time"
)

func main() {
	//Please replace this with your own browser path.
	chromiumPath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"


	err, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/Documents/workspace/golang/cache"),
		Headless:  false,
		ProxyStr:  "https://username:password@42.179.57.60:39349",
		WindowSize: &devtools.WindowSize{
			Width:  1496,
			Height: 967,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	//case 1
	err, pageObj := browserObj.OpenPage("https://www.baidu.com")
	err, x, y := pageObj.GetElementPositionByXpathOnPage("your xpath selector")
	pageObj.SimulateMouseClickOnPage(x, y)

	//case 2
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err, x, y := pageObj.GetElementPositionByCssOnPage("your css selector")
		if err == nil {

			pageObj.SimulateMouseMoveToTarget(x, y)
			pageObj.SimulateMouseClickOnPage(x, y)
			pageObj.SimulateKeyboardInputOnPage("your text")
			pageObj.SimulateEnterKey()
			//pageObj.SimulateBackspaceKey()
		}
	}
	//case 3
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err, node := pageObj.GetElementByXpath("your xpath selector")
		if err == nil {
			fmt.Println(node.XPathSelector)
			fmt.Println(node.CSSSelector)
			fmt.Println(node.TextContent)
			fmt.Println(node.HTMLContent)
			pageObj.GetElementByXpath(node.XPathSelector)
			pageObj.GetAllChildElementByXpah(node.XPathSelector)
			pageObj.GetNextSiblingElementByXpath(node.XPathSelector)
			pageObj.GetFirstChildElementByXpath(node.XPathSelector)
			pageObj.GetLastChildElementByXpah(node.XPathSelector)
			pageObj.GetPreviousSiblingElementByXpath(node.XPathSelector)
		}
	}
	//case 4
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err, node := pageObj.GetElementByXpath("your xpath selector")
		if err == nil {
			fmt.Println(node.XPathSelector)
			err, node = pageObj.GetElementByCss(node.CSSSelector)
			if err == nil {
				err, node = pageObj.GetFirstChildElementByXpath(node.CSSSelector)
				err, x, y := pageObj.GetElementPositionByCssOnPage(node.CSSSelector)
				if err == nil {
					pageObj.SimulateScrollToElementBySelector(node.CSSSelector)
					pageObj.SimulateMouseScrollOnPage(x, y, 100, input.DOWN)
				}
			}
		}
	}

	time.Sleep(1 * time.Hour)
}

```



**设置指纹项 示例**--请注意，这里只是示例，你需要改成正确的指纹，这里只是展示在哪里改。具体指纹说明，参考完整文档。



```
package test_case

import (
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"time"
)

func main() {
	//Please replace this with your own browser path.
	chromiumPath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	

	err, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/Documents/workspace/golang/cache"),
		Headless:  false,
		Fingerprint: []string{
			"--luna_platform=Win32",
			"--luna_audio_random_int_number=981",
			"--luna_cavans_random_int_number=99981",
			"--luna_deviceMemory=8",
			"--luna_hardwareConcurrency=16",
			"--luna_devicePixelRatio=3",
			"--luna_header_set=true"
			"--luna_header_1=accept-language-lunareplace-en;q=0.9",
			"--luna_header_2=sec-ch-ua-arch-lunaremove-",
			"--luna_language=zh-CN",
			"--luna_languages=zh-CN",
			"--luna_timezone=Europe/London",
			"--luna_timezone_offset=3600000",
			"--luna_webrtc_public_ip=10.29.120.2",
			"--luna_webrtc_local_ip6_ip=0f0d8599-9999-4130-87ad-ec008a1c8d63.local",
			"--luna_user_agent=Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
			"--luna_userAgentData=Google Chrome:92-luna-Chromium:92-luna-Not-A.Brand:24-luna-platform:win32-luna-mobile:false-luna-platform_version:6.1-luna-ua_full_version:92.0.4515.186-luna-model:PC-luna-architecture:x86_64",
			"--luna_screen=height:803,width:360,availHeight:803,availWidth:360,availLeft:0,availTop:0,colorDepth:24,pixelDepth:24",
		},
		ProxyStr: "https://username:password@42.179.160.60:39349",
		WindowSize: &devtools.WindowSize{
			Width:  1496,
			Height: 967,
		},
	})
	if err == nil {
		browserObj.OpenPage("https://www.baidu.com")
	}

	time.Sleep(1 * time.Hour)
}

```



**其他 示例**-这些是多个子示例，你需要根据你自己的场景来搭配，这里仅仅是展示。



```
package test_case

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/emulation"
	"github.com/musiclover789/luna/base_devtools/network"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/devtools"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"github.com/tidwall/gjson"
	"sync"
	"time"
)

func main() {
	luna_utils.KillProcess()
	//Please replace this with your own browser path.
	chromiumPath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"


	err, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/Documents/workspace/golang/cache"),
		Headless:  false,
		Fingerprint: []string{
			"--touch-events",
		},
		ProxyStr: "https://username:password@42.179.160.60:39349",
		WindowSize: &devtools.WindowSize{
			Width:  1496,
			Height: 967,
		},
	})

	if err == nil {
		//case 1
		var wg sync.WaitGroup
		wg.Add(1)
		err, pagePro := browserObj.OpenPageAndListen("https://www.baidu.com", func(Session *protocol.Session) {
			page.PageEnable(Session)
			Session.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
				fmt.Println("Waiting for the page to fully load")
				wg.Done()
			})
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		wg.Wait()
		pagePro.RunJS("location.reload()")
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			err, result := pagePro.RunJSSync("your js", time.Minute)
			if err == nil {
				fmt.Println(result, err)
				break
			}
		}
		//case 2
		emulation.SetTouchEmulationEnabled(pagePro.Session, 5)
		var wg1 sync.WaitGroup
		wg1.Add(1)
		page.PageEnable(pagePro.Session)
		pagePro.Session.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
			fmt.Println("Waiting for the page to fully load")
			wg1.Done()
		})
		//case 3
		pagePro.Touch(100, 100)
		pagePro.TouchDrag(100, 100, 100, 200)
		pagePro.TouchTouchEvent(100, 100, "touchMove")
		wg1.Wait()
		//case 3
		_, p1 := browserObj.OpenPageAndListen("https://www.baidu.com/", func(devToolsConn *protocol.Session) {
			network.SetCookieByURL(devToolsConn, "luna_url", "luna-cookie", "https://www.baidu.com")
			network.SetCookie(devToolsConn, "luna_domain", "luna-cookie", "www.baidu.com")
			//network.ClearBrowserCookies(devToolsConn)
		})

		network.SetCookie(p1.Session, "luna_domain_abc", "luna-cookie", "www.baidu.com")

		urls := []string{"https://www.baidu.com"}
		cookies, _ := network.GetCookies(p1.Session, urls)

		for _, result := range gjson.Parse(luna_utils.FormatJSONAsString(cookies)).Get("result.cookies").Array() {
			fmt.Println(result.Get("name").String(), result.Get("value").String(), result.Get("domain").String())
		}
		//case 5
		_, ps := browserObj.GetPages()
		for _, pi := range ps {
			fmt.Println(pi.CurrentURL, pi.Title, pi.PageID)
			browserObj.SwitchPage(pi)
			browserObj.SwitchPageAndListen(pi, func(devToolsConn *protocol.Session) {
			})
		}
		//case 6
		err, p1 = browserObj.OpenPageAndListen("https://www.baidu.com/", func(session *protocol.Session) {
			network.EnableNetwork(session)
			network.RequestResponseAsync(session, func(requestId string, request, response map[string]interface{}) {
				fmt.Println(luna_utils.FormatJSONAsString(request), luna_utils.FormatJSONAsString(request))
				//平时用不上,并不是每个请求都有请求报体；需要根据请求的url自行判断是否需要使用
				network.GetResponseBody(session, requestId, time.Minute)
			})
		})
		//case 7
		fmt.Println(p1.GetHtml())
		fmt.Println(page.DecodeHTMLString(p1.GetHtml().Get("result.outerHTML").String()))
		
		time.Sleep(1 * time.Hour)
		pagePro.Close()
		browserObj.Close()
	}

	time.Sleep(1 * time.Hour)
}

```



##### 相关教程：

| [第一课-常见概念介绍.md](https://github.com/musiclover789/luna-browser/blob/main/第一课-常见概念介绍.md) |
| ------------------------------------------------------------ |
| [第三课-brower对象.md](https://github.com/musiclover789/luna-browser/blob/main/第三课-brower对象.md) |
| [第二课-第一个小例子.md](https://github.com/musiclover789/luna-browser/blob/main/第二课-第一个小例子.md) |
| [第五课.md](https://github.com/musiclover789/luna-browser/blob/main/第五课.md) |
| [第四课page对象.md](https://github.com/musiclover789/luna-browser/blob/main/第四课page对象.md) |



##### 框架相关文档

|                                                              |
| ------------------------------------------------------------ |
| [框架文档.md](https://github.com/musiclover789/luna-browser/tree/main/luna%E6%A1%86%E6%9E%B6-golang%E7%89%88) |
| [如何设置指纹.md](https://github.com/musiclover789/luna-browser/tree/main/%E5%A6%82%E4%BD%95%E4%BF%AE%E6%94%B9%E6%8C%87%E7%BA%B9) |

完整文档参考:  https://github.com/musiclover789/luna-browser

备注: 具体指纹修改项，请参阅上面表格部分。

-----





----------------------



#### 常见问题回复



1、**可以自己随便修改指纹吗？**

​     是的、理论上无限指纹;



2、**目前支持Linux 系统吗？**

​     暂时不支持



3、**为什么我测试基于视觉时候发现，出现bug**

​     如果需要使用视觉方式、请下载原代码的方式，并且注意项目的名称叫luna

​     不要更改因为他是会找寻luna这个文件夹来定位视觉的程序。



4、**我用了这个框架就可以换指纹吗？**

​	您需要下载并使用luna浏览器，结合框架才可以，单独的框架是不能达到效果的。



5、**为什么我测试browserscan时候并不是100%**

​     说明你设置的指纹有问题，或者您根本就不是授权用户，指纹设置不生效。











##### 免责声明：



请在使用本框架之前仔细阅读并理解以下内容。本框架仅用于合法目的，并且作者不承担任何因非法或滥用本框架而导致的责任或后果。通过使用本框架，您同意自行承担风险，并对使用本框架的后果负全部责任。

1. 合法使用：本框架旨在为用户提供便利和支持，并帮助用户完成特定的任务。用户应确保在使用本框架时遵守所有适用的法律、法规和政策。禁止将本框架用于非法目的，包括但不限于侵犯他人隐私、违反知识产权、进行网络攻击等行为。
2. 自担风险：使用本框架的风险完全由用户自行承担。作者不对因使用本框架而导致的任何直接或间接损失或后果承担责任，包括但不限于数据损失、设备故障、业务中断或其他经济损失。
3. 免责声明的范围：本免责声明适用于本框架的所有功能和服务，无论是明示的还是暗示的。作者不提供任何形式的保证，明示或暗示，包括但不限于适销性、特定用途适用性、安全性和准确性。用户对于本框架的选择和使用应自行审慎考虑并承担相应风险。
4. 第三方链接：本框架可能包含指向第三方网站或资源的链接。这些链接仅作为方便提供，不代表作者对这些网站或资源的认可或控制。用户访问任何第三方链接所造成的风险由用户自行承担。
5. 法律适用：使用本框架和解释本免责声明的所有争议均受到适用法律的管辖。

请在使用本框架之前仔细阅读并理解本免责声明的内容。如果您不同意本免责声明的任何部分，请立即停止使用本框架。

如果您有任何问题或疑虑，请与作者联系。谢谢您的合作和理解！

