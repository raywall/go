package main

import (
	"errors"
	"fmt"
)

func main() {
	err := processar(45)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Erro é 'não encontrado'") // Saída: Erro é 'não encontrado'
	}
	fmt.Println(err) // Saída: falha ao processar id 45: não encontrado
}
