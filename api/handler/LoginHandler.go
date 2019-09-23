package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	services "VMCreationWorkflow/api/services"

	helpers "VMCreationWorkflow/api/helpers"

	models "VMCreationWorkflow/api/model"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		c.Bind(&user)
		fmt.Println(user)
		user, isValidUser, err := services.ValidatedUser(user)
		if isValidUser != true {
			fmt.Println("Invalid user!", err)
			// return

			c.JSON(http.StatusOK, "invalid user")
		}

		if isValidUser {
			token, err := helpers.GenerateToken(user.UserName, user.Password, user.UserRole, 24*time.Hour)
			if err != nil {
				fmt.Print("error while generating token:", err)
				// return err
			}
			fmt.Println("Token:", token)

			c.Header("Authorization", token)
			c.JSON(http.StatusOK, user)
		}

	}
}
