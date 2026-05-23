/******************************************************************************/
/***************高压电源网口Modbus通信CPP文件（作者：ShiPeipei）***************/
/******************************************************************************/
#include "pch.h" 
#include "ModbusHVPS.h"

ModbusHVPS::ModbusHVPS()
{    
    HVPS_socket = INVALID_SOCKET;
    WSADATA wsa;
    WSAStartup(MAKEWORD(2, 2), &wsa);
}

ModbusHVPS::~ModbusHVPS()
{
    DisconnectHVPS();
    WSACleanup();
}

// 连接高压电源设备
bool ModbusHVPS::ConnectHVPS(const char* HVPS_ip, int HVPS_port)
{
    HVPS_socket = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (HVPS_socket == INVALID_SOCKET)
    {
        return false;
    }

    sockaddr_in addr1{};
    addr1.sin_family = AF_INET;
    addr1.sin_port = htons(HVPS_port);
    inet_pton(AF_INET, HVPS_ip, &addr1.sin_addr);

    if (::connect(HVPS_socket, (sockaddr*)&addr1, sizeof(addr1)) == SOCKET_ERROR)
    {
        closesocket(HVPS_socket);
        HVPS_socket = INVALID_SOCKET;
        return false;
    }
    return true;
}
// 断开连接高压电源设备
bool ModbusHVPS::DisconnectHVPS()
{
    if (HVPS_socket != INVALID_SOCKET)
    {
        closesocket(HVPS_socket);
        HVPS_socket = INVALID_SOCKET;
        return true;
    }
    else
    {
        return false;
    }
}
// 发送数据 ---- 高压电源设备
bool ModbusHVPS::sendHVData(BYTE* HV_packet, int HV_len)
{
    return send(HVPS_socket, (char*)HV_packet, HV_len, 0) > 0;
}
// 接收数据 ---- 高压电源设备
bool ModbusHVPS::recvHVData(BYTE* HV_buf, int HV_len)
{
    return recv(HVPS_socket, (char*)HV_buf, HV_len, 0) > 0;
}
// 设置高压电源远程Remote模式指令
bool ModbusHVPS::setHV_Remote()
{
    BYTE RemCmd[] = { 0xA1, 0x65, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x0D };
    if (!sendHVData(RemCmd, 11)) return false;

    // 清空接收缓冲区
    memset(HV_Recv, 0, sizeof(HV_Recv));
    // 接收信息并进行验证
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);
    // 高压电源返回数据：如果长度为11且第6个字节为0x01，说明远程模式设置成功
    if (HV_RecvLen == 11 && HV_Recv[5] == 1)
    {
        return true;
    }
    else
    {
        return false;
    }
}
// 设置高压电源的电压并进行验证
bool ModbusHVPS::setHV_VI(double setvoltage_kV, double setcurrent_uA)
{
    // 读取用户输入的电压和电流值，转换成高压电源协议要求的格式，并计算Check位
    VolValue = static_cast<int>(setvoltage_kV * 4095 / 50 + 0.5);			// 将设置的电压值转为int型，四舍五入，电压最大为50 kV
    CurValue = static_cast<int>(setcurrent_uA * 4095 / 1000 + 0.5);			// 将设置的电压值转为int型，四舍五入，电流最大为1000 uA

    // 16位简化版累加和校验算法（双字节），计算每个值的二进制形式中数字“1”的总数
    CheckSum_Vol = _mm_popcnt_u32(0xA1) + _mm_popcnt_u32(0x61) + _mm_popcnt_u32(VolValue);
    CheckSum_Cur = _mm_popcnt_u32(0xA1) + _mm_popcnt_u32(0x62) + _mm_popcnt_u32(CurValue);

    // 设置电压指令形式为 A1 61 00 00 00 HIBYTE(VolValue) LOBYTE(VolValue) HIBYTE(CheckSum_Vol) LOBYTE(CheckSum_Vol) 00 0D
    BYTE Vol_Array[11] = { 
        0xA1, 0x61, 0x00, 0x00, 0x00, 
        HIBYTE(VolValue), LOBYTE(VolValue), HIBYTE(CheckSum_Vol), LOBYTE(CheckSum_Vol),
        0x00, 0x0D 
    };
    // 设置电流指令形式为 A1 62 00 00 00 HIBYTE(CurValue) LOBYTE(CurValue) HIBYTE(CheckSum_Cur) LOBYTE(CheckSum_Cur) 00 0D
    BYTE Cur_Array[11] = { 
        0xA1, 0x62, 0x00, 0x00, 0x00, 
        HIBYTE(CurValue), LOBYTE(CurValue), HIBYTE(CheckSum_Cur), LOBYTE(CheckSum_Cur),
        0x00, 0x0D 
    };

    // 清空缓存区，发送设置电压的指令，接收指令，判断指令
    memset(HV_Recv, 0, sizeof(HV_Recv));                                    // 清空接收缓冲区
    if (!sendHVData(Vol_Array, 11))             return false;
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);
    if (HV_RecvLen < 0 || HV_Recv[2] != 10)     return false;               // 返回数据第3个字节为0x0A，说明格式正确

    // 清空缓存区，发送设置电流的指令，接收指令，判断指令    
    memset(HV_Recv, 0, sizeof(HV_Recv));                                    // 清空接收缓冲区
    if (!sendHVData(Cur_Array, 11))             return false;
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);
    if (HV_RecvLen < 0 || HV_Recv[2] != 10)     return false;               // 返回数据第3个字节为0x0A，说明格式正确

    return true;                                                            // 返回0表示设置成功
}

