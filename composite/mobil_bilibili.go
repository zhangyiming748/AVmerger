package composite

import (
	"AVmerger/log"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
当单集多P和多单集混在一起
按照多单集转换
异常文件名使用
find . -name "*.json" | xargs grep "MV" | tee find.txt
查找
重新使用多P命令转换
*/

type Info struct {
	video string
	audio string
	title string
}
type entry struct {
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

// 批量转换Android端哔哩哔哩下载文件

func getDir(pwd string) (partname []string) {
	//获取文件或目录相关信息
	fileInfoList, err := os.ReadDir(pwd)
	if err != nil {
		log.Debug.Panicln(err)
	}
	for i := range fileInfoList {
		partname = append(partname, fileInfoList[i].Name())
	}
	return partname
}
func readEntry(dir string) (e entry) {
	bytes, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}

	err = json.Unmarshal(bytes, &e)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
	log.Info.Printf("获取到的partname:%s\n", e.PageData.Part)
	log.Info.Printf("获取到的title:%s\n", e.Title)

	return e
}
func merge(infos []Info, dst string) {
	for _, info := range infos {
		command(info.title, info.video, info.audio, dst)
	}
}

// func command(title, dst string) {
func command(title, video, audio, dst string) {
	newname := strings.Join([]string{replace(title), "mp4"}, ".")
	output := strings.Join([]string{dst, newname}, "/")
	cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-codec", "copy", output)
	log.Debug.Printf("生成的命令是:%s", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Panicf("cmd.Run产生的错误:%v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		log.Info.Println(replace(t))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Panicln("命令执行中有错误产生", err)
	}
}

func replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "，", ",", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "_", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "、", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)

	///usr/local/bin/ffmpeg -threads 3 -i download/207257026/c_386723432/80/video.m4s -i download/207257026/c_386723432/80/audio.m4s -codec copy -thread新三国29曹操真是奸诈无比，连自己的发小许攸，都一骗再骗.mp4

	return str
}

// 删除当前目录下的DS_store文件
func rmds(dir string) {
	path := strings.Join([]string{dir, ".DS_store"}, "/")
	os.RemoveAll(path)
}
