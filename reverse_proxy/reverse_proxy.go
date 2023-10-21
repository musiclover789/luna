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
	"time"
)

func StartProxyServer(proxyType, host, port, user, password string) (string, error) {
	// 设置代理IP
	var auth *proxy.Auth
	if user != "" && password != "" {
		auth = &proxy.Auth{
			User:     user,
			Password: password,
		}
	}

	var dialFunc func(string, string) (net.Conn, error)

	switch proxyType {
	case "http":
		proxyURL := fmt.Sprintf("http://%s:%s@%s:%s", user, password, host, port)
		dialer, err := proxyclient.NewProxyClient(proxyURL)
		if err != nil {
			return "", fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		dialFunc = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	case "https":
		proxyURL := fmt.Sprintf("https://%s:%s@%s:%s", user, password, host, port)
		dialer, err := proxyclient.NewProxyClient(proxyURL)
		if err != nil {
			return "", fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		dialFunc = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	case "socks5":
		dialer, err := proxy.SOCKS5("tcp", net.JoinHostPort(host, port), auth, proxy.Direct)
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
		Dial: dialFunc,
	}

	// 创建 ProxyServer
	proxyServer := goproxy.NewProxyHttpServer()
	proxyServer.Tr = tr

	// 获取一个随机端口
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", fmt.Errorf("failed to listen on a random port: %w", err)
	}
	port = strconv.Itoa(l.Addr().(*net.TCPAddr).Port)

	// 启动 ProxyServer
	ch := make(chan error)
	go func() {
		log.Printf("Proxy server is listening on :%s", port)
		ch <- http.Serve(l, proxyServer)
	}()

	// 等待代理服务器启动完成
	select {
	case err := <-ch:
		return "", err
	case <-time.After(time.Second):
	}

	return port, nil
}

