package classify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DefaultKeywords 默认的歌手关键词列表
var DefaultKeywords = []string{
	"Adele",
	"Aqua",
	"Avril Lavigne",
	"Beyonce",
	"Blue",
	"Ed Sheeran",
	"Ellie Goulding",
	"F.I.R",
	"F.I.R/F.I.R",
	"FIR",
	"GALA",
	"GARNiDELiA",
	"GEM",
	"Groove Coverage",
	"High School Musical",
	"Katy Perry",
	"LadyGaga",
	"Mariah Carey",
	"Marionette",
	"Maroon 5",
	"Michael Learns To Rock",
	"NieR",
	"PB阿星儿",
	"Ricky Martin",
	"S.H.E",
	"Stellar",
	"Sweety",
	"Tank",
	"Taylor Swift",
	"Twins",
	"Victoria's Secret Fashion Show",
	"Yunomi",
	"kpop",
	"lady gaga",
	"中岛美嘉",
	"庾澄庆",
	"张栋梁",
	"张韶涵",
	"中岛美雪",
	"乐器",
	"伍佰",
	"凤凰传奇",
	"刘德华",
	"刘欢",
	"刘若英",
	"卡拉OK",
	"卢冠延",
	"古典",
	"叶倩文",
	"吴克群",
	"吴莫愁",
	"吴青峰",
	"周传雄",
	"周华健",
	"周杰伦",
	"唐磊",
	"坂井泉水",
	"夏川里美",
	"大肉嘎",
	"姚贝娜",
	"孙俪",
	"孙悦",
	"孙楠",
	"孙燕姿",
	"宋祖英",
	"容祖儿",
	"尼尔",
	"屠洪刚",
	"布兰妮",
	"庞龙",
	"康熙王朝",
	"庾澄庆",
	"张也",
	"张信哲",
	"张国荣",
	"张学友",
	"张宇",
	"张柏芝",
	"张栋梁",
	"张雨生",
	"张靓颖",
	"张韶涵",
	"徐怀钰",
	"徐歌阳",
	"徐良",
	"慕容晓晓",
	"成龙",
	"戴佩妮",
	"数码宝贝",
	"斯琴高丽",
	"曹格",
	"朗朗",
	"未整理",
	"李丽芬",
	"李春波",
	"李玖哲",
	"李玟",
	"李琛",
	"李翊君",
	"李贞贤",
	"杜德伟",
	"杨坤",
	"杨钰莹",
	"林俊杰",
	"林子祥",
	"林忆莲",
	"林志炫",
	"林肯公园",
	"林逸欣",
	"柯南",
	"柳岩",
	"梁咏琪",
	"梁心颐",
	"梁静茹",
	"梅艳芳",
	"歌剧",
	"歌舞青春音乐剧",
	"毛宁",
	"毛晓彤",
	"江涛",
	"江语晨",
	"沙宝亮",
	"海明威",
	"海鸣威",
	"满文军",
	"潘玮柏",
	"灌篮高手",
	"王力宏",
	"王强",
	"王心凌",
	"王菲",
	"理查德克莱德曼",
	"瑞奇马丁",
	"瓢三爷的小喇叭",
	"祖海",
	"秦霄贤",
	"罗大佑",
	"罗志祥",
	"群星",
	"羽泉",
	"至上励合",
	"花泽香菜",
	"苏永康",
	"苑冉",
	"莫文蔚",
	"萧亚轩",
	"蔡依林",
	"蔡健雅",
	"蔡妍",
	"蕾哈娜",
	"誓言",
	"许冠杰",
	"许慧欣",
	"费玉清",
	"费翔",
	"贾玲",
	"赵本山",
	"邓紫棋",
	"那英",
	"郑伊健",
	"郑少秋",
	"郭静",
	"金莎",
	"钟镇涛",
	"钢琴",
	"间谍过家家",
	"阿木",
	"陈佳",
	"陈坤",
	"陈奕迅",
	"陈慧娴",
	"陈慧琳",
	"陈明",
	"陈淑桦",
	"陈琳",
	"陈瑞",
	"陈百强",
	"陶喆",
	"青山",
	"青山黛玛",
	"韩磊",
	"音乐视听室",
	"香香",
	"马天宇",
	"高林生",
	"高枫",
	"黄品源",
	"黄圣依",
	"黄大炜",
	"黄明志",
	"黄渤",
	"黄霑",
	"齐秦",
}

// 这个包的功能是用来整理已经merged的文件,按照给定的关键词创建目标文件夹并移动到目标文件夹
// 文件会根据类型放置在Audio或Video子目录中，然后按歌手名进一步分类

/*
srcDir表示merged文件夹路径
dstDir表示目标根文件夹路径
keywords表示要匹配的歌手名关键词

处理逻辑：
1. 遍历srcDir目录下的所有文件
2. 根据文件扩展名确定类型：
  - .mp4文件放入dstDir/Video目录
  - .mp3文件放入dstDir/Audio目录

3. 根据文件名中的关键词创建子目录并移动文件：
  - 例如文件"1080P修复张韶涵 呐喊 MV修复版 TGCH张韶涵 呐喊 1080P修复版.mp4"
  - 匹配到关键词"张韶涵"
  - 最终路径为 dstDir/Video/张韶涵/1080P修复张韶涵 呐喊 MV修复版 TGCH张韶涵 呐喊 1080P修复版.mp4

示例:
merged文件夹下有如下文件:
1080P修复张韶涵  呐喊 MV修复版 TGCH张韶涵  呐喊 1080P修复版.mp4
1080P修复萧亚轩  我的男朋友MV 修复版 我的男朋友.mp3
keywords列表里有"张韶涵"
keywords列表里有"萧亚轩"

处理结果：
- 第一个文件会移动到 dstDir/Video/张韶涵/
- 第二个文件会移动到 dstDir/Audio/萧亚轩/
*/
func Classify(srcDir, dstDir string, keywords []string) {
	// 遍历源目录中的所有文件
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，跳过目录
		if info.IsDir() {
			return nil
		}

		// 获取文件名
		filename := info.Name()

		// 获取文件扩展名并确定类型目录
		ext := strings.ToLower(filepath.Ext(filename))
		var typeDir string
		switch ext {
		case ".mp4":
			typeDir = "Video"
		case ".mp3":
			typeDir = "Audio"
		default:
			// 不处理非mp3/mp4文件
			return nil
		}

		// 遍历关键词列表
		for _, keyword := range keywords {
			// 检查文件名是否包含关键词
			if strings.Contains(filename, keyword) {
				// 创建目标目录路径：dstDir/typeDir/keyword
				targetDir := filepath.Join(dstDir, typeDir, keyword)

				// 确保目标目录存在
				err := os.MkdirAll(targetDir, os.ModePerm)
				if err != nil {
					fmt.Printf("创建目录失败 %s: %v\n", targetDir, err)
					continue
				}

				// 构造目标文件路径
				targetPath := filepath.Join(targetDir, filename)

				// 移动文件
				err = os.Rename(path, targetPath)
				if err != nil {
					fmt.Printf("移动文件失败 %s 到 %s: %v\n", path, targetPath, err)
				} else {
					fmt.Printf("已移动文件 %s 到 %s\n", filename, targetDir)
				}
				break // 一个文件只移动一次
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}
