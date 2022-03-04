package middlewares

import (
	"github.com/gin-gonic/gin"
)

var ValidMethod = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
}

func CorsAuth(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token,Content-Language,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}

		// _,ok := ValidMethod[method]
		// if !ok{
		// 	c.JSON(200,gin.H{
		// 		"status":false,
		// 		"msg":"无权访问",
		// 	})
		// 	c.Abort()
		// }
		c.Next()
	}

}
