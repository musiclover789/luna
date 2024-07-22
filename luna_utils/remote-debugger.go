package luna_utils

import (
	"fmt"
	"github.com/musiclover789/luna/log"
	"github.com/musiclover789/luna/reverse_proxy"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//--window-size=800,600

var StartChromiumWithUserDataDir = func(chromiumPath, userDataDirFullPath string, proxy *string, isHeadless bool, size func() (bool, int, int), customArgs ...string) (int, *reverse_proxy.ProxyServer, int) {
	// 获取随机端口
	port, err := getRandomPort()
	if err != nil {
		luna_log.LogError("failed to get random port: %v\n", err)
		return -1, nil, -1
	}

	// 检查端口是否被占用
	if isPortOpen(port) {
		luna_log.LogError("port %d is already in use\n", port)
		return -1, nil, -1
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
	//这里是全部的指纹信息
	fingerprintArgs := []string{}
	for _, arg := range customArgs {
		if len(arg) > 0 {
			fingerprintArgs = append(fingerprintArgs, arg)
		}
	}

	/***
	根据指纹信息、
	1、判断chromiumCmdArgs是否有值；
	2、如果有在次判断是否有指定的盘符;
	3、拼接数据到字符串,然后写文件到指定到目录。就可以了
	4、如果超过256个,就删除前面的
	*/

	switch os_item := runtime.GOOS; os_item {
	case "windows":
		fmt.Println("您的操作系统  Windows-设置指纹信息")
		filePath := "C:\\luna-temp"
		err = writeFile(filePath, fingerprintArgs, strconv.Itoa(port))
		if err != nil {
			fmt.Println("写入指纹信息-错误:", err)
		}
		path := filePath
		//指纹文件超过256个自动删除
		num := 256
		err = deleteExcessFiles(path, num)
		if err != nil {
			fmt.Println("指纹信息路径错误:", err)
		}
	}
	//----

	chromiumCmdArgs = append(chromiumCmdArgs, "--remote-debugging-port="+strconv.Itoa(port))
	var proxyServer *reverse_proxy.ProxyServer
	if proxy != nil && len(*proxy) > 0 {
		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			luna_log.LogError("Failed to parse proxy URL: %v", err)
			return -1, nil, -1
		}
		pwd, _ := proxyURL.User.Password()

		// 创建一个代理服务器实例
		fmt.Println(proxyURL.Scheme, proxyURL.Hostname(), proxyURL.Port(), proxyURL.User.Username(), pwd)
		proxyServer = reverse_proxy.NewProxyServer(proxyURL.Scheme, proxyURL.Hostname(), proxyURL.Port(), proxyURL.User.Username(), pwd)

		// 启动代理服务器
		proxy_port, err := proxyServer.Start()
		if err != nil {
			luna_log.LogError("Failed to start proxy server: %v", err)
		}
		log.Printf("Proxy server started on port %s", port)

		if err != nil {
			luna_log.LogError("Failed to start proxy server: %v", err)
			return -1, nil, -1
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
		return -1, nil, -1
	}
	if chromiumCmd.ProcessState != nil && chromiumCmd.ProcessState.Exited() {
		luna_log.LogError("Failed to start process, exit code %d\n", chromiumCmd.ProcessState.ExitCode())
		return -1, nil, -1
	}
	pid := chromiumCmd.Process.Pid
	fmt.Println("Browser process PID:", pid)
	return port, proxyServer, pid
}

var mutex sync.Mutex
var CreateCacheDirInSubDir = func(basePath string) string {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	rand.Seed(time.Now().UnixNano())

	// 生成随机字母
	letters := make([]rune, 3)
	for i := 0; i < 3; i++ {
		letters[i] = rune('a' + rand.Intn(26))
	}

	// 获取当前时间戳的中间 9 到 16 位数字
	timestamp := time.Now().UnixNano()
	middleDigits := (timestamp / 1e6) % 1e8

	randFolderName := fmt.Sprintf("user_%08d%s", middleDigits, string(letters))

	cacheDirFullPath := filepath.Join(basePath, randFolderName)

	// 检查文件夹是否已存在
	if _, err := os.Stat(cacheDirFullPath); err == nil {
		return cacheDirFullPath
	}

	if err := os.MkdirAll(cacheDirFullPath, 0777); err != nil {
		fmt.Printf("创建缓存目录失败: %v\n", err)
		return ""
	}
	time.Sleep(time.Millisecond * 10)
	fmt.Println("当前缓存目录为:", cacheDirFullPath)
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

/****

 */

// 函数1：根据给定的文件路径和字符串数组，将字符串拼接后写入文件
func writeFile(filePath string, args []string, name string) error {
	// 检查参数是否为空
	if len(args) == 0 {
		return fmt.Errorf("args不能为空")
	}
	// 在路径末尾添加斜杠
	if !strings.HasSuffix(filePath, string(filepath.Separator)) {
		filePath += string(filepath.Separator)
	}

	// 检查文件目录是否存在
	//dir := filepath.Dir(filePath)
	//if _, err := os.Stat(dir); os.IsNotExist(err) {
	//	return fmt.Errorf("目录 %s 不存在", dir)
	//}
	// 检查目录是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(filePath, 0755) // 0755 表示权限，具体权限可以根据需要调整
		if err != nil {
			fmt.Println("创建目录失败:", err)
		} else {
			fmt.Println("目录创建成功:", filePath)
		}
	}

	// 拼接参数字符串
	var content string
	for _, arg := range args {
		arg = strings.TrimLeft(arg, "-")
		content += arg + "\n"
	}

	// 获取当前时间的秒数作为文件名
	//fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	//fileName = fileName[2 : len(fileName)-3]
	filePath = filepath.Join(filePath, name)

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("删除文件失败:", err)
	}

	// 写入文件
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败：%v", err)
	}

	return nil
}

// 函数2：根据给定路径判断文件数量是否超过指定数量，并删除多余的文件
func deleteExcessFiles(path string, num int) error {

	// 在路径末尾添加斜杠
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}

	// 检查文件目录是否存在
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("目录 %s 不存在", dir)
	}

	// 打开目录
	dirEntries, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("打开目录失败：%v", err)
	}

	// 排除特定文件
	var files []os.FileInfo
	for _, entry := range dirEntries {
		if entry.Name() != "uname.txt" && entry.Name() != "license.txt" {
			files = append(files, entry)
		}
	}

	// 排序文件列表
	sortFilesByNumber(files)

	// 删除多余文件
	for i := num; i < len(files); i++ {
		err := os.Remove(filepath.Join(path, files[i].Name()))
		if err != nil {
			return fmt.Errorf("删除文件失败：%v", err)
		}
	}

	return nil
}

// 根据文件名中的数字进行反向排序
func sortFilesByNumber(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		numI, _ := strconv.Atoi(strings.TrimSuffix(files[i].Name(), filepath.Ext(files[i].Name())))
		numJ, _ := strconv.Atoi(strings.TrimSuffix(files[j].Name(), filepath.Ext(files[j].Name())))
		return numI > numJ // 反向排序
	})
}
