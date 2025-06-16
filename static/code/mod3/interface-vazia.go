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
