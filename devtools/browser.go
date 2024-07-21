package devtools

import (
	"errors"
	"fmt"
	"github.com/musiclover789/luna/base_devtools/browser"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/base_devtools/target"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"github.com/musiclover789/luna/reverse_proxy"
	"github.com/tidwall/gjson"
	"sync"
)

type BrowserOptions struct {
	CachePath   string //缓存目录
	ImgPath     string //存放目标图片的基础目录
	Headless    bool
	ProxyStr    string
	Fingerprint []string
	WindowSize  *WindowSize
	// 其他可选参数...
}

type Browser struct {
	ChromiumPath   string         //可执行文件存放路径
	Port           int            //端口
	BrowserOptions BrowserOptions //可选参数
	SessionManager *protocol.DevtoolsRoot
	Session        *protocol.Session
	First          bool //是不是第一次打开窗口
	mu             sync.Mutex
	Pages          []*Page
	ImgPath        string                     //存放目标图片的基础目录
	Proxy          *reverse_proxy.ProxyServer //proxy对象
	TargetID       string
	Pid            int
}

func initBrowser(chromiumPath string, options *BrowserOptions) (int, *reverse_proxy.ProxyServer, int) {
	if options != nil {
		return luna_utils.StartChromiumWithUserDataDir(chromiumPath, options.CachePath, &options.ProxyStr, options.Headless, func() (bool, int, int) {
			if options.WindowSize == nil {
				return false, 0, 0
			}
			return true, options.WindowSize.Width, options.WindowSize.Height
		}, options.Fingerprint[:]...)
	} else {
		return luna_utils.StartChromiumWithUserDataDir(chromiumPath, "", nil, false, nil)
	}
	return 0, nil, -1
}

var mutex sync.Mutex

func NewBrowser(chromiumPath string, options *BrowserOptions) (error, *Browser) {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	port, proxy, pid := initBrowser(chromiumPath, options)
	droot := protocol.NewDevtoolsRoot(port)
	if port == 0 || port == -1 {
		return errors.New("NewBrowser异常-获取不到正确到端口"), nil
	}
	err, firstConn, targetID := droot.FirstConn()
	if err != nil {
		return errors.New("NewBrowser-FirstConn异常-连接不到端口"), nil
	}
	imgPath := ""
	if options != nil {
		imgPath = options.ImgPath
	}
	return nil, &Browser{
		ChromiumPath:   chromiumPath,
		Port:           port,
		SessionManager: droot,
		First:          true,
		Session:        firstConn,
		ImgPath:        imgPath,
		Proxy:          proxy,
		TargetID:       targetID,
		Pid:            pid,
	}
}

func (browser *Browser) isFirst() bool {
	browser.mu.Lock()
	defer browser.mu.Unlock()
	if browser.First {
		browser.First = false
		return true
	}
	return false
}

func (browser *Browser) addPage(p *Page) {
	browser.Pages = append(browser.Pages, p)
}

func (browser *Browser) RemovePage(p *Page) {
	for i, page := range browser.Pages {
		if page == p {
			browser.Pages = append(browser.Pages[:i], browser.Pages[i+1:]...)
			break
		}
	}
}

func (browser *Browser) OpenPage(url string) (error, *Page) {
	//打开网址、返回页面ID
	frameId := target.CreateTarget(browser.Session, url).Get("result.targetId").String()
	resultEndpoint, err := protocol.GetPageEndpointByID(browser.Port, frameId)
	if err != nil {
		return err, nil
	}
	//返回这个page对对象
	err, pageConn := protocol.CreteSession(resultEndpoint.WebSocketDebuggerURL)
	if err != nil {
		return err, nil
	}
	pageObj := NewPage(pageConn, frameId, resultEndpoint.URL, resultEndpoint.Title, resultEndpoint.WebSocketDebuggerURL, browser.Port, browser.ImgPath)
	browser.addPage(pageObj)
	return nil, pageObj
}

func (browser *Browser) OpenPageAndListen(url string, fns ...func(Session *protocol.Session)) (error, *Page) {
	//打开网址、返回页面ID
	frameId := target.CreateTarget(browser.Session, "").Get("result.targetId").String()
	resultEndpoint, err := protocol.GetPageEndpointByID(browser.Port, frameId)
	if err != nil {
		return err, nil
	}
	//返回这个page对对象
	err, pageConn := protocol.CreteSession(resultEndpoint.WebSocketDebuggerURL)
	if err != nil {
		return err, nil
	}
	for _, fn := range fns {
		fn(pageConn)
	}
	page.PageNavigate(pageConn, url)
	pageObj := NewPage(pageConn, frameId, resultEndpoint.URL, resultEndpoint.Title, resultEndpoint.WebSocketDebuggerURL, browser.Port, browser.ImgPath)
	browser.addPage(pageObj)
	return nil, pageObj
}

func (browser *Browser) GetPages() (error, []*Page) {
	/***
	1、循环当前所有有多少页面
	2、检测当前的浏览器对象里面是否有、如果有就可以叠加、如果没有的可以删除掉、然后返回
	*/
	//先筛查出或者的元素;
	retainedPages := []*Page{}
	for _, po := range browser.Pages {
		//说明这个页面还建在
		if po.Alive {
			retainedPages = append(retainedPages, po)
		}
	}
	browser.Pages = retainedPages
	//然后我们在同步元素
	ers, err := protocol.GetPageEndpoints(browser.Port)
	if err != nil {
		return err, nil
	}
	/***
	ers实际的页面
	browser.Pages缓存的页面。
	循环缓存的页面，将实际已经没有的从缓存中去掉
	*/
	for _, er := range browser.Pages {
		bl := true
		for _, po := range *ers {
			//说明这个页面还建在
			if er.PageID == po.ID {
				bl = false
			}
		}
		//说明缓存总有,但是页面没有
		if bl {
			browser.RemovePage(er)
		}
	}

	for _, er := range *ers {
		bl := true
		for _, po := range browser.Pages {
			//说明这个页面还建在
			if er.ID == po.PageID {
				po.CurrentURL = er.URL
				po.Title = er.Title
				bl = false
			}
		}
		//说明实际有但是你的browser没有
		if bl {
			err, pageConn := protocol.CreteSession(er.WebSocketDebuggerURL)
			if err == nil {
				pageObj := NewPage(pageConn, er.ID, er.URL, er.Title, er.WebSocketDebuggerURL, browser.Port, browser.ImgPath)
				browser.addPage(pageObj)
			} else {
				fmt.Println("新建页面发送错误", err)
			}
		}
	}
	return nil, browser.Pages
}

func (browser *Browser) SwitchPage(currentPage *Page) gjson.Result {
	return page.BringToFront(currentPage.Session)
}

func (browser *Browser) SwitchPageAndListen(currentPage *Page, fns ...func(Session *protocol.Session)) gjson.Result {
	for _, fn := range fns {
		fn(currentPage.Session)
	}
	page.PageNavigate(currentPage.Session, currentPage.CurrentURL)
	return page.BringToFront(currentPage.Session)
}

func (b *Browser) Close() {
	//for _, p := range b.GetPages() {
	//	p.Close()
	//}
	proxy := b.Proxy
	if proxy != nil {
		// 停止代理服务器
		if err := proxy.Stop(); err != nil {
			fmt.Println("Failed to stop proxy server:", err.Error())
		}
	}
	browser.CloseBrowser(b.Session)
}

func (b *Browser) SetWindowBounds(left, top, width, height int) {
	browser.SetWindowBounds(b.Session, left, top, width, height)
}
