package helpers

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitDb() {
	fmt.Println("In db function")
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("pt2019")
}

func Collection() *mgo.Collection {
	return db.C("VmRequest")
}
func Collection1() *mgo.Collection {
	return db.C("users")
}
func Collection2() *mgo.Collection {
	return db.C("programs")
}
