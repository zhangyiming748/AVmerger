package convert

import (
	"log"
	"os/exec"
	"testing"
)

func init() {
	// 检查 mediainfo 命令是否可用
	// mediainfo用于获取媒体文件的详细信息
	if _, err := exec.LookPath("mediainfo"); err != nil {
		log.Fatal("未找到 mediainfo 命令，请先安装 mediainfo")
	}

	// 检查 ffmpeg 命令是否可用
	// ffmpeg用于音视频文件的处理（合并、转码等）
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("未找到 ffmpeg 命令，请先安装 ffmpeg")
	}
	log.Println("系统环境检查通过: mediainfo 和 ffmpeg 命令可用")
}
// go test -v -timeout 10h -run TestConvert
func TestConvert(t *testing.T) {
	src := "D:\\Users\\Public\\Videos\\bilibili"
	dst := "D:\\Users\\Public\\Videos\\DONE"
	Convert(src, dst)
}
