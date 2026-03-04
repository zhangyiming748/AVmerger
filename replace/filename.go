package replace

import (
	"log"
	"regexp"
	"strings"
)

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
