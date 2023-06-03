package AVmerger

import (
	"fmt"
	"golang.org/x/exp/slog"
)

func AllInAAC(root string) {
	infos := get(root)
	slog.Debug("解析后", slog.Any("返回的视频", infos))
	for i, info := range *infos {
		slog.Info(fmt.Sprintf("正在合并第 %d/%d 个视频\n", i+1, len(*infos)))
		aac(root, info)
	}
}
