package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/pretty"
)

// BasicInfo 存储视频和音频文件的基本信息
// 主要用于处理B站下载的视频文件，包含entry.json文件路径和对应的音视频文件路径
type BasicInfo struct {
	// EntryFullPath 存储entry.json文件的完整路径
	// 例如：/Users/zen/gitea/AVmerge/AVmerger/download/1656385690/c_1624397638/entry.json
	EntryFullPath string

	// EntryPurgePath 存储entry.json文件所在目录的路径（不包含文件名）
	// 例如：/Users/zen/gitea/AVmerge/AVmerger/download/1656385690/c_1624397638
	EntryPurgePath string

	// Video 存储视频文件的完整路径
	Video string

	// Audio 存储音频文件的完整路径
	Audio string

	// Effect 标识当前条目是否有效
	// 当视频和音频文件都存在时为true，否则为false
	Effect bool
}

// GetEntryFilesWithExt 在指定目录中查找特定扩展名的文件，并解析相关的音视频信息
// 主要用于查找B站下载的视频文件，查找entry.json并解析对应的音视频文件路径
// 参数:
//   - dir: 要搜索的目录路径
//   - ext: 要查找的文件扩展名（例如".json"）
//
// 返回值:
//   - []BasicInfo: 包含所有找到的entry文件及其对应音视频文件信息的切片
//   - error: 如果在查找过程中发生错误，返回相应的错误信息
func GetEntryFilesWithExt(dir, ext string) (bs []BasicInfo, err error) {
	var files []string
	// 递归遍历目录
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只处理指定扩展名的文件
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
			// 特别处理entry.json文件
			if strings.Contains(path, "entry.json") {
				var b BasicInfo
				// 设置entry.json的完整路径和所在目录
				b.EntryFullPath = path
				b.EntryPurgePath = filepath.Dir(path)

				// 查找资源文件夹（通常包含音视频文件）
				source, err := findSingleDirectory(b.EntryPurgePath)
				if err != nil {
					fmt.Println(err)
					b.Effect = false
				} else {
					fmt.Printf("资源文件夹:%v\n", source)
					// 构建并验证视频文件路径
					b.Video = filepath.Join(source, "video.m4s")
					if !isExist(b.Video) || b.Video == "" {
						b.Effect = false
					} else {
						b.Effect = true
					}
					// 构建并验证音频文件路径
					b.Audio = filepath.Join(source, "audio.m4s")
					if !isExist(b.Audio) || b.Audio == "" {
						b.Effect = false
					} else {
						b.Effect = true
					}
				}
				// 输出调试信息
				pretty.P(b)
				fmt.Println(b)
				bs = append(bs, b)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// findSingleDirectory 在给定的根目录下查找唯一的子目录
// 主要用于在B站下载的视频目录结构中找到包含音视频文件的目录
// 参数:
//   - root: 要搜索的根目录路径
//
// 返回值:
//   - string: 找到的唯一子目录的绝对路径
//   - error: 如果发生以下情况会返回错误：
//     1. 找到多个子目录
//     2. 没有找到任何子目录
//     3. 目录遍历过程中发生错误
//     4. 转换为绝对路径时发生错误
func findSingleDirectory(root string) (string, error) {
	var singleDir string
	// 遍历根目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只处理目录，且跳过根目录本身
		if info.IsDir() && path != root {
			// 如果已经找到一个目录，说明有多个目录，返回错误
			if singleDir != "" {
				return fmt.Errorf("more than one directory found")
			}
			singleDir = path
		}
		return nil
	})

	// 处理遍历错误
	if err != nil {
		return "", err
	}

	// 处理未找到目录的情况
	if singleDir == "" {
		return "", fmt.Errorf("no directory found")
	}

	// 转换为绝对路径
	absPath, err := filepath.Abs(singleDir)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

// isExist 检查指定路径是否存在
// 参数:
//   - path: 要检查的文件或目录路径
//
// 返回值:
//   - bool: 如果路径存在返回true，否则返回false
//     注意：当发生权限错误等其他错误时也会返回false
func isExist(path string) bool {
	// 尝试获取文件信息来判断路径是否存在
	if _, err := os.Stat(path); err == nil {
		fmt.Println("路径存在")
		return true
	} else if os.IsNotExist(err) {
		// 路径不存在
		fmt.Println("路径不存在")
		return false
	} else {
		// 其他错误（如权限问题）
		fmt.Println("发生错误：", err)
		return false
	}
}
