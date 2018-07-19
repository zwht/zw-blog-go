package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "zhaowei"
	password = "123"
	dbname   = "zhaowei"
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
