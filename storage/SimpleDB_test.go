package storage

import (
	"testing"
)

func TestFindHistoryFile(t *testing.T) {
	h := findHistoryFile()
	t.Log(h)
}
