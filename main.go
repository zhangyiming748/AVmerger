package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func init() {
	setLog()
}
func main() {
	defer func() {
		err := os.Chmod("merge.db", 0666)
		err = os.Chmod("AVmerger.log", 0666)
		if err != nil {
			log.Println("修改权限错误")
		}
		//sms.SendMessage()
	}()
	found := false
	if merge.IsExist(constant.BILI) {
		merge.Merge(constant.BILI)
		found = true
	}
	if merge.IsExist(constant.HD) {
		merge.Merge(constant.HD)
		found = true
	}
	if merge.IsExist(constant.GLOBAL) {
		merge.Merge(constant.GLOBAL)
		found = true
	}
	src := strings.Join([]string{getRoot(), "download"}, string(os.PathSeparator))
	if merge.IsExist(src) {
		if !found {
			merge.Merge(src)
		}
	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	log.Printf("当前的工作目录:%v\n", filename)
	util.SetRoot(path.Dir(filename))
	return util.GetRoot()
}

func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   "/data/data/com.termux/files/home/storage/movies/bili/AVmerge.log",
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

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
func NumsOfGoroutine() {
	for {
		fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
