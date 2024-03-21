package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findLunaDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		absPath, err := filepath.Abs("")
		if err != nil {
			return "", err
		}
		dir = absPath
	}

	for {
		if filepath.Base(dir) == "luna" {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达根目录，未找到 "luna" 文件夹
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("luna directory not found")
}

func main() {
	lunaDir, err := findLunaDirectory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Luna directory:", filepath.Join(lunaDir, "devtools", "luna_lib", "img_mac_arm_01"))
}
