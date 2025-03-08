package client

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// go test -v -run TestMacOS
func TestMacOS(t *testing.T) {
	files, err := FindVideoInfoFiles("/Users/zen/Movies")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Logf("json file is %s", file)
		dir := filepath.Dir(file)
		t.Logf("dir is %s", dir)
		media, err := FindM4sFiles(dir)
		if err != nil {
			t.Fatal(err)
		}
		if len(media) != 2 {
			t.Fatal("no media file")
		}
		t.Logf("media file is %s", media)
		RemoveEncryptionHeader(media[0])
		RemoveEncryptionHeader(media[1])
		vi, _ := ReadVideoInfo(file)
		home, _ := os.UserHomeDir()
		baseDir := filepath.Join(home, "Movies", vi.Uname)
		os.MkdirAll(baseDir, 0755)
		title := strings.Join([]string{vi.Title, "mp4"}, ".")
		target := filepath.Join(baseDir, title)
		cmd := exec.Command("ffmpeg", "-i", media[0], "-i", media[1], "-c:v", "libx265", "-tag:v", "hvc1", target)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)

		}
		log.Printf("out is %s", out)
	}
	//RemoveEncryptionHeader("/Users/zen/Movies/bilibili/28734260477/28734260477-1-100046.m4s")
}

func TestFindJson(t *testing.T) {

	//t.Logf("%d,%v",len(files), files)
}
