package routersBlog

import (
	. "../../services"
	"errors"
	"github.com/kataras/iris/mvc"
)

type BlogDetailController struct{}

func (c *BlogDetailController) GetBy(id string) mvc.Result {

	detail, err := NewsSelect(id)
	if err != nil {
		return mvc.Response{Err: errors.New("错误"), Code: 400}
	}
	data := map[string]interface{}{
		"Detail": detail,
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
