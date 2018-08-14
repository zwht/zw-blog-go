package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type News struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	Author      string `json:"author"`
	TypeId      string `json:"typeId"`
	UserGroupId string `json:"userGroupId"`
	SeeSum      string `json:"seeSum"`
	Index       string `json:"index"`
	Img         string `json:"img"`
	IsHot       int    `json:"isHot"`
	State       int    `json:"state"`
	Abstract    string `json:"abstract"`
	Labels      string `json:"labels"`
	ReviewSum   int    `json:"reviewSum"`
}
type NewsSearchVo struct {
	ID          string `column:"and,id,="`
	Title       string `column:"and,title,like"`
	Author      string `column:"and,author,like"`
	TypeId      string `column:"and,typeId,like"`
	UserGroupId string `column:"and,userGroupId,like"`
	IsHot       int    `column:"and,isHot,="`
	State       int    `column:"and,state,="`
	// StartTime   time.Time `column:"and,createTime,between"`
	// EndTime     time.Time `column:"and,endTime,between"`
}

func (news *News) NewsInsert() (err error) {
	sql := "insert into news(id,title,content,create_time,author,type_id,img,user_group_id,state,abstract,labels) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), news.Title, news.Content, time.Now(), news.Author, news.TypeId, news.Img, news.UserGroupId, news.State, news.Abstract, news.Labels)
	return
}

func NewsDelete(id string) (err error) {
	sql := "delete from news where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (news *News) NewsUpdate() (err error) {
	sql := "update news set title=$1,content=$2,type_id=$3,img=$4,state=$5,abstract=$6,labels=$7 where id=$8"
	_, err = Db.Exec(sql, news.Title, news.Content, news.TypeId, news.Img, news.State, news.Abstract, news.Labels, news.ID)
	return
}

func NewsSelect(id string) (news News, err error) {
	sql := "select id,title,content,create_time,author,type_id,img,state,abstract,labels from news where id=$1"
	news = News{}
	err = Db.QueryRow(sql, id).Scan(&news.ID, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.State, &news.Abstract, &news.Labels)
	return
}
func NewsSelectCount(search NewsSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from news"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewsSelectList(pageSize int, pageNum int, search NewsSearchVo) (newss []News, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,title,content,create_time,author,type_id,img,user_group_id,state,abstract,labels from news "+whereStr+" limit ? offset ?", 0)
	newss = []News{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var news News
		err = rows.Scan(&news.ID, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.UserGroupId, &news.State, &news.Abstract, &news.Labels)
		if err != nil {
			panic(err.Error)
		}
		newss = append(newss, news)
	}
	return
}
