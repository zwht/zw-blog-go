package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type NewsReview struct {
	ID         string `json:"id"`
	ParentId   string `json:"parentId"`
	UserId     string `json:"userId"`
	NewsId     string `json:"newId"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Ip         string `json:"ip"`
}
type NewsReviewSearchVo struct {
	ID        string    `column:"and,id,="`
	ParentId  string    `column:"and,parentId,like"`
	UserId    string    `column:"and,userId,like"`
	NewsId    string    `column:"and,newId,like"`
	Content   string    `column:"and,content,like"`
	StartTime time.Time `column:"and,createTime,between"`
	EndTime   time.Time `column:"and,endTime,between"`
	Ip        string    `column:"and,ip,="`
}

func (newsReview *NewsReview) NewsReviewInsert() (err error) {
	sql := "insert into new_review(id,user_id,new_id,content,create_time,ip,parent_id) values($1,$2,$3,$4,$5,$6,$7)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), newsReview.UserId, newsReview.NewsId, newsReview.Content, newsReview.CreateTime, newsReview.Ip, newsReview.ParentId)
	return
}

func NewsReviewDelete(id string) (err error) {
	sql := "delete from new_review where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func NewsReviewSelect(id string) (newsReview NewsReview, err error) {
	sql := "select id,user_id,new_id,content,create_time,ip,parent_id from new_review where id=$1"
	newsReview = NewsReview{}
	err = Db.QueryRow(sql, id).Scan(&newsReview.ID, &newsReview.UserId, &newsReview.NewsId, &newsReview.Content, &newsReview.CreateTime, &newsReview.Ip, &newsReview.ParentId)
	return
}
func NewsReviewSelectCount(search NewsReviewSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from new_review"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewsReviewSelectList(pageSize int, pageNum int, search NewsReviewSearchVo) (newsReviews []NewsReview, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,user_id,new_id,content,create_time,ip,parent_id from new_review "+whereStr+" limit ? offset ?", 0)
	newsReviews = []NewsReview{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var newsReview NewsReview
		err = rows.Scan(&newsReview.ID, &newsReview.UserId, &newsReview.NewsId, &newsReview.Content, &newsReview.CreateTime, &newsReview.Ip, &newsReview.ParentId)
		if err != nil {
			panic(err.Error)
		}
		newsReviews = append(newsReviews, newsReview)
	}
	return
}
