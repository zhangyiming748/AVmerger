package AVmerger

import (
	"io"
	"log"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/zhangyiming748/lumberjack"
)

func init() {
	SetLog("avmerge.log")
}

func TestMerge(t *testing.T) {
	src := "C:\\Users\\zen\\Videos\\download"
	dst := "C:\\Users\\zen\\Videos\\done"
	archive := "C:\\Users\\zen\\Videos\\archive"
	if runtime.GOOS == "android" && runtime.GOARCH == "arm64" {
		src = "/sdcard/bilibili/download"
		dst = "/sdcard/bilibili/done"
		archive = "/sdcard/bilibili/archive"
	}
	//执行操作之前的目录情况
	log.Printf("src = %v\tdst = %v\tarchive = %v\n", src, dst, archive)
	//AVmerger.Client(src, dst)
	Android2PC(src, dst)
	// folders := finder.FindAllFolders(dst)
	// for i, folder := range folders {
	// 	log.Printf("正在处理第%d/%d个文件夹:%s\n", i+1, len(folders), folder)
	// 	vFiles := finder.FindAllVideosInRoot(folder)
	// 	for j, vFile := range vFiles {
	// 		log.Printf("正在处理第%d/%d个文件夹:%s中的第个%d/%d文件:%s\n", i+1, len(folders), folder, j+1, len(vFiles), vFile)
	// 		archive.Convert2H265(vFile)
	// 	}
	// }
	ClassifyAfterMerge(dst, archive, nil)
}

func SetLog(l string) {
	// 设置全局时区为Asia/Shanghai
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Printf("无法加载时区 Asia/Shanghai: %v", err)
	} else {
		time.Local = location
	}
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   l,
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
	}
	err = fileLogger.Rotate()
	if err != nil {
		log.Println("转换新日志文件失败", err)
	}
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.Ltime | log.Lshortfile)
}
