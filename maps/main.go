package main

import "fmt"

func main() {

	// Criação de um map já com valores iniciais
	// map[tipoDaChave]tipoDoValor
	salarios := map[string]int{"Iza": 20000, "Jack": 15000, "Lino": 3000}

	fmt.Println(salarios)         // Imprime todo o map
	fmt.Println(salarios["Lino"]) // Acessa o valor da chave "Lino"

	fmt.Println("----------")
	salarios["Marisa"] = 50000 // Adiciona uma nova chave ao map
	fmt.Println(salarios)

	// Formas alternativas de criar um map vazio
	// salarios := make(map[string]int)
	// salarios := map[string]int{}

	fmt.Println("----------")
	// Percorre o map pegando chave e valor
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	fmt.Println("----------")
	// Percorre o map pegando apenas os valores
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	fmt.Println("---------- Remove Iza ---------")
	delete(salarios, "Iza") // Remove uma chave do map
	fmt.Println(salarios)

	println("------------")
	valor, ok := salarios["Iza"]

	if ok { // Valida de a chave Iza existe
		fmt.Println("Existe:", valor)
	} else {
		fmt.Println("Não existe")
	}

}
