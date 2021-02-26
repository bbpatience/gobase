package models

import (
	"gobase/global/my_errors"
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
	variables.GormDbMysql.Where("id=?", id).First(todo)
	if  todo.ID <= 0 {
		err = &my_errors.MyError{ErrorString: "id invalid."}
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = variables.GormDbMysql.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	affectRows := variables.GormDbMysql.Debug().Where("id=?", id).Delete(&Todo{}).RowsAffected
	if affectRows <= 0 {
		err = &my_errors.MyError{ErrorString: "id invalid."}
	}
	return
}
