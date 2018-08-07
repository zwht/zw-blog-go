package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
	"strconv"
)

func NewsReviewCreate(ctx iris.Context) {
	var newsReview NewsReview
	ctx.ReadJSON(&newsReview)

	err := newsReview.NewsReviewInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存newsReview信息"
	}

	ctx.JSON(result)
}

func NewsReviewGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	newsReview, err := NewsReviewSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取newsReview信息"
		result.Data = newsReview
	}

	ctx.JSON(result)
}

func NewsReviewGetList(ctx iris.Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var newsReviewSearchVo NewsReviewSearchVo
	ctx.ReadJSON(&newsReviewSearchVo)
	count, _ := NewsReviewSelectCount(newsReviewSearchVo)
	newsReviews, err := NewsReviewSelectList(pageSize, pageNum, newsReviewSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = newsReviews

		result.Code = 200
		result.Msg = "成功获取newsReview列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func NewsReviewDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := NewsReviewDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除newsReview信息"
	}

	ctx.JSON(result)
}
