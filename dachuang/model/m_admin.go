package model

import (
	"dachuang/db"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Salt	 string
	PassWord string `json:"password"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Power	 int `json:"power"`
}

//这里有个User2是因为前端获取数据的时候，我只想给它部分数据，像password,salt不能给
//即使是空的，字段名也不能暴露
type Admin2 struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Power	 int `json:"power"`
}


//Add Admin
func AddUser(u *Admin) error {
	sqlStr := "insert into admin(username,password,salt,nickname,email) values(?,?,?,?,?)"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.UserName, u.PassWord,u.Salt,u.NickName,u.Email)
	if err != nil {
		return err
	}
	return nil
}

//Query Admin from admin id
func QueryOneUser(id int) (*Admin2, error) {
	sqlStr := "select id,username,nickname,email,power from admin where id=?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	var admin Admin2
	err = stmt.QueryRow(id).Scan(&admin.Id,&admin.UserName, &admin.NickName, &admin.Email,&admin.Power)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

//Query All Users Information
func QueryAllUser() (user_list []*Admin2,total int,err error) {
	sqlStr := "select id,username,nickname,email,power from admin;"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil,0,err
	}

	for rows.Next() {
		var admin Admin2
		err = rows.Scan(&admin.Id, &admin.UserName, &admin.NickName, &admin.Email,&admin.Power)
		if err != nil {
			return nil,0,err
		}
		user_list = append(user_list, &admin)
	}

	//
	sqlStr = "select count(*) from admin;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	return user_list, total, err
}

//Delete Admin From ID
func DeleteUser(id_list []int) error {
	sqlStr := "delete from admin where id = ?;"
	sqlStr2 := "select count(*) from admin where id = ?;"
	stmt2, err := db.Db.Prepare(sqlStr2)
	if err != nil {
		return err
	}
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	for _,id := range id_list{
		var count int
		err = stmt2.QueryRow(id).Scan(&count)
		if err != nil {
			return err
		}
		if count == 0 {
			return errors.New("Admin ID Exception")
		} else {

			_, err = stmt.Exec(id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


//Update Admin
func UpdateUser(u *Admin) error {
	sqlStr := "update admin set nickname = ?,email = ? where id = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}

	_,err = stmt.Exec(u.NickName,u.Email,u.Id)
	if err != nil{
		return err
	}

	return nil
}

//更新密码
func UpdateAdminPasswd(id int,old string,new string) error {
	//1.先验证旧密码是否正确
	sqlStr := "select count(*) from admin where id=? and password = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	var count int
	err = stmt.QueryRow(id,old).Scan(&count)
	if err != nil{
		return err
	}
	if count == 0 {
		return errors.New("密码错误")
	}

	//2.再更新密码
	sqlStr2 := "update admin set password = ? where id = ?;"
	stmt2,err := db.Db.Prepare(sqlStr2)
	if err != nil{
		return err
	}
	_,err = stmt2.Exec(new,id)
	if err != nil{
		return err
	}
	return nil
}

//
func Login(username string,passwd string) (*Admin,error) {
	sqlStr := "select id,username,power from admin where username = ? and password = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	var admin Admin
	err = stmt.QueryRow(username,passwd).Scan(&admin.Id,&admin.UserName,&admin.Power)
	if err != nil{
		return nil,err
	}
	return &admin,nil
}


//Get Admin Salt From UserID
func GetSalt(id int) (salt string,err error) {
	sqlStr := "select salt from admin where id = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return salt,err
	}

	err = stmt.QueryRow(id).Scan(&salt)
	if err != nil{
		return salt,err
	}
	return salt,nil
}

//Get UserID from UserName
func GetUserID(username string) (user_id int,err error) {
	sqlStr := "select id from admin where username = ?"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return user_id,err
	}
	err = stmt.QueryRow(username).Scan(&user_id)
	if err != nil{
		return user_id,err
	}
	return user_id,nil
}

