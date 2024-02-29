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
		//./danmaku2ass.py /Users/zen/Github/AVmerger/downl994070/c_1444525192/danmaku.xml -s 1280x720 -dm 15 -o /Users/zen/Github/AVmerger/download/1600994070/c_1444525192/danmaku.xml
		danmaku := strings.Replace(info.FullPath, ".xml", ".ass", 1)
		output, err := exec.Command("python3", "danmaku2ass.py", info.FullPath, "-s", "1280x720", "-dm", "15", "-o", danmaku).CombinedOutput()
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
