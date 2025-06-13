Módulo 9 – Persistência com Banco de Dados em Go

Este módulo aborda a persistência de dados em Go, usando o pacote padrão database/sql, o ORM GORM, migrações com golang-migrate, e testes de integração com bancos de dados. O conteúdo é voltado para engenheiros Java, com comparações ao Spring Data e Hibernate, e mantém-se objetivo com exemplos práticos e casos de uso. O lab prático implementa o CRUD dos módulos anteriores com persistência em PostgreSQL, incluindo migrações e testes de integração.

  

1. Drivers e `database/sql`

Pacote `database/sql`

- Parte da biblioteca padrão de Go, fornece uma interface genérica para bancos de dados SQL.
- Requer um driver específico (ex.: github.com/lib/pq para PostgreSQL).
- Usa conexão pooling e é thread-safe.

Exemplo (PostgreSQL):

package main

  

import (

    "database/sql"

    _ "github.com/lib/pq"

    "log"

)

  

func main() {

    connStr := "user=postgres password=secret dbname=mydb sslmode=disable"

    db, err := sql.Open("postgres", connStr)

    if err != nil {

        log.Fatal(err)

    }

    defer db.Close()

  

    var nome string

    err = db.QueryRow("SELECT nome FROM produtos WHERE id = $1", 1).Scan(&nome)

    if err != nil {

        log.Fatal(err)

    }

    log.Println("Nome:", nome)

}

Instalação do driver PostgreSQL:

go get github.com/lib/pq

Comparação com Java:

- Java: Usa JDBC ou Spring Data para abstrair acesso a bancos.
- Go: database/sql é mais explícito, com menos abstrações que Hibernate.

Caso de uso: Ideal para aplicações que precisam de controle fino sobre consultas SQL, como relatórios ou sistemas legados.

  

2. ORM com `GORM`

GORM

- Biblioteca popular que fornece uma API ORM, semelhante ao Hibernate em Java.
- Suporta PostgreSQL, MySQL, SQLite, etc., com funcionalidades como mapeamento de structs, associações e migrações automáticas.

Instalação:

go get gorm.io/gorm

go get gorm.io/driver/postgres

Exemplo:

package main

  

import (

    "gorm.io/driver/postgres"

    "gorm.io/gorm"

    "log"

)

  

