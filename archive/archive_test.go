package archive

import (
	"testing"
)
func TestAll(t *testing.T) {
	files,_:=GetAllFiles("/Users/zen/Movies/唯一8090")
	for _,file:=range files {
		Convert(file)

	}
}