package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/AVmerger/util"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	constant.SetLogLevel("Debug")
	files, _ := util.GetMKVFilesWithExt("/mnt/f/alist/bilibili")
	for _, file := range files {
		slog.Debug(fmt.Sprintf("获取到的mkv%+v\n", file))
		codec, width, height := GetCodec(file.FullPath)
		if codec == "HEVC" {
			slog.Info("跳过已经是h265编码的视频", slog.String("视频名", file.FullPath))
			continue
		}
		if codec == "VP9" {
			slog.Info("跳过已经是VP9编码的视频", slog.String("视频名", file.FullPath))
			continue
		}

		crf := GetCrfForVP9(width, height)

		after := strings.Replace(file.FullPath, ".mkv", "vp9.mkv", 1)
		cmd := exec.Command("ffmpeg", "-i", file.FullPath, "-map", "0:v:0", "-map", "0:a:0", "-map", "0:s:0", "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", after)
		slog.Debug(fmt.Sprintf("命令原文:%s", cmd.String()))
		output, err := cmd.CombinedOutput()
		if err != nil {
			return
		} else {
			if compere(file.FullPath, after) == "smaller" {
				err := os.Remove(file.FullPath)
				if err != nil {
					return
				} else {
					slog.Warn(fmt.Sprintf("删除文件%v\n", file.FullPath))
				}
			} else {
				slog.Warn("转换后的文件比源文件更大", slog.String("源文件", file.FullPath), slog.String("目标文件", after), slog.String("命令原文", cmd.String()))
			}
		}
		slog.Debug(fmt.Sprintln("命令输出", string(output)))
	}
}
func compere(before, after string) (result string) {
	file1 := before
	file2 := after

	info1, err1 := os.Stat(file1)
	info2, err2 := os.Stat(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("获取文件信息失败：", err1, err2)
		return
	}

	size1 := info1.Size()
	size2 := info2.Size()

	if size1 > size2 {
		fmt.Printf("%s 文件大小为 %d 字节，大于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "smaller"
	} else if size1 < size2 {
		fmt.Printf("%s 文件大小为 %d 字节，小于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "bigger"
	} else {
		fmt.Printf("%s 文件大小为 %d 字节，等于 %s 文件大小 %d 字节", file1, size1, file2, size2)
		result = "equal"
	}
	return result
}

type MI struct {
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
			UniqueID                 string `json:"UniqueID"`
			UniqueIDString           string `json:"UniqueID_String,omitempty"`
			VideoCount               string `json:"VideoCount,omitempty"`
			AudioCount               string `json:"AudioCount,omitempty"`
			TextCount                string `json:"TextCount,omitempty"`
			VideoFormatList          string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList  string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList           string `json:"Video_Codec_List,omitempty"`
			AudioFormatList          string `json:"Audio_Format_List,omitempty"`
			AudioFormatWithHintList  string `json:"Audio_Format_WithHint_List,omitempty"`
			AudioCodecList           string `json:"Audio_Codec_List,omitempty"`
			TextFormatList           string `json:"Text_Format_List,omitempty"`
			TextFormatWithHintList   string `json:"Text_Format_WithHint_List,omitempty"`
			TextCodecList            string `json:"Text_Codec_List,omitempty"`
			CompleteName             string `json:"CompleteName,omitempty"`
			FileNameExtension        string `json:"FileNameExtension,omitempty"`
			FileName                 string `json:"FileName,omitempty"`
			FileExtension            string `json:"FileExtension,omitempty"`
			Format                   string `json:"Format"`
			FormatString             string `json:"Format_String"`
			FormatUrl                string `json:"Format_Url,omitempty"`
			FormatExtensions         string `json:"Format_Extensions,omitempty"`
			FormatCommercial         string `json:"Format_Commercial"`
			FormatVersion            string `json:"Format_Version,omitempty"`
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
			DurationString4          string `json:"Duration_String4,omitempty"`
			DurationString5          string `json:"Duration_String5"`
			OverallBitRate           string `json:"OverallBitRate,omitempty"`
			OverallBitRateString     string `json:"OverallBitRate_String,omitempty"`
			FrameRate                string `json:"FrameRate,omitempty"`
			FrameRateString          string `json:"FrameRate_String,omitempty"`
			FrameCount               string `json:"FrameCount,omitempty"`
			IsStreamable             string `json:"IsStreamable,omitempty"`
			Description              string `json:"Description,omitempty"`
			FileModifiedDate         string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal    string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication       string `json:"Encoded_Application,omitempty"`
			EncodedApplicationString string `json:"Encoded_Application_String,omitempty"`
			EncodedLibrary           string `json:"Encoded_Library,omitempty"`
			EncodedLibraryString     string `json:"Encoded_Library_String,omitempty"`
			Extra                    struct {
				ErrorDetectionType string `json:"ErrorDetectionType,omitempty"`
				VENDORID           string `json:"VENDOR_ID,omitempty"`
			} `json:"extra,omitempty"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			IDString                       string `json:"ID_String,omitempty"`
			FormatInfo                     string `json:"Format_Info,omitempty"`
			FormatProfile                  string `json:"Format_Profile,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatTier                     string `json:"Format_Tier,omitempty"`
			InternetMediaType              string `json:"InternetMediaType,omitempty"`
			CodecID                        string `json:"CodecID,omitempty"`
			Width                          string `json:"Width,omitempty"`
			WidthString                    string `json:"Width_String,omitempty"`
			Height                         string `json:"Height,omitempty"`
			HeightString                   string `json:"Height_String,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString       string `json:"DisplayAspectRatio_String,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString            string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
			FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString        string `json:"ChromaSubsampling_String,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			BitDepthString                 string `json:"BitDepth_String,omitempty"`
			Delay                          string `json:"Delay,omitempty"`
			DelayString3                   string `json:"Delay_String3,omitempty"`
			DelayString4                   string `json:"Delay_String4,omitempty"`
			DelayString5                   string `json:"Delay_String5,omitempty"`
			DelaySource                    string `json:"Delay_Source,omitempty"`
			DelaySourceString              string `json:"Delay_Source_String,omitempty"`
			Default                        string `json:"Default,omitempty"`
			DefaultString                  string `json:"Default_String,omitempty"`
			Forced                         string `json:"Forced,omitempty"`
			ForcedString                   string `json:"Forced_String,omitempty"`
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
			FormatAdditionalFeatures       string `json:"Format_AdditionalFeatures,omitempty"`
			Channels                       string `json:"Channels,omitempty"`
			ChannelsString                 string `json:"Channels_String,omitempty"`
			ChannelPositions               string `json:"ChannelPositions,omitempty"`
			ChannelPositionsString2        string `json:"ChannelPositions_String2,omitempty"`
			ChannelLayout                  string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame                string `json:"SamplesPerFrame,omitempty"`
			SamplingRate                   string `json:"SamplingRate,omitempty"`
			SamplingRateString             string `json:"SamplingRate_String,omitempty"`
			SamplingCount                  string `json:"SamplingCount,omitempty"`
			CompressionMode                string `json:"Compression_Mode,omitempty"`
			CompressionModeString          string `json:"Compression_Mode_String,omitempty"`
			VideoDelay                     string `json:"Video_Delay,omitempty"`
			VideoDelayString3              string `json:"Video_Delay_String3,omitempty"`
			VideoDelayString5              string `json:"Video_Delay_String5,omitempty"`
			CodecIDInfo                    string `json:"CodecID_Info,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

/*
mediainfo --Output=JSON --Full
*/
func GetCodec(fp string) (codec, width, height string) {
	cmd := exec.Command("mediainfo", "--Output=JSON", "--Full", fp)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	var mi MI
	err = json.Unmarshal(output, &mi)
	if err != nil {
		return
	}
	for _, v := range mi.Media.Track {
		if v.Type == "Video" {
			if v.Format == "HEVC" {
				codec = "HEVC"
			}
			if v.Format == "VP9" {
				codec = "VP9"
			}
			width = v.Width
			height = v.Height
		}
	}
	return codec, width, height
}
func GetCrfForVP9(width, height string) (crf string) {
	width_int, _ := strconv.Atoi(width)
	height_int, _ := strconv.Atoi(height)
	stand := 0
	if width_int > height_int {
		stand = height_int
	} else if height_int > width_int {
		stand = width_int
	} else {
		stand = width_int
	}
	if stand >= 2160 {
		crf = "15"
	} else if stand >= 1440 {
		crf = "24"
	} else if stand >= 1080 {
		crf = "31"
	} else if stand >= 720 {
		crf = "32"
	} else if stand >= 480 {
		crf = "33"
	} else if stand >= 135 {
		crf = "34"
	} else {
		crf = "35"
	}
	return crf
}
