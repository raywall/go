import "github.com/google/uuid"

type Produto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome" binding:"required"`
	Preco float64   `json:"preco" binding:"required,gt=0"`
}