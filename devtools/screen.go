package devtools

import (
	"errors"
	"luna/base_devtools/runtime"
	"luna/protocol"
	"luna/script"
	"time"
)

type WindowSize struct {
	Width  int
	Height int
}

/***
这里记录、屏幕尺寸、缩放等级等全局信息
*/

type BrowserContext struct {
	DevicePixelRatio  float64 //缩放因子
	ScreenWidth       float64   //屏幕的宽度
	ScreenHeight      float64   //返回屏幕的高度，以像素为单位
	ScreenAvailWidth  float64   //返回屏幕的可用宽度，即去除操作系统工具栏等后的宽度，以像素为单位
	ScreenAvailHeight float64   //返回屏幕的可用高度，即去除操作系统工具栏等后的高度，以像素为单位
}

var BrowserGlobal BrowserContext

func initBrowserScreen(devtoolsRoot *protocol.DevToolsConn) {
	err, jsonStr := runtime.EvaluateWithResultSync(devtoolsRoot, script.ScreenInfo(), time.Minute)
	if err != nil {
		errors.New("BrowserScreen error")
	}
	width := jsonStr.Get("result.result.preview.properties.0.value").Float()
	height := jsonStr.Get("result.result.preview.properties.1.value").Float()
	availWidth := jsonStr.Get("result.result.preview.properties.2.value").Float()
	availHeight := jsonStr.Get("result.result.preview.properties.3.value").Float()
	devicePixelRatio := jsonStr.Get("result.result.preview.properties.4.value").Float()

	BrowserGlobal.ScreenWidth = width
	BrowserGlobal.ScreenHeight = height
	BrowserGlobal.ScreenAvailWidth = availWidth
	BrowserGlobal.ScreenAvailHeight = availHeight
	BrowserGlobal.DevicePixelRatio = devicePixelRatio
}

func GetBrowserScreen(devtoolsRoot *protocol.DevToolsConn)*BrowserContext  {
	err, jsonStr := runtime.EvaluateWithResultSync(devtoolsRoot, script.ScreenInfo(), time.Minute)
	if err != nil {
		errors.New("BrowserScreen error")
	}
	width := jsonStr.Get("result.result.preview.properties.0.value").Float()
	height := jsonStr.Get("result.result.preview.properties.1.value").Float()
	availWidth := jsonStr.Get("result.result.preview.properties.2.value").Float()
	availHeight := jsonStr.Get("result.result.preview.properties.3.value").Float()
	devicePixelRatio := jsonStr.Get("result.result.preview.properties.4.value").Float()

	return &BrowserContext{
		ScreenWidth: width,
		ScreenHeight: height,
		ScreenAvailWidth: availWidth,
		ScreenAvailHeight: availHeight,
		DevicePixelRatio: devicePixelRatio,
	}
}