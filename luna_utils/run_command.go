package luna_utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(executablePath string, args ...string) (string, error) {
	cmd := exec.Command(executablePath, args...)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行命令出错: %v", err)
	}
	result := string(out)
	return result, nil
}



