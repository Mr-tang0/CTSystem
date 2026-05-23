package CSDK

/*
#cgo CFLAGS: -I../../Manager/SDK/INCLUDE
#cgo LDFLAGS: -L../../Manager/SDK/LIB -lNetCom

#include "NetCom.h"
*/
import "C"
import "unsafe"

// NetCom 封装了CSDK的所有网络通信功能
type NetCom struct{}

// 基本初始化函数

func (n *NetCom) COM_Init() bool {
	return bool(C.COM_Init())
}

func (n *NetCom) COM_Uninit() bool {
	return bool(C.COM_Uninit())
}

func (n *NetCom) COM_SetCfgFilePath(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetCfgFilePath(cPath))
}

// 设备列表管理
func (n *NetCom) COM_List(ptComFpList *TComFpList) bool {
	cList := (*C.TComFpList)(unsafe.Pointer(ptComFpList))
	return bool(C.COM_List(cList))
}

func (n *NetCom) COM_ListAdd(psn string) bool {
	cSn := C.CString(psn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_ListAdd(cSn))
}

func (n *NetCom) COM_ListDel(psn string) bool {
	cSn := C.CString(psn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_ListDel(cSn))
}

func (n *NetCom) COM_ListClr() bool {
	return bool(C.COM_ListClr())
}

// 设备连接控制
func (n *NetCom) COM_Open(psn string) bool {
	cSn := C.CString(psn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_Open(cSn))
}

func (n *NetCom) COM_Close() bool {
	return bool(C.COM_Close())
}

func (n *NetCom) COM_StopNet() bool {
	return bool(C.COM_StopNet())
}

func (n *NetCom) COM_StartNet() bool {
	return bool(C.COM_StartNet())
}

// 事件回调注册

func (n *NetCom) COM_RegisterEvCallBack(nEvent byte, funcallback unsafe.Pointer) bool {
	return bool(C.COM_RegisterEvCallBack(C.CHAR(nEvent), (C.FP_EVENT_CALLBACK)(funcallback)))
}

// 校准模式

func (n *NetCom) COM_SetPreCalibMode(nCalMode byte) bool {
	return bool(C.COM_SetPreCalibMode(C.CHAR(nCalMode)))
}

func (n *NetCom) COM_GetPreCalibMode() byte {
	return byte(C.COM_GetPreCalibMode())
}

func (n *NetCom) COM_SetCalibMode(nCalMode byte) bool {
	return bool(C.COM_SetCalibMode(C.CHAR(nCalMode)))
}

func (n *NetCom) COM_GetCalibMode() byte {
	return byte(C.COM_GetCalibMode())
}

// 采集控制

func (n *NetCom) COM_HstAcq() bool {
	return bool(C.COM_HstAcq())
}

func (n *NetCom) COM_AedAcq() bool {
	return bool(C.COM_AedAcq())
}

func (n *NetCom) COM_Trigger() bool {
	return bool(C.COM_Trigger())
}

func (n *NetCom) COM_Trigger2() bool {
	return bool(C.COM_Trigger2())
}

func (n *NetCom) COM_Prep() bool {
	return bool(C.COM_Prep())
}

func (n *NetCom) COM_Acq() bool {
	return bool(C.COM_Acq())
}

func (n *NetCom) COM_PrepAcq() bool {
	return bool(C.COM_PrepAcq())
}

func (n *NetCom) COM_SetAcq() bool {
	return bool(C.COM_SetAcq())
}

func (n *NetCom) COM_ComAcq() bool {
	return bool(C.COM_ComAcq())
}

func (n *NetCom) COM_ExposeReq() bool {
	return bool(C.COM_ExposeReq())
}

func (n *NetCom) COM_AedTrigger() bool {
	return bool(C.COM_AedTrigger())
}

func (n *NetCom) COM_AedPrep() bool {
	return bool(C.COM_AedPrep())
}

func (n *NetCom) COM_Aed2Acq() bool {
	return bool(C.COM_Aed2Acq())
}

func (n *NetCom) COM_Stop() bool {
	return bool(C.COM_Stop())
}

func (n *NetCom) COM_Dst() bool {
	return bool(C.COM_Dst())
}

func (n *NetCom) COM_Dacq() bool {
	return bool(C.COM_Dacq())
}

func (n *NetCom) COM_Dacqaed() bool {
	return bool(C.COM_Dacqaed())
}

func (n *NetCom) COM_Cbct() bool {
	return bool(C.COM_Cbct())
}

func (n *NetCom) COM_Cbct2() bool {
	return bool(C.COM_Cbct2())
}

func (n *NetCom) COM_Dexit() bool {
	return bool(C.COM_Dexit())
}

func (n *NetCom) COM_Dprep() bool {
	return bool(C.COM_Dprep())
}

func (n *NetCom) COM_Cprep() bool {
	return bool(C.COM_Cprep())
}

func (n *NetCom) COM_Exprep() bool {
	return bool(C.COM_Exprep())
}

func (n *NetCom) COM_SetConfigId(ucConfigId byte) bool {
	return bool(C.COM_SetConfigId(C.UCHAR(ucConfigId)))
}

func (n *NetCom) COM_GetConfigId(ucConfigId *byte) bool {
	return bool(C.COM_GetConfigId((*C.UCHAR)(unsafe.Pointer(ucConfigId))))
}

func (n *NetCom) COM_SetModeId(ucModeId byte) bool {
	return bool(C.COM_SetModeId(C.UCHAR(ucModeId)))
}

func (n *NetCom) COM_GetModeId(ucModeId *byte) bool {
	return bool(C.COM_GetModeId((*C.UCHAR)(unsafe.Pointer(ucModeId))))
}

func (n *NetCom) COM_LoadFullCfg(ucModeId byte) bool {
	return bool(C.COM_LoadFullCfg(C.UCHAR(ucModeId)))
}

func (n *NetCom) COM_SaveFullCfg(ucModeId byte) bool {
	return bool(C.COM_SaveFullCfg(C.UCHAR(ucModeId)))
}

func (n *NetCom) COM_SetMetaData(tMetaData *TMetaData) bool {
	cMeta := (*C.TMetaData)(unsafe.Pointer(tMetaData))
	return bool(C.COM_SetMetaData(*cMeta))
}

func (n *NetCom) COM_GetMetaData(ptMetaData *TMetaData) bool {
	cMeta := (*C.TMetaData)(unsafe.Pointer(ptMetaData))
	return bool(C.COM_GetMetaData(cMeta))
}

func (n *NetCom) COM_GetPreImg() bool {
	return bool(C.COM_GetPreImg())
}

func (n *NetCom) COM_AedAcqOffLine() bool {
	return bool(C.COM_AedAcqOffLine())
}

