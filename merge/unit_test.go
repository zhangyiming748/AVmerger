package merge

import "testing"

func TestGetFolder(t *testing.T) {
	ret := getFolder("D:/git/AV\\download\\923131529\\c_1391127616\\")
	t.Log(ret)
}