type Produto struct {

    ID    uint   `gorm:"primaryKey"`

    Nome  string `gorm:"not null"`

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

Comparação com Java:

- Java: Hibernate/Spring Data usa anotações como @Entity e @Column.
- Go: GORM usa tags (gorm:"...") e é mais leve, com menos overhead.

Caso de uso: GORM é ideal para APIs RESTful que precisam de mapeamento objeto-relacional rápido, como sistemas de e-commerce.

  

3. Migrations com `golang-migrate`

`golang-migrate`

- Biblioteca para gerenciar migrações de banco de dados, semelhante ao Flyway ou Liquibase em Java.
- Suporta SQL ou arquivos Go para definir migrações.

Instalação:

go get -u github.com/golang-migrate/migrate/v4

Exemplo (Migração SQL):

1. Crie arquivos de migração:  
    migrations/
2. ├── 202506120001_create_produtos.up.sql
3. └── 202506120001_create_produtos.down.sql
4.   
    

- 202506120001_create_produtos.up.sql:  
    CREATE TABLE produtos (
-     id UUID PRIMARY KEY,
-     nome VARCHAR(255) NOT NULL,
-     preco DOUBLE PRECISION NOT NULL
- );
-   
    
- 202506120001_create_produtos.down.sql:  
    DROP TABLE produtos;
-   
    

6. Código para aplicar migrações:

package main

  

import (

    "github.com/golang-migrate/migrate/v4"

    _ "github.com/golang-migrate/migrate/v4/database/postgres"

    _ "github.com/golang-migrate/migrate/v4/source/file"

    "log"

)

  

func main() {

    m, err := migrate.New(

        "file://migrations",

        "postgres://postgres:secret@localhost:5432/mydb?sslmode=disable",

    )

    if err != nil {

        log.Fatal(err)

    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {

        log.Fatal(err)

    }

    log.Println("Migrações aplicadas")

}

Comparação com Java:

- Java: Flyway/Liquibase gerenciam migrações com SQL ou Java.
- Go: golang-migrate é mais simples, focado em SQL ou Go.

Caso de uso: Migrações garantem consistência do schema em ambientes de desenvolvimento, teste e produção.

  

4. Repositórios e Testes de Integração com Banco de Dados

Repositórios

- Abstraem o acesso ao banco, usando database/sql ou GORM.
- Seguem o padrão Repository, como no Spring Data.

Testes de Integração

- Testam a interação com o banco, usando containers (ex.: Testcontainers) ou bancos em memória (ex.: SQLite).
- Requerem configuração do ambiente de teste.

Caso de uso: Testes de integração garantem que queries e mapeamentos funcionem corretamente em um banco real.

  

📌 Lab: Persistir o CRUD em Banco Real (PostgreSQL)

Objetivo

Refatorar o CRUD do Módulo 8 para usar PostgreSQL com GORM, aplicar migrações com golang-migrate, e implementar testes de integração.

Pré-requisitos

- PostgreSQL instalado e rodando (ex.: docker run -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres).
- Banco mydb criado:psql -U postgres -c "CREATE DATABASE mydb;"
-   
    

Passo a Passo

1. Configuração:

- Use a estrutura do Módulo 8:  
    lab6/
- ├── go.mod
- ├── cmd/
- │   └── api/
- │       └── main.go
- ├── internal/
- │   └── repo/
- │       ├── memoria.go
- │       ├── postgres.go
- │       └── postgres_test.go
- ├── models/
- │   └── produto.go
- ├── migrations/
- │   ├── 202506120001_create_produtos.up.sql
- │   └── 202506120001_create_produtos.down.sql
-   
    
- Adicione dependências:  
    go get gorm.io/gorm
- go get gorm.io/driver/postgres
- go get github.com/golang-migrate/migrate/v4
-   
    
- Atualize go.mod:  
    module github.com/seu-usuario/lab6
-   
    
- go 1.21
-   
    
- require (
-     github.com/gin-gonic/gin v1.10.0
-     github.com/go-playground/validator/v10 v10.22.0
-     github.com/google/uuid v1.6.0
-     gorm.io/driver/postgres v1.5.9
-     gorm.io/gorm v1.25.12
-     github.com/golang-migrate/migrate/v4 v4.17.1
- )
-   
    

3. Implementação:

- migrations/202506120001_create_produtos.up.sql:  
    CREATE TABLE produtos (
-     id UUID PRIMARY KEY,
-     nome VARCHAR(255) NOT NULL,
-     preco DOUBLE PRECISION NOT NULL
- );
-   
    
- migrations/202506120001_create_produtos.down.sql:  
    DROP TABLE produtos;
-   
    
- models/produto.go:  
    package models
-   
    
- import (
-     "github.com/google/uuid"
-     "gorm.io/gorm"
- )
-   
    
- type Produto struct {
-     ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
-     Nome  string    `json:"nome" gorm:"not null" binding:"required,min=3"`
-     Preco float64   `json:"preco" gorm:"not null" binding:"required,gt=0"`
- }
-   
    
- // Antes de criar, gerar UUID
- func (p *Produto) BeforeCreate(tx *gorm.DB) error {
-     p.ID = uuid.New()
-     return nil
- }
-   
    
- internal/repo/postgres.go:  
    package repo
-   
    
- import (
-     "errors"
-     "fmt"
-     "github.com/seu-usuario/lab6/models"
-     "gorm.io/gorm"
-     "log/slog"
- )
-   
    
- var (
-     ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
-     ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
- )
-   
    
- type PostgresRepositorio struct {
-     db     *gorm.DB
-     logger *slog.Logger
- }
-   
    
- func NovoPostgresRepositorio(db *gorm.DB, logger *slog.Logger) *PostgresRepositorio {
-     return &PostgresRepositorio{db: db, logger: logger}
- }
-   
    
- func (r *PostgresRepositorio) Criar(nome string, preco float64) (models.Produto, error) {
-     if preco < 0 {
-         r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)
-         return models.Produto{}, ErrPrecoInvalido
-     }
-     produto := models.Produto{Nome: nome, Preco: preco}
-     if err := r.db.Create(&produto).Error; err != nil {
-         r.logger.Error("Falha ao criar produto no banco", "error", err)
-         return models.Produto{}, fmt.Errorf("criar produto: %w", err)
-     }
-     r.logger.Info("Produto criado", "id", produto.ID, "nome", nome, "preco", preco)
-     return produto, nil
- }
-   
    
- func (r *PostgresRepositorio) Buscar(id uuid.UUID) (models.Produto, error) {
-     var produto models.Produto
-     if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
-         if errors.Is(err, gorm.ErrRecordNotFound) {
-             r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)
-             return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
-         }
-         r.logger.Error("Falha ao buscar produto no banco", "error", err)
-         return models.Produto{}, fmt.Errorf("buscar produto: %w", err)
-     }
-     r.logger.Info("Produto encontrado", "id", id)
-     return produto, nil
- }
-   
    
