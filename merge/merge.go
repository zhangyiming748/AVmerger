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
	"strings"
	"sync"
	"time"

	"github.com/zhangyiming748/AVmerger/constant"

	"github.com/zhangyiming748/AVmerger/replace"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/FastMediaInfo"
)

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

func Merge(bs []util.BasicInfo) (warning bool) {
	var wg sync.WaitGroup
	for _, b := range bs {
		wg.Add(1)
		fname, subFolder, err := getName(b.EntryFullPath)
		if err != nil {
			log.Printf("文件%v在最终处理文件名的过程中出错%v跳过\n", b.EntryPurgePath, err)
			time.Sleep(10 * time.Second)
			warning = true
			continue
		}
		dir := filepath.Join(constant.ANDROIDVIDEO, subFolder)
		os.MkdirAll(dir, 0777)
		fname = strings.Join([]string{fname, "mp4"}, ".")
		fullName := filepath.Join(dir, fname)
		mp3Dir := filepath.Join(constant.ANDROIDAUDIO, subFolder)
		os.MkdirAll(mp3Dir, 0777)
		mp3Name := strings.Replace(fname, "mp4", "mp3", 1)
		mp3Name = filepath.Join(mp3Dir, mp3Name)
		mi := FastMediaInfo.GetStandMediaInfo(b.Video)
		args := []string{"-i", b.Video, "-i", b.Audio, "-c:v", "copy", "-c:a", "copy", "-map_chapters", "0"}
		if mi.Video.Format == "HEVC" {
			args = append(args, "-tag:v", "hvc1")
		}
		args = append(args, fullName)
		mp4 := exec.Command("ffmpeg", args...)
		mp3 := exec.Command("ffmpeg", "-i", b.Audio, "-c:a", "libmp3lame", mp3Name)
		log.Printf("mp3产生的命令:%s\n", mp3.String())
		go func() {
			defer wg.Done()
			if out, err := mp3.CombinedOutput(); err != nil {
				log.Printf("mp3命令执行输出%s出错:%v\n", out, err) // 使用 Printf 而不是 Panic
			}
		}()
		log.Printf("mp4产生的命令:%s\n", mp4.String())
		frame := FastMediaInfo.GetStandMediaInfo(b.Video).Video.FrameCount
		if err := util.ExecCommandWithBar(mp4, frame); err != nil {
			log.Printf("命令执行失败\n")
			time.Sleep(10 * time.Second)
			warning = true
		}
		if !warning {
			os.RemoveAll(b.Audio)
			os.RemoveAll(b.Video)
			os.RemoveAll(b.EntryFullPath)
		}
	}
	log.Println("等待mp3合并完成")
	ctx, cancel := context.WithCancel(context.Background())
	go NumsOfGoroutine(ctx)
	wg.Wait()
	cancel() // 直接取消上下文
	return warning
}

func MergeLocal(bs []util.BasicInfo) (warning bool) {
	var wg sync.WaitGroup
	wg.Add(len(bs))
	for _, b := range bs {
		fname, subFolder, err := getName(b.EntryFullPath)
		if err != nil {
			log.Printf("文件%v在最终处理文件名的过程中出错%v跳过\n", b.EntryPurgePath, err)
			warning = true
			continue
		}
		dir := subFolder
		os.MkdirAll(dir, 0777)
		fname = strings.Join([]string{fname, "mp4"}, ".")
		fullName := filepath.Join(dir, fname)
		mp3Name := strings.Replace(fullName, ".mp4", ".mp3", 1)

		mp4 := exec.Command("ffmpeg", "-i", b.Video, "-i", b.Audio, "-c:v", "copy", "-c:a", "copy", "-map_chapters", "0", fullName)
		if format := FastMediaInfo.GetStandMediaInfo(b.Video).Video.Format; format == "hvc1" || format == "hevc" {
			log.Printf("视频格式为%s\n", format)
			mp4 = exec.Command("ffmpeg", "-i", b.Video, "-i", b.Audio, "-c:v", "copy", "-tag:v", "hvc1", "-c:a", "copy", "-map_chapters", "0", fullName)
		}
		mp3 := exec.Command("ffmpeg", "-i", b.Audio, "-c:a", "copy", mp3Name)
		go func() {
			defer wg.Done()
			mp3.CombinedOutput()
		}()
		frame := FastMediaInfo.GetStandMediaInfo(b.Video).Video.FrameCount
		if err := util.ExecCommandWithBar(mp4, frame); err != nil {
			log.Printf("命令执行失败\n")
			warning = true
		} else {
			if err := os.RemoveAll(b.EntryPurgePath); err != nil {
				log.Printf("目录%s删除失败\n", b.EntryPurgePath)
			} else {
				log.Printf("目录%s删除成功\n", b.EntryPurgePath)
			}
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	go NumsOfGoroutine(ctx)
	wg.Wait()
	cancel() // 直接取消上下文
	return warning
}

/*
获取文件结构基础信息
*/
func GetBasicInfo(rootPath string) []util.BasicInfo {
	entrys, _ := util.GetEntryFilesWithExt(rootPath, ".json")
	var bs []util.BasicInfo
	for _, entryFile := range entrys {
		if entryFile.Effect {
			bs = append(bs, entryFile)
		}
		log.Printf("开始处理包含entry的文件%+v\n", entryFile)
	}
	return bs
}

/*
解析并返回文件名和owner_name
*/
func getName(jackson string) (string, string, error) {
	var entry Entry
	file, err := os.ReadFile(jackson)
	if err != nil {
		return "", "", err
	}
	err = json.Unmarshal(file, &entry)

	if err != nil {
		return "", "", err
	}

	var name string
	if entry.PageData.Part == entry.Title {
		name = entry.Title
	} else if entry.PageData.Part == "" || entry.Title == "" {
		var b PlanB
		json.Unmarshal(file, &b)
		index_title := b.Ep.IndexTitle
		index := b.Ep.Index
		name = strings.Join([]string{index, entry.Title, index_title}, " ")
	} else {
		name = strings.Join([]string{entry.Title, entry.PageData.Part}, " ")
	}
	name = replace.ForFileName(name)
	//slog.Debug("解析之后拼接", slog.String("名称", name))
	//record.SetOne()
	name = strings.TrimRight(name, " ")
	name = replace.RemoveLeadingSpace(name)
	return name, entry.OwnerName, nil
}

/*
截取合理长度的标题
*/
func CutName(before string) (after string) {
	for i, char := range before {
		//slog.Debug(fmt.Sprintf("第%d个字符:%v", i+1, string(char)))
		if i >= 124 {
			//slog.Debug("截取124之前的完整字符")
			break
		} else {
			after = strings.Join([]string{after, string(char)}, "")
		}
	}
	//slog.Debug("截取后", slog.String("before", before), slog.String("after", after))
	return after
}

func NumsOfGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("NumsOfGoroutine exiting...")
			return
		default:
			fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}
}
