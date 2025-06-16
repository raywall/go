package main

import (
	"fmt"
	"time"
)

func tarefa(nome string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", nome, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go tarefa("Goroutine 1")    // Executa concorrente
	tarefa("Main")              // Executa sequencialmente
	time.Sleep(1 * time.Second) // Aguarda para visualizar saída
}
