package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "120.79.171.251"
	port     = 5432
	user     = "zhaowei"
	password = "123"
	dbname   = "blogDev"
)

var Db *sql.DB

func init() {
	var err error

	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", pgInfo)

	if err != nil {
		panic(err)
	}
}
