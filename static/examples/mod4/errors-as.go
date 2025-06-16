package main

import (
	"errors"
	"fmt"
)

type MeuErro struct {
	Mensagem string
}

func (e *MeuErro) Error() string {
	return e.Mensagem
}

func operacao() error {
	return fmt.Errorf("contexto: %w", &MeuErro{Mensagem: "falha específica"})
}

func main() {
	err := operacao()

	var meuErro *MeuErro
	if errors.As(err, &meuErro) {
		fmt.Println("Erro capturado:", meuErro.Mensagem) // Saída: falha específica
	}
}
