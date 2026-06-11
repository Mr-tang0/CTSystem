/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: C SDK包装器，提供Go与C之间的接口转换和回调机制
 */
package CSDK

//此文件为SDK的包装器，尽量不要直接使用它，可以进行二次包装

/*
#cgo CFLAGS: -I${SRCDIR}/SDK
#cgo LDFLAGS: -L${SRCDIR}/SDK -lComApi

#include "NetCom.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// 声明将在 Go 侧实现的导出函数
extern int goOnLinkCallBack(char nEvent);
extern int goOnBreakCallBack(char nEvent);
extern int goOnImageCallBack(char nEvent);
extern int goOnHeartBeatCallBack(char nEvent);
extern int goOnReadyCallBack(char nEvent);
extern int goOnErrorCallBack(char nEvent);


// 这里的 C 回调函数被第三方 DLL 触发
// 注意：必须带上 WINAPI (即 __stdcall)，否则调用时堆栈会崩溃闪退！
static BOOL WINAPI CFuncLinkCallBack(char nEvent) {
    goOnLinkCallBack(nEvent);
    return 0;
}
static BOOL WINAPI CFuncBreakCallBack(char nEvent) {
    goOnBreakCallBack(nEvent);
    return 0;
}
static BOOL WINAPI CFuncImageCallBack(char nEvent) {
    goOnImageCallBack(nEvent);
    return 0;
}
static BOOL WINAPI CFuncHeartBeatCallBack(char nEvent) {
    goOnHeartBeatCallBack(nEvent);
    return 0;
}
static BOOL WINAPI CFuncReadyCallBack(char nEvent) {
    goOnReadyCallBack(nEvent);
    return 0;
}
static BOOL WINAPI CFuncErrorCallBack(char nEvent) {
    goOnErrorCallBack(nEvent);
    return 0;
}

// 桥接注册函数，将带 WINAPI 的 C 回调指针安全传入 DLL
static int RegisterLinkCallBackBridge() {
    return COM_RegisterEvCallBack(EVENT_LINKUP, CFuncLinkCallBack);
}
static int RegisterBreakCallBackBridge() {
    return COM_RegisterEvCallBack(EVENT_LINKDOWN, CFuncBreakCallBack);
}
static int RegisterImageCallBackBridge() {
    return COM_RegisterEvCallBack(EVENT_IMAGEVALID, CFuncImageCallBack);
}
static int RegisterHeartBeatCallbackBridge() {
    return COM_RegisterEvCallBack(EVENT_HEARTBEAT, CFuncHeartBeatCallBack);
}
static int RegisterReadyCallBackBridge() {
    return COM_RegisterEvCallBack(EVENT_READY, CFuncReadyCallBack);
}
static int RegisterErrorCallBackBridge() {
    return COM_RegisterEvCallBack(EVENT_TrigErr, CFuncErrorCallBack);
}
*/
import "C"
import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"sync"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type NetCom struct {
	ctx context.Context
}

var (
	imgBufferPool = sync.Pool{
		New: func() interface{} { return new(bytes.Buffer) },
	}
)

// 全局变量，用于 C 回调拿到 NetCom 实例发送 Wails 信号
var globalNetCom *NetCom

func NewNetCom(ctx context.Context) *NetCom {
	n := &NetCom{ctx: ctx}
	globalNetCom = n // 赋值给全局变量
	return n
}

// --- 导出给 C 调用的 Go 函数 ---

//export goOnLinkCallBack
func goOnLinkCallBack(nEvent C.char) C.int {
	val := int(nEvent)

	if globalNetCom != nil && globalNetCom.ctx != nil {
		if val == 1 {
			runtime.EventsEmit(globalNetCom.ctx, "ct_linked", map[string]bool{"ct_linked": true})
		}
		return 1
	} else {
		fmt.Println("[Go 捕获成功] 未初始化 NetCom 实例")
		return 0
	}
}

//export goOnBreakCallBack
func goOnBreakCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 断开连接: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		if val == 1 {
			runtime.EventsEmit(globalNetCom.ctx, "ct_linked", map[string]bool{"ct_linked": false})
		}
	}
	return 0
}

