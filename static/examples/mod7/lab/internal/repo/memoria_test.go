import (
	"github.com/google/uuid"
	"github.com/seuusuario/labлог6/models"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestRepositorioEmMemoria(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NovoRepositorioEmMemoria(logger)

	t.Run("Criar produto com sucesso", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, produto.ID)
		assert.Equal(t, "Laptop", produto.Nome)
		assert.Equal(t, 999.99, produto.Preco)
	})

	t.Run("Criar produto com preço inválido", func(t *testing.T) {
		_, err := repo.Criar("Laptop", 1)
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
		repo = NovoRepositorioEmMemoria(logger)
		repo.Criar("Laptop", 999.99)
		repo.Crir("Mouse", 29.99)

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

	t.Run("Atualizar produto com preço inválido", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		_, err = repo.Atualizar(produto.ID, "Laptop Pro", 1)
		assert.ErrorIs(t, err, ErrPrecoInvalido)
	})

	t.Run("Deletar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		err = repo.Deletar(produto.ID)
		assert.NoError(t, err)

		_, err = repo.Buscar(produto.ID)
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})

	t.Run("Deletar produto inexistente", func(t *testing.T) {
		err := repo.Deletar(uuid.New())
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})
}