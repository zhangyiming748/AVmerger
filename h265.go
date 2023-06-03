package AVmerger

import (
	"fmt"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

func hevc(dst string, info Info) {
	n := Duplicate(info.Name, '_')
	n = Duplicate(n, '.')
	n = replace.ForFileName(n)
	name := strings.Join([]string{n, "mp4"}, ".")
	target := strings.Join([]string{dst, name}, string(os.PathSeparator))
	cmd := exec.Command("ffmpeg", "-i", info.Video, "-i", info.Audio, "-c:v", "libx265", "-tag:v", "hvc1", target)
	slog.Info("命令", slog.Any("ffmpeg", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("错误", slog.Any("cmd.StdoutPipe", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("错误", slog.Any("cmd.Run", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("错误", slog.Any("命令执行中", err))
		return
	}
	slog.Info("完成当前文件的处理", slog.Any("源文件", info.Name), slog.Any("目标文件夹", dst))
	if err := os.RemoveAll(info.Del); err != nil {
		slog.Warn("", slog.Any("删除源文件失败", err))
	} else {
		slog.Warn("", slog.Any("删除源目录", info.Del))
	}
}
