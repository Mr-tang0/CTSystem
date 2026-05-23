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
            <!-- 探测器序列号 -->
            <div class="monitor-item">
                <span class="item-label">序列号</span>
                <span class="item-value serial">{{ monitorData.serialNumber }}</span>
            </div>
            
            <!-- 当前北京时间 -->
            <div class="monitor-item">
                <span class="item-label">时间</span>
                <span class="item-value time">{{ monitorData.currentTime }}</span>
            </div>
            
            <!-- 当前拍摄角度 -->
            <div class="monitor-item">
                <span class="item-label">角度</span>
                <span class="item-value angle">{{ monitorData.angle.toFixed(2) }}°</span>
            </div>
            
            <!-- 图像尺寸 -->
            <div class="monitor-item">
                <span class="item-label">尺寸</span>
                <span class="item-value size">{{ monitorData.imageWidth }} × {{ monitorData.imageHeight }}</span>
            </div>
            
            <!-- 当前采集状态 -->
            <div class="monitor-item status-item">
                <span class="item-label">状态</span>
                <span class="item-value" :class="getStatusClass(monitorData.status)">
                    {{ getStatusText(monitorData.status) }}
                </span>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue';

const monitorData = reactive({
    serialNumber: 'CT-2024-001',
    currentTime: '',
    angle: 0.00,
    imageWidth: 3072,
    imageHeight: 3072,
    status: 'idle' // 'idle', 'exposing', 'success'
});

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

const getStatusClass = (status) => {
    switch (status) {
        case 'exposing':
            return 'status-exposing';
        case 'success':
            return 'status-success';
        case 'idle':
        default:
            return 'status-idle';
    }
};

const getStatusText = (status) => {
    switch (status) {
        case 'exposing':
            return '正在曝光';
        case 'success':
            return '采集成功';
        case 'idle':
        default:
            return '待命中';
    }
};

// 模拟角度变化
const simulateAngleChange = () => {
    monitorData.angle = (monitorData.angle + 0.5) % 360;
};

onMounted(() => {
    updateTime();
    timer = setInterval(() => {
        updateTime();
        simulateAngleChange();
    }, 1000);
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