package main

import (
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sql"
	"log/slog"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	constant.SetLogLevel("Info")
	sql.SetEngine()
}

func main() {
	l := len(os.Args)
	var src string
	root := getRoot()
	if l == 1 {
		slog.Info("启动参数", slog.String("1", os.Args[0]))
		src = strings.Join([]string{root, "download"}, string(os.PathSeparator))
	} else if l == 2 {
		slog.Info("启动参数", slog.String("1", os.Args[0]), slog.String("2", os.Args[1]))
		switch os.Args[1] {
		case "bili":
			src = constant.BILI
			constant.SetSecParam(os.Args[1])
		case "hd":
			src = constant.HD
			constant.SetSecParam(os.Args[1])
		case "global":
			src = constant.GLOBAL
			constant.SetSecParam(os.Args[1])
		case "-h", "--help":
			println("go run main.go { bili | hd | global | <path/to/file> |none(./download) }")
			return
		default:
			src = strings.Join([]string{root, "download"}, string(os.PathSeparator))
			slog.Warn("启动参数错误,修改为默认目录", slog.String("src", src))
		}
	}
	merge.Merge(src)
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)

	return path.Dir(filename)
}
