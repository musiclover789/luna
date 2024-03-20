# Luna - 基于视觉的抗指纹爬虫第三方库


## <如果您有抗指纹方面需求、可以联系QQ：80258153>

Luna 是一款强大的第三方库，专为抗指纹自动化爬虫而设计。通过利用视觉特征和先进的算法，Luna 提供了一种有效的方法来对抗现代爬虫检测技术，保护您的网络资源免受恶意爬取和滥用。

## 功能亮点

- **强大的抗指纹技术：** Luna 提供了先进的抗指纹技术，使您的爬虫程序难以被识别。
- **视觉特征解析：** 基于视觉特征的页面解析和操作，使爬虫更智能。
- **简单易用的接口：** Luna 提供简单易用的接口，轻松集成和使用它的功能。
- **智能化行为模拟：** 模拟用户行为，有效应对现代爬虫检测技术。
- **绕过检测技术：** 具备绕过常见爬虫检测技术的能力，确保您的爬虫不容易被拦截。

  效果展示-加载可能有些慢
![效果展示-加载可能有些慢](https://i.ibb.co/yPkZLd0/mnggiflab-compressed-20231026-215253-min.gif)

## 为什么选择 Luna？

使用 Luna，您可以快速构建出智能、高效、难以被识别的爬虫程序。不论是在开发自动化测试脚本、数据采集应用还是其他需要模拟用户行为的场景中，Luna 都能为您提供可靠的解决方案。

不论您是开发人员、数据科学家还是网络安全专家，Luna 都是您在抗指纹爬虫领域的得力助手。让 Luna 成为您的选择，保护您的网络资源，确保您的数据安全。



目前支持的操作系统是 Windows，且仅限于 x86-64 架构。已经在 Windows x86-64 硬件环境下进行了测试。其他操作系统或平台的测试尚不充分，因此不建议在这些系统上使用。



## 开始使用 Luna

详细的使用说明和示例代码，请查看本项目的[文档](https://musiclover789.github.io/lunadocs/)。

示例代码部分也可以查看源码的test_case包下内容。


## Luna浏览器部分
浏览器部分是 Luna 的核心功能之一，它使您能够执行抗指纹爬虫任务。请注意，您需要下载适用于 Luna 的专用浏览器才能实现指纹防识别。该浏览器的大小约为 2GB，因此需要一些时间来下载。如果您没有抗指纹需求、可以直接用您的chrome或其他浏览器即可。

目前，我们已经将浏览器文件上传到 百度 网盘，并提供了下载链接：



链接: https://pan.baidu.com/s/14EZw9DvCtO998LOwo_epvA 提取码: mm6s



### 如何测试



正常使用情况下，这部分完全代码来控制。但是为了方便您测试luna浏览器的基础功能，可以使用手动的方式来测试。
<非付费用户，只能测试useragent 部分效果>
原理：这个是基于chromium的源代码，对内核进行修改编译的。

测试步骤：

1、在您的c盘根目录，手工建立一个文件夹luna-temp
示例:           C:\luna-temp

2、打开浏览器，有些同学不太知道怎么打开浏览器，就是下载好后，找到目录里面的 chrome.exe 用鼠标双击打开.
ps:" I 服了 you ，如果不知道怎么双击打开，我建议你直接放弃研究本项目"

3、你会发现，C:\luna-temp 目录里面会有一个uname.txt文件，这个是做授权用的，你测试的时候无需关心。

4、在这个目录下，手工创建一个文件，命名为1696987203497907900 



也就是路径应该是 C:\luna-temp\1696987203497907900



注意不要有.txt的扩展名。



然后用记事本打开，黏贴如下内容：

```
luna_user_agent=Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
luna_platform=win64
luna_timezone=Europe/London
luna_timezone_offset=3600000
luna_languages=en-GB
luna_userAgentData=Google Chrome:92-luna-Chromium:92-luna-Not-A.Brand:24-luna-platform:win32-luna-mobile:false-luna-platform_version:6.1-luna-ua_full_version:92.0.4515.186-luna-model:PC-luna-architecture:x86_64
luna_header_1=User-Agent-lunareplace-Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
luna_header_2=sec-ch-ua-arch-lunaremove-
luna_header_3=sec-ch-ua-platform-lunaremove-
luna_header_4=accept-language-lunareplace-en;q=0.9
luna_deviceMemory=8
luna_hardwareConcurrency=16
luna_UNMASKED_VENDOR_WEBGL=Intel Corporation
luna_UNMASKED_RENDERER_WEBGL=Intel(R) UHD Graphics 620
luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 Intel(R) UHD Graphics 620)
luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
luna_GL_VENDOR=WebKit
luna_GL_RENDERER=WebKit WebGL
luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
luna_cavans_random_str=B3B4
remote-debugging-port=55392
user-data-dir=C:\workspace\tempcatch\chromium_user_data_1696987203497907900
```

退出浏览器，再次打开。理论上你可以抓包看一下，无论你访问任何网站，无论你的电脑是什么配置，你的useragent会改成配置文件里面的 

Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36

当然这个内容，你自己测试随意修改。 (未付费用户，其他指纹不会生效)

还记得这个配置文件的命名么？1696987203497907900

这个是当前毫秒数而已，如果多个配置文件，他只会选用最新的。也就是最新的当前毫秒数，这些配置文件里面的其他信息，*授权后也是可以修改的* ，但是正常情况下，我们都是用 程序去调用，并不会人工去测试这些东西。

仅提供 手工测试 爱好者。





## 快速入门

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
    browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
        // 设置非隐身模式
        Headless: false,
    })
    // 打开一个页面  如果你想看更多示例 请参考文档 https://musiclover789.github.io/lunadocs/docs/category/case
    browserObj.OpenPage("https://www.baidu.com")
    fmt.Println("恭喜你，非常nice的第一个案例")
    time.Sleep(1 * time.Hour)

}

