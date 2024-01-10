package luna_utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateCacheDirectory(path string) (error, string) {
	cachePath := filepath.Join(path, "cache")

	// 检查目录是否存在
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.Mkdir(cachePath, 0755)
		if err != nil {
			return fmt.Errorf("无法创建目录：%s", err), cachePath
		}
		fmt.Printf("已成功创建缓存目录：%s\n", cachePath)
	} else {
		fmt.Printf("缓存目录已存在：%s\n", cachePath)
	}

	return nil, cachePath
}

func ConcatenateFileName(filePath, dirPath, extName string) string {
	// 获取文件名
	fileName := filepath.Base(filePath)

	// 拼接目标文件路径
	targetFilePath := filepath.Join(dirPath, fileName)

	// 获取文件扩展名
	fileExt := filepath.Ext(targetFilePath)

	// 去除文件扩展名后的文件名
	fileNameWithoutExt := strings.TrimSuffix(targetFilePath, fileExt)

	// 拼接文件名和指定的后缀以及扩展名
	resultFilePath := fileNameWithoutExt + extName + fileExt

	return resultFilePath
}


func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}

	return nil
}

