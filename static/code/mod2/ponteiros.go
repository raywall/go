package main

import "fmt"

func main() {
	x := 45
	p := &x         // Ponteiro para x
	fmt.Println(*p) // 45 (desreferenciação)

	// Modificar via ponteiro
	*p = 100
	fmt.Println(x) // 100
}

func incrementar(p *int) {
	*p++
}