func (n *NetCom) COM_AcqOffLineImage() bool {
	return bool(C.COM_AcqOffLineImage())
}

func (n *NetCom) COM_GetNumOffLineImg() uint32 {
	return uint32(C.COM_GetNumOffLineImg())
}

func (n *NetCom) COM_ImgSaveStart() bool {
	return bool(C.COM_ImgSaveStart())
}

func (n *NetCom) COM_ImgSaveExit() bool {
	return bool(C.COM_ImgSaveExit())
}

func (n *NetCom) COM_ImgSavedRead() bool {
	return bool(C.COM_ImgSavedRead())
}

// 图像获取相关

func (n *NetCom) COM_GetImageMode(ptImageMode *TImageMode) bool {
	cMode := (*C.TImageMode)(unsafe.Pointer(ptImageMode))
	return bool(C.COM_GetImageMode(cMode))
}

func (n *NetCom) COM_GetImageModeV(ptImageMode *TImageMode) bool {
	cMode := (*C.TImageMode)(unsafe.Pointer(ptImageMode))
	return bool(C.COM_GetImageModeV(cMode))
}

func (n *NetCom) COM_GetImageShiftMode(ptImageShiftMode *TImageShiftMode) bool {
	cShift := (*C.TImageShiftMode)(unsafe.Pointer(ptImageShiftMode))
	return bool(C.COM_GetImageShiftMode(cShift))
}

func (n *NetCom) COM_GetImageName(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return bool(C.COM_GetImageName(cName))
}

func (n *NetCom) COM_ClrImageID() bool {
	return bool(C.COM_ClrImageID())
}

func (n *NetCom) COM_GetImageID(pimgID *uint32) bool {
	return bool(C.COM_GetImageID((*C.UINT32)(unsafe.Pointer(pimgID))))
}

func (n *NetCom) COM_GetImage(pImageBuff []byte) bool {
	return bool(C.COM_GetImage((*C.CHAR)(unsafe.Pointer(&pImageBuff[0]))))
}

func (n *NetCom) COM_GetImageV(pImageBuff []byte) bool {
	return bool(C.COM_GetImageV((*C.CHAR)(unsafe.Pointer(&pImageBuff[0]))))
}

func (n *NetCom) COM_ResetFP() bool {
	return bool(C.COM_ResetFP())
}

func (n *NetCom) COM_FpTurnOff() bool {
	return bool(C.COM_FpTurnOff())
}

func (n *NetCom) COM_GetErrNo() int32 {
	return int32(C.COM_GetErrNo())
}

func (n *NetCom) COM_DhcpActivate(enableflag byte) bool {
	return bool(C.COM_DhcpActivate(C.CHAR(enableflag)))
}

func (n *NetCom) COM_DhcpSetCfg(tDhcpCfg *TDhcpCfg) bool {
	cCfg := (*C.TDhcpCfg)(unsafe.Pointer(tDhcpCfg))
	return bool(C.COM_DhcpSetCfg(cCfg))
}

func (n *NetCom) COM_DhcpGetCfg(tDhcpCfg *TDhcpCfg) bool {
	cCfg := (*C.TDhcpCfg)(unsafe.Pointer(tDhcpCfg))
	return bool(C.COM_DhcpGetCfg(cCfg))
}

func (n *NetCom) COM_ImgOverlayEnable(on byte) bool {
	return bool(C.COM_ImgOverlayEnable(C.UCHAR(on)))
}

// 配置设置

func (n *NetCom) COM_SetFPConf(ptFpUserCfg *TFPUserCfg) bool {
	cCfg := (*C.TFPUserCfg)(unsafe.Pointer(ptFpUserCfg))
	return bool(C.COM_SetFPConf(cCfg))
}

func (n *NetCom) COM_GetFPConf(ptFpUserCfg *TFPUserCfg) bool {
	cCfg := (*C.TFPUserCfg)(unsafe.Pointer(ptFpUserCfg))
	return bool(C.COM_GetFPConf(cCfg))
}

func (n *NetCom) COM_SetRBConf(ptRbConf *TRBConf) bool {
	cCfg := (*C.TRBConf)(unsafe.Pointer(ptRbConf))
	return bool(C.COM_SetRBConf(cCfg))
}

func (n *NetCom) COM_GetRBConf(ptRbConf *TRBConf) bool {
	cCfg := (*C.TRBConf)(unsafe.Pointer(ptRbConf))
	return bool(C.COM_GetRBConf(cCfg))
}

func (n *NetCom) COM_SetWifiMode(ApMode byte) bool {
	return bool(C.COM_SetWifiMode(C.CHAR(ApMode)))
}

func (n *NetCom) COM_GetWifiMode() byte {
	return byte(C.COM_GetWifiMode())
}

func (n *NetCom) COM_SetWifiConf(ptWifiConf *TWifiConf) bool {
	cCfg := (*C.TWifiConf)(unsafe.Pointer(ptWifiConf))
	return bool(C.COM_SetWifiConf(cCfg))
}

func (n *NetCom) COM_GetWifiConf(ptWifiConf *TWifiConf) bool {
	cCfg := (*C.TWifiConf)(unsafe.Pointer(ptWifiConf))
	return bool(C.COM_GetWifiConf(cCfg))
}

func (n *NetCom) COM_SetWifiConfig(ptWifiConfAp *TWifiConf, ptWifiConfSta *TWifiConf) bool {
	cCfgAp := (*C.TWifiConf)(unsafe.Pointer(ptWifiConfAp))
	cCfgSta := (*C.TWifiConf)(unsafe.Pointer(ptWifiConfSta))
	return bool(C.COM_SetWifiConfig(cCfgAp, cCfgSta))
}

func (n *NetCom) COM_GetWifiConfig(ptWifiConfAp *TWifiConf, ptWifiConfSta *TWifiConf) bool {
	cCfgAp := (*C.TWifiConf)(unsafe.Pointer(ptWifiConfAp))
	cCfgSta := (*C.TWifiConf)(unsafe.Pointer(ptWifiConfSta))
	return bool(C.COM_GetWifiConfig(cCfgAp, cCfgSta))
}

func (n *NetCom) COM_SetWifiCountry(pCountry string) bool {
	cCountry := C.CString(pCountry)
	defer C.free(unsafe.Pointer(cCountry))
	return bool(C.COM_SetWifiCountry(cCountry))
}

func (n *NetCom) COM_GetWifiCountry(pCountry string) bool {
	cCountry := C.CString(pCountry)
	defer C.free(unsafe.Pointer(cCountry))
	return bool(C.COM_GetWifiCountry(cCountry))
}

