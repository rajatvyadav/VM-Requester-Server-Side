package promanagerservice

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2/bson"
)

func GetAllVmRequestFromDB() []models.VMRequest {
	res := []models.VMRequest{}
	if err := helper.Collection().Find(nil).All(&res); err != nil {
		fmt.Println("error is: ", err)
	}
	fmt.Println("4444444444444444444")
	models.VMRequestList = res
	fmt.Println(models.VMRequestList)
	return res
}

func SaveRequestToDB(vmrequest models.VMRequest) error {
	fmt.Println("33333333333333333333")
	// uuidVar, _ := uuid.NewV4() // 47891-sskjsknfdsf-89624-dhnkskfhdkf
	// vmrequest.Id = uuidVar.String()
	//fmt.Println(user)
	vmrequest.RequestId = strings.ToLower(vmrequest.ProgramManager) + strconv.Itoa(rand.Intn(1000))
	return helper.Collection().Insert(vmrequest)
}

//User Login check
func GetUserFromDb(luser models.Login) models.Login {
	fmt.Println("3333333333333333")
	res := models.Login{}
	fmt.Println(luser)
	filter := bson.D{{"password", luser.Password}}
	//{"userId", luser.UserId},
	//filter := bson.D{{"name", "Ash"}}
	//(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0})
	if err := helper.Collection1().Find(filter).One(&res); err != nil {
		fmt.Println(err)
		fmt.Println("44444444444")
		return models.Login{}
	}

	//models.LoginList = res
	fmt.Println("^^^^^^^^^^^^^^^^^", res)
	return res
}
