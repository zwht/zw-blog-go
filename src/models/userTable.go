package models

import ()

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Roles     string `json:"roles"`
	State     bool   `json:"state"`
	ParentId  string `json:"parentId"`
}
type UserSearchVo struct {
	Phone     string `column:"and,phone,like"`
	Name      string `column:"and,name,like"`
	Email     string `column:"and,email,like"`
	LoginName string `column:"and,loginName,like"`
	Roles     string `column:"and,roles,like"`
}
type SearchRepeatVo struct {
	Phone     string `column:"or,phone,="`
	Email     string `column:"or,email,="`
	LoginName string `column:"or,loginName,="`
}
type LoginVo struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
type LoginUser struct {
	User
	Token string `json:"token"`
}
