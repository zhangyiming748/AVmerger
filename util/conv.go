package util

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

func Conv(xml, ass string) (string, error) {
	// danmaku2ass danmaku.xml -s 1280x720 -dm 15 -fs 45 -a 50 -o danmaku.ass
	bat := strings.Join([]string{"danmaku2ass", xml, "-s", "640x480", "-dm", "15", "-fs", "35", "-r", "-o", ass}, " ")
	cmd := exec.Command("bash", "-c", bat)
	//cmd := exec.Command("danmaku2ass", xml.FullPath, "-s", "1280x720", "-dm", "15", "-fs", "45", "-a", "50", "-r", "-o", ass)
	output, err := cmd.CombinedOutput()
	slog.Debug("生成命令", slog.String("命令原文", fmt.Sprint(cmd)))
	if err != nil {
		slog.Warn("当前弹幕文件转换错误", slog.Any("文件信息", xml), slog.Any("错误原文", err))
		return "", err
	} else {
		slog.Info("当前弹幕文件转换成功", slog.Any("命令输出", string(output)))
		return ass, nil
	}
}
