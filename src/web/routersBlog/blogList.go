package routersBlog

import (
	. "../../services"
	"github.com/kataras/iris/mvc"
)

type BlogListController struct{}

func (c *BlogListController) GetBy(name string, pageSize int64, pageNum int64) mvc.Result {
	var newsSearchVo NewsSearchVo
	count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(10, 1, newsSearchVo)
	data := map[string]interface{}{
		"Title":    "列表",
		"List":     newss,
		"Count":    count,
		"PageSize": pageSize,
		"PageNum":  pageNum,
	}
	var helloView = mvc.View{
		Name: "blog/pages/index.html",
		Data: data,
	}
	return helloView
}

// var errBadName = errors.New("bad name")
// var badName = mvc.Response{Err: errBadName, Code: 400}