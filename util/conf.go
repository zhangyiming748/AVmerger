package util

import (
	"errors"
	"fmt"
	"github.com/zhangyiming748/AVmerger/util/goini"
	"log/slog"
)

const confPath = "./conf.ini"

var (
	conf *goini.Config
)

/*
*
  - 初始化
    init函数的主要作用：
    初始化不能采用初始化表达式初始化的变量。
    程序运行前的注册。
    实现sync.Once功能。
    其他
*/
func init() {
	initConfig()
}

func initConfig() {
	conf = goini.SetConfig(confPath)
	slog.Debug("读取配置文件", slog.String("文件名", confPath))
}

/**
 * 根据键获取值
 */
func GetVal(section, name string) string {
	val, _ := conf.GetValue(section, name)
	return val
}

func SetVal(section, key, value string) error {
	if err := conf.SetValue(section, key, value); err {
		slog.Debug("修改配置文件成功")
		return nil
	} else {
		slog.Warn("修改配置文件失败")
		return errors.New(fmt.Sprintf("修改配置文件失败\tsextion:%s\tkey:%s\tvalue:%s\n", section, key, value))
	}
}
