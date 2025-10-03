package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/zhangyiming748/AVmerger/client"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/storage"
	"github.com/zhangyiming748/AVmerger/util"
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

	// 初始化日志系统，设置日志文件路径和格式
	util.SetLog()
}

// main 程序的主入口函数
// 主要功能：
// 1. 检查并处理各种B站客户端的下载目录
// 2. 合并音视频文件
// 3. 清理处理完成的源文件
// 4. 特殊平台（如Termux、macOS）的额外处理
func main() {

	// 处理标准B站客户端的下载目录
	if isExist(constant.BILI) {
		// 获取目录中的基本信息（音视频文件路径等）
		bs := merge.GetBasicInfo(constant.BILI)
		// 尝试合并音视频文件
		if merge.Merge(bs) {
			// 合并过程中出现错误，保留源文件目录
			log.Printf("程序有错误,%s目录不会被删除\n", constant.BILI)
		} else {
			// 合并成功，删除源文件目录
			err := os.RemoveAll(constant.BILI)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.BILI)
			}
		}

	}

	// 处理B站HD版（平板）客户端的下载目录
	if isExist(constant.HD) {
		bs := merge.GetBasicInfo(constant.HD)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.HD)
		} else {
			err := os.RemoveAll(constant.HD)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.HD)
			}
		}

	}

	// 处理B站国际版客户端的下载目录
	if isExist(constant.GLOBAL) {
		bs := merge.GetBasicInfo(constant.GLOBAL)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.GLOBAL)
		} else {
			err := os.RemoveAll(constant.GLOBAL)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.GLOBAL)
			}
		}

	}

	// 处理B站海外版客户端的下载目录
	if isExist(constant.BLUE) {
		bs := merge.GetBasicInfo(constant.BLUE)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.BLUE)
		} else {
			err := os.RemoveAll(constant.BLUE)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.BLUE)
			}
		}

	}

	// 处理标准B站客户端在第二存储空间的下载目录
	if isExist(constant.BILI999) {
		bs := merge.GetBasicInfo(constant.BILI999)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.BILI999)
		} else {
			err := os.RemoveAll(constant.BILI999)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.BILI999)
			}
		}

	}

	// 处理B站HD版在第二存储空间的下载目录
	if isExist(constant.HD999) {
		bs := merge.GetBasicInfo(constant.HD999)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.HD999)
		} else {
			err := os.RemoveAll(constant.HD999)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.HD999)
			}
		}

	}

	// 处理B站国际版在第二存储空间的下载目录
	if isExist(constant.GLOBAL999) {
		bs := merge.GetBasicInfo(constant.GLOBAL999)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.GLOBAL999)
		} else {
			err := os.RemoveAll(constant.GLOBAL999)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.GLOBAL999)
			}
		}

	}

	// 处理B站海外版在第二存储空间的下载目录
	if isExist(constant.BLUE999) {
		bs := merge.GetBasicInfo(constant.BLUE999)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.BLUE999)
		} else {
			err := os.RemoveAll(constant.BLUE999)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.BLUE)
			}
		}

	}
	if runtime.GOOS == "darwin" {
		log.Printf("检测到 macOS 系统，开始处理 macOS 相关任务")
		home, _ := os.UserHomeDir()
		root := filepath.Join(home, "Movies", "bilibili")
		// defer func() {
		// 	archiveVideos.ArchiveVideos(root)
		// }()
		if !isExist(root) {
			log.Printf("未找到macos客户端目录%v跳过\n", root)
			return
		}
		if err := client.Convert(root); err != nil {
			log.Println(err)
		} else {
			if err := os.RemoveAll(root); err != nil {
				log.Printf("删除失败%s\n", root)
			} else {
				log.Printf("删除成功%s\n", root)
			}
		}
	}

	if runtime.GOOS == "windows" {
		log.Printf("检测到 windows 系统，开始处理 windows 相关任务")
		storage.SetMysql("root", "163453", "192.168.5.2", "3306", "merge")
		home, _ := os.UserHomeDir()
		root := filepath.Join(home, "Videos", "bilibili")
		//defer func() {
		//	archiveVideos.ArchiveVideos(root)
		//}()
		if !isExist(root) {
			log.Printf("未找到Windows客户端目录%v跳过\n", root)
			return
		}
		if err := client.Convert(root); err != nil {
			log.Println(err)
		} else {
			if err := os.RemoveAll(root); err != nil {
				log.Printf("删除失败%s\n", root)
			} else {
				log.Printf("删除成功%s\n", root)
			}
		}
	}
}

func getRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	log.Printf("当前的工作目录:%v\n", filename)
	util.SetRoot(path.Dir(filename))
	return util.GetRoot()
}

func NumsOfGoroutine() {
	for {
		fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
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
