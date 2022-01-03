package db

import (
	"dachuang/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	username := config.GetValue("mysql", "username")
	password := config.GetValue("mysql", "password")
	host := config.GetValue("mysql", "host")
	port := config.GetValue("mysql", "port")
	database := config.GetValue("mysql", "database")

	Db, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		panic(err)
	}
}
