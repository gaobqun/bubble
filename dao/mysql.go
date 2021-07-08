package dao

import (
	"gin/bubble/tool"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql()(err error)  {
	// 获取配置
	cfg, err := tool.ParseConfig("bubble/config/app.json")
	if err != nil {
		return err
		//panic(err.Error())
	}
	// 链接数据库
	database := cfg.Database
	conn := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	if DB, err = gorm.Open("mysql", conn); err != nil {
		//panic(err.Error()) // 执行出错，程序退出
		return err
	}
	return
}

func Close()  {
	DB.Close()
}