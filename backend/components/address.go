/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: PLC地址映射定义，包含电机控制相关的寄存器地址
 */
package components

// PLC 地址映射结构
type AddressMap struct {
	XCW   uint16
	XCCW  uint16
	XSTOP uint16
	XABS  uint16

	YCW   uint16
	YCCW  uint16
	YSTOP uint16
	YABS  uint16

	ZCW   uint16
	ZCCW  uint16
	ZSTOP uint16
	ZABS  uint16

	RCW   uint16
	RCCW  uint16
	RSTOP uint16
	RABS  uint16

	ALLABS uint16
}

// ADDRESS 全局变量，用于访问地址
var ADDRESS = AddressMap{
	XCW:   3,
	XCCW:  4,
	XSTOP: 6,
	XABS:  5,

	YCW:   13,
	YCCW:  14,
	YSTOP: 16,
	YABS:  15,

	ZCW:   23,
	ZCCW:  24,
	ZSTOP: 26,
	ZABS:  25,

	RCW:   33,
	RCCW:  34,
	RSTOP: 36,
	RABS:  35,

	ALLABS: 0,
}
