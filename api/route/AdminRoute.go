package routes

import (
	handlers "VMCreationWorkflow/api/handler/admin"

	"github.com/gin-gonic/gin"
)

func AdminInit(o *gin.RouterGroup) {

	o.POST("admin/adduser", handlers.AddUser())
	o.GET("admin/getuser", handlers.ShowUser())
	o.POST("admin/createProgram", handlers.CreateProgram())
	o.GET("admin/showPrograms", handlers.ShowProgram())
	o.POST("admin/createProject", handlers.CreateProject())
}
