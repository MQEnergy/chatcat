package service

import (
	"chatcat/backend/config"
	"chatcat/backend/pkg/clog"
	"chatcat/backend/pkg/csqlite"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	Cfg *config.Conf
	Log *logrus.Logger
	DB  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	// 读取配置
	config.ConfEnv = "test"
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	// 连接sqlite
	db, err := csqlite.WithConnect()
	if err != nil {
		panic(err)
	}
	return &App{
		Cfg: conf,
		Log: clog.NewLogger(conf.Log.DirPath, conf.Log.FileName, conf.Log.Debug),
		DB:  db,
	}
}

func (a *App) OnStartUp(ctx context.Context) {
	a.ctx = ctx
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

// LogInfo ...
func (t *App) LogInfo(args ...interface{}) {
	t.Log.Info(args)
	clog.PrintInfo(args...)
}

// LogInfof ...
func (t *App) LogInfof(format string, args ...interface{}) {
	t.Log.Infof(format, args...)
	clog.PrintInfo(args...)
}

// LogError ...
func (t *App) LogError(args ...interface{}) {
	t.Log.Error(args...)
	clog.PrintError(args...)
}

// LogErrorf ...
func (t *App) LogErrorf(format string, args ...interface{}) {
	t.Log.Errorf(format, args...)
	clog.PrintError(args...)
}

// Println ...
func (t *App) Println(level logrus.Level, args ...interface{}) {
	if level == logrus.InfoLevel {
		t.LogInfo(args)
	} else {
		t.Log.Log(level, args)
		clog.PrintError(args...)
	}
}
