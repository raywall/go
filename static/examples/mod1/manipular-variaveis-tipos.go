package main

import "fmt"

func main() {
	// Declaração de variáveis
	var nome string = "Raywall"
	idade := 45

	const versao = "1.0.0"

	// Tipos
	var preco float64 = 99.99
	ativo := true

	// Exibição
	fmt.Printf("Nome: %s, Idade: %d, Versão: %s\n", nome, idade, versao)
	fmt.Printf("Preço: %.2f, Ativo: %t\n", preco, ativo)

	// Função com múltiplos retornos
	x, y := troca("Go", "Java")
	fmt.Printf("Troca: %s, %s\n", x, y)
}

func troca(a, b string) (string, string) {
	return b, a
}
