package util

// Root 存储项目的根目录路径
// 用于在整个应用程序中保持一致的文件路径引用
var Root string

// SetRoot 设置项目的根目录路径
// 参数:
//   - r: 要设置的根目录路径
func SetRoot(r string) {
	Root = r
}

// GetRoot 获取当前设置的根目录路径
// 返回值:
//   - string: 项目的根目录路径
func GetRoot() string {
	return Root
}
