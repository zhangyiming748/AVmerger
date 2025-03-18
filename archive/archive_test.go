package archive

import (
	"testing"
)

//go test -timeout 1h -v -run TestAll
func TestAll(t *testing.T) {
	files,_:=GetAllFiles("/Users/zen/Movies")
	for _,file:=range files {
		Convert(file)
	}
}