package replace

import (
	"log"
	"regexp"
	"strings"
)

// 以下是旧版本的文件名处理函数，通过字符串替换来处理特殊字符
// 由于可能存在未覆盖的特殊字符，已改用正则表达式方式实现
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
//	str = strings.Replace(str, " ", "", -1)
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
//	str = strings.Replace(str, "\"", "", -1)
//	str = strings.Replace(str, "\"", "", -1)
//	str = strings.Replace(str, "--", "", -1)
//	str = strings.Replace(str, "_", "", -1)
//	str = strings.Replace(str, "：", "", -1)
//	return str
//}

// Package replace 提供了文件名处理相关的功能
// 主要用于清理和规范化文件名，确保文件名只包含合法字符
// ForFileName 处理文件名，只保留数字、字母和中文字符
// 通过遍历字符串的每个字符，使用正则表达式判断字符的有效性
// 参数:
//   - name: 原始文件名字符串
//
// 返回值:
//   - string: 处理后的文件名，只包含数字、字母、中文和空格
func ForFileName(name string) string {
	nStr := ""
	// 遍历原始文件名的每个字符
	for _, v := range name {
		// 检查字符是否有效
		if Effective(string(v)) {
			// 将有效字符添加到结果字符串中
			nStr = strings.Join([]string{nStr, string(v)}, "")
		}
	}
	log.Printf("正则表达式匹配数字字母汉字:%v\n", nStr)
	return nStr
}

// Effective 判断字符是否为有效字符（数字、字母、中文或空格）
// 使用正则表达式匹配不同类型的字符
// 参数:
//   - s: 待检查的单个字符
//
// 返回值:
//   - bool: 如果是有效字符返回true，否则返回false
func Effective(s string) bool {
	// 空格作为特殊情况处理，允许保留
	if s == " " {
		return true
	}
	// 编译正则表达式用于匹配不同类型的字符
	num := regexp.MustCompile(`\d`)          // 匹配任意一个数字
	letter := regexp.MustCompile(`[a-zA-Z]`) // 匹配任意一个字母
	char := regexp.MustCompile(`[\p{Han}]`)  // 匹配任意一个汉字
	// 如果字符匹配任意一种类型，则认为是有效字符
	if num.MatchString(s) || letter.MatchString(s) || char.MatchString(s) {
		return true
	}
	return false
}

// RemoveLeadingSpace 移除字符串开头的空格字符
// 用于清理字符串前导空格，保持字符串格式整洁
// 参数:
//   - s: 需要处理的原始字符串
//
// 返回值:
//   - string: 移除开头空格后的字符串
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
