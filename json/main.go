package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s" validade:"gt=0"`
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}
	res, err := json.Marshal(conta) // Marshal é a função que converte um struct em JSON
	if err != nil {
		panic(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	println("------------")
	dados2 := []byte(`{"n":2323,"s":1000}`)
	var conta2 Conta
	err = json.Unmarshal(dados2, &conta2) // Unmarshal é a função que converte um JSON em struct
	if err != nil {
		panic(err)
	}

	println(conta2.Numero)
}
