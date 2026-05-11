package merge

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"AVmerger/replace"

	"github.com/zhangyiming748/stand"
)

type BLV struct {
	MediaType                  int         `json:"media_type"`
	HasDashAudio               bool        `json:"has_dash_audio"`
	IsCompleted                bool        `json:"is_completed"`
	TotalBytes                 int64       `json:"total_bytes"`
	DownloadedBytes            int64       `json:"downloaded_bytes"`
	Title                      string      `json:"title"`
	TypeTag                    string      `json:"type_tag"`
	Cover                      string      `json:"cover"`
	VideoQuality               int         `json:"video_quality"`
	PreferedVideoQuality       int         `json:"prefered_video_quality"`
	GuessedTotalBytes          int         `json:"guessed_total_bytes"`
	TotalTimeMilli             int64       `json:"total_time_milli"`
	DanmakuCount               int         `json:"danmaku_count"`
	TimeUpdateStamp            int64       `json:"time_update_stamp"`
	TimeCreateStamp            int64       `json:"time_create_stamp"`
	CanPlayInAdvance           bool        `json:"can_play_in_advance"`
	InterruptTransformTempFile bool        `json:"interrupt_transform_temp_file"`
	QualityPithyDescription    string      `json:"quality_pithy_description"`
	QualitySuperscript         string      `json:"quality_superscript"`
	VariableResolutionRatio    bool        `json:"variable_resolution_ratio"`
	CacheVersionCode           int         `json:"cache_version_code"`
	PreferredAudioQuality      int         `json:"preferred_audio_quality"`
	AudioQuality               int         `json:"audio_quality"`
	Avid                       int64       `json:"avid"`
	Spid                       int         `json:"spid"`
	SeasonID                   int         `json:"season_id"`
	Bvid                       string      `json:"bvid"`
	OwnerID                    int64       `json:"owner_id"`
	OwnerName                  string      `json:"owner_name"`
	IsChargeVideo              bool        `json:"is_charge_video"`
	VerificationCode           int         `json:"verification_code"`
	PageData                   PageData    `json:"page_data"`
	Ep                         interface{} `json:"ep"`
}

type PageData struct {
	Cid              int64  `json:"cid"`
	Page             int    `json:"page"`
	From             string `json:"from"`
	Part             string `json:"part"`
	Link             string `json:"link"`
	RichVid          string `json:"rich_vid"`
	HasAlias         bool   `json:"has_alias"`
	Tid              int    `json:"tid"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	Rotate           int    `json:"rotate"`
	DownloadTitle    string `json:"download_title"`
	DownloadSubtitle string `json:"download_subtitle"`
}

type VideoMeta struct {
	title   string // 视频标题
	srcPath string // 原始flv视频文件绝对路径
	dstPath string // 转码mp4后视频文件绝对路径
}

func Flv(root, dst string) {
	flvs, err := findBlv(root)
	if err != nil {
		log.Printf("查找blv文件路径发生错误:%v\n", err)
	}
	for _, flv := range flvs {
		vm, err := parseEntryInfoFromBlv(flv, dst)
		if err != nil {
			log.Printf("解析blv文件%s发生错误:%v\n", flv, err)
			continue
		}
		log.Printf("正在处理视频：%s\n", vm.title)
		err = flv2mp4(vm)
		if err != nil {
			log.Printf("转换视频%s发生错误:%v\n", vm.srcPath, err)
			continue
		}
	}
}

// findBlv 给定一个root目录,找到所有的blv文件,返回一个blv文件绝对路径的切片
// 相当于find . -name "*.blv"
func findBlv(root string) ([]string, error) {
	var blvFiles []string
	// 递归遍历目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只处理文件，跳过目录
		if !info.IsDir() {
			// 检查文件扩展名是否为.blv
			if filepath.Ext(path) == ".blv" {
				// 转换为绝对路径
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				blvFiles = append(blvFiles, absPath)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return blvFiles, nil
}

// parseEntryInfoFromBlv 根据给定的blv绝对路径,找到对应的entry.json文件并解析
// 例如:/path/to/video/123/lua.mp4.bili2api.80.blv -> 解析 /path/to/video/123/entry.json
func parseEntryInfoFromBlv(blvPath, dst string) (VideoMeta, error) {
	// 获取blv文件所在目录
	dir := filepath.Dir(blvPath)
	// 获取上一级目录
	parentDir := filepath.Dir(dir)
	// 拼接entry.json路径
	entryPath := filepath.Join(parentDir, "entry.json")
	// 读取JSON文件
	data, err := os.ReadFile(entryPath)
	if err != nil {
		return VideoMeta{}, err
	}
	// 解析JSON到BLV结构体
	var blv BLV
	err = json.Unmarshal(data, &blv)
	if err != nil {
		return VideoMeta{}, err
	}
	videoMeta := VideoMeta{
		title:   replace.ForFileName(blv.Title),
		srcPath: blvPath,
		dstPath: filepath.Join(dst, replace.ForFileName(blv.Title)+".mp4"),
	}
	return videoMeta, nil
}

func flv2mp4(videoMeta VideoMeta) error {
	var args []string
	var cmd *exec.Cmd
	args = append(args, "-i", videoMeta.srcPath)
	args = append(args, "-c:v", "libx265")
	args = append(args, "-c:a", "aac")
	args = append(args, "-tag:v", "hvc1")
	args = append(args, videoMeta.dstPath)
	cmd = exec.Command("ffmpeg", args...)
	fmt.Printf("\n命令原文 : %s\n", cmd.String())
	return stand.ExecuteCommandWithRealtimeOutput(cmd)
}
