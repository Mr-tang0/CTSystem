package main

import (
	update "CTSystem/backend"
	DEVICE "CTSystem/backend/components"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx     context.Context
	updater *update.UpdateService

	ct *DEVICE.CTDevice
	// 其他字段...
}

// NewApp creates a new App application struct
func NewApp() *App {

	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//创建更新服务
	a.updater = &update.UpdateService{}

	//创建CT设备
	a.ct = DEVICE.NewCTDevice(a.ctx)
	//初始化CT
	a.ct.InitCT()
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
