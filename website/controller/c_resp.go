package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

//simplify error && success response
func ErrorResp(c *gin.Context,errMsg string) {
	Resp(c,200,errMsg,nil)
}

func SuccessResp(c *gin.Context,data interface{}) {
	Resp(c,200,"success",data)
}


//Article Response
func Resp(c *gin.Context,status int,msg string,data interface{}){
	c.JSON(200,gin.H{
		"status":status,
		"message":msg,
		"data":data,
		"timestamp":time.Now().UnixMilli(),
	})
}