package routes

import (
	handlers "VMCreationWorkflow/api/handler/vmstaff"

	"github.com/gin-gonic/gin"
)

func VMStaffInit(o *gin.RouterGroup) {

	o.POST("vmstaff/updateRequest", handlers.UpdateRequest())
	o.GET("staff/show/pendingrequest", handlers.ShowPendingRequestVM())
	//o.GET("admin/getuser", handlers.ShowUser())
	//o.GET("todos/view", handlers.ViewTodos())
	//c.POST("todos/delete", handlers.DeleteTodo())
}
