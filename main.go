package main

import (
	"fmt"
	"github.com/zhangyiming748/AVmerger/client"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/archiveVideos"
	"github.com/zhangyiming748/AVmerger/util"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	// 检查 mediainfo 命令
	if _, err := exec.LookPath("mediainfo"); err != nil {
		log.Fatal("未找到 mediainfo 命令，请先安装 mediainfo")
	}

	// 检查 ffmpeg 命令
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("未找到 ffmpeg 命令，请先安装 ffmpeg")
	}

	log.Println("系统环境检查通过：mediainfo 和 ffmpeg 命令可用")

	util.SetLog()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("程序运行最终收集的panic:%v\n", err)
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
				err1 := util.UploadWithRsync(constant.ANDROIDVIDEO)
				err2 := util.UploadWithRsync(constant.ANDROIDAUDIO)
				if err1 != nil || err2 != nil {
					log.Printf("上传视频文件夹失败:%v\n", err1)
					log.Printf("上传音频文件夹失败:%v\n", err2)
				} else {
					log.Printf("上传文件夹成功:%v\n", constant.BILI)
				}
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
				err1 := util.UploadWithRsync(constant.ANDROIDVIDEO)
				err2 := util.UploadWithRsync(constant.ANDROIDAUDIO)
				if err1 != nil || err2 != nil {
					log.Printf("上传视频文件夹失败:%v\n", err1)
					log.Printf("上传音频文件夹失败:%v\n", err2)
				} else {
					log.Printf("上传文件夹成功:%v\n", constant.BILI)
				}
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
				err1 := util.UploadWithRsync(constant.ANDROIDVIDEO)
				err2 := util.UploadWithRsync(constant.ANDROIDAUDIO)
				if err1 != nil || err2 != nil {
					log.Printf("上传视频文件夹失败:%v\n", err1)
					log.Printf("上传音频文件夹失败:%v\n", err2)
				} else {
					log.Printf("上传文件夹成功:%v\n", constant.BILI)
				}
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
				err1 := util.UploadWithRsync(constant.ANDROIDVIDEO)
				err2 := util.UploadWithRsync(constant.ANDROIDAUDIO)
				if err1 != nil || err2 != nil {
					log.Printf("上传视频文件夹失败:%v\n", err1)
					log.Printf("上传音频文件夹失败:%v\n", err2)
				} else {
					log.Printf("上传文件夹成功:%v\n", constant.BILI)
				}
			}
		}
		found = true
	}
	src := filepath.Join(getRoot(), "downloads")
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

	if runtime.GOOS == "darwin" {
		log.Printf("检测到 macOS 系统，开始处理 macOS 相关任务")
		home, _ := os.UserHomeDir()
		root := filepath.Join(home, "Movies", "bilibili")
		defer func() {
			archiveVideos.ArchiveVideos(root)
		}()
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
		home, _ := os.UserHomeDir()
		root := filepath.Join(home, "Videos", "bilibili")
		defer func() {
			archiveVideos.ArchiveVideos(root)
		}()
		if !isExist(root) {
			log.Printf("未找到linux客户端目录%v跳过\n", root)
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
