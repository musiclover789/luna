package main

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 指定字体文件夹路径
	fontsDir := "/Users/hongyuji/Documents/workspace/testFont"

	// 读取字体文件夹中的所有文件
	files, err := ioutil.ReadDir(fontsDir)
	if err != nil {
		fmt.Println("Error reading fonts directory:", err)
		return
	}
	fonts := strings.Builder{}

	// 遍历每个文件并解析字体文件
	for _, file := range files {
		// 构建完整路径
		fontPath := filepath.Join(fontsDir, file.Name())

		// 打开字体文件
		fontFile, err := os.Open(fontPath)
		if err != nil {
			fmt.Println("Error opening font file:", err)
			continue
		}
		defer fontFile.Close()

		// 读取字体文件内容
		fontData, err := ioutil.ReadAll(fontFile)
		if err != nil {
			fmt.Println("Error reading font file:", err)
			continue
		}

		// 解析字体文件
		font, err := truetype.Parse(fontData)
		if err != nil {
			fmt.Println("Error parsing font file:", err)
			continue
		}

		// 获取字体名称和字体族
		fontName := font.Name(truetype.NameIDFontFullName)
		fontFamily := font.Name(truetype.NameIDFontFamily)

		// 打印结果
		fonts.WriteString(fontName)
		fonts.WriteString(",")
		fmt.Println("Font Name:>>>", fontName, "<<<")
		fmt.Println("Font Family:", fontFamily)
		fmt.Println("Font File:", fontPath)
		fmt.Println("----------------------------------")

	}
	fmt.Println(fonts.String())
}
