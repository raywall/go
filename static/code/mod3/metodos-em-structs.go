package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

// Método com receiver por valor
func (p Produto) Descricao() string {
	return fmt.Sprintf("%s: R$%.2f", p.Nome, p.Preco)
}

// Método com receiver por ponteiro
func (p *Produto) AplicarDesconto(desconto float64) {
	p.Preco -= p.Preco * desconto
}

func main() {
	p := Produto{Nome: "Laptop", Preco: 1000}

	fmt.Println(p.Descricao()) // Saída: Laptop: R$1000.00
	p.AplicarDesconto(0.1)

	fmt.Println(p.Descricao()) // Saída: Laptop: R$900.00
}
