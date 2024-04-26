# Luna - 基于视觉的抗指纹爬虫第三方库



​		Luna是专为抗指纹自动化爬虫设计的工具，包含抗指纹浏览器和自动化框架，让您能够自由实现所需功能。



作者QQ: 80258153



python版本框架:https://github.com/musiclover789/luna_python



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



##  Luna文档部分



详细的使用说明和示例代码，请查看本项目的[文档](https://musiclover789.github.io/lunadocs/)。

示例代码部分也可以查看源码的test_case包下内容。




## Luna浏览器部分

目前，我们已经将浏览器文件上传到 百度 网盘，并提供了下载链接：

老版本-win-[2GB]链接: https://pan.baidu.com/s/14EZw9DvCtO998LOwo_epvA 提取码: mm6s

新版本-win-[670MB]连接: https://pan.baidu.com/s/1kxnZO6BaF6cE3e8Ugrzdog 提取码: tbwg

Mac-arm版[114MB]:链接: https://pan.baidu.com/s/1au226sENM5XcoB7SPhEYZA 提取码: lbfs

<Mac版本仅供开发测试使用，部分抗指纹功能不可用，方便Mac开发人员进行开发-完全免费-无限制>

<win版本-没有授权文件的用户,仅可以测试useragent指纹部分,其他指纹不会生效>

如何获取授权文件联系作者获取;



请查阅以下内容，了解老版和新版本框架使用上差异:
https://github.com/musiclover789/luna-browser/edit/main/README.md



****

### 目前支持指纹项:

|      | 指纹项                      | 技术方案                            |      | win  | mac  |
| ---- | --------------------------- | ----------------------------------- | ---- | ---- | ---- |
|      | user_agent指纹              | headless模式下、也会生效            |      |      |      |
|      | cavnvas指纹                 | 真实指纹库、难以识别                |      |      |      |
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
|      | webRTC                      | 已处理                              |      |      |      |
|      | screen                      | 已处理                              |      |      |      |



###### 技术交流Q群: 524592021



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
    // 打开一个页面  如果你想看更多示例 请参考文档 https://musiclover789.github.io/lunadocs/docs/category/case
    browserObj.OpenPage("https://www.baidu.com")
    fmt.Println("恭喜你，非常nice的第一个案例")
    time.Sleep(1 * time.Hour)

}

```



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



代码调用抗指纹部分示例，最好您参考文档里面的详细内容，这里仅黏贴一部分代码，提供参考。

## 

```bash
package test_case

import (
	"fmt"
	"luna/devtools"
	"luna/luna_utils"
	"testing"
	"time"
)

/***
这个例子是一个入门例子,其目的是希望你可以通过这个例子,成功的打开浏览器、并且成功的访问一个网址仅此而已.
//测试我们是否被发觉出来的https://undetectable.io/zh-cn/blog/post/browser-fingerprinting-test-services
https://uutool.cn/browser/
*/

