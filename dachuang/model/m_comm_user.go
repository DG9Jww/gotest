package model

import (
	"dachuang/db"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type CommUser struct {
	ID       int		`json:"id"`
	UserName string		`json:"username"`
	Salt     string		
	Password string		`json:"password"`
	NickName string		`json:"nickname"`
	Photo	 string		`json:"photo"`
}

type CommUser2 struct{
	ID       int		`json:"id"`
	UserName string		`json:"username"`
	NickName string		`json:"nickname"`
	Photo	 string		`json:"photo"`
}

//查询出所有用户
func QueryAllCommUser() (user_list []*CommUser2,total int,err error) {
	sqlStr := "select id,user_name,user_nickname,user_photo from community_user"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil,0,err
	}

	for rows.Next(){
		var comm_user CommUser2
		err = rows.Scan(&comm_user.ID,&comm_user.UserName,&comm_user.NickName,&comm_user.Photo)
		if err != nil{
			return nil,0,err
		}
		user_list = append(user_list, &comm_user)
	}

	//获取总数量
	sqlStr = "select count(*) from community_user;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	if err != nil{
		return nil,0,err
	}
	return user_list,total,err
}


//添加用户
func AddCommUser(u *CommUser) error {
	sqlStr := "insert into community_user(user_name,user_password,user_salt,user_nickname,user_photo) values(?,?,?,?,?);"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	_,err = stmt.Exec(u.UserName,u.Password,u.Salt,u.NickName,u.Photo)
	if err != nil{
		return err
	}
	return nil
}

//删除用户
func DeleteCommUser(id_list []int) error {
	sqlStr := "delete from community_user where id = ?;"
	sqlStr2 := "select count(*) from community_user where id = ?;"
	stmt2, err := db.Db.Prepare(sqlStr2)
	if err != nil {
		return err
	}
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	//删除每一个之前都要检查ID是否合法
	for _,id := range id_list{
		var count int
		err = stmt2.QueryRow(id).Scan(&count)
		if err != nil {
			return err
		}
		if count == 0 {
			return errors.New("CommUser ID Exception")
		}
	}

	//合法则执行删除
	for _,id := range id_list{
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
		}
	return nil
}

//更新用户(昵称和头像)
func UpdateCommUser(u *CommUser) error {
	sqlStr := "update community_user set user_nickname = ?,user_photo = ? where id = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}

	_,err = stmt.Exec(u.NickName,u.Photo,u.ID)
	if err != nil{
		return err
	}
	return nil
}

//查询单个用户信息 from uid
func QueryOneCommUser(id int) (*CommUser2,error) {
	sqlStr := "select id,user_name,user_nickname,user_photo from community_user where id=?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	var comm_user CommUser2
	err = stmt.QueryRow(id).Scan(&comm_user.ID,&comm_user.UserName, &comm_user.NickName,&comm_user.Photo)
	if err != nil {
		return nil, err
	}
	return &comm_user, nil
}

//更新密码
func UpdateCommUserPasswd(id int,new string) error {
	sqlStr2 := "update community_user set user_password = ? where id = ?;"
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

//验证密码是否正确
func VerifyCommUserPasswd(p string,id int) error {
	sqlStr := "select count(*) from community_user where id = ? and user_password = ?"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	var count int
	err = stmt.QueryRow(id,p).Scan(&count)
	if err != nil{
		return errors.New("网络异常")
	}
	if count == 0{
		return errors.New("密码错误")
	}
	return nil
}

//GetID From UserName
func GetCommUserID(name string) (id int,err error) {
	sqlStr := "select id from community_user where user_name = ?"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return id,err
	}
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil{
		return id,err
	}
	return id,nil
}

//Get Salt
func GetCommUserSalt(id int) (salt string,err error) {
	sqlStr := "select user_salt from community_user where id = ?;"
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
