package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/seu-usuario/lab6/internal/repo"
)

func exibirProdutos(repo repo.RepositorioProdutos) error {
	produtos, err := repo.Listar()
	if err != nil {
		return fmt.Errorf("exibir produtos: %w", err)
	}

	fmt.Println("Lista de produtos:")
	for _, p := range produtos {
		fmt.Printf("ID: %s, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NovoRepositorioEmMemoria(logger)

	// Criar produtos
	p1, err := repo.Criar("Laptop", 999.99)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	repo.Criar("Mouse", 29.99)

	// Listar produtos
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

	// Deletar produto
	if err := repo.Deletar(p1.ID); err == nil {
		fmt.Println("Produto deletado com sucesso")
	} else {
		fmt.Println("Erro:", err)
	}

	// Listar novamente
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
	}
}
