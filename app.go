/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: CTSystem应用主结构体，管理设备连接和更新服务
 */
package main

import (
	update "CTSystem/backend"
	sdk "CTSystem/backend/CSDK"
	DEVICE "CTSystem/backend/components"
	"context"
	"fmt"
	"time"
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
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		motor: DEVICE.NewMotorDevice(),
		hvps:  DEVICE.NewHVPSDevice(),
		ray:   DEVICE.NewRaySource160keV(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//创建更新服务
	a.updater = &update.UpdateService{}

	//创建CT设备
	a.ct = DEVICE.NewCTDevice(a.ctx)
	//初始化CT
	a.ct.InitCT()
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
func (a *App) StageOpenDevice(ip string) string {
	err := a.motor.Connect(ip)
	if err != nil {
		fmt.Printf("位移台连接失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("位移台连接成功: %s\n", ip)
	return "成功"
}

// 关闭位移台
func (a *App) StageCloseDevice() {
	a.motor.Disconnect()
	fmt.Println("位移台已关闭")
}

// 位移台轴(轴号：XYZR)点动
func (a *App) StageAxisPulse(axis string, dir bool) {
	speed := float32(1.0)
	if !dir {
		speed = -1.0
	}
	a.motor.MotorJogMove(axis, speed)
	fmt.Printf("位移台[%s]轴点动: %s\n", axis, map[bool]string{true: "正向", false: "反向"}[dir])
}

// 位移台轴相对位置运动，传入轴号与距离
func (a *App) StageMoveRel(axis string, distance float64) {
	a.motor.MotorRelMove(axis, float32(distance), float32(100))
	fmt.Printf("位移台[%s]轴相对运动: %.2f\n", axis, distance)
}

// 位移台停止运动
func (a *App) StageStop(axis string) {
	a.motor.MotorStop(axis)
	fmt.Printf("位移台[%s]轴停止\n", axis)
}

// 位移台回零
func (a *App) StageAxisAbs(axis string) {
	a.motor.MotorAbs(axis)
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

// 获取探测器列表
func (a *App) DetectorGetList() []string {
	// devices := a.ct.GetDeviceNames()
	// var names []string
	// for _, dev := range devices.FpList {
	// 	names = append(names, fmt.Sprintf("%s", dev.SzName))
	// }
	return []string{}
}

// 连接探测器
func (a *App) DetectorConnect(index int) string {
	err := a.ct.Connect(sdk.Char(index))
	if err != nil {
		fmt.Printf("探测器连接失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Printf("探测器连接成功: index=%d\n", index)
	return "成功"
}

// 断开探测器
func (a *App) DetectorDisconnect() string {
	err := a.ct.Disconnect()
	if err != nil {
		fmt.Printf("探测器断开失败: %v\n", err)
		return fmt.Sprintf("失败: %v", err)
	}
	fmt.Println("探测器已断开")
	return "成功"
}

// 探测器软件触发一张
func (a *App) DetectorSoftwareTrigger() {
	fmt.Println("探测器软件触发")
	// TODO: 实现触发逻辑
}

// 获取探测器连接状态
func (a *App) DetectorIsConnected() bool {
	return a.ct.IsConnecting
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

// 开启CT连续扫描
func (a *App) DetectorEnableContinuousScan() {
	//开启循环，初始采集一张探测器图片，后将R轴运动1度，再采集一张图片，直到R轴运动360度
	for i := 0; i < 360; i++ {
		// 采集一张图片
		a.motor.MotorRelMove("R", 1, 1)
		time.Sleep(time.Second) //等待R轴运动完成
		a.DetectorSoftwareTrigger()
		time.Sleep(2 * time.Second) //等待探测器软件触发完成
	}
	fmt.Println("CT连续扫描已开启")
}
