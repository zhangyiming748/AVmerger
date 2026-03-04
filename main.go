package main

import (
	"log"

	"AVmerger/core"

	"github.com/spf13/cobra"
)

var (
	src     string
	dst     string
	archive string
)

var rootCmd = &cobra.Command{
	Use:   "AVmerger",
	Short: "哔哩哔哩缓存视频合并工具",
	Long:  `AVmerger 是一个用于合并哔哩哔哩（Bilibili）手机版缓存视频的命令行工具`,
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "处理 B 站客户端缓存视频",
	Long:  `从 B 站客户端缓存目录中转换并处理视频文件到目标目录`,
	Run: func(cmd *cobra.Command, args []string) {
		if src == dst {
			log.Fatal("src 不能和 dst 相同，程序运行后 src 目录会被删除")
		}
		core.Client(src, dst)
		if archive != "" {
			core.ClassifyAfterMerge(dst, archive, nil)
		}
	},
}

var android2pcCmd = &cobra.Command{
	Use:   "android2pc",
	Short: "转换安卓客户端下载目录",
	Long:  `将安卓客户端下载目录中的音视频文件合并并转换到 PC 端格式`,
	Run: func(cmd *cobra.Command, args []string) {
		if src == "" {
			log.Fatal("安卓源目录 (--src) 必须指定")
		}
		if dst == "" {
			log.Fatal("目标输出目录 (--dst) 必须指定")
		}
		if src == dst {
			log.Fatal("src 不能和 dst 相同，程序运行后 src 目录会被删除")
		}
		core.Android2PC(src, dst)
		if archive != "" {
			core.ClassifyAfterMerge(dst, archive, nil)
		}
	},
}

func init() {
	// 为 client 命令添加标志
	clientCmd.Flags().StringVarP(&src, "src", "i", "", "B 站客户端缓存目录基础路径 (可选，为空则使用默认路径)")
	clientCmd.Flags().StringVarP(&dst, "dst", "o", "", "输出目录基础路径 (必填)")
	clientCmd.Flags().StringVarP(&archive, "archive", "a", "", "归档目录基础路径 (可选，用于分类整理合并后的文件)")
	clientCmd.MarkFlagRequired("dst")

	// 为 android2pc 命令添加标志
	android2pcCmd.Flags().StringVarP(&src, "src", "i", "", "安卓客户端下载目录路径 (必填)")
	android2pcCmd.Flags().StringVarP(&dst, "dst", "o", "", "输出目录基础路径 (必填)")
	android2pcCmd.Flags().StringVarP(&archive, "archive", "a", "", "归档目录基础路径 (可选，用于分类整理合并后的文件)")
	android2pcCmd.MarkFlagRequired("src")
	android2pcCmd.MarkFlagRequired("dst")

	// 将子命令添加到根命令
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(android2pcCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
