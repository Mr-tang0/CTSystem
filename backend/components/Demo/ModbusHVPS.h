/*****************************************************************************/
/***************高压电源网口Modbus通信头文件（作者：ShiPeipei）***************/
/*****************************************************************************/
#ifndef MODBUS_HVPS_CLIENT_H
#define MODBUS_HVPS_CLIENT_H

class ModbusHVPS
{
public:
    ModbusHVPS();
    ~ModbusHVPS();
    
    // 连接高压电源设备：IP、端口;
    bool ConnectHVPS(const char* HVPS_ip, int HVPS_port);
    bool DisconnectHVPS();
    // 远程Remote模式设置
    bool setHV_Remote();
    // 设置高压电源的电压电流值
    bool setHV_VI(double setvoltage_kV, double setcurrent_uA);
    // 设置灯丝预设电流值及限制值
	bool setHV_Filament(double filament_pre, double filament_lim);
	// 打开灯丝输出
	bool FIL_ON();
	// 关闭灯丝输出
	bool FIL_OFF();
    // 打开高压电源输出
    bool HV_ON();
    // 关闭高压电源输出
    bool HV_OFF();
    // 高压电源状态读取函数，返回实际电压与电流值
    int readHV_State();
    // 灯丝状态读取函数，返回实际电压与电流值
    int readFIL_State();      
   
private:    
    SOCKET HVPS_socket;

     // 给高压电源发送数据和接收数据
    bool sendHVData(BYTE* HV_packet, int HV_len);
    bool recvHVData(BYTE* HV_buf, int HV_len);

    BYTE						HV_Recv[1024];			// 高压电源接收数据缓冲区
    int                         HV_RecvLen;             // 高压电源接收数据的长度，用于验证接收数据的完整性

    int							VolValue;				// 电压设定值，转换成协议要求的整数值（0 ~ 4096）
    int							CheckSum_Vol;			// 设置电压值，checksum位
    int							CurValue;				// 电流设定值，转换成协议要求的整数值（0 ~ 4096）
    int							CheckSum_Cur;			// 设置电流值，checksum位

    int                         HVrealVol;				// 高压电源实际电压对应数据
    int                         HVrealCur;				// 高压电源实际电流对应数据

	int						    Fil_Pre;				// 灯丝电流设定值，转换成协议要求的整数值（0 ~ 4095）
    int							checkSum_filpre;		// 灯丝电流设定值，checksum位
    int						    Fil_Lim;				// 灯丝电流设定极限值，转换成协议要求的整数值（0 ~ 4095）
    int							checkSum_fillim;		// 灯丝电流设定极限值，checksum位

    int                         FILrealVol;				// 高压电源实际电压对应数据
    int                         FILrealCur;				// 高压电源实际电流对应数据

    int                         errCheck;				// 校验是否有错误 
	BOOL						HVPS_noErr;		        // 高压电源读取无错误，TRUE —— 无错误；FALSE —— 有错误
};

#endif
