package service

import (
	. "../config"
	"github.com/satori/go.uuid"
	"strconv"
)

type Code struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type SearchCodeVo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (code *Code) CodeInsert() (err error) {
	sql := "insert into _code(id,name,description) values($1,$2,$3)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), code.Name, code.Description)
	return
}

func CodeDelete(id string) (err error) {
	sql := "delete from _code where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (code *Code) CodeUpdate() (err error) {
	sql := "update _code set name=$1,description=$2 where id=$3"
	_, err = Db.Exec(sql, code.Name, code.Description, code.ID)
	return
}

func CodeSelect(id string) (code Code, err error) {
	sql := "select id,name,description from _code where id=$1"
	code = Code{}
	err = Db.QueryRow(sql, id).Scan(&code.ID, &code.Name, &code.Description)
	return
}

func CodeSelectList(pageSize int, pageNum int, search SearchCodeVo) (codes []Code, err error) {
	sql := "select id,name,description from _code limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa(pageSize*(pageNum-1))
	codes = []Code{}
	rows, err := Db.Query(sql)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var code Code
		err = rows.Scan(&code.ID, &code.Name, &code.Description)
		if err != nil {
			panic(err.Error)
		}
		codes = append(codes, code)
	}
	return
}
