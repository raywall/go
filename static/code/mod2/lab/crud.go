package main

import (
	"errors"
	"fmt"
)

// Struct para Produto
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// Banco de dados em memória (slice de produtos)
var produtos []Produto
var idCounter = 1

// Create: Adiciona um produto
func criarProduto(nome string, preco float64) Produto {
	produto := Produto{ID: idCounter, Nome: nome, Preco: preco}
	idCounter++

	produtos = append(produtos, produto)
	return produto
}

// Read: Busca todos os produtos
func listarProdutos() []Produto {
	return produtos
}

// Read: Busca produto por ID
func buscarProduto(id int) (Produto, error) {
	for _, p := range produtos {
		if p.ID == id {
			return p, nil
		}
	}

	return Produto{}, errors.New("Produto não encontrado")
}

// Update: Atualiza um produto
func atualizarProduto(id int, nome string, preco float64) (Produto, error) {
	for i, p := range produtos {
		if p.ID == id {
			produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			return produtos[i], nil
		}
	}

	return Produto{}, errors.New("Produto não encontrado")
}

// Delete: Remove um produto
func deletarProduto(id int) error {
	for i, p := range produtos {
		if p.ID == id {
			produtos = append(produtos[:i], produtos[i+1:]...)
			return nil
		}
	}
	return errors.New("Produto não encontrado")
}

func main() {
	// Criar produtos
	criarProduto("Laptop", 999.99)
	criarProduto("Mouse", 29.99)

	// Listar produtos
	fmt.Println("Lista de produtos:")
	for _, p := range listarProdutos() {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}

	// Buscar produto
	if p, err := buscarProduto(1); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	}

	// Atualizar produto
	if p, err := atualizarProduto(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	}

	// Deletar produto
	if err := deletarProduto(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	}

	// Listar novamente
	fmt.Println("Lista final:")

	for _, p := range listarProdutos() {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
}
