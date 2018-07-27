package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func CodeCreate(ctx iris.Context) {
	var code Code
	ctx.ReadJSON(&code)

	err := code.CodeInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存code信息"
	}

	ctx.JSON(result)
}
func CodeUpdate(ctx iris.Context) {
	var code Code
	ctx.ReadJSON(&code)

	err := code.CodeUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存code信息"
	}

	ctx.JSON(result)
}

func CodeGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	code, err := CodeSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取code信息"
		result.Data = code
	}

	ctx.JSON(result)
}

func CodeGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var codeSearchVo CodeSearchVo
	ctx.ReadJSON(&codeSearchVo)
	count, _ := CodeSelectCount(codeSearchVo)
	codes, err := CodeSelectList(pageSize, pageNum, codeSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = codes

		result.Code = 200
		result.Msg = "成功获取code列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func CodeDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := CodeDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除code信息"
	}

	ctx.JSON(result)
}
