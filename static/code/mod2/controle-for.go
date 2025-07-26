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

	// Range
	numeros := []int{1, 2, 3}

	fmt.Println("Números do slice:")
	for indice, valor := range numeros {
		fmt.Printf("%d. %d\n", indice, valor)
	}
}
