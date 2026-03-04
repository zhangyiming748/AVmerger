package cover

import (
	"testing"
)

// go test -v -run TestFindAllCover
func TestFindAllCover(t *testing.T) {
	src := "/Volumes/Untitled/代码测试用例/download"
	covers, err := FindAllCover(src)
	if err != nil {
		t.Error(err)
	}
	for i, cover := range covers {
		t.Log(i, cover)
	}
}

// go test -v -run TestArchiveCover
func TestArchiveCover(t *testing.T) {
	src := "/Volumes/Untitled/代码测试用例/download"
	ArchiveCovers(src, "/Volumes/Untitled/cover")
}
