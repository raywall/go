package main

import "fmt"

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado) // Resultado: 5
}
