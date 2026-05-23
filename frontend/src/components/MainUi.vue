<template>
    <main-container> 
        <!-- 菜单栏 -->
        <div class='top-bar' @dblclick="handleTitleBarDblClick">
            <div class='logo-section'>
                <logo>NIMTE CT CONTROL</logo>
            </div>
            <menu-bar>
                <!-- 项目 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('项目')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
                    项目
                </button>

                <!-- 设备 -->
                 <button class="menu-item" type="button" @click="onMenuItemClick('设备')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"></path></svg>
                    设备
                </button>

                <!-- 设置 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('设置')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="3"></circle><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"></path></svg>
                    设置
                </button>

                <!-- 帮助 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('帮助')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
                    帮助
                </button>

                <!-- 关于 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('关于')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                    关于
                </button>
            </menu-bar>

            <div class="top-right-controls">
                <!-- <div class="status-indicator">
                    <div class="status-dot status-online"></div>
                    <span class="status-text">在线</span>
                </div> -->
                <user-header>
                    <button class="user-avatar">
                        <svg viewBox="0 0 24 24" fill="currentColor" width="30" height="30">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z"/>
                        </svg>
                    </button>
                </user-header>
                <div class='window-actions'>    
                    <button class="window-btn minimize-btn" type="button" title="最小化" @click="minimize">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="5" y1="12" x2="19" y2="12"/>
                        </svg>
                    </button>
                    <button class="window-btn maximize-btn" type="button" title="最大化" @click="toggleMaximize">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <rect x="3" y="3" width="18" height="18" rx="2"/>
                        </svg>
                    </button>
                    <button class="window-btn close-btn" type="button" title="关闭" @click="closeApp">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="18" y1="6" x2="6" y2="18"/>
                            <line x1="6" y1="6" x2="18" y2="18"/>
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <contollr-container>
            <!-- 左控制面板 -->
            <left-panel>
                <!-- <div class="panel-header">
                    <div class="panel-title">系统状态</div>
                </div> -->

                <div class="status-group">
                    <div class="status-card power-card">
                        <div class="card-header">
                            <span class="card-title">电子枪状态</span>
                            <span class="card-status status-running">{{ Status.Gun.runing ? '运行中' : '未运行' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">灯丝电压</span>
                                <span class="info-value highlight">{{ Status.Gun.V.toFixed(2) }} kV</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">灯丝电流</span>
                                <span class="info-value highlight">{{ Status.Gun.I.toFixed(2) }} mA</span>
                            </div>
                        </div>
                    </div>

                    <div class="status-card power-card">
                        <div class="card-header">
                            <span class="card-title">高压电源状态</span>
                            <span class="card-status status-running">{{ Status.Power.runing ? '运行中' : '未运行' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">管电压</span>
                                <span class="info-value highlight">{{ Status.Power.V.toFixed(2) }} kV</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">管电流</span>
                                <span class="info-value highlight">{{ Status.Power.I.toFixed(2) }} mA</span>
                            </div>
                        </div>
                    </div>

                    <div class="status-card detector-card">
                        <div class="card-header">
                            <span class="card-title">位移台状态</span>
                            <span class="card-status status-ready">{{ Status.Stage.runing ? '运行中' : '未就绪' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">X</span>
                                <span class="info-value">{{ Status.Stage.X.toFixed(3) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">Y</span>
                                <span class="info-value">{{ Status.Stage.Y.toFixed(3) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">Z</span>
                                <span class="info-value">{{ Status.Stage.Z.toFixed(3) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">R</span>
                                <span class="info-value">{{ Status.Stage.R.toFixed(3) }} °</span>
                            </div>
                        </div>
                    </div>

                    <div class="status-card network-card">
                        <div class="card-header">
                            <span class="card-title">探测器状态</span>
                            <span class="card-status status-connected">已连接</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">IP地址</span>
                                <span class="info-value">192.168.1.100</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">连接类型</span>
                                <span class="info-value">Ethernet</span>
                            </div>
                        </div>
                    </div>
                </div>
            </left-panel>

            <!-- 中间显示面板 -->
            <div class='mid-panel-wrapper'>
                <!-- 左侧控制面板按钮 -->
                <button 
                    class="panel-toggle-btn" 
                    :class="{ 'expanded': showDetectorPanel }"
                    @click="toggleDetectorPanel"
                >{{ showDetectorPanel ? '《' : '》' }}</button>
                
                <!-- 左侧探测器控制面板 -->
                <transition name="slide">
                    <div v-if="showDetectorPanel" class="detector-panel-container">
                        <DetectorPanel @close="showDetectorPanel = false" @apply="handleApplyParams" />
                    </div>
                </transition>
                
                <!-- 中间显示区域 -->
                <div class='mid-panel'>
                    <div class='detector-container'>

                        <!-- 探测器显示 -->
                        <div class='detector-display'>

                            <!-- 无图像显示 -->
                            <div v-if="!ctImage" class='detector-placeholder'>
                                <div class='placeholder-icon'>
                                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                                        <rect x="3" y="3" width="18" height="18" rx="2"/>
                                        <circle cx="12" cy="12" r="6"/>
                                        <circle cx="12" cy="12" r="2"/>
                                    </svg>
                                </div>
                                <div class='placeholder-text'>等待图像数据...</div>
                            </div>

                            <!-- 图像显示 -->
                            <img v-else :src="ctImage" class='detector-image' alt='探测器图像' />
                        </div>
                            
                            <!-- 探测器监控面板 -->
                            <DetectorMonitor />
                        </div>

                        <!-- 历史预览图片 -->
                        <div class='detector-footer'>
                            <!-- <div class='detector-info'>
                                <span class='info-tag'>分辨率: 3072×3072</span>
                                <span class='info-tag'>像素深度: 16bit</span>
                                <span class='info-tag'>Binning: 1×1</span>
                            </div> -->
                        </div>
                </div>
            </div>

            <!-- 右控制面板 -->
            <div class='right-panel'>
                <!-- 采集控制按钮 -->
                 <div class="control-section">
                    <div class="section-title">采集控制</div>
                    <div class="button-grid">
                        <button class="action-btn primary-btn">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="currentColor">
                                <path d="M8 5v14l11-7z"/>
                            </svg>
                            <span class="btn-text">连续CT采集</span>
                        </button>
                        <button class="action-btn secondary-btn">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                                <line x1="10" y1="15" x2="10" y2="9"/><line x1="14" y1="15" x2="14" y2="9"/>
                            </svg>
                            <span class="btn-text">暂停采集</span>
                        </button>
                        <button class="action-btn danger-btn">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="currentColor">
                                <rect x="6" y="6" width="12" height="12" rx="1"/>
                            </svg>
                            <span class="btn-text">停止采集</span>
                        </button>
                        <button class="action-btn info-btn">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                                <circle cx="12" cy="13" r="4"/>
                            </svg>
                            <span class="btn-text">单帧采集</span>
                        </button>
                    </div>
                </div>


                <!-- 2. 新增：定位运动控制（上侧） -->
            <div class="control-section motion-section">
                <div class="section-title">定位运动</div>
                
                <!-- 模式切换 -->
                <div class="mode-selector">
                    <button 
                        type="button" 
                        class="mode-tab" 
                        :class="{ active: motionMode === 'absolute' }" 
                        @click="motionMode = 'absolute'"
                    >
                        绝对运动
                    </button>
                    <button 
                        type="button" 
                        class="mode-tab" 
                        :class="{ active: motionMode === 'relative' }" 
                        @click="motionMode = 'relative'"
                    >
                        相对运动
                    </button>
                </div>

                <!-- 四轴数值输入 -->
                <div class="axis-inputs-grid">
                    <div class="axis-input-item">
                        <span class="axis-input-label">X</span>
                        <div class="axis-input-wrapper">
                            <input type="number" step="0.001" v-model.number="motionTargets.X" />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">Y</span>
                        <div class="axis-input-wrapper">
                            <input type="number" step="0.001" v-model.number="motionTargets.Y" />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">Z</span>
                        <div class="axis-input-wrapper">
                            <input type="number" step="0.001" v-model.number="motionTargets.Z" />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">R</span>
                        <div class="axis-input-wrapper">
                            <input type="number" step="0.001" v-model.number="motionTargets.R" />
                            <span class="axis-unit">°</span>
                        </div>
                    </div>
                </div>

                <!-- 动作执行按钮 -->
                <div class="motion-actions">
                    <button type="button" class="motion-btn exec-btn" @click="handleExecuteMotion">
                        <svg class="btn-svg-sm" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                            <polyline points="20 6 9 17 4 12"/>
                        </svg>
                        执行
                    </button>
                    <button type="button" class="motion-btn stop-btn" @click="handleStopMotion">
                        <svg class="btn-svg-sm" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                            <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
                        </svg>
                        停止
                    </button>
                </div>
            </div>
                    



                <!-- 位移台点动按钮 -->
                <div class="sector-container">
                    <div class="bigbox s5" @click="handleSectorClick('Z+')">
                        <span class="sector-text">Z+</span>
                    </div>
                    <div class="bigbox s6" @click="handleSectorClick('R+')">
                        <span class="sector-text">R+</span>
                    </div>
                    <div class="bigbox s7" @click="handleSectorClick('R-')">
                        <span class="sector-text">R-</span>
                    </div>
                    <div class="bigbox s8" @click="handleSectorClick('Z-')">
                        <span class="sector-text">Z-</span>
                    </div>
                    <div class="sector">
                        <div class="box s1" @click="handleSectorClick('Y+')">
                            <span class="sector-text">Y+</span>
                        </div>
                        <div class="box s2" @click="handleSectorClick('X+')">
                            <span class="sector-text">X+</span>
                        </div>
                        <div class="box s3" @click="handleSectorClick('Y-')">
                            <span class="sector-text">Y-</span>
                        </div>
                        <div class="box s4" @click="handleSectorClick('X-')">
                            <span class="sector-text">X-</span>
                        </div>
                        <div class="center" @click="handleSectorClick('STOP')">
                            <span class="sector-text">STOP </span>
                        </div>
                    </div>

                </div>

                
                <!-- <div class="control-group">
                    <div class="control-section">
                        <div class="section-title">采集控制</div>
                        <div class="button-grid">
                            <button class="action-btn primary-btn">
                                <span class="btn-icon">▶️</span>
                                <span class="btn-text">开始采集</span>
                            </button>
                            <button class="action-btn secondary-btn">
                                <span class="btn-icon">⏸️</span>
                                <span class="btn-text">暂停</span>
                            </button>
                            <button class="action-btn danger-btn">
                                <span class="btn-icon">⏹️</span>
                                <span class="btn-text">停止</span>
                            </button>
                            <button class="action-btn info-btn">
                                <span class="btn-icon">📸</span>
                                <span class="btn-text">单帧</span>
                            </button>
                        </div>
                    </div>

                    <div class="control-section">
                        <div class="section-title">参数设置</div>
                        <div class="param-item">
                            <label class="param-label">曝光时间</label>
                            <div class="param-input">
                                <input type="range" min="1" max="100" value="50" class="param-slider"/>
                                <span class="param-value">50ms</span>
                            </div>
                        </div>
                        <div class="param-item">
                            <label class="param-label">管电压</label>
                            <div class="param-input">
                                <input type="range" min="40" max="150" value="120" class="param-slider"/>
                                <span class="param-value">120kV</span>
                            </div>
                        </div>
                        <div class="param-item">
                            <label class="param-label">管电流</label>
                            <div class="param-input">
                                <input type="range" min="10" max="500" value="200" class="param-slider"/>
                                <span class="param-value">200mA</span>
                            </div>
                        </div>
                    </div>

                    <div class="control-section">
                        <div class="section-title">快速操作</div>
                        <div class="quick-actions">
                            <button class="quick-btn">🔄 自动校准</button>
                            <button class="quick-btn">📊 图像分析</button>
                            <button class="quick-btn">💾 保存图像</button>
                            <button class="quick-btn">📤 导出数据</button>
                        </div>
                    </div>
                </div> -->
            </div>
        </contollr-container>

        <!-- 底部栏 -->
        <div class='bottom-bar'>
            <div class="bottom-left">
                <span class="status-item">系统状态: <span class="status-normal">{{Status.SystemStatus}}</span></span>
                <span class="status-item">运行时间: <span class="time-display">{{Status.RunTime}}</span></span>
            </div>
            <!-- <div class="bottom-right">
                <span class="status-item">内存使用: <span class="memory-display">{{Status.MemoryUsage}}</span></span>
                <span class="status-item">CPU使用: <span class="cpu-display">{{Status.CPUUsage}}</span></span>
            </div> -->
        </div>
    </main-container>
    
    <!-- 设备连接模态框 -->
    <DeviceConnectModal :visible="showDeviceModal" @close="showDeviceModal = false" />

    <!-- 更新提示模态框 -->
    <teleport to="body">
        <transition name="modal">
            <div v-if="showUpdateModal" class="modal-overlay" @click.self="showUpdateModal = false">
                <div class="modal-container update-modal">
                    <div class="modal-header">
                        <h3 class="modal-title">发现新版本</h3>
                        <button class="modal-close" @click="showUpdateModal = false">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    <div class="modal-body">
                        <div class="update-content">
                            <div class="update-icon">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/>
                                    <path d="M3 3v5h5"/>
                                    <path d="M3 16a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/>
                                    <path d="M16 21h5v-5"/>
                                </svg>
                            </div>
                            <div class="update-info">
                                <p class="update-version">新版本: <strong>{{ updateInfo.tagName }}</strong></p>
                                <p class="update-desc">发现应用程序更新，建议及时升级以获得更好的体验。</p>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button class="btn btn-secondary" @click="showUpdateModal = false">
                            稍后更新
                        </button>
                        <button class="btn btn-primary" @click="handleUpdate">
                            立即更新
                        </button>
                    </div>
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, watch, computed, nextTick } from 'vue';
import { WindowMinimise,
          WindowToggleMaximise,
          Quit,
          WindowIsMaximised,
          WindowUnmaximise,
          EventsOn,
          BrowserOpenURL
        } from '../../wailsjs/runtime/runtime'
import DetectorPanel from './DetectorPanel.vue';
import DetectorMonitor from './DetectorMonitor.vue';
import DeviceConnectModal from './DeviceConnectModal.vue';

import { APIUpdate } from '../../wailsjs/go/main/App'

//项目管理模态框状态
const showProjectModal = ref(false);
// 设备连接模态框状态
const showDeviceModal = ref(false);
// 系统状态模态框状态
const showSettingModal = ref(false);
//帮助模态框状态
const showHelpModal = ref(false);
//关于模态框状态
const showAboutModal = ref(false);

// 更新模态框状态
const showUpdateModal = ref(false);

// 更新处理函数
const handleUpdate = () => {
    if (updateInfo.value.htmlUrl) {
        // window.open(updateInfo.value.htmlUrl, '_blank');
        BrowserOpenURL(updateInfo.value.htmlUrl) 
        showUpdateModal.value = false;
    }
};

// 图像数据状态
const ctImage = ref('');

const onMenuItemClick = (option) => {
    if (option === '项目') {
        showProjectModal.value = true;
    }else if (option === '设备') {
        showDeviceModal.value = true;
    }else if (option === '设置') {
        showSettingModal.value = true;
    }else if (option === '帮助') {
        showHelpModal.value = true;
    }else if (option === '关于') {
        showAboutModal.value = true;
    }
    console.log(option);
};

// 左侧控制面板状态
const showDetectorPanel = ref(false);
// 切换控制面板显示
const toggleDetectorPanel = () => {
    showDetectorPanel.value = !showDetectorPanel.value;
};

const handleTitleBarDblClick = async () => {
    const isMax = await WindowIsMaximised()
    if (isMax) {
      WindowUnmaximise()
    } else {
      WindowToggleMaximise()
    }
  }

const minimize = () => WindowMinimise()
const toggleMaximize = () => WindowToggleMaximise()
const closeApp = async () => {
Quit()
// const result = await HandleClose()
// if (result) {
//   Quit()
// }
// else{
//   notify.error('未知错误,系统未能正常退出', "error", 3000)
// }
}



// 应用参数
const handleApplyParams = (params) => {
    console.log('[探测器参数] 应用参数:', params);
    // 这里可添加你的 Wails API 调用，将参数发送给后端
    // window.go.main.App.SetDetectorParams(params);
};

const Status = reactive({
    Power:{
        V: 1.11,
        I: 2.22,
        runing: false
    },
    Gun:{
        V: 3.33,
        I: 4.44,
        runing: false
    },
    Stage:{
        X:0.001,
        Y:0.002,
        Z:0.003,
        R:0.004,
        runing: false
    },
    SystemStatus: '正常',
    RunTime: '00:00:00',
    StartTime: '',
    MemoryUsage: '1.2GB / 8GB',
    CPUUsage: '23%'
})

const motionMode = ref('absolute'); // 'absolute' (绝对运动) 或 'relative' (相对运动)

const motionTargets = reactive({
    X: 0.000,
    Y: 0.000,
    Z: 0.000,
    R: 0.000
});

// 执行定位运动
const handleExecuteMotion = () => {
    console.log(`[定位运动] 模式: ${motionMode.value} | 目标位置:`, { ...motionTargets });
    // 这里可添加你的 Wails API 调用，例如：
    // window.go.main.App.MoveTo(motionMode.value, motionTargets.X, motionTargets.Y, motionTargets.Z, motionTargets.R);
};

// 紧急停止运动
const handleStopMotion = () => {
    console.log('[定位运动] 触发停止');
    // window.go.main.App.EmergencyStop();
};


const updateInfo = ref({
  tagName: '无',
  htmlUrl: 'https://github.com'
})

onMounted(async () =>{
    const release = await APIUpdate()
    if (release) {
        updateInfo.value.tagName = release.tag_name
        updateInfo.value.htmlUrl = release.html_url
        if (release.assets.length > 0) {
            updateInfo.value.htmlUrl = release.assets[0].browser_download_url
        }
        // 显示更新模态框
        showUpdateModal.value = true
    }

    Status.StartTime = new Date().getTime()
    setInterval(()=>{
        Status.RunTime = new Date(new Date().getTime()-Status.StartTime-8*60*60*1000).toLocaleTimeString()
    },1000)

    // 监听后端ct_image事件
    EventsOn('ct_image', (imageData) => {
        console.log('[MainUi] 收到图像数据:', imageData ? imageData.substring(0, 100) + '...' : '空');
        ctImage.value = imageData;
    });
});





</script>

<style>
/* 容器重置 */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body, html { 
    margin: 0; 
    padding: 0; 
    height: 100%; 
    font-family: 'Segoe UI', 'Microsoft YaHei', -apple-system, BlinkMacSystemFont, sans-serif;
}

main-container { 
    display: flex;
    flex-direction: column;
    height: 100vh;
    background: linear-gradient(135deg, #0f172a 0%, #1e293b 50%, #0f172a 100%);
    color: #e2e8f0;
    overflow: hidden;
}

/* 顶部栏优化 */
.top-bar { 
    display: flex;
    align-items: center;
    gap: 24px;
    padding: 0 20px;
    height: 64px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.9) 100%);
    border-bottom: 1px solid rgba(148, 163, 184, 0.1);
    backdrop-filter: blur(10px);
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
    --wails-draggable: drag;
}

.logo-section {
    display: flex;
    align-items: center;
    gap: 12px;
}

.logo-icon {
    width: 32px;
    height: 32px;
    color: #38bdf8;
    animation: rotate 20s linear infinite;
}

@keyframes rotate {
    from { transform: rotate(0deg); }
    to { transform: rotate(360deg); }
}

logo {
    padding: 8px 20px;
    border-radius: 8px;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.15) 0%, rgba(56, 189, 248, 0.05) 100%);
    color: #38bdf8;
    font-weight: 700;
    font-size: 18px;
    letter-spacing: 0.1em;
    text-shadow: 0 0 20px rgba(56, 189, 248, 0.3);
    border: 1px solid rgba(56, 189, 248, 0.2);
}

menu-bar { 
    display: flex; 
    align-items: center; 
    gap: 8px;
}

.menu-item { 
    display: flex;
    align-items: center;
    gap: 8px;
    color: #ffffff;
    background: transparent;
    border: none;
    cursor: pointer;
    font-size: 14px;
    padding: 10px 16px;
    border-radius: 8px;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
    overflow: hidden;
}

.menu-item::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.1) 0%, transparent 100%);
    opacity: 0;
    transition: opacity 0.3s;
}

.menu-item:hover {
    color: #ffffff;
    background: rgba(56, 189, 248, 0.1);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(56, 189, 248, 0.2);
}

.menu-item:hover::before {
    opacity: 1;
}

.menu-item {
    display: flex;
    align-items: center;
}

.menu-svg {
    width: 20px;
    height: 20px;
    margin-right: 8px;
    color: #9ba2ac; /* 默认深灰 */
    transition: color 0.2s ease;
}

/* 悬停时随文字一起变色 */
.menu-item:hover .menu-svg {
    color: #1d4ed8; 
}

.top-right-controls {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 20px;
}

.status-indicator {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 16px;
    background: rgba(16, 185, 129, 0.1);
    border-radius: 20px;
    border: 1px solid rgba(16, 185, 129, 0.2);
}

.status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    animation: pulse 2s infinite;
}

.status-online {
    background: #10b981;
    box-shadow: 0 0 10px rgba(16, 185, 129, 0.5);
}

@keyframes pulse {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.5; }
}

.status-text {
    font-size: 12px;
    color: #10b981;
    font-weight: 600;
}

user-header { 
    display: flex; 
    align-items: center; 
}

.user-avatar {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background: linear-gradient(135deg, #334155 0%, #1e293b 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    border: 2px solid rgba(56, 189, 248, 0.3);
    cursor: pointer;
    transition: all 0.3s;
    color: #38bdf8;
}

.user-avatar:hover {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
    transform: scale(1.1);
    box-shadow: 0 0 20px rgba(56, 189, 248, 0.4);
}

.window-actions { 
    display: flex; 
    align-items: center; 
    gap: 4px;
}

.window-btn {
    width: 36px;
    height: 36px;
    border: none;
    background: transparent;
    color: #94a3b8;
    cursor: pointer;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s;
    padding: 0;
}

.window-btn svg {
    width: 18px;
    height: 18px;
}

.window-btn:hover { 
    background: rgba(56, 189, 248, 0.1);
    color: #38bdf8;
    transform: scale(1.1);
}

.window-btn.close-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    color: #ef4444;
}

/* 中间区域布局 */
contollr-container {
    display: flex;
    flex: 1;
    overflow: hidden;
    gap: 16px;
    padding: 16px;
}

left-panel {
    width: 280px;
    background: linear-gradient(180deg, rgba(30, 41, 59, 0.8) 0%, rgba(15, 23, 42, 0.9) 100%);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(148, 163, 184, 0.1);
}

.panel-title {
    font-size: 16px;
    font-weight: 600;
    color: #38bdf8;
}

.panel-icon {
    font-size: 20px;
}

.status-group {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.status-card {
    background: linear-gradient(135deg, rgba(15, 23, 42, 0.6) 0%, rgba(30, 41, 59, 0.4) 100%);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 10px;
    overflow: hidden;
    transition: all 0.3s;
}

.status-card:hover {
    border-color: rgba(56, 189, 248, 0.3);
    box-shadow: 0 4px 20px rgba(56, 189, 248, 0.1);
    transform: translateY(-2px);
}

.card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 16px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(148, 163, 184, 0.05);
}

.card-title {
    font-size: 13px;
    font-weight: 600;
    color: #e2e8f0;
}

.card-status {
    font-size: 11px;
    padding: 4px 10px;
    border-radius: 12px;
    font-weight: 600;
}

.status-running {
    background: rgba(16, 185, 129, 0.2);
    color: #10b981;
}

.status-ready {
    background: rgba(56, 189, 248, 0.2);
    color: #38bdf8;
}

.status-connected {
    background: rgba(168, 85, 247, 0.2);
    color: #a855f7;
}

.card-body {
    padding: 12px 16px;
}

.info-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid rgba(148, 163, 184, 0.05);
}

.info-row:last-child {
    border-bottom: none;
}

.info-label {
    font-size: 12px;
    color: #94a3b8;
}

.info-value {
    font-size: 13px;
    font-weight: 600;
    color: #e2e8f0;
}

.info-value.highlight {
    color: #38bdf8;
    font-family: 'Courier New', monospace;
}

.mid-panel-wrapper {
    flex: 1;
    display: flex;
    align-items: stretch;
    gap: 0;
    position: relative;
    overflow: hidden;
}

/* 左侧控制面板切换按钮 */
.panel-toggle-btn {
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    z-index: 10;
    width: 28px;
    height: 160px;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.3) 0%, rgba(14, 165, 233, 0.3) 100%);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-left: none;
    border-radius: 0 8px 8px 0;
    color: #38bdf8;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    backdrop-filter: blur(10px);
}

