package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/util"
	"log/slog"
	"os"
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
		} else {
			if compere(file.FullPath, after) == "smaller" {
				err := os.Remove(file.FullPath)
				if err != nil {
					return
				} else {
					slog.Warn(fmt.Sprintf("删除文件%v\n", file.FullPath))
				}
			} else {
				slog.Warn("转换后的文件比源文件更大", slog.String("源文件", file.FullPath), slog.String("目标文件", after), slog.String("命令原文", cmd.String()))
			}

		}
		slog.Debug(fmt.Sprintln("命令输出", string(output)))
	}
}
func compere(before, after string) (result string) {
	file1 := before
	file2 := after

	info1, err1 := os.Stat(file1)
	info2, err2 := os.Stat(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("获取文件信息失败：", err1, err2)
		return
	}

	size1 := info1.Size()
	size2 := info2.Size()

	if size1 > size2 {
		fmt.Printf("%s 文件大小为 %d 字节，大于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "smaller"
	} else if size1 < size2 {
		fmt.Printf("%s 文件大小为 %d 字节，小于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "bigger"
	} else {
		fmt.Printf("%s 文件大小为 %d 字节，等于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "equal"
	}
	return result
}
