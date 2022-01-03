package controller

import (
	"dachuang/common"
	"dachuang/middlewares"
	"dachuang/model"
	"github.com/gin-gonic/gin"
)

const (
	SuffixTips = "图片后缀必须为jpg,jpeg,png"
)

//登录
func CommUserLogin(c *gin.Context){
	//1.获取明文账号密码，并转为加密的
	var param = make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil{
		ErrorResp(c,"绑定失败")
		return
	}
	username := param["username"]
	password := param["password"]

	id,err := model.GetCommUserID(username)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}

	salt,err := model.GetCommUserSalt(id)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}

	encrypt_passwd := common.MD5encrypt(password+salt)
	//2.验证密码
	err = model.VerifyCommUserPasswd(encrypt_passwd,id)
	if err != nil{
		ErrorResp(c,"账号或密码错误")
		return
	}

	token,err := middlewares.CreateToken(id,3)
	if err != nil{
		ErrorResp(c,"创建Token失败")
		return
	}
	SuccessResp(c,gin.H{"token":token})
}

//显示个人资料
func ShowCommUserInfo(c *gin.Context){
	val,ok := c.Get("userID")
	if !ok{
		ErrorResp(c,"网络异常")
		return
	}

	id := val.(int)
	info,err := model.QueryOneCommUser(id)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	SuccessResp(c,info)
}

//修改密码
func CommUserChangePasswd(c *gin.Context){
	//1.获取明文旧密码和新密码
	val,ok:= c.Get("userID")
	if !ok{
		ErrorResp(c,"网络异常，修改失败")
		return
	}
	id := val.(int)
	var param = make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil{
		ErrorResp(c,"参数绑定失败")
		return
	}

	old := param["old"]
	new := param["new"]
	//2.得到加密密码
	salt,err := model.GetCommUserSalt(id)
	if err != nil{
		ErrorResp(c,"修改密码失败")
		return
	}
	
	old_passwd := common.MD5encrypt(old+salt)

	new_passwd := common.MD5encrypt(new+salt)
	//3.先验证旧密码
	err = model.VerifyCommUserPasswd(old_passwd,id)
	if err != nil{
		ErrorResp(c,"密码错误")
		return
	}
	//4.更新密码,这里可以加规则
	err = model.UpdateCommUserPasswd(id,new_passwd)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	SuccessResp(c,"成功修改密码")
}

//修改资料（仅仅昵称和头像可更改）
func UpdateCommUserInfo(c *gin.Context){
	//1.绑定参数
	var comm_user model.CommUser
	err := c.ShouldBind(&comm_user)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	id := c.GetInt("userID")
	comm_user.ID = id
	//2.头像处理
	var file_path string
	var url_path string
	f,err := c.FormFile("img")
	if f != nil{
		if err != nil{
			ErrorResp(c,"网络异常")
			return
		}else{
			file_path,url_path,err = common.GetFilePath(f)
			if err != nil{
				ErrorResp(c,SuffixTips)
				return
			}
			err = c.SaveUploadedFile(f,file_path)
			if err != nil{
				ErrorResp(c,"图片上传失败")
				return
			}
		}
		comm_user.Photo = url_path
	}

	err = model.UpdateCommUser(&comm_user)
	if err != nil{
		ErrorResp(c,"更新失败")
		return
	}
	SuccessResp(c,"更新成功")
}

//删除用户,这里要传数组
func DeleteCommUser(c *gin.Context){
	var param = make(map[string][]int)
	err := c.BindJSON(&param)
	if err != nil{
		ErrorResp(c,"网络异常")
	}
	id_list := param["id"]
	err = model.DeleteCommUser(id_list)
	if err != nil{
		ErrorResp(c,"网络异常")
		return
	}
	SuccessResp(c,"删除成功!")
}

//添加社区用户
func AddCommUser(c *gin.Context){
	comm_user := new(model.CommUser)
	c.ShouldBind(&comm_user)

	//验证用户名是否存在
	_,err := model.GetCommUserID(comm_user.UserName)
	if err == nil{
		ErrorResp(c,"用户名已存在")
		return
	}

	//信息不能为空
	if comm_user.NickName != "" && comm_user.Password != "" && comm_user.UserName != ""{
		//盐以及密码加密
		comm_user.Salt = common.RandomNum()
		comm_user.Password = common.MD5encrypt(comm_user.Password+comm_user.Salt)
		//头像处理。刚创建用户时为默认
		comm_user.Photo = "/static/photo/default.jpg"

		err = model.AddCommUser(comm_user)
		if err != nil{
			ErrorResp(c,"添加失败")
			return
		}
	}else{
		ErrorResp(c,"信息不能为空")
		return
	}
	SuccessResp(c,"添加成功")
}

//显示所有用户资料
func ShowAllCommUsers(c *gin.Context){
	var u_list []*model.CommUser2
	var total int
	u_list,total,err := model.QueryAllCommUser()
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	SuccessResp(c,map[string]interface{}{
		"total":total,
		"users":u_list,
	})
}

