package main

import (
	"fmt"
	"time"
)

/*
*
* Exemplo 6 de channel (LOAD BALANCER)
 */
// Thread 1
func main() {
	data := make(chan int)
	qtdWorkers := 100

	// inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	//recebendo
	for i := 0; i < 1000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Workerd %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
