package cover

import (
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
		name := strings.Join([]string{strconv.Itoa(i + 1), ". cover.jpg"}, "")
		target := filepath.Join(dst, name)
		log.Printf("源文件:%v\t移动后:%v\n", cover, target)
		//开始移动文件
		err = os.Rename(cover, target)
		if err != nil {
			log.Printf("移动文件失败：%v -> %v, 错误：%v\n", cover, target, err)
			continue
		}
		log.Printf("移动成功：%v\n", target)
	}
}
