package main

import "fmt"

// Declaração de constante no package utilizando
// o modo de declaração curta (por inferência de tipo)
const versao = "1.0.0"

// Declaração de variável no package utilizando
// o modo de declaração explícita
var nome string = "Raywall"

func main() {
	// Declaração de variável por inferência de tipo
	idade := 45

	// Declaração de variáveis de outros tipos no modo
	// explícito e curto
	var preco float64 = 99.99
	ativo := true

	// Exibição dos valores das variáveis e constantes
	fmt.Printf("Nome: %s, Idade: %d, Versão: %s\n", nome, idade, versao)
	fmt.Printf("Preço: %.2f, Ativo: %t\n", preco, ativo)

	// Utilizando funções com múltiplos retornos
	x, y := troca("Go", "Let's")
	fmt.Printf("Troca: %s, %s!\n", x, y)
}

// troca realiza a troca da ordem dos valores informados
// exemplo:
//
//	x, y := troca("Go", "Let's")
//	fmt.Printf("%s, %s\n", x, y)
//
// resultado:
//
//	Let's, Go
func troca(a, b string) (string, string) {
	return b, a
}
