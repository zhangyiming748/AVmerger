package merge

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
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

// todo 添加视频属性的字段
// todo 测试defer 会不会正确写入数据库
func Merge(rootPath string) {
	entrys := GetFileInfo.GetAllFilesInfo(rootPath, "json")
	for i, entryFile := range entrys {
		if entryFile.FullName == "entry.json" {
			mergeOne(i, rootPath, entryFile)
		}
	}
}
func mergeOne(index int, rootPath string, entryFile GetFileInfo.BasicInfo) {
	danmakuAss := strings.Join([]string{entryFile.PurgePath, "danmaku.ass"}, "")
	record := new(sql.Bili)
	defer func() {
		if err := recover(); err != nil {
			record.Success = false
			record.Reason = fmt.Sprint(err)
		} else {
			record.Success = true
		}
		record.SetOne()
	}()
	slog.Debug(fmt.Sprintf("正在处理第%d个文件%+v", index+1, entryFile))
	content := getFolder(entryFile.PurgePath)
	video := strings.Join([]string{content, "video.m4s"}, string(os.PathSeparator))
	audio := strings.Join([]string{content, "audio.m4s"}, string(os.PathSeparator))
	jname, errJ := getName(entryFile.FullPath, record)
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
	var (
		vname string
		aname string
	)
	os.MkdirAll(constant.ANDROIDVIDEO, 0777)
	os.MkdirAll(constant.ANDROIDAUDIO, 0777)
	os.MkdirAll(constant.ANDROIDDANMAKU, 0777)
	vname = strings.Join([]string{constant.ANDROIDVIDEO, string(os.PathSeparator), jname, ".mkv"}, "")
	aname = strings.Join([]string{constant.ANDROIDAUDIO, string(os.PathSeparator), jname, ".aac"}, "")
	cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-i", danmakuAss, "-c:v", "copy", "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", vname)
	record.Format = "hevc"
	if mi.VideoCodecID == "avc1" {
		cmd = exec.Command("ffmpeg", "-i", video, "-i", audio, "-i", danmakuAss, "-c:v", "libx265", "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", vname)
		record.Format = "avc1 to hvc1"
	}
	ogg := exec.Command("ffmpeg", "-i", audio, "-c:a", "aac", "-ac", "1", aname)
	slog.Debug("音视频所在文件夹", slog.String("json文件名", jname), slog.String("音频所在文件夹", audio), slog.String("视频所在文件夹", video), slog.String("vname", vname), slog.String("cmd", fmt.Sprint(cmd)))
	slog.Info("开始写入弹幕")
	//ass文件名和视频一致
	//assName := strings.Replace(vname, ".mp4", ".ass", -1)
	//xml2ass(danmakuXml, assName)
	errV := util.ExecCommand(cmd)
	errA := util.ExecCommand(ogg)
	if errV != nil || errA != nil || errJ != nil {
		slog.Error("最终命令执行出错", slog.String("视频错误", errV.Error()), slog.String("音频错误", errA.Error()), slog.String("json错误", errV.Error()))
		return
	} else {
		if err := os.RemoveAll(entryFile.PurgePath); err != nil {
			slog.Warn("删除失败", slog.String("要删除的文件夹", entryFile.PurgePath), slog.String("错误原文", err.Error()))
		} else {
			slog.Warn("删除成功")
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
func getName(jackson string, record *sql.Bili) (string, error) {
	var entry Entry
	file, err := os.ReadFile(jackson)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(file, &entry)

	//record := new(sql.Bili)
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
	//record.SetOne()
	if err != nil {
		return "", err
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
		record.PartName = index_title
		record.BvID = strings.Join([]string{"https://www.bilibili.com/video/", b.Ep.Bvid}, "")
		record.AvID = strings.Join([]string{"https://www.bilibili.com/video/av", strconv.Itoa(b.Ep.AvId)}, "")
	} else {
		name = strings.Join([]string{entry.Title, entry.PageData.Part}, " ")
	}
	name = replace.ForFileName(name)
	slog.Debug("解析之后拼接", slog.String("名称", name))
	//record.SetOne()
	return name, nil
}

/*
判断路径是否存在
*/
func IsExist(path string) bool {
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
			fname = file.Name()
		}
	}

	return strings.Join([]string{dir, fname}, "")
}

type Danmaku struct {
	Chatserver string `xml:"chatserver"`
	Chatid     int64  `xml:"chatid"`
	Mission    int    `xml:"mission"`
	Maxlimit   int    `xml:"maxlimit"`
	State      int    `xml:"state"`
	RealName   int    `xml:"real_name"`
	Source     string `xml:"source"`
	D          []struct {
		P string `xml:",innerxml"`
	} `xml:"d"`
}

func ReadDanmaku(xmlFile string, record *sql.Bili) {
	file, err := os.ReadFile(xmlFile)
	if err != nil {
		return
	}

	var d Danmaku
	var dans []sql.Danmaku
	xml.Unmarshal(file, &d)
	for _, v := range d.D {
		var dan sql.Danmaku
		dan.AvID = record.AvID
		dan.BvID = record.BvID
		dan.Title = record.Title
		dan.Content = v.P
		dans = append(dans, dan)
	}
	new(sql.Danmaku).SetMany(&dans)
}

func xml2ass(path, name string) {
	//danmaku2ass danmaku.xml -s 1280x720  -dm 15 -o 1.ass
	//assName := strings.Join([]string{name, ".ass"}, "")
	//py := strings.Join([]string{util.GetRoot(), "danmaku2ass.py"}, "")
	cmd := exec.Command("danmaku2ass.py", path, "-s", "1280x720", "-dm", "15", "-o", name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		slog.Warn("字幕转换失败", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", fmt.Sprint(err)))
	}
}
