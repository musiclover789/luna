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
	"math"
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

// 上传文件 selector:css选择器;需要注意的是 这个要是type=file的元素的css选择器,  files:需要上传的文件路径
func (p Page) UploadFiles(selector string, files []string) {
	nodeId := dom.QuerySelector(p.DevToolsConn, p.GetDocument().Get("result.root.nodeId").Int(), selector).Get("result.nodeId").Int()
	dom.SetFileInputFiles(p.DevToolsConn, nodeId, files)
}

// 运行js 同步
// 执行并等待结果反馈
func (p Page) RunJSSync(js string, timeout time.Duration) (error, gjson.Result) {
	return runtime.EvaluateWithResultSync(p.DevToolsConn, js, timeout)
}

// 运行js 异步
// 执行，不关心结果反馈，或者不关心结果什么时候反馈
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

/*
**
模拟拖拽
*/
func (p Page) SimulateDrag(startX, startY, endX, endY float64) {
	input.SimulateMousePressed(p.DevToolsConn, startX, startY)
	input.SimulateMoveMouse(p.DevToolsConn, startX, startY, endX, endY)
	input.SimulateMouseReleased(p.DevToolsConn, endX, endY)
}

//鼠标移动 //从当前鼠标所在位置 移动到目标位置

func (p Page) SimulateMouseMoveToTarget(endX, endY float64) error {
	err, result := p.RunJSSync(script.JSGetRandomCoordinates(), time.Minute)
	if err == nil {
		result = result.Get("result.result.value")
		if !result.Exists() {
			return fmt.Errorf("未找到目标元素")
		}
		result = gjson.Parse(result.String())
		startX := result.Get("x")
		startY := result.Get("y")
		if !startX.Exists() || !startY.Exists() {
			return fmt.Errorf("未找到目标元素")
		}
		speed := p.DevToolsConn.GetSpeed()
		if speed != 0 {
			p.DevToolsConn.ReduceSpeed(0)
		}
		input.SimulateMoveMouse(p.DevToolsConn, startX.Float(), startY.Float(), endX, endY)
		p.DevToolsConn.ReduceSpeed(speed)
		return nil
	}
	return fmt.Errorf("未找到目标元素")
}

//鼠标移动到给定元素 边框内随机位置
/***
1、滚轮滚动到目标元素
2、获取元素坐标
3、钓鱼鼠标移动到目标坐标
*/
func (p Page) SimulateMouseMoveToElement(selector string) (err error, randomX, randomY float64) {
	randomX = 0
	randomY = 0
	err = p.SimulateScrollToElementBySelector(selector)
	if err != nil {
		return err, randomX, randomY
	}
	time.Sleep(1 * time.Second)
	err, randomX, randomY = p.GetElementPositionByCssOnPage(selector)
	if err != nil {
		return err, randomX, randomY
	}
	err = p.SimulateMouseMoveToTarget(randomX, randomY)
	if err != nil {
		return err, randomX, randomY
	}
	return nil, randomX, randomY
}

/**
鼠标点击 点击当前坐标,相对于当前可视窗口的,也就是说 如果你需要人类是可以看见的区域,也就是如果需要滚轮滑动才能看见
的区域,你需要先滚动到那个区域，否则坐标可能是错误的。
*/

func (p Page) SimulateMouseClickOnPage(x, y float64) {

	input.SimulateMouseClick(p.DevToolsConn, x, y)
}

//滚轮

/*
**
x, y float64, totalDistance int, direction input.Direction
开始位置x,y
滚动举例 totalDistance
input.Direction 滚动方向
*/
func (p Page) SimulateMouseScrollOnPage(x, y float64, totalDistance int, direction input.Direction) {
	input.SimulateMouseScroll(p.DevToolsConn, x, y, totalDistance, direction)
}

func (p Page) ScrollToTargetImagePosition(x, y float64, direction input.Direction, smallImgPath string, matchScore float64, timeout time.Duration) (error, bool) {
	return input.ScrollMouseToTargetImage(p.DevToolsConn, x, y, 2000, direction, smallImgPath, p.ImgPath, BrowserGlobal.DevicePixelRatio, matchScore, timeout)
}