//export goOnImageCallBack
func goOnImageCallBack(nEvent C.char) C.int {
	fmt.Println("[Go 捕获成功] 收到图像有效事件")
	if globalNetCom == nil || globalNetCom.ctx == nil {
		return 0
	}
	var tImageMode C.TImageMode
	C.COM_GetImageMode(&tImageMode)
	row := int(tImageMode.usRow) // 高度
	col := int(tImageMode.usCol) // 宽度

	if row <= 0 || col <= 0 {
		return 0
	}
	bufSize := 2 * row * col
	pPicBuff := C.malloc(C.size_t(bufSize))
	if pPicBuff == nil {
		return 0
	}
	defer C.free(pPicBuff) // 确保在转换逻辑执行完毕后释放 C 内存，防止内存泄漏

	// 从底层的共享内存或硬通道中抓取图像
	C.COM_GetImage((*C.char)(pPicBuff))

	// 将 C 内存指针映射为 Go 的无拷贝 []byte 切片
	raw16Data := (*[1 << 30]byte)(pPicBuff)[:bufSize:bufSize]

	// 创建标准的 8 位灰度图像容器
	grayImg := image.NewGray(image.Rect(0, 0, col, row))

	// 16位 RAW 转 8位灰度：使用固定参考值转换
	const maxRef uint32 = 255 // 16位最大值作为固定参考
	for i := 0; i < col*row; i++ {
		// Little-endian: 低字节在前，高字节在后
		val := uint32(raw16Data[i*2]) | (uint32(raw16Data[i*2+1]) << 8)
		// 直接线性映射到 0-255
		normalized := val * 255 / maxRef
		grayImg.Pix[i] = uint8(normalized)
	}

	// 从对象池获取 Buffer，对灰度图执行高效 JPEG 编码
	buf := imgBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer imgBufferPool.Put(buf)

	// 质量设为 60：既保证了图像的高还原度，又大幅压减了 Base64 的传输体积，提升帧率
	if err := jpeg.Encode(buf, grayImg, &jpeg.Options{Quality: 60}); err != nil {
		fmt.Println("[Go 图像处理] JPEG 编码失败:", err)
		return 0
	}

	// 转换为 Base64（添加 JPEG 前缀，以便浏览器正确识别）
	encodedStr := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	data := map[string]string{
		"image":  encodedStr,
		"width":  fmt.Sprintf("%d", col),
		"height": fmt.Sprintf("%d", row),
	}

	// 将完整的 Base64 图像异步广播至前端监听器
	go runtime.EventsEmit(globalNetCom.ctx, "ct_image", data)

	return 0
}

//export goOnHeartBeatCallBack
func goOnHeartBeatCallBack(nEvent C.char) C.int {
	// val := int(nEvent)
	// fmt.Printf("[Go 捕获成功] 收到心跳事件: %d\n", val)

	// 获取序列号
	acSnTmp := make([]C.char, 32)
	C.COM_GetFPsn(&acSnTmp[0])
	sn := C.GoString(&acSnTmp[0])

	// 获取当前状态
	cFpCurStat := C.COM_GetFPCurStatus()
	statusText := ""
	switch cFpCurStat {
	case 0:
		statusText = "AED1"
	case 1:
		statusText = "IDLE"
	case 2:
		statusText = "HST"
	case 3:
		statusText = "AED2"
	case 4:
		statusText = "RECOVER"
	default:
		statusText = "ERR"
	}
	// fmt.Printf("[心跳数据] 状态: %s (%d)\n", statusText, cFpCurStat)

	// 获取完整状态信息
	wifiSignal := 0
	temp := 0.0
	hum := 0.0
	var tFPStat C.TFPStat
	result := C.COM_GetFPStatus(&tFPStat)
	if result == 1 {
		// WiFi信号强度
		wifiSignal = int(tFPStat.tWifiStatus.ucSignal_level)
		// fmt.Printf("[心跳数据] WiFi信号强度: %d\n", wifiSignal)

		// 温湿度 (需要除以10)
		temp = float64(tFPStat.tFpTempHum.Temp) / 10.0
		hum = float64(tFPStat.tFpTempHum.Hum) / 10.0
		// fmt.Printf("[心跳数据] 温度: %.1f°C, 湿度: %.1f%%\n", temp, hum)

	} else {
		fmt.Printf("[心跳数据] 获取状态信息失败\n")
	}

	// var pxwin C.UINT32
	// var prepeat C.UINT16
	// var pbinMode C.CHAR
	// var psync C.CHAR
	// C.COM_GetDynamicPara(&pxwin, &prepeat, &pbinMode, &psync)
	// fmt.Printf("[心跳数据] XWin: %d, 重复次数: %d, 采集模式: %d, 同步模式: %d\n", xwin, prepeat, pbinMode, psync)

	data := map[string]interface{}{
		"sn":         sn,
		"mode":       statusText,
		"wifi":       wifiSignal,
		"tempreture": temp,
		"humidity":   hum,
		"battery":    tFPStat.tBatInfo1.Remain,
	}

	// fmt.Println("[Go 捕获成功] 获取心跳数据:", data)

	runtime.EventsEmit(globalNetCom.ctx, "ct_heartbeat", data)

	return 0
}

//export goOnReadyCallBack
func goOnReadyCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 获取设备就绪状态: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_ready", val)
	}
	return 0
}

