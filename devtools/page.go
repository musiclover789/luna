package devtools

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/dom"
	"github.com/musiclover789/luna/base_devtools/emulation"
	"github.com/musiclover789/luna/base_devtools/input"
	"github.com/musiclover789/luna/base_devtools/runtime"
	"github.com/musiclover789/luna/protocol"
	"github.com/musiclover789/luna/script"
	"github.com/tidwall/gjson"
	"math/rand"
	"time"
)

type Page struct {
	DevToolsConn         *protocol.DevToolsConn
	PageID               string //就是所谓的窗口ID
	CurrentURL           string //当前链接
	Title                string //title
	WebSocketDebuggerUrl string
	Port                 int //端口
	Alive                bool
	ImgPath              string //存放图片的基础目录
}

func NewPage(devToolsConn *protocol.DevToolsConn, pageID, currentURL, title, webSocketDebuggerUrl string, port int, ImgPath string) *Page {
	return &Page{
		DevToolsConn:         devToolsConn,
		PageID:               pageID,
		CurrentURL:           currentURL,
		Title:                title,
		WebSocketDebuggerUrl: webSocketDebuggerUrl,
		Port:                 port,
		Alive:                true,
		ImgPath:              ImgPath,
	}
}

func (p *Page) Close() {
	//设置成失效
	p.Alive = false

	//在关闭页面
	protocol.ClosePageEndpoint(p.Port, p.PageID)

	//先关闭socket
	p.DevToolsConn.Close()

}

func (p Page) GetDocument() gjson.Result {
	return dom.GetDocument(p.DevToolsConn, 1, false)
}

func (p Page) GetHtml() gjson.Result {
	return dom.GetOuterHTML(p.DevToolsConn, p.GetDocument().Get("result.root.nodeId").Int())
}

func (p Page) GetHTMLWithMap() map[string]interface{} {
	return dom.GetOuterHTMLWithMap(p.DevToolsConn, p.GetDocument().Get("result.root.nodeId").Int())
}

func (p Page) SetHtml(html string) {
	nodeId := p.GetDocument().Get("result.root.nodeId").Int()
	dom.SetOuterHTML(p.DevToolsConn, nodeId, html)
}

//运行js 同步

func (p Page) RunJSSync(js string, timeout time.Duration) (error, gjson.Result) {
	return runtime.EvaluateWithResultSync(p.DevToolsConn, js, timeout)
}

//运行js 异步

func (p Page) RunJS(js string) {
	runtime.Evaluate(p.DevToolsConn, js)
}

//鼠标移动

func (p Page) SimulateMouseMoveOnPage(startX, startY, endX, endY float64) {
	speed := p.DevToolsConn.GetSpeed()
	if speed != 0 {
		p.DevToolsConn.ReduceSpeed(0)
	}
	input.SimulateMoveMouse(p.DevToolsConn, startX, startY, endX, endY)
	p.DevToolsConn.ReduceSpeed(speed)
}

//鼠标点击

func (p Page) SimulateMouseClickOnPage(x, y float64) {
	input.SimulateMouseClick(p.DevToolsConn, x, y)
}

//滚轮

func (p Page) SimulateMouseScrollOnPage(x, y float64, totalDistance int, direction input.Direction) {
	input.SimulateMouseScroll(p.DevToolsConn, x, y, totalDistance, direction)
}

func (p Page) ScrollToTargetImagePosition(x, y float64, direction input.Direction, smallImgPath string, matchScore float64, timeout time.Duration) (error, bool) {
	return input.ScrollMouseToTargetImage(p.DevToolsConn, x, y, 2000, direction, smallImgPath, p.ImgPath, BrowserGlobal.DevicePixelRatio, matchScore, timeout)
}

//键盘

func (p Page) SimulateKeyboardInputOnPage(text string) {
	input.SimulateKeyboardInput(p.DevToolsConn, text)
}

//等等图片相似度

