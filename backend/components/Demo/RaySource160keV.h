/***************************************************************************/
/****************160keV射线源串口通信头文件（作者：ShiPeipei）**************/
/***************************************************************************/
#ifndef XRAY_160KEV_H
#define XRAY_160KEV_H

#include <afxwin.h>
#include <vector>
#include <string>

// 通讯协议相关定义
#define STX                         0x02                      // 帧头
#define ETX                         0x03                      // 帧尾
#define SEPARATOR                   0x2C                      // 间隔符（逗号）
#define DEFAULT_BAUDRATE CBR_9600                             // 默认波特率9600bps
#define RESPONSE_DELAY              200                       // 响应延迟（ms），协议规定控制盒2ms响应，预留冗余
#define COMM_INTERVAL               100                       // 通讯间隔（ms），协议建议不低于100ms

// 命令码定义
enum CmdCode
{
    SET_KV                  =       10,                       // 设置kV电压
    SET_MA                  =       11,                       // 设置mA电流
    READ_KV_SET             =       14,                       // 读取kV设定值
    READ_MA_SET             =       15,                       // 读取mA设定值
    READ_MODULE_STATUS      =       22,                       // 读取故障码
    READ_CODE_VERSION       =       23,                       // 读取软件版本号
    READ_HARDWRE_VERSION    =       24,                       // 读取硬件版本号
    READ_SERIAL_NO          =       25,                       // 读取出厂编码
    SERVICE_WATCHDOG        =       27,                       // 喂通讯看门狗
    WATCHDOG_ENA            =       28,                       // 使能看门狗
    RESET_FAULT_FLAG        =       32,                       // 清除故障标志位
    READ_NOW_KV             =       60,                       // 读取当前kV值
    READ_NOW_MA             =       61,                       // 读取当前mA值
    READ_OIL_TEMP           =       63,                       // 读取油温值
    READ_TOTAL_ARC_COUNT    =       72,                       // 读取设备总打火次数
    READ_LAST_ARC_COUNT     =       73,                       // 读取最近打火次数
    READ_XRAY_STATUS        =       96,                       // 读取射线源状态
    XRAY_ONOFF              =       99                        // 打开或关闭射线源
};

// 射线源状态定义
enum XRayStatus
{
    XRAY_OFF                =       0,                        // 关闭
    XRAY_ON                 =       1,                        // 打开
    XRAY_WARM               =       2                         // 预热
};

// 故障码位定义
enum FaultBit
{
    FAULT_LOCK              =       0,                        // B0：闭锁保护
    FAULT_OCP_LCC           =       1,                        // B1：OCP_LCC保护
    FAULT_KV_OVER           =       2,                        // B2：kV过压保护
    FAULT_KV_UNDER          =       3,                        // B3：kV欠压保护
    FAULT_ARC_LOCK          =       4,                        // B4：打火锁死保护
    FAULT_INPUT_VOLT        =       5,                        // B5：输入过欠压保护
    FAULT_TEMP              =       6,                        // B6：温度保护
    FAULT_REMOTE            =       7,                        // B7：Remote保护
    FAULT_MA_OVER           =       8,                        // B8：mA过流保护
    FAULT_MA_UNDER          =       9,                        // B9：mA欠流保护
    FAULT_POWER_OVER        =       10,                       // B10：过功率保护
    FAULT_WATCHDOG          =       11,                       // B11：看门狗溢出保护
    FAULT_XRAY_OFF          =       12,                       // B12：Xray关闭
    FAULT_COMM_TIMEOUT      =       13,                       // B13：通讯超时保护
    FAULT_FAST              =       14,                       // B14：快速保护
    FAULT_FAULT_LOCK        =       15                        // B15：故障锁死保护
};

// 射线源通讯核心类
class CRaySource160keVComm
{
public:
    CRaySource160keVComm();                                     // 构造函数
    ~CRaySource160keVComm();                                    // 析构函数

    // 核心接口（主程序调用入口）
    BOOL InitSerial(CString strCom);                            // 初始化串口（参数：串口名称，如COM1）
    BOOL CloseSerial();                                         // 关闭串口
    BOOL IsSerialOpen();                                        // 判断串口是否打开

