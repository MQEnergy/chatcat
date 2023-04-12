package service

import (
	"chatcat/backend/config"
	"chatcat/backend/pkg/clog"
	"chatcat/backend/pkg/csqlite"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PushResp struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

// App struct
type App struct {
	Ctx            context.Context
	Cfg            *config.Conf
	Log            *logrus.Logger
	DB             *gorm.DB
	ChatRecordChan chan bool // 保存聊天记录的管道
	ClientId       string
	WsPushChan     chan PushResp // 推送内容通道
}

// NewApp creates a new App application struct
func NewApp() *App {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	logger := clog.NewLogger(conf.Log.DirPath, conf.Log.FileName, conf.Log.Debug)
	db, err := csqlite.WithConnect(logger, conf)
	if err != nil {
		panic(err)
	}
	return &App{
		Cfg:            conf,
		Log:            logger,
		DB:             db,
		ChatRecordChan: make(chan bool),
		WsPushChan:     make(chan PushResp),
	}
}

func (a *App) OnStartUp(ctx context.Context) {
	a.Ctx = ctx
	return
}

func (a *App) OnDomReady(ctx context.Context) {
	a.Log.Info("DomReady")
	return
}

func (a *App) OnShutdown(ctx context.Context) {
	a.Log.Info("Shutdown")
	return
}

func (a *App) OnBeforeClose(ctx context.Context) bool {
	a.Log.Info("BeforeClose")
	// 返回true将阻止程序关闭
	return false
}

func (a *App) GetDatabasePath() string {
	return csqlite.GetDatabasePath()
}

func (a *App) GetIsAutoMigrate() bool {
	return csqlite.GetIsAutoMigrate()
}

// LogInfo ...
func (a *App) LogInfo(args ...interface{}) {
	a.Log.Info(args...)
	clog.PrintInfo(args...)
}

// LogInfof ...
func (a *App) LogInfof(format string, args ...interface{}) {
	a.Log.Infof(format, args...)
	clog.PrintInfo(fmt.Sprintf(format, args...))
}

// LogError ...
func (a *App) LogError(args ...interface{}) {
	a.Log.Error(args...)
	clog.PrintError(args...)
}

// LogErrorf ...
func (a *App) LogErrorf(format string, args ...interface{}) {
	a.Log.Errorf(format, args...)
	clog.PrintError(fmt.Sprintf(format, args...))
}

// Println ...
func (a *App) Println(level logrus.Level, args ...interface{}) {
	if level == logrus.InfoLevel {
		a.LogInfo(args...)
	} else {
		a.Log.Log(level, args...)
		clog.PrintError(args...)
	}
}
