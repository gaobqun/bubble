package main

import (
	"gin/bubble/tool"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var (
	DB *gorm.DB
)

// 访问html
func getHtml(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

// model
type Todo struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Status bool `json:"status"`
	//CreateTime int `json:"create_time"`
	//UpdateTime int `json:"update_time"`
}

func main() {
	//dir,_ := os.Getwd()
	//fmt.Println("当前路径：",dir)
	// 获取配置
	cfg, err := tool.ParseConfig("bubble/config/app.json")
	if err != nil {
		panic(err.Error())
	}
	// 链接数据库
	database := cfg.Database
	conn := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	if DB, err = gorm.Open("mysql", conn); err != nil {
		panic(err.Error()) // 执行出错，程序退出
	}
	defer DB.Close() // 延时关闭数据库
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	app := gin.Default()
	// 告诉gin，模版文件引用的静态文件去哪里找
	app.Static("/static", "bubble/static")
	// 告诉gin框架去哪里找模板文件
	app.LoadHTMLGlob("bubble/templates/*")
	app.GET("/", getHtml)

	// 定义api组 v1 版本
	GroupV1 := app.Group("v1")
	{
		// 待办事项列表
		GroupV1.GET("/todo", getList)
		// 添加
		GroupV1.POST("/todo", postAdd)
		// 查看
		GroupV1.GET("/todo/:id", getOne)
		// 修改
		GroupV1.PUT("/todo/:id", postUpdate)
		// 删除
		GroupV1.DELETE("/todo/:id", getDel)
	}

	// 使用自定义的ip和端口
	if err := app.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		log.Fatal(err.Error())
	}
}

// 查看列表
func getList(context *gin.Context) {
	var todoList []Todo //定义一个切片
	if err:=DB.Find(&todoList).Error;err!=nil{
		context.JSON(http.StatusOK, ResponseParam{Code: 500, Message: err.Error(), Data: err.Error()})
	}else{
		//context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "Ok", Data: todoList})
		context.JSON(http.StatusOK, todoList)
	}
}

// 添加
func postAdd(context *gin.Context) {
	// 1、获取参数
	var todo Todo
	context.BindJSON(&todo)
	// 2、存入数据库
	if err := DB.Create(&todo).Error; err != nil {
		context.JSON(http.StatusOK, ResponseParam{Code: 500, Message: err.Error(), Data: err.Error()})
	}else{
		//3、返回参数
		context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "Ok", Data: todo})
	}


}

// 查看单个
func getOne(context *gin.Context) {

}

// 编辑
func postUpdate(context *gin.Context) {
	// 获取到参数id
	id,ok:=context.Params.Get("id")
	if !ok {
		return
	}
	// 更新
	var todo Todo
	if err:= DB.Where("id=?", id).First(&todo).Error;err!=nil{
		context.JSON(http.StatusOK, ResponseParam{Code: 1001, Message: err.Error()})
		return
	}
	context.BindJSON(&todo)
	if err:= DB.Save(&todo).Error;err!=nil{
		context.JSON(http.StatusOK, ResponseParam{Code: 1002, Message: err.Error()})
	}else{
		context.JSON(http.StatusOK, todo)
	}
}

// 删除
func getDel(context *gin.Context) {
	// 获取到参数 id
	id,ok:=context.Params.Get("id")
	if !ok {
		return
	}
	if err:= DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
		context.JSON(http.StatusOK, ResponseParam{Code: 1001, Message: err.Error()})
		return
	}else{
		context.JSON(http.StatusOK, ResponseParam{Code: 200, Message: "删除成功"})
	}

}

// 参数返回定义
type ResponseParam struct {
	Code    int
	Message string
	Data    interface{}
}