.panel-toggle-btn:hover {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.5) 0%, rgba(14, 165, 233, 0.5) 100%);
    border-color: rgba(56, 189, 248, 0.5);
    box-shadow: 2px 0 15px rgba(56, 189, 248, 0.3);
}

.panel-toggle-btn.expanded {
    left: 280px;
}

.panel-toggle-btn svg {
    width: 18px;
    height: 18px;
    transition: transform 0.3s ease;
}

.panel-toggle-btn.expanded svg {
    transform: rotate(180deg);
}

/* 左侧控制面板容器 */
.detector-panel-container {
    width: 280px;
    flex-shrink: 0;
    height: 100%;
    overflow: hidden;
}

/* 滑动过渡动画 */
.slide-enter-active,
.slide-leave-active {
    transition: transform 0.3s ease;
}

.slide-enter-from,
.slide-leave-to {
    transform: translateX(-100%);
}

.mid-panel {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0;
    transition: margin-left 0.3s ease;
}

.detector-container {
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(15, 23, 42, 0.8) 0%, rgba(30, 41, 59, 0.6) 100%);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.detector-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(148, 163, 184, 0.1);
}

.detector-title {
    font-size: 16px;
    font-weight: 600;
    color: #38bdf8;
}

.detector-controls {
    display: flex;
    gap: 8px;
}

