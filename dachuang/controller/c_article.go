package controller

import (
	"dachuang/model"
	"github.com/gin-gonic/gin"
)

//Get one Article information
func GetOneArtile(c *gin.Context) {
	id := c.Query("id")
	var article *model.Article
	article, err := model.QueryOneArticle(id)
	if err != nil {
		ErrorResp(c, "error")
		return
	}
	SuccessResp(c, article)
}

//Add Article
func DoAddArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil {
		ErrorResp(c, "error:"+err.Error())
		return
	}
	err = model.AddArticle(&article)
	if err != nil {
		ErrorResp(c, "error:"+err.Error())
		return
	}
	SuccessResp(c, "添加成功")
}

//Delete Article
func DeleteArticle(c *gin.Context) {
	data := make(map[string][]int)
	c.BindJSON(&data)
	id_list := data["id"]
	err := model.DeleteArticle(id_list)
	if err != nil {
		ErrorResp(c, "error")
		return
	}

	SuccessResp(c, "删除成功!")
}

//Update Article
func UpdateArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil {
		ErrorResp(c, "bind error")
		return
	}

	err = model.UpdateArticle(&article)
	if err != nil {
		ErrorResp(c, "update error")
		return
	}

	SuccessResp(c, "更新成功")
}

//All Articles
func ShowAllArticles(c *gin.Context) {
	var a_list []*model.Article
	var total int
	a_list, total, err := model.QueryAllArticle()
	if err != nil {
		ErrorResp(c, "error")
		return
	}

	SuccessResp(c, map[string]interface{}{
		"total":    total,
		"articles": a_list,
	})
}

//Get News
func GetNews(c *gin.Context) {
	var a_list []*model.Article
	var total int
	a_list, total, err := model.GetNews()
	if err != nil {
		ErrorResp(c, "error")
		return
	}

	SuccessResp(c, map[string]interface{}{
		"total":    total,
		"articles": a_list,
	})
}

//Get Notices
func GetNotices(c *gin.Context) {
	var a_list []*model.Article
	var total int
	a_list, total, err := model.GetNotices()
	if err != nil {
		ErrorResp(c, "error")
		return
	}

	SuccessResp(c, map[string]interface{}{
		"total":    total,
		"articles": a_list,
	})
}

//Get Introduction
func GetIntroduction(c *gin.Context) {
	var a_list []*model.Article
	var total int
	a_list, total, err := model.GetIntroduction()
	if err != nil {
		ErrorResp(c, "error")
		return
	}

	SuccessResp(c, map[string]interface{}{
		"total":    total,
		"articles": a_list,
	})
}

//Search
func SearchArticle(c *gin.Context) {
	query := c.Query("query")
	var a_list []*model.Article
	var total int
	a_list, total, err := model.SearchArticle(query)
	if err != nil {
		ErrorResp(c, "网络异常")
		return
	}

	SuccessResp(c, map[string]interface{}{
		"total":    total,
		"articles": a_list,
	})
}
