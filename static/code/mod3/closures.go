package main

import "fmt"

func contador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c1 := contador()
	fmt.Println(c1()) // Saída: 1
	fmt.Println(c1()) // Saída: 2

	c2 := contador()
	fmt.Println(c2()) // Saída: 1 (novo contador)
}
