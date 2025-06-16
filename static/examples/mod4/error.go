package main

import (
	"errors"
	"fmt"
)

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}

	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 0)

	if err != nil {
		fmt.Println("Erro:", err) // Saída: Erro: divisão por zero
		return
	}

	fmt.Println("Resultado:", resultado)
}
