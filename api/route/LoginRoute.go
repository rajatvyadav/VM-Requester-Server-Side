package routes

import (
	handlers "VMCreationWorkflow/api/handler"

	"github.com/gin-gonic/gin"
)

func InitLoginRoute(o *gin.RouterGroup) {
	o.POST("/login/loginuser", handlers.LoginUser())
}
