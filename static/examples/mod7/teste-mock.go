package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type RepositorioMock struct {
	mock.Mock
}

func (m *RepositorioMock) Buscar(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestServico(t *testing.T) {
	repo := &RepositorioMock{}
	repo.On("Buscar", 1).Return("Produto", nil)

	resultado, err := repo.Buscar(1)
	if err != nil {
		t.Fatal(err)
	}

	if resultado != "Produto" {
		t.Errorf("esperado Produto, got %s", resultado)
	}
	repo.AssertExpectations(t)
}
