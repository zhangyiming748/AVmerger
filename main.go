package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sql"
	"io"
	"log/slog"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	setLog()
	sql.SetEngine()
}

func main() {
	l := len(os.Args)
	src := ""
	root := getRoot()
	if l == 1 {
		slog.Info("启动参数", slog.String("1", os.Args[0]))
		src = strings.Join([]string{root, "download"}, string(os.PathSeparator))
	} else if l == 2 {
		slog.Info("启动参数", slog.String("1", os.Args[0]), slog.String("2", os.Args[1]))
		switch os.Args[1] {
		case "bili":
			src = "/sdcard/Android/data/tv.danmaku.bili/download"
		case "hd":
			src = "/sdcard/Android/data/tv.danmaku.bilibilihd/download"
		case "global":
			src = "/sdcard/Android/data/com.bilibili.app.in/download"
		case "-h", "--help":
			println("go run main.go { bili | hd | global | <path/to/file> |none(./download) }")
			return
		default:
			src = strings.Join([]string{root, "download"}, string(os.PathSeparator))
			slog.Warn("启动参数错误,修改为默认目录", slog.String("src", src))
		}
	}
	merge.Merge(src)
}

/*
设置程序运行的日志等级
*/
func setLog() {
	var opt slog.HandlerOptions
	level := "Info"
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
	logf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0770)
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
func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)

	return path.Dir(filename)
}
