package util

import (
	"testing"
)

// go test -v -run TestRSync
func TestRSync(t *testing.T) {
	remoteDir := "/Volumes/ugreen/alist/bili/"
	UploadWithRsyncAll(remoteDir, "../constant", "../merge")
}
