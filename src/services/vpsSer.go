package service

import (
	. "../datamodels"
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type Vps struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ip          string `json:"ip"`
	Address     string `json:"address"`
	CreateTime  int64  `json:"createTime"`
	OverdueTime int64  `json:"overdueTime"`
}
type VpsGet struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ip          string    `json:"ip"`
	Address     string    `json:"address"`
	CreateTime  LocalTime `json:"createTime"`
	OverdueTime LocalTime `json:"overdueTime"`
}

type VpsSearchVo struct {
	ID   string `column:"and,id,="`
	Name string `column:"and,name,like"`
	Ip   string `column:"and,ip,like"`
	// StartTime time.Time `column:"and,overdue_time,between"`
	// EndTime   time.Time `column:"and,overdue_time,between"`
}

func (vps *Vps) VpsInsert() (err error) {
	sql := "insert into vps(id,name,description,ip,address,create_time,overdue_time) values($1,$2,$3,$4,$5,$6,$7)"
	overdueTime := time.Unix(vps.OverdueTime/1000, 0)
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), vps.Name, vps.Description, vps.Ip, vps.Address, time.Now(), overdueTime)
	return
}

func VpsDelete(id string) (err error) {
	sql := "delete from vps where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (vps *Vps) VpsUpdate() (err error) {
	sql := "update vps set name=$1,description=$2,ip=$3,address=$4,overdue_time=$5 where id=$6"
	overdueTime := time.Unix(vps.OverdueTime/1000, 0)
	_, err = Db.Exec(sql, vps.Name, vps.Description, vps.Ip, vps.Address, overdueTime, vps.ID)
	return
}

func VpsSelect(id string) (vps VpsGet, err error) {
	sql := "select id,name,description,ip,address,create_time,overdue_time from vps where id=$1"
	vps = VpsGet{}
	err = Db.QueryRow(sql, id).Scan(&vps.ID, &vps.Name, &vps.Description, &vps.Ip, &vps.Address, &vps.CreateTime, &vps.OverdueTime)
	return
}
func VpsSelectCount(search VpsSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from vps"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func VpsSelectList(pageSize int, pageNum int, search VpsSearchVo) (vpss []VpsGet, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,name,description,ip,address,create_time,overdue_time from vps "+whereStr+" limit ? offset ?", 0)
	vpss = []VpsGet{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var vps VpsGet
		err = rows.Scan(&vps.ID, &vps.Name, &vps.Description, &vps.Ip, &vps.Address, &vps.CreateTime, &vps.OverdueTime)
		if err != nil {
			return
		}
		vpss = append(vpss, vps)
	}
	return
}
