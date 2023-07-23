package model

import (
	"crypto/md5"
	"encoding/hex"
)

func NewUserDomain(email, password, name string, age int) *userDomainStruct {
	return &userDomainStruct{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomainStruct) EncryptPass() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomainStruct) GetEmail() string    { return ud.email }
func (ud *userDomainStruct) GetPassword() string { return ud.password }
func (ud *userDomainStruct) GetName() string     { return ud.name }
func (ud *userDomainStruct) GetAge() int         { return ud.age }
func (ud *userDomainStruct) GetId() string       { return ud.id }
func (ud *userDomainStruct) SetID(id string)     { ud.id = id }
