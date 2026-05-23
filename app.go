package main

import (
	DEVICE "CTSystem/backend/components"
	"context"
)

// App struct
type App struct {
	ctx context.Context

	ct *DEVICE.CTDevice
	// 其他字段...
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//创建CT设备
	a.ct = DEVICE.NewCTDevice(a.ctx)
	//初始化CT
	a.ct.InitCT()

}
