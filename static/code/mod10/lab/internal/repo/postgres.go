package repo

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
)

// PostgresRepositorio implementa o repositório com PostgreSQL.
type PostgresRepositorio struct {
	db     *gorm.DB
	logger *zap.Logger
}

// NovoPostgresRepositorio cria um novo repositório.
func NovoPostgresRepositorio(db *gorm.DB, logger *zap.Logger) *PostgresRepositorio {
	return &PostgresRepositorio{db: db, logger: logger}
}

// Criar adiciona um novo produto ao banco.
func (r *PostgresRepositorio) Criar(nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", zap.Error(ErrPrecoInvalido), zap.String("nome", nome))
		return models.Produto{}, ErrPrecoInvalido
	}

	produto := models.Produto{Nome: nome, Preco: preco}
	if err := r.db.Create(&produto).Error; err != nil {
		r.logger.Error("Falha ao criar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("criar produto: %w", err)
	}

	r.logger.Info("Produto criado", zap.String("id", produto.ID.String()), zap.String("nome", nome), zap.Float64("preco", preco))
	return produto, nil
}

// Buscar recupera um produto pelo ID.
func (r *PostgresRepositorio) Buscar(id uuid.UUID) (models.Produto, error) {
	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao buscar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
			return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

		r.logger.Error("Falha ao buscar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("buscar produto: %w", err)
	}

	r.logger.Info("Produto encontrado", zap.String("id", id.String()))
	return produto, nil
}

// Listar retorna todos os produtos.
func (r *PostgresRepositorio) Listar() ([]models.Produto, error) {
	var produtos []models.Produto
	if err := r.db.Find(&produtos).Error; err != nil {
		r.logger.Error("Falha ao listar produtos", zap.Error(err))
		return nil, fmt.Errorf("listar produtos: %w", err)
	}

	r.logger.Info("Listando produtos", zap.Int("total", len(produtos)))
	return produtos, nil
}

// Atualizar modifica um produto existente.
func (r *PostgresRepositorio) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", zap.Error(ErrPrecoInvalido), zap.String("id", id.String()))
		return models.Produto{}, ErrPrecoInvalido
	}
	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao atualizar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
			return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

		r.logger.Error("Falha ao buscar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

	produto.Nome = nome
	produto.Preco = preco
	if err := r.db.Save(&produto).Error; err != nil {
		r.logger.Error("Falha ao atualizar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

	r.logger.Info("Produto atualizado", zap.String("id", id.String()), zap.String("nome", nome), zap.Float64("preco", preco))
	return produto, nil
}

// Deletar remove um produto pelo ID.
func (r *PostgresRepositorio) Deletar(id uuid.UUID) error {
	result := r.db.Delete(&models.Produto{}, "id = ?", id)
	if result.Error != nil {
		r.logger.Error("Falha ao deletar produto no banco", zap.Error(result.Error))
		return fmt.Errorf("deletar produto: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		r.logger.Error("Falha ao deletar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
		return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

	r.logger.Info("Produto deletado", zap.String("id", id.String()))
	return nil
}
