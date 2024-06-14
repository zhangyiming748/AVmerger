package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Conv(xml, ass string) (string, error) {
	// danmaku2ass danmaku.xml -s 1280x720 -dm 15 -fs 45 -a 50 -o danmaku.ass
	bat := strings.Join([]string{"danmaku2ass", xml, "-s", "640x480", "-dm", "15", "-fs", "35", "-r", "-o", ass}, " ")
	cmd := exec.Command("bash", "-c", bat)
	//cmd := exec.Command("danmaku2ass", xml.FullPath, "-s", "1280x720", "-dm", "15", "-fs", "45", "-a", "50", "-r", "-o", ass)
	output, err := cmd.CombinedOutput()
	log.Printf("生成命令:%v\n", fmt.Sprint(cmd))
	if err != nil {
		log.Printf("当前弹幕文件转换错误:%v\n", err)
		return "", err
	} else {
		log.Printf("当前弹幕文件转换成功:%v\n", string(output))
		return ass, nil
	}
}
