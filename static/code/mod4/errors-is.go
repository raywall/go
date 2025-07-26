package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("não encontrado")

func buscarDado(id int) error {
	return ErrNotFound
}

func processar(id int) error {
	err := buscarDado(id)
	if err != nil {
		return fmt.Errorf("falha ao processar id %d: %w", id, err)
	}
	return nil
}

func main() {
	err := processar(45)
	if err := processar(45); errors.Is(err, ErrNotFound) {
		fmt.Println("Erro é 'não encontrado'") // Saída: Erro é 'não encontrado'
	}
	fmt.Println(err) // Saída: falha ao processar id 45: não encontrado
}
