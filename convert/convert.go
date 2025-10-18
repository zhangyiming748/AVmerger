package convert

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/zhangyiming748/AVmerger/sqlite"
	"github.com/zhangyiming748/FastMediaInfo"
)

type VideoInfo struct {
	Type           string      `json:"type"`
	Codecid        int         `json:"codecid"`
	GroupId        interface{} `json:"groupId"`
	ItemId         int64       `json:"itemId"`
	Aid            int64       `json:"aid"`
	Cid            int64       `json:"cid"`
	Bvid           string      `json:"bvid"`
	P              int         `json:"p"`
	TabP           int         `json:"tabP"`
	TabName        string      `json:"tabName"`
	Uid            string      `json:"uid"`
	Uname          string      `json:"uname"`
	Avatar         string      `json:"avatar"`
	CoverUrl       string      `json:"coverUrl"`
	Title          string      `json:"title"`
	Duration       int         `json:"duration"`
	GroupTitle     string      `json:"groupTitle"`
	GroupCoverUrl  string      `json:"groupCoverUrl"`
	Danmaku        int         `json:"danmaku"`
	View           int         `json:"view"`
	Pubdate        int         `json:"pubdate"`
	Vt             int         `json:"vt"`
	Status         string      `json:"status"`
	Active         bool        `json:"active"`
	Loaded         bool        `json:"loaded"`
	Qn             int         `json:"qn"`
	AllowHEVC      bool        `json:"allowHEVC"`
	CreateTime     int64       `json:"createTime"`
	CoverPath      string      `json:"coverPath"`
	GroupCoverPath string      `json:"groupCoverPath"`
	UpdateTime     int64       `json:"updateTime"`
	TotalSize      int         `json:"totalSize"`
	LoadedSize     int         `json:"loadedSize"`
	Progress       float64     `json:"progress"`
	Speed          int         `json:"speed"`
	CompletionTime int64       `json:"completionTime"`
	ReportedSize   int         `json:"reportedSize"`
}

