package merge

import "testing"

// go test -v -run TestBasicInfo
func TestBasicInfo(t *testing.T) {
	root := "/Users/zen/gitea/AVmerge/download"
	GetBasicInfo(root)
}
