package clog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// NewLogger 构造日志服务
func NewLogger(logPath, module string, debug bool) *logrus.Logger {
	once.Do(func() {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic("创建日志文件失败: " + err.Error())
		}
		// 定义文件前缀和日志名称
		prefix := logPath + "/" + module
		latestLogFile := prefix + ".log"
		logger = logrus.New()
		// 设置输出
		logger.Out = src
		// 设置日志级别
		if debug == true {
			logger.SetLevel(logrus.DebugLevel)
		}
		// 设置rotatelogs
		logWriter, err := rotatelogs.New(
			prefix+"-%Y-%m-%d.log",                    // 生成实际文件名的模式
			rotatelogs.WithLinkName(latestLogFile),    // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(30*24*time.Hour),    // 设置最大保存时间(30天)
			rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
		)
		if err != nil {
			panic("创建日志文件失败: " + err.Error())
		}
		logger.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: logWriter,
				logrus.InfoLevel:  logWriter,
				logrus.FatalLevel: logWriter,
				logrus.PanicLevel: logWriter,
				logrus.WarnLevel:  logWriter,
				logrus.ErrorLevel: logWriter,
			},
			&logrus.JSONFormatter{
				TimestampFormat: time.DateTime,
			},
		))
	})
	return logger
}

// PrintInfo ...
func PrintInfo(args ...interface{}) {
	fmt.Println(fmt.Sprintf("\u001B[34m[%s] %v\u001B[0m", time.Now().Format(time.DateTime), args))
}

// PrintError ...
func PrintError(args ...interface{}) {
	fmt.Println(fmt.Sprintf("\x1b[31m[%s] %v\x1b[0m", time.Now().Format(time.DateTime), args))
}
