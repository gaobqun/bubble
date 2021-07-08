package controller

import (
	"gin/bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 访问html
func GetHtml(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

// 查看列表
func GetList(context *gin.Context) {
	if todoList, err := models.ListTodo(); err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 500, Message: err.Error(), Data: err.Error()})
	} else {
		//context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "Ok", Data: todoList})
		context.JSON(http.StatusOK, todoList)
	}
}

// 添加
func PostAdd(context *gin.Context) {
	// 1、获取参数
	var todo models.Todo
	context.BindJSON(&todo)
	// 2、存入数据库
	if err := models.CreateTodo(&todo); err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 500, Message: err.Error(), Data: err.Error()})
	} else {
		//3、返回参数
		context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "Ok", Data: todo})
	}

}

// 查看单个
func GetOne(context *gin.Context) {

}

// 编辑
func PostUpdate(context *gin.Context) {
	// 获取到参数id
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "无效的id"})
		return
	}
	// 更新
	todo, err := models.OneTodo(id)
	if err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 1001, Message: err.Error()})
		return
	}
	context.BindJSON(&todo)
	if err := models.UpdateTodo(todo); err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 1002, Message: err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// 删除
func GetDel(context *gin.Context) {
	// 获取到参数 id
	id, ok := context.Params.Get("id")
	if !ok {
		return
	}
	if err := models.DelTodo(id); err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 1001, Message: err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "删除成功"})
	}

}

// 参数返回定义
type ResponseParam struct {
	Code    int
	Message string
	Data    interface{}
}
