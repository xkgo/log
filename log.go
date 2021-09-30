package log

import (
	"github.com/xkgo/log/api"
	"github.com/xkgo/log/driver"
)

const (
	// DebugLevel DEBUG
	DebugLevel = api.DebugLevel

	// InfoLevel INFO
	InfoLevel = api.InfoLevel

	// WarnLevel WARN
	WarnLevel = api.WarnLevel

	// ErrorLevel ERROR
	ErrorLevel = api.ErrorLevel

	// PanicLevel Panic
	PanicLevel = api.PanicLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = api.FatalLevel
)

/*
Init 日志初始化
Parameters:
- driver 驱动器，日志具体提供者
- options 参数选项，参考 options.go
*/
func Init(driver driver.Driver, options ...Option) {
	if len(options) > 0 {
		for _, option := range options {
			option()
		}
	}
	cfg.driver = driver
	// 应用默认配置
	cfg.resolveDefault()

	// 进行日志初始化，同时初始化默认日志
}

// GetLogger 获取Logger 对象，如果不存在则创建
func GetLogger(name string) api.Logger {
	// TODO GetLogger
	return nil
}
