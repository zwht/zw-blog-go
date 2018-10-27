package tools

import (
	"html/template"
)

func RenderHtml(content string) template.HTML {
	return template.HTML(content)
}
