package main

import "fmt"

type Descrivivel interface {
	Descricao() string
}

type Produto struct {
	Nome  string
	Preco float64
}

type Servico struct {
	Nome string
	Taxa float64
}

func (p Produto) Descricao() string {
	return fmt.Sprintf("Produto: %s, R$%.2f", p.Nome, p.Preco)
}

func (s Servico) Descricao() string {
	return fmt.Sprintf("Serviço: %s, R$%.2f", s.Nome, s.Taxa)
}

func exibir(d Descrivivel) {
	fmt.Println(d.Descricao())
}

func main() {
	p := Produto{Nome: "Laptop", Preco: 1000}
	s := Servico{Nome: "Consultoria", Taxa: 500}

	exibir(p) // Saída: Produto: Laptop, R$1000.00
	exibir(s) // Saída: Serviço: Consultoria, R$500.00
}
