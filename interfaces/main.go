package main

import "fmt"

type Endereco struct {
	Logradouro string
	Num        int
	Cidade     string
	Estado     string
}

// Só pode passar apenas métodos para implementar a interface basta ter o mesmo métodos
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
