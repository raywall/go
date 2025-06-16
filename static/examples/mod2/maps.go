package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["idade"] = 45
	m["preco"] = 99
	fmt.Println(m) // map[idade:42 preco:99]

	// Verificar existência
	valor, existe := m["idade"]
	if existe {
		fmt.Println("Idade:", valor)
	}

	// Deletar
	delete(m, "preco")
	fmt.Println(m) // map[idade:42]
}
