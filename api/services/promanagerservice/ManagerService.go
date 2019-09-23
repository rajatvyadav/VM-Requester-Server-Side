package promanagerservice

import (
	models "VMCreationWorkflow/api/model"
	"fmt"
)

func NewVMRequest() models.VMRequest {
	return models.VMRequest{}
}
func GetVMRequest_List() []models.VMRequest {
	if models.VMRequestList == nil || len(models.VMRequestList) == 0 {
		models.VMRequestList = []models.VMRequest{}
		return models.VMRequestList
	}
	return models.VMRequestList
}
func AddVMRequest_Service(vmrequest models.VMRequest) {
	// Mongo DB
	fmt.Println("22222222222222")
	SaveRequestToDB(vmrequest)
	//GetAllUFromDB()
}

func ShowRequestVM_Service() []models.VMRequest {
	// Mongo DB
	fmt.Println("22222222222222")
	models.VMRequestList = GetAllVmRequestFromDB()
	return models.VMRequestList

}
