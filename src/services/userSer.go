package service

import (
	. "../models"
	. "../tools"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
)

func UserInsert(user User) (err error) {
	sql := "insert into _user(id,loginName,name,password,email,phone,roles,state) values($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), user.LoginName, user.Name, user.Password, user.Email, user.Phone, user.Roles, true)
	return
}

func UserDelete(id string) (err error) {
	sql := "delete from _user where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func UserUpdate(user User) (err error) {
	sql := "update _user set name=$1,LoginName=$2,Phone=$3,Email=$4,Roles=$5 where id=$6"
	_, err = Db.Exec(sql, user.Name, user.LoginName, user.Phone, user.Email, user.Roles, user.ID)
	return
}
func UserUpdateState(id string, state bool) (err error) {
	sql := "update _user set state=$1 where id=$2"
	_, err = Db.Exec(sql, state, id)
	return
}
func UserUpdatePassoword(id string, oldPassword string, passoword string) (err error) {
	sql := "update _user set password=$1 where id=$2 and password=$3"
	_, err = Db.Exec(sql, passoword, id, oldPassword)
	return
}

func UserSelect(id string) (user User, err error) {
	sql := "select id,name,loginName,phone,email,roles,state from _user where id=$1"
	user = User{}
	err = Db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email, &user.Roles, &user.State)
	return
}
func UserSelectByName(loginName string) (user User, err error) {
	sql := "select id,name,loginName,phone,email,roles,state from _user where loginName=$1"
	user = User{}
	err = Db.QueryRow(sql, loginName).Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email, &user.Roles, &user.State)
	return
}

func UserSelectCount(search UserSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from _user"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}

func UserSelectList(pageSize int, pageNum int, search UserSearchVo) (users []User, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,loginName,phone,email,roles,state from _user "+whereStr+" limit ? offset ?", 0)
	fmt.Println(sql)
	users = []User{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email, &user.Roles, &user.State)
		if err != nil {
			panic(err.Error)
		}
		users = append(users, user)
	}
	return
}

func Logins(loginVo LoginVo) (user User, err error) {
	sql := "select id,name,loginName,phone,email,roles,state from _user where loginName=$1 AND password=$2"
	user = User{}
	err = Db.QueryRow(sql, loginVo.LoginName, loginVo.Password).Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email, &user.Roles, &user.State)
	return
}
