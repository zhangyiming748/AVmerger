package util

import (
	"testing"
)

// go test -v -run TestMaster

// go test -v -run TestGetEntryFilesWithExt
func TestGetEntryFilesWithExt(t *testing.T) {
	root := "/Users/zen/gitea/AVmerge/download"
	t.Log("he")
	GetEntryFilesWithExt(root, "json")
}
