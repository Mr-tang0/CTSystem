/*
 * @Author: tang
 * @Date: 2026-06-12
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 位移台电机设备控制模块，通过Modbus协议控制电机运动
 */
package components

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	XResolution float64 = 200
	YResolution float64 = 200
	ZResolution float64 = 200
	RResolution float64 = 1000
)

const (
	// 电机相对运动使能
	XREL uint16 = 0
	YREL uint16 = 1
	ZREL uint16 = 2
	RREL uint16 = 3

	// 电机JOG运动使能
	XJOG uint16 = 10
	YJOG uint16 = 11
	ZJOG uint16 = 12
	RJOG uint16 = 13

	// 电机停止
	XSTOP uint16 = 20
	YSTOP uint16 = 21
	ZSTOP uint16 = 22
	RSTOP uint16 = 23

	//电机当前位置HDRegister
	XPOS uint16 = 4
	YPOS uint16 = 14
	ZPOS uint16 = 24
	RPOS uint16 = 34

	// 电机速度HDRegister
	XSPEED uint16 = 2
	YSPEED uint16 = 12
	ZSPEED uint16 = 22
	RSPEED uint16 = 32

	//电机相对运动HDRegister
	XLEN uint16 = 0
	YLEN uint16 = 10
	ZLEN uint16 = 20
	RLEN uint16 = 30
)

type MotorDevice struct {
	ID  int
	plc *XinjieClient

	ctx     context.Context
	lengths map[string]float64
}

// NewMotorDevice 创建电机设备实例
func NewMotorDevice() *MotorDevice {
	return &MotorDevice{
		plc:     NewXinjieClient(),
		lengths: map[string]float64{"X": 0, "Y": 0, "Z": 0, "R": 0},
	}
}

// SetContent 设置上下文，用于事件触发
func (m *MotorDevice) SetContent(ctx context.Context) {
	m.ctx = ctx
}

// StageOpenDevice 打开位移台设备
func (m *MotorDevice) StageOpenDevice(ip string) error {
	fmt.Println("连接位移台", m.ctx == nil)
	err := m.plc.OpenTCP(ip, 1)
	if err != nil {
		return err
	}
	//开启子线程，循环获取电机状态
	go m.GetMotorDetails()
	runtime.EventsEmit(m.ctx, "stage_linked", map[string]bool{"stage_linked": true})

	return nil
}

// StageCloseDevice 关闭位移台设备
func (m *MotorDevice) StageCloseDevice() {
	m.plc.Close()
	runtime.EventsEmit(m.ctx, "stage_linked", map[string]bool{"stage_linked": false})
}

// GetLengths 获取当前各轴位置
func (m *MotorDevice) GetLengths() map[string]float64 {
	return m.lengths
}

// 连接成功后，子线程循环执行函数，获取电机状态
func (m *MotorDevice) GetMotorDetails() {
	for {
		X, _ := m.plc.ReadHDRegister(XPOS, Int32)
		Y, _ := m.plc.ReadHDRegister(YPOS, Int32)
		Z, _ := m.plc.ReadHDRegister(ZPOS, Int32)
		R, _ := m.plc.ReadHDRegister(RPOS, Int32)

		m.lengths["X"] = float64(int32(float64(X.(int32))/XResolution*1000)) / 1000.0
		m.lengths["Y"] = float64(int32(float64(Y.(int32))/YResolution*1000)) / 1000.0
		m.lengths["Z"] = float64(int32(float64(Z.(int32))/ZResolution*1000)) / 1000.0
		m.lengths["R"] = float64(int32(float64(R.(int32))/RResolution*1000)) / 1000.0

		// fmt.Println(X, Y, Z, R, lengths)
		runtime.EventsEmit(m.ctx, "motor_details", m.lengths)
		time.Sleep(100 * time.Millisecond)
	}
}

// SetMotorSpeed 设置电机速度
func (m *MotorDevice) SetMotorSpeed(axis string, speed float32) error {
	switch axis {
	case "X":
		speed_puls := int32(float64(speed) * XResolution)
		return m.plc.WriteHDRegister(XSPEED, speed_puls, Int32)
	case "Y":
		speed_puls := int32(float64(speed) * YResolution)
		return m.plc.WriteHDRegister(YSPEED, speed_puls, Int32)
	case "Z":
		speed_puls := int32(float64(speed) * ZResolution)
		return m.plc.WriteHDRegister(ZSPEED, speed_puls, Int32)
	case "R":
		speed_puls := int32(float64(speed) * RResolution)
		return m.plc.WriteHDRegister(RSPEED, speed_puls, Int32)
	default:
		return fmt.Errorf("axis %s not supported", axis)
	}
}

// SetMotorLength 设置电机相对运动长度
func (m *MotorDevice) SetMotorLength(axis string, rel_length float32) error {
	switch axis {
	case "X":
		length_puls := int32(float64(rel_length) * XResolution)
		return m.plc.WriteHDRegister(XLEN, length_puls, Int32)
	case "Y":
		length_puls := int32(float64(rel_length) * YResolution)
		return m.plc.WriteHDRegister(YLEN, length_puls, Int32)
	case "Z":
		length_puls := int32(float64(rel_length) * ZResolution)
		return m.plc.WriteHDRegister(ZLEN, length_puls, Int32)
	case "R":
		length_puls := int32(float64(rel_length) * RResolution)
		return m.plc.WriteHDRegister(RLEN, length_puls, Int32)
	default:
		return fmt.Errorf("axis %s not supported", axis)
	}
}

// StageStop 停止电机运动
func (m *MotorDevice) StageStop(axis string) error {
	switch axis {
	case "X":
		return m.plc.WriteMCoil(XSTOP, true)
	case "Y":
		return m.plc.WriteMCoil(YSTOP, true)
	case "Z":
		return m.plc.WriteMCoil(ZSTOP, true)
	case "R":
		return m.plc.WriteMCoil(RSTOP, true)
	default:
		return nil
	}
}

// StageMoveJog 电机JOG运动
func (m *MotorDevice) StageMoveJog(axis string, speed float32) error {
	// fmt.Println(axis, speed)
	switch axis {
	case "X":
		m.SetMotorSpeed(axis, speed)
		return m.plc.WriteMCoil(10, true)
	case "Y":
		m.SetMotorSpeed(axis, speed)
		return m.plc.WriteMCoil(11, true)
	case "Z":
		m.SetMotorSpeed(axis, speed)
		return m.plc.WriteMCoil(12, true)
	case "R":
		m.SetMotorSpeed(axis, speed)
		return m.plc.WriteMCoil(13, true)
	default:
		return nil
	}
}

// StageMoveRel 电机相对运动
func (m *MotorDevice) StageMoveRel(axis string, rel_length float32) error {
	// fmt.Println(axis, rel_length)
	switch axis {
	case "X":
		m.SetMotorLength(axis, rel_length)
		return m.plc.WriteMCoil(0, true)
	case "Y":
		m.SetMotorLength(axis, rel_length)
		return m.plc.WriteMCoil(1, true)
	case "Z":
		m.SetMotorLength(axis, rel_length)
		return m.plc.WriteMCoil(2, true)
	case "R":
		m.SetMotorLength(axis, rel_length)
		return m.plc.WriteMCoil(3, true)
	default:
		return fmt.Errorf("axis %s not supported", axis)
	}
}

// StageMoveAbs 电机绝对运动
func (m *MotorDevice) StageMoveAbs(axis string, abs_length float32) error {
	// fmt.Println(axis, abs_length)
	rel_length := abs_length - float32(m.lengths[axis])
	return m.StageMoveRel(axis, rel_length)
}
