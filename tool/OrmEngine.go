package tool

import "github.com/go-xorm/xorm"

// 将整个xorm的engine都包含在这个结构体中
type Orm struct {
	*xorm.Engine
}

// 数据库连接
func OrmEngine(cfg *Config)(*Orm, error)  {
	database:=cfg.Database
	conn:=database.User+":"+database.Pwd+"@tcp("+database.Host+":"+database.Port+")/"+database.DbName+"?charset="+database.Charset
	engine,err:=xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil,err
	}
	engine.ShowSQL(database.Show)

	orm:=new(Orm)
	orm.Engine = engine
	return orm,nil
}