// 打开高压电源输出
bool ModbusHVPS::HV_ON()
{
    BYTE HV_ON[11] = { 0xA1, 0x69, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x0D };

    // 发送设置电压的指令
    if (!sendHVData(HV_ON, 11)) return false;
    // 接收信息并进行验证
    memset(HV_Recv, 0, sizeof(HV_Recv));                                        // 清空接收缓冲区
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);

    // 高压电源返回数据：如果长度为11且第6个字节为0x01，说明HV_ON命令执行成功
    if (HV_RecvLen > 0 && HV_Recv[5] == 1)
    {
        return true;
    }
    else
    {
        return false;
    }
}
// 关闭高压电源输出
bool ModbusHVPS::HV_OFF()
{
    BYTE HV_OFF[11] = { 0xA1, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x0D };

    // 发送设置电压的指令
    if (!sendHVData(HV_OFF, 11)) return false;
    // 接收信息并进行验证
    memset(HV_Recv, 0, sizeof(HV_Recv));                                        // 清空接收缓冲区
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);

    // 高压电源返回数据：如果长度为11且第6个字节为0x01，说明HV_ON命令执行成功
    if (HV_RecvLen > 0 && HV_Recv[5] == 0)
    {
        return true;
    }
    else
    {
        return false;
    }
}
// 设置灯丝预设电流值及限制值，并进行验证
bool ModbusHVPS::setHV_Filament(double filament_pre, double filament_lim)
{
    // 读取用户输入的电压和电流值，转换成高压电源协议要求的格式，并计算Check位
    Fil_Pre = static_cast<int>(filament_pre * 4095 / 10 + 0.5);			    // 将灯丝的设置电流值转为int型，四舍五入，电流最大为10 A
    Fil_Lim = static_cast<int>(filament_lim * 4095 / 10 + 0.5);			    // 将灯丝的极限电流值转为int型，四舍五入，电流最大为10 A

    // 16位简化版累加和校验算法（双字节），计算每个值的二进制形式中数字“1”的总数
    checkSum_filpre = _mm_popcnt_u32(0xA1) + _mm_popcnt_u32(0x63) + _mm_popcnt_u32(Fil_Pre);
    checkSum_fillim = _mm_popcnt_u32(0xA1) + _mm_popcnt_u32(0x64) + _mm_popcnt_u32(Fil_Lim);
 
    // 设置灯丝电流值指令形式为 A1 63 00 00 00 HIBYTE(Fil_Pre) LOBYTE(Fil_Pre) HIBYTE(checkSum_filpre) LOBYTE(checkSum_filpre) 00 0D
    BYTE fil_pre[11] = {
       0xA1, 0x63,0x00, 0x00,0x00,
       HIBYTE(Fil_Pre), LOBYTE(Fil_Pre), HIBYTE(checkSum_filpre), LOBYTE(checkSum_filpre),
       0x00,0x0D
    };

    // 设置灯丝电流极限值指令形式为 A1 64 00 00 00 HIBYTE(Fil_Lim) LOBYTE(Fil_Lim) HIBYTE(checkSum_fillim) LOBYTE(checkSum_fillim) 00 0D
    BYTE fil_lim[11] = {
       0xA1, 0x64,0x00, 0x00,0x00,
       HIBYTE(Fil_Lim), LOBYTE(Fil_Lim), HIBYTE(checkSum_fillim), LOBYTE(checkSum_fillim),
       0x00,0x0D
    };

    // 清空缓存区，发送设置灯丝极限电流值的指令，接收指令，判断指令    
    memset(HV_Recv, 0, sizeof(HV_Recv));                                    // 清空接收缓冲区
    if (!sendHVData(fil_lim, 11))             return false;
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);
    if (HV_RecvLen < 0 || HV_Recv[2] != 10)     return false;               // 返回数据第3个字节为0x0A，说明格式正确

    // 清空缓存区，发送设置灯丝电流值的指令，接收指令，判断指令
    memset(HV_Recv, 0, sizeof(HV_Recv));                                    // 清空接收缓冲区
    if (!sendHVData(fil_pre, 11))             return false;
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);
    if (HV_RecvLen < 0 || HV_Recv[2] != 10)     return false;               // 返回数据第3个字节为0x0A，说明格式正确
   
    return true;                                                            // 返回0表示设置成功
}

