package storage

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

var history = make(map[string]struct{})

/* 读取历史文件到 map 如果调用这个函数的时候 root变量为空 则使用程序运行根目录下的 history.txt 文件
 * root为history文件所在的目录 不必加上 文件名
 */
func SetHistory2Map(root string) {
	// 确定历史文件路径
	historyFilePath := findHistoryFile()
	// 读取文件内容到 history map 中
	file, err := os.Open(historyFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			history[line] = struct{}{}
		}
	}
}

/*
返回整个的历史文件的map
*/
func GetHistory4Map() map[string]struct{} {
	return history
}

/*
把新增的内容追加写回history.txt 文件
*/
func AppendHistory(key string) {
	// 如果key已经存在，直接返回
	if _, exists := history[key]; exists {
		return
	}
	// 将新的key添加到内存map中
	history[key] = struct{}{}
	// 追加写入到文件
	historyFilePath := findHistoryFile()
	// 以追加模式打开文件
	file, err := os.OpenFile(historyFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("追加处理过的历史文件函数中打开文件发生错误: %v\n", err)
	}
	defer file.Close()
	// 追加新行
	_, err = file.WriteString(key + "\n")
	if err != nil {
		log.Fatalf("追加处理过的历史文件函数中写入文件发生错误: %v\n", err)
	}
}

/*
判断关键词是否已经存在于map
*/
func IsDownloaded(key string) bool {
	if _, ok := history[key]; ok {
		return true
	}
	return false
}

/*
*
实现查找程序运行目录下(任意层级)存在的第一个 文件名完全匹配 history.txt 的文件 返回绝对路径
类似于linux下find . -name "history.txt" -print
但是我这个代码要实现跨平台linux macos Windows 通用
*/
func findHistoryFile() string {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取处理过的历史文件函数中获取当前工作目录发生错误: %v\n", err)
	}
	// 使用filepath.Walk遍历目录树
	var result string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 继续遍历其他路径
		}
		// 如果找到名为history.txt的文件
		if info.Name() == "history.txt" {
			result = path
			return filepath.SkipDir // 跳过当前目录的遍历
		}
		return nil
	})
	// 返回绝对路径
	if result != "" {
		if absPath, err := filepath.Abs(result); err == nil {
			return absPath
		}
	}
	// 如果最后都没有找到，则在工作目录下创建history.txt并返回
	h := filepath.Join(dir, "history.txt")
	// 检查文件是否已经存在（双重检查）
	if _, err := os.Stat(h); os.IsNotExist(err) {
		file, err := os.Create(h)
		if err != nil {
			log.Fatalf("获取处理过的历史文件函数中在整个目录都不存在history.txt文件时创建新的history.txt文件发生错误: %v\n", err)
		}
		file.Close()
	}

	// 返回绝对路径
	absPath, err := filepath.Abs(h)
	if err != nil {
		log.Fatalf("获取处理过的历史文件函数中在整个目录都不存在history.txt文件创建新的history.txt文件后转换为绝对路径时发生错误: %v\n", err)
	}
	return absPath
}
