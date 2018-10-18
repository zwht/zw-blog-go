package routersBlog

import (
	. "../../services"
	"github.com/kataras/iris/mvc"
)

type BlogDetailController struct{}

func (c *BlogDetailController) Get() mvc.Result {
	var newsSearchVo NewsSearchVo
	// count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(10, 1, newsSearchVo)
	data := map[string]interface{}{
		"Title":     "Hello Page",
		"list":      newss,
		"MyMessage": "<div>999</div>",
	}
	var helloView = mvc.View{
		Name: "blog/pages/detail.html",
		Data: data,
	}
	return helloView
}

// var errBadName = errors.New("bad name")
// var badName = mvc.Response{Err: errBadName, Code: 400}

// func (c *BlogListController) GetBy(name string) mvc.Result {
// 	if name != "iris" {
// 		return badName
// 		// or
// 		// GetBy(name string) (mvc.Result, error) {
// 		//	return nil, errBadName
// 		// }
// 	}
// 	// return mvc.Response{Text: "Hello " + name} OR:
// 	return mvc.View{
// 		Name: "blog/pages/index.html",
// 		Data: map[string]interface{}{
// 			"Title":     "Hello Page",
// 			"MyMessage": "Welcome to my awesome website",
// 		},
// 	}
// }