    // 协议命令接口
    BOOL SetKV(double kV);                                      // 设置kV电压（单位：kV，内部自动转换为0.1kV的ASCII码）
    BOOL SetMA(int mA);                                         // 设置mA电流（单位：uA，直接转换为ASCII码）
    CString ReadKVSet();                                        // 读取kV设定值（返回：设定值字符串，单位0.1kV）
    CString ReadMASet();                                        // 读取mA设定值（返回：设定值字符串，单位1uA）
    CString ReadFaultCode();                                    // 读取故障码并解析（返回：故障描述字符串）
    CString ReadSoftwareVersion();                              // 读取软件版本号（返回：版本号字符串）
    CString ReadHardwareVersion();                              // 读取硬件版本号（返回：版本号字符串）
    CString ReadSerialNo();                                     // 读取出厂编码（返回：17位ASCII编码字符串）
    BOOL FeedWatchDog();                                        // 喂通讯看门狗（返回：TRUE=成功，FALSE=失败）
    BOOL EnableWatchDog(int time);                              // 使能看门狗（参数：时间0~99秒，0=不使能；返回：执行结果）
    BOOL ResetFaultFlag();                                      // 清除故障标志位（返回：TRUE=成功，FALSE=失败）
    CString ReadCurrentKV();                                    // 读取当前kV值（返回：当前值字符串，单位0.1kV）
    CString ReadCurrentMA();                                    // 读取当前mA值（返回：当前值字符串，单位1uA）
    CString ReadOilTemp();                                      // 读取油温值（返回：实际温度字符串，单位℃）
    CString ReadTotalArcCount();                                // 读取设备总打火次数（返回：次数字符串）
    CString ReadLastArcCount();                                 // 读取最近打火次数（返回：次数字符串）
    CString ReadXrayStatus();                                   // 读取射线源状态（返回：状态描述字符串）
    BOOL ControlXray(CString arg);                              // 打开/关闭射线源（参数：00=关闭，11=打开；返回：执行结果）

    // 日志获取接口（主程序可调用，获取通讯日志）
    CString GetLastLog();                                       // 获取最后一条通讯日志
    std::vector<CString> GetAllLog();                           // 获取所有通讯日志

private:
    // 串口相关成员（内部使用，主程序无需操作）
    HANDLE                          m_hCom;                     // 串口句柄
    DCB                             m_dcb;                      // 串口配置结构体
    COMMTIMEOUTS                    m_commTimeouts;             // 串口超时配置
    CString                         m_strCom;                   // 所选串口
    BOOL                            m_bComOpen;                 // 串口是否打开

    // 日志存储（内部使用）
    std::vector<CString>            m_vecLog;                   // 通讯日志列表

    // 内部核心功能（不对外暴露，主程序无需调用）
    BYTE CalculateCheckSum(const std::vector<BYTE>& data);      // 计算校验码（遵循协议规则）
    std::vector<BYTE> PackFrame(CmdCode cmd, const CString& arg = _T("")); // 封装数据帧
    BOOL SendCmd(CmdCode cmd, const CString& arg = _T(""));     // 发送命令
    BOOL ReceiveResponse(std::vector<BYTE>& response);          // 接收响应
    CString ParseResponse(CmdCode cmd, const std::vector<BYTE>& response); // 解析响应
    void AddLog(const CString& log);                            // 添加通讯日志（内部调用）
    BOOL IsValidFrame(const std::vector<BYTE>& frame);          // 验证帧格式合法性

    // 辅助函数（内部使用）
    CString ByteToHexStr(const BYTE& byte);                     // 单个字节转十六进制字符串
    CString BytesToHexStr(const std::vector<BYTE>& bytes);      // 字节数组转十六进制字符串
    int HexStrToInt(const CString& hexStr);                     // 十六进制字符串转十进制整数
    std::vector<BYTE> CStringToBytes(const CString& str);       // CString转字节数组
    void Delay(DWORD ms);                                       // 延时函数（适配协议通讯间隔要求）
};

#endif
