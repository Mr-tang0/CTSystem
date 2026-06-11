<template>
    <div class="detector-panel" >
        <div class="panel-header">
            <span class="panel-title">探测器参数</span>
            <!-- <button class="close-btn" @click="$emit('close')">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"/>
                    <line x1="6" y1="6" x2="18" y2="18"/>
                </svg>
            </button> -->
        </div>
        
        <div class="panel-content">
            <!-- 曝光时间 -->
            <div class="param-group">
                <label class="param-label">探测器曝光时间</label>
                <div class="param-input-wrap">
                    <input 
                        type="number" 
                        v-model.number="exposureTime" 
                        class="param-input"
                        min="1" 
                        max="1000"
                        placeholder="曝光时间(ms)"
                    />
                    <span class="param-unit">ms</span>
                </div>
                <div class="param-slider">
                    <input type="range" v-model.number="exposureTime" min="1" max="50000" />
                </div>
            </div>

            <!-- Binning模式 -->
            <div class="param-group">
                <label class="param-label">探测器Binning</label>
                <div class="binning-options">
                    <button 
                        v-for="bin in binningOptions" 
                        :key="bin.value"
                        class="binning-btn"
                            
                        :class="{ active: binningMode === bin.value }"
                        @click="binningMode = bin.value"
                    >
                        {{ bin.label }}
                    </button>
                </div>
            </div>

            <!-- 增益 -->
            <div class="param-group"> 
                <label class="param-label">探测器增益</label>
            </div>

            <!-- 管电压 -->
            <!-- <div class="param-group">
                <label class="param-label">管电压</label>
                <div class="param-input-wrap">
                    <input 
                        type="number" 
                        v-model.number="tubeVoltage" 
                        class="param-input"
                        min="40" 
                        max="150"
                        placeholder="管电压(kV)"
                    />
                    <span class="param-unit">kV</span>
                </div>
                <div class="param-slider">
                    <input type="range" v-model.number="tubeVoltage" min="40" max="150" />
                </div>
            </div> -->

            <!-- 管电流 -->
            <!-- <div class="param-group">
                <label class="param-label">管电流</label>
                <div class="param-input-wrap">
                    <input 
                        type="number" 
                        v-model.number="tubeCurrent" 
                        class="param-input"
                        min="10" 
                        max="500"
                        placeholder="管电流(mA)"
                    />
                    <span class="param-unit">mA</span>
                </div>
                <div class="param-slider">
                    <input type="range" v-model.number="tubeCurrent" min="10" max="500" />
                </div>
            </div> -->

            

            <!-- 帧率设置 -->
            <!-- <div class="param-group">
                <label class="param-label">帧率</label>
                <div class="param-input-wrap">
                    <input 
                        type="number" 
                        v-model.number="frameRate" 
                        class="param-input"
                        min="1" 
                        max="60"
                        placeholder="帧率(fps)"
                    />
                    <span class="param-unit">fps</span>
                </div>
            </div> -->

            <!-- 图像尺寸 -->
            <!-- <div class="param-group">
                <label class="param-label">图像尺寸</label>
                <div class="size-info">
                    <span>{{ imageWidth }} × {{ imageHeight }}</span>
                </div>
            </div> -->

            <!-- 应用按钮 -->
            <div class="action-buttons">
                <button class="get-btn" @click="getParams">
                    获取参数
                </button>
                <button class="apply-btn" @click="applyParams">
                    应用参数
                </button>
                <!-- <button class="reset-btn" @click="resetParams">
                    重置
                </button> -->
            </div>
            
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { 
    DetectorSetExposeTime, 
    DetectorGetExposeTime,
    DetectorSetBinning,
    DetectorGetBinning
} from '../../wailsjs/go/main/App'



const emit = defineEmits(['close', 'apply']);

// 参数状态
const exposureTime = ref(50);
const binningMode = ref('1×1');
const frameRate = ref(30);



const binningOptions = [
    { label: '1×1', value: '1×1' },
    { label: '2×2', value: '2×2' },
    // { label: '3×3', value: '3×3' },
    // { label: '4×4', value: '4×4' },
    // { label: '6×6', value: '6×6' },
    // { label: '8×8', value: '8×8' }
];

