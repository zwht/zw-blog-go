package service

import (
	. "../config"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
type UserSearchVo struct {
	Phone string `column:"and,phone,like"`
	Name  string `column:"and,name,like"`
	Emai  string `column:"and,email,like"`
}
type LoginVo struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
type LoginUser struct {
	User
	Token string `json:"token"`
}

func (user *User) UserInsert() (err error) {
	sql := "insert into _user(id,loginName,name,password,email,phone) values($1,$2,$3,$4,$5,$6)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), user.LoginName, user.Name, user.Password, user.Email, user.Phone)
	return
}

func UserDelete(id string) (err error) {
	sql := "delete from _user where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (user *User) UserUpdate() (err error) {
	sql := "update _user set name=$1,LoginName=$2,Phone=$3,Email=$4 where id=$5"
	_, err = Db.Exec(sql, user.Name, user.LoginName, user.Phone, user.Email, user.ID)
	return
}

func UserSelect(id string) (user User, err error) {
	sql := "select id,name,loginName,phone,email from _user where id=$1"
	user = User{}
	err = Db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email)
	return
}

func UserSelectList(pageSize int, pageNum int, search UserSearchVo) (users []User, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, key := ReplaceQuestionToDollarInherit("select id,name,loginName,phone,email from _user "+whereStr+" limit ? offset ?", 0)
	fmt.Println(sql, key)
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
		err = rows.Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email)
		if err != nil {
			panic(err.Error)
		}
		users = append(users, user)
	}
	return
}

func Logins(loginVo LoginVo) (user User, err error) {
	sql := "select id,name,loginName,phone,email from _user where loginName=$1 AND password=$2"
	user = User{}
	err = Db.QueryRow(sql, loginVo.LoginName, loginVo.Password).Scan(&user.ID, &user.Name, &user.LoginName, &user.Phone, &user.Email)
	return
}
