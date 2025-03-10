package client

import (
	"testing"
)

// go test -v -timeout 100h -run TestMacOS
func TestMacOS(t *testing.T) {
	Convert("/Users/zen/Movies")
}

func TestFindJson(t *testing.T) {

	//t.Logf("%d,%v",len(files), files)
}
