package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffer de tamanho 2
	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // Saída: 1
	fmt.Println(<-ch) // Saída: 2
}