func (n *NetCom) COM_SetApEssid(pessid string) bool {
	cEssid := C.CString(pessid)
	defer C.free(unsafe.Pointer(cEssid))
	return bool(C.COM_SetApEssid(cEssid))
}

func (n *NetCom) COM_GetApEssid(pessid string) bool {
	cEssid := C.CString(pessid)
	defer C.free(unsafe.Pointer(cEssid))
	return bool(C.COM_GetApEssid(cEssid))
}

func (n *NetCom) COM_DefFPConf() bool {
	return bool(C.COM_DefFPConf())
}

func (n *NetCom) COM_DefRBConf() bool {
	return bool(C.COM_DefRBConf())
}

func (n *NetCom) COM_DefFPTpl() bool {
	return bool(C.COM_DefFPTpl())
}

func (n *NetCom) COM_SetXwin(xwin uint32) bool {
	return bool(C.COM_SetXwin(C.UINT32(xwin)))
}

func (n *NetCom) COM_GetXwin(xwin *uint32) bool {
	return bool(C.COM_GetXwin((*C.UINT32)(unsafe.Pointer(xwin))))
}

func (n *NetCom) COM_SetXwin_us(xwin_us uint32) bool {
	return bool(C.COM_SetXwin_us(C.UINT32(xwin_us)))
}

func (n *NetCom) COM_GetXwin_us(pxwin_us *uint32) bool {
	return bool(C.COM_GetXwin_us((*C.UINT32)(unsafe.Pointer(pxwin_us))))
}

func (n *NetCom) COM_SetTrailTime(msec uint16) bool {
	return bool(C.COM_SetTrailTime(C.USHORT(msec)))
}

func (n *NetCom) COM_SetDynamicPara(xwin uint32, repeat uint16, binMode byte, sync byte) bool {
	return bool(C.COM_SetDynamicPara(C.UINT32(xwin), C.UINT16(repeat), C.CHAR(binMode), C.CHAR(sync)))
}

func (n *NetCom) COM_GetDynamicPara(pxwin *uint32, prepeat *uint16, pbinMode *byte, psync *byte) bool {
	return bool(C.COM_GetDynamicPara((*C.UINT32)(unsafe.Pointer(pxwin)), (*C.UINT16)(unsafe.Pointer(prepeat)), (*C.CHAR)(unsafe.Pointer(pbinMode)), (*C.CHAR)(unsafe.Pointer(psync))))
}

func (n *NetCom) COM_SetBinningMode(cbinningMode byte) bool {
	return bool(C.COM_SetBinningMode(C.CHAR(cbinningMode)))
}

func (n *NetCom) COM_GetBinningMode(cbinningMode *byte) bool {
	return bool(C.COM_GetBinningMode((*C.CHAR)(unsafe.Pointer(cbinningMode))))
}

func (n *NetCom) COM_SetDynamicFps(fps100Set uint32) bool {
	return bool(C.COM_SetDynamicFps(C.UINT32(fps100Set)))
}

func (n *NetCom) COM_GetDynamicFps(pfps100 *uint32) bool {
	return bool(C.COM_GetDynamicFps((*C.UINT32)(unsafe.Pointer(pfps100))))
}

func (n *NetCom) COM_SetDyncParaEnd() bool {
	return bool(C.COM_SetDyncParaEnd())
}

func (n *NetCom) COM_SetRoiPara(startRow uint16, endRow uint16, startCol uint16, endCol uint16) bool {
	return bool(C.COM_SetRoiPara(C.USHORT(startRow), C.USHORT(endRow), C.USHORT(startCol), C.USHORT(endCol)))
}

func (n *NetCom) COM_GetRoiPara(startRow *uint16, endRow *uint16, startCol *uint16, endCol *uint16) bool {
	return bool(C.COM_GetRoiPara((*C.USHORT)(unsafe.Pointer(startRow)), (*C.USHORT)(unsafe.Pointer(endRow)), (*C.USHORT)(unsafe.Pointer(startCol)), (*C.USHORT)(unsafe.Pointer(endCol))))
}

func (n *NetCom) COM_SetIfsRef(cbinningMode byte, cIfs byte, cRef byte) bool {
	return bool(C.COM_SetIfsRef(C.CHAR(cbinningMode), C.UCHAR(cIfs), C.UCHAR(cRef)))
}

func (n *NetCom) COM_GetIfsRef(cbinningMode byte, cIfs *byte, cRef *byte) bool {
	return bool(C.COM_GetIfsRef(C.CHAR(cbinningMode), (*C.UCHAR)(unsafe.Pointer(cIfs)), (*C.UCHAR)(unsafe.Pointer(cRef))))
}

func (n *NetCom) COM_SetGainValue(cbinningMode byte, ucGain byte) bool {
	return bool(C.COM_SetGainValue(C.CHAR(cbinningMode), C.UCHAR(ucGain)))
}

func (n *NetCom) COM_GetGainValue(cbinningMode byte, ucGain *byte) bool {
	return bool(C.COM_GetGainValue(C.CHAR(cbinningMode), (*C.UCHAR)(unsafe.Pointer(ucGain))))
}

func (n *NetCom) COM_SetFpIpNetmask(Ip uint32, netmask uint32) bool {
	return bool(C.COM_SetFpIpNetmask(C.UINT32(Ip), C.UINT32(netmask)))
}

func (n *NetCom) COM_GetFpIpNetmask(pIp *uint32, pnetmask *uint32) bool {
	return bool(C.COM_GetFpIpNetmask((*C.UINT32)(unsafe.Pointer(pIp)), (*C.UINT32)(unsafe.Pointer(pnetmask))))
}

func (n *NetCom) COM_SetSenValue(senValue uint16, ppassword string) bool {
	cPass := C.CString(ppassword)
	defer C.free(unsafe.Pointer(cPass))
	return bool(C.COM_SetSenValue(C.USHORT(senValue), cPass))
}

func (n *NetCom) COM_GetSenValue(psenValue *uint16) bool {
	return bool(C.COM_GetSenValue((*C.USHORT)(unsafe.Pointer(psenValue))))
}

func (n *NetCom) COM_SetClientSn(pClientSn string, ppassword string) bool {
	cSn := C.CString(pClientSn)
	cPass := C.CString(ppassword)
	defer C.free(unsafe.Pointer(cSn))
	defer C.free(unsafe.Pointer(cPass))
	return bool(C.COM_SetClientSn(cSn, cPass))
}

func (n *NetCom) COM_GetClientSn(pClientSn string) bool {
	cSn := C.CString(pClientSn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_GetClientSn(cSn))
}

