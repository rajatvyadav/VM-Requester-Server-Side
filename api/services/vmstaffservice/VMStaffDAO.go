package vmstaffservice

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"

	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2/bson"
)

func UpdateRequestFromDB(vmrequest models.VMRequest) error {
	fmt.Println("###################", vmrequest)
	filter := bson.M{"requestid": vmrequest.RequestId}
	//update := bson.M{"$set": {bson.M{"serverName": vmrequest.ServerName,"privateIp": vmrequest.PrivateIp}}
	update := bson.M{
		"$set": bson.M{
			"servername":    vmrequest.ServerName,
			"privateip":     vmrequest.PrivateIp,
			"publicip":      vmrequest.PublicIp,
			"currentstatus": vmrequest.CurrentStatus,
		},
	}
	//update := bson.M{"$set": bson.M{"task": todo.Task}}
	//ServerName     string        `json:"serverName"`
	//PrivateIp      string        `json:"privateIp"`
	//PublicIp       string        `json:"publicIp"`
	//CurrentStatus  string        `json:"currentStatus"`
	fmt.Println("UPdate done")
	err := helper.Collection().Update(filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
