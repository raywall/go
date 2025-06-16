package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Canal 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Canal 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido:", msg2)
		}
	}
}
