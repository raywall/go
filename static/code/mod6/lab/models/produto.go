package models

import "github.com/google/uuid"

// Produto representa um produto no sistema
type Produto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Preco float64   `json:"preco"`
}
