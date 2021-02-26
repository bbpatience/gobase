package routers

import (
	"github.com/gin-gonic/gin"
	"gobase/api/controller"
	"gobase/global/variables"
)

func SetupRouter() *gin.Engine {
	if !variables.ConfigYml.GetBool("AppDebug") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
		// 获取某一个待办事项
		v1Group.GET("/todo/:id", controller.GetTodoById)
	}
	return r
}
