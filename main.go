package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/AVmerger/util"
	"io"
	"log/slog"
	"os"
)

func init() {
	setLog()
	sql.SetEngine()
}
func main() {
	src := util.GetVal("merge", "src")
	bilibilihd := "/sdcard/Android/data/tv.danmaku.bilibilihd"
	bilibili := "/sdcard/Android/data/tv.danmaku.bili"
	if existFolder(bilibilihd) {
		src = bilibilihd
	} else if existFolder(bilibili) {
		src = bilibili
	}
	merge.Merge(src)
}

/*
设置程序运行的日志等级
*/
func setLog() {
	var opt slog.HandlerOptions
	level := util.GetVal("log", "level")
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
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
	file := "AVmerge.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

func existFolder(f string) bool {
	_, err := os.Stat(f)

	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
		return false
	} else if err == nil {
		fmt.Println("文件夹存在")
		return true
	} else {
		fmt.Println("发生错误：", err)
		return false
	}
}
