package cover

import (
	"testing"
)

// go test -v -run TestFindAllCover
func TestFindAllCover(t *testing.T) {
	src := "C:\\Users\\zen\\Videos\\download"
	covers, err := findAllCover(src)
	if err != nil {
		t.Error(err)
	}
	for i, cover := range covers {
		t.Log(i, cover)
	}
}

// go test -v -run TestArchiveCover
func TestArchiveCover(t *testing.T) {
	src := "C:\\Users\\zen\\Videos\\download"
	dst := "C:\\Users\\zen\\Videos\\cover"
	ArchiveCovers(src, dst)
}
