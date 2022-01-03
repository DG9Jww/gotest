package model

import (
	"dachuang/db"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  string `json:"author"`
	Type    int `json:"type"`
}

//Add Article
func AddArticle(a *Article) error {
	sqlStr := "insert into article(title,content,author,time,a_type) values(?,?,?,now(),?);"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	stmt.Exec(a.Title, a.Content, a.Author, a.Type)
	return nil
}

//Delete Article
func DeleteArticle(id_list []int) error {
	sqlStr := "delete from article where id = ?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	for _, id := range id_list {
		//1.Firstly Judge Whether The Specific Field Exist
		sqlStr2 := "select count(*) from article where id = ?;"
		stmt2, err := db.Db.Prepare(sqlStr2)
		if err != nil {
			return err
		}
		var count int
		err = stmt2.QueryRow(id).Scan(&count)
		if err != nil {
			return err
		}
		if count == 0 {
			return nil
		} else {
			_, err = stmt.Exec(id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//Get All Article,Return Article Total Amount and All Article Struct
func QueryAllArticle() (a_list []*Article, total int, err error) {

	//1.Get All Article Struct
	sqlStr := "select * from article;"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		article := new(Article)
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author,  &article.Type)
		if err != nil {
			return nil, 0, err
		} else {
			a_list = append(a_list, article)
		}
	}

	//2.Get Total Article Amount
	sqlStr = "select count(*) from article;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	if err != nil{
		return nil, 0, err
	}
	
	return a_list, total, nil
}

//Get One Article Information From id
func QueryOneArticle(id string) (*Article, error) {
	sqlStr := "select * from article where id = ?;"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	var article Article
	err = stmt.QueryRow(id).Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author, &article.Type)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

//Update Article Information
func UpdateArticle(article *Article) error {
	sqlStr := "update article set title = ?,content = ?,time = now(),author = ? where id = ?"
	stmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(article.Title, article.Content, article.Author, article.Id)
	if err != nil {
		return err
	}
	return nil
}


//Get News
func GetNews() (a_list []*Article, total int, err error) {
		//1.Get All Article Struct
		sqlStr := "select * from article where a_type = 1;"
		rows, err := db.Db.Query(sqlStr)
		if err != nil {
			return nil, 0, err
		}
	
		for rows.Next() {
			article := new(Article)
			err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author,  &article.Type)
			if err != nil {
				return nil, 0, err
			} else {
				a_list = append(a_list, article)
			}
		}
	
		//2.Get Total Article Amount
		sqlStr = "select count(*) from article where a_type = 1;"
		err = db.Db.QueryRow(sqlStr).Scan(&total)
		if err != nil{
			return nil, 0, err
		}
		return a_list, total, nil
}

//Get Notice
func GetNotices() (a_list []*Article, total int, err error) {
	//1.Get All Article Struct
	sqlStr := "select * from article where a_type = 2;"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		article := new(Article)
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author,  &article.Type)
		if err != nil {
			return nil, 0, err
		} else {
			a_list = append(a_list, article)
		}
	}

	//2.Get Total Article Amount
	sqlStr = "select count(*) from article where a_type = 2;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	if err != nil{
		return nil, 0, err
	}
	return a_list, total, nil
}


//Get Introductio
func GetIntroduction() (a_list []*Article, total int, err error) {
	//1.Get All Article Struct
	sqlStr := "select * from article where a_type = 3;"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		article := new(Article)
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author,  &article.Type)
		if err != nil {
			return nil, 0, err
		} else {
			a_list = append(a_list, article)
		}
	}

	//2.Get Total Article Amount
	sqlStr = "select count(*) from article where a_type = 3;"
	err = db.Db.QueryRow(sqlStr).Scan(&total)
	if err != nil{
		return nil, 0, err
	}
	return a_list, total, nil
}

//Search
func SearchArticle(query string) (a_list []*Article, total int, err error) {
	sqlStr := "select * from article where content like ? or title like ?;"
	stmt,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return a_list,total,err
	}
	rows,err := stmt.Query("%"+query+"%","%"+query+"%")
	if err != nil{
		return a_list,total,err
	}
	for rows.Next() {
		article := new(Article)
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time, &article.Author,  &article.Type)
		if err != nil {
			return nil, 0, err
		} else {
			a_list = append(a_list, article)
		}
	}

	//2.Get Total Article Amount
	sqlStr = "select count(*) from article where content like ? or title like ?;"
	stmt2,err := db.Db.Prepare(sqlStr)
	if err != nil{
		return nil, 0, err
	}
	err = stmt2.QueryRow("%"+query+"%","%"+query+"%").Scan(&total)
	if err != nil{
		return nil, 0, err
	}
	return a_list, total, nil
}