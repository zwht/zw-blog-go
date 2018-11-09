package routersBlog

import (
	. "../../services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
	"strings"
)

type ListNews struct {
	ID          string `json:"id"`
	UrlEn       string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	Author      string `json:"author"`
	TypeId      string `json:"typeId"`
	UserGroupId string `json:"userGroupId"`
	SeeSum      string `json:"seeSum"`
	Index       string `json:"index"`
	Img         string `json:"img"`
	IsHot       int    `json:"isHot"`
	State       int    `json:"state"`
	Abstract    string `json:"abstract"`
	Labels      string `json:"labels"`
	ReviewSum   int    `json:"reviewSum"`
	Year        string
	Month       string
}

type BlogIndexController struct{}

func (c *BlogIndexController) Get(ctx iris.Context) mvc.Result {
	pageSzie, _ := strconv.Atoi(ctx.FormValue("pageSzie"))
	pageNumber, _ := strconv.Atoi(ctx.FormValue("pageNumber"))
	state, _ := strconv.Atoi(ctx.FormValue("state"))
	typeId := ctx.FormValue("typeId")
	if pageSzie == 0 {
		pageSzie = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}
	if state == 0 {
		state = 1105
	}
	var newsSearchVo NewsSearchVo
	newsSearchVo.State = state
	newsSearchVo.TypeId = typeId
	count, _ := NewsSelectCount(newsSearchVo)
	newss, _ := NewsSelectList(pageSzie, pageNumber, newsSearchVo)

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
		"PageSize": pageSzie,
		"PageNum":  pageNumber,
	}
	var helloView = mvc.View{
		Name: "blog/pages/index.html",
		Data: data,
	}
	return helloView
}
