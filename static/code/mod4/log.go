package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("Iniciando aplicação...")
	err := errors.New("falha crítica")

	if err != nil {
		log.Fatal("Erro:", err) // Sai com código 1
	}
}
