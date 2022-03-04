package db

import (
	"database/sql"
	"errors"
	"fmt"
	"website/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)
var Err_PREPARE_FAILED = errors.New("SQL PREPARE FAILED")

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

//增加，删除，更新
func Exec(sqlStr string, args ...interface{}) error {
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return Err_PREPARE_FAILED
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Println("1",err)
		return err
	}
	return nil
}

//查询多条结果
func Query(sqlStr string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return nil, Err_PREPARE_FAILED
	}
	return stmt.Query(args...)
}

//查询单条结果
func QueryRow(sqlStr string, args ...interface{}) (*sql.Row, error) {
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil, Err_PREPARE_FAILED
	}
	row := stmt.QueryRow(args...)
	return row, nil
}

//查询所有,不需要预编译
func QueryAll(sqlStr string) (*sql.Rows, error) {
	return Db.Query(sqlStr)
}
