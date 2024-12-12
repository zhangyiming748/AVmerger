package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RsyncDir(localDir, remoteDir, user, server, password string) error {
	// 使用sshpass来传递密码并构建rsync命令
	args := []string{"sshpass", "-p", password, "rsync", "-avz", localDir, fmt.Sprintf("%s@%s:%s", user, server, remoteDir)}
	cmd := exec.Command("bash", "-c", strings.Join(args, " "))
	log.Printf("最终的rsync命令是:%s\n", cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