//如果对这种写法不熟悉,可以直接用main函数替代,也就是把里面的代码全部粘贴到main函数里面是一样的.
func TestFingerprint(t *testing.T) {
	//代码开始

	//第一步
	//杀死之前你可能测试过的chromium进程;
	//启动前先杀死其他的chromium进程;为了防止程序以及停止但是依然在内存中贮存;
	//其目的是为了反复测试的时候不会产生过多的内存驻留进程、实际使用时候请根据实际情况选择使用;
	//他会根据你的系统不同,使用命令行的命令进行杀死其他chromium进程
	luna_utils.KillProcess()

	//第二步
	//设置自己需要的指纹
	args := []string{
		/***
		a、luna_user_agent 这个参数目前仅会使得 navigator.userAgent 的值发生变化；也就是说http、http2、等协议层 仍然不会被替换;如果需要替换往下看。
		b、headless 模式下、navigator.userAgent 也是会被替换；并不会显示任何headless 的userAgent;
		*/
		"--luna_user_agent=Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
		/***
		a、navigator.platform 值被替换
		*/
		"--luna_platform=win64",
		/***
		a、timeZone 经过测试、时区均会被替换成指定的时区
		b、timeZone 的时区偏移量也是自动计算的
		c、测试js :
				const date = new Date();
				const timeZone = date.getTimezoneOffset();
				const timeZoneOffset = -timeZone / 60;
				console.log("Time Zone: " + Intl.DateTimeFormat().resolvedOptions().timeZone);
				console.log("Time Zone Offset: " + timeZoneOffset);
		*/
		"--luna_timezone=Europe/London",
		/***
		languages: 以下均会生效
		a、navigator.language
		b、new Intl.DateTimeFormat().resolvedOptions().locale
		说明:理论上、无论是国际API、还是navigator 均生效。

		*/
		"--luna_languages=en-GB",
		/***
			navigator.userAgentData:
			值示例:Google Chrome:92-luna-Chromium:92-luna-Not-A.Brand:24-luna-platform:win32-luna-mobile:false-luna-platform_version:6.1-luna-ua_full_version:92.0.4515.186-luna-model:PC-luna-architecture:x86_64
			格式、-luna- 为每组的分隔符
				: 为key、value分隔符
			举例说明:
				1：Google Chrome:92-luna-Chromium:92 分别为2组 Google Chrome:92 和 Chromium:92
				每组的值key、value分别是: 组1 key 是 Google Chrome value是 92 ，组2 Chromium:92 key是Chromium value是 92
				2：他们分别代表什么呢？代表执行"navigator.userAgentData" 后的、brands的值；你自己浏览器测试以下就知道了。
				3：以下几个分别是特别值、分别代表执行"navigator.userAgentData" 后参数值
		          mobile:false //是否是手机
		          ua_full_version:92.0.4515.186 //浏览器版本
				  model:PC	//设备的型号信息
		          architecture:x86_64 //芯片类型、比如 arm
				注:这4个特殊值反而要特别注意、因为你可能看不到他的具体值、但是第三方的指纹测试是可以看到他的值的
		        如：https://abrahamjuliot.github.io/creepjs/ 测试，你可以看以下
		*/
		"--luna_userAgentData=Google Chrome:92-luna-Chromium:92-luna-Not-A.Brand:24-luna-platform:win32-luna-mobile:false-luna-platform_version:6.1-luna-ua_full_version:92.0.4515.186-luna-model:PC-luna-architecture:x86_64",
		/***
		   luna_header_1:这个是为了替换http请求时候的数据包里面的header参数的;
			key 格式:--luna_header_1
				--luna_header_2
				...
				--luna_header_11
				取值范围[1-11]
				也就是说，你最多可以替换11个、为什么这么设计呢？因为太多了影响效率；个人感觉1-11够用了
			value 格式:
					key -lunareplace- value //代表将左侧匹配到的key、替换成右侧的value
					key -lunaremove- value //代表将左侧匹配到的key、删除这个header参数项
					key --lunaadd-- value //增加这个key、value header项
				注意格式中没有空格，这里只是方便看、正常应该是:key-lunareplace-value
			使用场景:
				1、为了配合navigator.userAgent变化而协议层没有变
				2、为了删除协议层暴露的隐私数据
				3、增加你想增加的请求头数据
		*/
		"--luna_header_1=User-Agent-lunareplace-Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
		"--luna_header_2=sec-ch-ua-arch-lunaremove-",
		"--luna_header_3=sec-ch-ua-platform-lunaremove-",
		"--luna_header_4=accept-language-lunareplace-en;q=0.9",
		/***
		deviceMemory 内存
			取值范围建议:[0.25, 0.5, 1, 2, 4, 8]
			表示设备的内存大小，以 GB 为单位。
		对应js的navigator.deviceMemory
		*/
		"--luna_deviceMemory=8", //[0.25, 0.5, 1, 2, 4, 8]
		/***
		deviceMemory 内存
			属性是一个数字，表示设备的逻辑处理器核心数。
		对应js的\navigator.hardwareConcurrency
		*/
		"--luna_hardwareConcurrency=16",
		/***
		显卡信息: 这些是与 WebGL 相关的属性和扩展名称，用于提供关于 WebGL 上下文和渲染器的信息：
		UNMASKED_VENDOR_WEBGL：表示 WebGL 上下文所使用的图形硬件的供应商。它通常返回一个字符串，表示硬件供应商的名称。
		UNMASKED_RENDERER_WEBGL：表示 WebGL 上下文所使用的图形硬件的渲染器。它通常返回一个字符串，表示图形渲染器的名称。
		GL_VERSION：表示 WebGL 上下文支持的 OpenGL 版本。它通常返回一个字符串，表示支持的 OpenGL 版本号。
		SupportedExtensions：表示 WebGL 上下文支持的扩展列表。它是一个数组，包含了当前 WebGL 上下文支持的各种扩展名称。
		GL_VENDOR：表示 WebGL 上下文所使用的图形硬件供应商的名称。它通常返回一个字符串，表示图形硬件供应商的名称。//默认值:WebKit,意思是我观察到chromium源代码中是写死到这个值,没特殊情况无需修改
		GL_RENDERER：表示 WebGL 上下文所使用的图形渲染器的名称。它通常返回一个字符串，表示图形渲染器的名称。//默认值:WebKit WebGL,意思是我观察到chromium源代码中是写死到这个值,没特殊情况无需修改
		GL_SHADING_LANGUAGE_VERSION：表示 WebGL 上下文支持的着色语言版本。它通常返回一个字符串，表示支持的着色语言版本号。
		对应的js、代码太长了我就不写了你可以观察下面的代码 runtime.Evaluate(conn, luna_script.TestWebDriver()) 里面基本上都测试了。
		*/
		"--luna_UNMASKED_VENDOR_WEBGL=Intel Corporation",           //Google Inc. (Apple)
		"--luna_UNMASKED_RENDERER_WEBGL=Intel(R) UHD Graphics 620", //ANGLE (Apple, Apple M1 Pro, OpenGL 4.2)
		"--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 Intel(R) UHD Graphics 620)",
		`--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]`,
		"--luna_GL_VENDOR=WebKit",         //默认值 WebKit
		"--luna_GL_RENDERER=WebKit WebGL", //默认值WebKit WebGL
		"--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)", //测试值 WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
		/***
			cavans指纹:
				原理&区别:
					luna_cavans_random_str
					luna_cavans_random_int
				a:首先这两个你只可以2选1，不能同时写，因为同时写基本上等同于只设置了luna_cavans_random_str、从而让luna_cavans_random_int变得时区作用了
				b:luna_cavans_random_str他只是在toDataURl的时候加入了随机字符串;这个字符串你需要自己传入,至于值 只要是字符串 剩下随便。
		          这有什么用呢？
						这样会使得整个浏览器、无论是标签页、窗口、隐藏窗口探测cavans指纹时、均报纸一致；
						而且你已经传入了随机字符串、所以他的指纹就变得随机了
					    这样的缺点是、因为是随机了结果值、导致很假，但是假的范围大。
				c：luna_cavans_random_int 他的原理只是调整了png的压缩level。
					这样有什么用呢？
					1、优点是比较真实、并不会影响任何结果
					2、缺点是你的取值范围只有[0-9]指纹数量有限、chromium源代码每个版本默认值可能会不太一样.我测试的版本是默认值是:3、我也观察到有些版本默认是4
					请注意下面是老版本浏览器的写法，新版本已经废弃这两个参数
					改成:
					--luna_canvas_random_int_number=4
					取值范围1-999 要求整型。
		*/
		//"--luna_cavans_random_str=B3B4",
		"--luna_cavans_random_int=1", //取值0-9 默认值3

		/***
			screen: 对象是一个表示用户屏幕信息的全局对象。下面是对 screen 对象中常见属性的解释：
				height：表示屏幕的总高度（以像素为单位）。
				width：表示屏幕的总宽度（以像素为单位）。
				availHeight：表示可用的屏幕高度，即减去操作系统任务栏或其他系统UI的高度后剩余的屏幕高度。
				availWidth：表示可用的屏幕宽度，即减去操作系统任务栏或其他系统UI的宽度后剩余的屏幕宽度。
				availLeft：表示可用屏幕的左边界相对于整个屏幕的左边界的偏移量。在多显示器设置中，这个属性可以用来确定屏幕的位置。
				availTop：表示可用屏幕的顶部边界相对于整个屏幕的顶部边界的偏移量。在多显示器设置中，这个属性可以用来确定屏幕的位置。
				internal：表示屏幕是否是内部显示器。如果屏幕是内部显示器，则返回 true，否则返回 false。//我个人理解、如果如果你是笔记本电脑买来就自带屏幕就是true、如果主机自己接入的外置显示器就是false
				primary：表示屏幕是否是主要显示器。如果屏幕是主要显示器，则返回 true，否则返回 false。
				top：表示屏幕的顶部边界相对于整个屏幕的顶部边界的偏移量。在多显示器设置中，这个属性可以用来确定屏幕的位置。
				left：表示屏幕的左边界相对于整个屏幕的左边界的偏移量。在多显示器设置中，这个属性可以用来确定屏幕的位置。
				scaleFactor：表示屏幕的缩放因子。如果操作系统设置了缩放级别，则返回缩放因子；否则，返回 1。
			//总之 如果没有必要、这个值我建议不要轻易修改、因为我这个框架是基于视觉的、如果你使用视觉点击、可能会受到影响。
		      还有这个参数的意义也并不大。
			格式:
				举例:height:800,width:978
				逗号","分隔每组
				冒号":"分隔key、value
				如果你不傻、应该可以看懂，我就不仔细说了。
		*/
		//"--luna_screen=height:800,width:978,availHeight:1024,availWidth:934,availLeft:0,availTop:28,internal=true,primary=true,top:34,left:34,scaleFactor=2",
		//"--luna_screen=height:1440,width:2560,availHeight:1440,availWidth:2560,availLeft:0,availTop:0,internal=true,primary=true,scaleFactor=2",
		"--luna_key=WERWER234234WERWEr345345",
		/***
			最后:
				1、这些均是常见的硬件信息、但是并不能代表所有、比如声卡、比如字体、比如每个版本的css特性、比如 语音合成器等
		          这完全取决于你面对的挑战到底有多大、（but、一般也够用了）
				2、从学术角度、我个人建议、要用真实值
				3、从使用的角度、如果你不许要修改某些指纹、你就不用写他,
		           对于里面的key、value 也是一样、如果你不需要可以不用写，
		           你不写他自己就是默认值，该是什么样就是什么样。
				4、你也不要以为你改了几个参数、他就认不出来你了,真实的情况是，这个是很有技巧的搭配指纹组,而且大数据统计的识别方式还是很可怕的，透肉看骨一般的
				5、仅从学术角度、关于鼠标移动、单击、双击、键盘事件 的 trusted 默认是true; 理论上无法识别出来是不是程序在点击、
				6、仅从学术角度、是否是 、headless 包括常见的 识别是否是playwright、puppeteer、webdriver、Selenium 理论上无法识别出来；因为我也没用这些框架、而是自己封装的。
		           并将已经默认做了特征剔除。
		*/
	}

	//初始化浏览器对象
	chromiumPath := "你自己的浏览器下载地址"
	_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		//设置缓存目录,
		CachePath: luna_utils.CreateCacheDirInSubDir("你自己想设置的浏览器缓存数据地址"),
		//设置你认为需要的指纹信息
		Fingerprint: args,
		//设置非隐身模式
		Headless: false,
	})

	//打开一个tap
	browserObj.OpenPage("https://www.baidu.com")

	fmt.Println("如果你可以成功的运行这个代码,那么恭喜你,已经基本上知道如何设置基本的指纹信息了.至少看起来很专业的样子了")
	time.Sleep(1 * time.Hour)
	//代码结束
}

