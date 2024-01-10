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

// 如果对这种写法不熟悉,可以直接用main函数替代,也就是把里面的代码全部粘贴到main函数里面是一样的.
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
	//chromiumPath := "C:\\src\\chromedev\\chromium\\src\\out\\Default/chrome.exe"
	chromiumPath := "/Users/hongyuji/Documents/workspace/golang/Chromium.app/Contents/MacOS/Chromium"
	browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
		//设置缓存目录,
		CachePath: luna_utils.CreateCacheDirInSubDir("/Users/hongyuji/Documents/workspace/golang/cache"),
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
