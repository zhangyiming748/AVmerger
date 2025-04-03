package test

import (
	"github.com/zhangyiming748/AVmerger/client"
	"runtime"
	"testing"
)

// go test -v -timeout 100h -run TestMacOS
/*
测试转换客户端缓存的视频
*/
func TestMacOS(t *testing.T) {
	client.Convert("/Users/zen/Movies")
}

/*
测试移除加密媒体的前缀
*/
func TestRemove(t *testing.T) {
	client.RemoveEncryptionHeader("/Users/zen/Movies/bilibili/28914616704/28914616704-1-30016.m4s")
	client.RemoveEncryptionHeader("/Users/zen/Movies/bilibili/28914616704/28914616704-1-30280.m4s")
	//	RemoveEncryptionHeader("/Users/zen/github/AVmerger/101391217/101391217-1-100027.m4s")
}

// go test -v -timeout 100h -run TestConvertClientCache
/*
测试转换客户端缓存的视频
*/
func TestConvertClientCache(t *testing.T) {
	switch runtime.GOOS {
	case "darwin":
		client.Convert("/Users/zen/Movies/bilibili")
	case "windows":
		client.Convert("C:\\Users\\zen\\Videos\\bilibili")
	}
}

/*
测试移除加密媒体的前缀
*/
func TestRemoveEncryptionPrefix(t *testing.T) {
	client.RemoveEncryptionHeader("/Users/zen/Movies/bilibili/28914616704/28914616704-1-30016.m4s")
	client.RemoveEncryptionHeader("/Users/zen/Movies/bilibili/28914616704/28914616704-1-30280.m4s")
	//	RemoveEncryptionHeader("/Users/zen/github/AVmerger/101391217/101391217-1-100027.m4s")
}
