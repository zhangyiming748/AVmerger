package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/util"
	"log/slog"
	"os/exec"
	"strings"
)

func main() {
	constant.SetLogLevel("Debug")
	files, _ := util.GetMKVFilesWithExt("/mnt/f/alist/bilibili/猫咪日常")
	for _, file := range files {
		slog.Debug(fmt.Sprintf("获取到的mkv%+v\n", file))
		after := strings.Replace(file.FullPath, ".mkv", "vp9.mkv", 1)
		cmd := exec.Command("ffmpeg", "-i", file.FullPath, "-map", "0:v:0", "-map", "0:a:0", "-map", "0:s:0", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", after)
		slog.Debug(fmt.Sprintf("命令原文%s\n", cmd.String()))
		output, err := cmd.CombinedOutput()
		if err != nil {
			return
		}
		slog.Debug(fmt.Sprintln("命令输出", string(output)))
	}
}
