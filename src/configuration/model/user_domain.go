package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	EncryptPass()
}

// Pacote utilizado para armazenar as regras de negócio daquele modelo. Tipo o service.
// Não pode ter as tag pq ele ñ pode ser exportável
// O obj que é conversado na reuest é sempre o que vai estar proximo ao controller
// Aqui ele é fechado dentro da regra de negócio
// Você n deixa o controller chamar o objeto, madna ele chamar a INTERFACE.
type userDomain struct {
	email    string
	password string
	name     string
	age      int
}

// Construtor para ser utilizado que vai instanciar este cara
func NewUserDomain(email, password, name string, age int) *userDomain {
	return &userDomain{email, password, name, age}
}

func (ud *userDomain) EncryptPass() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int {
	return ud.age
}