.control-btn {
    width: 32px;
    height: 32px;
    border: 1px solid rgba(148, 163, 184, 0.2);
    background: rgba(15, 23, 42, 0.5);
    border-radius: 6px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s;
    padding: 0;
}

.control-btn:hover {
    background: rgba(56, 189, 248, 0.2);
    border-color: rgba(56, 189, 248, 0.4);
    color: #38bdf8;
}

.control-btn svg {
    width: 16px;
    height: 16px;
}

.detector-display {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
    background: rgba(0, 0, 0, 0.3);
    position: relative;
}

.detector-placeholder {
    text-align: center;
    color: #64748b;
}

.placeholder-icon {
    width: 120px;
    height: 120px;
    margin: 0 auto 20px;
    opacity: 0.3;
    animation: float 3s ease-in-out infinite;
}

@keyframes float {
    0%, 100% { transform: translateY(0); }
    50% { transform: translateY(-10px); }
}

.placeholder-text {
    font-size: 14px;
    font-weight: 500;
}

.detector-image {
    width: 100%;
    height: 100%;
    object-fit: contain;
    background: #0f172a;
}

.detector-footer {
    padding: 12px 20px;
    height: 88px;
    background: rgba(56, 189, 248, 0.05);
    border-top: 1px solid rgba(148, 163, 184, 0.1);
}

