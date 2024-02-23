package util

import (
	"path"
	"runtime"
	"strings"
)

var root string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	r := strings.Replace(path.Dir(filename), "util", "", -1)
	SetRoot(r)
}
func SetRoot(root string) {
	root = root
}
func GetRoot() string {
	return root
}
