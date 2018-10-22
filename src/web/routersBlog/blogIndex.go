package routersBlog

import (
	. "../../services"
	"github.com/kataras/iris/mvc"
)

type BlogIndexController struct{}

func (c *BlogIndexController) Get() mvc.Result {
	var newsSearchVo NewsSearchVo
	count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(10, 1, newsSearchVo)
	data := map[string]interface{}{
		"Title":    "列表",
		"List":     newss,
		"Count":    count,
		"PageSize": 10,
		"PageNum":  1,
	}
	var helloView = mvc.View{
		Name: "blog/pages/index.html",
		Data: data,
	}
	return helloView
}
