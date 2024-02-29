package util

import (
	"log/slog"
	"path"
	"runtime"
	"strings"
)

var root string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	r := strings.Replace(path.Dir(filename), "util", "", -1)
	slog.Info("设置根目录", slog.String("根目录", r))
	SetRoot(r)
}
func SetRoot(r string) {
	root = r
}
func GetRoot() string {
	return root
}
