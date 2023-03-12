package AVmerger

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAllIn(t *testing.T) {
	AllIn("/Volumes/swap/download")
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

func TestDup(t *testing.T) {
	s1 := "a__s___1_2"
	// 期望 "a_s_1_2"
	ret := duplicate(s1, '_')
	t.Log(ret)
}
