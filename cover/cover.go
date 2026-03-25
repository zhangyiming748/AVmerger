package cover

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
这里实现查找给定src路径下的所有封面图片
类似 find src -type f -name "cover.jpg"
并返回符合条件文件的绝对路径切片
*/
func findAllCover(src string) (covers []string, err error) {
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，跳过目录
		if !info.IsDir() {
			// 检查文件名是否为 cover.jpg
			if strings.ToLower(info.Name()) == "cover.jpg" {
				// 转换为绝对路径
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				covers = append(covers, absPath)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return covers, nil
}

func ArchiveCovers(src, dst string) {
	os.MkdirAll(dst, 0755)
	covers, err := findAllCover(src)
	if err != nil {
		panic(err)
	}
	for i, cover := range covers {
		name := strings.Join([]string{strconv.Itoa(i + 1), ".cover.jpg"}, "")
		target := filepath.Join(dst, name)
		log.Printf("源文件:%v\t移动后:%v\n", cover, target)
		//开始移动文件：先复制再删除，避免跨分区移动导致文件损坏
		err = copyFile(cover, target)
		if err != nil {
			log.Printf("复制文件失败：%v -> %v, 错误：%v\n", cover, target, err)
			continue
		}
		// 复制成功后删除源文件
		err = os.Remove(cover)
		if err != nil {
			log.Printf("删除源文件失败：%v, 错误：%v\n", cover, err)
			// 删除失败不影响继续处理
		}
		log.Printf("移动成功：%v\n", target)
	}
}

// copyFile 复制文件内容从 src 到 dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
