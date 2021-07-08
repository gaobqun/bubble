package routers

import (
	"gin/bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine  {

	app := gin.Default()
	// 告诉gin，模版文件引用的静态文件去哪里找
	app.Static("/static", "bubble/static")
	// 告诉gin框架去哪里找模板文件
	app.LoadHTMLGlob("bubble/templates/*")
	app.GET("/", controller.GetHtml)

	// 定义api组 v1 版本
	GroupV1 := app.Group("v1")
	{
		// 待办事项列表
		GroupV1.GET("/todo", controller.GetList)
		// 添加
		GroupV1.POST("/todo", controller.PostAdd)
		// 查看
		GroupV1.GET("/todo/:id", controller.GetOne)
		// 修改
		GroupV1.PUT("/todo/:id", controller.PostUpdate)
		// 删除
		GroupV1.DELETE("/todo/:id", controller.GetDel)
	}

	return app
}
