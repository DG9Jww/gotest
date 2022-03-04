package model

import (
	"website/db"
)

type Admin struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Right    int    `json:"right"`	//0是普通，1为超级
}

type Admin2 struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

//添加用户
func (ad *Admin) Insert() error {
	sqlStr := "insert into admins(ad_name,ad_password,ad_right) values(?,HEX(AES_ENCRYPT(?,ad_name)),0);"
	err := db.Exec(sqlStr, ad.UserName, ad.PassWord)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//验证密码
func (ad *Admin) QueryPasswd() error {
	name := ad.UserName
	passwd := ad.PassWord
	sqlStr := "select * from admins where ad_name=? and AES_DECRYPT(UNHEX(ad_password),ad_name)=?;"
	row, err := db.QueryRow(sqlStr, name, passwd)
	if err != nil {
		return err
	} else {
		err := row.Scan(&ad.ID,&ad.UserName,&ad.PassWord,&ad.Right)
		if err != nil {
			return err
		}
		return nil
	}
}

//获取管理员ID
func (ad *Admin) QueryAdminID() (int, error) {
	name := ad.UserName
	var admin_id int
	sqlStr := "select id from admins where ad_name = ?;"
	row, err := db.QueryRow(sqlStr, name)
	if err != nil {
		return 0, err
	}
	err = row.Scan(&admin_id)
	if err != nil {
		return 0, err
	}
	return admin_id, err
}

//查询单个用户信息
func (ad *Admin) QueryAdmin() error {
	sqlStr := "select id,ad_name from admins where id = ?;"
	row, err := db.QueryRow(sqlStr, ad.ID)
	if err != nil {
		return err
	}
	err = row.Scan(&ad.ID,&ad.UserName)
	if err != nil {
		return err
	}
	return err
}

//查询所有管理员信息
func (*Admin) QueryAdmins(page int, page_size int) (total int, a_list []*Admin2, err error) {

	//查询出用户信息
	sqlStr := "select id,ad_name from admins limit ?,?;"
	rows, err := db.Query(sqlStr, (page-1)*page_size, page_size) //第一页从0开始，所以减一
	if err != nil {
		return 0, nil, err
	}

	for rows.Next() {
		var admin Admin2
		err = rows.Scan(&admin.ID, &admin.UserName)
		if err != nil {
			return 0, nil, err
		}
		a_list = append(a_list, &admin)
	}

	//统计总数
	sqlStr = "select count(*) from admins;"
	row, err := db.QueryRow(sqlStr)
	if err != nil {
		return 0, nil, err
	}
	err = row.Scan(&total)
	if err != nil {
		return 0, nil, err
	}

	return total, a_list, nil
}

//删除用户
func (ad *Admin) Delete() error {
	sqlStr := "delete from admins where id = ?;"
	err := db.Exec(sqlStr, ad.ID)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//更改密码
func (ad *Admin) UpdatePasswd(new string) error {
	sqlStr := "update admins set ad_password = HEX(AES_ENCRYPT(?,ad_name)) where ad_name =?;"
	err := db.Exec(sqlStr, new, ad.UserName)
	if err != nil {
		return err
	} else {
		return nil
	}
}