- func (r *PostgresRepositorio) Listar() ([]models.Produto, error) {
-     var produtos []models.Produto
-     if err := r.db.Find(&produtos).Error; err != nil {
-         r.logger.Error("Falha ao listar produtos", "error", err)
-         return nil, fmt.Errorf("listar produtos: %w", err)
-     }
-     r.logger.Info("Listando produtos", "total", len(produtos))
-     return produtos, nil
- }
-   
    
- func (r *PostgresRepositorio) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
-     if preco < 0 {
-         r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)
-         return models.Produto{}, ErrPrecoInvalido
-     }
-     var produto models.Produto
-     if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
-         if errors.Is(err, gorm.ErrRecordNotFound) {
-             r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)
-             return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
-         }
-         r.logger.Error("Falha ao buscar produto no banco", "error", err)
-         return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
-     }
-     produto.Nome = nome
-     produto.Preco = preco
-     if err := r.db.Save(&produto).Error; err != nil {
-         r.logger.Error("Falha ao atualizar produto no banco", "error", err)
-         return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
-     }
-     r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
-     return produto, nil
- }
-   
    
- func (r *PostgresRepositorio) Deletar(id uuid.UUID) error {
-     result := r.db.Delete(&models.Produto{}, "id = ?", id)
-     if result.Error != nil {
-         r.logger.Error("Falha ao deletar produto no banco", "error", result.Error)
-         return fmt.Errorf("deletar produto: %w", result.Error)
-     }
-     if result.RowsAffected == 0 {
-         r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)
-         return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
-     }
-     r.logger.Info("Produto deletado", "id", id)
-     return nil
- }
-   
    
- cmd/api/main.go (ajustado para PostgreSQL):package main
-   
    
- import (
-     "github.com/gin-gonic/gin"
-     "github.com/seu-usuario/lab6/internal/repo"
-     "github.com/seu-usuario/lab6/models"
-     "github.com/golang-migrate/migrate/v4"
-     _ "github.com/golang-migrate/migrate/v4/database/postgres"
-     _ "github.com/golang-migrate/migrate/v4/source/file"
-     "gorm.io/driver/postgres"
-     "gorm.io/gorm"
-     "log/slog"
-     "net/http"
-     "os"
-     "time"
- )
-   
    
