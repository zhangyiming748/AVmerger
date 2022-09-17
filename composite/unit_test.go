package composite

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestGetDir(t *testing.T) {
	getDir("/Users/zen/Github/Tools/download")
}

func TestReadEntry(t *testing.T) {
	dir := "/Users/zen/Github/Tools/download/凤凰传奇现场合集/c_748654562/entry.json"
	readEntry(dir)
}
func TestCmd(t *testing.T) {
	//find . -name "*DS_Store*" -exec rm {} \;
	cmd := exec.Command("/bin/bash", "-c", "/Users/zen/workplace/Tools/initialization.sh")
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
}

func TestMulti(t *testing.T) {
	Multi("/Volumes/exchange/file", "/Volumes/exchange/done")
}
func TestSingle(t *testing.T) {
	Single("/Users/zen/Movies/bilibili/single", "/Users/zen/Movies/bilibili/done")
}
