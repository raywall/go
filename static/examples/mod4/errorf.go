package main

import (
	"fmt"
)

func acessarIndice(slice []int, indice int) (int, error) {
	if indice >= len(slice) {
		return 0, fmt.Errorf("índice %d fora do limite (%d)", indice, len(slice))
	}

	return slice[indice], nil
}

func main() {
	values := []int{1, 5, 6, 8, 9}

	result, err := acessarIndice(values, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Result:", result)
}
