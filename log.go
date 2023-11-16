package zerolog

import (
	"context"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

type Config struct {
	Filename     string `json:"Filename"`
	MaxSize      int    `json:"MaxSize"`
	MaxAge       int    `json:"MaxAge"`
	MaxBackups   int    `json:"MaxBackups"`
	Compress     bool   `json:"Compress"`
	RotationTime int    `json:"RotationTime"`
}

func Init(cfg *Config) error {
	// 创建一个 lumberjack.Logger 对象，用于按照文件大小切割日志
	logSize := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}
	// 创建一个 rotatelogs 的 Handler 对象，用于按照时间切割日志
	logTime, err := rotatelogs.New(
		cfg.Filename+"%Y%m%d",                                                  // 日志文件名格式
		rotatelogs.WithLinkName(cfg.Filename),                                  // 最新日志文件的软链接
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*24*time.Hour),          // 旧日志文件的最大保留时间，单位：天
		rotatelogs.WithRotationTime(time.Duration(cfg.RotationTime)*time.Hour), // 日志切割时间间隔，单位：小时
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create rotatelogs handler")
		return err
	}

	// 创建一个 zerolog.ConsoleWriter 对象，用于输出日志到控制台
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	// 创建一个 zerolog.MultiLevelWriter 对象，将 logSize、logTime 和 consoleWriter 三个日志输出目标合并为一个
	logWriter := zerolog.MultiLevelWriter(
		zerolog.SyncWriter(logSize),
		zerolog.SyncWriter(logTime),
		consoleWriter,
	)

	// 将 logWriter 对象设置为 zerolog 的全局输出目标
	log.Logger = log.Output(logWriter)

	// 示例日志输出```go
	log.Info().Msg("log init success")

	return nil
}

// TODO: 以下为重新原生方法

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return log.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return log.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return log.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return log.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return log.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return log.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return log.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return log.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return log.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return log.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return log.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return log.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return log.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return log.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return log.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.Print(v)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
