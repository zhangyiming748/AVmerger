package AVmerge

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zhangyiming748/AVmerger/convert"
	"github.com/zhangyiming748/AVmerger/storage"
)

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

func Client(mc *MergeConfig) {
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
	storage.SetMysql(mc.MysqlUser, mc.MysqlPassword, mc.MysqlHost, mc.MysqlPort, "merge")
	storage.GetMysql().Sync2(storage.History{})
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
