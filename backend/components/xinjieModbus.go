/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: 信捷PLC Modbus通信客户端，实现与PLC的Modbus TCP通信
 */
package components

//created by Tanglifa

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

type XinjieClient struct {
	handler *modbus.TCPClientHandler
	client  modbus.Client
	address string
	mu      sync.Mutex
}

// NewXinjieClient 创建信捷PLC客户端实例
func NewXinjieClient() *XinjieClient {
	return &XinjieClient{}
}

// OpenTCP 连接到指定 IP 和 端口 (默认 502)
func (x *XinjieClient) OpenTCP(address string, slaveID byte) error {
	x.mu.Lock()
	defer x.mu.Unlock()

	if x.handler != nil {
		x.handler.Close()
	}

	x.handler = modbus.NewTCPClientHandler(address)
	x.handler.Timeout = 2 * time.Second
	x.handler.SlaveId = slaveID

	err := x.handler.Connect()
	if err != nil {
		return err
	}

	x.client = modbus.NewClient(x.handler)
	x.address = address
	return nil
}

// Close 关闭连接
func (x *XinjieClient) Close() {
	if x.handler != nil {
		x.handler.Close()
	}
}

func (x *XinjieClient) ConvertOctalToDecimal(octalAddr int) (uint16, error) {
	// 将数字转为字符串进行合法性检查
	s := strconv.Itoa(octalAddr)

	// ParseUint 的第二个参数 8 代表按八进制解析
	val, err := strconv.ParseUint(s, 8, 16)
	if err != nil {
		return 0, fmt.Errorf("非法八进制地址: %d (不能包含8或9)", octalAddr)
	}

	return uint16(val), nil
}

// 读取整数
func (x *XinjieClient) ReadInt16(address uint16, count uint16) []int16 {
	registers, err := x.ReadRegisters(address, count)
	if err != nil {
		return nil
	} else {
		results := make([]int16, len(registers)/2)
		for i := 0; i < len(registers); i += 2 {
			// 大端序 (BigEndian)
			val := binary.BigEndian.Uint16(registers[i : i+2])
			results[i/2] = int16(val)
		}
		return results
	}
}

// ReadFloat32 读取32位浮点数
func (x *XinjieClient) ReadFloat32(address uint16, count uint16) []float32 {
	// 1个浮点数 = 2个寄存器 = 4字节
	// 所以读取的寄存器数量是 count * 2
	registers, err := x.ReadRegisters(address, count*2)
	if err != nil {
		return nil
	}

	results := make([]float32, count)
	for i := 0; i < int(count); i++ {
		startIndex := i * 4

		low := binary.BigEndian.Uint16(registers[startIndex : startIndex+2])
		high := binary.BigEndian.Uint16(registers[startIndex+2 : startIndex+4])
		data := uint32(high)<<16 | uint32(low)

		results[i] = math.Float32frombits(data)
	}

	return results
}

// 读Y线圈
func (x *XinjieClient) Read_Y_Coils(address uint16, count uint16) []bool {
	address, err := x.ConvertOctalToDecimal(int(address))
	registers, err := x.ReadCoils(address+0x6000, count)
	if err != nil {
		return nil
	}
	return registers
}

// 读X线圈
func (x *XinjieClient) Read_X_Coils(address uint16, count uint16) []bool {
	registers, err := x.ReadCoils(address+0x5000, count)
	if err != nil {
		return nil
	}
	return registers
}

// 读M线圈
func (x *XinjieClient) Read_M_Coils(address uint16, count uint16) []bool {
	registers, err := x.ReadCoils(address, count)
	if err != nil {
		return nil
	}
	return registers
}

// ReadRegisters 读取保持寄存器
func (x *XinjieClient) ReadRegisters(address uint16, quantity uint16) ([]byte, error) {
	if x.client == nil {
		return nil, fmt.Errorf("客户端未初始化")
	}
	x.mu.Lock()
	defer x.mu.Unlock() // 结束时解锁

	return x.client.ReadHoldingRegisters(address, quantity)
}

