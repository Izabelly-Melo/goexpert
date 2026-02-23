package main

import "encoding/json"

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}
	res, err := json.Marshal(conta) // Marshal é a função que converte um struct em JSON
	if err != nil {
		panic(err)
	}
	println(string(res))
}
