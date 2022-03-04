package route

import (
	"website/controller"
	"website/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Static("/static","view/static")
	r.Use(gin.Logger(), gin.Recovery(), middlewares.CorsAuth, middlewares.LogerMiddleware())

	//无需登录
	{
		r.GET("/v1/article", controller.GetArticleDetails)      //获取单篇文章详细
		r.GET("/v1/articles", controller.GetAllArticles)        //获取所有文章
		r.GET("/v1/articleType", controller.GetArticleCategory) //根据分类ID查询不同类型文章
		r.POST("/v1/adminlogin", controller.AdminLogin)         //管理员登录
	}

	//需要登录后
	tokenGroup := r.Group("/v1/auth", middlewares.AuthenticateToken)
	{
		tokenGroup.POST("/article", controller.AddArticle)      //添加文章
		tokenGroup.PUT("/article", controller.UpdateArticle)    //更新文章
		tokenGroup.DELETE("/article", controller.DeleteArticle) //删除文章
		tokenGroup.PUT("/admin", controller.ChAdminPasswd)      //更改管理员密码
		tokenGroup.GET("/admin", controller.GetAdmin)           //获取单个用户信息
		tokenGroup.POST("/upload", controller.Upload)           //上传图片	

		//需要超级管理员
		superGroup := tokenGroup.Group("/admin", middlewares.AuthAdminRight)
		{
			superGroup.POST("/admin", controller.AddAdmin)      //添加管理员
			superGroup.DELETE("/admin", controller.DeleteAdmin) //删除管理员
			superGroup.GET("/admins", controller.GetAdmins)     //获取所有用户信息
		}
	}

	//测试
	{
		r.GET("v1/test", func(c *gin.Context) {
			c.File(`D:\Documents\WeChat Files\wxid_t08at3quahrs21\FileStorage\Video\2021-04\04f1adc645a061a4080a92bd23496391.mp4`)
		})
	}

	return r
}
