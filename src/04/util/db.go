package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func ini() {
	Db,err = sql.Open("mysql","root:123456@localhost:3306")
	if err != nil {
		panic(err.Error())
	}
}