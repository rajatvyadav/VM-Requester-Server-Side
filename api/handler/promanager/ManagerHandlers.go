package promanager

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "VMCreationWorkflow/api/model"
	services "VMCreationWorkflow/api/services/promanagerservice"
)

// AddTodo Todo Add Handler
func AddRequestVm() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.VMRequest{}
		c.Bind(&requestBody)
		services.AddVMRequest_Service(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}

// ViewTodos Todo View Handler
func ShowRequestVM() gin.HandlerFunc {
	return func(c *gin.Context) {
		//append(models.VMRequestList, requestBody)
		//return models.LoginList
		//requestBody := models.VMRequestList{}
		//requestBody = ShowRequestVM_Service()
		list := services.ShowRequestVM_Service()
		c.JSON(http.StatusOK, list)
	}
}
