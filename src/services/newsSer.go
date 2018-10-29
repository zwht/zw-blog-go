package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type News struct {
	ID          string `json:"id"`
	UrlEn       string `json:"urlEn"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	Author      string `json:"author"`
	AuthorId    string `json:"authorId"`
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
type NewsList struct {
	ID          string `json:"id"`
	UrlEn       string `json:"urlEn"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	Author      string `json:"author"`
	AuthorId    string `json:"authorId"`
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
	TypeName    string `json:"typeName"`
}
type NewsSearchVo struct {
	ID          string `column:"and,news.id,="`
	Title       string `column:"and,news.title,like"`
	Author      string `column:"and,news.author,like"`
	TypeId      string `column:"and,news.type_id,like"`
	UserGroupId string `column:"and,news.user_group_id,like"`
	IsHot       int    `column:"and,news.is_hot,="`
	State       int    `column:"and,news.state,="`
	AuthorId    string `column:"and,news.author_id,="`
	// StartTime   time.Time `column:"and,createTime,between"`
	// EndTime     time.Time `column:"and,endTime,between"`
}

func (news *News) NewsInsert() (err error) {
	sql := "insert into news(id,url_en,title,content,create_time,author,type_id,img,user_group_id,state,abstract,labels,author_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), news.UrlEn, news.Title, news.Content, time.Now(), news.Author, news.TypeId, news.Img, news.UserGroupId, news.State, news.Abstract, news.Labels, news.AuthorId)
	return
}

func NewsDelete(id string) (err error) {
	sql := "delete from news where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (news *News) NewsUpdate() (err error) {
	sql := "update news set title=$1,content=$2,type_id=$3,img=$4,state=$5,abstract=$6,labels=$7,url_en=$8 where id=$9"
	_, err = Db.Exec(sql, news.Title, news.Content, news.TypeId, news.Img, news.State, news.Abstract, news.Labels, news.UrlEn, news.ID)
	return
}

func NewsSelect(id string) (news News, err error) {
	sql := "select id,url_en,title,content,create_time,author,type_id,img,state,abstract,labels,author_id from news where id=$1"
	news = News{}
	err = Db.QueryRow(sql, id).Scan(&news.ID, &news.UrlEn, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.State, &news.Abstract, &news.Labels, &news.AuthorId)
	return
}
func NewsSelectByUrl(urlEn string, time1 string, time2 string) (news News, err error) {
	sql := "select id,url_en,title,content,create_time,author,type_id,img,state,abstract,labels from news where url_en=$1 and create_time >= $2 and create_time < $3"
	news = News{}
	err = Db.QueryRow(sql, urlEn, time1, time2).Scan(&news.ID, &news.UrlEn, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.State, &news.Abstract, &news.Labels)
	return
}
func NewsSelectCount(search NewsSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from news"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewsSelectList(pageSize int, pageNum int, search NewsSearchVo) (newss []NewsList, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select news.id,news.url_en,news.title,news.content,news.create_time,news.author,news.type_id,news.img,news.user_group_id,news.state,news.abstract,news.labels,new_type.name from news join new_type on news.type_id = new_type.id "+whereStr+" limit ? offset ?", 0)
	println(sql)
	newss = []NewsList{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var news NewsList
		err = rows.Scan(&news.ID, &news.UrlEn, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.UserGroupId, &news.State, &news.Abstract, &news.Labels, &news.TypeName)
		if err != nil {
			panic(err.Error)
		}
		newss = append(newss, news)
	}
	return
}