func (n *NetCom) COM_SetClientPn(pClientPn string, ppassword string) bool {
	cPn := C.CString(pClientPn)
	cPass := C.CString(ppassword)
	defer C.free(unsafe.Pointer(cPn))
	defer C.free(unsafe.Pointer(cPass))
	return bool(C.COM_SetClientPn(cPn, cPass))
}

func (n *NetCom) COM_GetClientPn(pClientPn string) bool {
	cPn := C.CString(pClientPn)
	defer C.free(unsafe.Pointer(cPn))
	return bool(C.COM_GetClientPn(cPn))
}

func (n *NetCom) COM_SetNickname(pNickname string) bool {
	cName := C.CString(pNickname)
	defer C.free(unsafe.Pointer(cName))
	return bool(C.COM_SetNickname(cName))
}

func (n *NetCom) COM_GetNickname(pNickname string) bool {
	cName := C.CString(pNickname)
	defer C.free(unsafe.Pointer(cName))
	return bool(C.COM_GetNickname(cName))
}

func (n *NetCom) COM_SetExtBattDefaultCapacity(iCapacity int) bool {
	return bool(C.COM_SetExtBattDefaultCapacity(C.int(iCapacity)))
}

func (n *NetCom) COM_GetExtBattDefaultCapacity(piCapacity *int) bool {
	return bool(C.COM_GetExtBattDefaultCapacity((*C.int)(unsafe.Pointer(piCapacity))))
}

func (n *NetCom) COM_SetAecEnable(aecGroup uint16) bool {
	return bool(C.COM_SetAecEnable(C.UINT16(aecGroup)))
}

func (n *NetCom) COM_SetAecThreshold(aecNum byte, valueset uint16) bool {
	return bool(C.COM_SetAecThreshold(C.CHAR(aecNum), C.UINT16(valueset)))
}

func (n *NetCom) COM_GetInfoModified(pInfo string) bool {
	cInfo := C.CString(pInfo)
	defer C.free(unsafe.Pointer(cInfo))
	return bool(C.COM_GetInfoModified(cInfo))
}

func (n *NetCom) COM_SetImgAverageNum(setvalue uint16) bool {
	return bool(C.COM_SetImgAverageNum(C.UINT16(setvalue)))
}

func (n *NetCom) COM_SetOffsetAverageNum(setvalue byte) bool {
	return bool(C.COM_SetOffsetAverageNum(C.CHAR(setvalue)))
}

func (n *NetCom) COM_GetOffsetAverageNum(psetvalue *byte) bool {
	return bool(C.COM_GetOffsetAverageNum((*C.CHAR)(unsafe.Pointer(psetvalue))))
}

func (n *NetCom) COM_GetFPPnandSn(pPnSn string) bool {
	cSn := C.CString(pPnSn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_GetFPPnandSn(cSn))
}

func (n *NetCom) COM_SetSdkLogLevel(sdklevel byte) bool {
	return bool(C.COM_SetSdkLogLevel(C.UCHAR(sdklevel)))
}

func (n *NetCom) COM_GetSdkLogLevel(psdklevel *byte) bool {
	return bool(C.COM_GetSdkLogLevel((*C.UCHAR)(unsafe.Pointer(psdklevel))))
}

// 设备状态获取

func (n *NetCom) COM_GetFPsn(psn string) bool {
	cSn := C.CString(psn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_GetFPsn(cSn))
}

func (n *NetCom) COM_GetFPCurStatus() byte {
	return byte(C.COM_GetFPCurStatus())
}

func (n *NetCom) COM_GetFPWireState() byte {
	return byte(C.COM_GetFPWireState())
}

func (n *NetCom) COM_GetFpPowerMode() uint32 {
	return uint32(C.COM_GetFpPowerMode())
}

func (n *NetCom) COM_GetFpWorkState() byte {
	return byte(C.COM_GetFpWorkState())
}

func (n *NetCom) COM_GetFpPendingState() byte {
	return byte(C.COM_GetFpPendingState())
}

func (n *NetCom) COM_ClearPendingState() bool {
	return bool(C.COM_ClearPendingState())
}

func (n *NetCom) COM_GetFPStatus(ptFPStat *TFPStat) bool {
	cStat := (*C.TFPStat)(unsafe.Pointer(ptFPStat))
	return bool(C.COM_GetFPStatus(cStat))
}

func (n *NetCom) COM_GetFPStatusP(ptFPStatex *TFPStatex) bool {
	cStat := (*C.TFPStatex)(unsafe.Pointer(ptFPStatex))
	return bool(C.COM_GetFPStatusP(cStat))
}

func (n *NetCom) COM_GetConnectEssid(pessid string) bool {
	cEssid := C.CString(pessid)
	defer C.free(unsafe.Pointer(cEssid))
	return bool(C.COM_GetConnectEssid(cEssid))
}

// 运动和冲击检测

func (n *NetCom) COM_QuaternionActivate(enableflag byte) bool {
	return bool(C.COM_QuaternionActivate(C.CHAR(enableflag)))
}

func (n *NetCom) COM_GetFPMotionFeatures(tMotionFeatures *TMotionFeatures) bool {
	cFeatures := (*C.TMotionFeatures)(unsafe.Pointer(tMotionFeatures))
	return bool(C.COM_GetFPMotionFeatures(cFeatures))
}

func (n *NetCom) COM_GetFPShock(pShockInfo *TShockInfo) bool {
	cShock := (*C.TShockInfo)(unsafe.Pointer(pShockInfo))
	return bool(C.COM_GetFPShock(cShock))
}

func (n *NetCom) COM_ClearFPShock(ppassword string) bool {
	cPass := C.CString(ppassword)
	defer C.free(unsafe.Pointer(cPass))
	return bool(C.COM_ClearFPShock(cPass))
}

func (n *NetCom) COM_SetPacketPauseTime(pauseT_us uint16) bool {
	return bool(C.COM_SetPacketPauseTime(C.USHORT(pauseT_us)))
}

func (n *NetCom) COM_SetFpSelfStart(enableflag byte) bool {
	return bool(C.COM_SetFpSelfStart(C.CHAR(enableflag)))
}

func (n *NetCom) COM_SetFpTime2TurnOffAfterDisc(time_mins uint32) bool {
	return bool(C.COM_SetFpTime2TurnOffAfterDisc(C.UINT32(time_mins)))
}

func (n *NetCom) COM_SetFpAutoAedOffline(enableflag byte) bool {
	return bool(C.COM_SetFpAutoAedOffline(C.CHAR(enableflag)))
}

// 模板操作

func (n *NetCom) COM_SetAllTpl() bool {
	return bool(C.COM_SetAllTpl())
}

func (n *NetCom) COM_UploadOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_UploadOffsetTpl(cPath))
}

func (n *NetCom) COM_DownloadOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_DownloadOffsetTpl(cPath))
}

