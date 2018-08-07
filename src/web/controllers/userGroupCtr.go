package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func UserGroupCreate(ctx iris.Context) {
	var userGroup UserGroup
	ctx.ReadJSON(&userGroup)

	err := userGroup.UserGroupInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存userGroup信息"
	}

	ctx.JSON(result)
}
func UserGroupUpdate(ctx iris.Context) {
	var userGroup UserGroup
	ctx.ReadJSON(&userGroup)

	err := userGroup.UserGroupUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存userGroup信息"
	}

	ctx.JSON(result)
}

func UserGroupGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userGroup, err := UserGroupSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取userGroup信息"
		result.Data = userGroup
	}

	ctx.JSON(result)
}

func UserGroupGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var userGroupSearchVo UserGroupSearchVo
	ctx.ReadJSON(&userGroupSearchVo)
	count, _ := UserGroupSelectCount(userGroupSearchVo)
	userGroups, err := UserGroupSelectList(pageSize, pageNum, userGroupSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = userGroups

		result.Code = 200
		result.Msg = "成功获取userGroup列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func UserGroupDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := UserGroupDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除userGroup信息"
	}

	ctx.JSON(result)
}
