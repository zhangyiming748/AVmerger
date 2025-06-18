package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// CheckRsync 检查系统中是否安装了rsync命令
// 如果未安装rsync，程序将终止并提示安装
func CheckRsync() {
	// 检查 rsync 命令是否存在
	if _, err := exec.LookPath("rsync"); err != nil {
		log.Fatal("未找到 rsync 命令，请先安装 rsync")
	}
	log.Println("系统环境检查通过: rsync 命令可用")
}

// CheckSshpass 检查系统中是否安装了sshpass命令
// sshpass用于在非交互式环境中提供SSH密码
// 如果未安装sshpass，程序将终止并提示安装
func CheckSshpass() {
	if _, err := exec.LookPath("sshpass"); err != nil {
		log.Fatal("未找到 sshpass 命令，请先安装 sshpass")
	}
	log.Println("系统环境检查通过: sshpass 命令可用")
}

// UploadWithRsyncAll 批量上传多个本地目录到远程目录
// 参数:
//   - remoteDir: 远程服务器上的目标目录
//   - localDirs: 要上传的本地目录列表（可变参数）
//
// 返回值:
//   - error: 如果任何目录上传失败，返回错误信息
func UploadWithRsyncAll(remoteDir string, localDirs ...string) error {
	for _, dir := range localDirs {
		if err := UploadWithRsync(dir, remoteDir); err != nil {
			return fmt.Errorf("上传文件夹 %s 失败: %v\n", dir, err)
		}
	}
	return nil
}

// UploadWithRsync 使用rsync将单个本地目录上传到远程服务器
// 参数:
//   - localDir: 要上传的本地目录路径
//   - remoteDir: 远程服务器上的目标目录
//
// 返回值:
//   - error: 如果上传过程中发生错误，返回错误信息
func UploadWithRsync(localDir, remoteDir string) error {
	// 远程服务器配置
	user := "zen"            // 服务器用户名
	server := "192.168.2.10" // 服务器地址
	password := "163453"     // 服务器密码
	// 注意：实际生产环境中应该从环境变量或配置文件读取敏感信息
	// password := os.Getenv("RSYNC_PASSWORD")

	// 构建rsync命令
	// -v: 显示详细信息
	// -z: 传输时压缩
	// -r: 递归处理目录
	// -c: 基于校验和比较文件
	// --progress: 显示传输进度
	cmdStr := fmt.Sprintf("sshpass -p '%s' rsync -vzrc --progress %s %s@%s:%s", password, localDir, user, server, remoteDir)
	cmd := exec.Command("bash", "-c", cmdStr)

	log.Printf("执行的 rsync 命令是: %s\n", cmd.String())

	// 将命令的输出重定向到当前进程的标准输出和标准错误
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行rsync命令
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("rsync 上传失败: %v", err)
	}

	log.Printf("成功上传文件夹 %s 到 %s@%s:%s\n", localDir, user, server, remoteDir)
	return nil
}

// OnTermux 检查当前程序是否在Termux环境中运行
// 返回值:
//   - bool: 如果在Termux环境中返回true，否则返回false
func OnTermux() bool {
	// 通过检查Termux特有的环境变量来判断
	if _, exists := os.LookupEnv("TERMUX_VERSION"); exists {
		return true
	}
	return false
}