func (n *NetCom) COM_UploadGainTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_UploadGainTpl(cPath))
}

func (n *NetCom) COM_DownloadDefectMap(pData []byte) bool {
	return bool(C.COM_DownloadDefectMap((*C.CHAR)(unsafe.Pointer(&pData[0]))))
}

func (n *NetCom) COM_DownloadDefectMapV(pData []byte) bool {
	return bool(C.COM_DownloadDefectMapV((*C.CHAR)(unsafe.Pointer(&pData[0]))))
}

func (n *NetCom) COM_DownloadGainTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_DownloadGainTpl(cPath))
}

func (n *NetCom) COM_UploadDefectTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_UploadDefectTpl(cPath))
}

func (n *NetCom) COM_UploadAedTOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_UploadAedTOffsetTpl(cPath))
}

func (n *NetCom) COM_DownloadAedTOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_DownloadAedTOffsetTpl(cPath))
}

func (n *NetCom) COM_DownloadDefectTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_DownloadDefectTpl(cPath))
}

func (n *NetCom) COM_UploadFPZMTpl(TplType byte, Tplpath string) bool {
	cPath := C.CString(Tplpath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_UploadFPZMTpl(C.CHAR(TplType), cPath))
}

func (n *NetCom) COM_DownLoadFPZMTpl(TplType byte, Tplpath string) bool {
	cPath := C.CString(Tplpath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_DownLoadFPZMTpl(C.CHAR(TplType), cPath))
}

func (n *NetCom) COM_SetOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetOffsetTpl(cPath))
}

func (n *NetCom) COM_SetGainTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetGainTpl(cPath))
}

func (n *NetCom) COM_SetDefectTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetDefectTpl(cPath))
}

func (n *NetCom) COM_GenOffsetTpl() bool {
	return bool(C.COM_GenOffsetTpl())
}

func (n *NetCom) COM_GenGainTpl() bool {
	return bool(C.COM_GenGainTpl())
}

func (n *NetCom) COM_GenDefectTpl() bool {
	return bool(C.COM_GenDefectTpl())
}

func (n *NetCom) COM_CalibOffsetTpl(pData []byte) bool {
	return bool(C.COM_CalibOffsetTpl((*C.CHAR)(unsafe.Pointer(&pData[0]))))
}

func (n *NetCom) COM_CalibGainTpl(pData []byte) bool {
	return bool(C.COM_CalibGainTpl((*C.CHAR)(unsafe.Pointer(&pData[0]))))
}

func (n *NetCom) COM_CalibDefectTpl(pData []byte) bool {
	return bool(C.COM_CalibDefectTpl((*C.CHAR)(unsafe.Pointer(&pData[0]))))
}

func (n *NetCom) COM_TplPathSet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_TplPathSet(cPath))
}

func (n *NetCom) COM_TplPathGet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_TplPathGet(cPath))
}

func (n *NetCom) COM_ImgPathSet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_ImgPathSet(cPath))
}

func (n *NetCom) COM_ImgPathGet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_ImgPathGet(cPath))
}

func (n *NetCom) COM_LogPathSet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_LogPathSet(cPath))
}

func (n *NetCom) COM_LogPathGet(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_LogPathGet(cPath))
}

func (n *NetCom) COM_BatLow1Get() int32 {
	return int32(C.COM_BatLow1Get())
}

func (n *NetCom) COM_BatLow2Get() int32 {
	return int32(C.COM_BatLow2Get())
}

func (n *NetCom) COM_BatLow1Set(iBatLow int32) bool {
	return bool(C.COM_BatLow1Set(C.INT32(iBatLow)))
}

func (n *NetCom) COM_BatLow2Set(iBatLow int32) bool {
	return bool(C.COM_BatLow2Set(C.INT32(iBatLow)))
}

func (n *NetCom) COM_SdkLog(on bool) bool {
	return bool(C.COM_SdkLog(C.BOOL(0)))
}

func (n *NetCom) COM_EditDefectTpl(x uint16, y uint16, valid bool) bool {
	return bool(C.COM_EditDefectTpl(C.UINT16(x), C.UINT16(y), C.BOOL(0)))
}

func (n *NetCom) COM_EditLineDefectTpl(line uint16, bRow bool, valid bool) bool {
	return bool(C.COM_EditLineDefectTpl(C.UINT16(line), C.BOOL(0), C.BOOL(0)))
}

func (n *NetCom) COM_StartTplMakeProcess() bool {
	return bool(C.COM_StartTplMakeProcess())
}

func (n *NetCom) COM_EndTplMakeProcess() bool {
	return bool(C.COM_EndTplMakeProcess())
}

// 版本和信息获取

func (n *NetCom) COM_GetDllVer(pcSDKVer string) bool {
	cVer := C.CString(pcSDKVer)
	defer C.free(unsafe.Pointer(cVer))
	return bool(C.COM_GetDllVer(cVer))
}

func (n *NetCom) COM_GetRBInfo(ptRBInfo *TRBInfo) bool {
	cInfo := (*C.TRBInfo)(unsafe.Pointer(ptRBInfo))
	return bool(C.COM_GetRBInfo(cInfo))
}

func (n *NetCom) COM_GetFPInfo(ptFPInfo *TFPInfo) bool {
	cInfo := (*C.TFPInfo)(unsafe.Pointer(ptFPInfo))
	return bool(C.COM_GetFPInfo(cInfo))
}

func (n *NetCom) COM_GetRBStatus(pcRBStatus *byte) bool {
	return bool(C.COM_GetRBStatus((*C.CHAR)(unsafe.Pointer(pcRBStatus))))
}

func (n *NetCom) COM_GetFPType() byte {
	return byte(C.COM_GetFPType())
}

func (n *NetCom) COM_SetFPType(ucFpType byte) bool {
	return bool(C.COM_SetFPType(C.CHAR(ucFpType)))
}

func (n *NetCom) COM_GetFPCompatibleVer() byte {
	return byte(C.COM_GetFPCompatibleVer())
}

// 固件升级

func (n *NetCom) COM_FpVerUpgrade(pcVerPath string) bool {
	cPath := C.CString(pcVerPath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_FpVerUpgrade(cPath))
}

func (n *NetCom) COM_FpgaVerUpgrade(pcVerPath string) bool {
	cPath := C.CString(pcVerPath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_FpgaVerUpgrade(cPath))
}

func (n *NetCom) COM_McuVerUpgrade(pcVerPath string) bool {
	cPath := C.CString(pcVerPath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_McuVerUpgrade(cPath))
}

func (n *NetCom) COM_RbVerUpgrade(pcVerPath string) bool {
	cPath := C.CString(pcVerPath)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_RbVerUpgrade(cPath))
}

