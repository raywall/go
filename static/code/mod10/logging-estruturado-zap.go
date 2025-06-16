package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Aplicação iniciada",
		zap.Int("port", 8080),
		zap.String("env", "prod"),
	)
}
