package model

import (
	"dachuang/db"
	"errors"
)

type Question struct {
	QID  int    `json:"qid"`
	Ques string `json:"question"`
	UID  int    `json:"uid"`
	Time string `json:"time"`
}

//这个是用来返回信息的
type Question2 struct {
	QID      int    `json:"qid"`
	Ques     string `json:"question"`
	UID      int    `json:"uid"`
	NickName string `json:"nickname"`
	Photo    string `json:"photo"`
	Time     string `json:"time"`
}

//ADD Question
func AddQuestion(comm *Question) error {
	sqlStr := "insert into community_question(com_question,com_uid,com_time) values(?,?,now());"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comm.Ques, comm.UID)
	if err != nil {
		return err
	}
	return nil
}

//Admin Delete Question
func AdminDeleteQues(id_list []int) error {
	sqlStr := "delete from community_question where com_qid = ?;"
	sqlStr2 := "select count(*) from community_question where com_qid = ?;"
	stmt2, err := db.Db.Prepare(sqlStr2)
	if err != nil {
		return err
	}
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	//删除每一个之前都要检查ID是否合法
	for _, id := range id_list {
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
	for _, id := range id_list {
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}
	return nil
}

//用户删除话题，一次只能删一个
func UserDeleteQues(qid int) error {
	sqlStr := "delete from community_question where com_qid = ?;"
	sqlStr2 := "select count(*) from community_question where com_qid = ?;"
	stmt2, err := db.Db.Prepare(sqlStr2)
	if err != nil {
		return err
	}
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	//删除每一个之前都要检查ID是否合法
	var count int
	err = stmt2.QueryRow(qid).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Question ID Exception")
	}

	//合法则执行删除
	_, err = stmt.Exec(qid)
	if err != nil {
		return err
	}
	return nil
}

//Show All Question
func QueryAllQues() (qlist []*Question2, total int, err error) {
	sqlStr := "select * from community_question;"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, 0, err
	}

	sqlStr = "select user_nickname,user_photo from community_user where id =?"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var ques Question2
		err = rows.Scan(&ques.QID, &ques.Ques, &ques.Time, &ques.UID)
		if err != nil {
			return nil, 0, err
		}
		err = stmt.QueryRow(ques.UID).Scan(&ques.NickName, &ques.Photo)
		if err != nil {
			return nil, 0, err
		}
		qlist = append(qlist, &ques)
	}

	//获取总数量
	sqlStr = "select count(*) from community_question;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	if err != nil {
		return nil, 0, err
	}
	return qlist, total, err
}

//query questions from user id
func QueryUserAllQues(uid int) (qlist []*Question, err error) {
	sqlStr := "select * from community_question where com_uid = ?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ques Question
		err = rows.Scan(&ques.QID, &ques.Ques, &ques.Time, &ques.UID)
		if err != nil {
			return nil, err
		}
		qlist = append(qlist, &ques)
	}

	return qlist, err
}


//Search Question
func SearchQuestion(query string) (qlist []*Question, total int, err error) {
	sqlStr := "select * from community_question where com_question like ?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return qlist, total, err
	}
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return qlist, total, err
	}
	for rows.Next() {
		question := new(Question)
		err = rows.Scan(&question.QID, &question.Ques, &question.Time, &question.UID)
		if err != nil {
			return nil, 0, err
		} else {
			qlist = append(qlist, question)
		}
	}

	//2.Get Total Article Amount
	sqlStr = "select count(*) from community_question where com_question like ?;"
	stmt2, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, 0, err
	}
	err = stmt2.QueryRow("%" + query + "%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}
	return qlist, total, nil
}
