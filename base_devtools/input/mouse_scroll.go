package input

import (
	"math"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UNKNOWN
)

type ScrollPoint struct {
	Distance int
	Duration int
}

func getScrollPoint(totalDistance int, isNegative bool) []ScrollPoint {
	//鼠标滚动样本数据
	dataSamples := []ScrollPoint{
		{2, 0},
		{8, 7},
		{8, 0},
		{18, 8},
		{34, 8},
		{44, 9},
		{78, 8},
		{86, 8},
		{98, 8},
		{116, 8},
		{124, 8},
		{124, 8},
		{150, 8},
		{172, 16},
		{198, 8},
		{190, 8},
		{182, 9},
		{174, 8},
		{168, 8},
		{164, 8},
		{148, 8},
		{146, 8},
		{142, 9},
		{140, 8},
		{140, 8},
		{134, 9},
		{132, 8},
		{130, 8},
		{128, 8},
		{124, 8},
		{124, 8},
		{120, 8},
		{118, 8},
		{116, 8},
		{114, 9},
		{110, 8},
		{112, 8},
		{106, 8},
		{104, 9},
		{104, 8},
		{98, 8},
		{98, 8},
		{94, 9},
		{92, 7},
		{90, 8},
		{90, 8},
		{86, 8},
		{82, 8},
		{82, 8},
		{76, 8},
		{76, 8},
		{76, 8},
		{72, 12},
		{68, 5},
		{68, 8},
		{64, 9},
		{62, 8},
	}
	return interpolateScrollData(dataSamples, totalDistance, isNegative)
}
func interpolateScrollData(dataSamples []ScrollPoint, totalDistance int, isNegative bool) []ScrollPoint {
	interpolatedData := []ScrollPoint{}
	currentDistance := 0

	for i := 0; i < len(dataSamples)-1; i++ {
		currSample := dataSamples[i]
		nextSample := dataSamples[i+1]

		// 计算两个样本之间的时间间隔差
		durationDiff := nextSample.Duration - currSample.Duration

		// 计算当前段的距离和时间间隔
		distance := nextSample.Distance - currSample.Distance

		// 计算当前段的距离比例
		distanceRatio := float64(distance) / float64(totalDistance)
		// 调整距离比例，增加起伏
		distanceRatio *= float64(transformNumber(totalDistance))
		// 插值生成滚动数据
		for t := 0.0; t <= 1.0; t += 0.1 {
			// 根据贝塞尔曲线插值算法计算距离和时间间隔
			interpolatedDistance := int(math.Round(float64(currSample.Distance) + distanceRatio*bezierInterpolation(t)*float64(totalDistance)/4))
			interpolatedDuration := int(math.Round(float64(currSample.Duration) + float64(durationDiff)*bezierInterpolation(t)))

			// 检查生成的距离是否超过了目标距离的剩余量
			if currentDistance+interpolatedDistance > totalDistance {
				interpolatedDistance = totalDistance - currentDistance
			}

			// 将当前点的距离和时间间隔添加到结果中
			interpolatedDistanceV := interpolatedDistance
			if !isNegative {
				interpolatedDistanceV = convertToNegative(interpolatedDistanceV)
			}
			interpolatedData = append(interpolatedData, ScrollPoint{Distance: interpolatedDistanceV, Duration: interpolatedDuration})
			currentDistance += interpolatedDistance

			// 如果已达到目标距离，则退出循环
			if currentDistance >= totalDistance {
				break
			}
		}

		// 如果已达到目标距离，则退出循环
		if currentDistance >= totalDistance {
			break
		}
	}

	return interpolatedData
}
func convertToNegative(num int) int {
	return num * -1
}

func transformNumber(num int) int {
	k := 0.5 // 调整斜率参数 k 的值，根据需要进行适当调整
	// 应用 S 形曲线函数
	result := int(160 / (1 + math.Exp(-k*float64(num))))
	return result
}

func bezierInterpolation(t float64) float64 {
	// 调整贝塞尔曲线的控制点，增加起伏程度
	controlPoints := []float64{0.0, 0.25, 0.25, 1.0}
	return math.Pow(1-t, 3)*controlPoints[0] + 3*math.Pow(1-t, 2)*t*controlPoints[1] + 3*(1-t)*math.Pow(t, 2)*controlPoints[2] + math.Pow(t, 3)*controlPoints[3]
}
