package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type MediaInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                     string `json:"@type"`
			Count                    string `json:"Count"`
			StreamCount              string `json:"StreamCount"`
			StreamKind               string `json:"StreamKind"`
			StreamKindString         string `json:"StreamKind_String"`
			StreamKindID             string `json:"StreamKindID"`
			VideoCount               string `json:"VideoCount,omitempty"`
			VideoFormatList          string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList  string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList           string `json:"Video_Codec_List,omitempty"`
			CompleteName             string `json:"CompleteName,omitempty"`
			FileNameExtension        string `json:"FileNameExtension,omitempty"`
			FileName                 string `json:"FileName,omitempty"`
			FileExtension            string `json:"FileExtension,omitempty"`
			Format                   string `json:"Format"`
			FormatString             string `json:"Format_String"`
			FormatExtensions         string `json:"Format_Extensions,omitempty"`
			FormatCommercial         string `json:"Format_Commercial"`
			FormatProfile            string `json:"Format_Profile"`
			InternetMediaType        string `json:"InternetMediaType"`
			CodecID                  string `json:"CodecID"`
			CodecIDString            string `json:"CodecID_String,omitempty"`
			CodecIDUrl               string `json:"CodecID_Url"`
			CodecIDCompatible        string `json:"CodecID_Compatible,omitempty"`
			FileSize                 string `json:"FileSize,omitempty"`
			FileSizeString           string `json:"FileSize_String,omitempty"`
			FileSizeString1          string `json:"FileSize_String1,omitempty"`
			FileSizeString2          string `json:"FileSize_String2,omitempty"`
			FileSizeString3          string `json:"FileSize_String3,omitempty"`
			FileSizeString4          string `json:"FileSize_String4,omitempty"`
			Duration                 string `json:"Duration"`
			DurationString           string `json:"Duration_String"`
			DurationString1          string `json:"Duration_String1"`
			DurationString2          string `json:"Duration_String2"`
			DurationString3          string `json:"Duration_String3"`
			DurationString4          string `json:"Duration_String4"`
			DurationString5          string `json:"Duration_String5"`
			OverallBitRate           string `json:"OverallBitRate,omitempty"`
			OverallBitRateString     string `json:"OverallBitRate_String,omitempty"`
			FrameRate                string `json:"FrameRate"`
			FrameRateString          string `json:"FrameRate_String"`
			FrameCount               string `json:"FrameCount"`
			StreamSize               string `json:"StreamSize"`
			StreamSizeString         string `json:"StreamSize_String"`
			StreamSizeString1        string `json:"StreamSize_String1"`
			StreamSizeString2        string `json:"StreamSize_String2"`
			StreamSizeString3        string `json:"StreamSize_String3"`
			StreamSizeString4        string `json:"StreamSize_String4"`
			StreamSizeString5        string `json:"StreamSize_String5"`
			StreamSizeProportion     string `json:"StreamSize_Proportion"`
			HeaderSize               string `json:"HeaderSize,omitempty"`
			DataSize                 string `json:"DataSize,omitempty"`
			FooterSize               string `json:"FooterSize,omitempty"`
			IsStreamable             string `json:"IsStreamable,omitempty"`
			Description              string `json:"Description,omitempty"`
			FileModifiedDate         string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal    string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication       string `json:"Encoded_Application,omitempty"`
			EncodedApplicationString string `json:"Encoded_Application_String,omitempty"`
			Extra                    struct {
				FileExtensionInvalid  string `json:"FileExtension_Invalid,omitempty"`
				CodecConfigurationBox string `json:"CodecConfigurationBox,omitempty"`
			} `json:"extra"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			IDString                       string `json:"ID_String,omitempty"`
			FormatInfo                     string `json:"Format_Info,omitempty"`
			FormatUrl                      string `json:"Format_Url,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatTier                     string `json:"Format_Tier,omitempty"`
			CodecIDInfo                    string `json:"CodecID_Info,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			BitRateString                  string `json:"BitRate_String,omitempty"`
			Width                          string `json:"Width,omitempty"`
			WidthString                    string `json:"Width_String,omitempty"`
			Height                         string `json:"Height,omitempty"`
			HeightString                   string `json:"Height_String,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString       string `json:"DisplayAspectRatio_String,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString            string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
			FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
			FrameRateMinimum               string `json:"FrameRate_Minimum,omitempty"`
			FrameRateMinimumString         string `json:"FrameRate_Minimum_String,omitempty"`
			FrameRateMaximum               string `json:"FrameRate_Maximum,omitempty"`
			FrameRateMaximumString         string `json:"FrameRate_Maximum_String,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString        string `json:"ChromaSubsampling_String,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			BitDepthString                 string `json:"BitDepth_String,omitempty"`
			BitsPixelFrame                 string `json:"BitsPixel_Frame,omitempty"`
			ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
			ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
			ColourRange                    string `json:"colour_range,omitempty"`
			ColourRangeSource              string `json:"colour_range_Source,omitempty"`
			ColourPrimaries                string `json:"colour_primaries,omitempty"`
			ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
			TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
			TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
			MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
			MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
		} `json:"track"`
	} `json:"media"`
}
type VideoInfo struct {
	Width          string `json:"width"`     // 1080 宽
	Height         string `json:"height"`    // 1920 高
	Format         string `json:"format"`    // HEVC 格式
	CodeId         string `json:"codeId"`    // hev1 标签
	FrameRate_Num  string `json:"frameRate"` // 30 帧率
	BitRate_String string `json:"bitRate"`   // 752 kb/s 比特率

}
type ffParam struct {
	crf     string
	bitrate string
	vcode   string
	acode   string
	vtag    string
}

