package main

import "fmt"

func main() {
	// Laço tradicional
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Como while
	contador := 0
	for contador < 3 {
		fmt.Println(contador)
		contador++
	}

	// Infinito (com break)
	for {
		fmt.Println("Infinito")
		break
	}
}
