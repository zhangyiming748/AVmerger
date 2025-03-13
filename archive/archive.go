package archive

import (
	"github.com/h2non/filetype"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().Unix()))

// GetAllFiles 获取指定目录及其子目录下的所有文件的绝对路径
func GetAllFiles(rootDir string) ([]string, error) {
	var files []string

	// 使用 filepath.Walk 递归遍历目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录，只处理文件
		if !info.IsDir() {
			// 获取文件的绝对路径
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if isVideo(absPath) {
				log.Printf("video file is %s\n", absPath)
				files = append(files, absPath)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func isVideo(fp string) bool {
	file, err := os.Open(fp)
	if err != nil {
		log.Printf("打开文件失败: %v\n", err)
		return false
	}
	defer file.Close()
	
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Printf("读取文件头失败: %v\n", err)
		return false
	}
	
	return filetype.IsVideo(head)
}

func Convert(file string) error {
	base := filepath.Base(file)
	dir := filepath.Dir(file)
	tmp := strings.Join([]string{strconv.Itoa(seed.Intn(2000)), "mp4"}, ".")
	newPath := filepath.Join(dir, tmp)
	
	cmd := exec.Command("ffmpeg", "-i", file, "-c:v", "libx264", "-tag:v", "hvc1", "-c:a", "aac", newPath)
	log.Printf("base is %v\tdir is %v\ttmp is %v\tnewPath is %v\n", base, dir, tmp, newPath)
	log.Printf("cmd is %v\n", cmd.String())
	
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("转换失败: %s\n", err)
		// 清理可能存在的临时文件
		os.Remove(newPath)
		return err
	} else {
		log.Printf("转换输出: %s\n", out)
		// 创建备份
		backupPath := file + ".bak"
		if err := os.Rename(file, backupPath); err != nil {
			log.Printf("创建备份失败: %s\n", err)
			os.Remove(newPath)
			return err
		}
		
		// 移动新文件到原位置
		if err := os.Rename(newPath, file); err != nil {
			log.Printf("重命名失败: %s\n", err)
			// 恢复原文件
			os.Rename(backupPath, file)
			os.Remove(newPath)
			return err
		}
		
		// 删除备份
		os.Remove(backupPath)
		log.Printf("转换完成: %s\n", file)
	}
	return nil
}