.detector-info {
    display: flex;
    gap: 16px;
    flex-wrap: wrap;
}

.info-tag {
    font-size: 11px;
    color: #94a3b8;
    padding: 4px 12px;
    background: rgba(15, 23, 42, 0.5);
    border-radius: 12px;
    border: 1px solid rgba(148, 163, 184, 0.1);
}

.right-panel {
    width: 280px;
    background: linear-gradient(180deg, rgba(30, 41, 59, 0.8) 0%, rgba(15, 23, 42, 0.9) 100%);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.control-group {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.control-section {
    background: rgba(15, 23, 42, 0.4);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 10px;
    padding: 16px;
}

.section-title {
    font-size: 14px;
    font-weight: 600;
    color: #38bdf8;
    margin-bottom: 16px;
    padding-bottom: 8px;
    border-bottom: 1px solid rgba(148, 163, 184, 0.1);
}

.button-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
}

.action-btn {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    padding: 16px 12px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s;
    font-size: 12px;
    font-weight: 600;
    height: 75px;
    width: 120px;
}

.action-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.primary-btn {
    background: rgba(14, 165, 233, 0.1); 
    border: 1px solid rgba(14, 165, 233, 0.3); 
    color: #38bdf8;
}

.primary-btn:hover {
    background: linear-gradient(135deg, #0ea5e9 0%, #0284c7 100%);
    box-shadow: 0 4px 12px rgba(56, 189, 248, 0.4);
}

.secondary-btn {
    background: rgba(168, 85, 247, 0.1); border: 1px solid rgba(168, 85, 247, 0.25); color: #c084fc;
}

.secondary-btn:hover {
    background: linear-gradient(135deg, #9333ea 0%, #7c3aed 100%);
    box-shadow: 0 4px 12px rgba(168, 85, 247, 0.4);
}

.danger-btn {
    background: rgba(239, 68, 68, 0.08); border: 1px solid rgba(239, 68, 68, 0.25); color: #f87171;
}

.danger-btn:hover {
    background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
}

.info-btn {
    background: rgba(16, 185, 129, 0.1); border: 1px solid rgba(16, 185, 129, 0.25); color: #34d399;
}

.info-btn:hover {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4);
}

.btn-icon {
    font-size: 20px;
}

.param-item {
    margin-bottom: 16px;
}

.param-item:last-child {
    margin-bottom: 0;
}

.param-label {
    display: block;
    font-size: 12px;
    color: #94a3b8;
    margin-bottom: 8px;
}

.param-input {
    display: flex;
    align-items: center;
    gap: 12px;
}

.param-slider {
    flex: 1;
    -webkit-appearance: none;
    height: 6px;
    background: rgba(148, 163, 184, 0.2);
    border-radius: 3px;
    outline: none;
}

.param-slider::-webkit-slider-thumb {
    -webkit-appearance: none;
    width: 16px;
    height: 16px;
    background: #38bdf8;
    border-radius: 50%;
    cursor: pointer;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.5);
}

.param-value {
    font-size: 12px;
    font-weight: 600;
    color: #38bdf8;
    min-width: 50px;
    text-align: right;
}

.quick-actions {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 8px;
}

.quick-btn {
    padding: 12px;
    background: rgba(15, 23, 42, 0.5);
    border: 1px solid rgba(148, 163, 184, 0.2);
    border-radius: 8px;
    color: #e2e8f0;
    font-size: 12px;
    cursor: pointer;
    transition: all 0.3s;
}

.quick-btn:hover {
    background: rgba(56, 189, 248, 0.1);
    border-color: rgba(56, 189, 248, 0.3);
    color: #38bdf8;
    transform: translateY(-2px);
}

/* 底部栏 */
.bottom-bar {
    height: 40px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.9) 100%);
    border-top: 1px solid rgba(148, 163, 184, 0.1);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
    backdrop-filter: blur(10px);
}

