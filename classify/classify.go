package classify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//这个包的功能是用来整理已经merged的文件,按照给定的关键词创建目标文件夹并移动到目标文件夹

/*
srcDir表示merged文件夹路径
dstDir表示目标根文件夹路径
keywords表示 1.要创建的目标文件夹的关键词 2.要移动的文件名包含的关键词
假设:
merged文件夹下有如下文件:
1080P修复张韶涵  呐喊 MV修复版 TGCH张韶涵  呐喊 1080P修复版.mp3
1080P修复萧亚轩  我的男朋友MV 修复版 我的男朋友.mp4
keywords列表里有"张韶涵"
就在dstDir下创建张韶涵文件夹并移动文件
keywords列表里有"萧亚轩"
就在dstDir下创建萧亚轩文件夹并移动文件
*/
func Classify(srcDir, dstDir string, keywords []string) {
	// 遍历源目录中的所有文件
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，跳过目录
		if info.IsDir() {
			return nil
		}

		// 获取文件名
		filename := info.Name()

		// 遍历关键词列表
		for _, keyword := range keywords {
			// 检查文件名是否包含关键词
			if strings.Contains(filename, keyword) {
				// 创建目标目录路径
				targetDir := filepath.Join(dstDir, keyword)

				// 确保目标目录存在
				err := os.MkdirAll(targetDir, os.ModePerm)
				if err != nil {
					fmt.Printf("创建目录失败 %s: %v\n", targetDir, err)
					continue
				}

				// 构造目标文件路径
				targetPath := filepath.Join(targetDir, filename)

				// 移动文件
				err = os.Rename(path, targetPath)
				if err != nil {
					fmt.Printf("移动文件失败 %s 到 %s: %v\n", path, targetPath, err)
				} else {
					fmt.Printf("已移动文件 %s 到 %s\n", filename, targetDir)
				}
				break // 一个文件只移动一次
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}
