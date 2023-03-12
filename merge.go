package AVmerger

import (
	"encoding/json"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// 最终生成前的结构体
type Info struct {
	Video string // 视频文件的绝对路径
	Audio string // 音频文件的绝对路径
	Name  string // 最终文件的全名
	Del   string // 标记删除的目录
}
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
		Vid              string `json:"vid"`
		HasAlias         bool   `json:"has_alias"`
		Tid              int    `json:"tid"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Rotate           int    `json:"rotate"`
		DownloadTitle    string `json:"download_title"`
		DownloadSubtitle string `json:"download_subtitle"`
	} `json:"page_data"`
}

// todo 目录使用序数词
// /Users/zen/Github/AVmerger/file 包括单p和多p
func AllIn(root string) {
	infos := get(root)
	log.Debug.Printf("返回的视频:%v\n", infos)
	for i, info := range infos {
		log.Debug.Printf("正在合并第 %d/%d 个视频\n", i+1, len(infos))
		merge(root, info)
	}
}

func get(root string) []Info {
	var infos []Info
	dirs, err := os.ReadDir(root)
	if err != nil {
		log.Warn.Panicf("读取根目录发生错误:%v\n", dirs)
	}
	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), ".") {
			continue
		}
		if !dir.IsDir() {
			log.Info.Printf("跳过文件:%v\n", dir.Name())
			continue
		}
		//log.Debug.Printf("%+v\n", dir.Name())
		afterRoot := strings.Join([]string{root, dir.Name()}, string(os.PathSeparator))
		log.Debug.Printf("获取到的视频目录%+v\n", afterRoot)
		//获取到的视频目录/Users/zen/Github/AVmerger/file/307014787
		seconds, err := os.ReadDir(afterRoot)
		if err != nil {
			log.Warn.Panicf("获取视频目录出错:%v\n", err)
		}
		folderNum := 0
		for i := range seconds {
			log.Debug.Printf("i = %v\n", i)
			folderNum++
		}
		log.Debug.Printf("文件夹数%v\n", folderNum)

		switch folderNum {
		case 1:
			for _, afterSecond := range seconds {
				if afterSecond.IsDir() {
					third := strings.Join([]string{afterRoot, afterSecond.Name()}, string(os.PathSeparator))
					log.Debug.Printf("连接的三级目录:%v\n", third)
					entry := strings.Join([]string{third, "entry.json"}, string(os.PathSeparator))
					js, err := os.ReadFile(entry)
					if err != nil {
						log.Warn.Panicf("读取json文件发生错误%v\n", err)
					}
					var e Entry
					err = json.Unmarshal(js, &e)
					if err != nil {
						log.Warn.Panicf("json反序列化出现问题:%v\n", err)
					}
					//pretty.P(e)
					title := replace.ForFileName(e.Title)
					log.Debug.Printf("单文件的entry标题:%v\n", title)
					afterThird, err := os.ReadDir(third)
					if err != nil {
						log.Warn.Panicf("获取随机数目录出错:%v\n", err)
					}
					for _, fourth := range afterThird {
						if fourth.IsDir() {
							afterFourth := strings.Join([]string{third, fourth.Name()}, string(os.PathSeparator))
							log.Debug.Printf("连接的四级目录:%v\n", afterFourth)
							audio := strings.Join([]string{afterFourth, "audio.m4s"}, string(os.PathSeparator))
							video := strings.Join([]string{afterFourth, "video.m4s"}, string(os.PathSeparator))
							info := Info{
								Video: video,
								Audio: audio,
								Name:  title,
								Del:   third,
							}
							infos = append(infos, info)
						}
					}
				}
			}
		default:
			for _, afterSecond := range seconds {
				if afterSecond.IsDir() {
					third := strings.Join([]string{afterRoot, afterSecond.Name()}, string(os.PathSeparator))
					log.Debug.Printf("连接的三级目录:%v\n", third)
					entry := strings.Join([]string{third, "entry.json"}, string(os.PathSeparator))
					js, err := os.ReadFile(entry)
					if err != nil {
						log.Warn.Panicf("读取json文件发生错误%v\n", err)
					}
					var e Entry
					err = json.Unmarshal(js, &e)
					if err != nil {
						log.Warn.Panicf("json反序列化出现问题:%v\n", err)
					}
					//pretty.P(e)
					title := strings.Join([]string{e.Title, e.PageData.Part}, "-")
					title = replace.ForFileName(title)
					log.Debug.Printf("混合文件的entry标题:%v\n", title)
					afterThird, err := os.ReadDir(third)
					if err != nil {
						log.Warn.Panicf("获取随机数目录出错:%v\n", err)
					}
					for _, fourth := range afterThird {
						if fourth.IsDir() {
							afterFourth := strings.Join([]string{third, fourth.Name()}, string(os.PathSeparator))
							log.Debug.Printf("连接的四级目录:%v\n", afterFourth)
							audio := strings.Join([]string{afterFourth, "audio.m4s"}, string(os.PathSeparator))
							video := strings.Join([]string{afterFourth, "video.m4s"}, string(os.PathSeparator))
							info := Info{
								Video: video,
								Audio: audio,
								Name:  title,
								Del:   third,
							}
							infos = append(infos, info)
						}
					}
				}
			}
		}
	}
	for _, info := range infos {
		pretty.P(info)
	}
	return infos
}

func merge(dst string, info Info) {
	n := duplicate(info.Name, '_')
	n = duplicate(n, '.')
	name := strings.Join([]string{n, "mp4"}, ".")
	target := strings.Join([]string{dst, name}, string(os.PathSeparator))
	var cmd *exec.Cmd
	//cmd := exec.Command("ffmpeg", "-i", info.Video, "-i", info.Audio, target)
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("ffmpeg", "-hwaccel", "videotoolbox", "-i", info.Video, "-i", info.Audio, target)
	default:
		cmd = exec.Command("ffmpeg", "-i", info.Video, "-i", info.Audio, target)
	}
	log.Debug.Printf("生成的命令是%v\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.TTY.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	log.Info.Printf("完成当前文件的处理:源文件是%s\t目标文件夹%s\n", info.Name, dst)
	if err := os.RemoveAll(info.Del); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源目录:%v\n", info.Del)
	}
}

// todo 我有一个业务是对字符串中指定重复字符去重，我看网上都是利用map实现，所以自己写了一个，但是感觉不够优雅，你有更好的方式吗
/*
s: 原字符串
dup: 需要被去重的字符
*/
func duplicate(s string, dup byte) string {
	sb := []byte(s)
	var nb []byte
	for i := 0; i < len(sb); i++ {
		if i == 0 {
			// 如果是第一个字符,直接原样写入字节数组
			nb = append(nb, sb[i])
		} else {
			// 如果不是第一个字符
			if sb[i] == dup && sb[i-1] == dup {
				//如果本身和前一个字符都是dup则跳过
				continue
			} else {
				//否则写入新字节数组
				nb = append(nb, sb[i])
			}
		}
	}
	return string(nb)
}
