package constant

// Package constant 定义了B站视频处理相关的常量
// 包括各种B站客户端的缓存目录和音视频输出目录的路径常量

// 定义了B站各版本客户端的缓存目录和音视频输出目录的常量
// 分为三类：
// 1. 主存储空间的B站客户端缓存目录（/sdcard/Android/data/...）
// 2. 第二存储空间的B站客户端缓存目录（/storage/emulated/999/...）
// 3. 音视频文件的输出目录（本地和远程）
const (
	// ===== 主存储空间的B站客户端缓存目录 =====

	// BILI 标准B站安卓客户端的下载缓存目录
	// 用于存放普通版B站APP下载的视频缓存文件
	BILI = "/sdcard/Android/data/tv.danmaku.bili/download"

	// HD B站HD版（平板）安卓客户端的下载缓存目录
	// 针对平板设备优化的HD版本的缓存目录
	HD = "/sdcard/Android/data/tv.danmaku.bilibilihd/download"

	// GLOBAL B站国际版安卓客户端的下载缓存目录
	// 用于存放哔哩哔哩国际版APP的下载缓存
	GLOBAL = "/sdcard/Android/data/com.bilibili.app.in/download"

	// BLUE B站海外版安卓客户端的下载缓存目录
	// 专门针对海外用户的版本的缓存目录
	BLUE = "/sdcard/Android/data/com.bilibili.app.blue/download"

	// ===== 第二存储空间的B站客户端缓存目录 =====

	// BILI999 标准B站安卓客户端在第二存储空间的下载缓存目录
	// 当设备支持多存储空间时，在第二存储空间中的缓存目录
	BILI999 = "/storage/emulated/999/Android/data/tv.danmaku.bili/download"

	// HD999 B站HD版安卓客户端在第二存储空间的下载缓存目录
	// HD版在第二存储空间中的缓存目录
	HD999 = "/storage/emulated/999/Android/data/tv.danmaku.bilibilihd/download"

	// GLOBAL999 B站国际版安卓客户端在第二存储空间的下载缓存目录
	// 国际版在第二存储空间中的缓存目录
	GLOBAL999 = "/storage/emulated/999/Android/data/com.bilibili.app.in/download"

	// BLUE999 B站海外版安卓客户端在第二存储空间的下载缓存目录
	// 海外版在第二存储空间中的缓存目录
	BLUE999 = "/storage/emulated/999/Android/data/com.bilibili.app.blue/download"

	// ===== 音视频文件的输出目录 =====

	// ANDROIDVIDEO 合并后视频文件的本地输出目录
	// 处理完成的视频文件将保存在此目录
	ANDROIDVIDEO = "/sdcard/Movies/bili_video"

	// ANDROIDAUDIO 提取的音频文件的本地输出目录
	// 从视频中提取的音频文件将保存在此目录
	ANDROIDAUDIO = "/sdcard/Music/bili_audio"

	// REMOTEVIDEO 远程服务器上的视频存储目录
	// 用于将处理后的视频文件同步到远程服务器
	REMOTEVIDEO = "/home/zen/ugreen/alist/bili/video"

	// REMOTEAUDIO 远程服务器上的音频存储目录
	// 用于将提取的音频文件同步到远程服务器
	REMOTEAUDIO = "/home/zen/ugreen/alist/bili/audio"
)
