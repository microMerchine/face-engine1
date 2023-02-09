package main

import (
	"fmt"

	. "github.com/loouss/face-engine1/v4"
	"github.com/loouss/face-engine1/v4/util"
)

var imageInfo = util.GetResizedImageInfo("./mask.jpg")

func main() {
	// 激活SDK
	if err := OnlineActivation("2eCYi7C1SiTMWCDWXZzt27CVbMTQcUov9452yhphUisF", "9gc3DQBq93eiSxGUvXFgTLmvv8Xp1ZAKua7RVKZ8i1jo", "82G1-11GA-B13Z-B2Z5"); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeImage,
		OrientPriority0,
		10, // 4.0最大支持10个人脸
		EnableFaceDetect|EnableFaceRecognition|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender|EnableMaskDetect)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	deviceInfo, err := GetActiveDeviceInfo()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("设备信息：%s\n", deviceInfo)
	// 检测人脸
	info, err := engine.DetectFaces(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 处理人脸数据
	if err = engine.Process(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, info, EnableAge|EnableGender|EnableLiveness|EnableMaskDetect); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)
	// 获取口罩信息
	maskInfo, err := engine.GetMask()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("口罩信息：%#v\n", maskInfo)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
}
