package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/zhangyiming748/lumberjack"
)

var logLocation string

func GetLog() string {
	return logLocation
}
func SetLog() {
	var local string
	goos := runtime.GOOS
	arch := runtime.GOARCH
	switch goos {
	case "windows", "darwin": // "darwin" 是 macOS 的标识
		local = "AVmerge.log"
	case "linux":
		switch arch {
		case "amd64":
			local = "AVmerge.log"
		case "arm64":
			local = "/sdcard/AVmerge.log"
		default:
			local = "AVmerge.log"
		}
	case "android":
		local = "/sdcard/AVmerge.log"
	default:
		local = "AVmerge.log"
	}

	logLocation = local
	fmt.Println("Local path:", local)
	fileLogger := &lumberjack.Logger{
		Filename:   local,
		MaxSize:    1, // MB
		MaxBackups: 30,
		MaxAge:     28, // days
	}
	fileLogger.Rotate()
	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
