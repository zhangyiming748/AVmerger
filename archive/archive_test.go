package archive

import (
	"testing"

	"github.com/zhangyiming748/AVmerger/util"
)
func init() {
	util.SetLog()
}
//go test -timeout 1000h -v -run TestAll
func TestAll(t *testing.T) {
	files,_:=GetAllFiles("/Users/zen/Movies")
	for _,file:=range files {
		Convert(file)
	}
}