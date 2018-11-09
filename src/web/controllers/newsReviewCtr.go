package controllers

import (
	. "../../datamodels"
	. "../../services"
	. "../../tools/http"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

func NewsReviewCreate(ctx *Context) {
	var news_review NewsReview
	ctx.ReadJSON(&news_review)
	news_review.State = 1201
	id := uuid.Must(uuid.NewV4()).String()
	news_review.ID = id
	news_review.CreateTime = time.Now()

	ip := ctx.RemoteAddr()
	if ip == "::1" {
		ip = ""
	}
	news_review.Ip = ip

	err := news_review.NewsReviewInsert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存news_review信息"
		result.Data = news_review
	}

	ctx.JSON(result)
}
func NewsReviewUpdate(ctx *Context) {
	var news_review NewsReview
	ctx.ReadJSON(&news_review)

	err := news_review.NewsReviewUpdate()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功保存news_review信息"
	}

	ctx.JSON(result)
}

func NewsReviewUpdateState(ctx *Context) {
	result := Result{}
	id := ctx.FormValue("id")
	newsId := ctx.FormValue("newsId")
	key, _ := strconv.Atoi(ctx.FormValue("key"))
	state, _ := strconv.Atoi(ctx.FormValue("state"))
	var newsReview NewsReview
	newsReview.ID = id
	newsReview.State = state
	err1 := newsReview.NewsReviewUpdateState()
	if err1 != nil {
		result.Code = 0
		result.Msg = err1.Error()
	} else {
		var news News
		news.ID = newsId
		news.ReviewSum = key
		news.NewsUpdateReviewSum()
		result.Code = 200
		result.Msg = "修改状态成功"
	}
	ctx.JSON(result)
}

func NewsReviewGetById(ctx *Context) {
	id := ctx.Params().Get("id")
	news_review, err := NewsReviewSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取news_review信息"
		result.Data = news_review
	}

	ctx.JSON(result)
}

func NewsReviewGetList(ctx *Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var news_reviewSearchVo NewsReviewSearchVo
	ctx.ReadJSON(&news_reviewSearchVo)
	count, _ := NewsReviewSelectCount(news_reviewSearchVo)
	news_reviews, err := NewsReviewSelectList(pageSize, pageNum, news_reviewSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = news_reviews

		result.Code = 200
		result.Msg = "成功获取news_review列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func NewsReviewDeleteById(ctx *Context) {
	id := ctx.Params().Get("id")
	err := NewsReviewDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除news_review信息"
	}

	ctx.JSON(result)
}
