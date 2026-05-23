/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 高压电源设备控制模块，实现Modbus协议的高压电源通信
 */
package components

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

// HVPSDevice 高压电源网口通信类
type HVPSDevice struct {
	conn net.Conn

	// 高压电源接收数据缓冲区
	hvRecv    []byte
	hvRecvLen int

	// 电压电流设定值
	volValue    int
	checkSumVol int
	curValue    int
	checkSumCur int

	// 高压电源实际值
	hvRealVol int
	hvRealCur int

	// 灯丝参数
	filPre         int
	checkSumFilPre int
	filLim         int
	checkSumFilLim int

	// 灯丝实际值
	filRealVol int
	filRealCur int

	// 错误状态
	errCheck  int
	hvpsNoErr bool
}

// NewHVPSDevice 创建高压电源设备实例
func NewHVPSDevice() *HVPSDevice {
	return &HVPSDevice{
		hvRecv: make([]byte, 1024),
	}
}

// ConnectHVPS 连接高压电源设备
func (h *HVPSDevice) ConnectHVPS(HVPS_ip string, HVPS_port int) error {
	addr := fmt.Sprintf("%s:%d", HVPS_ip, HVPS_port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	h.conn = conn
	return nil
}

// DisconnectHVPS 断开连接
func (h *HVPSDevice) DisconnectHVPS() error {
	if h.conn != nil {
		err := h.conn.Close()
		h.conn = nil
		return err
	}
	return errors.New("未连接")
}

// sendHVData 发送数据
func (h *HVPSDevice) sendHVData(HV_packet []byte) error {
	if h.conn == nil {
		return errors.New("未连接")
	}
	_, err := h.conn.Write(HV_packet)
	return err
}

// recvHVData 接收数据
func (h *HVPSDevice) recvHVData(HV_buf []byte) (int, error) {
	if h.conn == nil {
		return 0, errors.New("未连接")
	}
	return h.conn.Read(HV_buf)
}

// popcount 计算二进制中1的个数（替代C++的_mm_popcnt_u32）
func popcount(x uint32) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

// setHV_Remote 设置远程模式
func (h *HVPSDevice) SetHV_Remote() error {
	remCmd := []byte{0xA1, 0x65, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x0D}
	if err := h.sendHVData(remCmd); err != nil {
		return err
	}

	// 清空接收缓冲区
	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}

	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}

	// 高压电源返回数据：如果长度为11且第6个字节为0x01，说明远程模式设置成功
	if hvRecvLen == 11 && h.hvRecv[5] == 1 {
		return nil
	}
	return errors.New("远程模式设置失败")
}

// setHV_VI 设置高压电源的电压电流值
func (h *HVPSDevice) SetHV_VI(setvoltage_kV, setcurrent_uA float64) error {
	// 转换成协议要求的整数值
	h.volValue = int(setvoltage_kV*4095/50 + 0.5)   // 电压最大为50 kV
	h.curValue = int(setcurrent_uA*4095/1000 + 0.5) // 电流最大为1000 uA

	// 计算校验和
	h.checkSumVol = popcount(0xA1) + popcount(0x61) + popcount(uint32(h.volValue))
	h.checkSumCur = popcount(0xA1) + popcount(0x62) + popcount(uint32(h.curValue))

	// 构建电压指令
	volArray := make([]byte, 11)
	volArray[0] = 0xA1
	volArray[1] = 0x61
	volArray[2] = 0x00
	volArray[3] = 0x00
	volArray[4] = 0x00
	binary.BigEndian.PutUint16(volArray[5:7], uint16(h.volValue))
	binary.BigEndian.PutUint16(volArray[7:9], uint16(h.checkSumVol))
	volArray[9] = 0x00
	volArray[10] = 0x0D

	// 构建电流指令
	curArray := make([]byte, 11)
	curArray[0] = 0xA1
	curArray[1] = 0x62
	curArray[2] = 0x00
	curArray[3] = 0x00
	curArray[4] = 0x00
	binary.BigEndian.PutUint16(curArray[5:7], uint16(h.curValue))
	binary.BigEndian.PutUint16(curArray[7:9], uint16(h.checkSumCur))
	curArray[9] = 0x00
	curArray[10] = 0x0D

	// 发送电压指令
	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	if err := h.sendHVData(volArray); err != nil {
		return err
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}
	if hvRecvLen < 0 || h.hvRecv[2] != 10 {
		return errors.New("电压设置失败")
	}

	// 发送电流指令
	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	if err := h.sendHVData(curArray); err != nil {
		return err
	}
	hvRecvLen, err = h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}
	if hvRecvLen < 0 || h.hvRecv[2] != 10 {
		return errors.New("电流设置失败")
	}

	return nil
}

// HV_ON 打开高压电源输出
func (h *HVPSDevice) HV_ON() error {
	hvOn := []byte{0xA1, 0x69, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x0D}

	if err := h.sendHVData(hvOn); err != nil {
		return err
	}

	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}

	if hvRecvLen > 0 && h.hvRecv[5] == 1 {
		return nil
	}
	return errors.New("HV_ON命令执行失败")
}

// HV_OFF 关闭高压电源输出
func (h *HVPSDevice) HV_OFF() error {
	hvOff := []byte{0xA1, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x0D}

	if err := h.sendHVData(hvOff); err != nil {
		return err
	}

	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}

	if hvRecvLen > 0 && h.hvRecv[5] == 0 {
		return nil
	}
	return errors.New("HV_OFF命令执行失败")
}

