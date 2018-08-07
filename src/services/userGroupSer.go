package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
)

type UserGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type UserGroupSearchVo struct {
	ID          string `column:"and,id,="`
	Name        string `column:"and,name,like"`
	Description string `column:"and,description,like"`
}

func (userGroup *UserGroup) UserGroupInsert() (err error) {
	sql := "insert into users_group(id,name,description) values($1,$2,$3)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), userGroup.Name, userGroup.Description)
	return
}

func UserGroupDelete(id string) (err error) {
	sql := "delete from users_group where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (userGroup *UserGroup) UserGroupUpdate() (err error) {
	sql := "update users_group set name=$1,description=$2 where id=$3"
	_, err = Db.Exec(sql, userGroup.Name, userGroup.Description, userGroup.ID)
	return
}

func UserGroupSelect(id string) (userGroup UserGroup, err error) {
	sql := "select id,name,description from users_group where id=$1"
	userGroup = UserGroup{}
	err = Db.QueryRow(sql, id).Scan(&userGroup.ID, &userGroup.Name, &userGroup.Description)
	return
}
func UserGroupSelectCount(search UserGroupSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from users_group"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func UserGroupSelectList(pageSize int, pageNum int, search UserGroupSearchVo) (userGroups []UserGroup, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,description from users_group "+whereStr+" limit ? offset ?", 0)
	userGroups = []UserGroup{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var userGroup UserGroup
		err = rows.Scan(&userGroup.ID, &userGroup.Name, &userGroup.Description)
		if err != nil {
			panic(err.Error)
		}
		userGroups = append(userGroups, userGroup)
	}
	return
}