/*

./Chromium.app/Contents/MacOS/Chromium

设置user-agent;
--luna_user_agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.48 Safari/537.36"

设置Platform;
--luna_platform="Win32"
--Linux armv81  MacIntel  Win32 "" Linux x86_64 iPhone Unsupported platform


设置时区；
--luna_timezone="Pacific/Midway"

设置语言:
--luna_languages="es-AR en-GB"
已经默认修改了 国际化接口 和 navigator了
en 中文
zh-CN 中国中文
en-CA 加拿大英语
en-US 美国英语
en-GB 英国英语
es-AR 西班牙语阿根廷

设置navigator.userAgentData的js值
格式-luna-是每组的分隔符 : 是key和value的分隔符
--luna_userAgentData=Google Chrome:92-luna-Chromium:92-luna-Not-A.Brand:24-luna-platform:win32-luna-mobile:false-luna-platform_version:92

替换http请求时数据包的header
慎重使用,因为会稍微影响性能
luna_header_1
格式:luna_header_<根数组1-11>也就是最多你可以更改11个header
其中-lunaremove-、-lunareplace-、-lunaadd- 是三种分隔符，分别代表 删除这个header、替换这个header、和增加这个header
"--luna_header_1=User-Agent-lunareplace-Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
		"--luna_header_2=sec-ch-ua-arch-lunaremove-",
		"--luna_header_3=sec-ch-ua-platform-lunaremove-",
		"--luna_header_4=accept-language-lunareplace-en;q=0.9",

//无需设置、自动鼠标移动、点击、键盘事件的trusted为true;

内存
--luna_deviceMemory="16"

逻辑核
--luna_hardwareConcurrency="10"



*/