const getParams = async() => {
    console.log('getParams');
    try {
        const exposeTime = await DetectorGetExposeTime();
        console.log(exposeTime);
        exposureTime.value = exposeTime;
    
        const binning = await DetectorGetBinning();
        binningMode.value = binning;
        console.log('binningMode.value:', binningMode.value);
        console.log('binningMode.value == "2×2":', binningMode.value == '2×2');
    } catch (error) {
        console.error('获取参数失败:', error);
        return;
    }
}

// 应用参数
const applyParams = async () => {
    console.log('applyParams');
    try {
        await DetectorSetExposeTime(exposureTime.value);

        await DetectorSetBinning(binningMode.value);
    } catch (error) {
        console.error('应用参数失败:', error);
        return;
    }
    emit('apply', params);
};

defineExpose({
    getParams
});


</script>

<style scoped>
.detector-panel {
    width: 280px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.95) 100%);
    border-radius: 0 12px 12px 0;
    box-shadow: 4px 0 20px rgba(0, 0, 0, 0.3);
    display: flex;
    flex-direction: column;
    height: 100%;
    backdrop-filter: blur(10px);
    border-left: 1px solid rgba(56, 189, 248, 0.2);
}

.panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 20px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
    background: rgba(56, 189, 248, 0.1);
}

.panel-title {
    font-size: 14px;
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
    transition: all 0.2s ease;
}

.close-btn:hover {
    background: rgba(239, 68, 68, 0.3);
    transform: rotate(90deg);
}

.panel-content {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
}

.param-group {
    margin-bottom: 20px;
}

.param-label {
    display: block;
    font-size: 12px;
    color: #94a3b8;
    margin-bottom: 8px;
    font-weight: 500;
}

.param-input-wrap {
    display: flex;
    align-items: center;
    gap: 8px;
}

.param-input {
    flex: 1;
    height: 36px;
    padding: 0 12px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 8px;
    color: #ffffff;
    font-size: 14px;
    outline: none;
    transition: all 0.2s ease;
}

.param-input:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.3);
}

.param-unit {
    font-size: 12px;
    color: #64748b;
    min-width: 30px;
}

.param-slider {
    margin-top: 8px;
}

.param-slider input[type="range"] {
    width: 100%;
    height: 4px;
    -webkit-appearance: none;
    appearance: none;
    background: rgba(56, 189, 248, 0.2);
    border-radius: 2px;
    outline: none;
}

.param-slider input[type="range"]::-webkit-slider-thumb {
    -webkit-appearance: none;
    appearance: none;
    width: 16px;
    height: 16px;
    background: #38bdf8;
    border-radius: 50%;
    cursor: pointer;
    transition: all 0.2s ease;
}

.param-slider input[type="range"]::-webkit-slider-thumb:hover {
    transform: scale(1.2);
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.5);
}

.binning-options {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 6px;
}

.binning-btn {
    padding: 8px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 6px;
    color: #94a3b8;
    font-size: 12px;
    cursor: pointer;
    transition: all 0.2s ease;
}

.binning-btn:hover {
    border-color: #38bdf8;
    color: #ffffff;
}

.binning-btn.active {
    background: rgba(56, 189, 248, 0.3);
    border-color: #38bdf8;
    color: #38bdf8;
}

.size-info {
    padding: 10px 12px;
    background: rgba(56, 189, 248, 0.1);
    border-radius: 8px;
    color: #94a3b8;
    font-size: 13px;
}

.action-buttons {
    display: flex;
    gap: 10px;
    margin-top: 24px;
}

.get-btn,
.apply-btn 
{
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    padding: 12px;
    border: none;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
}

.apply-btn {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
}

.apply-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 15px rgba(56, 189, 248, 0.4);
}

.get-btn {
    background: rgba(148, 163, 184, 0.2);
    color: #94a3b8;
    border: 1px solid rgba(148, 163, 184, 0.3);
}

.get-btn:hover {
    background: rgba(148, 163, 184, 0.3);
    color: #ffffff;
}
</style>
