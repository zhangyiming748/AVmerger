// Package merge 提供了音视频合并的核心功能，包括从B站下载的音视频文件的合并和本地音视频文件的合并
package merge

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/zhangyiming748/AVmerger/replace"
	"github.com/zhangyiming748/AVmerger/sqlite"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/FastMediaInfo"
)

// Entry 定义了B站视频条目的数据结构，用于解析下载文件的entry.json
type Entry struct {
	MediaType                  int    `json:"media_type"`
	HasDashAudio               bool   `json:"has_dash_audio"`
	IsCompleted                bool   `json:"is_completed"`
	TotalBytes                 int    `json:"total_bytes"`
	DownloadedBytes            int    `json:"downloaded_bytes"`
	Title                      string `json:"title"`
	TypeTag                    string `json:"type_tag"`
	Cover                      string `json:"cover"`
	VideoQuality               int    `json:"video_quality"`
	PreferedVideoQuality       int    `json:"prefered_video_quality"`
	GuessedTotalBytes          int    `json:"guessed_total_bytes"`
	TotalTimeMilli             int    `json:"total_time_milli"`
	DanmakuCount               int    `json:"danmaku_count"`
	TimeUpdateStamp            int64  `json:"time_update_stamp"`
	TimeCreateStamp            int64  `json:"time_create_stamp"`
	CanPlayInAdvance           bool   `json:"can_play_in_advance"`
	InterruptTransformTempFile bool   `json:"interrupt_transform_temp_file"`
	QualityPithyDescription    string `json:"quality_pithy_description"`
	QualitySuperscript         string `json:"quality_superscript"`
	CacheVersionCode           int    `json:"cache_version_code"`
	PreferredAudioQuality      int    `json:"preferred_audio_quality"`
	AudioQuality               int    `json:"audio_quality"`
	Avid                       int    `json:"avid"`
	Spid                       int    `json:"spid"`
	SeasionId                  int    `json:"seasion_id"`
	Bvid                       string `json:"bvid"`
	OwnerId                    int    `json:"owner_id"`
	OwnerName                  string `json:"owner_name"`
	OwnerAvatar                string `json:"owner_avatar"`
	PageData                   struct {
		Cid              int    `json:"cid"`
		Page             int    `json:"page"`
		From             string `json:"from"`
		Part             string `json:"part"`
		Link             string `json:"link"`
		RichVid          string `json:"rich_vid"`
		Vid              string `json:"vid"`
		HasAlias         bool   `json:"has_alias"`
		Weblink          string `json:"weblink"`
		Offsite          string `json:"offsite"`
		Tid              int    `json:"tid"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Rotate           int    `json:"rotate"`
		DownloadTitle    string `json:"download_title"`
		DownloadSubtitle string `json:"download_subtitle"`
	} `json:"page_data"`
}

// PlanB 定义了B站番剧视频条目的备选数据结构，当Entry结构无法满足时使用
type PlanB struct {
	MediaType                  int    `json:"media_type"`
	HasDashAudio               bool   `json:"has_dash_audio"`
	IsCompleted                bool   `json:"is_completed"`
	TotalBytes                 int    `json:"total_bytes"`
	DownloadedBytes            int    `json:"downloaded_bytes"`
	Title                      string `json:"title"`
	TypeTag                    string `json:"type_tag"`
	Cover                      string `json:"cover"`
	VideoQuality               int    `json:"video_quality"`
	PreferedVideoQuality       int    `json:"prefered_video_quality"`
	GuessedTotalBytes          int    `json:"guessed_total_bytes"`
	TotalTimeMilli             int    `json:"total_time_milli"`
	DanmakuCount               int    `json:"danmaku_count"`
	TimeUpdateStamp            int64  `json:"time_update_stamp"`
	TimeCreateStamp            int64  `json:"time_create_stamp"`
	CanPlayInAdvance           bool   `json:"can_play_in_advance"`
	InterruptTransformTempFile bool   `json:"interrupt_transform_temp_file"`
	QualityPithyDescription    string `json:"quality_pithy_description"`
	QualitySuperscript         string `json:"quality_superscript"`
	CacheVersionCode           int    `json:"cache_version_code"`
	PreferredAudioQuality      int    `json:"preferred_audio_quality"`
	AudioQuality               int    `json:"audio_quality"`
	Ep                         struct {
		AvId       int    `json:"av_id"`
		Page       int    `json:"page"`
		Danmaku    int    `json:"danmaku"`
		Cover      string `json:"cover"`
		EpisodeId  int    `json:"episode_id"`
		Index      string `json:"index"`
		IndexTitle string `json:"index_title"`
		From       string `json:"from"`
		SeasonType int    `json:"season_type"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Rotate     int    `json:"rotate"`
		Link       string `json:"link"`
		Bvid       string `json:"bvid"`
		SortIndex  int    `json:"sort_index"`
	} `json:"ep"`
	SeasonId string `json:"season_id"`
}

