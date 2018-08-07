package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func NewsTypeCreate(ctx iris.Context) {
	var news NewsType
	ctx.ReadJSON(&news)

	err := news.NewsTypeInsert()

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
func NewsTypeUpdate(ctx iris.Context) {
	var news NewsType
	ctx.ReadJSON(&news)

	err := news.NewsTypeUpdate()

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

func NewsTypeGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	news, err := NewsTypeSelect(id)

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

func NewsTypeGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var newsSearchVo NewsTypeSearchVo
	ctx.ReadJSON(&newsSearchVo)
	count, _ := NewsTypeSelectCount(newsSearchVo)
	newss, err := NewsTypeSelectList(pageSize, pageNum, newsSearchVo)
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

func NewsTypeDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := NewsTypeDelete(id)

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
