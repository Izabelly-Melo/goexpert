package main

import "fmt"

// type assertion é a forma de extrair um valor concreto a partir de uma interface, ou seja, ela permite recuperar o tipo real guardado dentro da interface
func main() {
	/*	var minhaVar interface{} = "Iza"
		println(minhaVar.(string))
		res, ok := minhaVar.(int)
		println(res)
		println(ok)
	*/0,,,,,,,,,,,,,,,,,,
	tipos("Iza")
	tipos(20)
	tipos(true)
}

func tipos(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("String:", v)
		return
	}

	if v, ok := x.(int); ok {
		fmt.Println("Int:", v)
		return
	}

	fmt.Println("Tipo desconhecido")
}
