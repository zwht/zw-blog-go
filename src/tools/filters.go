package tools

import (
	"html/template"
)

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func RenderHtml(content string) template.HTML {
	return template.HTML(content)
}
func DateFk(str string) string {
	// str = "2018-10-30T23:18:11.673298Z"
	return Substr(str, 0, 10) + " " + Substr(str, 11, 8)
}
