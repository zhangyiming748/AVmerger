package merge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhangyiming748/AVmerger/replace"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/GetFileInfo/mediaInfo"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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

func Merge(rootPath string) {
	roots := getall(rootPath)
	slog.Debug("根目录", slog.Any("roots", roots))
	for _, root := range roots {
		slog.Info("1", slog.String("1", root))
		secs := getall(root)
		for _, sec := range secs {
			slog.Info("2", slog.String("2", sec))
			entry := strings.Join([]string{sec, "entry.json"}, string(os.PathSeparator))
			name := getName(entry)
			name = CutName(name)
			slog.Info("entry", slog.String("获取到的文件名", name))
			thirds := getall(sec)
			for _, third := range thirds {
				slog.Info("3", slog.String("3", third))
				prefix := util.GetVal("merge", "prefix")

				video := strings.Join([]string{third, "video.m4s"}, string(os.PathSeparator))
				audio := strings.Join([]string{third, "audio.m4s"}, string(os.PathSeparator))

				fname := strings.Join([]string{name, "mp4"}, ".")
				if isExist(prefix) {
					aim := strings.Join([]string{prefix, "bili"}, string(os.PathSeparator))
					os.Mkdir(aim, 0777)
					fname = strings.Join([]string{aim, fname}, string(os.PathSeparator))
				} else {
					slog.Warn("目标文件夹不存在,退出")
					os.Exit(-1)
				}
				if isFileExist(fname) {
					perfix := strings.Replace(fname, ".mp4", "", 1)
					middle := strings.Join([]string{perfix, time.Now().Format("20060102")}, "-")
					fname = strings.Join([]string{middle, "mp4"}, ".")
				}
				slog.Info("命令执行前最终名称", slog.String("文件名", fname), slog.String("视频", video), slog.String("音频", audio))
				vInfo := GetFileInfo.GetFileInfo(video)
				mi, ok := vInfo.MediaInfo.(mediaInfo.VideoInfo)
				if ok {
					slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
				} else {
					slog.Warn("断言视频mediainfo结构体失败")
				}
				slog.Info("WARNING", slog.String("vTAG", mi.VideoCodecID))
				cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-c:v", "copy", "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", fname)

				voiceName := strings.Replace(fname, ".mp4", ".ogg", 1)
				voice := exec.Command("ffmpeg", "-i", audio, "-vn", "-ac", "1", voiceName)
				if VoiceErr := util.ExecCommand(voice); VoiceErr != nil {
					slog.Warn("转换有声书失败")
				}

				if mi.VideoCodecID == "avc1" {
					cmd = exec.Command("ffmpeg", "-i", video, "-i", audio, "-c:v", "copy", "-c:a", "copy", "-ac", "1", fname)
				}
				err := util.ExecCommand(cmd)
				if err != nil {
					slog.Warn("哔哩哔哩合成出错", slog.Any("错误原文", err), slog.Any("命令原文", fmt.Sprint(cmd)))
					continue
				}
			}
		}
	}
	for _, sec := range roots {
		if err := os.RemoveAll(sec); err != nil {
			slog.Warn("删除文件夹失败", slog.Any("错误原文", err))
			if errors.Is(err, errors.New("exit status 1")) {
				fmt.Println("经典错误")
			}
		} else {
			slog.Info("删除文件夹成功")
		}
	}
}
func clean(dir string) {
	delFile := exec.Command("find", dir, "-type", "f", "-exec", "rm", "{}", "\\;").Run()
	fmt.Println("删除文件错误", delFile)
	delFolders := exec.Command("find", dir, "-type", "d", "-exec", "rmdir", "{}", "\\;").Run()
	fmt.Println("删除文件夹错误", delFolders)
}
func isDir(path string) bool {
	fileInfo, _ := os.Stat(path)
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

func getall(rootPath string) (realFolders []string) {
	folders, _ := os.ReadDir(rootPath)
	for _, folder := range folders {
		folderPath := strings.Join([]string{rootPath, folder.Name()}, string(os.PathSeparator))
		if isDir(folderPath) {
			realFolders = append(realFolders, folderPath)
		}
	}
	return realFolders
}

/*
解析并返回文件名和entry原始文件
*/
func getName(jackson string) (name string) {
	var entry Entry
	file, err := os.ReadFile(jackson)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &entry)

	record := new(sql.Bili)
	record.Title = entry.Title
	record.Cover = strings.Replace(entry.Cover, "\\/", "//", -1)
	record.CreatedAt = sql.S2T(strconv.FormatInt(entry.TimeCreateStamp, 10))
	record.UpdatedAt = sql.S2T(strconv.FormatInt(entry.TimeUpdateStamp, 10))
	record.Owner = entry.OwnerName
	record.PartName = entry.PageData.Part
	// https://www.bilibili.com/video/av229337132
	record.AvID = strings.Join([]string{"https://www.bilibili.com/video/av", strconv.Itoa(entry.Avid)}, "")
	// https://www.bilibili.com/video/BV
	record.BvID = strings.Join([]string{"https://www.bilibili.com/video/BV", entry.Bvid}, "")
	record.Original = string(file)
	record.SetOne()

	if err != nil {
		return
	}
	name = strings.Join([]string{entry.PageData.Part, entry.Title}, "")
	name = replace.ForFileName(name)
	slog.Debug("解析之后拼接", slog.String("名称", name))
	return name
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

/*
判断文件是否存在
*/
func isFileExist(fp string) bool {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

/*
截取合理长度的标题
*/
func CutName(before string) (after string) {
	for i, char := range before {
		slog.Debug(fmt.Sprintf("第%d个字符:%v", i+1, string(char)))
		if i >= 124 {
			slog.Debug("截取124之前的完整字符")
			break
		} else {
			after = strings.Join([]string{after, string(char)}, "")
		}
	}
	slog.Debug("截取后", slog.String("before", before), slog.String("after", after))
	return after
}
func kindesOfPrefix() string {
	switch runtime.GOOS {
	case "linux":
		if uname, _ := exec.Command("uname", "-a").CombinedOutput(); strings.Contains(string(uname), "microsoft") {
			return "/mnt/c/Users/zen/Videos"
		}
	case "windows":
		return ""
	case "darwin":
	case "android":
		return "/sdcard/Movies"
	default:
		os.Exit(-1)
	}
	return ""
}
