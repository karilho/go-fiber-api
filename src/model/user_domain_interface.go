package model

// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetId() string
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
	id       string
	email    string
	password string
	name     string
	age      int
}
