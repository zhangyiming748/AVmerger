package classify

import (
	"testing"
)

func TestClassify(t *testing.T) {
	keywords := []string{"张韶涵","李玟","萧亚轩","Tank","FIR","孙燕姿","SHE","周杰伦","曹格","梁静茹","林俊杰","马天宇","刘若英","张栋梁","潘玮柏","Sweety","祖海","庾澄庆","凤凰传奇","Katy Perry","许慧欣","蔡依林","郭静","李玖哲","林俊杰","群星","王力宏","王心凌","吴克群","张信哲","Taylor Swift"}
	Classify("/Users/zen/Movies/bilibili/merged", "/Users/zen/Movies/bilibili/merged/zyl2012_音乐无限", keywords)
}
