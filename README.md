# AVmerger

合并哔哩哔哩手机版缓存的视频

## 2025/11/12 更新

放弃直接使用，变成作为 go modules 提供

### 作为 Go Module 使用

使用方法参考

```go
package main

import (
    "log"

    "github.com/zhangyiming748/AVmerger"
    "github.com/zhangyiming748/archive"
    "github.com/zhangyiming748/finder"
)

func init() {
    log.SetFlags(log.Ltime | log.Lshortfile)
}
func main() {
    dst := "C:\\Users\\zen\\Gitea\\UseAVmerge"
    AVmerge.Client(dst)
    AVmerge.Android2PC("",dst)
    folders := finder.FindAllFolders(dst)
    for i, folder := range folders {
        log.Printf("正在处理第%d/%d个文件夹:%s\n", i+1, len(folders), folder)
        vFiles := finder.FindAllVideosInRoot(folder)
        for j, vFile := range vFiles {
            log.Printf("正在处理第%d/%d个文件夹:%s中的第个%d/%d文件:%s\n", i+1, len(folders), folder, j+1, len(vFiles), vFile)
            archive.Convert2H265(vFile)
        }
    }
}
```

### 作为命令行工具使用

AVmerger 提供了基于 Cobra 的命令行界面，可以方便地通过命令行调用核心功能。

#### 安装依赖

```bash
cd AVmerger
go get github.com/spf13/cobra
go build -o avmerger
```

#### 命令行参数

**全局参数：**

- `-i, --src`: 源目录路径（B 站客户端缓存目录或安卓下载目录）
- `-o, --dst`: 目标输出目录路径（必填）
- `-a, --archive`: 归档目录路径（可选，用于分类整理合并后的文件）

**子命令：**

1. **client** - 处理 B 站客户端缓存视频
   - 适用于 macOS/Linux 的 B 站客户端缓存目录
   - 支持自动检测默认缓存路径

2. **android2pc** - 转换安卓客户端下载目录
   - 适用于从安卓设备复制的 download 目录
   - 需要明确指定源目录和目标目录

#### 使用示例

**处理 B 站客户端缓存：**

```bash
# 基本用法（自动检测默认缓存路径）
./avmerger client -o ~/Videos/output

# 指定源目录和目标目录
./avmerger client -i ~/Movies/bilibili -o ~/Videos/output

# 带归档目录（自动分类整理）
./avmerger client -i ~/Movies/bilibili -o ~/Videos/output -a ~/Videos/archive
```

**处理安卓客户端下载：**

```bash
# 基本用法
./avmerger android2pc -i /sdcard/download -o ~/Videos/output

# 带归档目录（自动分类整理）
./avmerger android2pc -i /sdcard/download -o ~/Videos/output -a ~/Videos/archive
```

**查看帮助信息：**

```bash
# 查看全局帮助
./avmerger --help

# 查看子命令帮助
./avmerger client --help
./avmerger android2pc --help
```

#### 功能说明

- **client 命令**: 调用 `core.Client(src, dst)` 函数，处理 B 站客户端缓存目录，自动检测操作系统并设置默认路径
- **android2pc 命令**: 调用 `core.Android2PC(src, dst)` 函数，专门处理安卓设备的 download 目录结构
- **--archive 参数**: 在视频合并完成后，自动调用 `core.ClassifyAfterMerge()` 函数将文件按类型分类到归档目录

#### 注意事项

1. 源目录和目标目录不能相同（防止误删数据）
2. `android2pc` 命令必须指定 `--src` 和 `--dst` 参数
3. `client` 命令的 `--src` 参数可选，为空时使用系统默认路径
4. 需要系统已安装 `ffmpeg` 和 `mediainfo` 命令行工具

### download 文件夹结构
