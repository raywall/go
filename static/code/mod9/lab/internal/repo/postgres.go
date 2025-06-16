package repo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
	"gorm.io/gorm"
)

var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
)

type PostgresRepositorio struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NovoPostgresRepositorio(db *gorm.DB, logger *slog.Logger) *PostgresRepositorio {
	return &PostgresRepositorio{db: db, logger: logger}
}

func (r *PostgresRepositorio) Criar(nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)
		return models.Produto{}, ErrPrecoInvalido
	}

	produto := models.Produto{Nome: nome, Preco: preco}
	if err := r.db.Create(&produto).Error; err != nil {
		r.logger.Error("Falha ao criar produto no banco", "error", err)
		return models.Produto{}, fmt.Errorf("criar produto: %w", err)
	}

	r.logger.Info("Produto criado", "id", produto.ID, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *PostgresRepositorio) Buscar(id uuid.UUID) (models.Produto, error) {
	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)
			return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

		r.logger.Error("Falha ao buscar produto no banco", "error", err)
		return models.Produto{}, fmt.Errorf("buscar produto: %w", err)
	}

	r.logger.Info("Produto encontrado", "id", id)
	return produto, nil
}

func (r *PostgresRepositorio) Listar() ([]models.Produto, error) {
	var produtos []models.Produto
	if err := r.db.Find(&produtos).Error; err != nil {
		r.logger.Error("Falha ao listar produtos", "error", err)
		return nil, fmt.Errorf("listar produtos: %w", err)
	}

	r.logger.Info("Listando produtos", "total", len(produtos))
	return produtos, nil
}

func (r *PostgresRepositorio) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)
		return models.Produto{}, ErrPrecoInvalido
	}

	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)

			return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

		r.logger.Error("Falha ao buscar produto no banco", "error", err)
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

	produto.Nome = nome
	produto.Preco = preco
	if err := r.db.Save(&produto).Error; err != nil {
		r.logger.Error("Falha ao atualizar produto no banco", "error", err)
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

	r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *PostgresRepositorio) Deletar(id uuid.UUID) error {
	result := r.db.Delete(&models.Produto{}, "id = ?", id)
	if result.Error != nil {
		r.logger.Error("Falha ao deletar produto no banco", "error", result.Error)
		return fmt.Errorf("deletar produto: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)
		return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

	r.logger.Info("Produto deletado", "id", id)
	return nil
}
