# CTSystem 操作文档

## 目录

- [项目简介](#项目简介)
- [系统要求](#系统要求)
- [快速开始](#快速开始)
- [设备连接](#设备连接)
  - [探测器连接](#探测器连接)
  - [高压电源连接](#高压电源连接)
  - [射线源连接](#射线源连接)
  - [位移台连接](#位移台连接)
- [操作指南](#操作指南)
  - [探测器操作](#探测器操作)
  - [高压电源操作](#高压电源操作)
  - [射线源操作](#射线源操作)
  - [位移台操作](#位移台操作)
- [API参考](#api参考)
- [故障排除](#故障排除)
- [附录](#附录)

---

## 项目简介

CTSystem 是一个基于 Wails 框架开发的 CT 设备控制系统，提供图形化界面控制探测器、高压电源、射线源和位移台等设备。

### 技术栈

| 组件 | 技术 |
|------|------|
| 后端 | Go 1.23+ |
| 前端 | Vue 3 + TypeScript |
| 框架 | Wails v2 |
| 通信协议 | Modbus TCP / 自定义串口协议 |

### 项目结构

```
CTSystem/
├── backend/                 # 后端代码
│   ├── CSDK/               # 探测器SDK封装
│   │   ├── SDK/            # C SDK动态库
│   │   ├── NetComGo.go     # Go封装层
│   │   └── types.go        # 类型定义
│   └── components/         # 设备组件
│       ├── CTDevice.go     # 探测器设备
│       ├── HVPSDevice.go   # 高压电源设备
│       ├── RaySource160keV.go  # 射线源设备
│       ├── MotorDevice.go  # 位移台设备
│       └── xinjieModbus.go # Modbus通信客户端
├── frontend/               # 前端代码
│   └── src/components/     # Vue组件
├── res/                    # 资源文件
├── app.go                  # 应用主结构
├── main.go                 # 程序入口
└── wails.json              # Wails配置
```

---

## 系统要求

### 硬件要求

| 项目 | 最低配置 | 推荐配置 |
|------|---------|---------|
| CPU | 双核 2.0GHz | 四核 3.0GHz |
| 内存 | 4GB | 8GB+ |
| 存储 | 500MB | 1GB |
| 网络 | 100Mbps | 1000Mbps |

### 软件要求

- 操作系统：Windows 10/11 (64位)
- 运行时：无需额外安装，程序自带运行时

### 开发环境要求

- Go 1.23+
- Node.js 18+
- Wails CLI v2.11+

---

## 快速开始

### 安装运行

1. 双击运行 `CTSystem.exe`
2. 等待程序启动完成
3. 进入主界面

### 开发模式

```bash
# 安装依赖
go mod tidy
cd frontend && npm install

# 开发模式运行
wails dev

# 构建生产版本
wails build
```

---

## 设备连接

### 探测器连接

![探测器连接界面](res/images/detector_connect.png)

1. 点击左侧 **设备** 按钮，打开设备连接面板
2. 在探测器区域点击 **刷新** 按钮获取设备列表
3. 从下拉框选择探测器设备
4. 点击 **连接** 按钮建立连接
5. 连接成功后状态指示灯变绿

**参数说明：**

| 参数 | 说明 | 示例 |
|------|------|------|
| 设备列表 | 系统检测到的探测器 | Detector_001 |

### 高压电源连接

![高压电源连接](res/images/hvps_connect.png)

1. 在设备面板中找到高压电源区域
2. 输入高压电源 IP 地址
3. 点击 **连接** 按钮
4. 等待连接成功提示

**参数说明：**

| 参数 | 说明 | 默认值 |
|------|------|--------|
| IP地址 | 高压电源网络地址 | 192.168.1.100 |
| 端口 | Modbus端口 | 502 |

### 射线源连接

![射线源连接](res/images/raysource_connect.png)

1. 在设备面板中找到射线源区域
2. 选择串口号（如 COM1, COM3）
3. 点击 **连接** 按钮

**参数说明：**

| 参数 | 说明 | 默认值 |
|------|------|--------|
| 串口 | 通信串口号 | COM1 |
| 波特率 | 通信速率 | 9600 |

### 位移台连接

![位移台连接](res/images/motor_connect.png)

1. 在设备面板中找到位移台区域
2. 输入 PLC IP 地址
3. 点击 **连接** 按钮

**参数说明：**

| 参数 | 说明 | 默认值 |
|------|------|--------|
| IP地址 | PLC网络地址 | 192.168.1.10 |
| 端口 | Modbus端口 | 502 |

---

## 操作指南

### 探测器操作

#### 图像采集

![探测器面板](res/images/detector_panel.png)

1. **曝光时间设置**
   - 在左侧控制面板设置曝光时间（单位：ms）
   - 范围：1-10000ms

2. **软件触发**
   - 点击 **采集** 按钮触发一次图像采集
   - 图像将显示在中央区域

3. **连续采集**
   - 勾选 **连续采集** 选项
   - 设置采集间隔时间
   - 点击开始连续采集

#### 参数监控

右侧监控面板显示实时参数：

| 参数 | 说明 |
|------|------|
| 序列号 | 探测器序列号 |
| 当前时间 | 北京时间 |
| 拍摄角度 | 当前旋转角度 |
| 图像尺寸 | 采集图像尺寸 |
| 采集状态 | 曝光中/采集成功/待命中 |

### 高压电源操作

#### 电压电流设置

1. 在控制面板输入目标电压（kV）
2. 输入目标电流（uA）
3. 点击 **设置** 按钮

**安全提示：**
- 电压范围：0-50 kV
- 电流范围：0-1000 uA
- 请确保灯丝已预热后再开启高压

#### 高压开关操作

```
操作流程：
1. 连接设备 → 2. 设置远程模式 → 3. 设置电压电流 → 4. 开启灯丝 → 5. 开启高压
```

| 操作 | 按钮 | 说明 |
|------|------|------|
| 开启高压 | HV ON | 输出高压 |
| 关闭高压 | HV OFF | 停止高压输出 |
| 开启灯丝 | FIL ON | 开启灯丝预热 |
| 关闭灯丝 | FIL OFF | 关闭灯丝 |

### 射线源操作

#### 基本控制

| 功能 | 操作 | 说明 |
|------|------|------|
| 设置电压 | 输入kV值后点击设置 | 范围：0-160kV |
| 设置电流 | 输入uA值后点击设置 | 范围：0-1000uA |
| 开启射线 | 点击"开启"按钮 | 射线源开始工作 |
| 关闭射线 | 点击"关闭"按钮 | 射线源停止工作 |

#### 状态监控

- **当前kV**：实时电压值
- **当前mA**：实时电流值
- **油温**：设备油温（℃）
- **故障码**：设备故障信息

### 位移台操作

#### 点动控制

![位移台控制](res/images/motor_control.png)

1. 选择运动轴（X/Y/Z/R）
2. 按住方向按钮进行点动
3. 松开按钮停止运动

#### 定位控制

| 功能 | 操作 |
|------|------|
| 相对运动 | 输入距离后点击移动 |
| 绝对定位 | 输入目标位置后点击定位 |
| 回零 | 点击回零按钮回到原点 |
| 停止 | 点击停止按钮紧急停止 |

**轴说明：**

| 轴 | 方向 | 行程 |
|----|------|------|
| X | 左右 | 0-500mm |
| Y | 前后 | 0-500mm |
| Z | 上下 | 0-300mm |
| R | 旋转 | 0-360° |

---

## API参考

### 探测器API

```typescript
// 获取探测器列表
DetectorGetList(): string[]

// 连接探测器
DetectorConnect(index: number): string

// 断开探测器
DetectorDisconnect(): string

// 软件触发
DetectorSoftwareTrigger(): void

// 获取连接状态
DetectorIsConnected(): boolean
```

### 高压电源API

```typescript
// 连接
HighVoltageConnect(ip: string): string

// 断开
HighVoltageDisconnect(): string

// 设置电压
HighVoltageSetVoltage(voltage: number): string

// 设置电流
HighVoltageSetCurrent(current: number): string

// 设置电压电流
HighVoltageSetVI(voltage: number, current: number): string

// 开启高压
HighVoltageEnable(): string

// 关闭高压
HighVoltageDisable(): string

// 设置远程模式
HighVoltageSetRemote(): string

// 读取状态
HighVoltageReadState(): string

// 灯丝控制
HighVoltageFilamentOn(): string
HighVoltageFilamentOff(): string
HighVoltageSetFilament(pre: number, lim: number): string
```

### 射线源API

```typescript
// 连接
RaySourceConnect(com: string): string

// 断开
RaySourceDisconnect(): void

// 设置参数
RaySourceSetKV(kV: number): string
RaySourceSetMA(mA: number): string

// 开关控制
RaySourceEnable(): string
RaySourceDisable(): string

// 状态读取
RaySourceReadStatus(): string
RaySourceReadFaultCode(): string
RaySourceReadCurrentKV(): string
RaySourceReadCurrentMA(): string
RaySourceReadOilTemp(): string
RaySourceReadVersion(): string

// 其他
RaySourceFeedWatchDog(): string
RaySourceResetFaultFlag(): string
RaySourceIsConnected(): boolean
```

### 位移台API

```typescript
// 连接
StageOpenDevice(ip: string): string

// 断开
StageCloseDevice(): void

// 运动
StageAxisPulse(axis: string, dir: boolean): void
StageMoveRel(axis: string, distance: number): void
StageStop(axis: string): void
StageAxisAbs(axis: string): void
```

---

## 故障排除

### 常见问题

#### 探测器无法连接

| 问题 | 解决方案 |
|------|---------|
| 设备列表为空 | 检查USB连接，重新插拔设备 |
| 连接超时 | 检查设备驱动是否正确安装 |
| 连接后无图像 | 检查曝光时间设置是否正确 |

#### 高压电源问题

| 错误码 | 含义 | 解决方案 |
|--------|------|---------|
| B0 | 闭锁保护 | 检查安全联锁装置 |
| B2 | kV过压保护 | 降低电压设置值 |
| B6 | 温度保护 | 等待设备冷却 |
| B11 | 看门狗溢出 | 重启设备并检查通信 |

#### 射线源故障码

| 故障位 | 含义 | 处理方法 |
|--------|------|---------|
| FAULT_LOCK | 闭锁保护 | 检查安全门状态 |
| FAULT_KV_OVER | 过压保护 | 降低电压 |
| FAULT_TEMP | 温度保护 | 检查散热系统 |
| FAULT_MA_OVER | 过流保护 | 降低电流 |

#### 位移台问题

| 问题 | 可能原因 | 解决方案 |
|------|---------|---------|
| 无法连接 | 网络不通 | 检查网线和IP设置 |
| 运动异常 | 参数错误 | 检查运动参数范围 |
| 不响应 | PLC故障 | 重启PLC |

### 日志查看

程序日志保存在：
```
%APPDATA%/CTSystem/logs/
```

---

## 附录

### 设备参数表

#### 高压电源参数

| 参数 | 范围 | 单位 | 精度 |
|------|------|------|------|
| 电压 | 0-50 | kV | 0.1 |
| 电流 | 0-1000 | uA | 1 |
| 灯丝电流 | 0-10 | A | 0.1 |

#### 射线源参数

| 参数 | 范围 | 单位 |
|------|------|------|
| 电压 | 0-160 | kV |
| 电流 | 0-1000 | uA |
| 油温 | -30-100 | ℃ |

### 通信协议

#### 高压电源协议

- 类型：自定义TCP协议
- 端口：502
- 帧格式：`[地址][命令][数据][校验][结束符]`

#### 射线源协议

- 类型：自定义串口协议
- 波特率：9600
- 数据位：8
- 校验：无
- 停止位：1
- 帧格式：`STX + 命令码 + 分隔符 + 参数 + 分隔符 + XOR校验 + ETX`

#### 位移台协议

- 类型：Modbus TCP
- 端口：502
- 从站ID：1

### 更新日志

| 版本 | 日期 | 更新内容 |
|------|------|---------|
| v1.0.0 | 2026-05-23 | 初始版本发布 |

---

## 联系支持

- **项目地址**：https://github.com/Mr-tang0/CTSystem
- **问题反馈**：https://github.com/Mr-tang0/CTSystem/issues

---

*文档版本：v1.0.0*
*最后更新：2026-05-23*