package vmstaffservice

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"

	// "gopkg.in/mgo.v2/bson"
)

// func NewVMRequest() models.VMRequest {
// 	return models.VMRequest{}
// }
// func GetVMRequest_List() []models.VMRequest {
// 	if models.VMRequestList == nil || len(models.VMRequestList) == 0 {
// 		models.VMRequestList = []models.VMRequest{}
// 		return models.VMRequestList
// 	}
// 	return models.VMRequestList
// }

func UpdateRequest_Service(vmrequest models.VMRequest) {

	fmt.Println("!!!!!!!!!!!!!!!!!!", vmrequest)
	//MongoDb
	UpdateRequestFromDB(vmrequest)
	//GetAllUserFromDB()
}
func ShowPendingRequestService() []models.VMRequest {
	res := []models.VMRequest{}
	if err := helper.Collection().Find(nil).All(&res); err != nil {
		fmt.Println("error is: ", err)
	}
	fmt.Println("4444444444444444444")
	models.VMRequestList = res
	fmt.Println(models.VMRequestList)

	return res
}
