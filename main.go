package main

import (
	"gin/bubble/dao"
	"gin/bubble/models"
	"gin/bubble/routers"
	"gin/bubble/tool"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	//dir,_ := os.Getwd()
	//fmt.Println("当前路径：",dir)
	// 获取配置
	cfg, err := tool.ParseConfig("bubble/config/app.json")
	if err != nil {
		panic(err.Error())
	}

	if err:=dao.InitMysql(cfg);err!=nil{
		panic(err.Error()) // 执行出错，程序退出
	}

	defer dao.Close() // 延时关闭数据库
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 调用路由
	app:= routers.SetUpRouter()
	// 使用自定义的ip和端口
	if err := app.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		log.Fatal(err.Error())
	}
}


