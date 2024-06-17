package main

//
//import (
//	"fmt"
//	"github.com/musiclover789/luna/reverse_proxy"
//	"net/url"
//	"os"
//	"sync"
//)
//
//func main() {
//	var wg1 sync.WaitGroup
//	wg1.Add(1)
//	if len(os.Args) < 2 {
//		fmt.Println("-1")
//		return
//	}
//	command := os.Args[1]
//	switch command {
//	case "start":
//		if len(os.Args) < 3 {
//			fmt.Println("-1")
//			return
//		}
//		message := os.Args[2]
//		result := startService(message)
//		fmt.Println(result) // 返回值输出到标准输出
//	default:
//		fmt.Println("-1")
//	}
//	wg1.Wait()
//}
//
//func startService(message string) string {
//	proxy := message
//	if len(proxy) > 0 {
//		proxyURL, err := url.Parse(proxy)
//		if err != nil {
//			return "-1"
//		}
//		pwd, _ := proxyURL.User.Password()
//		proxyServer := reverse_proxy.NewProxyServer(proxyURL.Scheme, proxyURL.Hostname(), proxyURL.Port(), proxyURL.User.Username(), pwd)
//		// 启动代理服务器
//		proxy_port, err := proxyServer.Start()
//		if err != nil {
//			return "-1"
//		}
//		if err != nil {
//			return "-1"
//		}
//		return proxy_port
//	}
//	return "-1"
//}
