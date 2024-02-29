package main

import (
	"os"
	"testing"
)

// go test -v -run  TestRemoveSrc
func TestRemoveSrc(t *testing.T) {
	err := os.RemoveAll("/mnt/d/git/AVmerger/downloads")
	if err != nil {
		panic(err)
	}
}
