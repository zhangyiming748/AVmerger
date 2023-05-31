package AVmerger

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAllIn(t *testing.T) {
	AllIn("C:\\Users\\zen\\Github\\AVmerger\\testfile")
	AllInH265("")
}

func TestGet(t *testing.T) {
	ret := get("/Volumes/swap/download")
	file, err := os.OpenFile("list.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	marshal, err := json.Marshal(ret)
	if err != nil {
		return
	}
	file.WriteString(string(marshal))

}
