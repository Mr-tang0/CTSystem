# CTSystem

CT 设备控制系统 - 基于 Wails 框架的跨平台桌面应用

## 项目简介

CTSystem 是一个用于控制 CT 扫描设备的桌面应用程序，提供图形化界面控制以下设备：

- **探测器**：CT 图像采集设备
- **高压电源**：X 射线高压发生器
- **射线源**：160keV X 射线源
- **位移台**：多轴精密运动平台

## 技术栈

| 组件 | 技术 |
|------|------|
| 后端 | Go 1.23+ |
| 前端 | Vue 3 + TypeScript |
| 框架 | Wails v2.11 |
| 通信协议 | Modbus TCP / 自定义串口协议 |

## 项目结构

```
CTSystem/
├── backend/                 # 后端代码
│   ├── CSDK/               # 探测器SDK封装
│   └── components/         # 设备组件
├── frontend/               # 前端代码
│   └── src/components/     # Vue组件
├── docs/                   # 文档
│   └── OPERATION_GUIDE.md  # 操作手册
├── res/                    # 资源文件
│   └── images/             # 图片资源
├── app.go                  # 应用主结构
├── main.go                 # 程序入口
└── wails.json              # Wails配置
```

## 快速开始

### 运行程序

双击 `CTSystem.exe` 启动程序

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

## 功能特性

### 设备连接

- 支持多设备同时连接
- 自动检测探测器设备
- 实时显示连接状态

### 探测器控制

- 曝光时间设置
- 软件触发采集
- 实时图像显示
- 连续采集模式

### 高压电源控制

- 电压/电流设置
- 高压开关控制
- 灯丝预热控制
- 远程模式切换

### 射线源控制

- kV/mA 参数设置
- 射线开关控制
- 状态监控
- 故障码读取

### 位移台控制

- 多轴点动控制
- 相对/绝对定位
- 回零操作
- 紧急停止

## 文档

- [操作手册](docs/OPERATION_GUIDE.md)
- [API参考](docs/OPERATION_GUIDE.md#api参考)
- [故障排除](docs/OPERATION_GUIDE.md#故障排除)

## 开发环境

### 前置要求

- Go 1.23+
- Node.js 18+
- Wails CLI v2.11+

### 安装 Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 开发命令

```bash
# 查看项目信息
wails doctor

# 开发模式
wails dev

# 构建
wails build

# 构建为单文件
wails build -s
```

## 配置文件

### wails.json

Wails 项目配置文件，包含：
- 应用名称
- 输出目录
- 前端构建配置

### backend/app.json

应用配置文件，包含：
- 版本信息
- 更新地址

## 许可证

MIT License

## 作者

- **tang**
- GitHub: [Mr-tang0](https://github.com/Mr-tang0/CTSystem)

## 更新日志

### v1.0.0 (2026-05-23)

- 初始版本发布
- 支持探测器、高压电源、射线源、位移台设备控制
- 图形化操作界面
- 实时状态监控