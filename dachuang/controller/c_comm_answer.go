package controller

import (
	"dachuang/model"
	"github.com/gin-gonic/gin"
)

//查看话题的回答
func ShowAnswers(c *gin.Context) {
	qid := c.Query("qid")
	alist,err := model.QueryAnswer(qid)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	SuccessResp(c,alist)
}

//用户增添回答
func AddAnswer(c *gin.Context) {
	var answ model.Answer
	err := c.ShouldBind(&answ)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	uid := c.GetInt("userID")
	answ.UID = uid

	err = model.AddAnswer(&answ)
	if err != nil{
		ErrorResp(c,"回答失败")
		return
	}
	SuccessResp(c,"回答成功")
}

//发布者删除回答,一次只能删一个
func DeleteAnswer(c *gin.Context) {
	//1.获取uid和aid
	uid := c.GetInt("userID")
	data := make(map[string]int)
	err := c.BindJSON(&data)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	qid := data["qid"]
	aid := data["aid"]

	//2.需要验证是否合法，即是不是用户创建的话题下面的回答
	err = model.VerifyAnswer(uid,qid,aid)
	if err != nil{
		ErrorResp(c,"非法请求")
		return
	}

	//3.执行删除
	err = model.UserDeleteAnswer(aid)
	if err != nil{
		ErrorResp(c,"删除失败")
	}
	SuccessResp(c,"删除成功")
}

//管理员删除回答，可批量删除
func AdminDeleteAnswer(c *gin.Context){
	data := make(map[string][]int)
	err := c.BindJSON(&data)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	aid_list := data["aid"]
	err = model.AdminDeleteAnswer(aid_list)
	if err != nil{
		ErrorResp(c,"删除失败")
		return
	}
	SuccessResp(c,"删除成功")
}