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

// ExecCommand 执行命令并实时输出命令执行结果
// 通过管道获取命令的标准输出，并在执行过程中实时打印输出内容
// 参数:
//   - c: 要执行的命令对象
//
// 返回值:
//   - error: 执行过程中的错误，如果执行成功则返回nil
func ExecCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	// 获取命令的标准输出管道
	stdout, err := c.StdoutPipe()
	// 将标准错误重定向到标准输出
	c.Stderr = c.Stdout
	if err != nil {
		log.Fatalf("连接Stdout产生错误:%v\n", err)
		return err
	}
	// 启动命令
	if err = c.Start(); err != nil {
		log.Fatalf("启动cmd命令产生错误:%v\n", err)
		return err
	}
	// 循环读取并打印命令输出
	for {
		tmp := make([]byte, 1024)
		_, e := stdout.Read(tmp)
		t := string(tmp)
		// 清理输出中的空字符
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Println(t)
		if e != nil {
			break
		}
	}
	// 等待命令执行完成
	if err = c.Wait(); err != nil {
		log.Fatalf("命令执行中产生错误:%v\n", err)
		return err
	}
	return nil
}

// ExecCommandWithBar 执行命令并显示进度条
// 在执行过程中通过解析输出获取处理帧数，并更新进度条显示
// 参数:
//   - c: 要执行的命令对象
//   - totalFrame: 总帧数，用于初始化进度条
//
// 返回值:
//   - error: 执行过程中的错误，如果执行成功则返回nil
func ExecCommandWithBar(c *exec.Cmd, totalFrame string) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	// 将总帧数转换为整数并创建进度条
	total, _ := strconv.Atoi(totalFrame)
	bar := progressbar.New(total)
	defer bar.Finish()
	// 获取命令的标准输出管道
	stdout, err := c.StdoutPipe()
	// 将标准错误重定向到标准输出
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	// 启动命令
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%V\n", err)
		return err
	}
	// 循环读取输出并更新进度条
	for {
		tmp := make([]byte, 1024)
		_, e := stdout.Read(tmp)
		t := string(tmp)
		// 从输出中提取当前帧数并更新进度条
		if frame, none := GetFrameNum(t); none == nil {
			bar.Set(frame)
		}
		if e != nil {
			break
		}
	}
	// 等待命令执行完成
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	log.Printf("命令结束:%v\n", c.String())
	return nil
}

// GetFrameNum 从ffmpeg输出中提取当前处理的帧数
// 使用正则表达式匹配输出中的frame字段
// 参数:
//   - s: ffmpeg命令的输出字符串
//
// 返回值:
//   - int: 提取到的帧数
//   - error: 提取过程中的错误，如果提取成功则返回nil
func GetFrameNum(s string) (int, error) {
	// 使用正则表达式匹配frame=后面的数字
	re := regexp.MustCompile(`frame=\s*(\d+)`)
	matches := re.FindStringSubmatch(s)
	// 如果匹配成功，将匹配到的数字转换为整数返回
	if len(matches) > 1 {
		frameNumber := matches[1]
		frame, _ := strconv.Atoi(frameNumber)
		return frame, nil
	} else {
		return 0, errors.New("not found")
	}
}