/*
鼠标滚轮 滚动到给定元素位置、会默认在居中位置
至于上滚动 还是下滚动，这个是自动的 无需干涉
滚动结束后 返回
*/
func (p Page) SimulateScrollToElementBySelector(selector string) error {
	err, result := p.RunJSSync(script.JSGetElementPositionAndWindowViewportByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !result.Exists() {
			return fmt.Errorf("未找到目标元素")
		}
		value := gjson.Parse(resultStr)
		positionX := value.Get("elementPosition.x")
		positionY := value.Get("elementPosition.y")
		viewportPositionTop := value.Get("viewportPosition.top")
		viewportPositionBottom := value.Get("viewportPosition.bottom")
		if !result.Exists() {
			return fmt.Errorf("未找到目标元素")
		}
		//计算方向以及距离 --如果目标位置Y坐标 大于 可视窗口下限;向下移动
		if positionY.Int() > viewportPositionBottom.Int() {
			distance := (viewportPositionBottom.Int()-viewportPositionTop.Int())/2 + (positionY.Int() - viewportPositionBottom.Int())
			input.SimulateMouseScroll(p.DevToolsConn, float64(positionX.Int()), float64(positionY.Int()), int(distance), input.DOWN)
			//--如果目标位置Y坐标  小于 可视窗口上限;向上移动
		} else if positionY.Int() < viewportPositionTop.Int() {
			distance := (viewportPositionBottom.Int()-viewportPositionTop.Int())/2 + (viewportPositionTop.Int() - positionY.Int())
			input.SimulateMouseScroll(p.DevToolsConn, float64(positionX.Int()), float64(positionY.Int()), int(distance), input.UP)
		} else {
			//如果处在他们之间,理论上无需移动，但是我们还是会移动到中间位置--
			//首先我们应该判断 到底距离上面近还是下面近,方便我们控制方向
			if math.Abs(float64(positionY.Int()-viewportPositionTop.Int())) > math.Abs(float64(positionY.Int()-viewportPositionBottom.Int())) {
				fmt.Println("应该向上")
				//应该向上
				distance := ((viewportPositionBottom.Int() - viewportPositionTop.Int()) / 2) - (viewportPositionBottom.Int() - positionY.Int())
				input.SimulateMouseScroll(p.DevToolsConn, float64(positionX.Int()), float64(positionY.Int()), int(distance), input.DOWN)
			} else if math.Abs(float64(positionY.Int()-viewportPositionTop.Int())) < math.Abs(float64(positionY.Int()-viewportPositionBottom.Int())) {
				//应该向下
				distance := ((viewportPositionBottom.Int() - viewportPositionTop.Int()) / 2) - (positionY.Int() - viewportPositionTop.Int())
				input.SimulateMouseScroll(p.DevToolsConn, float64(positionX.Int()), float64(positionY.Int()), int(distance), input.UP)
			}
		}
		return nil
	}
	return fmt.Errorf("未找到目标元素")
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

// 设置窗口大小
func (p Page) SetViewportSize(width, height int64) {
	emulation.SetDeviceMetricsOverride(p.DevToolsConn, width, height, 0)
}

// 设置窗口大小
func (p Page) SetViewportSizeAndScale(width, height, scale int64) {
	emulation.SetDeviceMetricsOverride(p.DevToolsConn, width, height, scale)
}

/*
通过css选择器获取元素位置信息
返回的是元素范围内的随机坐标
*/
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

/*
通过xpath选择器获取元素位置信息
返回的是元素范围内的随机坐标
*/
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

// Node 表示节点信息的数据结构
type Node struct {
	NodeType      int64
	NodeName      string
	NodeValue     string
	TextContent   string
	HTMLContent   string
	CSSSelector   string
	XPathSelector string
}

/*
**
获取当前节点信息
*/
func (p Page) GetElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSGetElementBycss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/**
firstChild: 第一个子节点
*/

func (p Page) GetFirstChildElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSGetFirstChildElementByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/*
*
lastChild: 最后一个子节点
*/
func (p Page) GetLastChildElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSGetLastChildElementByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/*
*
nextSibling: 下一个兄弟节点
*/
func (p Page) GetNextSiblingElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSGetNextSiblingElementByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/**
previousSibling: 上一个兄弟节点
*/

func (p Page) GetPreviousSiblingElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSPreviousSiblingElementByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/***
parentNode: 父节点
*/

func (p Page) GetParentElementByCss(selector string) (error, Node) {
	err, result := p.RunJSSync(script.JSParentElementByCss(selector), time.Minute)
	if err == nil {
		resultStr := result.Get("result.result.value").String()
		if !gjson.Get(resultStr, "nodeType").Exists() {
			return fmt.Errorf("未找到元素"), Node{}
		}
		return nil, Node{
			NodeType:      gjson.Get(resultStr, "nodeType").Int(),
			NodeName:      gjson.Get(resultStr, "nodeName").String(),
			NodeValue:     gjson.Get(resultStr, "nodeValue").String(),
			TextContent:   gjson.Get(resultStr, "textContent").String(),
			HTMLContent:   gjson.Get(resultStr, "htmlContent").String(),
			CSSSelector:   gjson.Get(resultStr, "cssSelector").String(),
			XPathSelector: gjson.Get(resultStr, "xpathSelector").String(),
		}

	}
	return fmt.Errorf("未找到元素"), Node{}
}

/*
获取所有子节点
*/
func (p Page) GetAllChildElementByCss(selector string) (error, []Node) {
	err, result := p.RunJSSync(script.JSGetAllChildElementByCss(selector), time.Minute)
	if err == nil {
		if !result.Get("result.result.value").Exists() {
			return fmt.Errorf("未找到元素"), []Node{}
		}
		resultStr := result.Get("result.result.value").String()
		array := gjson.Parse(resultStr).Array()
		nodes := make([]Node, len(array))
		for i, item := range array {
			if !item.Get("nodeType").Exists() {
				return fmt.Errorf("未找到元素"), []Node{}
			}
			nodes[i] = Node{
				NodeType:      item.Get("nodeType").Int(),
				NodeName:      item.Get("nodeName").String(),
				NodeValue:     item.Get("nodeValue").String(),
				TextContent:   item.Get("textContent").String(),
				HTMLContent:   item.Get("htmlContent").String(),
				CSSSelector:   item.Get("cssSelector").String(),
				XPathSelector: item.Get("xpathSelector").String(),
			}
		}
		return nil, nodes
	}
	return fmt.Errorf("未找到元素"), []Node{}
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

func (p Page) GetCurrentURL() (error, string) {
	err, result := p.RunJSSync("window.location.href;", time.Minute)
	if err == nil {
		result := result.Get("result.result.value")
		if !result.Exists() {
			return fmt.Errorf("未找到元素"), ""
		}
		return nil, result.String()

	}
	return fmt.Errorf("未找到元素"), ""
}
