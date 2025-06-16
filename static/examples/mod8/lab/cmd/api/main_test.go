package main

import (
	"github.com/seuusuario/lab6/internal/repo"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NovoRepositorioEmMemoria(logger)

	t.Run("Fluxo completo do CRUD", func(t *testing.T) {

		// Criar
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		// Listar
		produtos, err := repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 1)

		// Buscar
		encontrado, err := repo.Buscar(produto.ID)
		assert.NoError(t, err)
		assert.Equal(t, produto, encontrado)

		// Atualizar
		atualizado, err := repo.Atualizar(produto.ID, "Laptop Pro", 1299.99)
		assert.NoError(t,, err)
		assert.Equal(t, "Laptop Pro", atualizado.Nome)

		// Deletar
		err = repo.Deletar(produto.ID)
		assert.NoError(t, err)

		// Verificar exclusão
		produtos, err = repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 0)
	})
}