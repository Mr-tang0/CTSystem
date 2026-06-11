/*
 * @Author: tang
 * @Date: 2026-05-23
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

type MotorDevice struct {
	ID  int
	plc *XinjieClient

	ctx     context.Context
	lengths map[string]float64
}

// NewMotorDevice 创建电机设备实例
func NewMotorDevice(ctx context.Context) *MotorDevice {
	return &MotorDevice{
		plc:     NewXinjieClient(),
		ctx:     ctx,
		lengths: map[string]float64{"X": 0, "Y": 0, "Z": 0, "R": 0},
	}
}

func (m *MotorDevice) Connect(ip string) error {
	err := m.plc.OpenTCP(ip, 1)
	if err != nil {
		return err
	}
	//开启子线程，循环获取电机状态
	go m.GetMotorDetails()
	runtime.EventsEmit(m.ctx, "stage_linked", map[string]bool{"stage_linked": true})

	return nil
}

func (m *MotorDevice) Disconnect() {
	m.plc.Close()
	runtime.EventsEmit(m.ctx, "stage_linked", map[string]bool{"stage_linked": false})
}

// GetLengths 获取当前各轴位置
func (m *MotorDevice) GetLengths() map[string]float64 {
	return m.lengths
}

// 连接成功后，子线程循环执行这个函数，获取电机状态
func (m *MotorDevice) GetMotorDetails() {
	for {
		X, _ := m.plc.ReadHDRegister(4, Int32)
		Y, _ := m.plc.ReadHDRegister(14, Int32)
		Z, _ := m.plc.ReadHDRegister(24, Int32)
		R, _ := m.plc.ReadHDRegister(34, Int32)

		// lengths := map[string]float64{
		// 	"X": float64(X.(int32)) / 1600.0 * 5.0,
		// 	"Y": float64(Y.(int32)) / 1600.0 * 5.0,
		// 	"Z": float64(Z.(int32)) / 1600.0 * 5.0,
		// 	"R": float64(R.(int32)) / 1600.0 * 5.0,
		// }
		m.lengths["X"] = float64(X.(int32)) / 1600.0 * 5.0
		m.lengths["Y"] = float64(Y.(int32)) / 1600.0 * 5.0
		m.lengths["Z"] = float64(Z.(int32)) / 1600.0 * 5.0
		m.lengths["R"] = float64(R.(int32)) / 1600.0 * 5.0

		// fmt.Println(X, Y, Z, R, lengths)
		runtime.EventsEmit(m.ctx, "motor_details", m.lengths)
		time.Sleep(100 * time.Millisecond)
	}
}

// 电机设置速度
func (m *MotorDevice) SetMotorSpeed(axis string, speed float32) error {
	switch axis {
	case "X":
		speed_puls := int32(float64(speed) * 1600 / 5)
		return m.plc.WriteHDRegister(2, speed_puls, Int32)
	case "Y":
		speed_puls := int32(float64(speed) * 1600 / 5)
		return m.plc.WriteHDRegister(12, speed_puls, Int32)
	case "Z":
		speed_puls := int32(float64(speed) * 1600 / 5)
		return m.plc.WriteHDRegister(22, speed_puls, Int32)
	case "R":
		speed_puls := int32(float64(speed) * 1600 / 5)
		return m.plc.WriteHDRegister(32, speed_puls, Int32)
	default:
		return fmt.Errorf("axis %s not supported", axis)
	}
}

func (m *MotorDevice) SetMotorLength(axis string, rel_length float32) error {
	switch axis {
	case "X":
		length_puls := int32(float64(rel_length) * 1600 / 5)
		return m.plc.WriteHDRegister(0, length_puls, Int32)
	case "Y":
		length_puls := int32(float64(rel_length) * 1600 / 5)
		return m.plc.WriteHDRegister(10, length_puls, Int32)
	case "Z":
		length_puls := int32(float64(rel_length) * 1600 / 5)
		return m.plc.WriteHDRegister(20, length_puls, Int32)
	case "R":
		length_puls := int32(float64(rel_length) * 1600 / 5)
		return m.plc.WriteHDRegister(30, length_puls, Int32)
	default:
		return fmt.Errorf("axis %s not supported", axis)
	}

}

func (m *MotorDevice) MotorStop(axis string) error {
	switch axis {
	case "X":
		return m.plc.WriteMCoil(20, true)
	case "Y":
		return m.plc.WriteMCoil(21, true)
	case "Z":
		return m.plc.WriteMCoil(22, true)
	case "R":
		return m.plc.WriteMCoil(23, true)
	default:
		return nil
	}
}

func (m *MotorDevice) MotorJogMove(axis string, speed float32) error {
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

func (m *MotorDevice) MotorRelMove(axis string, rel_length float32) error {
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

func (m *MotorDevice) MotorAbsMove(axis string, abs_length float32) error {
	rel_length := abs_length - float32(m.lengths[axis])
	return m.MotorRelMove(axis, rel_length)
}
