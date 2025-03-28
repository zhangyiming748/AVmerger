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

// go test -timeout 1000h -v -run TestAll
func TestAll(t *testing.T) {
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
