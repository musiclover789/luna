package test_case
//
//import (
//	"fmt"
//	"luna/luna_devtools/input"
//	"luna/luna_devtools/page"
//	"luna/luna_network"
//	"luna/luna_utils"
//	"path/filepath"
//	"testing"
//	"time"
//)
//
///***
//官方文档
//https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
//*/
//
//func TestQichacha(t *testing.T) {
//
//	//启动前先杀死其他的chromium进程;为了防止程序以及停止但是依然在内存中贮存;
//	//他会根据你的系统不同,使用命令行的命令进行杀死进程
//	luna_utils.KillProcess()
//
//	//如果你的路径有空格,或者权限不足,你可以创建快捷方式放在这个目录就可以了;mac和linux就是软连接,一个意思.
//	chromiumPath := "/Users/hongyuji/Documents/workspace/golang/Chromium.app/Contents/MacOS/Chromium"
//	//这个是存储你的浏览器缓存文件的目录
//	chromiumCachePath := "/Users/hongyuji/Documents/workspace/golang/cache"
//
//	/***
//	如果你需要使用代理参考 case_proxy_test.go
//	*/
//	/***
//	你也可以不要传入这个函数luna_utils.CreateCacheDirInSubDir(chromiumCachePath)
//	直接这样写.
//	port := luna_utils.StartChromiumWithUserDataDir(chromiumPath, chromiumCachePath,&proxyStr)
//	这样的意思是直接使用chromiumCachePath作为缓存目录、如果你使用了这个函数唯一的区别是,他会每次运行这行代码的时候都创建一个临时目录
//	在你给定的chromiumCachePath目录下.
//	*/
//	chromiumCachePath = luna_utils.CreateCacheDirInSubDir(chromiumCachePath)
//	port := luna_utils.StartChromiumWithUserDataDir(chromiumPath, chromiumCachePath, nil, false)
//
//	/***
//		这个函数只是和浏览器建立链接而已;没什么特别的.
//	**/
//
//	conn := luna_network.NewDevToolsConn(port)
//
//	conn.ShowLog(false)
//
//	conn.ShowLogJson(false)
//
//	err, coefficient := page.CalculateScalingFactorSync(conn, time.Second*160)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("缩放因子是:", coefficient, "因为你实际截图的图片大小和你的屏幕css像素尺寸是 缩放因子倍数的关系")
//
//	/**
//	下面的每个命令、执行完1秒内，不运行其他命令运行。
//	*/
//	//conn.SetBlockingTimeout(time.Second)
//	/***
//	打开一个页面、
//	*/
//	page.PageNavigate(conn, "https://www.qcc.com")
//	/***
//	判断这个页面是否是正常打开
//	*/
//	/**
//		TODO :
//		1、滚轮的功能、比如我想看第二屏 如何操作呢？
//		2、比如页面打开了5个标签页、我当前到底在第几个?我如何可以关闭其他的？ 包括当前到底有几个标签页，我如何切换？
//		3、我要查看当前的页面的源代码、则么查看。
//	 */
//
//	basePath := "/Users/hongyuji/Documents/workspace/golang/luna/luna_test_img/qichacha_test_img/"
//	err, input_obj := input.NewInput(conn, basePath, 2)
//	err, b := input_obj.WaitForMatchSync(filepath.Join(basePath, "check_home_01.png"), 0.5, time.Second*60) //判断已经到达了首页
//	if b && err == nil {
//		fmt.Println("成功打开页面")
//		page.PageNavigate(conn,"view-source:https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Viewport")
//		//err, step01 := input_obj.GetSmallImageCoordinatesOnPageWithMargin(filepath.Join(basePath, "step01.png"), 10, 50, 5, 5, time.Second*60)
//		//if err == nil && step01.MatchScore > 0.8 {
//		//	runtime.Evaluate(conn,luna_script.ShowMousePosition())
//		//	input_obj.SimulateMoveMouse(-1, -1, step01.RandomX, step01.RandomY)
//		//	input_obj.SimulateMouseClick(step01.RandomX, step01.RandomY)
//		//	input_obj.SimulateKeyboardInput("广联达")
//		//	//点击按钮
//		//	//conn.SetBlockingTimeout(time.Second)
//		//	err,step02:=input_obj.GetSmallImageCoordinatesOnPage(filepath.Join(basePath, "step02.png"),time.Second*60)
//		//	if err == nil && step02.MatchScore > 0.8{
//		//		input_obj.SimulateMouseClick(step02.RandomX,step02.RandomY)
//		//	}
//		//} else {
//		//	fmt.Println(err)
//		//}
//	} else {
//		fmt.Println(err)
//	}
//
//	time.Sleep(time.Hour)
//	time.Sleep(1 * time.Hour)
//}
