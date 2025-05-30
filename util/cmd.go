package util

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
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
		_, e := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Println(t)
		if e != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Fatalf("命令执行中产生错误:%v\n", err)
		return err
	}
	return nil
}

/*
执行命令过程中可以循环打印消息
*/
func ExecCommandWithBar(c *exec.Cmd, totalFrame string) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	total, _ := strconv.Atoi(totalFrame)
	bar := progressbar.New(total)
	defer bar.Finish()
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%V\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, e := stdout.Read(tmp)
		t := string(tmp)
		if frame, none := GetFrameNum(t); none == nil {
			bar.Set(frame)
		}
		if e != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	log.Printf("命令结束:%v\n", c.String())
	return nil
}

func GetFrameNum(s string) (int, error) {
	re := regexp.MustCompile(`frame=\s*(\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		frameNumber := matches[1]
		frame, _ := strconv.Atoi(frameNumber)
		return frame, nil
	} else {
		return 0, errors.New("not found")
	}
}
