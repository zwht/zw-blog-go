package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func VpsCreate(ctx iris.Context) {
	var vps Vps
	ctx.ReadJSON(&vps)
	err := vps.VpsInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存vps信息"
	}

	ctx.JSON(result)
}
func VpsUpdate(ctx iris.Context) {
	var vps Vps
	ctx.ReadJSON(&vps)

	err := vps.VpsUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存vps信息"
	}

	ctx.JSON(result)
}

func VpsGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	vps, err := VpsSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取vps信息"
		result.Data = vps
	}

	ctx.JSON(result)
}

func VpsGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var vpsSearchVo VpsSearchVo
	ctx.ReadJSON(&vpsSearchVo)
	count, _ := VpsSelectCount(vpsSearchVo)
	vpss, err := VpsSelectList(pageSize, pageNum, vpsSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = vpss
		result.Code = 200
		result.Msg = "成功获取vps列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func VpsDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := VpsDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除vps信息"
	}

	ctx.JSON(result)
}
