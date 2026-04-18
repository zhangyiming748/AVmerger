package rename

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/finder"
)

// fp 文件路径
// src 需要被替换的文件名中的子字符串
// dst 需要替换成的文件名中的子字符串
func renameOne(fp, src, dst string) error {
	// 获取文件所在目录和文件名
	dir := filepath.Dir(fp)
	baseName := filepath.Base(fp)

	// 检查文件名中是否包含要替换的字符串
	if !strings.Contains(baseName, src) {
		return nil // 如果不包含，直接返回，不进行重命名
	}

	// 替换文件名中的字符串
	newBaseName := strings.ReplaceAll(baseName, src, dst)

	// 构建新的完整文件路径
	newPath := filepath.Join(dir, newBaseName)
	log.Printf("原文件绝对路径:%v\t重命名后文件绝对路径:%v\n", fp, newPath)
	// 执行重命名操作
	return os.Rename(fp, newPath)
}

func RenameAll(root, src, dst string) {
	files := finder.FindAllFiles(root)
	l := len(files)
	for i, file := range files {
		log.Printf("开始重命名第 %d/%d 个文件: %s", i+1, l, file)
		if err := renameOne(file, src, dst); err != nil {
			log.Printf("重命名文件 %s 失败: %v", file, err)
		}
	}
}