package luna_utils

import (
	"fmt"
	"luna/log"
	"luna/reverse_proxy"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//--window-size=800,600

var StartChromiumWithUserDataDir = func(chromiumPath, userDataDirFullPath string, proxy *string, isHeadless bool, size func() (bool, int, int), customArgs ...string) int {
	// 获取随机端口
	port, err := getRandomPort()
	if err != nil {
		luna_log.LogError("failed to get random port: %v\n", err)
		return -1
	}

	// 检查端口是否被占用
	if isPortOpen(port) {
		luna_log.LogError("port %d is already in use\n", port)
		return -1
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
	if proxy != nil && len(*proxy) > 0 {
		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			luna_log.LogError("Failed to parse proxy URL: %v", err)
			return -1
		}
		pwd, _ := proxyURL.User.Password()

		proxy_port, err := reverse_proxy.StartProxyServer(proxyURL.Scheme, proxyURL.Hostname(), proxyURL.Port(), proxyURL.User.Username(), pwd)
		if err != nil {
			luna_log.LogError("Failed to start proxy server: %v", err)
			return -1
		}
		chromiumCmdArgs = append(chromiumCmdArgs, "--proxy-server=127.0.0.1:"+proxy_port)
	}
	fmt.Println("启动参数:", chromiumCmdArgs)
	if len(userDataDirFullPath) > 0 {
		chromiumCmdArgs = append(chromiumCmdArgs, "--user-data-dir="+userDataDirFullPath)
	}
	fmt.Println(chromiumCmdArgs)
	switch os_item := runtime.GOOS; os_item {
	case "windows":
		writeToLogFile("C:\\luna-temp", chromiumCmdArgs)
	}

	chromiumCmd := exec.Command(chromiumPath, chromiumCmdArgs...)

	err = chromiumCmd.Start()
	if err != nil {
		luna_log.LogError("Failed to start process: %v\n", err)
		return -1
	}

	if chromiumCmd.ProcessState != nil && chromiumCmd.ProcessState.Exited() {
		luna_log.LogError("Failed to start process, exit code %d\n", chromiumCmd.ProcessState.ExitCode())
		return -1
	}

	return port
}

func writeToLogFile(filePath string, customArgs []string) error {

	// 获取当前毫秒数
	milliseconds := time.Now().UnixNano() / int64(time.Millisecond)
	// 将毫秒数转换为字符串
	fileName := fmt.Sprintf("%010d", milliseconds)
	// 创建文件并打开
	fullPath := filepath.Join(filePath, fileName)
	// 创建文件
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 遍历参数数组，写入每个参数
	for _, arg := range customArgs {
		// 去掉参数开头的"--"并进行trim
		line := strings.TrimPrefix(arg, "--")
		line = strings.TrimSpace(line)

		// 将处理后的参数写入文件
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
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
