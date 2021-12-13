package model

import (
	"learnGin/001/db"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//If err != nil which means success
func Add(todo *Todo) error {
	sqlStr := "insert into todo(id,title,status) values(?,?,?)"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(todo.ID,todo.Title,todo.Status)
	if err != nil {
		return err
	}
	return nil
}

//If err != nil which means success
func Delete(id string) error {
	sqlStr := "delete from todo where id = ?"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//Get All Infomation
func QueryAll() ([]*Todo,error) {
	sqlStr := "select * from todo"
	rows,err := db.Db.Query(sqlStr)
	if err != nil {
		return nil,err
	}
	todo_list := []*Todo{}
	for rows.Next(){
		single_todo := &Todo{}
		err := rows.Scan(&single_todo.ID,&single_todo.Title,&single_todo.Status)
		if err != nil{
			return nil,err
		}
		todo_list = append(todo_list, single_todo)
	}
	return todo_list,nil
}



//Get One Info
func QueryOne(id string) (*Todo,error) {
	sqlStr := "select * from todo where id = ?"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	var todo = new(Todo)
	err = stmt.QueryRow(id).Scan(&todo.ID,&todo.Title,&todo.Status)
	if err != nil{
		return nil,err
	}
	return todo,nil
}


//Update Info
func Update(todo *Todo) error {
	sqlStr := "update todo set status = ? where id = ?"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	_,err = stmt.Exec(todo.Status,todo.ID)
	if err != nil{
		return err
	}
	return nil
}