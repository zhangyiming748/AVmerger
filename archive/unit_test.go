package archive

import (
	"testing"

	"github.com/zhangyiming748/AVmerger/util"
	"github.com/zhangyiming748/archiveVideos"
)

// go test -v -timeout 30h -run TestArchiveByGoModule
func TestArchiveByGoModule(t *testing.T) {
	util.SetLog()
	root := "/Users/zen/Movies"
	archiveVideos.ArchiveVideos(root)
}
