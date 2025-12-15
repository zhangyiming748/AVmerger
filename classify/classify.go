package classify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DefaultKeywords 默认的歌手关键词列表
var DefaultKeywords = []string{
	"张韶涵", "李玟", "萧亚轩", "Tank", "FIR", "孙燕姿", "SHE", "周杰伦", "曹格", "梁静茹",
	"林俊杰", "马天宇", "刘若英", "张栋梁", "潘玮柏", "Sweety", "祖海", "庾澄庆", "凤凰传奇",
	"Katy Perry", "许慧欣", "蔡依林", "郭静", "李玖哲", "群星", "王力宏", "王心凌", "吴克群",
	"张信哲", "Taylor Swift",
}

// 这个包的功能是用来整理已经merged的文件,按照给定的关键词创建目标文件夹并移动到目标文件夹
// 文件会根据类型放置在Audio或Video子目录中，然后按歌手名进一步分类

/*
srcDir表示merged文件夹路径
dstDir表示目标根文件夹路径
keywords表示要匹配的歌手名关键词

处理逻辑：
1. 遍历srcDir目录下的所有文件
2. 根据文件扩展名确定类型：
  - .mp4文件放入dstDir/Video目录
  - .mp3文件放入dstDir/Audio目录

3. 根据文件名中的关键词创建子目录并移动文件：
  - 例如文件"1080P修复张韶涵 呐喊 MV修复版 TGCH张韶涵 呐喊 1080P修复版.mp4"
  - 匹配到关键词"张韶涵"
  - 最终路径为 dstDir/Video/张韶涵/1080P修复张韶涵 呐喊 MV修复版 TGCH张韶涵 呐喊 1080P修复版.mp4

示例:
merged文件夹下有如下文件:
1080P修复张韶涵  呐喊 MV修复版 TGCH张韶涵  呐喊 1080P修复版.mp4
1080P修复萧亚轩  我的男朋友MV 修复版 我的男朋友.mp3
keywords列表里有"张韶涵"
keywords列表里有"萧亚轩"

处理结果：
- 第一个文件会移动到 dstDir/Video/张韶涵/
- 第二个文件会移动到 dstDir/Audio/萧亚轩/
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

		// 获取文件扩展名并确定类型目录
		ext := strings.ToLower(filepath.Ext(filename))
		var typeDir string
		switch ext {
		case ".mp4":
			typeDir = "Video"
		case ".mp3":
			typeDir = "Audio"
		default:
			// 不处理非mp3/mp4文件
			return nil
		}

		// 遍历关键词列表
		for _, keyword := range keywords {
			// 检查文件名是否包含关键词
			if strings.Contains(filename, keyword) {
				// 创建目标目录路径：dstDir/typeDir/keyword
				targetDir := filepath.Join(dstDir, typeDir, keyword)

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
