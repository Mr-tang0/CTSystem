package CSDK

// 所有类型定义
type (
    // 基础类型
    Char       = int8
    Bool       = int32
    Short      = int16
    Int16      = int16
    Int32      = int32
    Long       = int32
    Uchar      = uint8
    Uint8      = uint8
    Uint16     = uint16
    Uint32     = uint32
    Ushort     = uint16
    Ulong      = uint32
    Uint64     = uint64

    // 结构体定义
    TImageMode struct {
        UsRow    uint16
        UsCol    uint16
        UsPix    uint16
        UsType   uint16
    }

    TImageShiftMode struct {
        UsTopRowShift     uint16
        UsLeftColShift    uint16
        UsBottomRowShift  uint16
        UsRightColShift   uint16
    }

    TFpTempHum struct {
        S16Temp   int16
        S16Hum    int16
    }

    TFpBatInfo struct {
        UsRemain  uint16
        UsFull    uint16
    }

    TWifiStatus struct {
        UcSignalLevel   uint8
        UcNoiseLevel    uint8
        UsFrequency     uint16
        UsPowerLevel    uint16
        I32Bitrates     int32
    }

    TFPStat struct {
        TWifiStatus  TWifiStatus
        TFpTempHum   TFpTempHum
        TFpBatInfo   TFpBatInfo1
        TFpBatInfo   TFpBatInfo2
    }

    TFpBatInfoEx struct {
        UsRemain      uint16
        UsFull        uint16
        UsCycles      uint16
        UsSerialNum   uint16
        UcPartNum     [16]uint8
        UcReserve     [16]uint8
    }

    TFPStatex struct {
        TWifiStatus   TWifiStatus
        TFpTempHum    TFpTempHum
        TFpBatInfoEx  TFpBatInfo1Ex
        TFpBatInfoEx  TFpBatInfo2Ex
    }

    TFPUserCfg struct {
        UsRepeatTimes      uint16
        UsCmdDelayTime     uint16
        UsXwinTime         uint16
        UsAutoScrubEnable  uint16
        UsXwinTimeH        uint16
        UsRefreshEnable    uint16
        UsXwinGDelay       uint16
        UsPreProcessing    uint16
        UdRefreshCycle     uint32
    }

    TFPInfo struct {
        UcBoardSn     [32]uint8
        UcMcuVer      [16]uint8
        UcFpgaVer     [16]uint8
        UcSwVer       [16]uint8
        UdUsage       uint32
        UdFreeFallTimes  uint32
        UdExposureDose   uint32
        UdFpIp        uint32
        UcFpMac       [6]uint8
        UcConMod      uint8
    }

    TRBConf struct {
        UsTubeReady  uint16
        UsFpAck      uint16
        UsXwin       uint16
        UsExposure   uint16
        UsHsMode     uint16
        UsPort       uint16
        UsXwinH      uint16
    }

    TRBInfo struct {
        UcRBInfo [17]uint8
    }

    TFpgaFullCfg struct {
        TFPUserCfg    TFPUserCfg
        AwFpgaReg     [256]uint16
        AwFpgaReg1    [256]uint16
        AwFpgaReg2    [256]uint16
        AwFpgaReg3    [256]uint16
        AwFpgaReg4    [256]uint16
    }

    TComFpNode struct {
        FPPsn    [32]int8
        FPIP     uint32
        Connect  int8
        Opened   int8
    }

    TComFpList struct {
        TFpNode    [16]TComFpNode
        Ncount     int8
    }

    TWifiConf struct {
        Essid    [64]int8
        Key      [64]int8
        Channel  [64]int8
    }

    TLicenseInfo struct {
        License      int8
        RemainDays   int8
        Permission   int8
    }

    TImageSize struct {
        UsRow    uint16
        UsCol    uint16
    }

    TCBConf struct {
        UsTubeReadyTime  uint16
        UsExposureTime   uint16
        UsPort           uint16
    }

    TFPBaseInfo struct {
        CFpSn          [32]int8
        UcFpIp         [4]uint8
        UcFpMac        [6]uint8
        UdUsage        uint32
        UdExposureDose uint32
        UsFreeFallTimes uint16
    }

    TQuaternionVectorData struct {
        I32W   int32
        I32X   int32
        I32Y   int32
        I32Z   int32
    }

    TEulerVectorData struct {
        I32H   int32
        I32R   int32
        I32P   int32
    }

    TAngularVelocityData struct {
        I32X   int32
        I32Y   int32
        I32Z   int32
    }

    TGravityVectorData struct {
        I32X   int32
        I32Y   int32
        I32Z   int32
    }

    TMotionFeatures struct {
        TQuaternionVectorData  TQuaternionVectorData
        TEulerVectorData       TEulerVectorData
        TAngularVelocityData   TAngularVelocityData
        TGravityVectorData     TGravityVectorData
    }

    TShockDate struct {
        UsYear        uint16
        UcMon         uint8
        UcDay         uint8
        UcHour        uint8
        UcMin         uint8
        UcShockLevel  uint8
    }

    TShockInfo struct {
        UsShockCount   uint16
        DateInfo       [50]TShockDate
    }

    TMetaData struct {
        UcMeta [40]uint8
    }

    TDhcpCfg struct {
        UcIsOn      uint8
        UdIpStr     uint32
        UdIpEnd     uint32
    }
)

