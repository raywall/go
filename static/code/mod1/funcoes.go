package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func troca(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(soma(5, 3)) // Saída: 8
	a, b := troca("hello", "world")
	fmt.Println(a, b) // Saída: world hello
}
