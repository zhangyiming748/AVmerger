package AVmerge

import (
	"io"
	"log"
	"os"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	/*
		日志同时打印到控制台 并且保存到log文件
	*/
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("打开日志文件失败：", err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}

// go test -timeout 24h -v -run TestAndroid2PC
func TestAndroid2PC(t *testing.T) {
	//root := "F:\\alist\\alist\\bilibili"
	//Android2PC(root)
	//Client()
}