func Convert(root, dst string) (err error) {
	files, err := FindVideoInfoFiles(root)
	if err != nil {
		return err
	} else {
		log.Printf("files is %+v\n", files)
	}
	for _, file := range files {
		log.Printf("json file is %s\n", file)
		dir := filepath.Dir(file)
		log.Printf("dir is %s\n", dir)
		media, err := FindM4sFiles(dir)
		if err != nil {
			log.Fatal(err)
		}
		if len(media) != 2 {
			if len(media) > 2 {
				log.Printf("m4s文件多于两个,判断真正需要的两个文件\n")
				for i := len(media) - 1; i >= 0; i-- { // 从后向前遍历
					RemoveEncryptionHeader(media[i])
					mi := FastMediaInfo.GetStandMediaInfo(media[i])
					log.Printf("m4s真实的媒体文件信息为%+v\n", mi)
					if mi.Audio.Format == "" {
						if mi.Video.Format != "HEVC" && mi.Video.Format != "AVC" {
							log.Printf("跳过非 HEVC/AVC 格式的文件: %s", media[i])
							if i == 0 {
								media = media[1:]
							} else if i == len(media)-1 {
								media = media[:len(media)-1]
							} else {
								media = append(media[:i], media[i+1:]...)
							}
						}
					}
				}
			} else {
				log.Printf("m4s文件少于两个,跳过\n")
				continue
			}
		}
		log.Printf("media file is %s\n", media)
		err = RemoveEncryptionHeader(media[0])
		if err != nil {
			return err
		}
		err = RemoveEncryptionHeader(media[1])
		if err != nil {
			return err
		}
		vi, _ := ReadVideoInfo(file)
		log.Printf("videoInfo = %+v\n", vi)
		baseDir := filepath.Join(dst, vi.Uname)
		os.MkdirAll(baseDir, 0755)
		title := strings.Join([]string{vi.Title, "mp4"}, ".")
		h:=new(sqlite.History)
		h.Title = vi.Title
		if has ,_ := h.ExistsByTitle();has{
			log.Printf("已存在%s,跳过\n", title)
			continue
		} 
		
		// if storage.IsDownloaded(key) {
		// 	log.Printf("已存在%s,跳过\n", key)
		// 	continue
		// } else {
		// 	storage.AppendHistory(key)
		// }
		target := filepath.Join(baseDir, title)
		mi1 := FastMediaInfo.GetStandMediaInfo(media[0])
		mi2 := FastMediaInfo.GetStandMediaInfo(media[1])
		args := []string{"-i", media[0], "-i", media[1], "-c:v", "copy", "-c:a", "copy"}
		// args = append(args, "-vf", "minterpolate=fps=60:mi_mode=mci:mc_mode=aobmc:me_mode=bidir:vsbmc=1")
		if mi1.Video.Format == "HEVC" || mi2.Video.Format == "HEVC" {
			log.Printf("视频格式为hevc,添加hvc1")
			args = append(args, "-tag:v", "hvc1")
		}
		{
			title := strings.Join([]string{"title", vi.Title}, "=")
			args = append(args, "-metadata", title)

			artist := strings.Join([]string{"artist", vi.Uname}, "=")
			args = append(args, "-metadata", artist)

			timeStamp := int64(vi.CompletionTime)
			t := time.Unix(timeStamp/1000, 0)
			formattedTime := t.Format("2006-01-02 15:04:05")
			comment := strings.Join([]string{"comment", formattedTime}, "=")
			args = append(args, "-metadata", comment)
		}
		args = append(args, target)
		cmd := exec.Command("ffmpeg", args...)
		log.Printf("开始转换 %s\n", cmd.String())
		out, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		
		h.Insert()
		fmt.Printf("out is %s", out)
		if audio, err := GetMusicFile(media[0], media[1]); err != nil {
			log.Printf("音频转换失败%v\n", err)
		} else {
			mp3 := strings.Replace(target, ".mp4", ".mp3", 1)
			args := []string{"-i", audio, "-c:a", "libmp3lame"}
			{
				title := strings.Join([]string{"title", vi.Title}, "=")
				args = append(args, "-metadata", title)

				artist := strings.Join([]string{"artist", vi.Uname}, "=")
				args = append(args, "-metadata", artist)

				timeStamp := int64(vi.CompletionTime)
				t := time.Unix(timeStamp/1000, 0)
				formattedTime := t.Format("2006-01-02 15:04:05")
				comment := strings.Join([]string{"comment", formattedTime}, "=")
				args = append(args, "-metadata", comment)
			}
			args = append(args, mp3)
			cmd := exec.Command("ffmpeg", args...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
			log.Printf("out is %s", out)
		}
	}
	return nil
}

func FindVideoInfoFiles(rootDir string) ([]string, error) {
	var results []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "videoInfo.json" {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			results = append(results, absPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}

func FindM4sFiles(rootDir string) ([]string, error) {
	var results []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".m4s" {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			results = append(results, absPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}

func ReadVideoInfo(jsonPath string) (*VideoInfo, error) {
	// 打开 JSON 文件
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}

	// 解析 JSON 到 VideoInfo 结构体
	var videoInfo VideoInfo
	if err := json.Unmarshal(content, &videoInfo); err != nil {
		log.Fatalf("无法解析%v\t JSON: %v", jsonPath, err)
	}

	return &videoInfo, nil
}

func RemoveEncryptionHeader(filePath string) error {
	// 打开加密的视频文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// 查找第一个非"30"的位置，并计数
	var cleanContent []byte
	count := 0
	for i := 0; i < len(content); {
		if content[i] == 0x30 {
			count++
			i++
			continue
		}
		cleanContent = content[i:]
		break
	}

	// 打印删除的数量
	log.Printf("成功删除 %d 个'30'字符\n", count)

	// 写回文件
	return os.WriteFile(filePath, cleanContent, 0644)
}

func GetMusicFile(file1, file2 string) (string, error) {
	stat1, err := os.Stat(file1)
	if err != nil {
		return "", err
	}

	stat2, err := os.Stat(file2)
	if err != nil {
		return "", err
	}

	if stat1.Size() <= stat2.Size() {
		return file1, nil
	}
	return file2, nil
}
