package model

type User struct {
	LoginID  string `json:"loginId"`
	UserName string `json:"username"`
	Password string `json:"password"`
	UserRole string `json:"userrole"`
	EmailId  string `json:"emailid"`
}

var UserList []User

type Login struct {
	LoginID  string `json:"loginid"`
	Password string `json:"password"`
	UserRole string `json:"userrole"`
}
type Temp struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserRole string `json:"userrole"`
	EmailID  string `json:"emailid"`
}
