package controller

import (
	"dachuang/common"
	"dachuang/model"
	"dachuang/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	UserExist	=	20
)

//Get one User information
func GetOneUser(c *gin.Context) {
	id,ok:= c.Get("userID")
	if !ok {
		ErrorResp(c,"获取信息失败")
		return
	}
	var user *model.Admin2
	user,err := model.QueryOneUser(id.(int))
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	SuccessResp(c,user)
}

//Add User 
func DoAddUser(c *gin.Context) {
	//请求只需要username,password,nickname,email。盐和ID 后端处理
	var user model.Admin
	err := c.ShouldBind(&user)
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	//先验证是用户名否存在
	_,err = model.GetUserID(user.UserName)
	if err == nil{
		ErrorResp(c,"用户已存在")
		return
	}
	//再验证信息是否合法
	if user.Email != "" && user.NickName != "" && user.PassWord != "" && user.UserName != ""{
		//再进行添加操作
		user.Salt = common.RandomNum()
		user.PassWord = common.MD5encrypt(user.PassWord+user.Salt)
		err = model.AddUser(&user)
		if err != nil{
			ErrorResp(c,"error")
			return
		}
	}else{
		ErrorResp(c,"信息不能为空")
		return
	}

	SuccessResp(c,"添加成功")
}

//Delete User
func DeleteUser(c *gin.Context){
	data := make(map[string][]int)
	c.BindJSON(&data)
	id_list := data["id"]
	err := model.DeleteUser(id_list)
	if err != nil{
		ErrorResp(c,"error")
		return
	}

	SuccessResp(c,"删除成功!")
}


//Update User
func UpdateUser(c *gin.Context){
	var user model.Admin
	err := c.ShouldBind(&user)
	val,ok := c.Get("userID")
	if !ok {
		ErrorResp(c,"更新失败")
		return
	}
	user.Id = val.(int)
	if err != nil{
		ErrorResp(c,"bind parameter error")
		return
	}
	user.Salt = common.RandomNum()
	err = model.UpdateUser(&user)
	if err != nil{
		ErrorResp(c,"update error")
		return
	}
	SuccessResp(c,"更新成功")
}


//All Users
func ShowAllUsers(c *gin.Context){
	var u_list []*model.Admin2
	var total int
	u_list,total,err := model.QueryAllUser()
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	SuccessResp(c,map[string]interface{}{
		"total":total,
		"users":u_list,
	})
}


//Login
func Login(c *gin.Context){
	//1.验证账号密码是否正确,正确则添加JWT
	var req_data = make(map[string]string)
	err := c.BindJSON(&req_data)
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	admin_id,err := model.GetUserID(req_data["username"])
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	salt,err := model.GetSalt(admin_id)
	if err != nil{
		ErrorResp(c,"error")
		return
	}
	
	passwd :=common.MD5encrypt(req_data["password"]+salt) 
	user,err := model.Login(req_data["username"],passwd)
	if err != nil{
		ErrorResp(c,"账号或者密码错误")
		return
	}

	//创建JWT并返回
	token,err := middlewares.CreateToken(user.Id,user.Power)
	if err != nil{
		ErrorResp(c,"创建Token失败")
		return
	}

	SuccessResp(c,gin.H{"token":token})
}


//Change Password
func AdminChangePassword(c *gin.Context){
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
	//先根据明文密码得到加密密码
	salt,err := model.GetSalt(id)
	if err != nil{
		ErrorResp(c,"修改密码失败")
		return
	}
	old_passwd := common.MD5encrypt(old+salt)
	new_passwd := common.MD5encrypt(new+salt)
	
	err = model.UpdateAdminPasswd(id,old_passwd,new_passwd)
	if err != nil{
		ErrorResp(c,"密码修改失败")
		return
	}
	SuccessResp(c,"成功修改密码")
}