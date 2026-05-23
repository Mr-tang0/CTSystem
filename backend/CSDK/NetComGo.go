package CSDK

//此文件为SDK的包装器，尽量不要直接使用它，可以进行二次包装

/*
#cgo CFLAGS: -I${SRCDIR}/SDK
#cgo LDFLAGS: -L${SRCDIR}/SDK -lComApi

#include "NetCom.h"
#include <stdlib.h>
#include <string.h>

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

func NewNetCom() *NetCom {
	n := &NetCom{}
	globalNetCom = n // 赋值给全局变量
	return n
}

func (n *NetCom) SetContext(ctx context.Context) {
	n.ctx = ctx
}

// --- 导出给 C 调用的 Go 函数 ---

//export goOnLinkCallBack
func goOnLinkCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到链路连接事件: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_linked", val)
	}
	return 0
}

//export goOnBreakCallBack
func goOnBreakCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到链路断开事件: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_breaked", val)
	}
	return 0
}

//export goOnImageCallBack
func goOnImageCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到图像就绪事件: %d\n", val)
	if globalNetCom == nil || globalNetCom.ctx == nil {
		return 0
	}
	var tImageMode C.TImageMode
	C.COM_GetImageMode(&tImageMode)
	row := int(tImageMode.usRow) // 高度
	col := int(tImageMode.usCol) // 宽度

	if row <= 0 || col <= 0 {
		fmt.Println("[Go 图像处理] 异常的长宽数据: Row=", row, "Col=", col)
		return 0
	}
	bufSize := 2 * row * col
	pPicBuff := C.malloc(C.size_t(bufSize))
	if pPicBuff == nil {
		fmt.Println("[Go 图像处理] 内存分配失败")
		return 0
	}
	defer C.free(pPicBuff) // 确保在转换逻辑执行完毕后释放 C 内存，防止内存泄漏

	// 从底层的共享内存或硬通道中抓取图像
	// 💡 修复点：将参数类型强转为 (*C.char) 以适配 C 接口中的 CHAR* 类型定义
	C.COM_GetImage((*C.char)(pPicBuff))

	// 将 C 内存指针映射为 Go 的无拷贝 []byte 切片
	raw16Data := (*[1 << 30]byte)(pPicBuff)[:bufSize:bufSize]

	// 创建标准的 8 位灰度图像容器
	grayImg := image.NewGray(image.Rect(0, 0, col, row))

	// 进行 Raw16 转 Gray8 的归一化/降采样处理
	for i := 0; i < col*row; i++ {
		grayImg.Pix[i] = raw16Data[i*2+1]
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

	// 转换为 Base64
	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 将完整的 Base64 图像异步广播至前端监听器
	go runtime.EventsEmit(globalNetCom.ctx, "ct_image", encodedStr)

	return 0
}

//export goOnHeartBeatCallBack
func goOnHeartBeatCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到心跳事件: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_heartbeat", val)
	}
	return 0
}

//export goOnReadyCallBack
func goOnReadyCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到就绪就绪事件: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_ready", val)
	}
	return 0
}

//export goOnErrorCallBack
func goOnErrorCallBack(nEvent C.char) C.int {
	val := int(nEvent)
	fmt.Printf("[Go 捕获成功] 收到错误异常事件: %d\n", val)
	if globalNetCom != nil && globalNetCom.ctx != nil {
		go runtime.EventsEmit(globalNetCom.ctx, "ct_error", val)
	}
	return 0
}

// --- 设备控制函数，通过 go 调用 C 函数 ---
// 初始化SDK
func (n *NetCom) COM_Init() bool {
	return C.COM_Init() == 1
}

// 反初始化SDK
func (n *NetCom) COM_Uninit() bool {
	return C.COM_Uninit() == 1
}

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

// 列出所有设备
func (n *NetCom) COM_List(ptComFpList *TComFpList) bool {
	cStructPtr := (*C.TComFpList)(unsafe.Pointer(ptComFpList))
	return C.COM_List(cStructPtr) == 1
}

// 连接探测器
func (n *NetCom) COM_Open(cSn Char) bool {
	return C.COM_Open((*C.CHAR)(unsafe.Pointer(&cSn))) == 1
}

// 断开连接
func (n *NetCom) COM_Close() bool {
	return C.COM_Close() == 1
}

// 软件触发一张图片
func (n *NetCom) COM_SoftTrigger() bool {
	return C.COM_ExposeReq() == 1
}

// AED 触发一张图片
func (n *NetCom) COM_AedTrigger() bool {
	return C.COM_AedTrigger() == 1
}

// 设置高分辨率模式
func (m *NetCom) COM_SetHstMode() bool {
	return C.COM_HstAcq() == 1
}

// SetAedMode 进入自动曝光检测模式？ (AED)
func (m *NetCom) COM_SetAedMode() bool {
	return C.COM_AedAcq() == 1
}

// COM_StopNet 停止采集并进入空闲状态
func (m *NetCom) COM_StopNet() bool {
	return C.COM_StopNet() == 1
}

// COM_StartNet 启动采集
func (m *NetCom) COM_StartNet() bool {
	return C.COM_StartNet() == 1
}
