package main

import (
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/xml2ass/conv"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	constant.SetLogLevel("Info")
	sql.SetEngine()
	conv.GetXmls()
}

func main() {
	found := false
	if merge.IsExist(constant.BILI) {
		merge.Merge(constant.BILI)
		found = true
	}
	if merge.IsExist(constant.HD) {
		merge.Merge(constant.HD)
		found = true
	}
	if merge.IsExist(constant.GLOBAL) {
		merge.Merge(constant.GLOBAL)
		found = true
	}
	src := strings.Join([]string{getRoot(), "download"}, string(os.PathSeparator))
	if merge.IsExist(src) {
		if !found {
			merge.Merge(src)
		}
	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}
