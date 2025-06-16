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

// Interface do repositório
type RepositorioProdutos interface {
	Criar(nome string, preco float64) (Produto, error)
	Buscar(id int) (Produto, error)
	Listar() ([]Produto, error)
	Atualizar(id int, nome string, preco float64) (Produto, error)
	Deletar(id int) error
}

// Repositório em memória
type RepositorioEmMemoria struct {
	produtos  []Produto
	idCounter int
}

// Implementação da interface
func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (Produto, error) {
	if preco < 0 {
		return Produto{}, errors.New("preço não pode ser negativo")
	}

	r.idCounter++
	produto := Produto{ID: r.idCounter, Nome: nome, Preco: preco}

	r.produtos = append(r.produtos, produto)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id int) (Produto, error) {
	for _, p := range r.produtos {
		if p.ID == id {
			return p, nil
		}
	}

	return Produto{}, errors.New("produto não encontrado")
}

func (r *RepositorioEmMemoria) Listar() ([]Produto, error) {
	return r.produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id int, nome string, preco float64) (Produto, error) {
	if preco < 0 {
		return Produto{}, errors.New("preço não pode ser negativo")
	}

	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			return r.produtos[i], nil
		}
	}

	return Produto{}, errors.New("produto não encontrado")
}

func (r *RepositorioEmMemoria) Deletar(id int) error {
	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos = append(r.produtos[:i], r.produtos[i+1:]...)
			return nil
		}
	}

	return errors.New("produto não encontrado")
}

// Função para exibir produtos
func exibirProdutos(repo RepositorioProdutos) {
	produtos, _ := repo.Listar()
	fmt.Println("Lista de produtos:")

	for _, p := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
}

func main() {
	// Inicializar repositório
	repo := &RepositorioEmMemoria{}

	// Criar produtos
	repo.Criar("Laptop", 999.99)
	repo.Criar("Mouse", 29.99)

	// Listar produtos
	exibirProdutos(repo)

	// Buscar produto
	if p, err := repo.Buscar(1); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	}

	// Atualizar produto
	if p, err := repo.Atualizar(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	}

	// Deletar produto
	if err := repo.Deletar(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	}

	// Listar novamente
	exibirProdutos(repo)
}
