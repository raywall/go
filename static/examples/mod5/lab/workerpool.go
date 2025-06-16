package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

// Tarefa representa uma tarefa a ser processada
type Tarefa struct {
	ID    int
	Valor int
}

// Resultado representa o resultado de uma tarefa
type Resultado struct {
	ID       int
	Quadrado int
}

// Worker processa tarefas de um channel
func worker(id int, tarefas chan Tarefa, resultados chan Resultado, wg *sync.WaitGroup, logger *slog.Logger) {
	defer wg.Done()
	for t := range tarefas {
		logger.Info("Processando tarefa", "worker", id, "tarefa_id", t.ID, "valor", t.Valor)
		time.Sleep(100 * time.Millisecond) // Simula trabalho
		resultados <- Resultado{ID: t.ID, Quadrado: t.Valor * t.Valor}
	}
}

func main() {
	// Configurar logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Canais para tarefas e resultados
	tarefas := make(chan Tarefa, 10)
	resultados := make(chan Resultado, 10)
	var wg sync.WaitGroup

	// Iniciar 3 workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tarefas, resultados, &wg, logger)
	}

	// Enviar tarefas
	numTarefas := 5
	go func() {
		for i := 1; i <= numTarefas; i++ {
			tarefas <- Tarefa{ID: i, Valor: i}
			logger.Info("Tarefa enviada", "tarefa_id", i, "valor", i)
		}
		close(tarefas)
	}()

	// Coletar resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Exibir resultados
	for r := range resultados {
		fmt.Printf("Tarefa %d: Quadrado = %d\n", r.ID, r.Quadrado)
	}
	logger.Info("Processamento concluído", "total_tarefas", numTarefas)
}
