package main

import (
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sms"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/AVmerger/util"
	"log/slog"
	"os"
	"path"
	"runtime"
)

func init() {
	constant.SetLogLevel("Debug")
	sql.SetEngine()
}

func main() {
	defer func() {
		os.Chmod("merge.db", 0666)
		os.Chmod("AVmerger.log", 0666)
		sms.SendMessage()
	}()
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
	src := "/mnt/e/video/download"
	//src := strings.Join([]string{getRoot(), "download"}, string(os.PathSeparator))
	if merge.IsExist(src) {
		if !found {
			merge.Merge(src)
		}
	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	slog.Info("getRoot", slog.String("当前的工作目录", filename))
	util.SetRoot(path.Dir(filename))
	return util.GetRoot()
}