// setHV_Filament 设置灯丝预设电流值及限制值
func (h *HVPSDevice) SetHV_Filament(filament_pre, filament_lim float64) error {
	// 转换成协议要求的整数值
	h.filPre = int(filament_pre*4095/10 + 0.5) // 电流最大为10 A
	h.filLim = int(filament_lim*4095/10 + 0.5) // 电流最大为10 A

	// 计算校验和
	h.checkSumFilPre = popcount(0xA1) + popcount(0x63) + popcount(uint32(h.filPre))
	h.checkSumFilLim = popcount(0xA1) + popcount(0x64) + popcount(uint32(h.filLim))

	// 构建灯丝电流极限值指令
	filLimArray := make([]byte, 11)
	filLimArray[0] = 0xA1
	filLimArray[1] = 0x64
	filLimArray[2] = 0x00
	filLimArray[3] = 0x00
	filLimArray[4] = 0x00
	binary.BigEndian.PutUint16(filLimArray[5:7], uint16(h.filLim))
	binary.BigEndian.PutUint16(filLimArray[7:9], uint16(h.checkSumFilLim))
	filLimArray[9] = 0x00
	filLimArray[10] = 0x0D

	// 构建灯丝电流设定值指令
	filPreArray := make([]byte, 11)
	filPreArray[0] = 0xA1
	filPreArray[1] = 0x63
	filPreArray[2] = 0x00
	filPreArray[3] = 0x00
	filPreArray[4] = 0x00
	binary.BigEndian.PutUint16(filPreArray[5:7], uint16(h.filPre))
	binary.BigEndian.PutUint16(filPreArray[7:9], uint16(h.checkSumFilPre))
	filPreArray[9] = 0x00
	filPreArray[10] = 0x0D

	// 发送灯丝极限电流指令
	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	if err := h.sendHVData(filLimArray); err != nil {
		return err
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}
	if hvRecvLen < 0 || h.hvRecv[2] != 10 {
		return errors.New("灯丝极限电流设置失败")
	}

	// 发送灯丝电流指令
	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	if err := h.sendHVData(filPreArray); err != nil {
		return err
	}
	hvRecvLen, err = h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}
	if hvRecvLen < 0 || h.hvRecv[2] != 10 {
		return errors.New("灯丝电流设置失败")
	}

	return nil
}

// FIL_ON 打开灯丝输出
func (h *HVPSDevice) FIL_ON() error {
	filOn := []byte{0xA1, 0x70, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x07, 0x00, 0x0D}

	if err := h.sendHVData(filOn); err != nil {
		return err
	}

	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}

	if hvRecvLen > 0 && h.hvRecv[5] == 1 {
		return nil
	}
	return errors.New("FIL_ON命令执行失败")
}

// FIL_OFF 关闭灯丝输出
func (h *HVPSDevice) FIL_OFF() error {
	filOff := []byte{0xA1, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x0D}

	if err := h.sendHVData(filOff); err != nil {
		return err
	}

	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return err
	}

	if hvRecvLen > 0 && h.hvRecv[5] == 0 {
		return nil
	}
	return errors.New("FIL_OFF命令执行失败")
}

// HVState 高压电源状态
type HVState struct {
	RealVol float64
	RealCur float64
	Error   error
}

// ReadHV_State 读取高压电源状态
func (h *HVPSDevice) ReadHV_State() (*HVState, error) {
	readHV := []byte{0xA1, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x0D}

	if err := h.sendHVData(readHV); err != nil {
		return nil, err
	}

	for i := range h.hvRecv {
		h.hvRecv[i] = 0
	}
	hvRecvLen, err := h.recvHVData(h.hvRecv)
	if err != nil {
		return nil, err
	}

	if hvRecvLen == 19 && h.hvRecv[2] == 10 {
		h.errCheck = int(h.hvRecv[3] % 8)
		if h.errCheck == 0 {
			h.hvpsNoErr = true
			h.hvRealVol = int((int(h.hvRecv[5])*256 + int(h.hvRecv[6])) * 50 * 100 / 4095)
			h.hvRealCur = int((int(h.hvRecv[7])*256 + int(h.hvRecv[8])) * 1000 * 100 / 4095)

			return &HVState{
				RealVol: float64(h.hvRealVol) / 100,
				RealCur: float64(h.hvRealCur) / 100,
			}, nil
		} else {
			h.hvpsNoErr = false
			return nil, errors.New(fmt.Sprintf("电源错误: %d", -int(h.hvRecv[3])))
		}
	}
	return nil, errors.New("读取状态失败")
}

// FILState 灯丝状态
type FILState struct {
	RealVol float64
	RealCur float64
}

// ReadFIL_State 读取灯丝状态（必须紧跟ReadHV_State调用）
func (h *HVPSDevice) ReadFIL_State() (*FILState, error) {
	if !h.hvpsNoErr {
		return nil, errors.New("高压电源状态错误")
	}

	h.filRealVol = int(float64(int(h.hvRecv[11])*256+int(h.hvRecv[12])) * 5.5 * 1000 / 4095)
	h.filRealCur = int(float64(int(h.hvRecv[9])*256+int(h.hvRecv[10])) * 3.6 * 1000 / 4095)

	return &FILState{
		RealVol: float64(h.filRealVol) / 1000,
		RealCur: float64(h.filRealCur) / 1000,
	}, nil
}

// IsConnected 检查是否已连接
func (h *HVPSDevice) IsConnected() bool {
	return h.conn != nil
}
