package classify

import (
	"testing"

	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/finder"
)

func TestClassify(t *testing.T) {
	Classify("/Users/zen/Movies/bilibili/merged", "/Users/zen/Movies/bilibili/merged/zyl2012_音乐无限", DefaultKeywords)
}

func TestGen(t *testing.T) {
	var lines []string
	foldersName := finder.FindAllFolders("/Volumes/A2/music/video")
	for _, folder := range foldersName {
		lines = append(lines, "\""+folder+"\",")
	}
	util.WriteByLine("string.txt", lines)
}
