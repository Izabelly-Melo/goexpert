package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			time.Sleep(time.Second)
			c2 <- msg
		}
	}()

	//for i := 0; i < 5; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("Message: %d - %s\n", msg1.id, msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("Message: %d - %s\n", msg2.id, msg2.Msg)
		case <-time.After(time.Second * 3):
			fmt.Printf("timeout\n")
			//default:
			//	fmt.Printf("default\n")
		}
	}
}
