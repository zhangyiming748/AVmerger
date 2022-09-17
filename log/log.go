package log

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger // 仅打印到屏幕
	Debug *log.Logger // 打印屏幕并保存到文件
)

func init() {
	file := "bilibili.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	Info = log.New(os.Stdout, "INFO:", log.Ltime)
	Debug = log.New(io.MultiWriter(logf, os.Stdout), "DEBUG:", log.Lshortfile)

}
