package main

import (
	"fmt"
)

func main() {
	fmt.Println("soma: ", sum(2, 2, 343, 55, 666, 642, 444, 324))
}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
