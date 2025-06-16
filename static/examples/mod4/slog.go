package main

import (
	"errors"
	"log/slog"
	"os"
)

func main() {
	// Configurar logger JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Log com campos estruturados
	logger.Info("Processando requisição", "id", 45, "metodo", "GET")
	err := errors.New("falha na operação")

	logger.Error("Erro encontrado", "error", err, "id", 45)
}
