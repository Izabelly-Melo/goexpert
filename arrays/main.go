package main

import "fmt"

func main() {
	var list [3]int // o meu array vai ter 3 posições fixas
	list[0] = 10
	list[1] = 20
	list[2] = 30

	// Valor do Array
	fmt.Println(list)

	// Tamanho do Array
	fmt.Println(len(list))

	// Imprimir o valor da posição 2 do array
	fmt.Println(list[2])

	for i, value := range list {
		fmt.Printf("index: %d, value: %d\n", i, value)
	}
}
