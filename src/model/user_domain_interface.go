package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetId() string
	EncryptPass()

	SetID(string)
}

type userDomainStruct struct {
	id       string
	email    string
	password string
	name     string
	age      int
}
