package constant

const (
	BILI           = "/sdcard/Android/data/tv.danmaku.bili/download"
	HD             = "/sdcard/Android/data/tv.danmaku.bilibilihd/download"
	GLOBAL         = "/sdcard/Android/data/com.bilibili.app.in/download"
	BLUE           = "/sdcard/Android/data/com.bilibili.app.blue/download"
	ANDROIDVIDEO   = "/sdcard/Movies/bili"
	ANDROIDAUDIO   = "/sdcard/Music/bili"
	ANDROIDDANMAKU = "/sdcard/Documents/bili"
)

var LogLevel string

func GetLogLevel() string {
	return LogLevel
}

var SecondParameter string

func SetSecParam(s string) {
	SecondParameter = s
}
func GetSecParam() string {
	return SecondParameter
}
