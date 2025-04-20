package util
import (
	"testing"
)
// go test -v -run TestRSync
func TestRSync(t *testing.T) {
	UploadWithRsync("./root.go")
	UploadWithRsync("../archive")
}