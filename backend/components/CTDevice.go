/*
 * @Author: tang
 * @Date: 2026-06-12
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 探测器设备管理模块，负责探测器连接和事件处理
 */
package components

import (
	sdk "CTSystem/backend/CSDK"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 探测器工作模式
type CTMODE int

const (
	IDLE    = 0
	HST     = 2
	DST     = 9
	AED1    = 3
	AED2    = 4
	RECOVER = 9
	ERR     = 99
)

type CTStatus struct {
	Mode          CTMODE // 当前工作模式
	ExposeTime    int    // 曝光时间
	Binning       string // 当前binning
	TriggerRepeat int    // 手动触发图片张数
	TriggerDelay  int    // 手动触发图片间隔
}

type CTDevice struct {
	CT    *sdk.NetCom
	Stage *MotorDevice // 电机设备

	IsConnecting bool
	comFpList    sdk.TComFpList
	ctx          context.Context

	status        CTStatus
	CTName        string
	CTPath        string
	CTRunning     bool
	CTImageCount  int // CT扫描图像计数
	CTImageMinVal int //用于图片归一化，最小值
	CTImageMaxVal int //用于图片归一化，最大值

}

// NewCTDevice 创建一个新探测器设备
func NewCTDevice() *CTDevice {
	return &CTDevice{
		CTImageMinVal: 0,
		CTImageMaxVal: 255,
		CTName:        "ct",
		CTPath:        "D:/appfile/code/vs/CTSystem/test",
	}
}

// SetContent 设置上下文
func (this *CTDevice) SetContent(ctx context.Context) {
	this.ctx = ctx
	Cnet := sdk.NewNetCom(ctx)
	this.CT = Cnet
}

// SetStage 设置电机设备
func (this *CTDevice) SetStage(stage *MotorDevice) {
	this.Stage = stage
}

// SetCTFileNameAndPath 设置探测器文件名和路径
func (this *CTDevice) SetCTFileNameAndPath(path string, name string) {
	fmt.Println("设置探测器文件名和路径", path, name)
	this.CTName = name
	this.CTPath = path
}

// InitCT 初始化探测器设备，自动初始化SDK, 并注册事件回调函数
func (this *CTDevice) InitCT() {
	fmt.Println("初始化探测器设备")
	this.CT.RegisterCallback()     // 先注册回调
	this.CT.COM_Init()             // 再初始化SDK
	this.CT.COM_StartNet()         // 启动网络监听
	this.CT.COM_SetCalibMode(0x06) // 设置校准模式 IMG_CALIB_GAIN | IMG_CALIB_DEFECT

	//监听ct_raw事件
	runtime.EventsOn(this.ctx, "ct_raw", func(raw16Data ...interface{}) {
		fmt.Println("接收到CT原始数据")
		if len(raw16Data) > 0 {
			this.handleCTImageEvent(raw16Data[0])
		}
	})
}

// DetectorConnect 连接探测器
func (this *CTDevice) DetectorConnect() error {
	fmt.Println("连接探测器")
	flag := this.CT.COM_Open(0)
	if flag {
		this.IsConnecting = true
		return nil
	} else {
		fmt.Println("无法连接探测器: ")
		return fmt.Errorf("无法连接探测器")
	}
}

// DetectorDisconnect 断开连接探测器
func (this *CTDevice) DetectorDisconnect() error {
	fmt.Println("断开连接器")
	flag := this.CT.COM_Close()
	if flag {
		this.IsConnecting = false
		return nil
	} else {
		fmt.Println("无法断开连接")
		return fmt.Errorf("无法断开连接")
	}
}

// DetectorIsConnected 检查连接状态
func (this *CTDevice) DetectorIsConnected() bool {
	fmt.Println("检查连接状态")
	return this.IsConnecting
}

// ContinuousTrigger 连续触发
func (this *CTDevice) ContinuousTrigger() error {
	fmt.Println("连续触发")

	return nil
}

// DetectorSingleScan 单次扫描
func (this *CTDevice) DetectorSingleScan() error {
	// fmt.Println("单次扫描", this.ctx == nil)
	if this.status.Mode != DST {
		this.CT.COM_Dst()
		time.Sleep(5000 * time.Millisecond)
		this.CT.COM_Dprep()
		time.Sleep(5000 * time.Millisecond) //等待拍背底完成
		this.status.Mode = DST
		fmt.Println("[DST] 启动DST")
	}
	this.CT.COM_Dacq()
	// fmt.Println("[HST] 获取图片")
	return nil
}

// DetectorSetExposeTime 设置曝光时间
func (this *CTDevice) DetectorSetExposeTime(exposeTime int) error {
	fmt.Println("设置曝光时间", exposeTime)
	if !this.CT.COM_SetExposeTime(exposeTime) {
		fmt.Println("无法设置曝光时间")
		return fmt.Errorf("无法设置曝光时间")
	}
	this.status.ExposeTime = exposeTime
	fmt.Println("设置曝光时间成功")
	return nil
}

// DetectorGetExposeTime 获取曝光时间
func (this *CTDevice) DetectorGetExposeTime() int {
	fmt.Println("获取曝光时间")
	this.status.ExposeTime = this.CT.COM_GetExposeTime()
	return this.status.ExposeTime
}

// DetectorSetBinning 设置像素数
func (this *CTDevice) DetectorSetBinning(binning string) error {
	fmt.Println("设置像素数", binning)
	if !this.CT.COM_SetBinning(binning) {
		fmt.Println("无法设置像素数", binning)
		return fmt.Errorf("无法设置像素数")
	}
	this.status.Binning = binning
	fmt.Println("设置像素数成功")
	return nil
}

// DetectorGetBinning 获取像素数
func (this *CTDevice) DetectorGetBinning() string {
	fmt.Println("获取像素数")
	this.status.Binning = this.CT.COM_GetBinning()
	return this.status.Binning
}

// DetectorSetRepeat 设置重复次数
func (this *CTDevice) DetectorSetRepeat(repeat int) error {
	return nil
}

// DetectorGetRepeat 获取重复次数
func (this *CTDevice) DetectorGetRepeat() int {
	return 0
}

// DetectorSetGain 设置增益
func (this *CTDevice) DetectorSetGain(gain int) error {
	return nil
}

// DetectorGetGain 获取增益
func (this *CTDevice) DetectorGetGain() int {
	return 0
}

// ****************************************   CT联测   ***************************************
// StartCTScan 开启CT连续扫描
func (this *CTDevice) StartCTScan() error {
	fmt.Printf("[CT扫描] 开始扫描到路径: %s\n", this.CTPath)
	// 确保输出目录存在
	if err := os.MkdirAll(this.CTPath, 0755); err != nil {
		fmt.Printf("[CT扫描] 创建输出目录失败: %v\n", err)
		return fmt.Errorf("创建输出目录失败: %v", err)
	}

	if this.CTRunning {
		return fmt.Errorf("CT扫描已经在运行中")
	}

	this.CTRunning = true
	this.CTImageCount = 0

	this.runSingleCTScan()

	return nil
}

// runSingleCTScan 运行单次CT扫描循环
func (this *CTDevice) runSingleCTScan() {

	// 检查是否暂停或完成
	if !this.CTRunning || this.CTImageCount >= 360 {
		fmt.Printf("[CT扫描] 暂停扫描或已扫描完成, 总扫描次数: %d\n", this.CTImageCount)
		this.CTRunning = false
		return // 暂停扫描或已扫描完成
	}

	err := this.Stage.StageMoveRel("R", 1.0)
	if err != nil {
		fmt.Printf("[CT扫描] 运动R轴失败: %v\n", err)
		// return
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("[CT扫描] 运动R轴完成, 当前R位置: %.2f°\n", this.Stage.GetLengths()["R"])

	// 执行单次扫描
	err = this.DetectorSingleScan()
	if err != nil {
		fmt.Printf("[CT扫描] 扫描失败: %v\n", err)
		// return
	}
	//等待handleCTImageEvent触发

	// //测试代码
	// time.Sleep(1 * time.Second)
	// // 增加图像计数
	// this.CTImageCount++
	// // 通知电机再次运动
	// this.runSingleCTScan()
}

// handleCTImageEvent 处理ct_raw事件数据
func (this *CTDevice) handleCTImageEvent(data interface{}) {
	// 解析事件数据
	rawData, ok := data.(map[string]interface{})
	if !ok {
		fmt.Printf("[CTDevice] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
		return
	}

	// 提取图片数据
	imageData, ok := rawData["image"].([]byte)
	if !ok {
		fmt.Printf("[CTDevice] 图片数据类型转换失败\n")
		return
	}

	// 提取宽高
	width, _ := rawData["width"].(int)
	height, _ := rawData["height"].(int)
	if width <= 0 || height <= 0 {
		fmt.Printf("[CTDevice] 图片尺寸无效: width=%d, height=%d\n", width, height)
		return
	}

	// 创建标准的8位灰度图像容器
	grayImg := image.NewGray(image.Rect(0, 0, width, height))

	// 16位RAW转8位灰度：使用CTImageMinVal和CTImageMaxVal进行归一化
	minVal := uint32(this.CTImageMinVal)
	maxVal := uint32(this.CTImageMaxVal)
	rangeVal := maxVal - minVal
	if rangeVal == 0 {
		rangeVal = 1 // 防止除零
	}

	for i := 0; i < width*height; i++ {
		// Little-endian
		val := uint32(imageData[i*2]) | (uint32(imageData[i*2+1]) << 8)
		// 归一化到0-255
		var normalized uint32
		if val <= minVal {
			normalized = 0
		} else if val >= maxVal {
			normalized = 255
		} else {
			normalized = (val - minVal) * 255 / rangeVal
		}
		grayImg.Pix[i] = uint8(normalized)
	}

	// JPEG编码
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, grayImg, &jpeg.Options{Quality: 60}); err != nil {
		fmt.Printf("[CTDevice] JPEG编码失败: %v\n", err)
		return
	}

	// 转换为Base64
	encodedStr := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	// 发送ct_image事件供前端显示
	runtime.EventsEmit(this.ctx, "ct_image", map[string]string{
		"image":  encodedStr,
		"width":  fmt.Sprintf("%d", width),
		"height": fmt.Sprintf("%d", height),
	})

	// 只有当CTRunning为true时才保存图片
	if this.CTRunning {
		// 增加图像计数
		this.CTImageCount++
		// 通知电机再次运动
		this.runSingleCTScan()
		// 保存原始图片到本地
		this.SaveRawToLocal(rawData)
	}
}

// SaveRawToLocal 保存原始图片到本地
func (this *CTDevice) SaveRawToLocal(data interface{}) error {
	fmt.Println("保存原始图片到本地")
	rawData, ok := data.(map[string]interface{})
	if !ok {
		fmt.Printf("[CTDevice] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
		return fmt.Errorf("数据类型转换失败")
	}
	filePath := this.CTPath + "/" + this.CTName + fmt.Sprintf("%d", this.CTImageCount) + ".raw"
	if _, err := os.Stat(this.CTPath); os.IsNotExist(err) {
		os.MkdirAll(this.CTPath, 0755)
	}
	// 将原始16位图片数据写入文件
	err := os.WriteFile(filePath, rawData["image"].([]byte), 0644)
	if err != nil {
		fmt.Printf("[CTDevice] 写入文件失败: %v\n", err)
	} else {
		fmt.Printf("[CTDevice] 图片已写入 %s，大小: %d 字节\n", filePath, len(rawData["image"].([]byte)))
	}
	return nil
}

// CTPauseScan 暂停CT扫描
func (this *CTDevice) CTPauseScan() error {
	this.CTRunning = false
	fmt.Printf("[CT扫描] 暂停扫描, 总扫描次数: %d\n", this.CTImageCount)
	return nil
}

// ContinueCTScan 继续CT扫描
func (this *CTDevice) ContinueCTScan() error {
	this.CTRunning = true
	fmt.Printf("[CT扫描] 继续扫描, 总扫描次数: %d\n", this.CTImageCount)
	this.runSingleCTScan()
	return nil
}

// StopCTScan 停止CT扫描
func (this *CTDevice) StopCTScan() error {
	this.CTRunning = false
	fmt.Printf("[CT扫描] 停止扫描, 总扫描次数: %d\n", this.CTImageCount)
	return nil
}
