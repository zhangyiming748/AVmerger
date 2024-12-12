package util

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestMaster

func TestMaster(t *testing.T) {
	localDir := "/mnt/c/Users/zen/Github/AVmerger" // 本地要传送的目录路径
	remoteDir := "/home/zen/rsync"                 // 远程服务器上的目标目录路径
	user := "zen"
	server := "192.168.1.9"
	passwd := "163453"

	if err := RsyncDir(localDir, remoteDir, user, server, passwd); err != nil {
		fmt.Println("Rsync failed:", err)
		os.Exit(1)
	}
}
