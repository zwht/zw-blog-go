package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func VpnCreate(ctx iris.Context) {
	var vpn Vpn
	ctx.ReadJSON(&vpn)

	err := vpn.VpnInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存vpn信息"
	}

	ctx.JSON(result)
}
func VpnUpdate(ctx iris.Context) {
	var vpn Vpn
	ctx.ReadJSON(&vpn)

	err := vpn.VpnUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存vpn信息"
	}

	ctx.JSON(result)
}

func VpnGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	vpn, err := VpnSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取vpn信息"
		result.Data = vpn
	}

	ctx.JSON(result)
}

func VpnGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var vpnSearchVo VpnSearchVo
	ctx.ReadJSON(&vpnSearchVo)
	count, _ := VpnSelectCount(vpnSearchVo)
	vpns, err := VpnSelectList(pageSize, pageNum, vpnSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = vpns

		result.Code = 200
		result.Msg = "成功获取vpn列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func VpnDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := VpnDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除vpn信息"
	}

	ctx.JSON(result)
}
