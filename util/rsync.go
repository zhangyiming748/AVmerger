package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func UploadWithRsync(localDir string) error {
	remoteDir := "/home/zen/ugreen/alist/bili/" // 服务器上的目标目录，请根据实际情况修改
	user := "zen"                               // 服务器用户名，请根据实际情况修改
	server := "192.168.1.9"
	password := "123456"
	// 强烈建议不要在代码中直接使用密码，这里只是为了演示
	// 实际应用中应该从环境变量或配置文件读取密码
	// password := os.Getenv("RSYNC_PASSWORD")

	cmdStr := fmt.Sprintf("sshpass -p '%s' rsync -vzrc --progress %s %s@%s:%s", password, localDir, user, server, remoteDir)
	cmd := exec.Command("bash", "-c", cmdStr)

	log.Printf("执行的 rsync 命令是: %s\n", cmd.String())

	// 将命令的标准输出和标准错误输出连接到当前进程，以便实时查看进度
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("rsync 上传失败: %v", err)
	}

	log.Printf("成功上传文件夹 %s 到 %s@%s:%s\n", localDir, user, server, remoteDir)
	return nil
}