.bottom-left,
.bottom-right {
    display: flex;
    gap: 24px;
}

.status-item {
    font-size: 12px;
    color: #94a3b8;
}

.status-normal {
    color: #10b981;
    font-weight: 600;
}

.time-display,
.memory-display,
.cpu-display {
    color: #38bdf8;
    font-weight: 600;
    font-family: 'Courier New', monospace;
}

/* 滚动条样式 */
::-webkit-scrollbar {
    width: 6px;
}

::-webkit-scrollbar-track {
    background: rgba(15, 23, 42, 0.5);
}

::-webkit-scrollbar-thumb {
    background: rgba(56, 189, 248, 0.3);
    border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
    background: rgba(56, 189, 248, 0.5);
}

.sector-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
  background: rgba(15, 23, 42, 0.5);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 0;
}

.sector {
  position: relative;
  width: 200px;
  height: 200px;
  margin: 0 auto;
  border: 2px solid rgba(56, 189, 248, 0.3);
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(15, 23, 42, 0.8) 0%, rgba(30, 41, 59, 0.6) 100%);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3), inset 0 0 20px rgba(56, 189, 248, 0.1);
}

.sector .box {
  position: absolute;
  width: 200px;
  height: 200px;
  border-radius: 50%;
  clip: rect(0px, 99px, 99px, 0);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}






