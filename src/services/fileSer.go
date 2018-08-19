package service

import (
	. "../tools"
	"strconv"
)

type FileVo struct {
	ID     string `json:"id"`
	Type   int    `json:"type"`
	UserId string `json:"userId"`
}
type FileGetVo struct {
	ID     string `json:"id"`
	Type   int    `json:"type"`
	UserId string `json:"userId"`
}

type FileSearchVo struct {
	ID     string `column:"and,id,="`
	Type   int    `column:"and,type,="`
	UserId string `column:"and,userId,="`
	// StartTime time.Time `column:"and,overdue_time,between"`
	// EndTime   time.Time `column:"and,overdue_time,between"`
}

func (file *FileVo) FileInsert() (err error) {
	sql := "insert into files(id,type,user_id) values($1,$2,$3)"
	_, err = Db.Exec(sql, file.ID, file.Type, file.UserId)
	return
}

func FileDelete(id string) (err error) {
	sql := "delete from files where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (file *FileVo) FileUpdate() (err error) {
	sql := "update files set type=$1,user_id=$2 where id=$3"
	_, err = Db.Exec(sql, file.Type, file.UserId)
	return
}

func FileSelect(id string) (file FileGetVo, err error) {
	sql := "select id,type,user_id from files where id=$1"
	file = FileGetVo{}
	err = Db.QueryRow(sql, id).Scan(&file.ID, &file.Type, &file.UserId)
	return
}
func FileSelectCount(search FileSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from files"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func FileSelectList(pageSize int, pageNum int, search FileSearchVo) (files []FileGetVo, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,type,user_id from files "+whereStr+" limit ? offset ?", 0)
	files = []FileGetVo{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var file FileGetVo
		err = rows.Scan(&file.ID, &file.Type, &file.UserId)
		if err != nil {
			return
		}
		files = append(files, file)
	}
	return
}
