package client

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type VideoInfo struct {
	Type           string `json:"type"`
	Codecid        int    `json:"codecid"`
	GroupId        string `json:"groupId"`
	ItemId         int64  `json:"itemId"`
	Aid            int64  `json:"aid"`
	Cid            int64  `json:"cid"`
	Bvid           string `json:"bvid"`
	P              int    `json:"p"`
	TabP           int    `json:"tabP"`
	TabName        string `json:"tabName"`
	Uid            string `json:"uid"`
	Uname          string `json:"uname"`
	Avatar         string `json:"avatar"`
	CoverUrl       string `json:"coverUrl"`
	Title          string `json:"title"`
	Duration       int    `json:"duration"`
	GroupTitle     string `json:"groupTitle"`
	GroupCoverUrl  string `json:"groupCoverUrl"`
	Danmaku        int    `json:"danmaku"`
	View           int    `json:"view"`
	Pubdate        int    `json:"pubdate"`
	Vt             int    `json:"vt"`
	Status         string `json:"status"`
	Active         bool   `json:"active"`
	Loaded         bool   `json:"loaded"`
	Qn             int    `json:"qn"`
	AllowHEVC      bool   `json:"allowHEVC"`
	CreateTime     int64  `json:"createTime"`
	CoverPath      string `json:"coverPath"`
	GroupCoverPath string `json:"groupCoverPath"`
	UpdateTime     int64  `json:"updateTime"`
	TotalSize      int    `json:"totalSize"`
	LoadedSize     int    `json:"loadedSize"`
	Progress       int    `json:"progress"`
	Speed          int    `json:"speed"`
	CompletionTime int64  `json:"completionTime"`
	ReportedSize   int    `json:"reportedSize"`
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
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 到 VideoInfo 结构体
	var videoInfo VideoInfo
	if err := json.Unmarshal(content, &videoInfo); err != nil {
		return nil, err
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
