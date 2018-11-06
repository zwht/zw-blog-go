package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
)

type NewsType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    string `json:"parentId"`
	UserGroupId string `json:"userGroupId"`
	Index       int    `json:"index"`
}
type NewsTypeSearchVo struct {
	ID          string `column:"and,id,="`
	Name        string `column:"and,name,like"`
	Description string `column:"and,description,like"`
	ParentId    string `column:"and,parentId,like"`
	UserGroupId string `column:"and,userGroupId,like"`
	Index       int    `column:"and,index,="`
}

func (newsType *NewsType) NewsTypeInsert() (err error) {
	sql := "insert into news_type(id,name,description,parent_id,index,user_group_id) values($1,$2,$3,$4,$5,$6)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), newsType.Name, newsType.Description, newsType.ParentId, newsType.Index, newsType.UserGroupId)
	return
}

func NewsTypeDelete(id string) (err error) {
	sql := "delete from news_type where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (newsType *NewsType) NewsTypeUpdate() (err error) {
	sql := "update news_type set name=$1,description=$2,parent_id=$3,index=$4 where id=$5"
	_, err = Db.Exec(sql, newsType.Name, newsType.Description, newsType.ParentId, newsType.Index, newsType.ID)
	return
}

func NewsTypeSelect(id string) (newsType NewsType, err error) {
	sql := "select id,name,description,parent_id,index from news_type where id=$1"
	newsType = NewsType{}
	err = Db.QueryRow(sql, id).Scan(&newsType.ID, &newsType.Name, &newsType.Description, &newsType.ParentId, &newsType.Index)
	return
}
func NewsTypeSelectCount(search NewsTypeSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from news_type"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewsTypeSelectList(pageSize int, pageNum int, search NewsTypeSearchVo) (newsTypes []NewsType, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,description,parent_id,index from news_type "+whereStr+" limit ? offset ?", 0)
	newsTypes = []NewsType{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var newsType NewsType
		err = rows.Scan(&newsType.ID, &newsType.Name, &newsType.Description, &newsType.ParentId, &newsType.Index)
		if err != nil {
			panic(err.Error)
		}
		newsTypes = append(newsTypes, newsType)
	}
	return
}
