package api

// Level log level
type Level int8

const (
	// DebugLevel DEBUG
	DebugLevel Level = iota

	// InfoLevel INFO
	InfoLevel

	// WarnLevel WARN
	WarnLevel

	// ErrorLevel ERROR
	ErrorLevel

	// PanicLevel Panic
	PanicLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// Logger logger interface def
type Logger interface {

	// Name 日志名称
	Name() string

	// Level 日志级别
	Level() Level

	// SetLevel 设置level 级别
	SetLevel(level Level)

	// IsLevelEnabled 给定日志级别是否需要打印日志
	IsLevelEnabled(level Level) bool

	// Debug debug log
	Debug(v ...interface{})

	// Info log
	Info(v ...interface{})

	// Warn log
	Warn(v ...interface{})

	// Error log
	Error(v ...interface{})

	// Panic log
	Panic(v ...interface{})

	// Fatal log
	Fatal(v ...interface{})

	// Debugf debug log
	Debugf(template string, v ...interface{})

	// Infof log
	Infof(template string, v ...interface{})

	// Warnf log
	Warnf(template string, v ...interface{})

	// Errorf log
	Errorf(template string, v ...interface{})

	// Panicf log
	Panicf(template string, v ...interface{})

	// Fatalf log
	Fatalf(template string, v ...interface{})
}
