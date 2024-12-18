package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/lumberjack"
)

func init() {
	setLog()
}

func main() {
	defer func() {
		if runtime.GOOS == "android" {
			if videoErr := util.RsyncDir(constant.ANDROIDVIDEO, constant.REMOTEVIDEO, "zen", "127.0.0.1", "163453"); videoErr != nil {
				log.Printf("rsync上传视频失败:%v\n", videoErr)
			}
			if audioErr := util.RsyncDir(constant.ANDROIDAUDIO, constant.REMOTEAUDIO, "zen", "127.0.0.1", "163453"); audioErr != nil {
				log.Printf("rsync上传音频失败:%v\n", audioErr)
			}
		}
	}()
	found := false
	if isExist(constant.BILI) {
		bs := merge.GetBasicInfo(constant.BILI)
		merge.Merge(bs)
		found = true
	}
	if isExist(constant.HD) {
		bs := merge.GetBasicInfo(constant.HD)
		merge.Merge(bs)
		found = true
	}
	if isExist(constant.GLOBAL) {
		bs := merge.GetBasicInfo(constant.GLOBAL)
		merge.Merge(bs)
		found = true
	}
	if isExist(constant.BLUE) {
		bs := merge.GetBasicInfo(constant.BLUE)
		merge.Merge(bs)
		found = true
	}
	src := strings.Join([]string{getRoot(), "download"}, string(os.PathSeparator))
	if isExist(src) {
		if !found {
			bs := merge.GetBasicInfo(src)
			merge.MergeLocal(bs)
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
		Filename:   "AVmerge.log",
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

/*
判断路径是否存在
*/
func isExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("路径存在")
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("路径不存在")
		return false
	} else {
		fmt.Println("发生错误：", err)
		return false
	}
}
