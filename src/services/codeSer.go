package service

import (
	. "../tools"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
)

type Code struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Groups      string `json:"groups"`
}
type CodeSearchVo struct {
	ID          string `column:"and,id,like"`
	Name        string `column:"and,name,like"`
	Description string `column:"and,description,like"`
	Code        string `column:"and,code,like"`
	Groups      string `column:"and,groups,like"`
}

func (code *Code) CodeInsert() (err error) {
	sql := "insert into code(id,name,description,groups,code) values($1,$2,$3,$4,$5)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), code.Name, code.Description, code.Groups, code.Code)
	return
}

func CodeDelete(id string) (err error) {
	sql := "delete from code where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (code *Code) CodeUpdate() (err error) {
	sql := "update code set name=$1,description=$2,groups=$3,code=$4 where id=$5"
	_, err = Db.Exec(sql, code.Name, code.Description, code.Groups, code.Code, code.ID)
	return
}

func CodeSelect(id string) (code Code, err error) {
	sql := "select id,name,description,groups,code from code where id=$1"
	code = Code{}
	err = Db.QueryRow(sql, id).Scan(&code.ID, &code.Name, &code.Description, &code.Groups, &code.Code)
	return
}
func CodeSelectCount(search CodeSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from code"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func CodeSelectList(pageSize int, pageNum int, search CodeSearchVo) (codes []Code, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,description,groups,code from code"+whereStr+" order by code desc limit ? offset ?", 0)
	fmt.Println(sql)
	codes = []Code{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)

	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var code Code
		err = rows.Scan(&code.ID, &code.Name, &code.Description, &code.Groups, &code.Code)
		if err != nil {
			panic(err.Error)
		}
		codes = append(codes, code)
	}
	return
}
