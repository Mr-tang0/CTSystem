/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 160keV射线源串口通信模块，实现自定义协议的射线源控制
 */
package components

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.bug.st/serial"
)

// 通讯协议相关定义
const (
	STX              = 0x02 // 帧头
	ETX              = 0x03 // 帧尾
	SEPARATOR        = 0x2C // 间隔符（逗号）
	DEFAULT_BAUDRATE = 9600
	RESPONSE_DELAY   = 200 // 响应延迟（ms）
	COMM_INTERVAL    = 100 // 通讯间隔（ms）
)

// 命令码定义
type CmdCode int

const (
	SET_KV               CmdCode = 10 // 设置kV电压
	SET_MA               CmdCode = 11 // 设置mA电流
	READ_KV_SET          CmdCode = 14 // 读取kV设定值
	READ_MA_SET          CmdCode = 15 // 读取mA设定值
	READ_MODULE_STATUS   CmdCode = 22 // 读取故障码
	READ_CODE_VERSION    CmdCode = 23 // 读取软件版本号
	READ_HARDWRE_VERSION CmdCode = 24 // 读取硬件版本号
	READ_SERIAL_NO       CmdCode = 25 // 读取出厂编码
	SERVICE_WATCHDOG     CmdCode = 27 // 喂通讯看门狗
	WATCHDOG_ENA         CmdCode = 28 // 使能看门狗
	RESET_FAULT_FLAG     CmdCode = 32 // 清除故障标志位
	READ_NOW_KV          CmdCode = 60 // 读取当前kV值
	READ_NOW_MA          CmdCode = 61 // 读取当前mA值
	READ_OIL_TEMP        CmdCode = 63 // 读取油温值
	READ_TOTAL_ARC_COUNT CmdCode = 72 // 读取设备总打火次数
	READ_LAST_ARC_COUNT  CmdCode = 73 // 读取最近打火次数
	READ_XRAY_STATUS     CmdCode = 96 // 读取射线源状态
	XRAY_ONOFF           CmdCode = 99 // 打开或关闭射线源
)

// 射线源状态定义
type XRayStatus int

const (
	XRAY_OFF  XRayStatus = 0 // 关闭
	XRAY_ON   XRayStatus = 1 // 打开
	XRAY_WARM XRayStatus = 2 // 预热
)

// 故障码位定义
const (
	FAULT_LOCK         = 0  // B0：闭锁保护
	FAULT_OCP_LCC      = 1  // B1：OCP_LCC保护
	FAULT_KV_OVER      = 2  // B2：kV过压保护
	FAULT_KV_UNDER     = 3  // B3：kV欠压保护
	FAULT_ARC_LOCK     = 4  // B4：打火锁死保护
	FAULT_INPUT_VOLT   = 5  // B5：输入过欠压保护
	FAULT_TEMP         = 6  // B6：温度保护
	FAULT_REMOTE       = 7  // B7：Remote保护
	FAULT_MA_OVER      = 8  // B8：mA过流保护
	FAULT_MA_UNDER     = 9  // B9：mA欠流保护
	FAULT_POWER_OVER   = 10 // B10：过功率保护
	FAULT_WATCHDOG     = 11 // B11：看门狗溢出保护
	FAULT_XRAY_OFF     = 12 // B12：Xray关闭
	FAULT_COMM_TIMEOUT = 13 // B13：通讯超时保护
	FAULT_FAST         = 14 // B14：快速保护
	FAULT_FAULT_LOCK   = 15 // B15：故障锁死保护
)

// 故障码描述映射
var faultDescriptions = map[int]string{
	FAULT_LOCK:         "闭锁保护",
	FAULT_OCP_LCC:      "OCP_LCC保护",
	FAULT_KV_OVER:      "kV过压保护",
	FAULT_KV_UNDER:     "kV欠压保护",
	FAULT_ARC_LOCK:     "打火锁死保护",
	FAULT_INPUT_VOLT:   "输入过欠压保护",
	FAULT_TEMP:         "温度保护",
	FAULT_REMOTE:       "Remote保护",
	FAULT_MA_OVER:      "mA过流保护",
	FAULT_MA_UNDER:     "mA欠流保护",
	FAULT_POWER_OVER:   "过功率保护",
	FAULT_WATCHDOG:     "看门狗溢出保护",
	FAULT_XRAY_OFF:     "Xray关闭",
	FAULT_COMM_TIMEOUT: "通讯超时保护",
	FAULT_FAST:         "快速保护",
	FAULT_FAULT_LOCK:   "故障锁死保护",
}

