package replace

import (
	"log"
	"regexp"
	"strings"
)

//func ForFileName(str string) string {
//	str = strings.Replace(str, "。", ".", -1)
//	str = strings.Replace(str, "，", ",", -1)
//	str = strings.Replace(str, "《", "(", -1)
//	str = strings.Replace(str, "》", ")", -1)
//	str = strings.Replace(str, "【", "(", -1)
//	str = strings.Replace(str, "】", ")", -1)
//	str = strings.Replace(str, "（", "(", -1)
//	str = strings.Replace(str, "）", ")", -1)
//	str = strings.Replace(str, "「", "(", -1)
//	str = strings.Replace(str, "」", ")", -1)
//	str = strings.Replace(str, "+", "_", -1)
//	str = strings.Replace(str, "`", "", -1)
//	str = strings.Replace(str, " ", "", -1)
//	str = strings.Replace(str, "\u00A0", "", -1)
//	str = strings.Replace(str, "\u0000", "", -1)
//	str = strings.Replace(str, "·", "", -1)
//	str = strings.Replace(str, "\uE000", "", -1)
//	str = strings.Replace(str, "\u000D", "", -1)
//	str = strings.Replace(str, "、", "", -1)
//	//str = strings.Replace(str, "/", "", -1)
//	str = strings.Replace(str, "！", "", -1)
//	str = strings.Replace(str, "|", "", -1)
//	str = strings.Replace(str, "｜", "", -1)
//	str = strings.Replace(str, ":", "", -1)
//	str = strings.Replace(str, " ", "", -1)
//	str = strings.Replace(str, "&", "", -1)
//	str = strings.Replace(str, "？", "", -1)
//	str = strings.Replace(str, "(", "", -1)
//	str = strings.Replace(str, ")", "", -1)
//	str = strings.Replace(str, "-", "", -1)
//	str = strings.Replace(str, " ", "", -1)
//	str = strings.Replace(str, "“", "", -1)
//	str = strings.Replace(str, "”", "", -1)
//	str = strings.Replace(str, "--", "", -1)
//	str = strings.Replace(str, "_", "", -1)
//	str = strings.Replace(str, "：", "", -1)
//	return str
//}

/*
仅保留文件名中的 数字 字母 和 中文
*/
func ForFileName(name string) string {
	nStr := ""
	for _, v := range name {
		if Effective(string(v)) {
			// fmt.Printf("%d\t有效%v\n", i, string(v))
			nStr = strings.Join([]string{nStr, string(v)}, "")
		}
	}
	log.Printf("正则表达式匹配数字字母汉字:%v\n", nStr)
	return nStr
}
func Effective(s string) bool {
	if s == " " {
		return true
	}
	num := regexp.MustCompile(`\d`)          // 匹配任意一个数字
	letter := regexp.MustCompile(`[a-zA-Z]`) // 匹配任意一个字母
	char := regexp.MustCompile(`[\p{Han}]`)  // 匹配任意一个汉字
	if num.MatchString(s) || letter.MatchString(s) || char.MatchString(s) {
		return true
	}
	return false
}
func RemoveLeadingSpace(s string) string {
	// 检查字符串是否为空
	if len(s) == 0 {
		return s
	}

	// 检查第一个字符是否为空格
	if s[0] == ' ' {
		// 返回去掉第一个字符的字符串
		return s[1:]
	}

	// 如果第一个字符不是空格，返回原字符串
	return s
}