```


#### 常见问题回复

1、可以自己随便修改指纹吗？

答：是的、理论上无限指纹;

2、目前支持Linux 系统吗？

答:暂时不支持、

3、原理是?

​	修改chromium内核。

4、有体积更小的浏览器么？

答：无、参考新版。

5、为什么我测试基于视觉时候发现，出现bug

答：下载代码后不要修我的项目名字 叫luna

6、第三封库可以用的么，如Selenium Pyppeteer Playwright 。

答：不支持、经过大量测试、发现第三方框架特别容易被识别；所以不再兼容第三方框架。

7、我没有找到如何类似xpath cssselecter选择器。

答：参考如下代码

```
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
		/***
	    鼠标移动到指定坐标,
		p1.SimulateMouseMoveToTarget(endX, endY float64) error
		鼠标移动到指定元素、动作：会先自动鼠标滚轮到指定元素位置，然后开始移动鼠标，最后返回 元素区域内 随机x，y坐标
		p1.SimulateMouseMoveToElement(selector string) (err error, randomX, randomY float64)
		鼠标 滚轮 滚动到指定 元素
	    p1.SimulateScrollToElementBySelector(selector string) error
		获取当前 选择器 指定 元素，返回如下数据结构元素数据
		//// Node 表示节点信息的数据结构
			//type Node struct {
			//	NodeType      int64
			//	NodeName      string
			//	NodeValue     string
			//	TextContent   string 节点文本数据
			//	HTMLContent   string 节点html数据
			//	CSSSelector   string css选择器内容
			//	XPathSelector string xpath选择器 内容
			//}
	    p1.GetElementByCss(selector string) (error, Node)
		//获取所有字节点 数据
		p1.GetAllChildElementByCss(selector string) (error, []Node)
		//获取第一个字节点 对象数据
		p1.GetFirstChildElementByCss(selector string) (error, Node)
		//获取最后一个字节点 对象数据
		p1.GetLastChildElementByCss(selector string) (error, Node)
		//下一个相邻节点
		p1.GetNextSiblingElementByCss()
		//上一个相邻节点
		p1.GetPreviousSiblingElementByCss()
		//父节点
		p1.GetParentElementByCss()
		//文档已经存在的,模拟人类点击
		p1.SimulateMouseClickOnPage(x, y float64)
		//拖动
		//p1.SimulateDrag(startX, startY, endX, endY float64)
		//其他参考文档https://musiclover789.github.io/lunadocs/docs/tutorial-basics/devtools.page
	*/
