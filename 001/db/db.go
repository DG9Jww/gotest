package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db		*sql.DB
	err		error
)

func init(){
	Db,err = sql.Open("mysql","root:845819308@tcp(localhost:3306)/gin")
	if err != nil{
		panic(err)
	}
}