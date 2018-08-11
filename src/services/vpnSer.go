package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type Vpn struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Ip         string `json:"ip"`
	VpsId      string `json:"vpsId"`
	CreateTime int64  `json:"createTime"`
	LockTime   int64  `json:"lockTime"`
	State      bool   `json:"state"`
}

type VpnSearchVo struct {
	ID        string    `column:"and,id,="`
	Name      string    `column:"and,name,like"`
	Ip        string    `column:"and,ip,like"`
	StartTime time.Time `column:"and,lockTime,between"`
	EndTime   time.Time `column:"and,lockTime,between"`
}

func (vpn *Vpn) VpnInsert() (err error) {
	sql := "insert into vpn(id,name,password,ip,vps_id,create_time,lock_time.state) values($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), vpn.Name, vpn.Password, vpn.Ip, vpn.VpsId, vpn.CreateTime, vpn.LockTime, vpn.State)
	return
}

func VpnDelete(id string) (err error) {
	sql := "delete from vpn where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (vpn *Vpn) VpnUpdate() (err error) {
	sql := "update vpn set name=$1,password=$2,ip=$3,vps_id=$4,create_time=$5,lock_time=$6,state=$7 where id=$8"
	_, err = Db.Exec(sql, vpn.Name, vpn.Password, vpn.Ip, vpn.VpsId, vpn.CreateTime, vpn.LockTime, vpn.State, vpn.ID)
	return
}

func VpnSelect(id string) (vpn Vpn, err error) {
	sql := "select id,name,password,ip,vps_id,create_time,lock_time,state from vpn where id=$1"
	vpn = Vpn{}
	err = Db.QueryRow(sql, id).Scan(&vpn.ID, &vpn.Name, &vpn.Password, &vpn.Ip, &vpn.VpsId, &vpn.CreateTime, &vpn.LockTime, &vpn.State)
	return
}

func VpnSelectCount(search VpnSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from vpn"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}

func VpnSelectList(pageSize int, pageNum int, search VpnSearchVo) (vpns []Vpn, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,password,ip,vps_id,create_time,lock_time,state from vpn "+whereStr+" limit ? offset ?", 0)
	vpns = []Vpn{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var vpn Vpn
		err = rows.Scan(&vpn.ID, &vpn.Name, &vpn.Password, &vpn.Ip, &vpn.VpsId, &vpn.CreateTime, &vpn.LockTime, &vpn.State)
		if err != nil {
			panic(err.Error)
		}
		vpns = append(vpns, vpn)
	}
	return
}
