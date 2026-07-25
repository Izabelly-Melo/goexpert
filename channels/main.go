package main

import (
	"fmt"
	"sync"
)

/*
*
* Exemplo 4 de channel (RANGE WAIT GROUPS)
*
 */
func main() {
	ch := make(chan int) // Canal Vazio
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

//-------------------------------------------------------
/*
*
* Exemplo 3 de channel (RANGE)
*
*
func main() {
	ch := make(chan int) // Canal Vazio
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
*/

//-------------------------------------------------------
/*
*
* Exemplo 2 de channel (FOREVER)
*
*
// Thread 1
func main() {
	forever := make(chan bool) // Canal Vazio

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever

}
*/

//-------------------------------------------------------
/*
*
* Exemplo 1 de channel
*
// Thread 1
func main() {
	channel := make(chan string) // Canal Vazio

	// Thread 2
	func() {
		channel <- "Olá Mundo!" // Canal cheio
	}()

	// Thread 1
	msg := <-channel // Canal esvazia
	fmt.Println(msg)
}
*/
