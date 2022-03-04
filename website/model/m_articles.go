package model

import (
	"time"
	"website/db"
)

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author    	string `json:"author"`
	PublishTime int64  `json:"publishTime"`
	CategoryID  int    `json:"categoryId"`	//1新闻，2通知，3学院简介，4合作交流
	Cover       string `json:"cover"`	//封面
	Status      int    `json:"status"` //1为发布了，0为草稿
}



//添加文章
func (a *Article) Insert() error {
	sqlStr := "insert into articles(a_title,a_content,a_author,a_publishTime,a_categoryID,a_cover,a_status) values(?,?,?,?,?,?,?)"
	err := db.Exec(sqlStr,a.Title,a.Content,a.Author,time.Now().UnixMilli(),a.CategoryID,a.Cover,a.Status)
	if err != nil{
		return err
	}else{
		return nil;
	}
}

//删除文章
func (a *Article) Delete() error {
	sqlStr := "delete from articles where id = ?"
	err := db.Exec(sqlStr,a.ID)
	if err != nil{
		return err
	}else{
		return nil;
	}
}

//更新文章
func (a *Article) Update() error {
	sqlStr := "update articles set a_title = ?,a_content = ?,a_author = ?,a_publishTime = ?,a_categoryID = ?,a_cover = ?,a_status = ? where id = ?"
	err := db.Exec(sqlStr,a.Title,a.Content,a.Author,time.Now().UnixMilli(),a.CategoryID,a.Cover,a.Status,a.ID)
	if err != nil{
		return err
	}else{
		return nil;
	}
}

//根据分类查询文章
func (a *Article)QueryCategory(page int,page_size int,cagyID int) (total int,a_list []*Article,err error) {

	//先查出当前分类的所有文章
	sqlStr := "select * from articles where a_categoryID = ? limit ?,?"
	rows,err := db.Query(sqlStr,cagyID,page,page_size)
	if err != nil{
		return 0,nil,err
	}
	for rows.Next(){
		var article Article
		err = rows.Scan(&article.ID,&article.Title,&article.Content,&article.Author,&article.PublishTime,&article.CategoryID,&article.Cover,&article.Status)
		if err != nil{
			return 0,nil,err
		}
		a_list = append(a_list,&article)
	}

	//统计总数
	sqlStr = "select count(*) from articles where a_categoryID = ?;"
	row,err := db.QueryRow(sqlStr,cagyID)
	if err != nil{
		return 0,nil,err
	}
	err = row.Scan(&total)
	if err != nil {
		return 0,nil,err
	}

	return total,a_list,nil
}

//查询所有文章
func (*Article) QueryArticles(page int,page_size int) (total int, a_list []*Article,err error){
	//查询出所有文章详细
	sqlStr := "select * from articles limit ?,?;"
	rows,err := db.Query(sqlStr,(page-1)*page_size, page_size)	
	if err != nil{
		return 0, nil, err
	}

	for rows.Next(){
		var article Article
		err = rows.Scan(&article.ID,&article.Title,&article.Content,&article.Author,&article.PublishTime,&article.CategoryID,&article.Cover,&article.Status)
		if err != nil{
			return 0,nil,err
		}
		a_list = append(a_list,&article)
	}

	//统计总数
	sqlStr = "select count(*) from articles;"
	row,err := db.QueryRow(sqlStr)
	if err != nil{
		return 0,nil,err
	}
	err = row.Scan(&total)
	if err != nil {
		return 0,nil,err
	}

	return total,a_list,nil
}

//查询文章详情
func (a *Article)QueryDetails(id int) (error) {
	sqlStr := "select * from articles where id=?;"
	row,err := db.QueryRow(sqlStr,id)
	if err != nil {
		return err
	}
	err = row.Scan(&a.ID,&a.Title,&a.Content,&a.Author,&a.PublishTime,&a.CategoryID,&a.Cover,&a.Status)
	if err != nil {
		return err
	}
	return nil
}



func (a *Article) Test() error {
	sqlStr := "select * from articles where id="
	err := db.Exec(sqlStr,a.ID)
	if err != nil{
		return err
	}else{
		return nil;
	}
}