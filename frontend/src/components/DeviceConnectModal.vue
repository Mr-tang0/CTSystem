<template>
    <teleport to="body">
        <transition name="modal">
            <div v-if="visible" class="modal-overlay">
                <div class="modal-container">
                    <div class="modal-header">
                        <span class="modal-title">设备连接</span>
                        <button class="close-btn" @click="handleClose">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    
                    <div class="modal-content">
                        <!-- 位移台 -->
                        <div class="device-item">
                            <label class="device-label">位移台</label>
                            <div class="device-input-group">
                                <input 
                                    v-model="devices.stage.ip" 
                                    type="text" 
                                    class="device-input" 
                                    placeholder="输入位移台IP地址"
                                />
                                <div class="device-buttons">
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.stage.connected }"
                                        @click="toggleDevice('stage')"
                                    >
                                        {{ devices.stage.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.stage.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.stage.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>

                        <!-- 高压电源 -->
                        <div class="device-item">
                            <label class="device-label">高压电源</label>
                            <div class="device-input-group">
                                <input 
                                    v-model="devices.power.ip" 
                                    type="text" 
                                    class="device-input" 
                                    placeholder="输入高压电源IP地址"
                                />
                                <div class="device-buttons">
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.power.connected }"
                                        @click="toggleDevice('power')"
                                    >
                                        {{ devices.power.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.power.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.power.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>

                        <!-- 放射源 -->
                        <div class="device-item">
                            <label class="device-label">电子源</label>
                            <div class="device-input-group">
                                <input 
                                    v-model="devices.source.ip" 
                                    type="text" 
                                    class="device-input" 
                                    placeholder="输入放射源IP地址"
                                />
                                <div class="device-buttons">
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.source.connected }"
                                        @click="toggleDevice('source')"
                                    >
                                        {{ devices.source.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.source.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.source.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>

                        <!-- 探测器 -->
                        <div class="device-item">
                            <label class="device-label">探测器</label>
                            <div class="device-input-group">
                                <select v-model="devices.detector.selected" class="device-select">
                                    <option value="">请选择探测器</option>
                                    <option v-for="det in detectorList" :key="det.id" :value="det.id">
                                        {{ det.name }} ({{ det.serial }})
                                    </option>
                                </select>
                                <div class="device-buttons">
                                    <button class="refresh-btn" @click="refreshDetectors">
                                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                            <polyline points="23 4 23 10 17 10"/>
                                            <polyline points="1 20 1 14 7 14"/>
                                            <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/>
                                        </svg>
                                    </button>
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.detector.connected }"
                                        @click="toggleDevice('detector')"
                                    >
                                        {{ devices.detector.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.detector.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.detector.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>
                    </div>

                    <div class="modal-footer">
                        <button class="action-btn primary-btn" @click="connectAll">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                            </svg>
                            一键连接
                        </button>
                        <button class="action-btn danger-btn" @click="disconnectAll">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                                <line x1="12" y1="19" x2="12" y2="10"/>
                            </svg>
                            断开全部
                        </button>
                    </div>
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { reactive, ref } from 'vue';

defineProps({
    visible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['close']);

const devices = reactive({
    stage: { ip: '192.168.11.6', connected: false },
    power: { ip: '192.168.11.4', connected: false },
    source: { ip: '192.168.11.5', connected: false },
    detector: { selected: '', connected: false }
});

const detectorList = ref([
    { id: 'det-1', name: 'CT-Detector-A', serial: 'SN-2024-001' },
    { id: 'det-2', name: 'CT-Detector-B', serial: 'SN-2024-002' },
    { id: 'det-3', name: 'CT-Detector-C', serial: 'SN-2024-003' }
]);

const handleClose = () => {
    emit('close');
};

const toggleDevice = (deviceType) => {
    const device = devices[deviceType];
    
    if (deviceType === 'detector' && !device.selected) {
        console.log('请先选择探测器');
        return;
    }
    
    if (deviceType !== 'detector' && !device.ip) {
        console.log('请输入IP地址');
        return;
    }
    
    device.connected = !device.connected;
    console.log(`${deviceType} ${device.connected ? '连接成功' : '已断开'}`);
};

const refreshDetectors = () => {
    console.log('刷新探测器列表');
    // 模拟刷新探测器列表
};

const connectAll = () => {
    devices.stage.connected = devices.stage.ip ? true : devices.stage.connected;
    devices.power.connected = devices.power.ip ? true : devices.power.connected;
    devices.source.connected = devices.source.ip ? true : devices.source.connected;
    devices.detector.connected = devices.detector.selected ? true : devices.detector.connected;
    console.log('一键连接所有设备');
};

const disconnectAll = () => {
    devices.stage.connected = false;
    devices.power.connected = false;
    devices.source.connected = false;
    devices.detector.connected = false;
    console.log('断开所有设备连接');
};
</script>

<style scoped>
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-container {
    width: 480px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.98) 0%, rgba(30, 41, 59, 0.98) 100%);
    border-radius: 16px;
    border: 1px solid rgba(56, 189, 248, 0.2);
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
    overflow: hidden;
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 20px;
    background: rgba(56, 189, 248, 0.1);
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
}

.modal-title {
    font-size: 16px;
    font-weight: 600;
    color: #38bdf8;
}

.close-btn {
    width: 28px;
    height: 28px;
    border: none;
    background: rgba(239, 68, 68, 0.2);
    border-radius: 6px;
    color: #ef4444;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
}

.close-btn:hover {
    background: rgba(239, 68, 68, 0.3);
    transform: rotate(90deg);
}

.modal-content {
    padding: 20px;
}

.device-item {
    margin-bottom: 20px;
}

.device-item:last-child {
    margin-bottom: 0;
}

.device-label {
    display: block;
    font-size: 12px;
    color: #94a3b8;
    margin-bottom: 8px;
    font-weight: 500;
}

.device-input-group {
    display: flex;
    align-items: center;
    gap: 10px;
}

.device-input {
    flex: 1;
    height: 40px;
    padding: 0 14px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 8px;
    color: #ffffff;
    font-size: 14px;
    outline: none;
    transition: all 0.2s ease;
}

.device-input:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.3);
}

.device-input::placeholder {
    color: #64748b;
}

.device-select {
    flex: 1;
    height: 40px;
    padding: 0 14px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 8px;
    color: #ffffff;
    font-size: 14px;
    outline: none;
    cursor: pointer;
    transition: all 0.2s ease;
    appearance: none;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
    background-repeat: no-repeat;
    background-position: right 12px center;
    background-size: 16px;
}

.device-select:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.3);
}

