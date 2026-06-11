/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 探测器设备管理模块，负责探测器连接和事件处理
 */
package components

import (
	sdk "CTSystem/backend/CSDK"
	"context"
	"fmt"
	"time"
)

type CTMODE int

const (
	IDLE CTMODE = iota
	HST
	AED1
	AED2
	RECOVER
	ERR
)

type CTStatus struct {
	Mode          CTMODE // 当前工作模式
	ExposeTime    int    // 曝光时间
	Binning       string // 当前binning
	TriggerRepeat int    // 手动触发图片张数
	TriggerDelay  int    // 手动触发图片间隔
}

type CTDevice struct {
	CT           *sdk.NetCom
	IsConnecting bool
	comFpList    sdk.TComFpList
	ctx          context.Context

	status CTStatus
}

// NewCTDevice 创建一个新探测器设备，自动初始化SDK, 并注册事件回调函数
func NewCTDevice(ctx context.Context) *CTDevice {
	Cnet := sdk.NewNetCom(ctx)

	return &CTDevice{
		CT:           Cnet,
		IsConnecting: false,
		ctx:          ctx,
		status: CTStatus{
			Mode:          IDLE,
			ExposeTime:    0,
			Binning:       "null",
			TriggerRepeat: 0,
			TriggerDelay:  0,
		},
	}
}

// SetCallBack 设置事件回调函数
func (this *CTDevice) InitCT() {
	fmt.Println("初始化探测器设备")
	this.CT.RegisterCallback()     // 先注册回调
	this.CT.COM_Init()             // 再初始化SDK
	this.CT.COM_StartNet()         // 启动网络监听
	this.CT.COM_SetCalibMode(0x06) // 设置校准模式 IMG_CALIB_GAIN | IMG_CALIB_DEFECT
}

// GetDeviceNames 获取所有设备
func (this *CTDevice) GetDeviceNames() sdk.TComFpList {
	this.CT.COM_List(&this.comFpList)
	return this.comFpList
}

// Connect 连接探测器
func (this *CTDevice) Connect() error {
	flag := this.CT.COM_Open(0)
	if flag {
		this.IsConnecting = true
		return nil
	} else {
		fmt.Println("无法连接探测器: ")
		return fmt.Errorf("无法连接探测器")
	}
}

func (this *CTDevice) Disconnect() error {
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

func (this *CTDevice) IsConnected() bool {
	return this.IsConnecting
}

func (this *CTDevice) ContinuousTrigger() error {
	// if(this.mode != HST){
	// 	this.CT.COM_HstAcq()
	// 	time.Sleep(2000*time.Millisecond)
	// }
	// this.CT.COM_ExposeReq()
	// if !flag {
	// 	fmt.Println("无法获取图片")
	// 	return fmt.Errorf("无法获取图片")
	// }
	// flag = this.CT.COM_ExposeReq()
	// if !flag {
	// 	fmt.Println("无法触发图片")
	// 	return fmt.Errorf("无法触发图片")
	// }
	// fmt.Println("开启连续触发")
	return nil
}

func (this *CTDevice) DetectorSingleScan() error {
	if this.status.Mode != HST {
		this.CT.COM_HstAcq()
		time.Sleep(2000 * time.Millisecond)
		this.status.Mode = HST
		fmt.Println("[HST] 启动HST")
	}
	this.CT.COM_ExposeReq()
	fmt.Println("[HST] 获取图片")
	return nil
}

func (this *CTDevice) DetectorSetExposeTime(exposeTime int) error {
	if !this.CT.COM_SetExposeTime(exposeTime) {
		fmt.Println("无法设置曝光时间")
		return fmt.Errorf("无法设置曝光时间")
	}
	this.status.ExposeTime = exposeTime
	fmt.Println("设置曝光时间成功")
	return nil
}

func (this *CTDevice) DetectorGetExposeTime() int {
	this.status.ExposeTime = this.CT.COM_GetExposeTime()
	return this.status.ExposeTime
}

func (this *CTDevice) DetectorSetBinning(binning string) error {
	if !this.CT.COM_SetBinning(binning) {
		fmt.Println("无法设置像素数", binning)
		return fmt.Errorf("无法设置像素数")
	}
	this.status.Binning = binning
	fmt.Println("设置像素数成功")
	return nil
}

func (this *CTDevice) DetectorGetBinning() string {
	this.status.Binning = this.CT.COM_GetBinning()
	return this.status.Binning
}
