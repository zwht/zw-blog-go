package routersBlog

import (
	. "../../services"
	"errors"
	"github.com/kataras/iris/mvc"
)

// HelloController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type HelloController struct{}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *HelloController) Get() mvc.Result {
	var newsSearchVo NewsSearchVo
	// count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(10, 1, newsSearchVo)
	newss[0].Content = "<div>999</div>"
	data := map[string]interface{}{
		"Title":     "Hello Page",
		"list":      newss,
		"MyMessage": "Welcome to my awesome website",
	}
	var helloView = mvc.View{
		Name: "blog/pages/index.html",
		Data: data,
	}
	return helloView
}

// you can define a standard error in order to re-use anywhere in your app.
var errBadName = errors.New("bad name")

// you can just return it as error or even better
// wrap this error with an mvc.Response to make it an mvc.Result compatible type.
var badName = mvc.Response{Err: errBadName, Code: 400}

// GetBy returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/hello/iris
// curl -i http://localhost:8080/hello/anything
func (c *HelloController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
		// or
		// GetBy(name string) (mvc.Result, error) {
		//	return nil, errBadName
		// }
	}

	// return mvc.Response{Text: "Hello " + name} OR:
	return mvc.View{
		Name: "blog/pages/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}
