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
)

type CTDevice struct {
	CT           *sdk.NetCom
	IsConnecting bool
	comFpList    sdk.TComFpList
	ctx          context.Context
}

// NewCTDevice 创建一个新探测器设备，自动初始化SDK, 并注册事件回调函数
func NewCTDevice(ctx context.Context) *CTDevice {
	Cnet := sdk.NetCom{}

	Cnet.SetContext(ctx)
	Cnet.COM_Init()
	defer Cnet.COM_Uninit()

	return &CTDevice{
		CT:           &Cnet,
		IsConnecting: false,
		ctx:          ctx,
	}
}

// SetCallBack 设置事件回调函数
func (this *CTDevice) InitCT() {
	fmt.Println("设置事件回调函数")
	this.CT.RegisterCallback()
}

// GetDeviceNames 获取所有设备
func (this *CTDevice) GetDeviceNames() sdk.TComFpList {
	fmt.Println("获取所有设备")
	this.CT.COM_List(&this.comFpList)
	return this.comFpList
}

// Connect 连接探测器
func (this *CTDevice) Connect(cSn sdk.Char) error {
	flag := this.CT.COM_Open(cSn)
	if flag {
		this.IsConnecting = true
		return nil
	} else {
		fmt.Println("无法连接探测器: ", cSn)
		return fmt.Errorf("无法连接探测器: %s", cSn)
	}
}

func (this *CTDevice) Disconnect() error {
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

func (this *CTDevice) SoftTrigger() error {
	flag := this.CT.COM_SoftTrigger()
	if flag {
		fmt.Println("软件触发一张图片")
		return nil
	} else {
		fmt.Println("无法触发一张图片")
		return fmt.Errorf("无法触发一张图片")
	}
}