```

8、如何操作cookie、文档中我并没有找到

答：参考如下代码

```
//示例
var wg sync.WaitGroup // 同步等待
wg.Add(1)             // 增加等待的数量
err, p1 := browserObj.OpenPageAndListen("https://www.baidu.com/", func(devToolsConn *protocol.DevToolsConn) {
   //第一个处理
   devToolsConn.ShowLogJson(true)
   network.EnableNetwork(devToolsConn)
network.RequestResponseAsync(devToolsConn, func(requestId string, request, response map[string]interface{}) {
            //这里的request & response都是一一对应的请求、header里面的cookie可以自己看一下
            fmt.Println(luna_utils.FormatJSONAsString(request),luna_utils.FormatJSONAsString(request))
            平时用不上,并不是每个请求都有请求报体；需要根据请求的url自行判断是否需要使用
            //network.GetResponseBody(devToolsConn,requestId,time.Minute)
        })
})

```

如果设计到启动时候就给固定的登陆cookie、参考

https://musiclover789.github.io/lunadocs/docs/tutorial-extras/case_fingerprint

这篇文章、中的luna_header 部分

或者您可以参考此代码段

``

```
_, p1 := browserObj.OpenPageAndListen("https://www.baidu.com/", func(devToolsConn *protocol.DevToolsConn) {
    devToolsConn.ShowLog(false)
    //page.PageEnable(devToolsConn) 可用于观察页面加载触发事件对日志,如果检测网络可以network.EnableNetwork(devToolsConn) 配合devToolsConn.ShowLog(true) 使用
    //第一个处理
    //设置cookie 以百度举例 此处的url如果设置为https://www.baidu.com 则仅对这个对应的url的cookie 设置值,如果以有值 则覆盖
    //可以根据你自己的需求,对任何url的cookie进行操作,再次举例 比如 希望对
    //https://gimg3.baidu.com 进行干预,你也可以写 url 为https://gimg3.baidu.com
    network.SetCookieByURL(devToolsConn, "luna_url", "luna-cookie", "https://www.baidu.com")
    //设置cookie,一样对只是换成了 domain 而已 ,相对于domain URL可能控制对更加细腻
    network.SetCookie(devToolsConn, "luna_domain", "luna-cookie", "www.baidu.com")
    //至于这里的devToolsConn 可以从任何页面对象获取，也不一定非要写在OpenPageAndListen里面,写外面也可以
    //用哪个页面的devToolsConn,和对应的 (domain 或url) 则对哪个页面生效
    /***
    ClearBrowserCookies
    */
    //清空cookie 这个要用在这个地方才可以有效,可以确保页面加载全周期 清空cookie,放在其他地方 就不好说了, 结合具体场景自行考虑使用
    //network.ClearBrowserCookies(devToolsConn)
})
//写在外面的示例
network.SetCookie(p1.DevToolsConn, "luna_domain_abc", "luna-cookie", "www.baidu.com")
//如何获取DevToolsConn的示例
//p1.DevToolsConn
//browserObj.DevToolsConn
fmt.Println("打印cookie")

