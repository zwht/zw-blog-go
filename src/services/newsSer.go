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
	SeeSum      int    `json:"seeSum"`
	Index       string `json:"index"`
	Img         string `json:"img"`
	IsHot       int    `json:"isHot"`
	State       int    `json:"state"`
	Abstract    string `json:"abstract"`
	Labels      string `json:"labels"`
	ReviewSum   int    `json:"reviewSum"`
	Source      string `json:"source"`
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
	SeeSum      int    `json:"seeSum"`
	Index       string `json:"index"`
	Img         string `json:"img"`
	IsHot       int    `json:"isHot"`
	State       int    `json:"state"`
	Abstract    string `json:"abstract"`
	Labels      string `json:"labels"`
	ReviewSum   int    `json:"reviewSum"`
	TypeName    string `json:"typeName"`
	Source      string `json:"source"`
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
	Source      string `column:"and,news.source,="`
	// StartTime   time.Time `column:"and,createTime,between"`
	// EndTime     time.Time `column:"and,endTime,between"`
}

func (news *News) NewsInsert() (err error) {
	sql := "insert into news(id,url_en,title,content,create_time,author,type_id,img,user_group_id,state,abstract,labels,author_id,see_sum,source,review_sum) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), news.UrlEn, news.Title, news.Content, time.Now(), news.Author, news.TypeId, news.Img, news.UserGroupId, news.State, news.Abstract, news.Labels, news.AuthorId, news.SeeSum, news.Source, 0)
	return
}

func NewsDelete(id string) (err error) {
	sql := "delete from news where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (news *News) NewsUpdate() (err error) {
	sql := "update news set title=$1,content=$2,type_id=$3,img=$4,state=$5,abstract=$6,labels=$7,url_en=$8, source=$9 where id=$10"
	_, err = Db.Exec(sql, news.Title, news.Content, news.TypeId, news.Img, news.State, news.Abstract, news.Labels, news.UrlEn, news.Source, news.ID)
	return
}
func (news *News) NewsUpdateSum() (err error) {
	sql := "update news set see_sum=$1 where id=$2"
	_, err = Db.Exec(sql, news.SeeSum, news.ID)
	return
}
func (news *News) NewsUpdateState() (err error) {
	sql := "update news set state=$1 where id=$2"
	_, err = Db.Exec(sql, news.State, news.ID)
	return
}
func (news *News) NewsUpdateReviewSum() (err error) {
	sql := "update news set review_sum=review_sum+$1 where id=$2"
	_, err = Db.Exec(sql, news.ReviewSum, news.ID)
	return
}

func NewsSelect(id string) (news News, err error) {
	sql := "select id,url_en,title,content,create_time,author,type_id,img,state,abstract,labels,author_id,see_sum,source from news where id=$1"
	news = News{}
	err = Db.QueryRow(sql, id).Scan(&news.ID, &news.UrlEn, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.State, &news.Abstract, &news.Labels, &news.AuthorId, &news.SeeSum, &news.Source)
	return
}
func NewsSelectByUrl(urlEn string, time1 string, time2 string) (news News, err error) {
	sql := "select id,url_en,title,content,create_time,author,type_id,img,state,abstract,labels,author_id,see_sum,source,review_sum from news where url_en=$1 and create_time >= $2 and create_time < $3"
	news = News{}
	err = Db.QueryRow(sql, urlEn, time1, time2).Scan(&news.ID, &news.UrlEn, &news.Title, &news.Content, &news.CreateTime, &news.Author, &news.TypeId, &news.Img, &news.State, &news.Abstract, &news.Labels, &news.AuthorId, &news.SeeSum, &news.Source, &news.ReviewSum)
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
	sql, _ := ReplaceQuestionToDollarInherit("select news.id,news.url_en,news.title,news.create_time,news.author,news.img,news.state,news.abstract,news.labels,news_type.name,news.source,news.review_sum,news.see_sum from news join news_type on news.type_id = news_type.id "+whereStr+" limit ? offset ?", 0)
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
		err = rows.Scan(&news.ID, &news.UrlEn, &news.Title, &news.CreateTime, &news.Author, &news.Img, &news.State, &news.Abstract, &news.Labels, &news.TypeName, &news.Source, &news.ReviewSum, &news.SeeSum)
		if err != nil {
			panic(err.Error)
		}
		newss = append(newss, news)
	}
	return
}
func NewsSelectHotList(typeId string) (newss []NewsList, err error) {
	newss = []NewsList{}
	rows, err := Db.Query("select news.url_en,news.title,news.create_time,news.author,news.type_id,news.state,news.source,news.review_sum from news join news_type on news.type_id = news_type.id where news.type_id = $1 limit 10 offset 0", typeId)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var news NewsList
		err = rows.Scan(&news.UrlEn, &news.Title, &news.CreateTime, &news.Author, &news.TypeId, &news.State, &news.Source, &news.ReviewSum)
		if err != nil {
			panic(err.Error)
		}
		newss = append(newss, news)
	}
	return
}
