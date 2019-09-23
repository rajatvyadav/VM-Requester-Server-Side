package adminservice

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	//"gopkg.in/mgo.v2/bson"
)

func GetAllUserFromDB() []models.User {
	res := []models.User{}
	if err := helper.Collection1().Find(nil).All(&res); err != nil {
		fmt.Println("error is: ", err)
	}
	fmt.Println("4444444444444444444")
	models.UserList = res
	fmt.Println(models.UserList)
	return res
}

func SaveUserToDB(user models.User) error {
	fmt.Println("33333333333333333333")
	// uuidVar, _ := uuid.NewV4() // 47891-sskjsknfdsf-89624-dhnkskfhdkf
	// vmrequest.Id = uuidVar.String()
	//fmt.Println(user)
	user.LoginID = strings.ToLower(user.UserName) + strconv.Itoa(rand.Intn(1000))
	return helper.Collection1().Insert(user)
}