func (p Page) WaitForMatchOnPageSync(smallImgPath string, matchScore float64, timeout time.Duration) (error, bool) {
	//初始化、获取浏览器屏幕信息
	initBrowserScreen(p.DevToolsConn)
	fmt.Println("屏幕缩放", BrowserGlobal.DevicePixelRatio)
	return input.WaitForMatchSync(p.DevToolsConn, smallImgPath, p.ImgPath, matchScore, BrowserGlobal.DevicePixelRatio, timeout)
}

//func (p Page) WaitForMatchMoreOnPageSync(targetCoordinatesItems[] *input.TargetCoordinatesItem, fn func(imageCoordinates *input.ImageCoordinates)) (error, bool) {
//	//初始化、获取浏览器屏幕信息
//	initBrowserScreen(p.DevToolsConn)
//	fmt.Println("屏幕缩放", BrowserGlobal.DevicePixelRatio)
//	imageCoordinates := &input.ImageCoordinates{}
//	fn(imageCoordinates)
//	return input.WaitForMatchSync(p.DevToolsConn, smallImgPath, p.ImgPath, matchScore, BrowserGlobal.DevicePixelRatio, timeout)
//}

//计算图片相似度

func (p Page) ImageSimilarity(smallImgPath string, timeout time.Duration) (error, *input.ImageCoordinates) {
	return input.GetSmallImageCoordinatesWithMargin(p.DevToolsConn, smallImgPath, p.ImgPath, BrowserGlobal.DevicePixelRatio, 1, 1, 1, 1, timeout)
}

//计算相似度

func (p Page) SimilarityWithMargin(smallImgPath string, leftMargin, rightMargin, topMargin, bottomMargin float64, timeout time.Duration) (error, *input.ImageCoordinates) {
	return input.GetSmallImageCoordinatesWithMargin(p.DevToolsConn, smallImgPath, p.ImgPath, BrowserGlobal.DevicePixelRatio, leftMargin, rightMargin, topMargin, bottomMargin, timeout)
}

//设置窗口大小
func (p Page) SetViewportSize(width, height int64) {
	emulation.SetDeviceMetricsOverride(p.DevToolsConn, width, height)
}

//获取位置信息

func (p Page) GetElementPositionByXpathOnPage(selector string) (err error, randomX, randomY float64) {
	js := script.GetElementPositionByXpath(selector)
	err, jsResult := p.RunJSSync(js, time.Minute)
	if err == nil {
		ok := jsResult.Get("result.result.preview.properties.0.value").Bool()
		top := jsResult.Get("result.result.preview.properties.1.value").Float()
		left := jsResult.Get("result.result.preview.properties.2.value").Float()
		width := jsResult.Get("result.result.preview.properties.3.value").Float()
		height := jsResult.Get("result.result.preview.properties.4.value").Float()
		x, y := getRandomPoint(top, left, width, height)
		if ok {
			return nil, x, y
		} else {
			return fmt.Errorf("未找到元素"), x, y
		}
	}
	return fmt.Errorf("未找到元素"), 0, 0
}

func (p Page) GetElementPositionByCssOnPage(selector string) (err error, randomX, randomY float64) {
	js := script.GetElementPositionByCss(selector)
	err, jsResult := p.RunJSSync(js, time.Minute)
	if err == nil {
		ok := jsResult.Get("result.result.preview.properties.0.value").Bool()
		top := jsResult.Get("result.result.preview.properties.1.value").Float()
		left := jsResult.Get("result.result.preview.properties.2.value").Float()
		width := jsResult.Get("result.result.preview.properties.3.value").Float()
		height := jsResult.Get("result.result.preview.properties.4.value").Float()
		x, y := getRandomPoint(top, left, width, height)
		if ok {
			return nil, x, y
		} else {
			return fmt.Errorf("未找到元素"), x, y
		}
	}
	return fmt.Errorf("未找到元素"), 0, 0
}

func getRandomPoint(top, left, width, height float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())

	// 计算矩形框的右下角坐标
	right := left + width
	bottom := top + height

	// 生成随机点的 x 和 y 坐标
	x := rand.Float64()*(right-left) + left
	y := rand.Float64()*(bottom-top) + top

	return x, y
}
