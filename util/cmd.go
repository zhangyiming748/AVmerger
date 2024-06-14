package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ExecCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Fatalf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Fatalf("启动cmd命令产生错误:%v\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Fatalf("命令执行中产生错误:%v\n", err)
		return err
	}
	return nil
}
