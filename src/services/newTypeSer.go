package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
)

type NewType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    string `json:"parentId"`
	Index       int    `json:"index"`
}
type NewTypeSearchVo struct {
	ID          string `column:"and,id,="`
	Name        string `column:"and,name,like"`
	Description string `column:"and,description,like"`
	ParentId    string `column:"and,parentId,like"`
	Index       int    `column:"and,index,="`
}

func (newType *NewType) NewTypeInsert() (err error) {
	sql := "insert into new_type(id,name,description,parent_id,index) values($1,$2,$3,$4,$5)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), newType.Name, newType.Description, newType.ParentId, newType.Index)
	return
}

func NewTypeDelete(id string) (err error) {
	sql := "delete from new_type where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (newType *NewType) NewTypeUpdate() (err error) {
	sql := "update new_type set name=$1,description=$2,parent_id=$3,index=$4 where id=$5"
	_, err = Db.Exec(sql, newType.Name, newType.Description, newType.ParentId, newType.Index, newType.ID)
	return
}

func NewTypeSelect(id string) (newType NewType, err error) {
	sql := "select id,name,description,parent_id,index from new_type where id=$1"
	newType = NewType{}
	err = Db.QueryRow(sql, id).Scan(&newType.ID, &newType.Name, &newType.Description, &newType.ParentId, &newType.Index)
	return
}
func NewTypeSelectCount(search NewTypeSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from new_type"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func NewTypeSelectList(pageSize int, pageNum int, search NewTypeSearchVo) (newTypes []NewType, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,description,parent_id,index from new_type "+whereStr+" limit ? offset ?", 0)
	newTypes = []NewType{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var newType NewType
		err = rows.Scan(&newType.ID, &newType.Name, &newType.Description, &newType.ParentId, &newType.Index)
		if err != nil {
			panic(err.Error)
		}
		newTypes = append(newTypes, newType)
	}
	return
}
