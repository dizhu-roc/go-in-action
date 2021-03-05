package words

import "strings"

func CountWord(text string) (count int) {
	// 将字符串按照空格分割，转换成 []string 类型的切片
	count = len(strings.Fields(text))
	return
}
