package xml2ass

import (
	"github.com/zhangyiming748/AVmerger/constant"
	"github.com/zhangyiming748/GetFileInfo"
	"testing"
)

// go test -v -run TestConv
func TestConv(t *testing.T) {
	xmls := GetFileInfo.GetAllFileInfo(constant.ANDROIDDANMAKU, "xml")
	for _, xml := range xmls {
		Conv(xml)
	}
}
