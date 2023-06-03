package AVmerger

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
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

// todo 目录使用序数词
// /Users/zen/Github/AVmerger/file 包括单p和多p
func AllIn(root string) {
	infos := get(root)
	slog.Debug("解析后", slog.Any("返回的视频", infos))
	for i, info := range *infos {
		slog.Info(fmt.Sprintf("正在合并第 %d/%d 个视频\n", i+1, len(*infos)))
		avc(root, info)
	}
}
func AllInH265(root string) {
	infos := get(root)
	slog.Debug("解析后", slog.Any("返回的视频", infos))
	for i, info := range *infos {
		slog.Info(fmt.Sprintf("正在合并第 %d/%d 个视频", i+1, len(*infos)))
		hevc(root, info)
	}
}

func get(root string) *[]Info {
	var infos []Info
	vs, err := getChildDir(root)
	if err != nil {
		slog.Warn("错误", slog.Any("读取视频根目录", err))
		return nil
	}
	for _, v := range vs {
		rootv := strings.Join([]string{root, v.Name()}, string(os.PathSeparator))
		p, err := getChildDir(rootv)
		if err != nil {
			slog.Warn("错误", slog.Any("读取视频根目录", err))
			return nil
		}
		for _, entry := range p {
			rootvp := strings.Join([]string{rootv, entry.Name()}, string(os.PathSeparator))
			// log.Info.Println(rootvp)
			entry := strings.Join([]string{rootvp, "entry.json"}, string(os.PathSeparator))
			j, err := os.ReadFile(entry)
			if err != nil {
				slog.Warn("错误", slog.Any("读取entry.json文件", err))
				return nil
			}
			var name Entry
			err = json.Unmarshal(j, &name)
			if err != nil {
				slog.Warn("错误", slog.Any("解析entry.json文件", err))
				return nil
			}
			avs, err := getChildDir(rootvp)
			if err != nil {
				slog.Warn("错误", slog.Any("读取分p视频目录", err))
				return nil
			}
			for _, av := range avs {
				audio := strings.Join([]string{rootvp, av.Name(), "audio.m4s"}, string(os.PathSeparator))
				video := strings.Join([]string{rootvp, av.Name(), "video.m4s"}, string(os.PathSeparator))
				info := Info{
					Video: strings.Replace(video, " ", "", -1),
					Audio: strings.Replace(audio, " ", "", -1),
					Name:  strings.Replace(strings.Join([]string{name.Title, name.PageData.Part}, ""), "|", "", -1),
					Del:   rootvp,
				}
				slog.Debug("一个完整视频的基本信息", slog.Any("视频", info.Video), slog.Any("音频", info.Audio), slog.Any("文件名", info.Name), slog.Any("删除后不会影响其他视频的目录", info.Del))
				infos = append(infos, info)
			}
		}
	}
	return &infos
}

/*
获取子目录
*/
func getChildDir(dir string) ([]os.DirEntry, error) {
	var cDir []os.DirEntry
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, child := range readDir {
		if strings.HasPrefix(child.Name(), ".") {
			slog.Info("跳过隐藏文件夹", slog.Any("文件名", child.Name()))
			continue
		}
		if !child.IsDir() {
			// log.Info.Printf("跳过文件:%v\n", child.Name())
			slog.Info("跳过文件", slog.Any("文件名", child.Name()))
			continue
		}
		cDir = append(cDir, child)
	}
	return cDir, err
}

// todo 我有一个业务是对字符串中指定重复字符去重，我看网上都是利用map实现，所以自己写了一个，但是感觉不够优雅，你有更好的方式吗
/*
s: 原字符串
dup: 需要被去重的字符
*/
func Duplicate(s string, dup byte) string {
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
