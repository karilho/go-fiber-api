package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
)

// Pacote utilizado para armazenar as regras de negócio daquele modelo. Tipo o service.
// Não pode ter as tag pq ele ñ pode ser exportável
// O obj que é conversado na reuest é sempre o que vai estar proximo ao controller
// Aqui ele é fechado dentro da regra de negócio
// Você n deixa o controller chamar o objeto, madna ele chamar a INTERFACE.
type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int
}

// Construtor para ser utilizado que vai instanciar este cara
func NewUserDomain(email, password, name string, age int) *UserDomain {
	return &UserDomain{email, password, name, age}
}

func (ud *UserDomain) EncryptPass() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestErr
	UpdateUser(string) *rest_errors.RestErr
	DeleteUser(string) *rest_errors.RestErr
	FindUser(string) (*UserDomain, *rest_errors.RestErr)
}