.device-select option {
    background: #1e293b;
    color: #ffffff;
}

.device-buttons {
    display: flex;
    gap: 6px;
}

.conn-btn {
    padding: 10px 16px;
    background: rgba(56, 189, 248, 0.2);
    border: 1px solid rgba(56, 189, 248, 0.4);
    border-radius: 6px;
    color: #38bdf8;
    font-size: 12px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
}

.conn-btn:hover {
    background: rgba(56, 189, 248, 0.3);
}

.conn-btn.connected {
    background: rgba(16, 185, 129, 0.2);
    border-color: rgba(16, 185, 129, 0.4);
    color: #10b981;
}

.conn-btn.connected:hover {
    background: rgba(16, 185, 129, 0.3);
}

.refresh-btn {
    width: 40px;
    height: 40px;
    background: rgba(148, 163, 184, 0.2);
    border: 1px solid rgba(148, 163, 184, 0.3);
    border-radius: 6px;
    color: #94a3b8;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
}

.refresh-btn:hover {
    background: rgba(148, 163, 184, 0.3);
    color: #ffffff;
}

.refresh-btn svg {
    width: 16px;
    height: 16px;
}

.device-status {
    display: flex;
    align-items: center;
    gap: 6px;
    margin-top: 8px;
    padding-left: 4px;
}

.status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: #64748b;
    transition: all 0.2s ease;
}

.device-status.connected .status-dot {
    background: #10b981;
    box-shadow: 0 0 8px rgba(16, 185, 129, 0.5);
}

.status-text {
    font-size: 11px;
    color: #64748b;
}

.device-status.connected .status-text {
    color: #10b981;
}

.modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    padding: 16px 20px;
    background: rgba(0, 0, 0, 0.2);
    border-top: 1px solid rgba(56, 189, 248, 0.1);
}

.action-btn {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 10px 20px;
    border: none;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
}

.action-btn svg {
    width: 16px;
    height: 16px;
}

.primary-btn {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
}

.primary-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 15px rgba(56, 189, 248, 0.4);
}

.danger-btn {
    background: rgba(239, 68, 68, 0.2);
    border: 1px solid rgba(239, 68, 68, 0.4);
    color: #ef4444;
}

.danger-btn:hover {
    background: rgba(239, 68, 68, 0.3);
}

/* 过渡动画 */
.modal-enter-active,
.modal-leave-active {
    transition: all 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
    transform: scale(0.95) translateY(-20px);
}
</style>