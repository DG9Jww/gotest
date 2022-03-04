package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"website/middlewares"
	"website/model"
	"github.com/gin-gonic/gin"
)

//添加用户
func AddAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		ErrorResp(c, "绑定失败")
		return
	}

	//验证用户名是否存在
	_, err = admin.QueryAdminID()
	if err != sql.ErrNoRows {
		ErrorResp(c, errors.New("用户名已存在").Error())
		return
	}

	//正式添加
	err = admin.Insert()
	if err != nil {
		ErrorResp(c, "管理员添加失败")
		return
	}
	SuccessResp(c, "添加成功")
}

//管理员登录
func AdminLogin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		ErrorResp(c, "绑定失败")
		return
	}
	//验证密码
	err = admin.QueryPasswd()
	if err != nil {
		ErrorResp(c, "11111"+err.Error())
		return
	}

	//创建JWT并返回
	fmt.Println(admin.ID)
	token, err := middlewares.CreateToken(admin.ID, admin.Right)
	if err != nil {
		ErrorResp(c, "创建Token失败")
		return
	}
	SuccessResp(c, gin.H{"token": token})
}

//得到单个用户信息
func GetAdmin(c *gin.Context) {
	var admin model.Admin
	var admin2 model.Admin2
	id := c.GetInt("userID")
	admin.ID = id

	err := admin.QueryAdmin()
	if err != nil {
		ErrorResp(c, "获取信息失败"+err.Error())
		return
	}

	admin2.ID = admin.ID
	admin2.UserName = admin.UserName
	SuccessResp(c, admin2)
}

//查询所有管理信息
func GetAdmins(c *gin.Context) {
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

	admin := new(model.Admin)
	//开始查询
	total,a_list, err := admin.QueryAdmins(p,p_s)
	if err != nil {
		ErrorResp(c, "查询失败"+err.Error())
		return
	}
	SuccessResp(c, map[string]interface{}{
		"page":p,
		"pageSize":p_s,
		"total":total,
		"list": a_list,
	})
}

//删除管理员用户
func DeleteAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		ErrorResp(c, "绑定失败")
		return
	}

	err = admin.Delete()
	if err != nil {
		ErrorResp(c, "删除失败")
		return
	}

	SuccessResp(c, "删除成功")
}

//更改密码
func ChAdminPasswd(c *gin.Context) {
	var admin model.Admin
	var param = make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		ErrorResp(c, "绑定失败2222"+err.Error())
		return
	}

	new := param["new"]
	admin.UserName = param["username"]
	admin.PassWord = param["password"]
	
	//先验证旧密码是否正确
	err = admin.QueryPasswd()
	if err != nil {
		ErrorResp(c, "原密码错误")
		return
	}

	//再进行更新
	err = admin.UpdatePasswd(new)
	if err != nil {
		ErrorResp(c, "更改失败")
		return
	}

	SuccessResp(c, "密码更改成功")
}
