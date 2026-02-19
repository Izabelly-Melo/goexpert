package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	iza := Client{
		Nome:  "Iza",
		Idade: 27,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", iza.Nome, iza.Idade, iza.Ativo)

	iza.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", iza.Nome, iza.Idade, iza.Ativo)

	lino := Client{
		Nome:  "Lino",
		Idade: 1,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", lino.Nome, lino.Idade, lino.Ativo)
}
