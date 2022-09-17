package composite

import (
	"AVmerger/log"
	"strings"
)

func Single(src, dst string) {
	rmds(src)
	var infos []Info
	head := getDir(src)
	log.Info.Printf("给定的目标文件下全部文件夹:%s\n", head)
	for _, first := range head {
		second := strings.Join([]string{src, first}, "/")
		if strings.Contains(second, "DS_store") {
			continue
		}
		log.Info.Printf("拼接第一级文件名:%s\n", second)
		thrid := strings.Join([]string{second, getDir(second)[0]}, "/")
		entry := strings.Join([]string{thrid, "entry.json"}, "/")
		log.Info.Printf("拼接entry文件名:%s\n", entry)
		e := readEntry(entry)
		random := getDir(thrid)[0]
		audio := strings.Join([]string{thrid, random, "audio.m4s"}, "/")
		log.Info.Printf("拼接audio文件名:%s\n", audio)
		video := strings.Join([]string{thrid, random, "video.m4s"}, "/")
		log.Info.Printf("拼接video文件名:%s\n", video)
		info := &Info{
			video: video,
			audio: audio,
			title: e.Title,
		}
		infos = append(infos, *info)

	}
	for _, value := range infos {
		log.Info.Printf("%+v\n", value)
	}
	merge(infos, dst)
}
