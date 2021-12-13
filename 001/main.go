package main

import (
	"fmt"
	"learnGin/001/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
)



func main(){
	r := gin.Default()
	r.Static("/static","static") //前面值的为自定义路由，后面为实际访问的文件目标
	r.LoadHTMLGlob("template/*")

	v1Group := r.Group("v1")
	{
		//Add Task
		v1Group.POST("/todo",func(c *gin.Context) {
			//Step1.Extract Parameter from Request
			todo := new(model.Todo)
			c.BindJSON(todo)
			//Step2.Save the data in database
			err := model.Add(todo)
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,gin.H{
					"msg":"success",
				})
			}
		})

		//Check All Tasks
		v1Group.GET("/todo",func(c *gin.Context) {
			var todo_list []*model.Todo
			todo_list,err := model.QueryAll()
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,todo_list)
			}
		})

		//Delete Task
		v1Group.DELETE("/todo/:id",func(c *gin.Context) {
			id,ok := c.Params.Get("id")
			if !ok{
				c.JSON(200,gin.H{"error":"Invalid ID"})
				return
			}
			err := model.Delete(id)
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,gin.H{"msg":"success"})
			}
		})

		//Update Status
		v1Group.PUT("todo/:id",func(c *gin.Context) {
			//Step1.Veryfy whether the id is vailid
			id,ok := c.Params.Get("id")
			if !ok{
				c.JSON(200,gin.H{"error":"Invalid ID"})
				return
			}
			_,err := model.QueryOne(id)
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}

			//Step2.Get The Information which should be changed from request
			var todo = new(model.Todo)
			c.BindJSON(todo)
			fmt.Println(todo)
			todo.ID,err = strconv.Atoi(id)
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}
			//Step3.Update the information
			err = model.Update(todo)
			if err != nil{
				c.JSON(200,gin.H{"error":err.Error()})
			}else{
				c.JSON(200,gin.H{"msg":"success"})
			}
			
		})
	}
	r.GET("/",func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})

	r.Run(":80")
}