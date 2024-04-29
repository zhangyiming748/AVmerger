package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BasicInfo struct {
	FullPath  string `json:"full_path,omitempty"`  // 文件的绝对路径
	FullName  string `json:"full_name,omitempty"`  // 文件名
	PurgeName string `json:"purge_name,omitempty"` // 单纯文件名
	PurgeExt  string `json:"purge_ext,omitempty"`  // 单纯扩展名
	PurgePath string `json:"purge_path,omitempty"` // 文件所在路径 不包含最后一个路径分隔符
}

/*
/mnt/e/video/download/1051383511/c_1456495552/80/index.json
/mnt/e/video/download/1051383511/c_1456495552/entry.json
/mnt/e/video/download/1103999468/c_1523388910/80/index.json
/mnt/e/video/download/1103999468/c_1523388910/entry.json
/mnt/e/video/download/1352273496/c_1487186880/80/index.json
/mnt/e/video/download/1352273496/c_1487186880/entry.json
/mnt/e/video/download/1403301913/c_1512244733/80/index.json
/mnt/e/video/download/1403301913/c_1512244733/entry.json
/mnt/e/video/download/1502538820/c_1494026446/80/index.json
*/
func GetEntryFilesWithExt(dir, ext string) (bs []*BasicInfo, err error) {
	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
			if strings.Contains(path, "entry.json") {
				b := new(BasicInfo)
				b.FullPath = path
				b.FullName = filepath.Base(path)
				b.PurgeName = "entry"
				b.PurgeExt = "json"
				b.PurgePath = filepath.Dir(path)
				fmt.Printf("%+v\n", b)
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

func GetMKVFilesWithExt(dir string) (bs []*BasicInfo, err error) {
	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".mkv" {
			files = append(files, path)
			b := new(BasicInfo)
			b.FullPath = path
			b.FullName = filepath.Base(path)
			lastIndex := strings.LastIndex(b.FullName, ".")
			b.PurgeName = b.FullName[:lastIndex]
			b.PurgeExt = b.FullName[lastIndex+1:]
			b.PurgePath = filepath.Dir(path)
			fmt.Printf("%+v\n", b)
			bs = append(bs, b)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return bs, nil
}
