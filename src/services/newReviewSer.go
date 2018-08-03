package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type NewReview struct {
	ID         string `json:"id"`
	UserId     string `json:"userId"`
	NewId      string `json:"newId"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Ip         string `json:"ip"`
}
type NewReviewSearchVo struct {
	ID        string    `column:"and,id,="`
	UserId    string    `column:"and,userId,like"`
	NewId     string    `column:"and,newId,like"`
	Content   string    `column:"and,content,like"`
	StartTime time.Time `column:"and,createTime,between"`
	EndTime   time.Time `column:"and,endTime,between"`
	Ip        string    `column:"and,ip,="`
}

func (newReview *NewReview) NewReviewInsert() (err error) {
	sql := "insert into new_review(id,user_id,new_id,content,create_time,ip) values($1,$2,$3,$4,$5,$6)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), newReview.UserId, newReview.NewId, newReview.Content, newReview.CreateTime, newReview.Ip)
	return
}

func NewReviewDelete(id string) (err error) {
	sql := "delete from new_review where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func NewReviewSelect(id string) (newReview NewReview, err error) {
	sql := "select id,user_id,new_id,content,create_time,ip from new_review where id=$1"
	newReview = NewReview{}
	err = Db.QueryRow(sql, id).Scan(&newReview.ID, &newReview.UserId, &newReview.NewId, &newReview.Content, &newReview.CreateTime, &newReview.Ip)
	return
}
func NewReviewSelectCount(search NewReviewSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from new_review"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewReviewSelectList(pageSize int, pageNum int, search NewReviewSearchVo) (newReviews []NewReview, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,user_id,new_id,content,create_time,ip from new_review "+whereStr+" limit ? offset ?", 0)
	newReviews = []NewReview{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var newReview NewReview
		err = rows.Scan(&newReview.ID, &newReview.UserId, &newReview.NewId, &newReview.Content, &newReview.CreateTime, &newReview.Ip)
		if err != nil {
			panic(err.Error)
		}
		newReviews = append(newReviews, newReview)
	}
	return
}