// 读线圈

func (x *XinjieClient) ReadCoils(address uint16, count uint16) ([]bool, error) {
	if x.client == nil {
		return nil, fmt.Errorf("客户端未初始化")
	}
	x.mu.Lock()
	defer x.mu.Unlock() // 结束时解锁

	results, err := x.client.ReadCoils(address, count)
	if err != nil {
		return nil, err
	}

	coils := make([]bool, count)
	for i := uint16(0); i < count; i++ {
		byteIndex := i / 8
		bitIndex := i % 8

		res := (results[byteIndex] >> bitIndex) & 0x01
		coils[i] = (res == 1)
	}

	return coils, nil
}

// WriteInt16 写单个数/多个整数
func (x *XinjieClient) WriteInt16(address uint16, values []int16) bool {
	count := uint16(len(values))
	payload := make([]byte, count*2)
	for i, v := range values {
		binary.BigEndian.PutUint16(payload[i*2:], uint16(v))
	}
	_, err := x.WriteRegisters(address, count, payload)
	return err == nil
}

// WriteFloat32 写32位浮点数
func (x *XinjieClient) WriteFloat32(address uint16, values []float32) bool {
	count := uint16(len(values))
	payload := make([]byte, count*4) // 1个浮点数占4字节
	for i, v := range values {
		bits := math.Float32bits(v)
		// 拆分为高16位和低16位
		high := uint16(bits >> 16)
		low := uint16(bits & 0xFFFF)

		binary.BigEndian.PutUint16(payload[i*4:], low)
		binary.BigEndian.PutUint16(payload[i*4+2:], high)
	}
	_, err := x.WriteRegisters(address, count*2, payload) // 寄存器数量需乘以2
	return err == nil
}

// Write_Y_Coils 写Y输出线圈 (偏移量 0x6000)
func (x *XinjieClient) Write_Y_Coils(address uint16, status []bool) bool {
	convertedAddress, err := x.ConvertOctalToDecimal(int(address))
	if err != nil {
		fmt.Printf("转换八进制地址失败: %v\n", err)
		return false
	}
	_, err = x.PackAndWriteCoils(convertedAddress+0x6000, status)
	return err == nil
}

// Write_M_Coils 写M辅助继电器
func (x *XinjieClient) Write_M_Coils(address uint16, status []bool) bool {
	convertedAddress, err := x.ConvertOctalToDecimal(int(address))
	if err != nil {
		fmt.Printf("转换八进制地址失败: %v\n", err)
		return false
	}
	_, err = x.PackAndWriteCoils(convertedAddress, status)
	return err == nil
}

// PackAndWriteCoils：将 bool 数组打包成字节并写入
func (x *XinjieClient) PackAndWriteCoils(address uint16, status []bool) ([]byte, error) {
	count := uint16(len(status))
	byteCount := (count + 7) / 8
	payload := make([]byte, byteCount)

	for i, s := range status {
		if s {
			payload[i/8] |= (1 << (uint(i) % 8))
		}
	}
	return x.WriteCoils(address, count, payload)
}

// WriteRegisters 通用写多个保持寄存器 (功能码 16)
func (x *XinjieClient) WriteRegisters(address uint16, quantity uint16, value []byte) ([]byte, error) {
	if x.client == nil {
		return nil, fmt.Errorf("客户端未初始化")
	}

	x.mu.Lock()
	defer x.mu.Unlock() // 结束时解锁

	return x.client.WriteMultipleRegisters(address, quantity, value)
}

// WriteCoils 通用写多个线圈 (功能码 15)
func (x *XinjieClient) WriteCoils(address uint16, quantity uint16, value []byte) ([]byte, error) {
	if x.client == nil {
		return nil, fmt.Errorf("客户端未初始化")
	}

	x.mu.Lock()
	defer x.mu.Unlock() // 结束时解锁

	return x.client.WriteMultipleCoils(address, quantity, value)
}
