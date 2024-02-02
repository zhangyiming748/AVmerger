package constant

import (
	"io"
	"log/slog"
	"os"
)

const (
	BILI         = "/sdcard/Android/data/tv.danmaku.bili/download"
	HD           = "/sdcard/Android/data/tv.danmaku.bilibilihd/download"
	GLOBAL       = "/sdcard/Android/data/com.bilibili.app.in/download"
	ANDROIDVIDEO = "/sdcard/Movies/bili"
	ANDROIDAUDIO = "/sdcard/Music/bili"
)

var LogLevel string

func GetLogLevel() string {
	return LogLevel
}

/*
设置程序运行的日志等级
*/
func SetLogLevel(s string) {
	file := "AVmerge.log"
	logf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0770)
	if err != nil {
		panic(err)
	}
	var opt slog.HandlerOptions
	switch s {
	case "Debug", "debug":
		LogLevel = "Debug"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info", "info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: false,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
		LogLevel = "Info"
	case "Warn", "warn":
		LogLevel = "Err"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err", "err":
		LogLevel = "Err"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

var SecondParameter string

func SetSecParam(s string) {
	SecondParameter = s
}
func GetSecParam() string {
	return SecondParameter
}
