package archive

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
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
	// 获取原始文件大小
	originalStat, err := os.Stat(file)
	if err != nil {
		return err
	}
	originalSize := uint64(originalStat.Size())
	originalMB := float64(originalSize) / 1024 / 1024
	log.Printf("原始文件大小: %.2f MB\n", originalMB)

	mi := FastMediaInfo.GetStandMediaInfo(file)
	base := filepath.Base(file)
	dir := filepath.Dir(file)
	tmp := strings.Join([]string{strconv.Itoa(seed.Intn(2000)), "mp4"}, ".")
	newPath := filepath.Join(dir, tmp)
	
	args := []string{"-i", file}
	if mi.Video.Format == "HEVC" {
		if mi.Video.CodecID == "hevc" {
			log.Printf("已经是hevc格式不需要转换")
			return nil
		} else {
			log.Printf("文件已经是hevc格式但没有正确的标签:%s\n", file)
			args = append(args, "-c:v", "copy", "-tag:v", "hvc1", "-c:a", "copy")
		}
	} else {
		log.Printf("文件%v不是hevc格式,开始转换\n", file)
		args = append(args, 
			"-c:v", "libx265",
			"-vf", "minterpolate=fps=60:mi_mode=mci:mc_mode=aobmc:me_mode=bidir:vsbmc=1",
			"-tag:v", "hvc1",
			"-c:a", "aac",
		)
	}
	args = append(args, newPath)
	cmd := exec.Command("ffmpeg", args...)
	
	log.Printf("base is %v\tdir is %v\ttmp is %v\tnewPath is %v\n", base, dir, tmp, newPath)
	log.Printf("cmd is %v\n", cmd.String())
	
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("转换失败: %s\n", err)
		os.Remove(newPath)
		return err
	} else {
		// 获取转换后的文件大小
		newStat, err := os.Stat(newPath)
		if err != nil {
			log.Printf("获取新文件大小失败: %s\n", err)
			os.Remove(newPath)
			return err
		}
		newSize := uint64(newStat.Size())
		newMB := float64(newSize) / 1024 / 1024
		ratio := float64(newSize) / float64(originalSize) * 100
		log.Printf("转换后文件大小: %.2f MB (%.1f%% of original)\n", newMB, ratio)
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
