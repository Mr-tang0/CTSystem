/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: CTSystem应用主结构体，管理设备连接和更新服务
 */
package main

import (
	update "CTSystem/backend"
	DEVICE "CTSystem/backend/components"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	updater *update.UpdateService

	// 设备实例
	ct    *DEVICE.CTDevice
	motor *DEVICE.MotorDevice
	hvps  *DEVICE.HVPSDevice
	ray   *DEVICE.RaySource160keV

	// CT扫描相关
	ctScanMutex         sync.Mutex
	ctScanRunning       bool
	ctScanPaused        bool
	ctScanStopRequested bool
	ctScanPausedAngle   float32
	ctScanImageCount    int
	ctScanPath          string

	// 图像事件通道（用于同步等待图像）
	ctImageChan chan struct{}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		hvps: DEVICE.NewHVPSDevice(),
		ray:  DEVICE.NewRaySource160keV(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//创建更新服务
	a.updater = &update.UpdateService{}

	//创建CT设备
	a.ct = DEVICE.NewCTDevice(a.ctx)

	a.motor = DEVICE.NewMotorDevice(a.ctx)
	//初始化CT
	a.ct.InitCT()

	// 监听ct_image事件
	runtime.EventsOn(ctx, "ct_image", func(optionalData ...interface{}) {
		if len(optionalData) > 0 {
			if data, ok := optionalData[0].(map[string]string); ok {
				a.handleCTImageEvent(data)
			}
		}
	})
}

func (a *App) APIUpdate() update.GitHubRelease {
	//获取更新信息
	release, err := a.updater.GetUpdateInfo()
	if err != nil {
		fmt.Printf("获取更新信息失败: %v\n", err)
		return update.GitHubRelease{}
	}
	fmt.Printf("更新信息: %v\n", release)
	return release
}

func (a *App) GetCachedRelease() update.GitHubRelease {
	return a.updater.GetCachedRelease()
}

// 与前端通信的API区域
// 8888888888888888888888888888888888888888    位移台    8888888888888888888888888888888888888888888888
// 打开位移台
func (a *App) StageOpenDevice(ip string) error {
	err := a.motor.Connect(ip)
	if err != nil {
		fmt.Printf("位移台连接失败: %v\n", err)
		return err
	}
	fmt.Printf("位移台连接成功: %s\n", ip)
	return nil
}

// 关闭位移台
func (a *App) StageCloseDevice() error {
	a.motor.Disconnect()
	fmt.Println("位移台已关闭")
	return nil
}

// 位移台轴(轴号：XYZR)点动
func (a *App) StageMoveJog(axis string, speed float32) error {
	return a.motor.MotorJogMove(axis, speed)
}

// 位移台停止运动
func (a *App) StageStop(axis string) error {
	return a.motor.MotorStop(axis)
}

// 位移台轴相对位置运动，传入轴号与距离
func (a *App) StageMoveRel(axis string, distance float32) {
	a.motor.MotorRelMove(axis, distance)
	fmt.Printf("位移台[%s]轴相对运动: %.2f\n", axis, distance)
}

// 位移台回零
func (a *App) StageMoveAbs(axis string, distance float32) {
	a.motor.MotorAbsMove(axis, distance)
	fmt.Printf("位移台[%s]轴回零\n", axis)
}

// 8888888888888888888888888888888888888888   高压电源   8888888888888888888888888888888888888888888888

// 连接高压源
func (a *App) HighVoltageConnect(ip string) string {
	err := a.hvps.ConnectHVPS(ip, 502)
	if err != nil {
		fmt.Printf("高压电源连接失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("高压电源连接成功: %s\n", ip)
	return "成功"
}

// 断开高压源
func (a *App) HighVoltageDisconnect() string {
	err := a.hvps.DisconnectHVPS()
	if err != nil {
		fmt.Printf("高压电源断开失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("高压电源已断开")
	return "成功"
}

// 设置高压源电压
func (a *App) HighVoltageSetVoltage(voltage float64) string {
	err := a.hvps.SetHV_VI(voltage, 0)
	if err != nil {
		fmt.Printf("高压电源设置电压失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("高压电源设置电压: %.2f kV\n", voltage)
	return "成功"
}

// 设置高压源电流
func (a *App) HighVoltageSetCurrent(current float64) string {
	err := a.hvps.SetHV_VI(0, current)
	if err != nil {
		fmt.Printf("高压电源设置电流失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("高压电源设置电流: %.2f uA\n", current)
	return "成功"
}

// 设置高压源电压和电流
func (a *App) HighVoltageSetVI(voltage, current float64) string {
	err := a.hvps.SetHV_VI(voltage, current)
	if err != nil {
		fmt.Printf("高压电源设置VI失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("高压电源设置VI: %.2f kV, %.2f uA\n", voltage, current)
	return "成功"
}

// 开启高压源
func (a *App) HighVoltageEnable() string {
	err := a.hvps.HV_ON()
	if err != nil {
		fmt.Printf("高压电源开启失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("高压电源已开启")
	return "成功"
}

// 关闭高压源
func (a *App) HighVoltageDisable() string {
	err := a.hvps.HV_OFF()
	if err != nil {
		fmt.Printf("高压电源关闭失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("高压电源已关闭")
	return "成功"
}

// 设置远程模式
func (a *App) HighVoltageSetRemote() string {
	err := a.hvps.SetHV_Remote()
	if err != nil {
		fmt.Printf("高压电源设置远程模式失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("高压电源已设置为远程模式")
	return "成功"
}

// 读取高压状态
func (a *App) HighVoltageReadState() string {
	state, err := a.hvps.ReadHV_State()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return fmt.Sprintf("电压: %.2f kV, 电流: %.2f uA", state.RealVol, state.RealCur)
}

// 设置灯丝参数
func (a *App) HighVoltageSetFilament(filamentPre, filamentLim float64) string {
	err := a.hvps.SetHV_Filament(filamentPre, filamentLim)
	if err != nil {
		fmt.Printf("高压电源设置灯丝失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("高压电源设置灯丝: 预设=%.2f, 限制=%.2f\n", filamentPre, filamentLim)
	return "成功"
}

// 开启灯丝
func (a *App) HighVoltageFilamentOn() string {
	err := a.hvps.FIL_ON()
	if err != nil {
		fmt.Printf("灯丝开启失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("灯丝已开启")
	return "成功"
}

// 关闭灯丝
func (a *App) HighVoltageFilamentOff() string {
	err := a.hvps.FIL_OFF()
	if err != nil {
		fmt.Printf("灯丝关闭失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("灯丝已关闭")
	return "成功"
}

// 8888888888888888888888888888888888888888   CT探测器   8888888888888888888888888888888888888888888888
// 连接探测器
func (a *App) DetectorConnect() error {
	err := a.ct.Connect()
	if err != nil {
		fmt.Printf("探测器连接失败: %v\n", err)
		return err
	}

	return nil
}

// 断开探测器
func (a *App) DetectorDisconnect() error {
	err := a.ct.Disconnect()
	if err != nil {
		fmt.Printf("探测器断开失败: %v\n", err)
		return err
	}
	return nil
}

// 获取探测器连接状态
func (a *App) DetectorIsConnected() bool {
	return a.ct.IsConnecting
}

// CT探测器连续触发
func (a *App) DetectorContinuousScan() error {
	fmt.Println("CT连续扫描触发")
	return a.ct.ContinuousTrigger()
}

// CT探测器单次触发
func (a *App) DetectorSingleScan() error {
	fmt.Println("CT单次触发")
	return a.ct.DetectorSingleScan()
}

// 设置CT探测器曝光时间
func (a *App) DetectorSetExposeTime(exposeTime int) error {
	fmt.Println("CT曝光曝光时间")
	return a.ct.DetectorSetExposeTime(exposeTime)
}

// 获取CT探测器曝光时间
func (a *App) DetectorGetExposeTime() int {
	return a.ct.DetectorGetExposeTime()
}

// 设置binning
func (a *App) DetectorSetBinning(binning string) error {
	fmt.Printf("CT设置binning: %s\n", binning)
	return a.ct.DetectorSetBinning(binning)
}

// 获取binning
func (a *App) DetectorGetBinning() string {
	return a.ct.DetectorGetBinning()
}

// 8888888888888888888888888888888888888888   射线源   8888888888888888888888888888888888888888888888

// 连接射线源
func (a *App) RaySourceConnect(com string) string {
	err := a.ray.InitSerial(com)
	if err != nil {
		fmt.Printf("射线源连接失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("射线源连接成功: %s\n", com)
	return "成功"
}

// 断开射线源
func (a *App) RaySourceDisconnect() {
	a.ray.CloseSerial()
	fmt.Println("射线源已断开")
}

// 设置射线源电压
func (a *App) RaySourceSetKV(kV float64) string {
	err := a.ray.SetKV(kV)
	if err != nil {
		fmt.Printf("射线源设置电压失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("射线源设置电压: %.2f kV\n", kV)
	return "成功"
}

// 设置射线源电流
func (a *App) RaySourceSetMA(mA int) string {
	err := a.ray.SetMA(mA)
	if err != nil {
		fmt.Printf("射线源设置电流失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("射线源设置电流: %d uA\n", mA)
	return "成功"
}

// 开启射线源
func (a *App) RaySourceEnable() string {
	err := a.ray.ControlXray("11")
	if err != nil {
		fmt.Printf("射线源开启失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("射线源已开启")
	return "成功"
}

// 关闭射线源
func (a *App) RaySourceDisable() string {
	err := a.ray.ControlXray("00")
	if err != nil {
		fmt.Printf("射线源关闭失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("射线源已关闭")
	return "成功"
}

// 读取射线源状态
func (a *App) RaySourceReadStatus() string {
	status, err := a.ray.ReadXrayStatus()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return status
}

// 读取射线源故障码
func (a *App) RaySourceReadFaultCode() string {
	fault, err := a.ray.ReadFaultCode()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return fault
}

// 读取射线源当前KV
func (a *App) RaySourceReadCurrentKV() string {
	kv, err := a.ray.ReadCurrentKV()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return kv
}

// 读取射线源当前MA
func (a *App) RaySourceReadCurrentMA() string {
	ma, err := a.ray.ReadCurrentMA()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return ma
}

// 读取射线源油温
func (a *App) RaySourceReadOilTemp() string {
	temp, err := a.ray.ReadOilTemp()
	if err != nil {
		return fmt.Sprintf("读取失败: %v", err)
	}
	return temp
}

// 读取射线源版本信息
func (a *App) RaySourceReadVersion() string {
	swVer, _ := a.ray.ReadSoftwareVersion()
	hwVer, _ := a.ray.ReadHardwareVersion()
	return fmt.Sprintf("软件版本: %s\n硬件版本: %s", swVer, hwVer)
}

// 喂射线源看门狗
func (a *App) RaySourceFeedWatchDog() string {
	err := a.ray.FeedWatchDog()
	if err != nil {
		return fmt.Sprintf("失败: %v", err)
	}
	return "成功"
}

// 清除射线源故障标志
func (a *App) RaySourceResetFaultFlag() string {
	err := a.ray.ResetFaultFlag()
	if err != nil {
		return fmt.Sprintf("失败: %v", err)
	}
	return "成功"
}

// 获取射线源连接状态
func (a *App) RaySourceIsConnected() bool {
	return a.ray.IsSerialOpen()
}

// 8888888888888888888888888888888888888888   CT联测   8888888888888888888888888888888888888888888888

// 开启CT连续扫描
func (a *App) StartCTScan(path string) error {
	fmt.Printf("[CT扫描] 开始扫描到路径: %s\n", path)
	a.ctScanMutex.Lock()
	defer a.ctScanMutex.Unlock()

	if a.ctScanRunning {
		return fmt.Errorf("CT扫描已经在运行中")
	}

	// 初始化扫描状态
	a.ctScanRunning = true
	a.ctScanPaused = false
	a.ctScanStopRequested = false
	a.ctScanPausedAngle = 0
	a.ctScanImageCount = 0
	a.ctScanPath = path

	// 创建图像事件通道（带缓冲，避免阻塞）
	a.ctImageChan = make(chan struct{}, 10)

	// 确保输出目录存在
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("[CT扫描] 创建输出目录失败: %v\n", err)
		return fmt.Errorf("创建输出目录失败: %v", err)
	}

	// 启动扫描goroutine
	go a.runCTScan()

	runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
		"running": true,
		"paused":  false,
		"message": "CT扫描已启动",
	})

	return nil
}

// 运行CT扫描循环
func (a *App) runCTScan() {
	// 获取当前R轴位置
	currentR := a.motor.GetLengths()["R"]
	fmt.Printf("[CT扫描] 当前R位置: %.2f°\n", currentR)

	// 计算需要扫描的角度数量
	anglesToScan := int(360.0 - currentR)
	fmt.Printf("[CT扫描] 需要扫描 %d 个角度\n", anglesToScan)

	for i := 0; i < anglesToScan; i++ {
		a.ctScanMutex.Lock()
		if a.ctScanStopRequested {
			a.ctScanRunning = false
			a.ctScanMutex.Unlock()
			fmt.Println("[CT扫描] 扫描已停止")
			runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
				"running": false,
				"paused":  false,
				"message": "CT扫描已停止",
			})
			return
		}

		if a.ctScanPaused {
			a.ctScanMutex.Unlock()
			fmt.Println("[CT扫描] 扫描已暂停")
			runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
				"running": false,
				"paused":  true,
				"message": "CT扫描已暂停",
			})
			return
		}
		a.ctScanMutex.Unlock()

		// 运动R轴1度
		fmt.Printf("[CT扫描] 运动R轴1度 (第%d次)\n", i+1)
		err := a.motor.MotorRelMove("R", 1.0)
		if err != nil {
			fmt.Printf("[CT扫描] 运动R轴失败: %v\n", err)
			continue
		}

		// 等待运动完成
		time.Sleep(1 * time.Second)

		// 执行单次扫描
		fmt.Printf("[CT扫描] 执行单次扫描 (第%d次)\n", i+1)
		err = a.ct.DetectorSingleScan()
		if err != nil {
			fmt.Printf("[CT扫描] 扫描失败: %v\n", err)
			continue
		}

		// 等待图像事件（最多等待10秒）
		fmt.Printf("[CT扫描] 等待图像 (第%d次)\n", i+1)
		select {
		case <-a.ctImageChan:
			// 图像已收到，继续下一次循环
			fmt.Printf("[CT扫描] 图像已收到 (第%d次)\n", i+1)
		case <-time.After(10 * time.Second):
			// 超时，报错并结束扫描
			fmt.Printf("[CT扫描] 等待图像超时(10秒)，结束扫描\n")
			a.ctScanMutex.Lock()
			a.ctScanRunning = false
			a.ctScanMutex.Unlock()
			runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
				"running": false,
				"paused":  false,
				"message": "CT扫描失败：等待图像超时",
			})
			return
		}
	}

	a.ctScanMutex.Lock()
	a.ctScanRunning = false
	a.ctScanMutex.Unlock()

	fmt.Println("[CT扫描] 扫描完成")
	runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
		"running": false,
		"paused":  false,
		"message": "CT扫描已完成",
	})
}

// 暂停CT扫描
func (a *App) CTPauseScan() error {
	a.ctScanMutex.Lock()
	defer a.ctScanMutex.Unlock()

	if !a.ctScanRunning {
		return fmt.Errorf("CT扫描未在运行")
	}

	if a.ctScanPaused {
		return fmt.Errorf("CT扫描已经暂停")
	}

	// 记录当前R位置
	currentR := a.motor.GetLengths()["R"]
	a.ctScanPausedAngle = float32(currentR)
	fmt.Printf("[CT扫描] 暂停扫描，记录R位置: %.2f°\n", a.ctScanPausedAngle)

	a.ctScanPaused = true

	runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
		"running": false,
		"paused":  true,
		"message": "CT扫描已暂停",
	})

	return nil
}

// 继续CT扫描
func (a *App) ContinueCTScan() error {
	a.ctScanMutex.Lock()
	defer a.ctScanMutex.Unlock()

	if !a.ctScanPaused {
		return fmt.Errorf("CT扫描未暂停")
	}

	// 恢复扫描状态
	a.ctScanPaused = false
	a.ctScanRunning = true

	// 启动扫描goroutine
	go a.runCTScan()

	runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
		"running": true,
		"paused":  false,
		"message": "CT扫描已继续",
	})

	return nil
}

// 停止CT扫描
func (a *App) StopCTScan() error {
	a.ctScanMutex.Lock()
	defer a.ctScanMutex.Unlock()

	if !a.ctScanRunning && !a.ctScanPaused {
		return fmt.Errorf("CT扫描未在运行")
	}

	fmt.Println("[CT扫描] 请求停止扫描")
	a.ctScanStopRequested = true
	a.ctScanPaused = false
	a.ctScanRunning = false

	runtime.EventsEmit(a.ctx, "ct_scan_status", map[string]interface{}{
		"running": false,
		"paused":  false,
		"message": "CT扫描已停止",
	})

	return nil
}

// 处理CT图像事件
func (a *App) handleCTImageEvent(data map[string]string) {
	var shouldSignal bool

	a.ctScanMutex.Lock()
	// 只有在扫描运行时才保存图像
	if !a.ctScanRunning || a.ctScanPaused {
		a.ctScanMutex.Unlock()
		return
	}

	imageBase64 := data["image"]
	width := data["width"]
	height := data["height"]

	// 解码Base64图像
	imageData, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(imageBase64, "data:image/jpeg;base64,"))
	if err != nil {
		fmt.Printf("[CT扫描] 解码图像失败: %v\n", err)
		a.ctScanMutex.Unlock()
		return
	}

	// 生成文件名
	a.ctScanImageCount++
	filename := fmt.Sprintf("ct_scan_%04d.jpg", a.ctScanImageCount)
	filepath := filepath.Join(a.ctScanPath, filename)

	// 保存图像
	err = ioutil.WriteFile(filepath, imageData, 0644)
	if err != nil {
		fmt.Printf("[CT扫描] 保存图像失败: %v\n", err)
		a.ctScanMutex.Unlock()
		return
	}

	fmt.Printf("[CT扫描] 已保存图像: %s (尺寸: %sx%s)\n", filename, width, height)

	// 发送保存进度事件
	runtime.EventsEmit(a.ctx, "ct_scan_progress", map[string]interface{}{
		"count":    a.ctScanImageCount,
		"filename": filename,
		"filepath": filepath,
		"width":    width,
		"height":   height,
	})

	shouldSignal = true
	a.ctScanMutex.Unlock()

	// 在互斥锁外发送信号，避免阻塞
	if shouldSignal && a.ctImageChan != nil {
		select {
		case a.ctImageChan <- struct{}{}:
			// 信号已发送
		default:
			// 通道已满或未创建，忽略
		}
	}
}
