package middlewares

import (
	"github.com/gin-gonic/gin"
)

//只有超级管理员才能通过验证
func AuthAdminRight(c *gin.Context){
	val,ok := c.Get("userRight")
	if !ok {
		c.JSON(200,gin.H{
			"status":false,
			"msg":"无权访问",
		})
		c.Abort()
	}

	right :=val.(int)
	if right==1{
		c.Next()
	}else{
		c.JSON(200,gin.H{
			"status":false,
			"msg":"无权访问",
		})
		c.Abort()
	}
}

//只有社区普通用户才能进行操作
func AuthCommUserRight(c *gin.Context){
	right:= c.GetInt("userRight")
	if right != 3{
		c.JSON(200,gin.H{
			"status":false,
			"msg":"无权访问",
		})
		c.Abort()
	}else{
		c.Next()
	}
}


//只有超级管理员才能通过验证
func AuthNormalAdminRight(c *gin.Context){
	val,ok := c.Get("userRight")
	if !ok {
		c.JSON(200,gin.H{
			"status":false,
			"msg":"无权访问",
		})
		c.Abort()
	}

	right :=val.(int)
	if right==0 || right==1 {
		c.Next()
	}else{
		c.JSON(200,gin.H{
			"status":false,
			"msg":"无权访问",
		})
		c.Abort()
	}
}