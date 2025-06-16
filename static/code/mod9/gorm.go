package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Produto struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string `gorm:"not null"`
	Preco float64
}

func main() {
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Criar tabela automaticamente
	db.AutoMigrate(&Produto{})

	// Criar produto
	db.Create(&Produto{Nome: "Laptop", Preco: 999.99})

	// Buscar produto
	var produto Produto
	db.First(&produto, "nome = ?", "Laptop")
	log.Printf("Produto: %+v", produto)
}
