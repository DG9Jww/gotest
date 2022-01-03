package controller

import (
	"dachuang/model"
	"github.com/gin-gonic/gin"
)

//用户查看所有话题
func ShowAllQuestions(c *gin.Context) {
	var total int
	var qlist []*model.Question2
	qlist,total,err := model.QueryAllQues()
	if err != nil{
		ErrorResp(c,"error")
		return
	}

	SuccessResp(c,map[string]interface{}{
		"total":total,
		"questions":qlist,
	})
}

//用户查看自己发布的所有话题
func ShowUserQues(c *gin.Context) {
	uid := c.GetInt("userID")
	var qlist []*model.Question
	qlist,err := model.QueryUserAllQues(uid)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	SuccessResp(c,map[string]interface{}{
		"questions":qlist,
	})
}

//用户添加话题
func AddQuestion(c *gin.Context){
	//前端只用传问题即可
	var ques model.Question
	err := c.ShouldBind(&ques)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	uid := c.GetInt("userID")
	ques.UID = uid

	err = model.AddQuestion(&ques)
	if err != nil{
		ErrorResp(c,"添加失败")
		return
	}
	SuccessResp(c,"添加成功")
}

//用户删除话题
func UserDeleteQues(c *gin.Context){
	data := make(map[string]int)
	err := c.BindJSON(&data)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	qid := data["qid"]
	err = model.UserDeleteQues(qid)
	if err != nil{
		ErrorResp(c,"删除失败")
		return
	}
	SuccessResp(c,"删除成功")
}

//管理员删除话题
func AdminDeleteQues(c *gin.Context){
	data := make(map[string][]int)
	err := c.BindJSON(&data)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	qid_list := data["qid"]
	err = model.AdminDeleteQues(qid_list)
	if err != nil{
		ErrorResp(c,"删除失败")
		return
	}
	SuccessResp(c,"删除成功")
}

//搜索
func SearchQuestion(c *gin.Context) {
	query := c.Query("query")
	var a_list []*model.Question
	var total int
	a_list,total,err := model.SearchQuestion(query)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}

	SuccessResp(c,map[string]interface{}{
		"total":total,
		"questions":a_list,
	})
}