.sector .box,
.sector .center {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

/* 扇形文字样式 */
.sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  white-space: nowrap;
}

.sector .s1 {
  background: linear-gradient(135deg, #4a595f 0%, #242d31 100%);
  transform: rotate(45deg);
}

.sector .s1 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(-45deg) translate(0px, -65px);
}

.sector .s1:hover,
.sector .s1.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #73858b 0%, #242d31 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s2 {
  background: linear-gradient(135deg, #4a595f 0%, #242d31 100%);
  transform: rotate(135deg);
}


.sector .s2 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(-135deg) translate(65px, 0px);
}

.sector .s2:hover,
.sector .s2.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #73858b 0%, #242d31 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s3 {
  background: linear-gradient(135deg, #4a595f 0%, #242d31 100%);
  transform: rotate(-135deg);
}

/* 给 s3 的文字也加上定位 */
.sector .s3 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(135deg) translate(0px, 65px);
}

.sector .s3:hover,
.sector .s3.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #73858b 0%, #242d31 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s4 {
  background: linear-gradient(135deg, #4a595f 0%, #242d31 100%);
  transform: rotate(-45deg);
}

/* 给 s4 的文字也加上定位 */
.sector .s4 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(45deg) translate(-65px, 0px);
}

.sector .s4:hover,
.sector .s4.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #73858b 0%, #242d31 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .center {
  width: 90px;
  height: 90px;
  position: absolute;
  left: 54px;
  top: 54px;
  border: 2px solid rgba(56, 189, 248, 1);
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(239, 68, 68, 1) 0%, rgba(220, 38, 38, 1) 100%);
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.3);
}

