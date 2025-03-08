package client

import (
    "os"
    "io"
    "log"
)

func RemoveEncryptionHeader(filePath string) error {
    // 打开加密的视频文件
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // 读取文件内容
    content, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    // 查找第一个非"30"的位置，并计数
    var cleanContent []byte
    count := 0
    for i := 0; i < len(content); {
        if content[i] == 0x30 {
            count++
            i++
            continue
        }
        cleanContent = content[i:]
        break
    }

    // 打印删除的数量
    log.Printf("成功删除 %d 个'30'字符\n", count)

    // 写回文件
    return os.WriteFile(filePath, cleanContent, 0644)
}