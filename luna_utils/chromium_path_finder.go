package luna_utils

/****
这里我希望
1、可以自己找到需要执行的chromium 可执行文件
2、可以自己执行命令行、启动chromium可执行文件、并且保证端口不可以冲突
3、可以自己获取devtools的 webSocketDebuggerUrl 地址;
*/
import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/***
逻辑:
1、找寻自己路径的父文件夹、查找Chromium的可执行程序、如果找不到就继续往上层目录找一直找到根目录为止;
2、如果mac系统就找寻Chromium.app
3、如果windows系统就找寻Chromium.exe
4、todo：如果是linux系统暂时不考虑,搞不过来。
return 返回Chromium这个可执行程序的具体路径.
*/
var FindChromiumPath = func() string {
	var appName string
	switch os_item := runtime.GOOS; os_item {
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Mac OS")
		appName = "Chromium.app"
	default:
		fmt.Printf("Unknown OS: %v\n", os_item)
		os.Exit(1)
	}
	chromiumPath := findChromiumPath(appName)
	if strings.EqualFold("", chromiumPath) {
		fmt.Println("没有找到", appName, "路径")
		os.Exit(1)
		return ""
	} else {
		fmt.Println("Chromium.app 路径:", chromiumPath)
		return chromiumPath
	}
}

func findChromiumPath(appName string) string {
	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Traverse up the directory tree until the root directory
	for dir != "/" {
		// Check if the current directory contains the desired application
		if containsApp(dir, appName) {
			return filepath.Join(dir, appName)
		}

		// Move up one directory
		dir = filepath.Dir(dir)
	}

	return ""
}

// Helper function to check if the given directory contains the desired application
func containsApp(dir string, appName string) bool {
	// Check if the directory contains a file or directory with the desired name
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	for _, file := range files {
		if file.Name() == appName {
			return true
		}
	}

	return false
}
