package repo

import (
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgresRepositorio(t *testing.T) {
	// Configurar banco de teste
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	defer db.Exec("TRUNCATE produtos RESTART IDENTITY CASCADE")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NovoPostgresRepositorio(db, logger)

	t.Run("Criar produto com sucesso", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, produto.ID)
		assert.Equal(t, "Laptop", produto.Nome)
		assert.Equal(t, 999.99, produto.Preco)
	})

	t.Run("Criar produto com preço inválido", func(t *testing.T) {
		_, err := repo.Criar("Laptop", -1)
		assert.ErrorIs(t, err, ErrPrecoInvalido)
	})

	t.Run("Buscar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Mouse", 29.99)
		assert.NoError(t, err)

		encontrado, err := repo.Buscar(produto.ID)
		assert.NoError(t, err)
		assert.Equal(t, produto, encontrado)
	})

	t.Run("Buscar produto inexistente", func(t *testing.T) {
		_, err := repo.Buscar(uuid.New())
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})

	t.Run("Listar produtos", func(t *testing.T) {
		db.Exec("TRUNCATE produtos RESTART IDENTITY CASCADE")
		repo.Criar("Laptop", 999.99)
		repo.Criar("Mouse", 29.99)

		produtos, err := repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 2)
	})

	t.Run("Atualizar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		atualizado, err := repo.Atualizar(produto.ID, "Laptop Pro", 1299.99)
		assert.NoError(t, err)
		assert.Equal(t, "Laptop Pro", atualizado.Nome)
		assert.Equal(t, 1299.99, atualizado.Preco)
	})

	t.Run("Deletar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		err = repo.Deletar(produto.ID)
		assert.NoError(t, err)

		_, err = repo.Buscar(produto.ID)
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})
}
