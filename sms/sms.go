package sms

import (
	"log"
	"os/exec"
	"runtime"
)

const NUM = "16710194256"

/*
termux-sms-send -n 18045977133 "程序运行结束"
*/
func SendMessage() {
	if runtime.GOARCH == "arm64" && runtime.GOOS == "android" {
		cmd := exec.Command("termux-sms-send", "-n", NUM, "AVmerger程序运行结束")
		output, err := cmd.CombinedOutput()
		if err != nil {
			//slog.Error("发短信命令执行出错 可能是未安装termux-api")
		}
		log.Println(string(output))
	}
}