.sector .center:hover {
  background: linear-gradient(135deg, rgba(220, 38, 38, 0.9) 0%, rgba(185, 28, 28, 0.9) 100%);
  border-color: rgba(239, 68, 68, 0.6);
  box-shadow: 0 0 25px rgba(239, 68, 68, 0.6);
  transform: scale(1.05);
}


.bigbox{
  position: absolute;
  width: 115px;
  height: 115px;
  background: linear-gradient(135deg, #374a52 0%, #242d31 100%);
  border: 2px solid rgba(56, 189, 248, 0.3);
  font-size: 18px;
  font-weight: 600;

  transform: translate(-50%, -50%); 
  border-radius: 5px;
  display: flex;
}

.bigbox:hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #73858b 0%, #242d31 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.bigbox.s5 {
  transform: rotate(0deg) translate(-65px, -65px);
}

/* 修正选择器：bigbox.s5 内部的 .sector-text */
.bigbox.s5 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(0deg) translate(10px, 10px);
}

.bigbox.s6 {
  transform: rotate(90deg) translate(-65px, -65px);
}
.bigbox.s6 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(-90deg) translate(-15px, 15px);
}

.bigbox.s7 {
  transform: rotate(180deg) translate(-65px, -65px);
}

.bigbox.s7 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(180deg) translate(-15px, -15px);
}



.bigbox.s8 {
  transform: rotate(270deg) translate(-65px, -65px);
}

.bigbox.s8 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  pointer-events: none;
  transform: rotate(-270deg) translate(15px, -15px);
}

.btn-svg {
    width: 20px;
    height: 20px;
    color: #ffffff;
    flex-shrink: 0;
}

.btn-svg-sm {
    width: 14px;
    height: 14px;
    color: #ffffff;
    flex-shrink: 0;
}

/* 运动控制面板容器 */
.motion-section {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

/* 绝对/相对 模式选择器 */
.mode-selector {
    display: flex;
    background: rgba(15, 23, 42, 0.6);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 8px;
    padding: 3px;
}

.mode-tab {
    flex: 1;
    background: transparent;
    border: none;
    color: #94a3b8;
    padding: 8px 0;
    font-size: 12px;
    font-weight: 600;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.2s ease;
}

.mode-tab:hover {
    color: #e2e8f0;
}

.mode-tab.active {
    background: rgba(56, 189, 248, 0.15);
    color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.1);
}

