package controllers

import (
	. "../../datamodels"
	. "../../services"
	. "../../tools"
	. "../../tools/http"
	"github.com/kataras/iris"
	"strconv"
	"strings"
)

func NewsCreate(ctx *Context) {
	var news News
	ctx.ReadJSON(&news)
	news.Author = ctx.User.Name
	news.AuthorId = ctx.User.ID
	news.SeeSum = 0
	err := news.NewsInsert()
	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存news信息"
	}

	ctx.JSON(result)
}
func NewsUpdate(ctx *Context) {
	var news News
	ctx.ReadJSON(&news)

	err := news.NewsUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存news信息"
	}

	ctx.JSON(result)
}

// 添加文章访问数量
func NewsUpdateSum(ctx iris.Context) {
	result := Result{}

	ip := ctx.RemoteAddr()
	if ip == "::1" {
		result.Code = 0
		result.Msg = "localhost访问，不加访问数量"
		ctx.JSON(result)
		return
	}
	id := ctx.FormValue("id")
	// 获取redis对象，时间为5分钟
	sess := GetSess(5).Start(ctx)
	// 查看ip+id是否已经有，已经有不用加预览数量
	hasSee := sess.GetString(ip + id)
	if hasSee != "" {
		result.Code = 0
		result.Msg = "5分钟内访问，不加访问数量"
		ctx.JSON(result)
		return
	}
	oldNews, err := NewsSelect(id)
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
		ctx.JSON(result)
		return
	}
	var news News
	news.ID = id
	news.SeeSum = oldNews.SeeSum + 1
	err1 := news.NewsUpdateSum()
	if err1 != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		sess.Set(ip+id, "ok")
		result.Code = 200
		result.Msg = "成功添加一次访问数量"
	}
	ctx.JSON(result)
}

// 修改文章状态
func NewsUpdateStateCtr(ctx *Context) {
	result := Result{}
	id := ctx.FormValue("id")
	state, _ := strconv.Atoi(ctx.FormValue("state"))
	var news News
	news.ID = id
	news.State = state
	err1 := news.NewsUpdateState()
	if err1 != nil {
		result.Code = 0
		result.Msg = err1.Error()
	} else {
		result.Code = 200
		result.Msg = "修改状态成功"
	}
	ctx.JSON(result)
}

func NewsGetById(ctx *Context) {
	id := ctx.Params().Get("id")
	news, err := NewsSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取news信息"
		result.Data = news
	}

	ctx.JSON(result)
}

func NewsGetList(ctx *Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var newsSearchVo NewsSearchVo
	ctx.ReadJSON(&newsSearchVo)
	if !strings.Contains(ctx.User.Roles, "1001") {
		newsSearchVo.AuthorId = ctx.User.ID
	}

	count, _ := NewsSelectCount(newsSearchVo)
	newss, err := NewsSelectList(pageSize, pageNum, newsSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = newss

		result.Code = 200
		result.Msg = "成功获取news列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func NewsGetHotList(ctx *Context) {
	typeId := ctx.FormValue("typeId")
	newss, err := NewsSelectHotList(typeId)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功"
		result.Data = newss
	}
	ctx.JSON(result)
}

func NewsDeleteById(ctx *Context) {
	id := ctx.Params().Get("id")
	err := NewsDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除news信息"
	}

	ctx.JSON(result)
}