// 常量定义
const (
    CONNECT_DIR   = int8(0)
    CONNECT_PC    = int8(1)
    CONNECT_FP1   = int8(2)
    CONNECT_FP2   = int8(3)
    CONNECT_WIFI  = int8(4)
    CONNECT_WIRE  = int8(5)
    CONNECT_NONE  = int8(0xFF)

    FP_OPEND    = int8(1)
    FP_CLOSED   = int8(0)

    EVENT_LINKUP       = int8(1)
    EVENT_LINKDOWN     = int8(2)
    EVENT_BUSY         = int8(3)
    EVENT_CMDSTART     = int8(4)
    EVENT_IMAGEVALID   = int8(5)
    EVENT_CMDEND       = int8(6)
    EVENT_READY        = int8(7)
    EVENT_EXPOSE       = int8(8)
    EVENT_EXPEND       = int8(9)
    EVENT_AED_A1       = int8(10)
    EVENT_AED_A2       = int8(11)
    EVENT_AEC_STOP     = int8(12)
    EVENT_OFFSETDONE   = int8(13)
    EVENT_XWINEND      = int8(14)
    EVENT_AED_PREP_DONE = int8(15)
    EVENT_INFO_MODIFIED = int8(16)
    EVENT_LICENSEOK    = int8(17)
    EVENT_LICENSESNERR = int8(18)
    EVENT_LICENSETYPEERR = int8(19)
    EVENT_HEARTBEAT    = int8(20)
    EVENT_CALINTERRUPT = int8(21)
    EVENT_LINKUPEX     = int8(26)
    EVENT_LINKDOWNEX   = int8(27)
    EVENT_HEARTBEATEX  = int8(28)
    EVENT_UPLOADFILESTEP = int8(30)
    EVENT_IMAGESTART   = int8(50)
    EVENT_IMAGEEND     = int8(51)
    EVENT_BATLOW1      = int8(52)
    EVENT_BATLOW2      = int8(53)
    EVENT_NETWORKBUSY  = int8(54)
    EVENT_ERRDDRFULL   = int8(98)
    EVENT_IMAGEFRAMELOST = int8(99)
    EVENT_TRIGERR      = int8(100)

    BINNING_1x1 = int8(0)
    BINNING_2x2 = int8(1)
    BINNING_3x3 = int8(2)
    BINNING_4x4 = int8(3)
    BINNING_6x6 = int8(4)
    BINNING_8x8 = int8(5)

    ACQ_SCAN        = int8(2)
    ACQ_READ_OFFSET = int8(3)
    ACQ_READ_OFFSET_READY = int8(4)
    ACQ_READ_IMAGE  = int8(6)
    ACQ_READ_IMAGE_READY = int8(7)

    FP_TPL_1x1 = int8(0x01)
    FP_TPL_2x2 = int8(0x02)
    FP_TPL_3x3 = int8(0x03)
    FP_TPL_4x4 = int8(0x04)
    FP_TPL_6x6 = int8(0x05)
    FP_TPL_8x8 = int8(0x06)

    IMG_CALIB_RAW    = int8(0x00)
    IMG_CALIB_OFFSET = int8(0x01)
    IMG_CALIB_GAIN   = int8(0x02)
    IMG_CALIB_DEFECT = int8(0x04)

    STATUS_NULL    = int8(0)
    STATUS_IDLE    = int8(0x01)
    STATUS_HST     = int8(0x02)
    STATUS_AED1    = int8(0x03)
    STATUS_AED2    = int8(0x04)
    STATUS_RECOVER = int8(0x05)
    STATUS_OLAED1  = int8(0x06)
    STATUS_OLAED2  = int8(0x07)
    STATUS_CBCT    = int8(0x08)
    STATUS_DST     = int8(0x09)
    STATUS_CBCT2   = int8(0x0A)

    COM_SUCCESS     = int32(1000)
    COM_INIT_FAIL   = int32(1001)
    COM_NO_TPL      = int32(1002)
    COM_FP_BUSY     = int32(1003)
    COM_LK_BREAK    = int32(1004)
    COM_LK_ERR      = int32(1005)
    COM_TIMEOUT     = int32(1006)
    COM_TP_ALARM    = int32(1007)

    FP_TYPE_ERR   = int8(10)
    FP_TYPE_3543  = int8(11)
    FP_TYPE_4343  = int8(12)
    FP_TYPE_1723  = int8(13)
    FP_TYPE_3030  = int8(14)
    FP_TYPE_2925  = int8(15)
    FP_TYPE_A843  = int8(16)
    FP_TYPE_1613  = int8(17)
    FP_TYPE_3543B = int8(18)
    FP_TYPE_4343ZF = int8(19)
    FP_TYPE_3543ZF = int8(20)
    FP_TYPE_3025ZF = int8(21)
    FP_TYPE_6543  = int8(22)
    FP_TYPE_2121  = int8(23)
    FP_TYPE_6561  = int8(24)
    FP_TYPE_4386  = int8(25)
    FP_TYPE_4365  = int8(26)
    FP_TYPE_6557  = int8(27)
    FP_TYPE_4314  = int8(28)
    FP_TYPE_1917  = int8(29)
    FP_TYPE_2929  = int8(30)
    FP_TYPE_2417  = int8(31)
    FP_TYPE_3543ZB = int8(32)
    FP_TYPE_2417Z = int8(33)
    FP_TYPE_1313  = int8(34)
    FP_TYPE_C243  = int8(35)

    FP_COMPATIBLE_VER = int8(1)
    ZM_COMPATIBLE_VER = int8(2)
    ERR_COMPATIBLE_VER = int8(0xFF)

    MAX_CONF_SIZE = 256
    META_DATA_LEN = 40
    SHOCK_REC_NUM = 50

    AEC_NUM_1 = int8(0x01)
    AEC_NUM_2 = int8(0x02)
    AEC_NUM_3 = int8(0x04)
    AEC_NUM_4 = int8(0x08)
    AEC_NUM_5 = int8(0x10)
)
