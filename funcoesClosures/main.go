package main

import "fmt"

func main() {
	total := func() int {
		return sum(2, 2, 40, 60) * 2
	}()

	fmt.Println("total: ", total)
}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