- func main() {
-     logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
-   
    
-     // Aplicar migrações
-     m, err := migrate.New(
-         "file://migrations",
-         "postgres://postgres:secret@localhost:5432/mydb?sslmode=disable",
-     )
-     if err != nil {
-         logger.Error("Falha ao inicializar migrações", "error", err)
-         os.Exit(1)
-     }
-     if err := m.Up(); err != nil && err != migrate.ErrNoChange {
-         logger.Error("Falha ao aplicar migrações", "error", err)
-         os.Exit(1)
-     }
-     logger.Info("Migrações aplicadas")
-   
    
-     // Conectar ao banco
-     dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
-     db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
-     if err != nil {
-         logger.Error("Falha ao conectar ao banco", "error", err)
-         os.Exit(1)
-     }
-   
    
-     // Inicializar repositório
-     repo := repo.NovoPostgresRepositorio(db, logger)
-   
    
-     r := gin.Default()
-     r.Use(func(c *gin.Context) {
-         start := time.Now()
-         c.Next()
-         logger.Info("Requisição processada",
-             "method", c.Request.Method,
-             "path", c.Request.URL.Path,
-             "status", c.Writer.Status(),
-             "duration", time.Since(start))
-     })
-   
    
-     // Rotas
-     produtos := r.Group("/produtos")
-     {
-         produtos.POST("", func(c *gin.Context) {
-             var p models.Produto
-             if err := c.ShouldBindJSON(&p); err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
-                 return
-             }
-             produto, err := repo.Criar(p.Nome, p.Preco)
-             if err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
-                 return
-             }
-             c.JSON(http.StatusCreated, produto)
-         })
-   
    
-         produtos.GET("", func(c *gin.Context) {
-             produtos, err := repo.Listar()
-             if err != nil {
-                 c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
-                 return
-             }
-             c.JSON(http.StatusOK, produtos)
-         })
-   
    
-         produtos.GET("/:id", func(c *gin.Context) {
-             id, err := uuid.Parse(c.Param("id"))
-             if err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
-                 return
-             }
-             produto, err := repo.Buscar(id)
-             if err != nil {
-                 c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
-                 return
-             }
-             c.JSON(http.StatusOK, produto)
-         })
-   
    
-         produtos.PUT("/:id", func(c *gin.Context) {
-             id, err := uuid.Parse(c.Param("id"))
-             if err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
-                 return
-             }
-             var p models.Produto
-             if err := c.ShouldBindJSON(&p); err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
-                 return
-             }
-             produto, err := repo.Atualizar(id, p.Nome, p.Preco)
-             if err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
-                 return
-             }
-             c.JSON(http.StatusOK, produto)
-         })
-   
    
-         produtos.DELETE("/:id", func(c *gin.Context) {
-             id, err := uuid.Parse(c.Param("id"))
-             if err != nil {
-                 c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
-                 return
-             }
-             if err := repo.Deletar(id); err != nil {
-                 c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
-                 return
-             }
-             c.Status(http.StatusNoContent)
-         })
-     }
-   
    
-     r.Run(":8080")
- }
-   
    
- internal/repo/postgres_test.go (Testes de Integração):package repo
-   
    
- import (
-     "github.com/google/uuid"
-     "github.com/seu-usuario/lab6/models"
-     "github.com/stretchr/testify/assert"
-     "gorm.io/driver/postgres"
-     "gorm.io/gorm"
-     "log/slog"
-     "os"
-     "testing"
- )
-   
    