// RaySource160keV 射线源通讯核心类
type RaySource160keV struct {
	port serial.Port
	logs []string
}

// NewRaySource160keV 创建射线源实例
func NewRaySource160keV() *RaySource160keV {
	return &RaySource160keV{
		logs: make([]string, 0),
	}
}

// InitSerial 初始化串口
func (r *RaySource160keV) InitSerial(com string) error {
	if r.port != nil {
		r.CloseSerial()
	}

	mode := &serial.Mode{
		BaudRate: DEFAULT_BAUDRATE,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(com, mode)
	if err != nil {
		r.addLog(fmt.Sprintf("打开串口失败: %s", com))
		return err
	}

	r.port = port
	r.addLog(fmt.Sprintf("串口打开成功: %s", com))
	return nil
}

// CloseSerial 关闭串口
func (r *RaySource160keV) CloseSerial() {
	if r.port != nil {
		r.port.Close()
		r.port = nil
		r.addLog("串口已关闭")
	}
}

// IsSerialOpen 判断串口是否打开
func (r *RaySource160keV) IsSerialOpen() bool {
	return r.port != nil
}

// CalculateCheckSum 计算校验码（XOR异或校验）
func (r *RaySource160keV) CalculateCheckSum(data []byte) byte {
	var checkSum byte = 0
	for _, b := range data {
		checkSum ^= b
	}
	return checkSum
}

// PackFrame 封装数据帧
func (r *RaySource160keV) PackFrame(cmd CmdCode, arg string) []byte {
	frame := make([]byte, 0)

	// 1. 添加帧头
	frame = append(frame, STX)

	// 2. 添加命令码（ASCII格式）
	cmdStr := strconv.Itoa(int(cmd))
	frame = append(frame, []byte(cmdStr)...)

	// 3. 添加分隔符
	frame = append(frame, SEPARATOR)

	// 4. 添加参数（如果有）
	if arg != "" {
		frame = append(frame, []byte(arg)...)
	}

	// 5. 添加分隔符
	frame = append(frame, SEPARATOR)

	// 6. 计算并添加校验码（校验范围：命令码+分隔符+参数+分隔符）
	checkData := frame[1:] // 排除帧头
	checkSum := r.CalculateCheckSum(checkData)
	frame = append(frame, checkSum)

	// 7. 添加帧尾
	frame = append(frame, ETX)

	return frame
}

// SendCmd 发送命令
func (r *RaySource160keV) SendCmd(cmd CmdCode, arg string) error {
	if r.port == nil {
		err := fmt.Errorf("串口未打开")
		r.addLog(err.Error())
		return err
	}

	// 封装帧
	frame := r.PackFrame(cmd, arg)
	if len(frame) == 0 {
		err := fmt.Errorf("帧封装失败")
		r.addLog(err.Error())
		return err
	}

	// 发送数据
	n, err := r.port.Write(frame)
	if err != nil {
		r.addLog(fmt.Sprintf("发送失败: cmd=%d, err=%v", cmd, err))
		return err
	}

	if n != len(frame) {
		err := fmt.Errorf("发送长度不匹配: 预期=%d, 实际=%d", len(frame), n)
		r.addLog(err.Error())
		return err
	}

	r.addLog(fmt.Sprintf("发送成功: cmd=%d, arg=%s, frame=%s", cmd, arg, hex.EncodeToString(frame)))

	// 通讯间隔延迟
	time.Sleep(COMM_INTERVAL * time.Millisecond)
	return nil
}

// ReceiveResponse 接收响应
func (r *RaySource160keV) ReceiveResponse() ([]byte, error) {
	if r.port == nil {
		err := fmt.Errorf("串口未打开")
		r.addLog(err.Error())
		return nil, err
	}

	// 响应延迟（协议规定控制盒2ms响应，预留冗余）
	time.Sleep(RESPONSE_DELAY * time.Millisecond)

	buf := make([]byte, 1024)
	n, err := r.port.Read(buf)
	if err != nil {
		r.addLog(fmt.Sprintf("接收失败: %v", err))
		return nil, err
	}

	if n == 0 {
		err := fmt.Errorf("未接收到数据")
		r.addLog(err.Error())
		return nil, err
	}

	response := buf[:n]
	r.addLog(fmt.Sprintf("接收成功: 长度=%d, 数据=%s", n, hex.EncodeToString(response)))

	// 验证帧格式
	if !r.IsValidFrame(response) {
		r.addLog("帧格式验证失败")
		return nil, fmt.Errorf("帧格式无效")
	}

	return response, nil
}

// IsValidFrame 验证帧格式合法性
func (r *RaySource160keV) IsValidFrame(frame []byte) bool {
	// 最小长度检查：帧头1 + 命令码2 + 分隔符2 + 校验码1 + 帧尾1 = 7
	if len(frame) < 7 {
		return false
	}

	// 帧头验证
	if frame[0] != STX {
		return false
	}

	// 帧尾验证
	if frame[len(frame)-1] != ETX {
		return false
	}

	// 校验码验证
	receivedCheckSum := frame[len(frame)-2]
	checkData := frame[1 : len(frame)-2] // 排除帧头和校验码、帧尾
	calculatedCheckSum := r.CalculateCheckSum(checkData)
	if receivedCheckSum != calculatedCheckSum {
		return false
	}

	return true
}

// ParseResponse 解析响应
func (r *RaySource160keV) ParseResponse(cmd CmdCode, response []byte) string {
	// 帧格式：STX + RPT(1字节) + 分隔符 + 数据(n字节) + 分隔符 + 校验码 + ETX
	if len(response) < 7 {
		return "响应数据过短"
	}

	// 找到第一个分隔符位置（命令码结束）
	startIdx := 1
	for startIdx < len(response) && response[startIdx] != SEPARATOR {
		startIdx++
	}

	// 跳过第一个分隔符
	startIdx++

	// 找到第二个分隔符位置（数据结束）
	endIdx := startIdx
	for endIdx < len(response) && response[endIdx] != SEPARATOR {
		endIdx++
	}

	// 提取RPT数据
	if startIdx >= endIdx {
		return "无法提取数据"
	}

	rpt := string(response[startIdx:endIdx])

	// 根据命令码解析响应
	switch cmd {
	case READ_MODULE_STATUS: // 读取故障码（16位二进制）
		faultCode, err := strconv.ParseInt(rpt, 16, 32)
		if err != nil {
			return fmt.Sprintf("故障码解析失败: %s", rpt)
		}
		return r.parseFaultCode(int(faultCode))

	case READ_XRAY_STATUS: // 读取射线源状态
		status, err := strconv.Atoi(rpt)
		if err != nil {
			return fmt.Sprintf("状态解析失败: %s", rpt)
		}
		return r.parseXrayStatus(status)

	case READ_OIL_TEMP: // 读取油温（实际温度 = RPT - 30）
		temp, err := strconv.Atoi(rpt)
		if err != nil {
			return fmt.Sprintf("温度解析失败: %s", rpt)
		}
		actualTemp := (float64(temp) - 30) / 10.0
		return fmt.Sprintf("%.1f℃", actualTemp)

	case SET_KV, READ_KV_SET, READ_NOW_KV: // kV相关（单位0.1kV）
		kvInt, err := strconv.Atoi(rpt)
		if err != nil {
			return fmt.Sprintf("kV解析失败: %s", rpt)
		}
		kv := float64(kvInt) / 10.0
		return fmt.Sprintf("%.1fkV (原始值: %s)", kv, rpt)

	case SET_MA, READ_MA_SET, READ_NOW_MA: // mA相关（单位1uA）
		ma, err := strconv.Atoi(rpt)
		if err != nil {
			return fmt.Sprintf("mA解析失败: %s", rpt)
		}
		return fmt.Sprintf("%duA (原始值: %s)", ma, rpt)

	case READ_TOTAL_ARC_COUNT: // 总打火次数
		return fmt.Sprintf("设备总打火次数: %s次", rpt)

	case READ_LAST_ARC_COUNT: // 最近打火次数
		return fmt.Sprintf("最近打火次数: %s次", rpt)

	case READ_CODE_VERSION: // 软件版本号
		return fmt.Sprintf("软件版本号: %s", rpt)

	case READ_HARDWRE_VERSION: // 硬件版本号
		return fmt.Sprintf("硬件版本号: %s", rpt)

	case READ_SERIAL_NO: // 出厂编码
		return fmt.Sprintf("出厂编码: %s", rpt)

	default: // 默认返回原始RPT数据
		return rpt
	}
}

// parseFaultCode 解析故障码
func (r *RaySource160keV) parseFaultCode(faultCode int) string {
	if faultCode == 0 {
		return "正常（无故障）"
	}

	var faults []string
	for i := 0; i < 16; i++ {
		if faultCode&(1<<i) != 0 {
			if desc, ok := faultDescriptions[i]; ok {
				faults = append(faults, desc)
			}
		}
	}

	if len(faults) == 0 {
		return fmt.Sprintf("未知故障码: 0x%04X", faultCode)
	}

	return strings.Join(faults, "; ")
}

// parseXrayStatus 解析射线源状态
func (r *RaySource160keV) parseXrayStatus(status int) string {
	switch XRayStatus(status) {
	case XRAY_OFF:
		return "射线源关闭"
	case XRAY_ON:
		return "射线源打开"
	case XRAY_WARM:
		return "射线源预热中"
	default:
		return fmt.Sprintf("状态未知: %d", status)
	}
}

// addLog 添加通讯日志
func (r *RaySource160keV) addLog(log string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	r.logs = append(r.logs, fmt.Sprintf("[%s] %s", timestamp, log))
}

// GetLastLog 获取最后一条通讯日志
func (r *RaySource160keV) GetLastLog() string {
	if len(r.logs) == 0 {
		return "无日志"
	}
	return r.logs[len(r.logs)-1]
}

// GetAllLog 获取所有通讯日志
func (r *RaySource160keV) GetAllLog() []string {
	return r.logs
}

// SetKV 设置kV电压（单位：kV，内部自动转换为0.1kV的ASCII码）
func (r *RaySource160keV) SetKV(kV float64) error {
	arg := int(kV * 10)
	return r.SendCmd(SET_KV, strconv.Itoa(arg))
}

// SetMA 设置mA电流（单位：uA）
func (r *RaySource160keV) SetMA(mA int) error {
	return r.SendCmd(SET_MA, strconv.Itoa(mA))
}

// ReadKVSet 读取kV设定值
func (r *RaySource160keV) ReadKVSet() (string, error) {
	if err := r.SendCmd(READ_KV_SET, ""); err != nil {
		return "读取kV设定值失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取kV设定值失败", err
	}

	return r.ParseResponse(READ_KV_SET, response), nil
}

// ReadMASet 读取mA设定值
func (r *RaySource160keV) ReadMASet() (string, error) {
	if err := r.SendCmd(READ_MA_SET, ""); err != nil {
		return "读取mA设定值失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取mA设定值失败", err
	}

	return r.ParseResponse(READ_MA_SET, response), nil
}

// ReadFaultCode 读取故障码
func (r *RaySource160keV) ReadFaultCode() (string, error) {
	if err := r.SendCmd(READ_MODULE_STATUS, ""); err != nil {
		return "读取故障码失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取故障码失败", err
	}

	return r.ParseResponse(READ_MODULE_STATUS, response), nil
}

// ReadSoftwareVersion 读取软件版本号
func (r *RaySource160keV) ReadSoftwareVersion() (string, error) {
	if err := r.SendCmd(READ_CODE_VERSION, ""); err != nil {
		return "读取软件版本号失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取软件版本号失败", err
	}

	return r.ParseResponse(READ_CODE_VERSION, response), nil
}

// ReadHardwareVersion 读取硬件版本号
func (r *RaySource160keV) ReadHardwareVersion() (string, error) {
	if err := r.SendCmd(READ_HARDWRE_VERSION, ""); err != nil {
		return "读取硬件版本号失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取硬件版本号失败", err
	}

	return r.ParseResponse(READ_HARDWRE_VERSION, response), nil
}

// ReadSerialNo 读取出厂编码
func (r *RaySource160keV) ReadSerialNo() (string, error) {
	if err := r.SendCmd(READ_SERIAL_NO, ""); err != nil {
		return "读取出厂编码失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取出厂编码失败", err
	}

	return r.ParseResponse(READ_SERIAL_NO, response), nil
}

// FeedWatchDog 喂通讯看门狗
func (r *RaySource160keV) FeedWatchDog() error {
	return r.SendCmd(SERVICE_WATCHDOG, "")
}

// EnableWatchDog 使能看门狗（时间0~99秒，0=不使能）
func (r *RaySource160keV) EnableWatchDog(timeSec int) error {
	if timeSec < 0 || timeSec > 99 {
		err := fmt.Errorf("看门狗时间超出范围(0~99秒)")
		r.addLog(err.Error())
		return err
	}
	return r.SendCmd(WATCHDOG_ENA, strconv.Itoa(timeSec))
}

// ResetFaultFlag 清除故障标志位
func (r *RaySource160keV) ResetFaultFlag() error {
	return r.SendCmd(RESET_FAULT_FLAG, "")
}

// ReadCurrentKV 读取当前kV值
func (r *RaySource160keV) ReadCurrentKV() (string, error) {
	if err := r.SendCmd(READ_NOW_KV, ""); err != nil {
		return "读取当前kV值失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取当前kV值失败", err
	}

	return r.ParseResponse(READ_NOW_KV, response), nil
}

// ReadCurrentMA 读取当前mA值
func (r *RaySource160keV) ReadCurrentMA() (string, error) {
	if err := r.SendCmd(READ_NOW_MA, ""); err != nil {
		return "读取当前mA值失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取当前mA值失败", err
	}

	return r.ParseResponse(READ_NOW_MA, response), nil
}

// ReadOilTemp 读取油温值
func (r *RaySource160keV) ReadOilTemp() (string, error) {
	if err := r.SendCmd(READ_OIL_TEMP, ""); err != nil {
		return "读取油温失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取油温失败", err
	}

	return r.ParseResponse(READ_OIL_TEMP, response), nil
}

// ReadTotalArcCount 读取设备总打火次数
func (r *RaySource160keV) ReadTotalArcCount() (string, error) {
	if err := r.SendCmd(READ_TOTAL_ARC_COUNT, ""); err != nil {
		return "读取总打火次数失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取总打火次数失败", err
	}

	return r.ParseResponse(READ_TOTAL_ARC_COUNT, response), nil
}

// ReadLastArcCount 读取最近打火次数
func (r *RaySource160keV) ReadLastArcCount() (string, error) {
	if err := r.SendCmd(READ_LAST_ARC_COUNT, ""); err != nil {
		return "读取最近打火次数失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取最近打火次数失败", err
	}

	return r.ParseResponse(READ_LAST_ARC_COUNT, response), nil
}

// ReadXrayStatus 读取射线源状态
func (r *RaySource160keV) ReadXrayStatus() (string, error) {
	if err := r.SendCmd(READ_XRAY_STATUS, ""); err != nil {
		return "读取射线源状态失败", err
	}

	response, err := r.ReceiveResponse()
	if err != nil {
		return "读取射线源状态失败", err
	}

	return r.ParseResponse(READ_XRAY_STATUS, response), nil
}

// ControlXray 打开/关闭射线源（arg: "00"=关闭, "11"=打开）
func (r *RaySource160keV) ControlXray(arg string) error {
	if arg != "00" && arg != "11" {
		err := fmt.Errorf("参数无效，仅支持00(关闭)或11(打开)")
		r.addLog(err.Error())
		return err
	}
	return r.SendCmd(XRAY_ONOFF, arg)
}