// Merge 合并从B站下载的音视频文件
// bs: 包含音视频文件基本信息的切片
// 返回warning: 表示处理过程中是否出现警告
func Merge(bs []util.BasicInfo, dst string) (warning bool) {
	// 创建等待组用于同步音频和视频的处理

	// 遍历每个基本信息条目
	for _, b := range bs {
		log.Printf("循环一次开始处理%+v\n", b.EntryFullPath)
		fname, subFolder, _, _ := getName(b.EntryFullPath)
		h := new(sqlite.History)
		h.Title = fname
		if has, err := h.ExistsByTitle(); has {
			log.Printf("已存在%s,跳过\n", fname)
			continue
		}else if err != nil {
			log.Fatalf("查询数据库出现错误:%+v,发生在%s\n", err, fname)
		}

		// 构建视频输出目录路径
		dir := filepath.Join(dst, subFolder)
		// 创建视频输出目录
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Printf("创建目录%s失败:%v\n", dir, err)
		}
		// 添加mp4扩展名
		fname = strings.Join([]string{fname, "mp4"}, ".")
		log.Printf("加入了第一次mp4扩展名处理后的文件名%v\n", fname)
		// 构建完整的视频输出路径
		fullName := filepath.Join(dir, fname)
		// 构建音频输出目录路径
		mp3Dir := filepath.Join(dst, subFolder)
		// 创建音频输出目录
		err = os.MkdirAll(mp3Dir, 0777)
		if err != nil {
			log.Printf("创建目录%s失败:%v\n", mp3Dir, err)
		}
		// 构建音频文件名和完整路径
		mp3Name := strings.Replace(fname, "mp4", "mp3", -1)
		mp3Name = filepath.Join(mp3Dir, mp3Name)
		// 获取视频的媒体信息
		mi := FastMediaInfo.GetStandMediaInfo(b.Video)
		// 构建基本的ffmpeg命令参数：复制视频和音频流，保留章节信息
		args := []string{"-i", b.Video, "-i", b.Audio, "-c:v", "copy", "-c:a", "copy", "-map_chapters", "0"}
		// 如果是HEVC格式，添加hvc1标签以确保兼容性
		if mi.Video.Format == "HEVC" {
			args = append(args, "-tag:v", "hvc1")
		}

		// 添加元数据信息
		{
			// 设置视频标题
			title := strings.Join([]string{"title", fname}, "=")
			args = append(args, "-metadata", title)

			// 设置艺术家（UP主）
			artist := strings.Join([]string{"artist", subFolder}, "=")
			args = append(args, "-metadata", artist)

			// 添加处理时间戳作为注释
			formattedTime := time.Now().Format("2006-01-02 15:04:05")
			comment := strings.Join([]string{"comment", formattedTime}, "=")
			args = append(args, "-metadata", comment)
		}

		// 添加输出文件路径到参数列表
		args = append(args, fullName)
		// 创建视频合并命令
		mp4 := exec.Command("ffmpeg", args...)
		// 创建音频提取命令，使用libmp3lame编码器
		mp3 := exec.Command("ffmpeg", "-i", b.Audio, "-c:a", "libmp3lame", mp3Name)
		log.Printf("mp3产生的命令:%s\n", mp3.String())
		// 在goroutine中异步处理音频提取

		// 执行音频提取命令并处理可能的错误
		if out, err := mp3.CombinedOutput(); err != nil {
			log.Printf("mp3命令执行输出%s出错:%v\n", out, err) // 使用 Printf 而不是 Panic
		}

		log.Printf("mp4产生的命令:%s\n", mp4.String())
		// 获取视频总帧数用于进度显示
		frame := FastMediaInfo.GetStandMediaInfo(b.Video).Video.FrameCount
		// 执行视频合并命令，显示进度条
		if err := util.ExecCommandWithBar(mp4, frame); err != nil {
			log.Printf("命令执行失败\n")
			// 出错时等待10秒并设置警告标志
			time.Sleep(10 * time.Second)
			warning = true
		} else {
			h.Insert()
		}
		// 如果处理成功，清理临时文件
		if !warning {
			// 删除原始音频文件
			os.RemoveAll(b.Audio)
			// 删除原始视频文件
			os.RemoveAll(b.Video)
			// 删除entry.json文件
			os.RemoveAll(b.EntryFullPath)
		}

	}
	log.Println("等待mp3合并完成")

	return warning
}