// 打开灯丝输出
bool ModbusHVPS::FIL_ON()
{
    BYTE FIL_ON[11] = { 0xA1, 0x70, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x07, 0x00, 0x0D };

    // 发送设置灯丝电流值的指令
    if (!sendHVData(FIL_ON, 11)) return false;
    // 接收信息并进行验证
    memset(HV_Recv, 0, sizeof(HV_Recv));                                        // 清空接收缓冲区
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);

    // 高压电源返回数据：如果长度为11且第6个字节为0x01，说明FIL_ON命令执行成功
    if (HV_RecvLen > 0 && HV_Recv[5] == 1)
    {
        return true;
    }
    else
    {
        return false;
    }
}

// 关闭灯丝输出
bool ModbusHVPS::FIL_OFF() 
{
    BYTE FIL_OFF[11] = { 0xA1, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x0D };

    // 发送设置灯丝电流值的指令
    if (!sendHVData(FIL_OFF, 11)) return false;
    // 接收信息并进行验证
    memset(HV_Recv, 0, sizeof(HV_Recv));                                        // 清空接收缓冲区
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);

    // 高压电源返回数据：如果长度为11且第6个字节为0x00，说明FIL_OFF命令执行成功
    if (HV_RecvLen > 0 && HV_Recv[5] == 0)
    {
        return true;
    }
    else
    {
        return false;
    }
}

// 高压电源状态读取函数，返回实际电压与电流值，需要解码
int ModbusHVPS::readHV_State()
{
    // 读取高压电源状态的指令，格式为 A1 60 00 00 00 00 00 00 05 00 0D
    BYTE ReadHV[11] = { 0xA1, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x0D };
    // 发送读取高压电源状态的指令，返回-100表示发送失败，
    if (!sendHVData(ReadHV, 11)) return -100;
    // 接收信息并进行验证
    memset(HV_Recv, 0, sizeof(HV_Recv));                                        // 清空接收缓冲区
    HV_RecvLen = recv(HVPS_socket, (char*)HV_Recv, sizeof(HV_Recv), 0);

    // 高压电源返回数据：如果长度为19且第2个字节为0x0A，说明返回数据格式正确，
    if (HV_RecvLen == 19 && HV_Recv[2] == 10)
    {
        errCheck = HV_Recv[3] % 8;
        if (errCheck == 0) // 说明电源未出错，正常传递高压电源的电压和电流值信息
        {
            HVPS_noErr = TRUE;
            HVrealVol = (int)((HV_Recv[5] * 256 + HV_Recv[6]) * 50 * 100 / 4095);
            HVrealCur = (int)((HV_Recv[7] * 256 + HV_Recv[8]) * 1000 * 100 / 4095);
            // 将反馈的电压电流值通过消息发送给主线程，更新界面显示
            // 返回值格式为：实际电压值*100 + 实际电流值*100*10000，
            // 这里*100是为了保留两位小数，*10000是为了将电流值放在整数部分的后四位，方便主程序分离电压和电流值
            return (int)(HVrealVol + HVrealCur * 10000);
        }
        else // 说明电源出错，将错误信息传输出去
        {
            // 具体错误通过-1 * HVrecvBuf[3]的值进行判断，错误类型的定义请参考高压电源的通讯协议
            HVPS_noErr = FALSE;
            return -1 * (int)HV_Recv[3];
        }
    }
}
// 读取灯丝的实际电压和电流值，需要解码，（此函数必须紧跟readHV_State()进行使用）
int ModbusHVPS::readFIL_State()
{
    // 高压电源返回数据：如果长度为19且第2个字节为0x0A，说明返回数据格式正确，
    if (HVPS_noErr)
    {       
        FILrealVol = (int)((HV_Recv[11] * 256 + HV_Recv[12]) * 5.5 * 1000 / 4095);
        FILrealCur = (int)((HV_Recv[9] * 256 + HV_Recv[10]) * 3.6 * 1000 / 4095);
        // 将反馈的灯丝电压电流值通过消息发送给主线程，更新界面显示
        // 返回值格式为：实际电压值*1000 + 实际电流值*1000*10000，
        // 这里*100是为了保留两位小数，*10000是为了将电流值放在整数部分的后四位，方便主程序分离电压和电流值
        return (int)(FILrealVol + FILrealCur * 10000);
    }    
}
