# AVmerger

合并哔哩哔哩手机版缓存的视频

# issue

Linux系统同一硬盘挂载到用户主目录的exfat分区不支持直接运行程序

尝试添加指定输出目录

# single文件结构

`./main Single /Users/zen/Movies/bilibili/single /Users/zen/Movies/bilibili/done`

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

# multi目录结构

`./main Multi /Users/zen/Movies/bilibili/multi/216750947 /Users/zen/Movies/bilibili/done`

```shell
.
├── c_796513071
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796516399
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796517384
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796517809
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796518132
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796518531
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796518775
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796519173
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796525400
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
├── c_796525959
│             ├── 80
│             │             ├── audio.m4s
│             │             ├── index.json
│             │             └── video.m4s
│             ├── danmaku.xml
│             └── entry.json
└── multi.md

20 directories, 51 files
```