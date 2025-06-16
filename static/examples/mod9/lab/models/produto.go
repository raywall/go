package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Produto struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Nome  string    `json:"nome" gorm:"not null" binding:"required,min=3"`
	Preco float64   `json:"preco" gorm:"not null" binding:"required,gt=0"`
}

// Antes de criar, gerar UUID
func (p *Produto) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
