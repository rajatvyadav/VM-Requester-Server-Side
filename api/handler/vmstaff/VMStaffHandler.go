package vmstaff

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"VMCreationWorkflow/api/model"
	services "VMCreationWorkflow/api/services/vmstaffservice"
)

// AddTodo Todo Add Handler
func UpdateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.VMRequest{}
		c.Bind(&requestBody)
		services.UpdateRequest_Service(requestBody)
	}
}

func ShowPendingRequestVM() gin.HandlerFunc {
	return func(c *gin.Context) {
		//append(models.VMRequestList, requestBody)
		//return models.LoginList
		//requestBody := models.VMRequestList{}
		//requestBody = ShowRequestVM_Service()
		list := services.ShowPendingRequestService()
		c.JSON(http.StatusOK, list)
	}
}
