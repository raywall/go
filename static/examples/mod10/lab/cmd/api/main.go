package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/internal/repo"
	"github.com/seu-usuario/lab6/models"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configurar logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Configurar OpenTelemetry
	traceExporter, err := stdouttrace.New()
	if err != nil {
		logger.Fatal("Falha ao configurar trace exporter", zap.Error(err))
	}

	tp := trace.NewTracerProvider(trace.WithBatcher(traceExporter))
	defer tp.Shutdown(context.Background())
	otel.SetTracerProvider(tp)

	// Configurar Prometheus
	metricExporter, err := prometheus.New()
	if err != nil {
		logger.Fatal("Falha ao configurar metric exporter", zap.Error(err))
	}
	mp := metric.NewMeterProvider(metric.WithReader(metricExporter))
	defer mp.Shutdown(context.Background())

	// Aplicar migrações
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:secret@postgres:5432/mydb?sslmode=disable",
	)
	if err != nil {
		logger.Fatal("Falha ao inicializar migrações", zap.Error(err))
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatal("Falha ao aplicar migrações", zap.Error(err))
	}
	logger.Info("Migrações aplicadas")

	// Conectar ao banco
	dsn := "host=postgres user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Falha ao conectar ao banco", zap.Error(err))
	}

	// Inicializar repositório
	repo := repo.NovoPostgresRepositorio(db, logger)

	// Configurar Gin
	r := gin.Default()
	tracer := otel.Tracer("api")

	// Middleware de tracing e logging
	r.Use(func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), c.Request.URL.Path)
		defer span.End()
		span.SetAttributes(attribute.String("method", c.Request.Method))

		start := time.Now()
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		logger.Info("Requisição processada",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
		)
	})

	// Expor métricas
	r.GET("/metrics", gin.WrapH(metricExporter))

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
