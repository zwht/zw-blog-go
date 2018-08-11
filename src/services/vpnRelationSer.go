package service

import (
	. "../tools"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type VpnRelation struct {
	ID          string `json:"id"`
	VpnId       string `json:"vpnId"`
	UserId      string `json:"userId"`
	CreateTime  int64  `json:"createTime"`
	OverdueTime int64  `json:"overdueTime"`
}
type VpnRelationSearchVo struct {
	ID        string    `column:"and,id,="`
	UserId    string    `column:"and,userId,like"`
	VpnId     string    `column:"and,vpnId,like"`
	StartTime time.Time `column:"and,overdueTime,between"`
	EndTime   time.Time `column:"and,overdueTime,between"`
}

func (vpnRelation *VpnRelation) VpnRelationInsert() (err error) {
	sql := "insert into vpn_relation(id,vpn_id,user_id,create_time,overdue_time) values($1,$2,$3,$4,$5)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), vpnRelation.VpnId, vpnRelation.UserId, vpnRelation.CreateTime, vpnRelation.OverdueTime)
	return
}

func VpnRelationDelete(id string) (err error) {
	sql := "delete from vpn_relation where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (vpnRelation *VpnRelation) VpnRelationUpdate() (err error) {
	sql := "update vpn_relation set vpn_id=$1,user_id=$2,create_time=$3,overdue_time=$4 where id=$5"
	_, err = Db.Exec(sql, vpnRelation.VpnId, vpnRelation.UserId, vpnRelation.CreateTime, vpnRelation.OverdueTime, vpnRelation.ID)
	return
}

func VpnRelationSelect(id string) (vpnRelation VpnRelation, err error) {
	sql := "select id,vpn_id,user_id,create_time,overdue_time from vpn_relation where id=$1"
	vpnRelation = VpnRelation{}
	err = Db.QueryRow(sql, id).Scan(&vpnRelation.ID, &vpnRelation.VpnId, &vpnRelation.UserId, &vpnRelation.CreateTime, &vpnRelation.OverdueTime)
	return
}
func VpnRelationSelectCount(search VpnRelationSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from vpn_relation"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func VpnRelationSelectList(pageSize int, pageNum int, search VpnRelationSearchVo) (vpnRelations []VpnRelation, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,vpn_id,user_id,create_time,overdue_time from vpn_relation "+whereStr+" limit ? offset ?", 0)
	vpnRelations = []VpnRelation{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var vpnRelation VpnRelation
		err = rows.Scan(&vpnRelation.ID, &vpnRelation.VpnId, &vpnRelation.UserId, &vpnRelation.CreateTime, &vpnRelation.OverdueTime)
		if err != nil {
			panic(err.Error)
		}
		vpnRelations = append(vpnRelations, vpnRelation)
	}
	return
}
