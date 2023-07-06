package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
)

// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int

	GetJSONValue() (string, error)
	EncryptPass()

	SetID(string)
}

// Pacote utilizado para armazenar as regras de negócio daquele modelo. Tipo o service.
// Não pode ter as tag pq ele ñ pode ser exportável
// O obj que é conversado na reuest é sempre o que vai estar proximo ao controller
// Aqui ele é fechado dentro da regra de negócio
// Você n deixa o controller chamar o objeto, madna ele chamar a INTERFACE.
// Quando eu trasnformo o struct privado mas deixo os atributos PUBLICOS, eu consigo acessar os atributos e exporta-los,
// mas não consigo acessar o struct, somente via instanciação.
type userDomainStruct struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int
}

func (ud *userDomainStruct) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error on marshalling userDomainStruct", err)
		return "", err
	}
	logger.Info(string(b))
	return string(b), nil
}

// Construtor para ser utilizado que vai instanciar este cara
func NewUserDomain(email, password, name string, age int) *userDomainStruct {
	return &userDomainStruct{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomainStruct) EncryptPass() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

// Getters.

func (ud *userDomainStruct) GetEmail() string {
	return ud.Email
}

func (ud *userDomainStruct) GetPassword() string {
	return ud.Password
}

func (ud *userDomainStruct) GetName() string {
	return ud.Name
}

func (ud *userDomainStruct) GetAge() int {
	return ud.Age
}

func (ud *userDomainStruct) SetID(id string) {
	ud.ID = id
}

func (ud *userDomainStruct) GetID() string {
	return ud.ID
}
