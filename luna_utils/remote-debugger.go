package luna_utils

import (
	"fmt"
	"github.com/musiclover789/luna/log"
	"github.com/musiclover789/luna/reverse_proxy"
	"log"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//--window-size=800,600

var StartChromiumWithUserDataDir = func(chromiumPath, userDataDirFullPath string, proxy *string, isHeadless bool, size func() (bool, int, int), customArgs ...string) (int, *reverse_proxy.ProxyServer) {
	// 获取随机端口
	port, err := getRandomPort()
	if err != nil {
		luna_log.LogError("failed to get random port: %v\n", err)
		return -1, nil
	}

	// 检查端口是否被占用
	if isPortOpen(port) {
		luna_log.LogError("port %d is already in use\n", port)
		return -1, nil
	}

	luna_log.Log("运行可执行文件的路径是:", chromiumPath)
	luna_log.Log("临时缓存目录是:", userDataDirFullPath)
	luna_log.Log("选用的临时端口是:", port)

	chromiumCmdArgs := []string{}

	if isHeadless {
		chromiumCmdArgs = append(chromiumCmdArgs, "--headless")
	}

	if size != nil {
		ok, Width, Height := size()
		if ok {
			chromiumCmdArgs = append(chromiumCmdArgs, strings.Join([]string{"--window-size=", strconv.Itoa(Width), ",", strconv.Itoa(Height)}, ""))
		}
	}

	for _, arg := range customArgs {
		if len(arg) > 0 {
			chromiumCmdArgs = append(chromiumCmdArgs, arg)
		}
	}

	chromiumCmdArgs = append(chromiumCmdArgs, "--remote-debugging-port="+strconv.Itoa(port))
	var proxyServer *reverse_proxy.ProxyServer
	if proxy != nil && len(*proxy) > 0 {
		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			luna_log.LogError("Failed to parse proxy URL: %v", err)
			return -1, nil
		}
		pwd, _ := proxyURL.User.Password()

		// 创建一个代理服务器实例
		proxyServer = reverse_proxy.NewProxyServer(proxyURL.Scheme, proxyURL.Hostname(), proxyURL.Port(), proxyURL.User.Username(), pwd)

		// 启动代理服务器
		proxy_port, err := proxyServer.Start()
		if err != nil {
			log.Fatalf("Failed to start proxy server: %v", err)
		}
		log.Printf("Proxy server started on port %s", port)

		if err != nil {
			luna_log.LogError("Failed to start proxy server: %v", err)
			return -1, nil
		}
		chromiumCmdArgs = append(chromiumCmdArgs, "--proxy-server=127.0.0.1:"+proxy_port)
	}
	fmt.Println("启动参数:", chromiumCmdArgs)
	if len(userDataDirFullPath) > 0 {
		chromiumCmdArgs = append(chromiumCmdArgs, "--user-data-dir="+userDataDirFullPath)
	}
	fmt.Println(chromiumCmdArgs)

	chromiumCmd := exec.Command(chromiumPath, chromiumCmdArgs...)

	err = chromiumCmd.Start()
	if err != nil {
		luna_log.LogError("Failed to start process: %v\n", err)
		return -1, nil
	}

	if chromiumCmd.ProcessState != nil && chromiumCmd.ProcessState.Exited() {
		luna_log.LogError("Failed to start process, exit code %d\n", chromiumCmd.ProcessState.ExitCode())
		return -1, nil
	}

	return port, proxyServer
}

var CreateCacheDirInSubDir = func(basePath string) string {
	// 组装完整的随机文件夹路径
	randFolderName := fmt.Sprintf("chromium_user_data_%d", time.Now().UnixNano())
	cacheDirFullPath := filepath.Join(basePath, randFolderName)
	// 创建随机文件夹
	if err := os.MkdirAll(cacheDirFullPath, 0700); err != nil {
		fmt.Printf("创建缓存目录失败：", err)
		return ""
	}
	return cacheDirFullPath
}

// 获取一个随机的未被占用的端口号
func getRandomPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// 检查指定的端口是否已经被占用
func isPortOpen(port int) bool {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return true
	}
	l.Close()
	return false
}

var ClearUserDataDir = func(userDataDirFullPath string) error {
	return os.RemoveAll(userDataDirFullPath)
}
