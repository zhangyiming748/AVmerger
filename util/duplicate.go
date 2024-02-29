package util

/*
根据输入的字符串切片返回去重的字符串切片
*/
func DuplicateBySlice(elements []string) (dup []string) {
	m := make(map[string]bool)
	for _, element := range elements {
		if _, ok := m[element]; !ok { // 如果元素不在map中，则添加到result和map中
			dup = append(dup, element)
			m[element] = true
		}
	}

	return dup
}