/* 四轴参数输入网格 */
.axis-inputs-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
}

.axis-input-item {
    display: flex;
    align-items: center;
    background: rgba(15, 23, 42, 0.4);
    border: 1px solid rgba(148, 163, 184, 0.1);
    border-radius: 6px;
    padding: 6px 10px;
    transition: all 0.3s;
}

.axis-input-item:focus-within {
    border-color: rgba(56, 189, 248, 0.5);
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.15);
}

.axis-input-label {
    font-size: 13px;
    font-weight: 700;
    color: #38bdf8;
    width: 18px;
}

.axis-input-wrapper {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.axis-input-wrapper input {
    width: 100%;
    background: transparent;
    border: none;
    color: #e2e8f0;
    font-family: 'Courier New', monospace;
    font-weight: 600;
    font-size: 12px;
    padding: 0 4px;
    outline: none;
    text-align: right;
}

/* 隐藏 Chrome 等浏览器 input[type=number] 的微调按钮 */
.axis-input-wrapper input::-webkit-outer-spin-button,
.axis-input-wrapper input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

.axis-unit {
    font-size: 11px;
    color: #64748b;
    margin-left: 6px;
    min-width: 22px;
    text-align: left;
}

/* 定位运动动作按钮 */
.motion-actions {
    display: grid;
    grid-template-columns: 1.2fr 1fr;
    gap: 10px;
    margin-top: 4px;
}

.motion-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    padding: 10px;
    border: none;
    border-radius: 6px;
    font-size: 12px;
    font-weight: 600;
    color: #ffffff;
    cursor: pointer;
    transition: all 0.3s;
}

.motion-btn:hover {
    transform: translateY(-1px);
}

/* 执行按钮 - 绿色 */
.exec-btn {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    box-shadow: 0 2px 8px rgba(16, 185, 129, 0.2);
}

.exec-btn:hover {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4);
}

/* 停止按钮 - 红色 */
.stop-btn {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    box-shadow: 0 2px 8px rgba(239, 68, 68, 0.2);
}

.stop-btn:hover {
    background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
}

/* 更新模态框样式 */
/* 遮罩层 */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
    backdrop-filter: blur(4px);
}

/* 模态框容器 */
.modal-container {
    background: linear-gradient(180deg, rgba(30, 41, 59, 0.95) 0%, rgba(15, 23, 42, 0.98) 100%);
    border-radius: 16px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
    border: 1px solid rgba(148, 163, 184, 0.1);
    overflow: hidden;
}

.modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
}

.modal-title {
    margin: 0;
    font-size: 15px;
    font-weight: 600;
    color: #cbd5e1;
}

.modal-close {
    background: transparent;
    border: none;
    color: #94a3b8;
    cursor: pointer;
    padding: 4px;
    border-radius: 8px;
    transition: all 0.2s;
    width: 28px;
    height: 28px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-close:hover {
    background: rgba(255, 255, 255, 0.1);
    color: #f1f5f9;
}

.modal-close svg {
    width: 18px;
    height: 18px;
}

.modal-body {
    padding: 20px;
}

.modal-footer {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 12px;
    padding: 16px 20px;
}

.btn {
    padding: 10px 20px;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    border: none;
    transition: all 0.2s;
}

.btn-primary {
    background: linear-gradient(135deg, rgba(56, 189, 248, 1) 0%, rgba(59, 130, 246, 1) 100%);
    color: white;
}

.btn-primary:hover {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.9) 0%, rgba(59, 130, 246, 0.9) 100%);
    transform: translateY(-1px);
}

.btn-secondary {
    background: rgba(255, 255, 255, 0.1);
    color: #cbd5e1;
}

.btn-secondary:hover {
    background: rgba(255, 255, 255, 0.15);
}

/* 更新模态框内容 */
.update-modal {
    width: 420px;
    max-width: 90vw;
}

.update-modal .modal-header {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.1) 0%, rgba(59, 130, 246, 0.1) 100%);
    border-bottom: 1px solid rgba(148, 163, 184, 0.1);
}

.update-modal .modal-title {
    color: #f1f5f9;
    font-size: 16px;
    font-weight: 600;
}

.update-content {
    display: flex;
    align-items: center;
    gap: 20px;
    padding: 10px 0;
}

.update-icon {
    width: 64px;
    height: 64px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.2) 0%, rgba(59, 130, 246, 0.2) 100%);
    border-radius: 16px;
    color: #38bdf8;
    flex-shrink: 0;
}

.update-icon svg {
    width: 32px;
    height: 32px;
}

.update-info {
    flex: 1;
}

.update-version {
    font-size: 15px;
    color: #f1f5f9;
    margin: 0 0 8px 0;
    font-weight: 500;
}

.update-version strong {
    color: #38bdf8;
}

.update-desc {
    font-size: 13px;
    color: #94a3b8;
    margin: 0;
    line-height: 1.5;
}

.update-modal .modal-footer {
    border-top: 1px solid rgba(148, 163, 184, 0.1);
}

/* 模态框动画 */
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-active .modal-container,
.modal-leave-active .modal-container {
    transition: transform 0.3s ease, opacity 0.3s ease;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
    transform: scale(0.95);
    opacity: 0;
}

</style>
