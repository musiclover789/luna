package reverse_proxy

import (
	"errors"
	"fmt"
	"github.com/gamexg/proxyclient"
	"golang.org/x/net/proxy"
	"gopkg.in/elazarl/goproxy.v1"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// ProxyServer 结构体包含代理服务器相关的信息和操作
type ProxyServer struct {
	server       *goproxy.ProxyHttpServer
	listener     net.Listener
	proxyType    string
	host         string
	port         string
	user         string
	password     string
	listenerLock sync.Mutex
}

// NewProxyServer 创建一个新的代理服务器实例
func NewProxyServer(proxyType, host, port, user, password string) *ProxyServer {
	return &ProxyServer{
		server:    goproxy.NewProxyHttpServer(),
		proxyType: proxyType,
		host:      host,
		port:      port,
		user:      user,
		password:  password,
	}
}

// Start 启动代理服务器
func (p *ProxyServer) Start() (string, error) {
	p.listenerLock.Lock()
	defer p.listenerLock.Unlock()

	if p.listener != nil {
		return "", errors.New("proxy server is already running")
	}

	// 设置代理IP
	var auth *proxy.Auth
	if p.user != "" && p.password != "" {
		auth = &proxy.Auth{
			User:     p.user,
			Password: p.password,
		}
	}

	var dialFunc func(string, string) (net.Conn, error)

	switch p.proxyType {
	case "http":
		proxyURL := fmt.Sprintf("http://%s:%s@%s:%s", p.user, p.password, p.host, p.port)
		dialer, err := proxyclient.NewProxyClient(proxyURL)
		if err != nil {
			return "", fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		dialFunc = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	case "https":
		proxyURL := fmt.Sprintf("https://%s:%s@%s:%s", p.user, p.password, p.host, p.port)
		dialer, err := proxyclient.NewProxyClient(proxyURL)
		if err != nil {
			return "", fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		dialFunc = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	case "socks5":
		dialer, err := proxy.SOCKS5("tcp", net.JoinHostPort(p.host, p.port), auth, proxy.Direct)
		if err != nil {
			return "", fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		dialFunc = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	default:
		return "", errors.New("invalid proxy type")
	}

	// 创建自定义的 Transport
	tr := &http.Transport{
		Dial:                dialFunc,
		MaxIdleConnsPerHost: 100, // 连接池大小
	}

	// 设置 ProxyServer 的 Transport
	p.server.Tr = tr

	// 获取一个随机端口
	var err error
	p.listener, err = net.Listen("tcp", ":0")
	if err != nil {
		return "", fmt.Errorf("failed to listen on a random port: %w", err)
	}
	p.port = strconv.Itoa(p.listener.Addr().(*net.TCPAddr).Port)

	// 启动 ProxyServer
	go func() {
		log.Printf("Proxy server is listening on :%s", p.port)
		err := http.Serve(p.listener, p.server)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Proxy server error: %v", err)
		}
	}()

	return p.port, nil
}

// Stop 停止代理服务器
func (p *ProxyServer) Stop() error {
	p.listenerLock.Lock()
	defer p.listenerLock.Unlock()

	if p.listener != nil {
		err := p.listener.Close()
		if err != nil {
			return err
		}
		log.Println("Proxy server stopped")
		p.listener = nil
		return nil
	}
	return errors.New("proxy server is not running")
}
