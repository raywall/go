package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Mensagem da goroutine"
	}()

	msg := <-ch      // Bloqueia até receber
	fmt.Println(msg) // Saída: Mensagem da goroutine
}
