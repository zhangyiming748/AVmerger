package AVmerge

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zhangyiming748/AVmerger/convert"
	"github.com/zhangyiming748/AVmerger/merge"
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

func Client(dst string) {
	OperatingSystem := runtime.GOOS
	var (
		root string
		home string
	)
	home, _ = os.UserHomeDir()
	switch OperatingSystem {
	case "darwin":
		root = filepath.Join(home, "Movies", "bilibili")
	case "linux", "windows":
		root = filepath.Join(home, "Videos", "bilibili")
	default:
		log.Println("不支持的操作系统")
		return
	}
	log.Printf("检测到 %v 系统，开始处理 %v 相关任务\n", OperatingSystem, root)
	if !isExist(root) {
		log.Printf("未找到%v客户端目录%v跳过\n", OperatingSystem, root)
		return
	}
	if err := convert.Convert(root); err != nil {
		log.Println(err)
	} else {
		if err := os.RemoveAll(root); err != nil {
			log.Printf("删除失败%s\n", root)
		} else {
			log.Printf("删除成功%s\n", root)
		}
	}
}
func Android2PC(root string) {
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
