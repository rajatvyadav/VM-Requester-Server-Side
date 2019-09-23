package services

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func ValidatedUser(userRequest models.User) (models.User, bool, error) {
	user := GetUser(userRequest)

	if user.UserName != "" {
		return user, true, nil
	}
	return user, false, nil
}

func GetUser(user models.User) models.User {
	temp := models.User{}
	// res := models.User{}
	err := helper.Collection1().Find(bson.M{"username": user.UserName, "password": user.Password, "userrole": user.UserRole}).One(&temp)
	// err := helper.Collection1().Find(nil).All(&temp)
	fmt.Println(err)
	if err == nil {

		fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%")
		fmt.Println(err, "888888", temp)

		return temp
	}

	return temp
}
