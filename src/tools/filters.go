package tools

import (
	"html/template"
	"time"
)

func RenderHtml(content string) template.HTML {
	return template.HTML(content)
}
func DateFk(date string) string {
	str, _ := time.Parse("2006-01-02 15:04:05", date)
	return str.Format("2006-01-02 15:04:05")
}
