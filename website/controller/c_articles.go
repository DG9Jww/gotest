package controller

import (
	"fmt"
	"strconv"
	"website/model"
	"github.com/gin-gonic/gin"
)

//添加文章
func AddArticle(c *gin.Context){
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil{
		ErrorResp(c,"绑定失败")
		return
	}

	err = article.Insert()
	if err != nil{
		ErrorResp(c,"文章添加失败")
		return
	}
	SuccessResp(c,"添加成功")
}

//更新文章
func UpdateArticle(c *gin.Context){
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil{
		ErrorResp(c,"绑定失败")
		return
	}

	err = article.Update()
	if err != nil{
		ErrorResp(c,"文章更新失败")
		return
	}
	SuccessResp(c,"更新成功")
}

//删除文章
func DeleteArticle(c *gin.Context){
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil{
		ErrorResp(c,"绑定失败")
		return
	}

	err = article.Delete()
	if err != nil{
		ErrorResp(c,"文章删除失败")
		return
	}
	SuccessResp(c,"删除成功")
}

//查询文章分类
func GetArticleCategory(c *gin.Context){
	//获取参数
	page := c.Query("page")
	page_size := c.Query("pageSize")
	cagyID := c.Query("categoryID")
	p,err := strconv.Atoi(page)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}
	p_s,err := strconv.Atoi(page_size)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}
	cid,err := strconv.Atoi(cagyID)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}

	a := new(model.Article)
	//开始查询
	total,a_list,err := a.QueryCategory(p,p_s,cid)
	if err != nil{
		ErrorResp(c,"文章查询失败")
		return
	}
	fmt.Println(total)
	SuccessResp(c, map[string]interface{}{
		"page":p,
		"pageSize":p_s,
		"total":total,
		"list": a_list,
	})
}

//查询文章详细
func GetArticleDetails(c *gin.Context){
	//获取参数
	param := c.Query("id")
	id,err := strconv.Atoi(param)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}

	//查询
	a := new(model.Article)
	err = a.QueryDetails(id)
	if err != nil{
		ErrorResp(c,"查询失败")
		return
	}
	SuccessResp(c,a)
}

//查询所有文章
func GetAllArticles(c *gin.Context){
	//获取参数
	page := c.Query("page")
	page_size := c.Query("pageSize")
	p,err := strconv.Atoi(page)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}
	p_s,err := strconv.Atoi(page_size)
	if err != nil{
		ErrorResp(c,"参数错误")
		return
	}

	//查询
	a := new(model.Article)
	total,list,err := a.QueryArticles(p,p_s)
	if err != nil{
		ErrorResp(c,"查询失败")
		return
	}
	SuccessResp(c, map[string]interface{}{
		"page":p,
		"pageSize":p_s,
		"total":total,
		"list": list,
	})
}