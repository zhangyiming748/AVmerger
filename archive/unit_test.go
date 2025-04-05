package archive

import (
	"github.com/zhangyiming748/archiveVideos"
	"testing"
)
// go test -v -timeout 30h -run TestArchiveByGoModule
func TestArchiveByGoModule(t *testing.T) {
	root := "/Users/zen/Movies"
	archiveVideos.ArchiveVideos(root)
}
