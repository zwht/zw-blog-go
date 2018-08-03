package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func NewTypeCreate(ctx iris.Context) {
	var newType NewType
	ctx.ReadJSON(&newType)

	err := newType.NewTypeInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存newType信息"
	}

	ctx.JSON(result)
}
func NewTypeUpdate(ctx iris.Context) {
	var newType NewType
	ctx.ReadJSON(&newType)

	err := newType.NewTypeUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存newType信息"
	}

	ctx.JSON(result)
}

func NewTypeGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	newType, err := NewTypeSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取newType信息"
		result.Data = newType
	}

	ctx.JSON(result)
}

func NewTypeGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var newTypeSearchVo NewTypeSearchVo
	ctx.ReadJSON(&newTypeSearchVo)
	count, _ := NewTypeSelectCount(newTypeSearchVo)
	newTypes, err := NewTypeSelectList(pageSize, pageNum, newTypeSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = newTypes

		result.Code = 200
		result.Msg = "成功获取newType列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func NewTypeDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := NewTypeDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除newType信息"
	}

	ctx.JSON(result)
}