//根据自己的需求,这个地方是一个数组，可以获取任意url的cookie
urls := []string{"https://www.baidu.com"}
cookies, _ := network.GetCookies(p1.DevToolsConn, urls)
//循环打印
for _, result := range gjson.Parse(luna_utils.FormatJSONAsString(cookies)).Get("result.cookies").Array() {
    fmt.Println(result.Get("name").String(), result.Get("value").String(), result.Get("domain").String())
    //可以将以上的cookie存储到数据库,根据你到逻辑
}
```



如何上传文件、代码片段

```
_, p1 := browserObj.OpenPageAndListen("https://graph.baidu.com/pcpage/index?tpl_from=pc", func(devToolsConn *protocol.DevToolsConn) {
    devToolsConn.ShowLog(false)
})
time.Sleep(3 * time.Second)
//点击文件上传的按钮,示例是Xpath选择器、也支持css选择器
_, x, y := p1.GetElementPositionByXpathOnPage("//*[@id=\"app\"]/div/div[1]/div[7]/div/span[1]/span[1]")
input.SimulateMouseClick(p1.DevToolsConn, x, y)
//获取对应的<input name="file" type="file" >
//文件css选择器路径 ,注意这里是css选择器,并不是其他的
seletor := "#app > div > div.page-banner > div.page-search > div > div > div.graph-d20-search-layer-contain > div.graph-d20-search-layer-choose > div > form > input"
//第一个参数就是css选择器路径,第二个参数是一个数组 存放自己需要上传的本地图片路径
p1.UploadFiles(seletor, []string{"/Users/Pictures/IMG_2614.JPG"})
```


当你准备写多线程调用的时候请注意2点；
1、请一定注意要设置CachePath；建议直接用我这种方式。
原理:每次都会在这个/golang/cache文件夹下，创建随机文件夹存储临时文件。
如果你想直接写死路径，那么请保持变化，如果是多线程，要保证每个线程不同的缓存目录

```
_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		CachePath: luna_utils.CreateCacheDirInSubDir("/golang/cache"),
		//设置非隐身模式
		Headless: false,
	})
```

2、请留意，因为忘记关闭浏览器导致的 进程贮存问题。
可以考虑
luna_utils.KillProcess()
当然这个要根据你的需求来处理，如果你正常关闭，不会遇到这个问题,但是大规模并发很难保证，所以还是建议考虑CachePath: luna_utils.CreateCacheDirInSubDir("/golang/cache"),的方式







**抗指纹部分需要授权 <非付费用户，只能测试useragent 部分效果>** 授权联系-QQ: 80158153
