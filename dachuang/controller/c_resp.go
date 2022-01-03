package controller

import (
	"github.com/gin-gonic/gin"
)
//simplify error && success response
func ErrorResp(c *gin.Context,data interface{}) {
	Resp(c,false,data)
}

func SuccessResp(c *gin.Context,data interface{}) {
	Resp(c,true,data)
}


//Response
func Resp(c *gin.Context,status bool,data interface{}){
	c.JSON(200,gin.H{
		"status":status,
		"msg":data,
	})
}