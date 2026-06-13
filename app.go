/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: CTSystem应用主结构体，管理设备连接和更新服务
 */
package main

import (
	"CTSystem/backend"
	update "CTSystem/backend"
	DEVICE "CTSystem/backend/components"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx     context.Context
	updater *update.UpdateService

	// 设备实例
	Stage *DEVICE.MotorDevice
	CT    *DEVICE.CTDevice
	HVPS  *DEVICE.HVPSDevice
	// 项目实例
	project *backend.Project
}

// NewApp creates a new App application struct
func NewApp(
	project *backend.Project,
	stage *DEVICE.MotorDevice,
	ct *DEVICE.CTDevice,
	hvps *DEVICE.HVPSDevice) *App {
	return &App{
		project: project,
		Stage:   stage,
		CT:      ct,
		HVPS:    hvps,
	}
}

func (a *App) startup(ctx context.Context) {
	//创建更新服务
	a.updater = &update.UpdateService{}

	a.ctx = ctx

	//初始化Stage
	a.Stage.SetContent(a.ctx)
	//初始化CT
	a.CT.SetContent(a.ctx)
	a.CT.SetStage(a.Stage)
	a.CT.InitCT()
	//初始化HVPS
	a.HVPS.SetContent(a.ctx)
}

func (a *App) APIUpdate() update.GitHubRelease {
	//获取更新信息
	release, err := a.updater.GetUpdateInfo()
	if err != nil {
		fmt.Printf("获取更新信息失败: %v\n", err)
		return update.GitHubRelease{}
	}
	fmt.Printf("更新信息: %v\n", release)
	return release
}

func (a *App) GetCachedRelease() update.GitHubRelease {
	return a.updater.GetCachedRelease()
}
