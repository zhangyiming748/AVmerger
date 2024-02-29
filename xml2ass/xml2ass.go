package xml2ass

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/GetFileInfo"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func Conv(info GetFileInfo.BasicInfo) {
	if isExist(constant.ANDROIDDANMAKU) {
		// danmaku2ass danmaku.xml -s 1280x720 -dm 15 -fs 45 -a 50 -o danmaku.ass
		danmaku := strings.Replace(info.FullPath, ".xml", ".ass", 1)
		output, err := exec.Command("danmaku2ass", info.FullPath, "-s", "1280x720", "-dm", "15", "-fs", "45", "-a", "50", "-r", "-o", danmaku).CombinedOutput()
		slog.Debug("生成命令", slog.String("命令原文", fmt.Sprint(danmaku)))
		if err != nil {
			slog.Warn("当前弹幕文件转换错误", slog.Any("文件信息", info))
			return
		} else {
			slog.Info("当前弹幕文件转换成功", slog.Any("命令输出", string(output)))
		}
	} else {
		slog.Info("未找到xml所在目录")
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
