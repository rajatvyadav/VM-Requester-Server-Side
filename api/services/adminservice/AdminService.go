package adminservice

import (
	helper "VMCreationWorkflow/api/helpers"
	models "VMCreationWorkflow/api/model"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func NewUser() models.User {
	return models.User{}
}
func GetUser_List() []models.User {
	if models.UserList == nil || len(models.UserList) == 0 {
		models.UserList = []models.User{}
		return models.UserList
	}
	return models.UserList
}
func AddUser_Service(user models.User) {
	// Mongo DB
	fmt.Println("22222222222222")
	SaveUserToDB(user)
	//GetAllUFromDB()
}

func ShowUser_Service() []models.User {
	// Mongo DB
	fmt.Println("22222222222222")
	models.UserList = GetAllUserFromDB()
	return models.UserList

}

func CreateProgramService(newProgram models.Program) error {

	return helper.Collection2().Insert(newProgram)

}
func ShowProgramService() []models.Program {
	temp := []models.Program{}
	err := helper.Collection2().Find(nil).All(&temp)
	if err == nil {
		fmt.Println("success")
	}

	return temp
}
func CreateProjectService(newProgram models.AddProject) error {

	//pro := models.Project{}
	//pro.ProjectId = newProgram.ProjectId
	//pro.ProjectName = newProgram.ProjectName
	fmt.Println(newProgram)
	newProgram.ProjectId = strings.ToLower(newProgram.ProjectName) + strconv.Itoa(rand.Intn(1000))
	filter := bson.M{"programname": newProgram.ProgramName}
	//update := bson.M{"$set": {bson.M{"serverName": vmrequest.ServerName,"privateIp": vmrequest.PrivateIp}}
	//update := bson.M{"$set": bson.M{"programid": newProgram.ProgramId}}
	update := bson.M{"$addToSet": bson.M{"project": bson.M{"projectId": newProgram.ProjectId, "projectName": newProgram.ProjectName}}}

	fmt.Println("Update done", newProgram)
	err := helper.Collection2().Update(filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
