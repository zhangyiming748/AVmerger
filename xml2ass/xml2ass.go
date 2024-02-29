package xml2ass

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

/*
转换弹幕文件到原始位置
*/
func Conv(xml GetFileInfo.BasicInfo) (string, error) {
	// danmaku2ass danmaku.xml -s 1280x720 -dm 15 -fs 45 -a 50 -o danmaku.ass
	ass := strings.Replace(xml.FullPath, ".xml", ".ass", 1)
	exe := exec.Command("danmaku2ass", xml.FullPath, "-s", "1280x720", "-dm", "15", "-fs", "45", "-a", "50", "-r", "-o", ass)
	output, err := exe.CombinedOutput()
	slog.Debug("生成命令", slog.String("命令原文", fmt.Sprint(exe)))
	if err != nil {
		slog.Warn("当前弹幕文件转换错误", slog.Any("文件信息", xml), slog.Any("错误原文", err))
		return "", err
	} else {
		slog.Info("当前弹幕文件转换成功", slog.Any("命令输出", string(output)))
		return ass, nil
	}
}
func isExist(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
		return false
	} else {
		fmt.Println("文件夹存在")
		return true
	}

}
