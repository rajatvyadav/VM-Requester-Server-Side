package routes

import (
	handlers "VMCreationWorkflow/api/handler/promanager"

	"github.com/gin-gonic/gin"
)

func Init(o *gin.RouterGroup) {

	o.POST("program_manager/request_vm", handlers.AddRequestVm())
	o.GET("program_manager/show_requestvm", handlers.ShowRequestVM())
	//o.GET("todos/view", handlers.ViewTodos())
	//c.POST("todos/delete", handlers.DeleteTodo())
}