- func TestPostgresRepositorio(t *testing.T) {
-     // Configurar banco de teste
-     dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
-     db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
-     assert.NoError(t, err)
-     defer db.Exec("TRUNCATE produtos RESTART IDENTITY CASCADE")
-   
    
-     logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
-     repo := NovoPostgresRepositorio(db, logger)
-   
    
-     t.Run("Criar produto com sucesso", func(t *testing.T) {
-         produto, err := repo.Criar("Laptop", 999.99)
-         assert.NoError(t, err)
-         assert.NotEqual(t, uuid.Nil, produto.ID)
-         assert.Equal(t, "Laptop", produto.Nome)
-         assert.Equal(t, 999.99, produto.Preco)
-     })
-   
    
-     t.Run("Criar produto com preço inválido", func(t *testing.T) {
-         _, err := repo.Criar("Laptop", -1)
-         assert.ErrorIs(t, err, ErrPrecoInvalido)
-     })
-   
    
-     t.Run("Buscar produto existente", func(t *testing.T) {
-         produto, err := repo.Criar("Mouse", 29.99)
-         assert.NoError(t, err)
-   
    
-         encontrado, err := repo.Buscar(produto.ID)
-         assert.NoError(t, err)
-         assert.Equal(t, produto, encontrado)
-     })
-   
    
-     t.Run("Buscar produto inexistente", func(t *testing.T) {
-         _, err := repo.Buscar(uuid.New())
-         assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
-     })
-   
    
-     t.Run("Listar produtos", func(t *testing.T) {
-         db.Exec("TRUNCATE produtos RESTART IDENTITY CASCADE")
-         repo.Criar("Laptop", 999.99)
-         repo.Criar("Mouse", 29.99)
-   
    
-         produtos, err := repo.Listar()
-         assert.NoError(t, err)
-         assert.Len(t, produtos, 2)
-     })
-   
    
-     t.Run("Atualizar produto existente", func(t *testing.T) {
-         produto, err := repo.Criar("Laptop", 999.99)
-         assert.NoError(t, err)
-   
    
-         atualizado, err := repo.Atualizar(produto.ID, "Laptop Pro", 1299.99)
-         assert.NoError(t, err)
-         assert.Equal(t, "Laptop Pro", atualizado.Nome)
-         assert.Equal(t, 1299.99, atualizado.Preco)
-     })
-   
    
-     t.Run("Deletar produto existente", func(t *testing.T) {
-         produto, err := repo.Criar("Laptop", 999.99)
-         assert.NoError(t, err)
-   
    
-         err = repo.Deletar(produto.ID)
-         assert.NoError(t, err)
-   
    
-         _, err = repo.Buscar(produto.ID)
-         assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
-     })
- }
-   
    

5. Execução:

- Inicie o PostgreSQL (ex.: via Docker):  
    docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres
- psql -U postgres -c "CREATE DATABASE mydb;"
-   
    
- Aplique migrações e execute a API:  
    go run cmd/api/main.go
-   
    
- Teste com curl:  
    curl -X POST http://localhost:8080/produtos -H "Content-Type: application/json" -d '{"nome":"Laptop","preco":999.99}'
- curl http://localhost:8080/produtos
-   
    
- Execute testes:  
    go test ./internal/repo -v
-   
    

7. Tarefa:

- Adicione validação no GORM para garantir que o nome tenha no mínimo 3 caracteres.
- Implemente um teste de integração para os endpoints da API usando httptest.
- Use golang-migrate para adicionar uma nova coluna (ex.: created_at) à tabela produtos.

Saída esperada (exemplo com curl):

# POST /produtos

{"id":"","nome":"Laptop","preco":999.99}

# GET /produtos

[{"id":"","nome":"Laptop","preco":999.99}]

Saída esperada (logs JSON, exemplo):

{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Migrações aplicadas"}

{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}

Caso de uso prático: Este lab simula um backend de e-commerce com persistência em PostgreSQL, migrações automatizadas e testes de integração, ideal para aplicações em produção.

  

Conclusão

Este módulo cobriu persistência com database/sql, ORM com GORM, migrações com golang-migrate, e testes de integração com bancos de dados. O lab prático integrou o CRUD com PostgreSQL, garantindo robustez e escalabilidade. Engenheiros Java notarão semelhanças com Spring Data/Hibernate, mas com a abordagem mais explícita e leve de Go.

Próximos passos: No próximo módulo, exploraremos deploy em produção, monitoramento e boas práticas para escalabilidade, consolidando o aprendizado para aplicações reais.
