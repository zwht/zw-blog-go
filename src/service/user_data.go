package service

import (
	. "../data"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"password"`
}

func (user *User) Insert() (err error) {
	sql := "insert into _user(name,age) values($1,$2)"
	_, err = Db.Exec(sql, user.Name, user.Age)
	return
}

func Delete(id int) (err error) {
	sql := "delete from _user where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (user *User) Update() (err error) {
	sql := "update _user set name=$1,age=$2 where id=$3"
	_, err = Db.Exec(sql, user.Name, user.Age, user.ID)
	return
}

func Select(id int) (user User, err error) {
	sql := "select id,name,age from _user where id=$1"
	user = User{}
	err = Db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Age)
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
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error)
		}
		users = append(users, user)
	}
	return
}
