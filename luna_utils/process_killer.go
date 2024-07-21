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

// CheckProcessRunning 检查指定PID的进程是否还在运行。
// 如果进程还在运行，返回false；如果进程已终止或不存在，返回true。
func CheckProcessRunning(pid int) bool {
	var checkCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		checkCmd = exec.Command("tasklist", "/FI", "PID eq "+strconv.Itoa(pid))
	} else {
		checkCmd = exec.Command("ps", "-p", strconv.Itoa(pid))
	}

	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return false // 命令执行失败，假设进程不存在
	}

	if runtime.GOOS == "windows" {
		// Windows下检查tasklist命令的输出
		outputStr := string(output)
		if strings.Contains(outputStr, "No tasks are running") {
			return false // 未找到该PID对应的进程
		}
	} else {
		// Unix/Linux下检查ps命令的输出
		outputStr := string(output)

		lines := strings.Split(outputStr, "\n")
		// 第二行是ps命令的输出，如果长度为1，表示没有找到进程
		if len(lines) == 1 {
			return false // 没有找到该PID对应的进程
		}
	}

	return true // 进程存在
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
