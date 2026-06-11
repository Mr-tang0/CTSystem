/*
 * @Author: tang
 * @Date: 2026-05-23
 * @GitHub: Mr-tang0/CTSystem
 * @Description: CTSystem主程序入口，使用Wails框架构建跨平台桌面应用
 */
package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CTSystem",
		Width:  1300,
		Height: 950,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// $env:PATH = "C:\msys64\ucrt64\bin;" + [Environment]::GetEnvironmentVariable("PATH", "Machine") + ";" + [Environment]::GetEnvironmentVariable("PATH", "User")
// $env:PATH += ";D:\appfile\code\vs\CTSystem\backend\CSDK\SDK"
// $env:PATH = "C:\msys64\ucrt64\bin;" + $env:PATH
