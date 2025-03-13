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
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	if filetype.IsVideo(head) {
		return true
	} else {
		return false
	}
}

func Convert(file string) {
	base := filepath.Base(file)                                              //文件名
	dir := filepath.Dir(file)                                                //路径名
	tmp := strings.Join([]string{strconv.Itoa(seed.Intn(2000)), "mp4"}, ".") //临时文件名 1.mp4
	newPath := filepath.Join(dir, tmp)                                       //输出全路径
	cmd := exec.Command("ffmpeg", "-i", file, "-c:v", "libx264", "-tag:v", "hvc1", "-c:a", "aac", newPath)
	log.Printf("base is %v\tdir is %v\ttmp is %v\tnewPath is %v\n", base, dir, tmp, newPath)
	log.Printf("cmd is %v\n", cmd.String())
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("cmd.CombinedOutput() failed with %s\n", err)
	} else {
		log.Printf("cmd.CombinedOutput() is %s\n", out)
		if err := os.RemoveAll(file); err != nil {
			log.Printf("os.RemoveAll(%s) failed with %s\n", file, err)
		} else {
			log.Printf("os.RemoveAll(%s) success\n", file)
			if err := os.Rename(newPath, file); err != nil {
				log.Printf("os.Rename(%s,%s) failed with %s\n", newPath, file, err)
			} else {
				log.Printf("os.Rename(%s,%s) success\n", newPath, file)
			}
		}
	}
}
