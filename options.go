package log

import (
	"github.com/xkgo/log/api"
	"github.com/xkgo/log/driver"
	"os"
	"strings"
	"sync"
)

const (
	// OneKB 1KB = 1024 字节
	OneKB int = 1024

	// OneMB 1MB = 1024 * 1024
	OneMB int = OneKB * 1024

	// Forever 永久
	Forever int = -1

	// DefaultMaxAge 默认旧日志保留天数
	DefaultMaxAge int = 15

	// DefaultMaxBackups 默认旧日志保留个数
	DefaultMaxBackups int = 30
)

// FileLogConfig file log config
type FileLogConfig struct {

	// 日志文件名称，默认没有填写的话就是进程的名称 os.Arg[0] + .log, 如 app.log
	Filename string `json:"filename" xml:"filename" yaml:"filename"`

	// 日志存储目录，默认是 os.Getwd() + /logs
	Dir string `json:"dir" xml:"dir" yaml:"dir"`

	// MaxSize 单个文件最大大小，单位：字节, 默认是 10M
	MaxSize int `json:"maxSize" xml:"maxSize" yaml:"maxSize"`

	// MaxAge 保留最近几天的日志，-1 表是永久保存，0 则使用默认值保留最近15天
	MaxAge int `json:"maxAge" xml:"maxAge" yaml:"maxAge"`

	// MaxBackups 最多保留多少个日志文件，-1 表示永久保存， 0 则表示保留最近30个
	MaxBackups int `json:"maxBackups" xml:"maxBackups" yaml:"maxBackups"`

	// Compress 是否需要压缩旧日志 使用 gzip 进行压缩，默认不进行压缩
	Compress bool `json:"compress" xml:"compress" yaml:"compress"`
}

// ResolveDefault 应用默认值
func (c *FileLogConfig) ResolveDefault() {
	if len(c.Filename) == 0 {
		c.Filename = os.Args[0]
		if pos := strings.LastIndex(c.Filename, "/"); pos >= 0 {
			c.Filename = c.Filename[pos+1:]
		}
	}

	if len(c.Dir) == 0 {
		c.Dir, _ = os.Getwd()
	}

	if c.MaxSize < 1 { // 默认 10MB
		c.MaxSize = OneMB * 10
	}

	if c.MaxAge != Forever && c.MaxAge < 1 {
		c.MaxAge = DefaultMaxAge
	}

	if c.MaxBackups != Forever && c.MaxBackups < 1 {
		c.MaxBackups = DefaultMaxBackups
	}
}

type config struct {
	// 是否启用控制台日志输出
	enabledStdout bool

	// 默认文件日志配置
	defFileLogConfig *FileLogConfig

	// 不同日志级别文件日志配置
	levelFileLogConfig sync.Map

	// 日志驱动
	driver driver.Driver
}

// 日志配置
var cfg = config{
	enabledStdout: true,
}

// 处理默认值
func (c *config) resolveDefault() {
	if nil != c.defFileLogConfig {
		c.defFileLogConfig.ResolveDefault()
	}

	c.levelFileLogConfig.Range(func(key, value interface{}) bool {
		value.(*FileLogConfig).ResolveDefault()
		return true
	})
}

// Option 选项
type Option func()

// WithStdout 启用 & 禁用控制台日志
func WithStdout(enabled bool) Option {
	return func() {
		cfg.enabledStdout = enabled
	}
}

// WithDefaultFileLogConfig 设置默认的文件日志配置
func WithDefaultFileLogConfig(fileLogCfg *FileLogConfig) Option {
	return func() {
		cfg.defFileLogConfig = fileLogCfg
	}
}

// WithSetLevelFileLogConfig 设置不同的日志级别对应的日志配置文件, 允许不同日志级别，使用不同的配置
func WithSetLevelFileLogConfig(level api.Level, fileLogCfg *FileLogConfig) Option {
	return func() {
		cfg.levelFileLogConfig.Store(level, fileLogCfg)
	}
}
