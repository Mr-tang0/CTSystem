<template>
    <div class="detector-monitor">
        <!-- <div class="monitor-header">
            <svg class="monitor-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <polyline points="12 6 12 12 16 14"/>
            </svg>
            <span class="monitor-title">探测器监控</span>
        </div> -->
        
        <div class="monitor-content">
            <!-- 当前北京时间 -->
            <div class="monitor-item">
                <span class="item-label">时间</span>
                <span class="item-value time">{{ monitorData.currentTime }}</span>
            </div>

            <!-- <div class="monitor-item">
                <span class="item-label">序列号</span>
                <span class="item-value serial">{{ monitorData.sn }}</span>
            </div> -->
            <!-- <div class="monitor-item">
                <span class="item-label">模式</span>
                <span class="item-value mode">{{ monitorData.mode }}</span>
            </div> -->

            <!-- 当前温度 -->
            <!-- <div class="monitor-item">
                <span class="item-label">温度</span>
                <span class="item-value">{{ monitorData.tempreture.toFixed(2) }}°C</span>
            </div> -->
            
            <!-- 当前湿度 -->
            <!-- <div class="monitor-item">
                <span class="item-label">湿度</span>
                <span class="item-value">{{ monitorData.humidity.toFixed(2) }}%</span>
            </div> -->
            
            <!-- 当前曝光时间 -->
            <div class="monitor-item">
                <span class="item-label">曝光时间</span>
                <span class="item-value">{{ monitorData.exposureTime.toFixed(2) }}ms</span>
            </div>
            
            <!-- 当前拍摄角度 -->
            <div class="monitor-item">
                <span class="item-label">角度</span>
                <span class="item-value angle">{{ monitorData.angle.toFixed(2) }}°</span>
            </div>
            
            <!-- 图像尺寸 -->
            <div class="monitor-item">
                <span class="item-label">图像尺寸</span>
                <span class="item-value size">{{ monitorData.imageWidth }} × {{ monitorData.imageHeight }}</span>
            </div>
            
            <!-- 当前采集状态 -->
            <div class="monitor-item status-item">
                <span class="item-label">状态</span>
                <span class="item-value">{{ monitorData.status }}</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, watch } from 'vue';

const props = defineProps({
    imageWidth: {
        type: Number,
        default: null
    },
    imageHeight: {
        type: Number,
        default: null
    },
    angle: {
        type: Number,
        default: 0.00
    },
    sn: {
        type: String,
        default: ''
    },
    mode: {
        type: String,
        default: 'idle'
    },
    tempreture: {
        type: Number,
        default: 0.00
    },
    humidity: {
        type: Number,
        default: 0.00
    },
    exposureTime: {
        type: Number,
        default: 0.00
    },
    status: {
        type: String,
        default: '未运行'
    },
});

const monitorData = reactive({
    sn: props.sn,//设备序列号
    mode: props.mode,//采集模式
    tempreture: props.tempreture,//温度
    humidity: props.humidity,//湿度
    exposureTime: props.exposureTime,//曝光时间
    currentTime: '',
    angle: props.angle,
    imageWidth: props.imageWidth,
    imageHeight: props.imageHeight,
    status: props.status,
});

// 监听props变化
// watch(() => props.imageWidth, (newVal) => {monitorData.imageWidth = newVal;});
// watch(() => props.imageHeight, (newVal) => {monitorData.imageHeight = newVal;});
// watch(() => props.angle, (newVal) => {monitorData.angle = newVal;});
// watch(() => props.sn, (newVal) => { monitorData.sn = newVal; });
// watch(() => props.mode, (newVal) => { monitorData.mode = newVal; });
// watch(() => props.tempreture, (newVal) => { monitorData.tempreture = newVal; });
// watch(() => props.humidity, (newVal) => { monitorData.humidity = newVal; });
// watch(() => props.exposureTime, (newVal) => { monitorData.exposureTime = newVal; });

let timer = null;
const updateTime = () => {
    const now = new Date();
    monitorData.currentTime = now.toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
    });
};


onMounted(() => {
    updateTime();
    timer = setInterval(() => {updateTime();}, 1000);
});

onUnmounted(() => {
    if (timer) {
        clearInterval(timer);
    }
});
</script>

<style scoped>
.detector-monitor {
    position: absolute;
    right: 16px;
    bottom: 16px;
    width: 220px;
    background: rgba(15, 23, 42, 0.55);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(56, 189, 248, 0.2);
    border-radius: 12px;
    padding: 14px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
    z-index: 10;
}

.monitor-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 12px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.15);
}

.monitor-icon {
    width: 16px;
    height: 16px;
    color: #38bdf8;
}

.monitor-title {
    font-size: 13px;
    font-weight: 600;
    color: #38bdf8;
}

.monitor-content {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.monitor-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.item-label {
    font-size: 11px;
    color: #64748b;
}

.item-value {
    font-size: 12px;
    color: #e2e8f0;
    font-family: 'Courier New', monospace;
}

.item-value.serial {
    color: #94a3b8;
    font-size: 11px;
}

.item-value.time {
    color: #38bdf8;
}

.item-value.angle {
    color: #fbbf24;
}

.item-value.size {
    color: #a78bfa;
}

.status-item .item-value {
    padding: 2px 8px;
    border-radius: 10px;
    font-size: 11px;
    font-weight: 500;
}

.status-idle {
    background: rgba(148, 163, 184, 0.2);
    color: #94a3b8;
}

.status-exposing {
    background: rgba(234, 88, 12, 0.2);
    color: #f97316;
    animation: pulse 1.5s infinite;
}

.status-success {
    background: rgba(16, 185, 129, 0.2);
    color: #10b981;
}

@keyframes pulse {
    0%, 100% {
        opacity: 1;
    }
    50% {
        opacity: 0.6;
    }
}
</style>