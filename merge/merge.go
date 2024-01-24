package merge

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/AVmerger/replace"
	"github.com/zhangyiming748/AVmerger/sql"
	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/GetFileInfo/mediaInfo"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func Merge(rootPath string) {
	entrys := GetFileInfo.GetAllFilesInfo(rootPath, "json")
	for i, entryFile := range entrys {
		if entryFile.FullName == "entry.json" {
			slog.Debug(fmt.Sprintf("正在处理第%d个文件%+v", i+1, entryFile))
			content := getFolder(entryFile.PurgePath)
			video := strings.Join([]string{content, "video.m4s"}, string(os.PathSeparator))
			audio := strings.Join([]string{content, "audio.m4s"}, string(os.PathSeparator))
			jname, sid, errJ := getName(entryFile.FullPath)
			slog.Info(fmt.Sprintf("数据库写入后获取的id = %d", sid))
			jname = replace.ForFileName(jname)
			// 替换连续空格
			jname = strings.Replace(jname, "  ", " ", -1)
			slog.Debug("音视频所在文件夹", slog.String("json文件名", jname), slog.String("音频所在文件夹", audio), slog.String("视频所在文件夹", video))
			vInfo := GetFileInfo.GetFileInfo(video)
			mi, ok := vInfo.MediaInfo.(mediaInfo.VideoInfo)
			if ok {
				slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
			} else {
				slog.Warn("断言视频mediainfo结构体失败")
			}
			slog.Info("WARNING", slog.String("vTAG", mi.VideoCodecID))
			vname := strings.Join([]string{rootPath, string(os.PathSeparator), jname, ".mp4"}, "")
			aname := strings.Join([]string{rootPath, string(os.PathSeparator), jname, ".ogg"}, "")
			cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-c:v", "copy", "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", vname)

			if mi.VideoCodecID == "avc1" {
				cmd = exec.Command("ffmpeg", "-i", video, "-i", audio, "-c:v", "copy", "-c:a", "copy", "-ac", "1", vname)
			}
			ogg := exec.Command("ffmpeg", "-i", audio, "-c:a", "libvorbis", "-ac", "1", aname)
			slog.Debug("音视频所在文件夹", slog.String("json文件名", jname), slog.String("音频所在文件夹", audio), slog.String("视频所在文件夹", video), slog.String("vname", vname), slog.String("cmd", fmt.Sprint(cmd)))
			errV := util.ExecCommand(cmd)
			errA := util.ExecCommand(ogg)
			if errV != nil || errA != nil || errJ != nil {
				slog.Error("最终命令执行出错", slog.String("视频错误", errV.Error()), slog.String("音频错误", errA.Error()), slog.String("json错误", errV.Error()))
				continue
			} else {
				if err := os.RemoveAll(entryFile.PurgePath); err != nil {
					slog.Warn("删除失败", slog.String("要删除的文件夹", entryFile.PurgePath), slog.String("错误原文", err.Error()))
				} else {
					slog.Warn("删除成功")
				}
			}
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

/*
解析并返回文件名和entry原始文件
*/
func getName(jackson string) (name string, id uint64, err error) {
	var entry Entry
	file, err := os.ReadFile(jackson)
	if err != nil {
		return "", 0, err
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
	fmt.Println("return id", record.ID)
	slog.Info("return id", slog.Uint64("id", record.ID))
	if err != nil {
		return "", 0, err
	}
	if entry.PageData.Part == entry.Title {
		name = entry.Title
	} else if entry.PageData.Part == "" || entry.Title == "" {
		var b PlanB
		json.Unmarshal(file, &b)
		index_title := b.Ep.IndexTitle
		name = strings.Join([]string{index_title, entry.Title}, "")
	} else {
		name = strings.Join([]string{entry.PageData.Part, entry.Title}, "")
	}
	name = replace.ForFileName(name)
	slog.Debug("解析之后拼接", slog.String("名称", name))
	return name, record.ID, nil
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

/*
获取指定文件夹下唯一一个文件夹
*/
func getFolder(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	var fname string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			fname = file.Name()
		}
	}

	return strings.Join([]string{dir, fname}, "")
}
