package routersBlog

import (
	. "../../services"
	"errors"
	"github.com/kataras/iris/mvc"
	"strconv"
	"strings"
)

type BlogListController struct{}

func (c *BlogListController) GetBy(name string, pageSize string, pageNum string) mvc.Result {
	var newsSearchVo NewsSearchVo
	pageSize1, error1 := strconv.Atoi(pageSize)
	pageNum1, error2 := strconv.Atoi(pageNum)
	if error1 != nil {
		return mvc.Response{Err: errors.New("参数错误"), Code: 400}
	}
	if error2 != nil {
		return mvc.Response{Err: errors.New("参数错误"), Code: 400}
	}
	count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(pageSize1, pageNum1, newsSearchVo)
	var out []interface{}
	for i := 0; i < len(newss); i++ {
		var listNews ListNews
		listNews.ID = newss[i].ID
		listNews.Title = newss[i].Title
		listNews.CreateTime = newss[i].CreateTime
		listNews.UrlEn = newss[i].UrlEn
		listNews.TypeId = newss[i].TypeId
		listNews.Abstract = newss[i].Abstract
		listNews.Year = strings.Split(newss[i].CreateTime, "-")[0]
		listNews.Month = strings.Split(newss[i].CreateTime, "-")[1]
		out = append(out, listNews)
	}
	data := map[string]interface{}{
		"Title":    "列表",
		"List":     out,
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
