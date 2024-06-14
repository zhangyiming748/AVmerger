package main

import (
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/AVmerger/util"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	sql.SetEngine()
}

func main() {
	defer func() {
		err := os.Chmod("merge.db", 0666)
		err = os.Chmod("AVmerger.log", 0666)
		if err != nil {
			log.Fatalln("修改权限错误")
		}
		//sms.SendMessage()
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
	src := strings.Join([]string{getRoot(), "download"}, string(os.PathSeparator))
	if merge.IsExist(src) {
		if !found {
			merge.Merge(src)
		}
	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	log.Printf("当前的工作目录:%v\n", filename)
	util.SetRoot(path.Dir(filename))
	return util.GetRoot()
}
