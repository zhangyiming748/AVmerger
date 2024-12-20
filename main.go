package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/lumberjack"
)

func init() {
	setLog()
}

func main() {
	defer func() {
		if runtime.GOOS != "windows" {
			if videoErr := util.RsyncDir(constant.ANDROIDVIDEO, constant.REMOTEVIDEO, "zen", "192.168.1.9", "163453"); videoErr != nil {
				log.Printf("rsync上传视频失败:%v\n", videoErr)
			}
			if audioErr := util.RsyncDir(constant.ANDROIDAUDIO, constant.REMOTEAUDIO, "zen", "192.168.1.9", "163453"); audioErr != nil {
				log.Printf("rsync上传音频失败:%v\n", audioErr)
			}
		}
	}()
	var (
		found bool
	)

	if isExist(constant.BILI) {
		bs := merge.GetBasicInfo(constant.BILI)
		if merge.Merge(bs) {
			log.Printf("程序有错误,%s目录不会被删除\n", constant.BILI)
		} else {
			err := os.RemoveAll(constant.BILI)
			if err != nil {
				log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
			} else {
				log.Printf("程序正确执行,删除文件夹:%v\n", constant.BILI)
			}
		}
		found = true
	}
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
		found = true
	}
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
		found = true
	}
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
		found = true
	}
	src := filepath.Join(getRoot(), "download")
	if isExist(src) {
		if !found {
			bs := merge.GetBasicInfo(src)
			if merge.MergeLocal(bs) {
				log.Printf("程序有错误,%s目录不会被删除\n", src)
			} else {
				err := os.RemoveAll(src)
				if err != nil {
					log.Printf("程序正确执行但删除文件夹失败:%v\n", err)
				} else {
					log.Printf("程序正确执行,删除文件夹:%v\n", src)
				}
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

func setLog() {
	// 创建一个用于写入文件的Logger实例
	local, err := os.UserHomeDir()
	if err != nil {
		local = "AVmerge.log"
		log.Printf("未找到家目录,日志保存到%s\n", local)
	} else {
		local = filepath.Join(local, "AVmerge.log")
		log.Printf("找到家目录,日志保存到%s\n", local)
		defer func() {
			revange(local)
		}()
	}
	fileLogger := &lumberjack.Logger{
		Filename:   local,
		MaxSize:    1, // MB
		MaxBackups: 30,
		MaxAge:     28, // days
	}
	fileLogger.Rotate()
	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
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
func revange(dir string) {
	// 要更改权限和所有者的文件路径
	filePath := dir

	// 设置新的用户和组
	userName := "u0_a331"
	groupName := "media_rw"

	// 获取用户信息
	usr, err := user.Lookup(userName)
	if err != nil {
		fmt.Printf("Error looking up user: %v\n", err)
		return
	}

	// 获取组信息
	group, err := user.LookupGroup(groupName)
	if err != nil {
		fmt.Printf("Error looking up group: %v\n", err)
		return
	}

	// 更改文件的拥有者
	uid, _ := strconv.Atoi(usr.Uid)
	gid, _ := strconv.Atoi(group.Gid)
	if err := os.Chown(filePath, uid, gid); err != nil {
		fmt.Printf("Error changing owner: %v\n", err)
		return
	}

	// 更改文件的权限
	if err := os.Chmod(filePath, 0777); err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
		return
	}

	fmt.Println("Successfully changed owner and permissions.")
}