```



###### 如果您执行到这一步遇到依赖包问题，那么请执行 go get -v -d ./... 命令即可。





## 特点

## 1. 抗指纹特性

Luna 强大的抗指纹技术可以模拟和对抗多种常见爬虫检测技术，包括但不限于：

- 时区指纹
- 显卡指纹
- User-Agent 指纹
- Platform 指纹
- Languages 指纹
- Device Memory 指纹
- Hardware Concurrency 指纹
- Canvas 指纹
- 鼠标滚动指纹（真实很难被识别）
- 鼠标移动轨迹（真实很难被识别）
- 键盘真实输入（包括内置转输入法等）

理论上，Luna 可以成功对抗这些指纹技术，使您的爬虫在操作时不容易被识别。更多详细信息请查看我们的[文档](https://musiclover789.github.io/lunadocs/)。

## 2. 基于视觉的操作

Luna 基于视觉的页面操作方法让您可以使用截图的方式来控制浏览器，也支持传统的 CSS 和 XPath 选择器等方式。这意味着您可以立即看到页面上的内容并执行操作，而不必等待特定事件触发。

这一特性的最大优势在于速度，因为您可以像人一样看到什么就可以操作什么。这样的交互方式使得 Luna 极为高效。

## 3. 代理 IP 多样性

Luna 支持市面上所有类型的代理 IP，包括 HTTP、HTTPS 和 SOCKS5，无论代理 IP 是否需要密码，Luna 都完全兼容。理论上，使用 Luna 进行爬取的请求将无法被追踪。

## 4. 多进程和多线程

Luna 考虑到了多进程和多线程的应用场景，使得您可以并发执行多个任务，提高了爬虫的效率。

## 5. 网络数据包过滤

Luna 考虑到了、可能会协议和浏览器混编的方式、和可能的协议采集需求,所以继承了比较完备的cookie方案，和数据包过滤方案、方便采集数据使用、已经封装了比较完善的 一对一 请求过滤。



##### 





付费部分：

- 如果您不用防指纹识别部分功能、就下载普通的chrome浏览器即可。其他部分完全免费、开源，可用。

- 唯一需要付费的部分是 Luna 浏览器的抗指纹功能。

- 请注意，Luna 浏览器的用户代理（User-Agent）功能是免费的，您可以使用它进行测试，以方便您更好地了解功能。

- 另外、鼠标移动轨迹、键盘输入、鼠标滚轮、如果没有luna浏览器配合、那么依然会被轻易识别。

- 联系我之前、希望您先自己跑通、至少可以运行指纹、视觉爬虫部分。luna浏览器也是免费下载的

  (除了高级指纹、其他都没有任何限制)

如果您有抗指纹方面需求、可以联系QQ：80258153

如果您有技术方便问题、或者bug反馈、也可以联系我。

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
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
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

设置canvas随机字符串
--luna_cavans_random_str="A3B4"
原理就是把产生的toDataURL拼接个随机字符串


设置user-agent;
--luna_user_agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.48 Safari/537.36"

设置Platform;
--luna_platform="Win32"
--Linux armv81  MacIntel  Win32 "" Linux x86_64 iPhone Unsupported platform


设置时区；
--luna_timezone=""
   "Pacific/Midway",
   "Pacific/Honolulu",
   "America/Anchorage",
   "America/Los_Angeles",
   "America/Vancouver",
   "America/Tijuana",
   "America/Phoenix",
   "America/Chihuahua",
   "America/Denver",
   "America/Edmonton",
   "America/Mazatlan",
   "America/Regina",
   "America/Costa_Rica",
   "America/Chicago",
   "America/Mexico_City",
   "America/Tegucigalpa",
   "America/Winnipeg",
   "Pacific/Easter",
   "America/Bogota",
   "America/Lima",
   "America/New_York",
   "America/Toronto",
   "America/Caracas",
   "America/Barbados",
   "America/Halifax",
   "America/Manaus",
   "America/Santiago",
   "America/St_Johns",
   "America/Araguaina",
   "America/Argentina/Buenos_Aires",
   "America/Argentina/San_Luis",
   "America/Montevideo",
   "America/Santiago",
   "America/Sao_Paulo",
   "America/Godthab",
   "Atlantic/South_Georgia",
   "Atlantic/Cape_Verde",
   "Atlantic/Azores",
   "Atlantic/Reykjavik",
   "Atlantic/St_Helena",
   "Africa/Casablanca",
   "Atlantic/Faroe",
   "Europe/Dublin",
   "Europe/Lisbon",
   "Europe/London",
   "Europe/Amsterdam",
   "Europe/Belgrade",
   "Europe/Berlin",
   "Europe/Bratislava",
   "Europe/Brussels",
   "Europe/Budapest",
   "Europe/Copenhagen",
   "Europe/Ljubljana",
   "Europe/Madrid",
   "Europe/Malta",
   "Europe/Oslo",
   "Europe/Paris",
   "Europe/Prague",
   "Europe/Rome",
   "Europe/Stockholm",
   "Europe/Sarajevo",
   "Europe/Tirane",
   "Europe/Vaduz",
   "Europe/Vienna",
   "Europe/Warsaw",
   "Europe/Zagreb",
   "Europe/Zurich",
   "Africa/Windhoek",
   "Africa/Lagos",
   "Africa/Brazzaville",
   "Africa/Cairo",
   "Africa/Harare",
   "Africa/Maputo",
   "Africa/Johannesburg",
   "Europe/Kaliningrad",
   "Europe/Athens",
   "Europe/Bucharest",
   "Europe/Chisinau",
   "Europe/Helsinki",
   "Europe/Istanbul",
   "Europe/Kiev",
   "Europe/Riga",
   "Europe/Sofia",
   "Europe/Tallinn",
   "Europe/Vilnius",
   "Asia/Amman",
   "Asia/Beirut",
   "Asia/Jerusalem",
   "Africa/Nairobi",
   "Asia/Baghdad",
   "Asia/Riyadh",
   "Asia/Kuwait",
   "Europe/Minsk",
   "Europe/Moscow",
   "Asia/Tehran",
   "Europe/Samara",
   "Asia/Dubai",
   "Asia/Tbilisi",
   "Indian/Mauritius",
   "Asia/Baku",
   "Asia/Yerevan",
   "Asia/Kabul",
   "Asia/Karachi",
   "Asia/Aqtobe",
   "Asia/Ashgabat",
   "Asia/Oral",
   "Asia/Yekaterinburg",
   "Asia/Calcutta",
   "Asia/Colombo",
   "Asia/Katmandu",
   "Asia/Omsk",
   "Asia/Almaty",
   "Asia/Dhaka",
   "Asia/Novosibirsk",
   "Asia/Rangoon",
   "Asia/Bangkok",
   "Asia/Jakarta",
   "Asia/Krasnoyarsk",
   "Asia/Novokuznetsk",
   "Asia/Ho_Chi_Minh",
   "Asia/Phnom_Penh",
   "Asia/Vientiane",
   "Asia/Shanghai",
   "Asia/Hong_Kong",
   "Asia/Kuala_Lumpur",
   "Asia/Singapore",
   "Asia/Manila",
   "Asia/Taipei",
   "Asia/Ulaanbaatar",
   "Asia/Makassar",
   "Asia/Irkutsk",
   "Asia/Yakutsk",
   "Australia/Perth",
   "Australia/Eucla",
   "Asia/Seoul",
   "Asia/Tokyo",
   "Asia/Jayapura",
   "Asia/Sakhalin",
   "Asia/Vladivostok",
   "Asia/Magadan",
   "Australia/Darwin",
   "Australia/Adelaide",
   "Pacific/Guam",
   "Australia/Brisbane",
   "Australia/Hobart",
   "Australia/Sydney",
   "Asia/Anadyr",
   "Pacific/Port_Moresby",
   "Asia/Kamchatka",
   "Pacific/Fiji",
   "Pacific/Majuro",
   "Pacific/Auckland",
   "Pacific/Tongatapu",
   "Pacific/Apia",
   "Pacific/Kiritimati"

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

答：是的

2、目前支持Linux 系统 or 苹果m1,m2芯片吗？

答:暂时不支持、

3、涉及服务器授权吗？

答：否、完全设备授权、离线的授权。

4、有体积更小的浏览器么？

答：无、本身就是抗指纹的，如果精简版 不利于抗指纹。

5、为什么我测试基于视觉时候发现，出现bug

答：下载代码后不要修我的项目名字

6、第三封库可以用的么，如Selenium Pyppeteer Playwright 。

答：可以，但是强烈不建议，因为这样基本上等于阉割了抗指纹的部分核心功能。

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
