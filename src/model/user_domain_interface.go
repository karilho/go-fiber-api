package model

import "github.com/karilho/go-fiber-api/src/configuration/rest_errors"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetId() string

	EncryptPass()
	GenerateToken() (string, *rest_errors.RestErr)

	SetID(string)
}

type userDomainStruct struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}
