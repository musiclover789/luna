

# Luna - 基于视觉的抗指纹爬虫第三方库



		Luna是专为抗指纹自动化爬虫设计的工具，包含抗指纹浏览器和自动化框架，让您能够自由实现所需功能。



作者QQ: 80258153



python版本框架:准备升级原生态版本；所以暂时删除掉基于golang版本封装的python版
任何问题直接咨询作者本人，不在设置QQ技术交流群。



## 功能亮点

- 强大的抗指纹技术

- 视觉特征解析

- 简单易用的接口

- 智能化行为模拟

- 绕过检测技术

  效果展示:
  ![效果展示-加载可能有些慢](https://i.ibb.co/yPkZLd0/mnggiflab-compressed-20231026-215253-min.gif)

## 为什么选择 Luna

![效果展示-加载可能有些慢](https://i.ibb.co/nftHyHW/511714127971-pic.jpg)

经过大量测试，目前基本可以过掉主流抗指纹识别;

```
测试网址:
https://www.browserscan.net/
https://uutool.cn/browser/
https://abrahamjuliot.github.io/creepjs/
```



## 使用限制

1、目前仅支持 Windows x86-64 架构，其他平台测试尚不充分。

2、mac arm版仅提供开发时测试、并不能用于生产环境、因为大部分指纹并不支持。



##  Luna文档部分



详细的使用说明和示例代码，请查看本项目的[文档](https://github.com/musiclover789/luna-browser)。

示例代码部分也可以查看源码的test_case包下内容。




## Luna浏览器部分

目前，我们已经将浏览器文件上传到 百度 网盘，并提供了下载链接：



新版本-win-[670MB]连接:链接：https://pan.baidu.com/s/1S3ZdbFHTtaZgW2dInc6JDA 提取码：3pmd

Mac-arm版[114MB]:链接: https://pan.baidu.com/s/1au226sENM5XcoB7SPhEYZA 提取码: lbfs



<Mac版本仅供开发测试使用，部分抗指纹功能不可用，方便Mac开发人员进行开发-完全免费-无限制>

<win版本-没有授权文件的用户,仅可以测试useragent指纹部分,其他指纹不会生效>

如何获取授权文件联系作者获取;





****

### 目前支持指纹项:

|      | 指纹项                      | 技术方案                            |      | win  | mac  |
| ---- | --------------------------- | ----------------------------------- | ---- | ---- | ---- |
|      | user_agent指纹              | headless模式下、也会生效            |      |      |      |
|      | canvas指纹                  | 真实指纹库、难以识别                |      |      |      |
|      | webgl指纹                   |                                     |      |      |      |
|      | platform平台                |                                     |      |      |      |
|      | timezone时区                |                                     |      |      |      |
|      | timezone_offset时区偏移量   |                                     |      |      |      |
|      | languages语言               | 无论是国际API、还是navigator 均生效 |      |      |      |
|      | userAgentData               |                                     |      |      |      |
|      | header 修改                 | 可以修改http请求协议层header        |      |      |      |
|      | deviceMemory                |                                     |      |      |      |
|      | hardwareConcurrency         |                                     |      |      |      |
|      | UNMASKED_VENDOR_WEBGL       | 显卡                                |      |      |      |
|      | UNMASKED_RENDERER_WEBGL     | 显卡                                |      |      |      |
|      | GL_VERSION                  | 显卡                                |      |      |      |
|      | GL_SupportedExtensions      | 显卡                                |      |      |      |
|      | GL_VENDOR                   | 显卡                                |      |      |      |
|      | GL_RENDERER                 | 显卡                                |      |      |      |
|      | GL_SHADING_LANGUAGE_VERSION | 显卡                                |      |      |      |
|      | 是否webdriver               | 已处理                              |      |      |      |
|      | 是否brave                   | 已处理                              |      |      |      |
|      | 是否selenium                | 已处理                              |      |      |      |
|      | 是否来自于真实键盘          | 已处理                              |      |      |      |
|      | 是否来自于真实鼠标          | 已处理                              |      |      |      |
|      | 鼠标移动轨迹                | 已处理                              |      |      |      |
|      | 其他机器人检测              | 已处理                              |      |      |      |
|      | webRTC                      | 可以自行设置出口ip                  |      |      |      |
|      | screen                      | 已处理                              |      |      |      |
|      | 声卡指纹                    | 0-1000任意整数                      |      |      |      |






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



相关文档：

| [第一课-常见概念介绍.md](https://github.com/musiclover789/luna-browser/blob/main/第一课-常见概念介绍.md) |
| ------------------------------------------------------------ |
| [第三课-brower对象.md](https://github.com/musiclover789/luna-browser/blob/main/第三课-brower对象.md) |
| [第二课-第一个小例子.md](https://github.com/musiclover789/luna-browser/blob/main/第二课-第一个小例子.md) |
| [第五课.md](https://github.com/musiclover789/luna-browser/blob/main/第五课.md) |
| [第四课page对象.md](https://github.com/musiclover789/luna-browser/blob/main/第四课page对象.md) |



完整文档参考:  https://github.com/musiclover789/luna-browser

备注: 具体指纹修改项，请参阅上面表格部分。

-----



## 2. 基于视觉的操作

Luna 基于视觉的页面操作方法让您可以使用截图的方式来控制浏览器，也支持传统的 CSS 和 XPath 选择器等方式。这意味着您可以立即看到页面上的内容并执行操作，而不必等待特定事件触发。

这一特性的最大优势在于速度，因为您可以像人一样看到什么就可以操作什么。这样的交互方式使得 Luna 极为高效。



## 3. 代理 IP 多样性

Luna 支持市面上所有类型的代理 IP，包括 HTTP、HTTPS 和 SOCKS5，无论代理 IP 是否需要密码，Luna 都完全兼容。理论上，使用 Luna 进行爬取的请求将无法被追踪。



## 4. 多进程和多线程

Luna 考虑到了多进程和多线程的应用场景，使得您可以并发执行多个任务，提高了爬虫的效率。



## 5. 网络数据包过滤

Luna 考虑到了、可能会协议和浏览器混编的方式、和可能的协议采集需求,所以继承了比较完备的cookie方案，和数据包过滤方案、方便采集数据使用、已经封装了比较完善的 一对一 请求过滤。



- 如果您不用防指纹识别部分功能、就下载普通的chrome浏览器即可。
- 另外、鼠标移动轨迹、键盘输入、鼠标滚轮、如果没有luna浏览器配合、那么依然会被轻易识别为机器人。



----------------------




#### 常见问题回复

1、可以自己随便修改指纹吗？

答：是的、理论上无限指纹;

2、目前支持Linux 系统吗？

答:暂时不支持、

3、原理是?

答:修改chromium内核。

4、有体积更小的浏览器么？

答：无、参考新版。

5、为什么我测试基于视觉时候发现，出现bug

答：下载代码后不要修我的项目名字 叫luna

6、第三方库可以用的么，如Selenium Pyppeteer Playwright 。

答：不支持、经过大量测试、发现第三方框架特别容易被识别为cdp控制；

不信，你就拿www.browserscan.net测试一下,所以不再兼容第三方框架,不要反复问了。





**抗指纹部分需要授权 <非付费用户，只能测试useragent 部分效果>** 授权联系-QQ: 80158153


