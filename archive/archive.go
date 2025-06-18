package archive

import (
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/FastMediaInfo"
)

// seed 用于生成随机数的种子，基于当前时间戳初始化
// 主要用于生成临时文件名，避免文件名冲突
var seed = rand.New(rand.NewSource(time.Now().Unix()))

// GetAllFiles 获取指定目录及其子目录下的所有视频文件的绝对路径
// 参数:
//   - rootDir: 要扫描的根目录路径
//
// 返回值:
//   - []string: 所有视频文件的绝对路径列表
//   - error: 如果遍历过程中发生错误，返回相应的错误信息
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
			// 检查文件是否为视频文件
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

// isVideo 检查指定文件是否为视频文件
// 通过读取文件头部信息并使用filetype库进行判断
// 参数:
//   - fp: 要检查的文件路径
//
// 返回值:
//   - bool: 如果是视频文件返回true，否则返回false
func isVideo(fp string) bool {
	// 打开文件
	file, err := os.Open(fp)
	if err != nil {
		log.Printf("打开文件失败: %v\n", err)
		return false
	}
	defer file.Close()

	// 读取文件头部261字节用于类型判断
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Printf("读取文件头失败: %v\n", err)
		return false
	}

	// 使用filetype库判断是否为视频文件
	return filetype.IsVideo(head)
}

// Convert 将视频文件转换为HEVC(H.265)格式，并进行帧率优化
// 主要功能：
// 1. 检查并转换视频编码为HEVC格式
// 2. 修正HEVC编码的标签为hvc1
// 3. 根据需要进行帧率插值处理
// 4. 保存转换后的文件并替换原文件
// 参数:
//   - file: 要转换的视频文件路径
//
// 返回值:
//   - error: 如果转换过程中发生错误，返回相应的错误信息
func Convert(file string) error {
	// 获取并记录原始文件大小
	originalStat, err := os.Stat(file)
	if err != nil {
		return err
	}
	originalSize := uint64(originalStat.Size())
	originalMB := float64(originalSize) / 1024 / 1024
	log.Printf("原始文件%s大小: %.2f MB\n", file, originalMB)

	// 获取媒体文件信息
	mi := FastMediaInfo.GetStandMediaInfo(file)

	// 生成临时文件路径
	base := filepath.Base(file)
	dir := filepath.Dir(file)
	tmp := strings.Join([]string{strconv.Itoa(seed.Intn(2000)), "mp4"}, ".")
	newPath := filepath.Join(dir, tmp)

	// 构建ffmpeg命令参数
	args := []string{"-i", file}

	// 根据视频格式决定转换策略
	if mi.Video.Format == "HEVC" {
		if mi.Video.CodecID == "hvc1" {
			// 已经是正确格式和标签的HEVC，无需转换
			log.Printf("已经是hevc格式不需要转换")
			return nil
		} else {
			// HEVC格式但标签不正确，只需修改标签
			log.Printf("文件已经是hevc格式但没有正确的标签:%s\n", file)
			args = append(args, "-c:v", "copy", "-tag:v", "hvc1", "-c:a", "copy")
		}
	} else {
		// 需要转换为HEVC格式
		log.Printf("文件%v不是hevc格式,开始转换\n", file)
		args = append(args,
			"-c:v", "libx265", // 使用x265编码器
			"-tag:v", "hvc1", // 设置正确的HEVC标签
			"-c:a", "aac", // 音频转换为AAC格式
		)

		// 处理帧率
		if fps, err := strconv.ParseFloat(mi.Video.FrameRate, 64); err == nil {
			log.Printf("帧率=%v,尝试插值", fps)
			if fps >= 29.97 {
				log.Printf("帧率大于等于30,不需要插值")
			} else {
				// 对低帧率视频进行插值处理，提升到30fps
				log.Printf("帧率小于30,需要插值")
				args = append(args, "-vf", "minterpolate=fps=30:mi_mode=mci:mc_mode=aobmc:me_mode=bidir:vsbmc=1")
			}
		} else {
			log.Printf("帧率解析失败: %v\n", err)
		}
	}

	// 设置输出文件路径
	args = append(args, newPath)
	cmd := exec.Command("ffmpeg", args...)

	// 输出调试信息
	log.Printf("base is %v\tdir is %v\ttmp is %v\tnewPath is %v\n", base, dir, tmp, newPath)
	log.Printf("cmd is %v\n", cmd.String())

	// 执行转换命令
	if err := util.ExecCommand(cmd); err != nil {
		log.Printf("转换失败: %s\n", err)
		os.Remove(newPath) // 转换失败时删除临时文件
		return err
	} else {
		// 获取并比较转换后的文件大小
		newStat, err := os.Stat(newPath)
		if err != nil {
			log.Printf("获取新文件大小失败: %s\n", err)
			// 这里不删除文件，因为转换可能是成功的
			log.Printf("继续进行文件替换操作\n")
		}
		newSize := uint64(newStat.Size())
		newMB := float64(newSize) / 1024 / 1024
		ratio := float64(newSize) / float64(originalSize) * 100
		log.Printf("转换后文件大小: %.2f MB (%.1f%% of original)\n", newMB, ratio)
		savedSize := float64(originalSize-newSize) / 1024 / 1024
		log.Printf("节省空间: %.2f MB\n", savedSize)

		// 替换原文件
		if err := os.Rename(newPath, file); err != nil {
			// 重命名失败时，尝试复制后删除的方式
			if err := copyFile(newPath, file); err != nil {
				log.Printf("复制文件失败: %s\n", err)
				os.Remove(newPath)
				return err
			}
			// 复制成功后删除临时文件
			os.Remove(newPath)
		}
		log.Printf("转换完成: %s\n", file)
	}
	return nil
}

// copyFile 将源文件复制到目标路径
// 主要用于在重命名操作失败时提供备选的文件替换方案
// 复制过程包括：
// 1. 打开源文件
// 2. 删除已存在的目标文件（如果有）
// 3. 创建新的目标文件
// 4. 复制文件内容
// 5. 确保数据写入磁盘
// 参数:
//   - src: 源文件路径
//   - dst: 目标文件路径
//
// 返回值:
//   - error: 如果复制过程中发生错误，返回相应的错误信息
func copyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 删除已存在的目标文件
	os.Remove(dst)

	// 创建新的目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// 确保数据写入磁盘
	err = destFile.Sync()
	return err
}
