package input

import (
	"fmt"
	"github.com/tidwall/gjson"
	"luna/luna_utils"
	"math/rand"
	"time"
)

type TargetCoordinates struct {
	BigImgPath    string //大图的路径
	SmallImgPath  string //小图的路径
	TestImgPath   string //大图里面找到匹配小图,并且用红色的方框框起来后,生成的图片//如果不需要可以传""空字符串
	LunaThreshold string //LunaThreshold-匹配的相似度阈值,如果你0-1之间,如果满足你给定的阈值就会返回对应匹配图片的位置. 我建议先给定一个小值,然后他返回值会说明他的相似度是多少,根据返回值在给定这个值。

	/****
	以下4个边距的作用均是，如果匹配到了小图在大图中的位置,然后你想在这个框之内再次把范围缩小就可以写这个值,平时写0即可;
	但是比如是个输入框，他几乎没什么特征，很难匹配准确，这个时候就可以截小图的时候，稍微截大一点，然后在给定边距值，他会根据框定的匹配位置
	和你给定的边距，再次把范围降低，这样就便于你找到你需要点击的位置。单位都是像素;写10就是10像素的意思;
	当然这个值会自动乘屏幕缩放因子,如果你人类看到的是10像素，如果你屏幕比较高清，可能实际图片是20像素，这个值会自动算，你无需太担心。
	*/
	LeftMargin   float64 //左边距
	RightMargin  float64 //右边距
	TopMargin    float64 //上边距
	BottomMargin float64 //下边距:

	Coefficient float64 //屏幕缩放因子，一般都是1或者2
}

type TargetCoordinatesItem struct {
	SmallImgPath  string //小图的路径
	MatchScore float64
}


type ImageCoordinates struct {
	Err         error
	RandomX     float64 //这个指的是 在大图中找到的小图位置;但是他是一个矩形,如果我们点击的时候需要点击这个矩形的任意点；所以这个是在这个矩形内随机产生的横坐标；
	RandomY     float64 //同上，只是是纵坐标
	ImageWidth  int64   //原始图片的真实宽度
	ImageHeight int64   //原始图片真实高度
	MatchScore  float64 //返回的相似度,也就是他找到的小图和大图之间的相似度是多少.你可以根据这个值来调整输入参数的LunaThreshold值;
}

func position(executablePath, bigImgPath, smallImgPath, testImgPath string, leftMargin, rightMargin, topMargin, bottomMargin float64, luna_threshold string, coefficient float64) (error, float64, float64, int64, int64, float64) {
	jsonData, err := luna_utils.RunCommand(executablePath, bigImgPath, smallImgPath, testImgPath, luna_threshold)
	if err != nil {
		return err, 0, 0, 0, 0, 0
	}
	fmt.Println("匹配输出:", jsonData)
	result := gjson.Parse(jsonData)
	if result.Get("result").Exists() && result.Get("result").String() == "no match" {
		return fmt.Errorf("匹配不到图片,请观察相似度是否太低,或者coefficient系数错误"), 0, 0, 0, 0, result.Get("matchScore").Float()
	}
	if result.Get("result").Exists() && result.Get("result").String() == "Failed to read image files." {
		return fmt.Errorf("图片路径错误,无法打开图片"), 0, 0, 0, 0, 0
	}
	// 解析结果
	topLeftX := result.Get("topLeft.x").Float() / coefficient
	topRightX := result.Get("topRight.x").Float() / coefficient
	topRightY := result.Get("topRight.y").Float() / coefficient
	bottomRightY := result.Get("bottomRight.y").Float() / coefficient

	// 选择随机坐标点
	randomX := randomBetween(topLeftX+(leftMargin), topRightX-(rightMargin))
	randomY := randomBetween(topRightY+(topMargin), bottomRightY-(bottomMargin))

	fmt.Println("计算结果:")
	ImageWidth := result.Get("ImageWidth").Int()
	ImageHeight := result.Get("ImageHeight").Int()
	matchScore := result.Get("matchScore").Float()
	fmt.Println("随机坐标点:", randomX, randomY, "图片大小:", ImageWidth, ImageHeight, "匹配相似度:", matchScore)
	return nil, randomX, randomY, ImageWidth, ImageHeight, matchScore
}

// 生成指定范围内的随机浮点数
func randomBetween(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
