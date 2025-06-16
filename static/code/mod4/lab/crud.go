package main

import (
	"errors"
	"fmt"
	"log/slog"
)

// Erros personalizados
var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
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
	logger    *slog.Logger
}

// Novo repositório com logger
func NovoRepositorioEmMemoria(logger *slog.Logger) *RepositorioEmMemoria {
	return &RepositorioEmMemoria{
		produtos:  make([]Produto, 0), // Inicializa um slice vazio para produtos
		idCounter: 0,                  // Inicializa o contador de IDs
		logger:    logger,             //Associa o logger ao repositório
	}
}

// Implementação da interface
func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)
		return Produto{}, ErrPrecoInvalido
	}

	r.idCounter++
	produto := Produto{ID: r.idCounter, Nome: nome, Preco: preco}
	r.produtos = append(r.produtos, produto)
	r.logger.Info("Produto criado", "id", produto.ID, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id int) (Produto, error) {
	for _, p := range r.produtos {
		if p.ID == id {
			r.logger.Info("Produto encontrado", "id", id)
			return p, nil
		}
	}

	r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return Produto{}, fmt.Errorf("buscar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

func (r *RepositorioEmMemoria) Listar() ([]Produto, error) {
	r.logger.Info("Listando produtos", "total", len(r.produtos))
	return r.produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id int, nome string, preco float64) (Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)
		return Produto{}, ErrPrecoInvalido
	}

	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
			return r.produtos[i], nil
		}
	}

	r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return Produto{}, fmt.Errorf("atualizar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

func (r *RepositorioEmMemoria) Deletar(id int) error {
	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos = append(r.produtos[:i], r.produtos[i+1:]...)
			r.logger.Info("Produto deletado", "id", id)
			return nil
		}
	}

	r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return fmt.Errorf("deletar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

// Função para exibir produtos
func exibirProdutos(repo RepositorioProdutos) error {
	produtos, err := repo.Listar()
	if err != nil {
		return fmt.Errorf("exibir produtos: %w", err)
	}
	for _, p := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}

	return nil
}

func main() {
	// Inicializar repositório com logger
	repo := NovoRepositorioEmMemoria()

	// Criar produtos
	p1, err := repo.Criar("Laptop", 999.99)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	repo.Criar("Mouse", 29.99)

	// Listar produtos
	fmt.Println("Lista de produtos:")
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Buscar produto
	if p, err := repo.Buscar(p1.ID); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

	// Atualizar produto
	if p, err := repo.Atualizar(p1.ID, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

	// Tentar atualizar com preço inválido
	if _, err := repo.Atualizar(p1.ID, "Laptop Pro", -100); err != nil {
		fmt.Println("Erro esperado:", err) // Saída: preço não pode ser negativo
	}

	// Deletar produto
	if err := repo.Deletar(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	} else {
		fmt.Println("Erro:", err)
	}

	// Listar novamente
	fmt.Println("Lista final:")
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
	}
}
