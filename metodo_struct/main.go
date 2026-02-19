package main

import "fmt"

type Endereco struct {
	Logradouro string
	Num        int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	iza := Client{
		Nome:  "Iza",
		Idade: 27,
		Ativo: true,
	}

	iza.Cidade = "São Paulo"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s\n", iza.Nome, iza.Idade, iza.Ativo, iza.Cidade)

	iza.Desativar()
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Println("Desativado")
}
