package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/zhangyiming748/lumberjack"
)

// logLocation 存储日志文件的路径
var logLocation string

// GetLog 获取当前日志文件的路径
// 返回值:
//   - string: 日志文件的完整路径
func GetLog() string {
	return logLocation
}

// SetLog 配置日志系统
// 根据不同的操作系统和架构设置日志文件路径
// 配置日志输出格式，同时输出到文件和控制台
func SetLog() {
	var local string
	// 获取当前操作系统和CPU架构信息
	goos := runtime.GOOS
	arch := runtime.GOARCH

	// 根据操作系统和架构确定日志文件路径
	switch goos {
	case "windows", "darwin": // "darwin" 是 macOS 的标识
		local = "AVmerge.log" // 在当前目录创建日志文件
	case "linux":
		switch arch {
		case "amd64":
			local = "AVmerge.log" // x86_64架构使用当前目录
		case "arm64":
			local = "/sdcard/AVmerge.log" // ARM架构使用sdcard目录
		default:
			local = "AVmerge.log"
		}
	case "android":
		local = "/sdcard/AVmerge.log" // Android系统使用sdcard目录
	default:
		local = "AVmerge.log"
	}

	// 保存日志路径并输出
	logLocation = local
	fmt.Println("Local path:", local)

	// 配置文件日志记录器
	fileLogger := &lumberjack.Logger{
		Filename:   local, // 日志文件路径
		MaxSize:    1,     // 每个日志文件最大尺寸，单位MB
		MaxBackups: 1,     // 保留的旧日志文件数量
		MaxAge:     28,    // 保留日志文件的最大天数
	}
	fileLogger.Rotate() // 触发日志轮转

	// 创建控制台日志记录器
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 配置日志输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	// 设置日志格式：显示时间和短文件名
	log.SetFlags(log.Ltime | log.Lshortfile)
}
