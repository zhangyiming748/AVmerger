package util

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(c *exec.Cmd) (e error) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("命令运行出现错误", slog.String("命令原文", fmt.Sprint(c)), slog.Any("错误原文", err))
			os.Exit(-1)
		}
	}()
	slog.Info("开始执行命令", slog.String("命令原文", fmt.Sprint(c)))
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		slog.Warn("连接Stdout产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
		return err
	}
	if err = c.Start(); err != nil {
		slog.Warn("启动cmd命令产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
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
		slog.Warn("命令执行中产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
		return err
	}
	return nil
}
