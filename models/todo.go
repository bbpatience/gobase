package models

import (
	"gobase/global/variables"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	err = variables.GormDbMysql.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = variables.GormDbMysql.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = variables.GormDbMysql.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = variables.GormDbMysql.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = variables.GormDbMysql.Where("id=?", id).Delete(&Todo{}).Error
	return
}
