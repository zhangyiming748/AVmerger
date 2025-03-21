package client

import (
	"bufio"

	"os"
	"strings"
	"testing"
)

// go test -v -timeout 100h -run TestMacOS
func TestMacOS(t *testing.T) {
	Convert("/Users/zen/Movies")
}
func TestRemove(t *testing.T){
	RemoveEncryptionHeader("/Users/zen/github/AVmerger/101391217/101391217-1-30112.m4s")
	RemoveEncryptionHeader("/Users/zen/github/AVmerger/101391217/101391217-1-30280.m4s")
	RemoveEncryptionHeader("/Users/zen/github/AVmerger/101391217/101391217-1-100027.m4s")
}
func TestFindJson(t *testing.T) {

	//t.Logf("%d,%v",len(files), files)
}

func TestExtractDM(t *testing.T) {
	content, err := os.ReadFile("dm1")
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	outFile, err := os.Create("dm1.txt")
	if err != nil {
		t.Fatalf("创建输出文件失败: %v", err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	// 遍历内容寻找以汉字开头的文本
	for i := 0; i < len(content); i++ {
		// 检查是否是汉字的起始字节
		if i+2 < len(content) && content[i] >= 0xE4 && content[i] <= 0xE9 {
			start := i
			// 向后查找直到遇到控制字符
			for j := start; j < len(content); j++ {
				if content[j] < 32 {
					// 找到一组完整的文本
					text := string(content[start:j])
					if text != "" {
						// 如果文本中包含@，只保留@之前的部分
						if idx := strings.Index(text, "@"); idx != -1 {
							text = text[:idx]
						}
						// 跳过包含 ??6? 的行
						if !strings.Contains(text, "??6?") && text != "" {
							writer.WriteString(text + "\n")
						}
					}
					i = j
					break
				}
			}
		}
	}

	writer.Flush()
	t.Log("弹幕提取完成")
}
