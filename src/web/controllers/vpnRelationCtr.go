package controllers

import (
	. "../../datamodels"
	. "../../services"
	. "../../tools/http"
	"strconv"
)

func VpnRelationCreate(ctx *Context) {
	var vpnRelation VpnRelation
	ctx.ReadJSON(&vpnRelation)

	err := vpnRelation.VpnRelationInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		err1 := VpnUpdateState(vpnRelation.VpnId, 2002)
		if err1 != nil {
			result.Code = 0
			result.Msg = err.Error()
		} else {
			result.Code = 200
			result.Msg = "成功保存vpnRelation信息"
		}
	}

	ctx.JSON(result)
}
func VpnRelationUpdate(ctx *Context) {
	var vpnRelation VpnRelation
	ctx.ReadJSON(&vpnRelation)

	err := vpnRelation.VpnRelationUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存vpnRelation信息"
	}

	ctx.JSON(result)
}
func VpnRelationUpdateStateCtr(ctx *Context) {
	id := ctx.URLParam("id")
	vpnId := ctx.URLParam("vpnId")
	state, _ := strconv.Atoi(ctx.URLParam("state"))
	err := VpnRelationUpdateState(id, state)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		if state == 2102 {
			err1 := VpnUpdateState(vpnId, 2001)
			if err1 != nil {
				result.Code = 0
				result.Msg = err.Error()
			} else {
				result.Code = 200
				result.Msg = "成功保存vpnRelation信息"
			}
		} else {
			result.Code = 200
			result.Msg = "成功保存vpnRelation信息"
		}
	}

	ctx.JSON(result)
}

func VpnRelationGetById(ctx *Context) {
	id := ctx.Params().Get("id")
	vpnRelation, err := VpnRelationSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取vpnRelation信息"
		result.Data = vpnRelation
	}

	ctx.JSON(result)
}

func VpnRelationGetList(ctx *Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var vpnRelationSearchVo VpnRelationSearchVo
	ctx.ReadJSON(&vpnRelationSearchVo)
	count, _ := VpnRelationSelectCount(vpnRelationSearchVo)
	vpnRelations, err := VpnRelationSelectList(pageSize, pageNum, vpnRelationSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = vpnRelations

		result.Code = 200
		result.Msg = "成功获取vpnRelation列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func VpnRelationDeleteById(ctx *Context) {
	id := ctx.Params().Get("id")
	err := VpnRelationDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除vpnRelation信息"
	}

	ctx.JSON(result)
}

func VpnRelationGetListByUserId(ctx *Context) {
	id := ctx.URLParam("id")
	vpnRelations, err := VpnRelationSelectListByUserId(id)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Data = vpnRelations
		result.Code = 200
		result.Msg = "成功获取vpnRelation列表信息"
	}

	ctx.JSON(result)
}
