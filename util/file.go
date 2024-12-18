package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/pretty"
)

type BasicInfo struct {
	EntryFullPath  string // entry文件所在的绝对路径 /Users/zen/gitea/AVmerge/AVmerger/download/1656385690/c_1624397638/entry.json
	EntryPurgePath string // 文件所在路径 不包含最后一个路径分隔符 /Users/zen/gitea/AVmerge/AVmerger/download/1656385690/c_1624397638
	Video          string // 视频路径
	Audio          string // 音频路径
	Effect         bool   // 判断是否是一个有效目录

}

/*
 */
func GetEntryFilesWithExt(dir, ext string) (bs []BasicInfo, err error) {
	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
			if strings.Contains(path, "entry.json") {
				var b BasicInfo
				b.EntryFullPath = path
				b.EntryPurgePath = filepath.Dir(path)
				source, err := findSingleDirectory(b.EntryPurgePath)
				if err != nil {
					fmt.Println(err)
					b.Effect = false
				} else {
					fmt.Printf("资源文件夹:%v\n", source)
					b.Video = filepath.Join(source, "video.m4s")
					if !isExist(b.Video) || b.Video == "" {
						b.Effect = false
					} else {
						b.Effect = true
					}
					b.Audio = filepath.Join(source, "audio.m4s")
					if !isExist(b.Audio) || b.Audio == "" {
						b.Effect = false
					} else {
						b.Effect = true
					}
				}
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

/*
查询给定目录下唯一的文件夹
*/
// findSingleDirectory returns the absolute path of the single directory in the given root directory.
func findSingleDirectory(root string) (string, error) {
	var singleDir string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root {
			if singleDir != "" {
				return fmt.Errorf("more than one directory found")
			}
			singleDir = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if singleDir == "" {
		return "", fmt.Errorf("no directory found")
	}
	absPath, err := filepath.Abs(singleDir)
	if err != nil {
		return "", err
	}
	return absPath, nil
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
