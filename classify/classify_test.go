package classify

import (
	"testing"
)

func TestClassify(t *testing.T) {
	Classify("/Users/zen/Movies/bilibili/merged", "/Users/zen/Movies/bilibili/merged/zyl2012_音乐无限", DefaultKeywords)
}
