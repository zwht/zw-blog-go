package controllers

import (
	. "../../datamodels"
	. "../../services"
	. "../../tools/http"
	"strconv"
	"strings"
)

func NewsCreate(ctx *Context) {
	var news News
	ctx.ReadJSON(&news)
	news.Author = ctx.User.Name
	news.AuthorId = ctx.User.ID
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
