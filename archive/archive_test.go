package archive

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/zhangyiming748/AVmerger/util"
)

func init() {
	util.SetLog()
}

// go test -timeout 1000h -v -run TestProcessBilibiliVideos
/*
这个测试用来手动执行哔哩哔哩在电脑上客户端已经转换成功视频的再加工
*/
func TestProcessBilibiliVideos(t *testing.T) {
	home, _ := os.UserHomeDir()
	switch runtime.GOOS {
	case "windows":
		// Windows 系统
		root := filepath.Join(home, "Videos")
		files, _ := GetAllFiles(root)
		for _, file := range files {
			// 清空终端屏幕
			print("\033[H\033[2J")
			Convert(file)
		}
	case "darwin":
		// macOS 系统
		root := filepath.Join(home, "Videos")
		files, _ := GetAllFiles(root)
		for _, file := range files {
			// 清空终端屏幕
			print("\033[H\033[2J")
			Convert(file)
		}
	case "linux":
		// Linux 系统
		root := filepath.Join(home, "Videos")
		files, _ := GetAllFiles(root)
		for _, file := range files {
			// 清空终端屏幕
			print("\033[H\033[2J")
			Convert(file)
		}
	}
}

// go test -timeout 1000h -v -run TestProcessSpecificDirectory
/*
这个测试用来手动执行指定目录已经转换成功视频的再加工
*/
func TestProcessSpecificDirectory(t *testing.T) {
	root := "E:\\Music"
	files, _ := GetAllFiles(root)
	for _, file := range files {
		t.Log(file)
		Convert(file)
	}
}