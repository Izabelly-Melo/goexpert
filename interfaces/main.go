package main

import "fmt"

type Endereco struct {
	Logradouro string
	Num        int
	Cidade     string
	Estado     string
}

// basta ter os mesmos métodos para implementar uma interface — não precisa declarar explicitamente
type Pessoa interface {
	Desativar()
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	iza := Client{
		Nome:  "Iza",
		Idade: 27,
		Ativo: true,
	}

	Desativacao(iza)

}

func (c Client) Desativar() { // só lê → valor
	c.Ativo = false
	fmt.Println("Desativado")
}