// 状态回调注册

func (n *NetCom) COM_Register_FP_STATUS(fpStatus unsafe.Pointer) bool {
	return bool(C.COM_Register_FP_STATUS((C.FP_STATUS_CALLBACK)(fpStatus)))
}

func (n *NetCom) COM_Register_RB_STATUS(rbStatus unsafe.Pointer) bool {
	return bool(C.COM_Register_RB_STATUS((C.RB_STATUS_CALLBACK)(rbStatus)))
}

func (n *NetCom) COM_Register_IMAGE_RECEIVED(imageReceived unsafe.Pointer) bool {
	return bool(C.COM_Register_IMAGE_RECEIVED((C.IMAGE_RECEIVED_CALLBACK)(imageReceived)))
}

func (n *NetCom) COM_Register_COMMAND_CB(cmdCallback unsafe.Pointer) bool {
	return bool(C.COM_Register_COMMAND_CB((C.COMMAND_CALLBACK)(cmdCallback)))
}

// 命令发送

func (n *NetCom) COM_SendCMD(pSendData string, nSize int32, us_MCmd uint16, us_SCmd uint16) bool {
	cData := C.CString(pSendData)
	defer C.free(unsafe.Pointer(cData))
	return bool(C.COM_SendCMD(cData, C.INT32(nSize), C.UINT16(us_MCmd), C.UINT16(us_SCmd)))
}

func (n *NetCom) COM_SetFPFullConf() bool {
	return bool(C.COM_SetFPFullConf())
}

func (n *NetCom) COM_GetFPFullConf() bool {
	return bool(C.COM_GetFPFullConf())
}

func (n *NetCom) COM_SetRBFullConf() bool {
	return bool(C.COM_SetRBFullConf())
}

func (n *NetCom) COM_GetFullConf() bool {
	return bool(C.COM_GetFullConf())
}

func (n *NetCom) COM_GenAEDParam() bool {
	return bool(C.COM_GenAEDParam())
}

func (n *NetCom) COM_SetAedCorrKB() bool {
	return bool(C.COM_SetAedCorrKB())
}

func (n *NetCom) COM_PrintLog(logText string) bool {
	cLog := C.CString(logText)
	defer C.free(unsafe.Pointer(cLog))
	return bool(C.COM_PrintLog(cLog))
}

func (n *NetCom) COM_SetCaliIntNum(usCount byte) bool {
	return bool(C.COM_SetCaliIntNum(C.uchar(usCount)))
}

func (n *NetCom) COM_AEDTriggerByHst() bool {
	return bool(C.COM_AEDTriggerByHst())
}

func (n *NetCom) COM_HstTriggerPre() bool {
	return bool(C.COM_HstTriggerPre())
}

func (n *NetCom) COM_AEDTriggerByHstStop() bool {
	return bool(C.COM_AEDTriggerByHstStop())
}

func (n *NetCom) COM_GetMeanDose(u16Image []uint16, size uint32) uint16 {
	return uint16(C.COM_GetMeanDose((*C.UINT16)(unsafe.Pointer(&u16Image[0])), C.UINT32(size)))
}

func (n *NetCom) COM_SetAedTOffsetTpl(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetAedTOffsetTpl(cPath))
}

func (n *NetCom) COM_CalibAedTTpl(pData []byte, u16LineNum uint16) bool {
	return bool(C.COM_CalibAedTTpl((*C.char)(unsafe.Pointer(&pData[0])), C.USHORT(u16LineNum)))
}

func (n *NetCom) COM_GetFPLicense(tLicenseInfo *TLicenseInfo, choose byte) bool {
	cInfo := (*C.TLicenseInfo)(unsafe.Pointer(tLicenseInfo))
	return bool(C.COM_GetFPLicense(cInfo, C.CHAR(choose)))
}

func (n *NetCom) COM_GetFPTypeFromHardware() byte {
	return byte(C.COM_GetFPTypeFromHardware())
}

func (n *NetCom) COM_SetRecoverTime(ucTime byte) bool {
	return bool(C.COM_SetRecoverTime(C.uchar(ucTime)))
}

func (n *NetCom) COM_GetRecoverTime() byte {
	return byte(C.COM_GetRecoverTime())
}

// 多FP API (Ex版本)

func (n *NetCom) COM_SetPreCalibModeEx(nCalMode byte, index byte) bool {
	return bool(C.COM_SetPreCalibModeEx(C.CHAR(nCalMode), C.CHAR(index)))
}

