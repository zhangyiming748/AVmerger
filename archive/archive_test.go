package archive

import (
	"testing"

	"github.com/zhangyiming748/AVmerger/util"
)
func init() {
	util.SetLog()
}
//go test -timeout 1000h -v -run TestAll
func TestAll(t *testing.T) {
	files, _ := GetAllFiles("/Users/zen/Movies")
	for _, file := range files {
		// 清空终端屏幕
		print("\033[H\033[2J")
		// 或者使用这个命令也可以
		// fmt.Print("\033[H\033[2J\033[3J")
		Convert(file)
	}
}