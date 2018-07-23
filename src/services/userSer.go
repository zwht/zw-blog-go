package service

import (
	. "../config"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type LoginVo struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
type LoginUser struct {
	User
	Token string `json:"token"`
}

func (user *User) Insert() (err error) {
	sql := "insert into _user(id,loginName,name,password,email,phone) values($1,$2,$3,$4,$5,$6)"
	_, err = Db.Exec(sql, "123", user.LoginName, user.Name, user.Password, user.Email, user.Phone)
	return
}

func Delete(id string) (err error) {
	sql := "delete from _user where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (user *User) Update() (err error) {
	sql := "update _user set name=$1,Password=$2 where id=$3"
	_, err = Db.Exec(sql, user.Name, user.Password, user.ID)
	return
}

func Select(id string) (user User, err error) {
	sql := "select id,name,password from _user where id=$1"
	user = User{}
	err = Db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Password)
	return
}

func SelectList() (users []User, err error) {
	sql := "select id,name,password from _user"
	users = []User{}
	rows, err := Db.Query(sql)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Password)
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
