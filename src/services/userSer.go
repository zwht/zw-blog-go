package service

import (
	. "../config"
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
type UserSearchUserVo struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Emai  string `json:"emai"`
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

func UserSelectList(pageSize int, pageNum int, search UserSearchUserVo) (users []User, err error) {
	sql := "select id,name,loginName,phone,email from _user limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa(pageSize*(pageNum-1))
	users = []User{}
	rows, err := Db.Query(sql)
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
