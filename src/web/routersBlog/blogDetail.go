package routersBlog

import (
	. "../../services"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type BlogDetailController struct{}
type TreeStruct struct {
	NewsReview
	Children []TreeStruct
}

// 根据月份和urlen去查询新闻详情
func (c *BlogDetailController) GetBy(ctx iris.Context, year string, month string, urlEn string) mvc.Result {
	month1, _ := strconv.Atoi(month)
	year1, _ := strconv.Atoi(year)
	if month1 == 12 {
		month1 = 1
		year1 = year1 + 1
	} else {
		month1 = month1 + 1
	}
	detail, err := NewsSelectByUrl(urlEn, year+"-"+month+"-01", strconv.Itoa(year1)+"-"+strconv.Itoa(month1)+"-01")
	if err != nil {
		return mvc.Response{Err: errors.New("未找到数据"), Code: 400}
	}
	if detail.State != 1105 {
		return mvc.Response{Err: errors.New("未找到数据!"), Code: 400}
	}
	var news_reviewSearchVo NewsReviewSearchVo
	news_reviewSearchVo.NewId = detail.ID
	news_reviewSearchVo.State = 1203
	reviewList, reviewErr := NewsReviewSelectList(10000, 1, news_reviewSearchVo)
	if reviewErr != nil {
	}

	newReviewList := []TreeStruct{}
	for _, v := range reviewList {
		var item TreeStruct
		item.ID = v.ID
		item.NewId = v.NewId
		item.Content = v.Content
		item.Ip = v.Ip
		item.CreateTime = v.CreateTime
		item.ParentId = v.ParentId
		item.Email = v.Email
		item.UserId = v.UserId
		item.UserName = v.UserName
		item.Url = v.Url
		item.Img = v.Img
		item.State = v.State
		newReviewList = append(newReviewList, item)
	}
	for _, v := range newReviewList {
		if v.ParentId != "" {
			for i, v1 := range newReviewList {
				if v1.ID == v.ParentId {
					var newR = []TreeStruct{}
					newR = append(newR, v)
					newReviewList[i].Children = append(newR, v1.Children...)
				}
			}
		}
	}
	newReviewList1 := []TreeStruct{}
	for _, v := range newReviewList {
		if v.ParentId == "" {
			newReviewList1 = append(newReviewList1, v)
		}
	}
	fmt.Println(newReviewList1)
	data := map[string]interface{}{
		"Detail":  detail,
		"Reviews": newReviewList1,
	}
	var helloView = mvc.View{
		Name: "blog/pages/detail.html",
		Data: data,
	}
	return helloView
}