/*
获取文件结构基础信息
*/
// GetBasicInfo 获取指定目录下所有entry.json文件的基本信息
// rootPath: 根目录路径
// 返回包含所有有效entry文件信息的切片
func GetBasicInfo(rootPath string) []util.BasicInfo {
	// 获取指定目录下所有.json文件
	entrys, _ := util.GetEntryFilesWithExt(rootPath, ".json")
	// 初始化基本信息切片
	var bs []util.BasicInfo
	// 遍历所有entry文件
	for _, entryFile := range entrys {
		// 只处理有效的entry文件
		if entryFile.Effect {
			bs = append(bs, entryFile)
		}
		log.Printf("开始处理包含entry的文件%+v\n", entryFile)
	}
	// 返回所有有效的基本信息
	return bs
}

/*
解析并返回文件名和owner_name
*/
// getName 从entry.json文件中解析并返回规范化的文件名和UP主名称
// jackson: entry.json文件的路径
// 返回值: 文件名, UP主名称, 错误信息
func getName(jackson string) (string, string, string, error) {
	// 定义Entry结构体变量
	var entry Entry
	// 读取entry.json文件内容
	file, err := os.ReadFile(jackson)
	if err != nil {
		return "", "", "", err
	}
	// 解析json内容到Entry结构体
	err = json.Unmarshal(file, &entry)

	if err != nil {
		return "", "", "", err
	}

	// 初始化文件名变量
	var name string
	// 根据不同情况构建文件名
	if entry.PageData.Part == entry.Title {
		// 如果分P标题和主标题相同，直接使用主标题
		name = entry.Title
	} else if entry.PageData.Part == "" || entry.Title == "" {
		// 如果分P标题或主标题为空，尝试使用PlanB结构解析
		var b PlanB
		json.Unmarshal(file, &b)
		index_title := b.Ep.IndexTitle
		index := b.Ep.Index
		// 使用剧集索引、标题和索引标题组合
		name = strings.Join([]string{index, entry.Title, index_title}, " ")
	} else {
		// 使用主标题和分P标题组合
		name = strings.Join([]string{entry.Title, entry.PageData.Part}, " ")
	}
	// 处理文件名中的非法字符
	name = replace.ForFileName(name)
	// 移除右侧空格
	name = strings.TrimRight(name, " ")
	// 移除开头空格
	name = replace.RemoveLeadingSpace(name)
	// 返回处理后的文件名和UP主名称
	var key string
	if key = entry.Bvid; key == "" {
		key = strconv.Itoa(entry.Avid)
	}
	return name, entry.OwnerName, key, nil
}

/*
截取合理长度的标题
*/
// CutName 将过长的标题截断到合适的长度（最大124个字符）
// before: 原始标题
// 返回after: 截断后的标题
func CutName(before string) (after string) {
	// 遍历原始字符串的每个字符
	for i, char := range before {
		// 如果已经达到124个字符的限制
		if i >= 124 {
			// 停止继续添加字符
			break
		} else {
			// 将当前字符添加到结果字符串
			after = strings.Join([]string{after, string(char)}, "")
		}
	}
	// 返回截断后的字符串
	return after
}

// NumsOfGoroutine 监控并打印当前程序运行时的goroutine数量
// ctx: 上下文，用于控制监控的停止
func NumsOfGoroutine(ctx context.Context) {
	// 无限循环监控goroutine数量
	for {
		select {
		// 如果收到取消信号
		case <-ctx.Done():
			fmt.Println("NumsOfGoroutine exiting...")
			return
		// 默认情况下每秒打印一次goroutine数量
		default:
			fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
			// 休眠1秒
			time.Sleep(1 * time.Second)
		}
	}
}
