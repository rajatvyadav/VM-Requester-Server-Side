package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	models "VMCreationWorkflow/api/model"
	services "VMCreationWorkflow/api/services/adminservice"
)

// AddTodo Todo Add Handler
func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.User{}
		c.Bind(&requestBody)
		services.AddUser_Service(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}

// ViewTodos Todo View Handler
func ShowUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//append(models.VMRequestList, requestBody)
		//return models.LoginList
		//requestBody := models.VMRequestList{}
		//requestBody = ShowRequestVM_Service()
		list := services.ShowUser_Service()
		c.JSON(http.StatusOK, list)
	}
}
func CreateProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.Program{}
		c.Bind(&requestBody)
		services.CreateProgramService(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}
func ShowProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		list := services.ShowProgramService()
		c.JSON(http.StatusOK, list)
	}
}

func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.AddProject{}
		c.Bind(&requestBody)
		fmt.Println("!!!!!!!!!!!!", requestBody)
		services.CreateProjectService(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}
