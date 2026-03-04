# AVmerger

合并哔哩哔哩手机版缓存的视频

## 2025/11/12 更新

放弃直接使用，变成作为go modules提供
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

### download文件夹结构

```shell
single/
├── 302039705
│             └── c_806115701
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 387104484
│             └── c_804016687
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 387590919
│             └── c_818879579
│                 ├── 80
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 602021125
│             └── c_805053161
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 644521662
│             └── c_805045481
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 686955054
│             └── c_802718147
│                 ├── 64
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 687225722
│             └── c_807042244
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 729608041
│             └── c_807005607
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 814501582
│             └── c_806123791
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
├── 899603455
│             └── c_806123799
│                 ├── 16
│                 │             ├── audio.m4s
│                 │             ├── index.json
│                 │             └── video.m4s
│                 ├── danmaku.xml
│                 └── entry.json
└── 941940043
    └── c_802198105
        ├── 64
        │             ├── audio.m4s
        │             ├── index.json
        │             └── video.m4s
        ├── danmaku.xml
        └── entry.json

33 directories, 55 files
```
