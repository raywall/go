package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/internal/repo"
	"github.com/seu-usuario/lab6/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Aplicar migrações
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:secret@localhost:5432/mydb?sslmode=disable",
	)
	if err != nil {
		logger.Error("Falha ao inicializar migrações", "error", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("Falha ao aplicar migrações", "error", err)
		os.Exit(1)
	}
	logger.Info("Migrações aplicadas")

	// Conectar ao banco
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Falha ao conectar ao banco", "error", err)
		os.Exit(1)
	}

	// Inicializar repositório
	repo := repo.NovoPostgresRepositorio(db, logger)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.Info("Requisição processada",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start))
	})

	// Rotas
	produtos := r.Group("/produtos")
	{
		produtos.POST("", func(c *gin.Context) {
			var p models.Produto
			if err := c.ShouldBindJSON(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			produto, err := repo.Criar(p.Nome, p.Preco)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, produto)
		})

		produtos.GET("", func(c *gin.Context) {
			produtos, err := repo.Listar()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, produtos)
		})

		produtos.GET("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			produto, err := repo.Buscar(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, produto)
		})

		produtos.PUT("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			var p models.Produto
			if err := c.ShouldBindJSON(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			produto, err := repo.Atualizar(id, p.Nome, p.Preco)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, produto)
		})

		produtos.DELETE("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}
			if err := repo.Deletar(id); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}

	r.Run(":8080")
}
