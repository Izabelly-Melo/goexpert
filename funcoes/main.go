package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("a + b: ", sum(2, 2))

	valor, ok := sub(100, 40)
	fmt.Println("a - b: ", valor, ok)

	fmt.Println("=======")
	valor, err := test(1000, 30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("total do teste: ", valor)

	fmt.Println("=======")
	valor, err = test(-100, 30)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("total do teste: ", valor)

}

func sum(a int, b int) int {
	return a + b
}

func sub(a, b int) (int, bool) {
	if a-b >= 60 {
		return a - b, true
	}

	return a - b, false
}

func test(a, b int) (int, error) {
	if a+b <= 0 {
		return 0, errors.New("o valor final não pode ser zero ou negativo")
	}

	return a + b, nil
}
