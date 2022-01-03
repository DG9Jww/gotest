package main

import (
	"dachuang/controller"
	"github.com/gin-gonic/gin"
	"dachuang/middlewares"
)

func main(){
	r := gin.Default()
	r.Static("/static","view/static")
	r.Use(middlewares.CorsAuth,middlewares.LogerMiddleware())
	tokenAuthGroup := r.Group("/v1/auth",middlewares.AuthenticateToken)		//是否带有合法token
	{	
		//nomal_admin
		user_group := tokenAuthGroup.Group("/user",middlewares.AuthNormalAdminRight)
			user_group.GET("/article",controller.GetOneArtile)		//Get One Article Information
			user_group.GET("/articles",controller.ShowAllArticles)	//Get All Articles Information
			user_group.POST("/article",controller.DoAddArticle)		//Add An Article
			user_group.PUT("/article",controller.UpdateArticle)		//Update Article
			user_group.DELETE("/article",controller.DeleteArticle)	//Delete Article

			user_group.PUT("/user",controller.UpdateUser)					//修改信息	
			user_group.GET("/user",controller.GetOneUser)					//查看个人信息
			user_group.PUT("/passwd",controller.AdminChangePassword) 		//修改密码

			user_group.GET("/comm_users",controller.ShowAllCommUsers)		//查看所有社区用户
			user_group.DELETE("/comm_user",controller.DeleteCommUser)		//删除用户
			user_group.POST("/comm_user",controller.AddCommUser)			//添加社区用户

			user_group.GET("/topics",controller.ShowAllQuestions)			//查看所有话题
			user_group.GET("/topic",controller.ShowAnswers)					//查看话题下的回答
			user_group.DELETE("/topic",controller.AdminDeleteQues)			//删除话题
			user_group.DELETE("/answer",controller.AdminDeleteAnswer)		//删除话题下的回答


		//super_admin
		admin_group := tokenAuthGroup.Group("/admin",middlewares.AuthAdminRight)		//管理员才有此权限
			admin_group.GET("/admins",controller.ShowAllUsers)		//查看所有用户
			admin_group.POST("/admin",controller.DoAddUser)			//添加用户
			admin_group.DELETE("/admin",controller.DeleteUser)		//删除用户

		//community_user
		comm_user_group := tokenAuthGroup.Group("/comm_user",middlewares.AuthCommUserRight)		//社区用户权限
			comm_user_group.GET("/user",controller.ShowCommUserInfo)		//个人信息
			comm_user_group.PUT("/user",controller.UpdateCommUserInfo)		//更改资料
			comm_user_group.PUT("/passwd",controller.CommUserChangePasswd)	//修改密码

			comm_user_group.POST("/topic",controller.AddQuestion)				//用户添加话题
			comm_user_group.POST("/answer",controller.AddAnswer)				//用户添加回答
			comm_user_group.DELETE("/topic",controller.UserDeleteQues)			//用户删除自己发布的话题
			comm_user_group.DELETE("/answer",controller.DeleteAnswer)			//用户删除自己发布话题下的回答
			comm_user_group.GET("/topics",controller.ShowAllQuestions)			//查看所有话题
			comm_user_group.GET("/topic",controller.ShowAnswers)				//查看话题的回答
			comm_user_group.GET("/my_topics",controller.ShowUserQues)			//查看自己发布的话题
			comm_user_group.GET("/search",controller.SearchQuestion)			//搜索话题

	}


	otherGroup := r.Group("/v1/oth")
	{
		otherGroup.POST("/admin_login",controller.Login)
		otherGroup.POST("/comm_login",controller.CommUserLogin)
		otherGroup.GET("/articles",controller.ShowAllArticles)		//所有文章
		otherGroup.GET("/article",controller.GetOneArtile)			//单个文章
		otherGroup.GET("/news",controller.GetNews)					//新闻
		otherGroup.GET("/notices",controller.GetNotices)			//通知
		otherGroup.GET("/introduction",controller.GetIntroduction)	//简介
		otherGroup.GET("/search",controller.SearchArticle)			//搜索
		
	}

	r.Run(":80")
}