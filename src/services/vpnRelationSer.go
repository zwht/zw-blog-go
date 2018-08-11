package service

import (
	. "../datamodels"
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
type VpnRelationGet struct {
	ID          string    `json:"id"`
	VpnId       string    `json:"vpnId"`
	UserId      string    `json:"userId"`
	State       int       `json:"state"`
	CreateTime  LocalTime `json:"createTime"`
	OverdueTime LocalTime `json:"overdueTime"`
}

type VpnRelationUserList struct {
	RelationId  string    `json:"relationId"`
	VpnId       string    `json:"vpnId"`
	UserId      string    `json:"userId"`
	LoginName   string    `json:"loginName"`
	Password    string    `json:"password"`
	Ip          string    `json:"ip"`
	State       int       `json:"state"`
	OverdueTime LocalTime `json:"overdueTime"`
	CreateTime  LocalTime `json:"createTime"`
}

type VpnRelationSearchVo struct {
	ID     string `column:"and,id,="`
	UserId string `column:"and,userId,like"`
	VpnId  string `column:"and,vpnId,like"`
	// StartTime time.Time `column:"and,overdueTime,between"`
	// EndTime   time.Time `column:"and,overdueTime,between"`
}

func (vpnRelation *VpnRelation) VpnRelationInsert() (err error) {
	sql := "insert into vpn_relation(id,vpn_id,user_id,create_time,overdue_time,state) values($1,$2,$3,$4,$5,$6)"
	_, err = Db.Exec(sql, uuid.Must(uuid.NewV4()), vpnRelation.VpnId, vpnRelation.UserId, time.Now(), time.Unix(vpnRelation.OverdueTime/1000, 0), 2101)
	return
}

func VpnRelationDelete(id string) (err error) {
	sql := "delete from vpn_relation where id=$1"
	_, err = Db.Exec(sql, id)
	return
}

func (vpnRelation *VpnRelation) VpnRelationUpdate() (err error) {
	sql := "update vpn_relation set vpn_id=$1,user_id=$2,overdue_time=$3 where id=$4"
	_, err = Db.Exec(sql, vpnRelation.VpnId, vpnRelation.UserId, vpnRelation.OverdueTime, time.Unix(vpnRelation.OverdueTime/1000, 0), vpnRelation.ID)
	return
}
func VpnRelationUpdateState(id1 string, state int) (err error) {
	sql := "update vpn_relation set state=$1 where id=$2"
	_, err = Db.Exec(sql, state, id1)
	return
}

func VpnRelationSelect(id string) (vpnRelation VpnRelationGet, err error) {
	sql := "select id,vpn_id,user_id,create_time,overdue_time from vpn_relation where id=$1"
	vpnRelation = VpnRelationGet{}
	err = Db.QueryRow(sql, id).Scan(&vpnRelation.ID, &vpnRelation.VpnId, &vpnRelation.UserId, &vpnRelation.CreateTime, &vpnRelation.OverdueTime)
	return
}
func VpnRelationSelectCount(search VpnRelationSearchVo) (count int, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select count(*) from vpn_relation"+whereStr, 0)
	err = Db.QueryRow(sql, args...).Scan(&count)
	return
}
func VpnRelationSelectList(pageSize int, pageNum int, search VpnRelationSearchVo) (vpnRelations []VpnRelationGet, err error) {
	whereStr, args := GenWhereByStruct(search)
	sql, _ := ReplaceQuestionToDollarInherit("select id,vpn_id,user_id,create_time,overdue_time from vpn_relation "+whereStr+" limit ? offset ?", 0)
	vpnRelations = []VpnRelationGet{}
	args = append(args, strconv.Itoa(pageSize), strconv.Itoa(pageSize*(pageNum-1)))
	rows, err := Db.Query(sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var vpnRelation VpnRelationGet
		err = rows.Scan(&vpnRelation.ID, &vpnRelation.VpnId, &vpnRelation.UserId, &vpnRelation.CreateTime, &vpnRelation.OverdueTime)
		if err != nil {
			return
		}
		vpnRelations = append(vpnRelations, vpnRelation)
	}
	return
}

//根据用户id查询关联vpn账号
func VpnRelationSelectListByUserId(id string) (vpnRelations []VpnRelationUserList, err error) {
	sql := "select vpn.id,vpn_relation.id,vpn_relation.user_id,vpn.name,vpn.password,vpn.ip,vpn_relation.overdue_time,vpn_relation.create_time,vpn_relation.state from vpn join vpn_relation on vpn.id = vpn_relation.vpn_id where vpn_relation.user_id=$1"
	rows, err := Db.Query(sql, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Columns()
		var vpnRelation VpnRelationUserList
		err = rows.Scan(&vpnRelation.VpnId, &vpnRelation.RelationId, &vpnRelation.UserId, &vpnRelation.LoginName, &vpnRelation.Password, &vpnRelation.Ip, &vpnRelation.OverdueTime, &vpnRelation.CreateTime, &vpnRelation.State)
		if err != nil {
			return
		}
		vpnRelations = append(vpnRelations, vpnRelation)
	}
	return
}
