package main

import "fmt"

func gerar(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
func dobrar(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

	return out
}

func main() {
	entrada := gerar(1, 2, 3)
	saida := dobrar(entrada)

	for n := range saida {
		fmt.Println(n) // Saída: 2, 4, 6
	}
}