func (n *NetCom) COM_HstAcqEx(index byte) bool {
	return bool(C.COM_HstAcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_AedAcqEx(index byte) bool {
	return bool(C.COM_AedAcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_TriggerEx(index byte) bool {
	return bool(C.COM_TriggerEx(C.CHAR(index)))
}

func (n *NetCom) COM_Trigger2Ex(index byte) bool {
	return bool(C.COM_Trigger2Ex(C.CHAR(index)))
}

func (n *NetCom) COM_PrepEx(index byte) bool {
	return bool(C.COM_PrepEx(C.CHAR(index)))
}

func (n *NetCom) COM_AcqEx(index byte) bool {
	return bool(C.COM_AcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_PrepAcqEx(index byte) bool {
	return bool(C.COM_PrepAcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_ExposeReqEx(index byte) bool {
	return bool(C.COM_ExposeReqEx(C.CHAR(index)))
}

func (n *NetCom) COM_SetAcqEx(index byte) bool {
	return bool(C.COM_SetAcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_StopEx(index byte) bool {
	return bool(C.COM_StopEx(C.CHAR(index)))
}

func (n *NetCom) COM_DstEx(index byte) bool {
	return bool(C.COM_DstEx(C.CHAR(index)))
}

func (n *NetCom) COM_DacqEx(index byte) bool {
	return bool(C.COM_DacqEx(C.CHAR(index)))
}

func (n *NetCom) COM_DacqaedEx(index byte) bool {
	return bool(C.COM_DacqaedEx(C.CHAR(index)))
}

func (n *NetCom) COM_CbctEx(index byte) bool {
	return bool(C.COM_CbctEx(C.CHAR(index)))
}

func (n *NetCom) COM_Cbct2Ex(index byte) bool {
	return bool(C.COM_Cbct2Ex(C.CHAR(index)))
}

func (n *NetCom) COM_DexitEx(index byte) bool {
	return bool(C.COM_DexitEx(C.CHAR(index)))
}

func (n *NetCom) COM_DprepEx(index byte) bool {
	return bool(C.COM_DprepEx(C.CHAR(index)))
}

func (n *NetCom) COM_CprepEx(index byte) bool {
	return bool(C.COM_CprepEx(C.CHAR(index)))
}

func (n *NetCom) COM_ExprepEx(index byte) bool {
	return bool(C.COM_ExprepEx(C.CHAR(index)))
}

func (n *NetCom) COM_AedPrepEx(index byte) bool {
	return bool(C.COM_AedPrepEx(C.CHAR(index)))
}

func (n *NetCom) COM_Aed2AcqEx(index byte) bool {
	return bool(C.COM_Aed2AcqEx(C.CHAR(index)))
}

func (n *NetCom) COM_ImgSaveStartEx(index byte) bool {
	return bool(C.COM_ImgSaveStartEx(C.CHAR(index)))
}

func (n *NetCom) COM_ImgSaveExitEx(index byte) bool {
	return bool(C.COM_ImgSaveExitEx(C.CHAR(index)))
}

func (n *NetCom) COM_ImgSavedReadEx(index byte) bool {
	return bool(C.COM_ImgSavedReadEx(C.CHAR(index)))
}

func (n *NetCom) COM_RegisterEvCallBackEx(nEvent int16, funcallbackex unsafe.Pointer) bool {
	return bool(C.COM_RegisterEvCallBackEx(C.INT16(nEvent), (C.FP_EVENT_CALLBACKEX)(funcallbackex)))
}

func (n *NetCom) COM_GetImageModeEx(ptImageMode *TImageMode, index byte) bool {
	cMode := (*C.TImageMode)(unsafe.Pointer(ptImageMode))
	return bool(C.COM_GetImageModeEx(cMode, C.CHAR(index)))
}

func (n *NetCom) COM_GetImageModeVEx(ptImageMode *TImageMode, index byte) bool {
	cMode := (*C.TImageMode)(unsafe.Pointer(ptImageMode))
	return bool(C.COM_GetImageModeVEx(cMode, C.CHAR(index)))
}

func (n *NetCom) COM_GetImageShiftModeEx(ptImageShiftMode *TImageShiftMode, index byte) bool {
	cShift := (*C.TImageShiftMode)(unsafe.Pointer(ptImageShiftMode))
	return bool(C.COM_GetImageShiftModeEx(cShift, C.CHAR(index)))
}

func (n *NetCom) COM_GetImageNameEx(name string, index byte) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return bool(C.COM_GetImageNameEx(cName, C.CHAR(index)))
}

func (n *NetCom) COM_GetImageIDEx(pimgID *uint32, index byte) bool {
	return bool(C.COM_GetImageIDEx((*C.UINT32)(unsafe.Pointer(pimgID)), C.CHAR(index)))
}

func (n *NetCom) COM_GetImageEx(pImageBuff []byte, index byte) bool {
	return bool(C.COM_GetImageEx((*C.CHAR)(unsafe.Pointer(&pImageBuff[0])), C.CHAR(index)))
}

func (n *NetCom) COM_GetImageVEx(pImageBuff []byte, index byte) bool {
	return bool(C.COM_GetImageVEx((*C.CHAR)(unsafe.Pointer(&pImageBuff[0])), C.CHAR(index)))
}

func (n *NetCom) COM_SetXwinEx(xwin uint32, index byte) bool {
	return bool(C.COM_SetXwinEx(C.UINT32(xwin), C.CHAR(index)))
}

func (n *NetCom) COM_SetXwin_usEx(xwin_us uint32, index byte) bool {
	return bool(C.COM_SetXwin_usEx(C.UINT32(xwin_us), C.CHAR(index)))
}

func (n *NetCom) COM_GetFPLicenseEx(tLicenseInfo *TLicenseInfo, choose byte, index byte) bool {
	cInfo := (*C.TLicenseInfo)(unsafe.Pointer(tLicenseInfo))
	return bool(C.COM_GetFPLicenseEx(cInfo, C.CHAR(choose), C.CHAR(index)))
}

func (n *NetCom) COM_SetMetaDataEx(tMetaData *TMetaData, index byte) bool {
	cMeta := (*C.TMetaData)(unsafe.Pointer(tMetaData))
	return bool(C.COM_SetMetaDataEx(*cMeta, C.CHAR(index)))
}

func (n *NetCom) COM_GetMetaDataEx(ptMetaData *TMetaData, index byte) bool {
	cMeta := (*C.TMetaData)(unsafe.Pointer(ptMetaData))
	return bool(C.COM_GetMetaDataEx(cMeta, C.CHAR(index)))
}

func (n *NetCom) COM_GetFpPendingStateEx(index byte) byte {
	return byte(C.COM_GetFpPendingStateEx(C.CHAR(index)))
}

func (n *NetCom) COM_ClearPendingStateEx(index byte) bool {
	return bool(C.COM_ClearPendingStateEx(C.CHAR(index)))
}

func (n *NetCom) COM_GetFPsnEx(index byte, psn string) bool {
	cSn := C.CString(psn)
	defer C.free(unsafe.Pointer(cSn))
	return bool(C.COM_GetFPsnEx(C.CHAR(index), cSn))
}

func (n *NetCom) COM_GetFPCurStatusEx(index byte) byte {
	return byte(C.COM_GetFPCurStatusEx(C.CHAR(index)))
}

func (n *NetCom) COM_GetFPWireStateEx(index byte) byte {
	return byte(C.COM_GetFPWireStateEx(C.CHAR(index)))
}

func (n *NetCom) COM_GetFpPowerModeEx(index byte) uint32 {
	return uint32(C.COM_GetFpPowerModeEx(C.CHAR(index)))
}

func (n *NetCom) COM_GetFpWorkStateEx(index byte) byte {
	return byte(C.COM_GetFpWorkStateEx(C.CHAR(index)))
}

func (n *NetCom) COM_GetFPStatusEx(ptFPStat *TFPStat, index byte) bool {
	cStat := (*C.TFPStat)(unsafe.Pointer(ptFPStat))
	return bool(C.COM_GetFPStatusEx(cStat, C.CHAR(index)))
}

func (n *NetCom) COM_GetFPStatusPEx(ptFPStatex *TFPStatex, index byte) bool {
	cStat := (*C.TFPStatex)(unsafe.Pointer(ptFPStatex))
	return bool(C.COM_GetFPStatusPEx(cStat, C.CHAR(index)))
}

func (n *NetCom) COM_GetConnectEssidEx(pessid string, index byte) bool {
	cEssid := C.CString(pessid)
	defer C.free(unsafe.Pointer(cEssid))
	return bool(C.COM_GetConnectEssidEx(cEssid, C.CHAR(index)))
}

func (n *NetCom) COM_SetDynamicParaEx(xwin uint32, repeat uint16, binMode byte, sync byte, index byte) bool {
	return bool(C.COM_SetDynamicParaEx(C.UINT32(xwin), C.UINT16(repeat), C.CHAR(binMode), C.CHAR(sync), C.CHAR(index)))
}

func (n *NetCom) COM_GetDynamicParaEx(pxwin *uint32, prepeat *uint16, pbinMode *byte, psync *byte, index byte) bool {
	return bool(C.COM_GetDynamicParaEx((*C.UINT32)(unsafe.Pointer(pxwin)), (*C.UINT16)(unsafe.Pointer(prepeat)), (*C.CHAR)(unsafe.Pointer(pbinMode)), (*C.CHAR)(unsafe.Pointer(psync)), C.CHAR(index)))
}

func (n *NetCom) COM_SetBinningModeEx(cbinningMode byte, index byte) bool {
	return bool(C.COM_SetBinningModeEx(C.CHAR(cbinningMode), C.CHAR(index)))
}

func (n *NetCom) COM_SetDynamicFpsEx(fps100Set uint32, index byte) bool {
	return bool(C.COM_SetDynamicFpsEx(C.UINT32(fps100Set), C.CHAR(index)))
}

func (n *NetCom) COM_GetDynamicFpsEx(pfps100 *uint32, index byte) bool {
	return bool(C.COM_GetDynamicFpsEx((*C.UINT32)(unsafe.Pointer(pfps100)), C.CHAR(index)))
}

func (n *NetCom) COM_SetDyncParaEndEx(index byte) bool {
	return bool(C.COM_SetDyncParaEndEx(C.CHAR(index)))
}

func (n *NetCom) COM_SetRoiParaEx(startRow uint16, endRow uint16, startCol uint16, endCol uint16, index byte) bool {
	return bool(C.COM_SetRoiParaEx(C.USHORT(startRow), C.USHORT(endRow), C.USHORT(startCol), C.USHORT(endCol), C.CHAR(index)))
}

func (n *NetCom) COM_GetRoiParaEx(startRow *uint16, endRow *uint16, startCol *uint16, endCol *uint16, index byte) bool {
	return bool(C.COM_GetRoiParaEx((*C.USHORT)(unsafe.Pointer(startRow)), (*C.USHORT)(unsafe.Pointer(endRow)), (*C.USHORT)(unsafe.Pointer(startCol)), (*C.USHORT)(unsafe.Pointer(endCol)), C.CHAR(index)))
}

func (n *NetCom) COM_SetIfsRefEx(cbinningMode byte, cIfs byte, cRef byte, index byte) bool {
	return bool(C.COM_SetIfsRefEx(C.CHAR(cbinningMode), C.UCHAR(cIfs), C.UCHAR(cRef), C.CHAR(index)))
}

func (n *NetCom) COM_GetIfsRefEx(cbinningMode byte, cIfs *byte, cRef *byte, index byte) bool {
	return bool(C.COM_GetIfsRefEx(C.CHAR(cbinningMode), (*C.UCHAR)(unsafe.Pointer(cIfs)), (*C.UCHAR)(unsafe.Pointer(cRef)), C.CHAR(index)))
}

func (n *NetCom) COM_SetOffsetAverageNumEx(setvalue byte, index byte) bool {
	return bool(C.COM_SetOffsetAverageNumEx(C.CHAR(setvalue), C.CHAR(index)))
}

func (n *NetCom) COM_GetOffsetAverageNumEx(psetvalue *byte, index byte) bool {
	return bool(C.COM_GetOffsetAverageNumEx((*C.CHAR)(unsafe.Pointer(psetvalue)), C.CHAR(index)))
}

func (n *NetCom) COM_SetAllTplEx(index byte) bool {
	return bool(C.COM_SetAllTplEx(C.CHAR(index)))
}

func (n *NetCom) COM_SetOffsetTplEx(path string, TplType byte, index byte) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetOffsetTplEx(cPath, C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_SetGainTplEx(path string, TplType byte, index byte) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetGainTplEx(cPath, C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_SetDefectTplEx(path string, TplType byte, index byte) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetDefectTplEx(cPath, C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_GenOffsetTplEx(TplType byte) bool {
	return bool(C.COM_GenOffsetTplEx(C.CHAR(TplType)))
}

func (n *NetCom) COM_GenGainTplEx(TplType byte) bool {
	return bool(C.COM_GenGainTplEx(C.CHAR(TplType)))
}

func (n *NetCom) COM_GenDefectTplEx(TplType byte) bool {
	return bool(C.COM_GenDefectTplEx(C.CHAR(TplType)))
}

func (n *NetCom) COM_CalibOffsetTplEx(pData []byte, TplType byte, index byte) bool {
	return bool(C.COM_CalibOffsetTplEx((*C.CHAR)(unsafe.Pointer(&pData[0])), C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_CalibGainTplEx(pData []byte, TplType byte, index byte) bool {
	return bool(C.COM_CalibGainTplEx((*C.CHAR)(unsafe.Pointer(&pData[0])), C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_CalibDefectTplEx(pData []byte, TplType byte, index byte) bool {
	return bool(C.COM_CalibDefectTplEx((*C.CHAR)(unsafe.Pointer(&pData[0])), C.CHAR(TplType), C.CHAR(index)))
}

func (n *NetCom) COM_GetFPTypeEx(index byte) byte {
	return byte(C.COM_GetFPTypeEx(C.CHAR(index)))
}

func (n *NetCom) COM_EditDefectTplEx(x uint16, y uint16, valid bool, TplType byte) bool {
	return bool(C.COM_EditDefectTplEx(C.UINT16(x), C.UINT16(y), C.BOOL(0), C.CHAR(TplType)))
}

func (n *NetCom) COM_EditLineDefectTplEx(line uint16, bRow bool, valid bool, TplType byte) bool {
	return bool(C.COM_EditLineDefectTplEx(C.UINT16(line), C.BOOL(0), C.BOOL(0), C.CHAR(TplType)))
}

func (n *NetCom) COM_SetAedTOffsetTplEx(path string, index byte) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.COM_SetAedTOffsetTplEx(cPath, C.CHAR(index)))
}

func (n *NetCom) COM_CalibAedTTplEx(pData []byte, u16LineNum uint16, index byte) bool {
	return bool(C.COM_CalibAedTTplEx((*C.char)(unsafe.Pointer(&pData[0])), C.USHORT(u16LineNum), C.CHAR(index)))
}

func (n *NetCom) COM_GetFPUseEx(pUse *byte, index byte) bool {
	return bool(C.COM_GetFPUseEx((*C.UCHAR)(unsafe.Pointer(pUse)), C.CHAR(index)))
}

func (n *NetCom) COM_UseActiveEx(index byte) bool {
	return bool(C.COM_UseActiveEx(C.CHAR(index)))
}
