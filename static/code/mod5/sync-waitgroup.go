package main

import (
	"fmt"
	"sync"
)

func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tarefa(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todas as tarefas concluídas")
}
