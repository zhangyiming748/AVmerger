# AVmerger

合并哔哩哔哩手机版缓存的视频

# 在安卓termux上编译

```shell
go build -o ../../../usr/bin/merge main.go 
```
其中`.shortcuts/bili.sh`文件内容

```shell
#!/data/data/com.termux/files/usr/bin/bash
sudo merge 
```


# Todo
- [ ] 找到entry的同级目录
- [ ] 保持之前的逻辑
- [ ] 同时支持main函数和包调用

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


# 优化


一把梭


自动判断方式