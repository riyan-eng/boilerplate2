package dto

type RegisterAdmin struct {
	UserName    string
	Email       string
	PhoneNumber string
	Password    string
}

type Login struct {
	UserName string
	Password string
}
