package main

import (
	"fmt"
	"sync"
)

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func main() {
	c := Contador{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Incrementar()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final:", c.valor) // Saída: Valor final: 1000
}