//export goOnErrorCallBack
func goOnErrorCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_error", val)
	}
	return 0
}

// --- 设备控制函数，通过 go 调用 C 函数 ---

// 一键注册所有事件回调
func (n *NetCom) RegisterCallback() bool {
	C.RegisterLinkCallBackBridge()
	C.RegisterBreakCallBackBridge()
	C.RegisterHeartBeatCallbackBridge()
	C.RegisterReadyCallBackBridge()
	C.RegisterErrorCallBackBridge()
	C.RegisterImageCallBackBridge()
	return true
}

// 初始化SDK
func (n *NetCom) COM_Init() bool {
	return C.COM_Init() == 1
}

// 反初始化SDK
func (n *NetCom) COM_Uninit() bool {
	return C.COM_Uninit() == 1
}

// 设置校准模式
func (n *NetCom) COM_SetCalibMode(mode int) bool {
	return C.COM_SetCalibMode(C.CHAR(mode)) == 1
}

// 获取设备序列号
func (n *NetCom) COM_GetFPsn() string {
	buf := make([]byte, 32)
	C.COM_GetFPsn((*C.CHAR)(unsafe.Pointer(&buf[0])))
	return string(buf)
}

// 列出所有设备
func (n *NetCom) COM_List(ptComFpList *TComFpList) bool {
	cStructPtr := (*C.TComFpList)(unsafe.Pointer(ptComFpList))
	return C.COM_List(cStructPtr) == 1
}

// 连接探测器
// 如果 cSn 为空（0），则自动连接（传入 NULL）
func (n *NetCom) COM_Open(cSn Char) bool {
	if cSn == 0 {
		// 传入 NULL，让SDK自动连接
		return C.COM_Open(nil) == 1
	}
	return C.COM_Open((*C.CHAR)(unsafe.Pointer(&cSn))) == 1
}

// 断开连接
func (n *NetCom) COM_Close() bool {
	return C.COM_Close() == 1
}

// 启动网络监听
func (n *NetCom) COM_StartNet() bool {
	return C.COM_StartNet() == 1
}

// COM_StopNet 停止采集并进入空闲状态
func (m *NetCom) COM_StopNet() bool {
	return C.COM_StopNet() == 1
}

//设置曝光时间

func (n *NetCom) COM_SetExposeTime(exposeTime int) bool {
	return C.COM_SetXwin(C.UINT32(exposeTime)) == 1
}

// 获取曝光时间
func (n *NetCom) COM_GetExposeTime() int {
	var xwin C.UINT32
	C.COM_GetXwin(&xwin)
	return int(xwin)
}

// 设置Binning
func (n *NetCom) COM_SetBinning(binning string) bool {
	switch binning {
	case "1×1":
		return C.COM_SetBinningMode(C.BINNING_1x1) == 1
	case "2×2":
		return C.COM_SetBinningMode(C.BINNING_2x2) == 1
	case "3×3":
		return C.COM_SetBinningMode(C.BINNING_3x3) == 1
	case "4×4":
		return C.COM_SetBinningMode(C.BINNING_4x4) == 1
	case "6×6":
		return C.COM_SetBinningMode(C.BINNING_6x6) == 1
	case "8×8":
		return C.COM_SetBinningMode(C.BINNING_8x8) == 1
	default:
		fmt.Printf("设置binning失败: %s\n", binning)
		return false
	}

}

// 获取Binning
func (n *NetCom) COM_GetBinning() string {
	var pbinMode C.CHAR
	C.COM_GetBinningMode(&pbinMode)

	switch pbinMode {
	case C.BINNING_1x1:
		return "1×1"
	case C.BINNING_2x2:
		return "2×2"
	case C.BINNING_3x3:
		return "3×3"
	case C.BINNING_4x4:
		return "4×4"
	case C.BINNING_6x6:
		return "6×6"
	case C.BINNING_8x8:
		return "8×8"
	default:
		return "error"
	}
}

// 连续触发流程COM_DstAcq->COM_Dprep->COM_Dacq->COM_StopNet
// 设置采集模式

// 设置DST模式
func (n *NetCom) COM_Dst() bool {
	return C.COM_Dst() == 1
}

func (n *NetCom) COM_Dprep() bool {
	return C.COM_Dprep() == 1
}

func (n *NetCom) COM_Dacq() bool {
	return C.COM_Dacq() == 1
}

// 单帧采集流程：COM_HstAcq->COM_ExposeReq->COM_StopNet
func (n *NetCom) COM_HstAcq() bool {
	return C.COM_HstAcq() == 1
}

func (n *NetCom) COM_ExposeReq() bool {
	return C.COM_ExposeReq() == 1
}
