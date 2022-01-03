package model

import (
	"dachuang/db"
	"errors"
)

type Answer struct{
	AID int `json:"aid"`
	QID int `json:"qid"`
	UID int `json:"uid"`
	Answ string `json:"answer"`
	Time string `json:"time"`
}

type Answer2 struct{
	AID int `json:"aid"`
	QID int `json:"qid"`
	UID int `json:"uid"`
	Answ string `json:"answer"`
	Time string `json:"time"`
	NickName string `json:"nickname"`
	Photo string `json:"photo"`
}

//用户添加回答
func AddAnswer(ans *Answer) error {
	sqlStr := "insert into community_answer(com_qid,com_uid,com_answer,com_ans_time) values(?,?,?,now());"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	_,err = stmt.Exec(ans.QID,ans.UID,ans.Answ)
	if err != nil{
		return err
	}
	return nil
}

//管理员删除回答，只有管理员才能批量删
func AdminDeleteAnswer(id_list []int) error {
	sqlStr := "delete from community_answer where com_qid = ? and com_aid = ?;"
	sqlStr2 := "select count(*) from community_answer where com_qid = ? and com_aid = ?;"
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
			return errors.New("Question ID Exception")
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

//用户删除回答，一次只能删一个
func UserDeleteAnswer(aid int) error {
	sqlStr := "delete from community_answer where com_aid = ?"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_,err = stmt.Exec(aid)
	if err != nil{
		return err
	}
	return nil
}

//查看话题的回答
func QueryAnswer(qid string) (a_list []*Answer2,err error) {
	//回答信息
	sqlStr := "select * from community_answer where com_qid = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	rows,err := stmt.Query(qid)
	if err != nil{
		return nil,err
	}

	//获取回答用户的信息
	sqlStr2 := "select user_nickname,user_photo from community_user where id = ?;"
	stmt2,err := db.Db.Prepare(sqlStr2)
	if err != nil{
		return nil,err
	}

	for rows.Next(){
		var answer Answer2
		err = rows.Scan(&answer.AID,&answer.QID,&answer.UID,&answer.Answ,&answer.Time)
		if err != nil{
			return nil,err
		}

		err = stmt2.QueryRow(answer.UID).Scan(&answer.NickName,&answer.Photo)
		if err != nil{
			return nil,err
		}
		a_list = append(a_list, &answer)
	}
	return a_list,nil
}

//验证合法
func VerifyAnswer(uid int,qid int,aid int) error {
	sqlStr := "select count(*) from community_answer where com_qid = ? and com_uid = ? and com_aid = ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return err
	}
	var count int
	err = stmt.QueryRow(qid,uid,aid).Scan(&count)
	if err != nil{
		return err
	}
	if count == 0{
		return errors.New("网络异常")
	}
	return nil
}