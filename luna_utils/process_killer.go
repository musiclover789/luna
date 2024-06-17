package luna_utils

import (
	"fmt"
	luna_log "github.com/musiclover789/luna/log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var KillProcess = func() {
	switch os_item := runtime.GOOS; os_item {
	case "windows":
		fmt.Println("您的操作系统  Windows")
		killProcess_Windows()
	case "linux":
		fmt.Println("您的操作系统  Linux")
		killProcess_Linux()
	case "darwin":
		fmt.Println("您的操作系统  Mac OS")
		killProcess_Mac()
	default:
		fmt.Printf("Unknown OS: %v\n", os_item)
		os.Exit(1)
	}
}

func KillProcessByPid(pid int) {
	if runtime.GOOS == "windows" {
		exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid)).Run()
	} else {
		exec.Command("kill", strconv.Itoa(pid)).Run()
	}
}

var killProcess_Mac = func() {
	// 查找特定的进程名称
	psCommand := "ps aux | grep Chromium | grep -v grep | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", psCommand).Output()
	if err != nil {
		luna_log.LogError("Error running command: %s\n", err)
		os.Exit(1)
	}
	// 解析出进程ID
	pids := strings.Split(string(out), "\n")
	for _, pidStr := range pids {
		if pidStr != "" {
			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				luna_log.Logf("Invalid PID: %s\n", pidStr)
			} else {
				// 结束进程
				killCommand := fmt.Sprintf("kill %d", pid)
				_, err := exec.Command("bash", "-c", killCommand).Output()
				if err != nil {
					luna_log.Logf("Error running command: %s\n", err)
				} else {
					luna_log.Logf("Process %d killed\n", pid)
				}
			}
		}
	}
}

var killProcess_Linux = func() {
	// 查找特定的进程名称
	cmd := "ps aux | grep chromium | grep -v grep | awk '{print $2}' | xargs kill -9"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		luna_log.LogFatal("Error running command: %s\n", err)
		os.Exit(1)
	}
	luna_log.Logf("Process killed: %s", out)
}

func killProcess_Windows() {
	// 执行任务列表命令获取进程信息
	out, err := exec.Command("tasklist", "/NH").Output()
	if err != nil {
		fmt.Printf("Error running command: %s\n", err)
		os.Exit(1)
	}
	// 解析出进程ID
	processes := strings.Split(string(out), "\n")
	for _, process := range processes {
		if strings.Contains(strings.ToLower(process), "chromium") || strings.Contains(strings.ToLower(process), "chrome") {
			fields := strings.Fields(process)
			if len(fields) >= 2 {
				pidStr := fields[1]
				pid, err := strconv.Atoi(pidStr)
				if err != nil {
					fmt.Printf("Invalid PID: %s\n", pidStr)
				} else {
					// 结束进程
					err := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid)).Run()
					if err != nil {
						fmt.Printf("Error killing process %d: %s\n", pid, err)
					} else {
						fmt.Printf("Process %d killed\n", pid)
					}
				}
			}
		}
	}
}
