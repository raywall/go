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
