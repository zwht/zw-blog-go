package service

import (
	. "../config"
)

func SelectCount(tableName string) (count int, err error) {
	sql := "select count(*) from " + tableName
	err = Db.QueryRow(sql).Scan(&count)
	return
}
