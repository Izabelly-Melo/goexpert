package main

import "fmt"

type ID int

var (
	a ID = 1
	b float32
)

/*
	fmt.Printf -> Quer formatar
	fmt.Println -> Quer simplicidade
	fmt.Print -> Quer imprimir sem quebrar linha
*/
func main() {
	fmt.Printf("type of variable 'a': %T\n", a)
	fmt.Printf("value of variable 'a': %v\n", a)
	fmt.Printf("type of variable 'b': %T\n", b)
}
