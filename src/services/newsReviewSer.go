package service

import (
	. "../tools"
	"fmt"
	"strconv"
	"time"
)

type NewsReview struct {
	ID         string    `json:"id"`
	NewId      string    `json:"newId"`
	Content    string    `json:"content"`
	Ip         string    `json:"ip"`
	CreateTime time.Time `json:"createTime"`
	ParentId   string    `json:"parentId"`
	Email      string    `json:"email"`
	UserId     string    `json:"userId"`
	UserName   string    `json:"userName"`
	Url        string    `json:"url"`
	Img        string    `json:"img"`
	State      int       `json:"state"`
}
type NewsReviewSearchVo struct {
	ID      string `column:"and,id,="`
	NewId   string `column:"and,new_id,="`
	Content string `column:"and,content,like"`
	State   string `column:"and,state,="`
	Ip      string `column:"and,ip,="`
	Email   string `column:"and,email,="`
	// StartTime   time.Time `column:"and,createTime,between"`
	// EndTime     time.Time `column:"and,endTime,between"`
}

func (news_review *NewsReview) NewsReviewInsert() (err error) {
	sql := "insert into news_review(id,new_id,content,create_time,ip,user_id,parent_id,email,user_name,url,img,state) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_, err = Db.Exec(sql, news_review.ID, news_review.NewId, news_review.Content, news_review.CreateTime, news_review.Ip,
		news_review.UserId, news_review.ParentId, news_review.Email, news_review.UserName, news_review.Url, news_review.Img, news_review.State)
	return
}

func NewsReviewDelete(id string) (err error) {
	sql := "delete from news_review where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (news_review *NewsReview) NewsReviewUpdate() (err error) {
	sql := "update news_review set new_id=$1,content=$2,create_time=$3,ip=$4,user_id=$5,parent_id=$6,email=$7,user_name=$8,url=$9,img=$10,state=$11 where id=$12"
	_, err = Db.Exec(sql, news_review.NewId, news_review.Content, news_review.CreateTime, news_review.Ip, news_review.UserId, news_review.ParentId, news_review.Email, news_review.UserName, news_review.Url, news_review.Img, news_review.State, news_review.ID)
	return
}
func (news_review *NewsReview) NewsReviewUpdateState() (err error) {
	sql := "update news_review set state=$1 where id=$2"
	_, err = Db.Exec(sql, news_review.State, news_review.ID)
	return
}

func NewsReviewSelect(id string) (news_review NewsReview, err error) {
	sql := "select id,new_id,content,create_time,ip,user_id,parent_id,email,user_name,url,img,state from news_review where id=$1"
	news_review = NewsReview{}
	err = Db.QueryRow(sql, id).Scan(&news_review.ID, &news_review.NewId, &news_review.Content, &news_review.CreateTime, &news_review.Ip, &news_review.UserId, &news_review.ParentId, &news_review.Email, &news_review.UserName, &news_review.Url, &news_review.Img, &news_review.State)
	return
}
func NewsReviewSelectCount(search NewsReviewSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from news_review"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewsReviewSelectList(pageSize int, pageNum int, search NewsReviewSearchVo) (news_reviews []NewsReview, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,new_id,content,create_time,ip,user_id,parent_id,email,user_name,url,img,state from news_review"+whereStr+" order by create_time desc limit ? offset ?", 0)
	fmt.Println(sql)
	news_reviews = []NewsReview{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)

	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var news_review NewsReview
		err = rows.Scan(&news_review.ID, &news_review.NewId, &news_review.Content, &news_review.CreateTime, &news_review.Ip, &news_review.UserId, &news_review.ParentId, &news_review.Email, &news_review.UserName, &news_review.Url, &news_review.Img, &news_review.State)
		if err != nil {
			panic(err.Error)
		}
		news_reviews = append(news_reviews, news_review)
	}
	return
}
