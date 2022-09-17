package main

import (
	"AVmerger/composite"
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
	os.RemoveAll(".DS_store")
}
func main() {
	start := time.Now()
	log.Println("程序开始时间:", time.Now().Format("2006-01-02 15:04:05"))
	defer func() {
		end := time.Now()
		log.Println("程序结束时间:", time.Now().Format("2006-01-02 15:04:05"))
		sub := end.Sub(start)
		log.Println("程序用时:", sub)
	}()
	if len(os.Args) != 4 {
		manual()
		return
	}
	log.Printf("程序名:%s\n功能:%s\n源文件夹:%s\n目标文件夹:%s\n", os.Args[0], os.Args[1], os.Args[2], os.Args[3])
	way := os.Args[1]
	src := os.Args[2]
	dst := os.Args[3]
	if isIllegal(src, dst) {
		manual()
		return
	}
	switch way {
	case "Single":
		composite.Single(src, dst)
	case "Multi":
		composite.Multi(src, dst)
	case "--help":
		manual()
	default:
		manual()
	}
}
func manual() {
	fmt.Println("可选参数")
	fmt.Println("Single|Multi")
	fmt.Println("src")
	fmt.Println("dst")
	fmt.Println("其中Single表示file文件夹中有一个或多个单集,")
	fmt.Println("Multi表示file文件夹中有且只能有一个合集")
	fmt.Println("src为文件所在目录")
	fmt.Println("dst为合成后输出目录")
	fmt.Println("无参数或参数为--help时会显示此说明")
}
func isIllegal(src, dst string) bool {
	if src == dst {
		log.Println("输入输出目录相同\n")
		return true
	}
	if !exists(src) {
		log.Println("src目录不存在\n")
		return true
	}
	if !exists(dst) {
		log.Println("dst目录不存在\n")
		return true
	}
	if !isDir(src) {
		log.Println("src不是目录\n")
		return true
	}
	if !isDir(dst) {
		log.Println("dst不是目录\n")
		return true
	}
	return false
}
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
