package AVmerger

import (
	"testing"
	"time"
)

var (
	cstZone = time.FixedZone("CST", 8*3600)
)

func TestSingelCMD(t *testing.T) {
	src := "/Users/zen/Github/AVmerger/file/single"
	dst := "/Users/zen/Github/AVmerger/file"
	Single(src, dst)
}

func TestMultiCMD(t *testing.T) {
	src := "/Users/zen/Github/AVmerger/file/multi"
	dst := "/Users/zen/Github/AVmerger/file"
	Multi(src, dst)
}
