package models

import "gin/bubble/dao"

// model
type Todo struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Status bool `json:"status"`
	//CreateTime int `json:"create_time"`
	//UpdateTime int `json:"update_time"`
}

// 新增
func CreateTodo(todo *Todo) (err error)  {
	err = dao.DB.Create(&todo).Error
	return
}

// 编辑
func UpdateTodo(todo *Todo)(err error)  {
	err = dao.DB.Save(todo).Error
	return
}


// 查询列表
func ListTodo() (todoList []*Todo, err error) {
	if err:=dao.DB.Find(&todoList).Error;err!=nil{
		return nil,err
	}else{
		return todoList,nil
	}
}


// 删除
func DelTodo(id string) (err error)  {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

// 单个查询
func OneTodo(id string) (todo *Todo,err error)  {
	todo = new(Todo)
	if err:= dao.DB.Where("id=?", id).First(&todo).Error;err!=nil{
		return nil,err
	}
	return todo,nil
}