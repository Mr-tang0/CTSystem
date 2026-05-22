package components

import "math"

type MotorDevice struct {
	ID  int
	plc *XinjieClient
}

func (m *MotorDevice) Connect(ip string) error {
	return m.plc.OpenTCP(ip, 1)
}

func (m *MotorDevice) Disconnect() {
	m.plc.Close()
}

func (m *MotorDevice) MotorJogMove(axis string, speed float32) {
	// 先判断轴和方向
	var cwAddr, ccwAddr uint16
	switch axis {
	case "X":
		cwAddr = ADDRESS.XCW
		ccwAddr = ADDRESS.XCCW
	case "Y":
		cwAddr = ADDRESS.YCW
		ccwAddr = ADDRESS.YCCW
	case "Z":
		cwAddr = ADDRESS.ZCW
		ccwAddr = ADDRESS.ZCCW
	case "R":
		cwAddr = ADDRESS.RCW
		ccwAddr = ADDRESS.RCCW
	default:
		return
	}

	// 统一写PLC
	if speed > 0 {
		m.plc.Write_M_Coils(cwAddr, []bool{true})
		m.plc.Write_M_Coils(cwAddr, []bool{false})
	} else {
		m.plc.Write_M_Coils(ccwAddr, []bool{true})
		m.plc.Write_M_Coils(ccwAddr, []bool{false})
	}

	//速度？
	//PLC未给出速度所对应地址

}

func (m *MotorDevice) MotorStop(axis string) {
	var stopAddr uint16
	switch axis {
	case "X":
		stopAddr = ADDRESS.XSTOP
	case "Y":
		stopAddr = ADDRESS.YSTOP
	case "Z":
		stopAddr = ADDRESS.ZSTOP
	case "R":
		stopAddr = ADDRESS.RSTOP
	default:
		return
	}

	m.plc.Write_M_Coils(stopAddr, []bool{true})
	m.plc.Write_M_Coils(stopAddr, []bool{false})
}

func (m *MotorDevice) MotorAbsMove(axis string, speed float32, abs_length float32) {
	var moveAddr uint16
	switch axis {
	case "X":
		moveAddr = ADDRESS.XABS
	case "Y":
		moveAddr = ADDRESS.YABS
	case "Z":
		moveAddr = ADDRESS.ZABS
	case "R":
		moveAddr = ADDRESS.RABS
	default:
		return
	}

	speed_puls := int16(math.Abs(float64(speed)) * 1600 / 5)
	length_puls := int16(math.Abs(float64(abs_length)) * 1600 / 5)

	if speed > 0 {
		m.plc.WriteInt16(ADDRESS.ALLABS, []int16{speed_puls, 0, length_puls}) //神奇参数？？？正为0，负为-1？？？
	} else {
		m.plc.WriteInt16(ADDRESS.ALLABS, []int16{speed_puls, -1, length_puls})
	}

	m.plc.Write_M_Coils(moveAddr, []bool{true})
	m.plc.Write_M_Coils(moveAddr, []bool{false})

}

func (m *MotorDevice) MotorRelMove(axis string, speed float32, rel_length float32) {
	// 暂未实现
}
