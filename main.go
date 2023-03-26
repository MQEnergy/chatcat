package main

import (
	"chatcat/backend/service"
	"chatcat/backend/service/prompt"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := service.NewApp()
	fremeless := true
	if runtime.GOOS == "darwin" {
		fremeless = false
	}
	// Create application with options
	err := wails.Run(&options.App{
		Title:             app.Cfg.App.AppName + " " + app.Cfg.App.Version,
		Width:             app.Cfg.App.Width,
		Height:            app.Cfg.App.Height,
		MinWidth:          app.Cfg.App.Width,       // 最小宽度
		MinHeight:         app.Cfg.App.Height,      // 最小高度
		MaxWidth:          app.Cfg.App.Width * 10,  // 最大宽度
		MaxHeight:         app.Cfg.App.Height * 10, // 最大高度
		DisableResize:     true,                    // 调整窗口尺寸
		Frameless:         fremeless,               // 无边框
		StartHidden:       false,                   // 启动后隐藏
		HideWindowOnClose: false,                   // 关闭窗口将隐藏而不退出应用程序
		LogLevel:          logger.DEBUG,            // 日志级别
		OnStartup:         app.OnStartUp,
		OnDomReady:        app.OnDomReady,
		OnBeforeClose:     app.OnBeforeClose, //
		OnShutdown:        app.OnShutdown,    //
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			prompt.New(app),
		},
	})

	if err != nil {
		app.LogErrorf("Error: %s", err.Error())
	}
}
