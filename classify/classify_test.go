package classify

import (
	"testing"
)

func TestClassify(t *testing.T) {
	keywords := []string{"张韶涵","李玟","萧亚轩","Tank","FIR","孙燕姿","SHE","周杰伦","曹格","梁静茹","林俊杰","马天宇","刘若英"}
	Classify("/Users/zen/Movies/bilibili/merged/zyl2012_音乐无限", "/Users/zen/Movies/bilibili/merged/zyl2012_音乐无限", keywords)
}