func GetParam(fp string) VideoInfo {
	cmd := exec.Command("mediainfo", "--Output=JSON", "--Full", fp)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("mediainfo命令执行失败,有可能没安装mediainfo,请安装mediainfo后再试")
		return VideoInfo{}
	}
	var mi MediaInfo
	var vi VideoInfo
	if err = json.Unmarshal(output, &mi); err != nil {
		log.Fatalln("mediainfo解析json失败", err)
		return VideoInfo{}
	}
	for _, track := range mi.Media.Track {
		if track.Type == "Video" {
			bitrate := strings.Split(track.BitRateString, " ")[0]
			bitrate = strings.Join([]string{bitrate, "k"}, "")
			vi = VideoInfo{
				Width:          track.Width,
				Height:         track.Height,
				Format:         track.Format,
				CodeId:         track.CodecID,
				FrameRate_Num:  track.FrameRateNum,
				BitRate_String: track.BitRateString,
			}
		}
	}
	return vi
}

//	type OneVideo struct {
//		vLocation      string // 视频绝对路径
//		aLocation      string // 音频绝对路径
//		sLocation      string // 字幕绝对路径
//		FinalVideoName string // 最终视频文件名
//		FinalAudioName string // 最终音频文件名
//	}
type One struct {
	VName       string // 最终视频文件名
	AName       string // 最终音频文件名
	JName       string // json 中获取的名称
	VLocation   string // video.m4s 文件位置
	ALocation   string // audio.m4s 文件位置
	XmlLocation string // xml 文件位置
	AssLocation string // ass 文件位置
}

/*
根据视频属性生成命令
*/
func ChooseParamByVideo(vi VideoInfo, o One) *exec.Cmd {
	var cmd *exec.Cmd
	// todo 如果存在平均比特率 优先使用比特率
	// todo 如果比特率为空 根据帧率选择crf值
	if vi.Width > vi.Height {
		fmt.Println("横屏视频")
	} else if vi.Width < vi.Height {
		fmt.Println("竖屏视频")
	} else {
		fmt.Println("正方形视频")
	}

	if isHEVC(vi) {
		fmt.Println("hevc视频")
		// 存在比特率的情况下
		if vi.BitRate_String != "" {
			cmd = exec.Command("ffmpeg", "-i", o.VLocation,
				"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "copy", "-b:v", vi.BitRate_String, "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
		} else {
			// 不存在比特率 需要按分辨率判断crf
			switch {
			case vi.Width == "1080" || vi.Height == "1080":
				fmt.Println("1080p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "copy", "-crf", "22", vi.BitRate_String, "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			case vi.Width == "720" || vi.Height == "720":
				fmt.Println("720p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "copy", "-crf", "23", "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			case vi.Width == "480" || vi.Height == "480":
				fmt.Println("480p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "copy", "-crf", "24", "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			}
		}

	} else {
		//不是x265视频
		if vi.BitRate_String != "" {
			cmd = exec.Command("ffmpeg", "-i", o.VLocation,
				"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "copy", "-b:v", vi.BitRate_String, "-c:a", "copy", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
		} else {
			// 不存在比特率 需要按分辨率判断crf
			switch {
			case vi.Width == "1080" || vi.Height == "1080":
				fmt.Println("1080p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "libx265", "-crf", "22", vi.BitRate_String, "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			case vi.Width == "720" || vi.Height == "720":
				fmt.Println("720p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "libx265", "-crf", "23", "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			case vi.Width == "480" || vi.Height == "480":
				fmt.Println("480p")
				cmd = exec.Command("ffmpeg", "-i", o.VLocation,
					"-i", o.ALocation, "-i", o.AssLocation, "-c:v", "libx265", "-crf", "24", "-c:a", "libvorbis", "-ac", "1", "-tag:v", "hvc1", "-c:s", "ass", o.VName)
			}
		}
	}
	return cmd
}

/*
判断是否为hevc视频
*/
func isHEVC(vi VideoInfo) bool {
	if vi.Format == "HEVC" {
		return true
	} else {
		return false
	}
}

/*
判断HEVC视频是否含有hvc1标签
*/
func isTagHevc(vi VideoInfo) bool {
	if vi.Format == "HEVC" {
		if vi.CodeId == "hev1" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
