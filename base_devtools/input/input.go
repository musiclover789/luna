package input

import (
	"encoding/base64"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

/***
当出现匹配当小图、和相似度条件满足后退出
//程序会在你给当base目录下创建临时目录；存放大图&测试图片
我们可以这样；我们默认创建的目录叫catch目录
如果发现已经有了，我们就不创建了。如果没有就创建
我们存放的大图、就是你小图同名的小图
我们存放的测试图片也是小图同名
但是我们加一个下划线_test  _big
*/

func WaitForMatchSync(conn *protocol.DevToolsConn, smallImgPath, imgCachePath string, matchScore, coefficient float64, timeout time.Duration) (error, bool) {
	start := time.Now() // 记录开始时间
	bigImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_big")
	testImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_test")
	smallImgPath = filepath.Join(imgCachePath, smallImgPath)
	fmt.Println(bigImgPath)
	fmt.Println(testImgPath)
	fmt.Println(smallImgPath)
	for {
		elapsed := time.Since(start)
		if elapsed >= timeout {
			return fmt.Errorf("Timeout"), false
		}
		//这个地方可能只能这样设计了,就是每秒截图一次;成功就成功，失败就继续。我们在超时之内不反回
		luna_utils.DeleteFile(bigImgPath)
		luna_utils.DeleteFile(testImgPath)
		page.BringToFront(conn)
		err := page.CaptureTestRetry(conn, bigImgPath, time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		coords := TargetCoordinates{
			BigImgPath:    bigImgPath,
			SmallImgPath:  smallImgPath,
			TestImgPath:   testImgPath,
			LunaThreshold: "0.1",
			LeftMargin:    1.0,
			RightMargin:   1.0,
			TopMargin:     1.0,
			BottomMargin:  1.0,
			Coefficient:   coefficient,
		}
		imageCoordinates_step_001 := getTargetCoordinates(&coords)
		if imageCoordinates_step_001.MatchScore > matchScore {
			return nil, true
		}
	}
	return nil, false
}

/*
*
在给定时间内返回当前页面的截图&小图的匹配度\和对应的点击区域坐标
*/
func GetSmallImageCoordinates(conn *protocol.DevToolsConn, smallImgPath, imgCachePath string, coefficient float64, timeout time.Duration) (error, *ImageCoordinates) {
	return GetSmallImageCoordinatesWithMargin(conn, smallImgPath, imgCachePath, coefficient, 1.0, 1.0, 1.0, 1.0, timeout)
}

func GetSmallImageCoordinatesWithMarginMore(conn *protocol.DevToolsConn, smallImgPaths []string, imgCachePath string, coefficient, leftMargin, rightMargin, topMargin, bottomMargin float64, timeout time.Duration) (error, *ImageCoordinates) {
	start := time.Now() // 记录开始时间
	//这个地方可能只能这样设计了,就是每秒截图一次;成功就成功，失败就继续。我们在超时之内不反回

	for {
		for _, smallImgPath := range smallImgPaths {
			bigImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_big")
			testImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_test")
			smallImgPath = filepath.Join(imgCachePath, smallImgPath)
			fmt.Println(bigImgPath)
			fmt.Println(testImgPath)
			fmt.Println(smallImgPath)
			elapsed := time.Since(start)
			if elapsed >= timeout {
				return fmt.Errorf("Timeout"), nil
			}
			luna_utils.DeleteFile(bigImgPath)
			luna_utils.DeleteFile(testImgPath)
			page.BringToFront(conn)
			err := page.CaptureTestRetry(conn, bigImgPath, time.Second)
			if err != nil {
				fmt.Println(err)
				continue
			}
			coords := TargetCoordinates{
				BigImgPath:    bigImgPath,
				SmallImgPath:  smallImgPath,
				TestImgPath:   testImgPath,
				LunaThreshold: "0.1",
				LeftMargin:    leftMargin,
				RightMargin:   rightMargin,
				TopMargin:     topMargin,
				BottomMargin:  bottomMargin,
				Coefficient:   coefficient,
			}
			return nil, getTargetCoordinates(&coords)
		}
	}
}

func GetSmallImageCoordinatesWithMargin(conn *protocol.DevToolsConn, smallImgPath, imgCachePath string, coefficient, leftMargin, rightMargin, topMargin, bottomMargin float64, timeout time.Duration) (error, *ImageCoordinates) {
	start := time.Now() // 记录开始时间
	//这个地方可能只能这样设计了,就是每秒截图一次;成功就成功，失败就继续。我们在超时之内不反回
	bigImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_big")
	testImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_test")
	smallImgPath = filepath.Join(imgCachePath, smallImgPath)
	fmt.Println(bigImgPath)
	fmt.Println(testImgPath)
	fmt.Println(smallImgPath)
	for {
		elapsed := time.Since(start)
		if elapsed >= timeout {
			return fmt.Errorf("Timeout"), nil
		}
		luna_utils.DeleteFile(bigImgPath)
		luna_utils.DeleteFile(testImgPath)
		page.BringToFront(conn)
		err := page.CaptureTestRetry(conn, bigImgPath, time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		coords := TargetCoordinates{
			BigImgPath:    bigImgPath,
			SmallImgPath:  smallImgPath,
			TestImgPath:   testImgPath,
			LunaThreshold: "0.1",
			LeftMargin:    leftMargin,
			RightMargin:   rightMargin,
			TopMargin:     topMargin,
			BottomMargin:  bottomMargin,
			Coefficient:   coefficient,
		}
		return nil, getTargetCoordinates(&coords)
	}
}

func SimulateMouseClick(conn *protocol.DevToolsConn, x, y float64) {
	params := mouseEventParams{
		Type:       "mousePressed",
		Button:     MouseButtonLeft,
		ClickCount: 1,
		X:          x,
		Y:          y,
	}
	dispatchMouseEvent(conn, params)
	// Sleep for a short duration to simulate a human's reaction time.
	time.Sleep(time.Duration(luna_utils.RandomInRange(1, 50)) * time.Millisecond)
	params.Type = "mouseReleased"
	dispatchMouseEvent(conn, params)
}

func ScrollMouseToTargetImage(conn *protocol.DevToolsConn, x, y float64, totalDistance int, direction Direction, smallImgPath, imgCachePath string, coefficient, matchScore float64, timeout time.Duration) (error, bool) {
	start := time.Now() // Record start time

	// Generate file paths
	bigImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_big")
	testImgPath := luna_utils.ConcatenateFileName(smallImgPath, imgCachePath, "_test")
	smallImgPath = filepath.Join(imgCachePath, smallImgPath)

	fmt.Println(bigImgPath)
	fmt.Println(testImgPath)
	fmt.Println(smallImgPath)

	// Determine reverse flag based on direction
	var reverse bool
	var deltaX, deltaY = 1, 1
	switch direction {
	case UP:
		fmt.Println("Moving Up")
		reverse = false
		deltaX = 0
	case DOWN:
		reverse = true
		deltaX = 0
	case LEFT:
		fmt.Println("Moving Left")
		reverse = false
		deltaY = 0
	case RIGHT:
		fmt.Println("Moving Right")
		reverse = true
		deltaY = 0
	default:
		return fmt.Errorf("Unknown Direction"), false
	}
	for {
		points := getScrollPoint(totalDistance, reverse)
		len := 0
		for _, point := range points {
			if deltaX == 0 {
				deltaY = point.Distance
			} else if deltaY == 0 {
				deltaX = point.Distance
			}
			mouseWheel(conn, x, y, deltaX, deltaY)

			len += point.Distance
			if len > 100 {
				len = 0
				elapsed := time.Since(start)
				if elapsed >= timeout {
					return fmt.Errorf("Timeout"), false
				}
				luna_utils.DeleteFile(bigImgPath)
				luna_utils.DeleteFile(testImgPath)
				page.BringToFront(conn)
				err := page.CaptureTestRetry(conn, bigImgPath, time.Second)
				if err != nil {
					fmt.Println(err)
					continue
				}
				coords := TargetCoordinates{
					BigImgPath:    bigImgPath,
					SmallImgPath:  smallImgPath,
					TestImgPath:   testImgPath,
					LunaThreshold: "0.1",
					LeftMargin:    1,
					RightMargin:   1,
					TopMargin:     1,
					BottomMargin:  1,
					Coefficient:   coefficient,
				}
				r := getTargetCoordinates(&coords)
				if r.MatchScore > matchScore {
					return nil, true
					//满足条件
				}
			} else {
				time.Sleep(time.Duration(2*math.Sqrt(float64(point.Duration))) * time.Millisecond)
			}
		}

	}
	return fmt.Errorf("未找到目标图片"), false

}

// 滚动
func SimulateMouseScroll(conn *protocol.DevToolsConn, x, y float64, totalDistance int, direction Direction) {
	var points []ScrollPoint
	switch direction {
	case UP:
		fmt.Println("Moving Up")
		points = getScrollPoint(totalDistance, false)
		for _, point := range points {
			mouseWheel(conn, x, y, 0, point.Distance)
			time.Sleep(time.Duration(2*math.Sqrt(float64(point.Duration))) * time.Millisecond)
		}
	case DOWN:
		fmt.Println("Moving Down")
		points = getScrollPoint(totalDistance, true)
		for _, point := range points {
			mouseWheel(conn, x, y, 0, point.Distance)
			time.Sleep(time.Duration(2*math.Sqrt(float64(point.Duration))) * time.Millisecond)
		}
	case LEFT:
		fmt.Println("Moving Left")
		points = getScrollPoint(totalDistance, false)
		for _, point := range points {
			mouseWheel(conn, x, y, point.Distance, 0)
			time.Sleep(time.Duration(2*math.Sqrt(float64(point.Duration))) * time.Millisecond)
		}
	case RIGHT:
		fmt.Println("Moving Right")
		points = getScrollPoint(totalDistance, true)
		for _, point := range points {
			mouseWheel(conn, x, y, point.Distance, 0)
			time.Sleep(time.Duration(2*math.Sqrt(float64(point.Duration))) * time.Millisecond)
		}
	default:
		fmt.Println("Unknown Direction")
	}
}

// 移动
func SimulateMoveMouse(conn *protocol.DevToolsConn, startX, startY, endX, endY float64) {
	targetSize := luna_utils.RandomInRange(1, 100)
	// Calculate the Fitts' Law index of difficulty
	// 计算Fitts' Law的困难指数
	a := math.Abs(endX - startX)
	b := math.Abs(endY - startY)
	d := math.Sqrt(a*a + b*b)
	id := math.Log2(d/targetSize + 1)

	// Calculate the number of interpolation points
	n := int(id * 10)
	if n < 5 {
		n = 5
	}
	// 设置多阶贝塞尔曲线的控制点
	// Set up the control points of the multi-order Bezier curve
	dx := endX - startX
	dy := endY - startY
	x2 := endX
	y2 := endY
	c1x := startX + dx*0.1
	c1y := startY + dy*0.5
	c2x := startX + dx*0.3
	c2y := startY + dy*0.9

	// Generate the interpolation points using a cubic Bezier curve
	points := make([]mouseEventParams, n+1)
	for i := 0; i <= n; i++ {
		t := float64(i) / float64(n)
		if t < 0.5 {
			t = 2 * t * t
		} else {
			t = -2*t*t + 4*t - 1
		}
		x := (1-t)*(1-t)*(1-t)*startX + 3*(1-t)*(1-t)*t*c1x + 3*(1-t)*t*t*c2x + t*t*t*x2
		y := (1-t)*(1-t)*(1-t)*startY + 3*(1-t)*(1-t)*t*c1y + 3*(1-t)*t*t*c2y + t*t*t*y2
		points[i] = mouseEventParams{
			Type:   "mouseMoved",
			X:      x,
			Y:      y,
			Button: MouseButtonNone,
		}
	}

	// Send the interpolation points using input.dispatchMouseEvent
	for _, point := range points {
		dispatchMouseEvent(conn, point)
		time.Sleep(time.Duration(10*math.Sqrt(float64(n)/id)) * time.Millisecond)
	}

}

// 键盘
func SimulateKeyboardInput(conn *protocol.DevToolsConn, text string) {
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	for i, item := range pinyin.Pinyin(text, a) {
		// Sleep for a short duration to simulate a human's reaction time.
		time.Sleep(time.Duration(luna_utils.RandomInRange(1.1, 10.8)) * time.Millisecond)
		for _, char := range strings.Split(strings.Join(item, ""), "") {
			keyParamsDown := keyEventParams{
				Type: "keyDown",
				Key:  char,
			}
			dispatchKeyEvent(conn, keyParamsDown)

			keyParamsChar := keyEventParams{
				Type: "char",
				Key:  char,
			}
			dispatchKeyEvent(conn, keyParamsChar)

			keyParamsUp := keyEventParams{
				Type: "keyUp",
				Key:  char,
			}
			dispatchKeyEvent(conn, keyParamsUp)
		}
		keyParamsDown := keyEventParams{
			Type: "keyDown",
			Key:  " ",
		}
		dispatchKeyEvent(conn, keyParamsDown)

		keyParamsChar := keyEventParams{
			Type: "char",
			Key:  " ",
		}
		dispatchKeyEvent(conn, keyParamsChar)

		keyParamsUp := keyEventParams{
			Type: "keyUp",
			Key:  " ",
		}
		dispatchKeyEvent(conn, keyParamsUp)
		//
		keyParamsDown = keyEventParams{
			Type: "keyDown",
		}
		dispatchKeyEvent(conn, keyParamsDown)

		keyParamsChar = keyEventParams{
			Type: "char",
			Text: strings.Split(text, "")[i],
		}
		dispatchKeyEvent(conn, keyParamsChar)

		keyParamsUp = keyEventParams{
			Type: "keyUp",
		}
	}
}

var exeablePath string

func init() {
	lunaDir, err := findLunaDirectory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch os_item := runtime.GOOS; os_item {
	case "windows":
		fmt.Println("您的操作系统  Windows")
		exeablePath = filepath.Join(lunaDir, "devtools", "luna_lib", "win_x64", "img_win_x86_64.exe")
	case "linux":
		fmt.Println("您的操作系统  Linux")
		exeablePath = filepath.Join(lunaDir, "devtools", "luna_lib", "img_mac_arm_01")
	case "darwin":
		fmt.Println("您的操作系统  Mac OS")
		exeablePath = filepath.Join(lunaDir, "devtools", "luna_lib", "mac_arm", "img_mac_arm_01")
	default:
		fmt.Printf("Unknown OS: %v\n", os_item)
		os.Exit(1)
	}
}

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

func getTargetCoordinates(params *TargetCoordinates) *ImageCoordinates {
	executablePath := exeablePath
	//fmt.Println("定位坐标的程序位置:", executablePath)
	err, randomX, randomY, imageWidth, imageHeight, matchScore := position(executablePath, params.BigImgPath, params.SmallImgPath, params.TestImgPath, params.LeftMargin, params.RightMargin, params.TopMargin, params.BottomMargin, params.LunaThreshold, params.Coefficient)
	return &ImageCoordinates{
		Err:         err,
		RandomX:     randomX,
		RandomY:     randomY,
		ImageWidth:  imageWidth,
		ImageHeight: imageHeight,
		MatchScore:  matchScore,
	}
}

func SetInterceptDrags(conn *protocol.DevToolsConn) {
	id := luna_utils.IdGen.NextID()
	// 发送拖拽事件
	params := map[string]interface{}{
		"enabled": true,
	}
	req := map[string]interface{}{
		"id":     id,
		"method": "Input.setInterceptDrags",
		"params": params,
	}
	conn.WriteMessage(req)
}

func DispatchDragEvent(conn *protocol.DevToolsConn, x, y int, event string) {
	id := luna_utils.IdGen.NextID()
	// 创建拖拽事件参数
	dragData := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{
				"mimeType": "image/jpeg",                                                                              // 指定图片类型
				"data":     base64.StdEncoding.EncodeToString([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}), // 你的图片数据
			},
		},
		"files":              []string{"/Users/hongyuji/Pictures/hxx/IMG_2545.JPG"}, // 指定文件路径
		"dragOperationsMask": 1,                                                     // 设置允许的拖拽操作
	}

	// 发送拖拽事件
	params := map[string]interface{}{
		"type": event,
		"x":    x,
		"y":    y,
		"data": dragData,
	}
	req := map[string]interface{}{
		"id":     id,
		"method": "Input.dispatchDragEvent",
		"params": params,
	}
	conn.WriteMessage(req)
}
