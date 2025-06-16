package repo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
)

var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
)

// RepositorioProdutos define a interface do repositório
type RepositorioProdutos interface {
	Criar(nome string, preco float64) (models.Produto, error)
	Buscar(id uuid.UUID) (models.Produto, error)
	Listar() ([]models.Produto, error)
	Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error)
	Deletar(id uuid.UUID) error
}

// RepositorioEmMemoria implementa o repositório em memória
type RepositorioEmMemoria struct {
	produtos map[uuid.UUID]models.Produto
	logger   *slog.Logger
}

// NovoRepositorioEmMemoria cria um novo repositório
func NovoRepositorioEmMemoria(logger *slog.Logger) *RepositorioEmMemoria {
	return &RepositorioEmMemoria{
		produtos: make(map[uuid.UUID]models.Produto),
		logger:   logger,
	}
}

func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)
		return models.Produto{}, ErrPrecoInvalido
	}

	id := uuid.New()
	produto := models.Produto{ID: id, Nome: nome, Preco: preco}
	r.produtos[id] = produto
	r.logger.Info("Produto criado", "id", id, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id uuid.UUID) (models.Produto, error) {
	produto, existe := r.produtos[id]
	if !existe {
		r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)
		return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

	r.logger.Info("Produto encontrado", "id", id)
	return produto, nil
}

func (r *RepositorioEmMemoria) Listar() ([]models.Produto, error) {
	var produtos []models.Produto
	for _, p := range r.produtos {
		produtos = append(produtos, p)
	}

	r.logger.Info("Listando produtos", "total", len(produtos))
	return produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)
		return models.Produto{}, ErrPrecoInvalido
	}

	produto, existe := r.produtos[id]
	if !existe {
		r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)
		return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

	produto.Nome = nome
	produto.Preco = preco

	r.produtos[id] = produto
	r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Deletar(id uuid.UUID) error {
	if _, existe := r.produtos[id]; !existe {
		r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)
		return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

	delete(r.produtos, id)
	r.logger.Info("Produto deletado", "id", id)
	return nil
}
