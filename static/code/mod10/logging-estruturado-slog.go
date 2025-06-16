package main

import (
	"log/slog"

	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Aplicação iniciada", "port", 8080)
}
