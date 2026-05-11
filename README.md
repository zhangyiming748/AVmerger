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
go build -o AVmerger.exe  # Windows
go build -o AVmerger      # macOS/Linux
```

#### 系统要求

- **ffmpeg**: 用于音视频合并和转码
- **mediainfo**: 用于获取媒体文件信息

确保这两个工具已安装并在系统 PATH 中。

#### 可用命令

```
Available Commands:
  a2p         转换安卓客户端下载目录
  archive     归档合并后的视频文件
  client      处理 B 站客户端缓存视频
  cover       归档封面图片
  rename      批量重命名文件
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
```

---

### 1. client - 处理 B 站客户端缓存视频

适用于 macOS/Windows/Linux 的 B 站客户端缓存目录，支持自动检测默认缓存路径。

**参数：**
- `-i, --src string`: B 站客户端缓存目录基础路径（可选，为空则使用默认路径）
- `-o, --dst string`: 输出目录基础路径（必填）
- `-a, --archive string`: 归档目录基础路径（可选，用于分类整理合并后的文件）

**使用示例：**

```bash
# 基本用法（自动检测默认缓存路径）
./AVmerger client -o ~/Videos/output

# 指定源目录和目标目录
./AVmerger client -i ~/Movies/bilibili -o ~/Videos/output

# 带归档目录（自动分类整理）
./AVmerger client -i ~/Movies/bilibili -o ~/Videos/output -a ~/Videos/archive
```

**默认路径：**
- macOS: `~/Movies/bilibili`
- Windows/Linux: `~/Videos/bilibili`

---

### 2. a2p - 转换安卓客户端下载目录

将安卓客户端下载目录中的音视频文件合并并转换到 PC 端格式。

**参数：**
- `-i, --src string`: 安卓客户端下载目录路径（必填）
- `-o, --dst string`: 输出目录基础路径（必填）
- `-a, --archive string`: 归档目录基础路径（可选，用于分类整理合并后的文件）

**使用示例：**

```bash
# 基本用法
./AVmerger a2p -i /sdcard/Android/data/tv.danmaku.bili/download -o ~/Videos/output

# 带归档目录（自动分类整理）
./AVmerger a2p -i /sdcard/Android/data/tv.danmaku.bili/download -o ~/Videos/output -a ~/Videos/archive
```

---

### 3. cover - 归档封面图片

将源目录下的所有 cover.jpg 文件移动到目标目录并按顺序重命名。

**参数：**
- `-i, --src string`: 源目录路径（必填）
- `-o, --dst string`: 目标目录路径（必填）

**使用示例：**

```bash
# 归档封面图片
./AVmerger cover -i ~/Videos/bilibili -o ~/Pictures/covers
```

---

### 4. archive - 归档合并后的视频文件

将源目录中合并后的视频文件按照分类规则归档到目标目录。

**参数：**
- `-i, --src string`: 源目录路径（必填）
- `-o, --dst string`: 目标目录路径（必填）

**使用示例：**

```bash
# 归档视频文件（按类型分类）
./AVmerger archive -i ~/Videos/output -o ~/Videos/archive
```

---

### 5. rename - 批量重命名文件

批量替换指定目录下所有文件名中的特定字符串。

**参数：**
- `-d, --dir string`: 要处理的根目录路径（必填）
- `-i, --src string`: 需要被替换的字符串（必填）
- `-o, --dst string`: 替换后的字符串（必填）

**使用示例：**

```bash
# 将所有文件名中的 "旧名称" 替换为 "新名称"
./AVmerger rename -d ~/Videos -i "旧名称" -o "新名称"

# 移除文件名中的特定字符
./AVmerger rename -d ~/Videos -i "[广告]" -o ""

# 批量修正拼写错误
./AVmerger rename -d ~/Videos -i "recieve" -o "receive"
```

---

### 通用说明

**查看帮助信息：**

```bash
# 查看全局帮助
./AVmerger --help

# 查看子命令帮助
./AVmerger client --help
./AVmerger a2p --help
./AVmerger cover --help
./AVmerger archive --help
./AVmerger rename --help
```

**注意事项：**

1. 源目录和目标目录不能相同（防止误删数据）
2. `a2p`、`cover`、`archive`、`rename` 命令必须指定所有必填参数
3. `client` 命令的 `--src` 参数可选，为空时使用系统默认路径
4. 需要系统已安装 `ffmpeg` 和 `mediainfo` 命令行工具
5. 所有命令都会记录详细日志到 `avmerge.log` 文件

### download 文件夹结构
