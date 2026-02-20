package main

import "fmt"

type Conta struct {
	saldo int
}

// função construtora que cria uma nova Conta e retorna um ponteiro
func NewConta() *Conta {
	return &Conta{
		saldo: 0,
	}
}

// método da struct Conta usando receiver ponteiro (*Conta)
// isso permite alterar o valor original da struct
func (c *Conta) simular(valor int) int {
	c.saldo += valor // altera o saldo da conta original
	return c.saldo
}

func main() {

	// cria uma instância da struct Conta
	conta := Conta{
		saldo: 20,
	}

	conta.simular(200) // altera o saldo
	fmt.Printf("Saldo da conta: %d", conta.saldo)
}
