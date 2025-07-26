package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

func inspecionar(v interface{}) {
	switch t := v.(type) {
	case Produto:
		fmt.Println("É um Produto:", t.Nome)
	case string:
		fmt.Println("É uma string:", t)
	default:
		fmt.Println("Tipo desconhecido")
	}
}

func main() {
	// inspecionar um Produto
	inspecionar(Produto{Nome: "Laptop", Preco: 1000})

	// inspecionar uma string
	inspecionar("Raywall")

	// inspecionar um inteiro
	inspecionar(45)
}
