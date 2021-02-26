package controller

import (
	"github.com/gin-gonic/gin"
	response "gobase/api/common"
	"gobase/api/models"
	"gobase/global/consts"
	log "gobase/global/variables"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	log.ZapLog.Sugar().Infof("CreateTodo：%v\n", todo)
	if todo.Title == "" {
		response.Fail(c, consts.CommonParamsCheckFailCode, consts.CommonParamsCheckFailMsg)
		return
	}
	err := models.CreateATodo(&todo)
	if err != nil {
		response.Fail(c, consts.DBCommonFailCode, consts.DBCommonFailMsg)
	} else {
		response.Success(c, todo)
	}
}

func GetTodoList(c *gin.Context) {
	log.ZapLog.Sugar().Infof("GetTodoList\n")
	todoList, err := models.GetAllTodo()
	if err != nil {
		response.Fail(c, consts.DBCommonFailCode, consts.DBCommonFailMsg)
	} else {
		response.Success(c, todoList)
	}
}

func GetTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	log.ZapLog.Sugar().Infof("GetTodoById：%s\n", id)
	if !ok {
		response.Fail(c, consts.CommonParamsCheckFailCode, consts.CommonParamsCheckFailMsg)
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		response.Fail(c, consts.DBInvalidIdCode, consts.DBInvalidIdMsg)
	} else {
		response.Success(c, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	log.ZapLog.Sugar().Infof("UpdateATodo：%s\n", id)
	if !ok {
		response.ErrorParam(c, nil)
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		response.Fail(c, consts.DBInvalidIdCode, consts.DBInvalidIdMsg)
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		response.Fail(c, consts.DBCommonFailCode, consts.DBCommonFailMsg)
	} else {
		response.Success(c, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	log.ZapLog.Sugar().Infof("DeleteATodo：%s\n", id)
	if !ok {
		response.Fail(c, consts.CommonParamsCheckFailCode, consts.CommonParamsCheckFailMsg)
	}
	if err := models.DeleteATodo(id); err != nil {
		response.Fail(c, consts.DBInvalidIdCode, consts.DBInvalidIdMsg)
	} else {
		response.Success(c, nil)
	}
}
