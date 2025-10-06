package AVmerge

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/storage"
)

// init 初始化函数，在程序启动时执行
// 主要完成以下工作：
// 1. 检查系统中是否安装了必要的命令行工具（mediainfo和ffmpeg）
// 2. 初始化日志系统
func init() {
	// 检查 mediainfo 命令是否可用
	// mediainfo用于获取媒体文件的详细信息
	if _, err := exec.LookPath("mediainfo"); err != nil {
		log.Fatal("未找到 mediainfo 命令，请先安装 mediainfo")
	}

	// 检查 ffmpeg 命令是否可用
	// ffmpeg用于音视频文件的处理（合并、转码等）
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("未找到 ffmpeg 命令，请先安装 ffmpeg")
	}
	log.Println("系统环境检查通过: mediainfo 和 ffmpeg 命令可用")
}

// main 程序的主入口函数
// 主要功能：
// 1. 检查并处理各种B站客户端的下载目录
// 2. 合并音视频文件
// 3. 清理处理完成的源文件
// 4. 特殊平台（如Termux、macOS）的额外处理
func Android2PC(mc *MergeConfig) {
	storage.SetMysql(mc.MysqlUser, mc.MysqlPassword, mc.MysqlHost, mc.MysqlPort, "merge")
	storage.GetMysql().Sync2(storage.History{})
	root := filepath.Dir(getRoot())
	src := filepath.Join(root, "download")
	dst := filepath.Join(root, "merged")
	// 处理标准B站客户端的下载目录
	if isExist(src) {
		// 获取目录中的基本信息（音视频文件路径等）
		bs := merge.GetBasicInfo(src)
		// 尝试合并音视频文件
		if merge.Merge(bs, dst) {
			// 合并过程中出现错误，保留源文件目录
			log.Printf("程序有错误,%s目录不会被删除\n", src)
		} else {
			// 合并成功，删除源文件目录
			err := os.RemoveAll(src)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", src)
			}
		}

	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	log.Printf("当前的工作目录:%v\n", filename)
	return filename
}

/*
判断路径是否存在
*/
func isExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("路径存在")
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("路径不存在")
		return false
	} else {
		fmt.Println("发生错误：", err)
		return false
	}
}
