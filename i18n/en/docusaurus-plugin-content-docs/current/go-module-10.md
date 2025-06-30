---
sidebar_position: 11
sidebar_label: Módulo 10
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

:::warning Attention
This page was not translated yet.
:::

# Deploy, observabilidade e boas ppráticas em Go

<div className="row">
<div className="col">

Este módulo aborda o _deploy de aplicações Go_, `observabilidade com logging`, `tracing` e `métricas`, e `boas práticas` para produção. Focado em engenheiros Java, compara com práticas do Spring Boot e ferramentas como `Prometheus`, com exemplos práticos e objetivos para consulta futura.

O lab prático containeriza a API CRUD dos módulos anteriores, configurando logging estruturado, tracing com `OpenTelemetry`, e exposição de métricas.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default} 
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-1rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Build com `go build` e cross-compilation

### `go build`

- Compila código Go em executáveis.
- Simples, com suporte a flags para otimização.

#### Exemplo

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/build.go)!

#### Executar

```bash
go build -o app
./app # Saída: Hello, World!
```

### Cross-compilation

- _Go suporta compilação para diferentes sistemas operacionais e arquiteturas_, como:

  - `Linux`
  - `Windows`
  - `ARM`

- Definido com variáveis de ambiente `GOOS` e `GOARCH`

#### Exemplo

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux
GOOS=windows GOARCH=amd64 go build -o app-windows.exe
```

#### Comparação com Java

##### Java

- Compila para bytecode (JVM), com dependências de runtime.

##### Go

- Gera **binários nativos**, _sem dependências externas_.

:::info Caso de uso
Cross-compilation é ideal para criar binários para containers ou dispositivos embarcados.
:::

<br />

## Docker com Go

### Docker

- Go é **altamente compatível com Docker**, gerando **binários pequenos** e **imagens leves**.
- Usa `multi-stage builds` para **reduzir o tamanho da imagem**.

#### Exemplo

```Dockerfile
# Etapa de build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/Dockerfile)!

#### Construir e executar

```bash
docker build -t meu-app .
docker run -p 8080:8080 meu-app
```

#### Comparação com Java

##### Java

- Imagens _Spring Boot são maiores **devido à JVM**_.

##### Go

- **Imagens são menores** (ex: `~10MB com Alpine`), ideais para microsserviços

:::info Caso de uso
**Containerizar APIs** para deploy em `Kubernetes` ou plataformas como `AWS ECS`.
:::

<br />

## Logging estruturado (`slog` e `zap`)

### `slog`

- Pacote padrão do Go (1.21+) para logging estruturado, com saída JSON.
- Já usado nos módulos anteriores.

#### Exemplo

```go
package main

import (
	"log/slog"

	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Aplicação iniciada", "port", 8080)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/logging-estruturado-slog.go)!

### `zap`

- Biblioteca de logging de alta performance, alternativa a slog.
- **Instalação**: `go get go.uber.org/zap`

#### Exemplo

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Aplicação iniciada",
		zap.Int("port", 8080),
		zap.String("env", "prod"),
	)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/logging-estruturado-zap.go)!

#### Comparação com Java

##### Java

- Usa `SLF4J`/`Logback` para _logging estruturado_.

##### Go

- `slog` **é nativo e simples**; `zap` é _mais rápido para sistemas de alta carga_.

:::info Caso de uso
Logging estruturado é essencial para monitoramento em ferramentas como `ELK Stack` ou `Grafana Loki`.
:::

<br />

## Tracing com OpenTelemetry

### OpenTelemetry

- Framework para _tracing distribuído, métricas e logs_, compatível com `Jaeger`, `Zipkin`, etc.
- Permite rastrear requisições em microsserviços.

#### Instalação

```bash
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace
go get go.opentelemetry.io/otel/sdk/trace
```

#### Exemplo

```go
package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	exporter, err := stdouttrace.New()
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)
	tracer := otel.Tracer("meu-app")
	_, span := tracer.Start(context.Background(), "operação-principal")
	defer span.End()

	log.Println("Operação rastreada")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/opentelemetry.go)!

#### Comparação com Java

##### Java

- Usa `Spring Cloud Sleuth` ou `Micrometer` para tracing.

##### Go

- `OpenTelemetry` é **mais flexível**, com integração nativa.

:::info Caso de uso
`Tracing` é _crucial para diagnosticar latências_ em sistemas distribuídos, como APIs em Kubernetes.
:::

<br />

## Linter, cobertura e documentação com `godoc`

### Linter

- Ferramentas como `golangci-lint` _verificam estilo, erros e boas práticas_.

#### Instalação

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### Uso

```bash
golangci-lint run
```

### Cobertura

- Go gera relatórios de cobertura de testes nativamente.

#### Exemplo

```bash
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```

### `godoc`

- Gera documentação automática a partir de comentários no código.

#### Uso local

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
```

> Acesse a documentação via browser através do endereço [http://localhost:6060/pkg](http://localhost:6060/pkg)

#### Exemplo de comentário

```go
// Criar adiciona um novo produto ao repositório.
// Retorna o produto criado ou um erro se o preço for inválido.
func (r *Repositorio) Criar(nome string, preco float64) (models.Produto, error) {
    // ...
}
```

#### Comparação com Java

##### Java

- Usa `Javadoc`, `SonarQube` e `Checkstyle`.

##### Go

- **Ferramentas nativas** como `godoc`, `go test` e `golangci-lint` são mais integradas.

:::info Caso de uso
`Linters` _garantem consistência_, cobertura valida testes, e `godoc` facilita manutenção.
:::

<br />

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

## Laboratório

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

### Pré-requisitos

- `PostgreSQL` instalado e rodando:

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres
```

- Banco `mydb` criado:

```bash
psql -U postgres -c "CREATE DATABASE mydb;"
```

### Configuração

1. Use a estrutura do [Módulo 09](go-module-9.md)

```dirtree
lab6/
├── go.mod
├── cmd/
│   └── api/
│		├── main.go
│       └── main_test.go
├── internal/
│   └── repo/
│		├── memoria.go
│		├── memoria_test.go
│		├── postgres.go
│       └── postgres_test.go
├── models/
│   └── produto.go
├── migrations/
│ 	├── 202506120001_create_produtos.up.sql
│   └── 202506120001_create_produtos.down.sql
├── docker-compose.yml
└── Dockerfile
```

2. Adicione as dependências necessárias:

```bash
go get go.uber.org/zap
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/prometheus
go get go.opentelemetry.io/otel/sdk/metric
go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace
go get go.opentelemetry.io/otel/sdk/trace
```

3. Atualize o arquivo `go.mod`:

```go
module github.com/seu-usuario/lab6

go 1.21

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/uuid v1.6.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
	github.com/golang-migrate/migrate/v4 v4.17.1
	go.uber.org/zap v1.27.0
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/exporters/prometheus v0.50.0
	go.opentelemetry.io/otel/sdk/metric v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0
	go.opentelemetry.io/otel/sdk/trace v1.28.0
)
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/lab/lab10.zip)!

</TabItem>
<TabItem value="app">

### Containerizar o serviço e expor métricas, trace e logs

#### Objetivo

Containerizar a API CRUD do [Módulo 09](go-module-9.md) (`PostgreSQL`), configurar _logging estruturado_ com `zap`, tracing com `OpenTelemetry`, e expor métricas com `Prometheus`. Gerar cobertura de testes e documentação com `godoc`.

#### Passo a passo

1. Crie o arquivo `Dockerfile` com o conteúdo:

```Dockerfile
# Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Package
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY migrations /root/migrations

EXPOSE 8080
CMD ["./app"]
```

2. Atualize o arquivo `models/produto.go` utilizando o código:

```go
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Produto representa um produto no sistema.
type Produto struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Nome  string    `json:"nome" gorm:"not null" binding:"required,min=3"`
	Preco float64   `json:"preco" gorm:"not null" binding:"required,gt=0"`
}

// BeforeCreate gera um UUID antes de salvar no banco.
func (p *Produto) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
```

3. Atualize o arquivo `internal/repo/postgres.go` utilizando o código:

```go
package repo

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
)

// PostgresRepositorio implementa o repositório com PostgreSQL.
type PostgresRepositorio struct {
	db     *gorm.DB
	logger *zap.Logger
}

// NovoPostgresRepositorio cria um novo repositório.
func NovoPostgresRepositorio(db *gorm.DB, logger *zap.Logger) *PostgresRepositorio {
	return &PostgresRepositorio{db: db, logger: logger}
}

// Criar adiciona um novo produto ao banco.
func (r *PostgresRepositorio) Criar(nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", zap.Error(ErrPrecoInvalido), zap.String("nome", nome))
		return models.Produto{}, ErrPrecoInvalido
	}

    produto := models.Produto{Nome: nome, Preco: preco}
	if err := r.db.Create(&produto).Error; err != nil {
		r.logger.Error("Falha ao criar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("criar produto: %w", err)
	}

    r.logger.Info("Produto criado", zap.String("id", produto.ID.String()), zap.String("nome", nome), zap.Float64("preco", preco))
	return produto, nil
}

// Buscar recupera um produto pelo ID.
func (r *PostgresRepositorio) Buscar(id uuid.UUID) (models.Produto, error) {
	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao buscar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
			return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

        r.logger.Error("Falha ao buscar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("buscar produto: %w", err)
	}

    r.logger.Info("Produto encontrado", zap.String("id", id.String()))
	return produto, nil
}

// Listar retorna todos os produtos.
func (r *PostgresRepositorio) Listar() ([]models.Produto, error) {
	var produtos []models.Produto
	if err := r.db.Find(&produtos).Error; err != nil {
		r.logger.Error("Falha ao listar produtos", zap.Error(err))
		return nil, fmt.Errorf("listar produtos: %w", err)
	}

    r.logger.Info("Listando produtos", zap.Int("total", len(produtos)))
	return produtos, nil
}

// Atualizar modifica um produto existente.
func (r *PostgresRepositorio) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", zap.Error(ErrPrecoInvalido), zap.String("id", id.String()))
		return models.Produto{}, ErrPrecoInvalido
	}
	var produto models.Produto
	if err := r.db.First(&produto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error("Falha ao atualizar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
			return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
		}

        r.logger.Error("Falha ao buscar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

    produto.Nome = nome
	produto.Preco = preco
	if err := r.db.Save(&produto).Error; err != nil {
		r.logger.Error("Falha ao atualizar produto no banco", zap.Error(err))
		return models.Produto{}, fmt.Errorf("atualizar produto: %w", err)
	}

    r.logger.Info("Produto atualizado", zap.String("id", id.String()), zap.String("nome", nome), zap.Float64("preco", preco))
	return produto, nil
}

// Deletar remove um produto pelo ID.
func (r *PostgresRepositorio) Deletar(id uuid.UUID) error {
	result := r.db.Delete(&models.Produto{}, "id = ?", id)
	if result.Error != nil {
		r.logger.Error("Falha ao deletar produto no banco", zap.Error(result.Error))
		return fmt.Errorf("deletar produto: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		r.logger.Error("Falha ao deletar produto", zap.Error(ErrProdutoNaoEncontrado), zap.String("id", id.String()))
		return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

    r.logger.Info("Produto deletado", zap.String("id", id.String()))
	return nil
}
```

4. Atualize o arquivo `cmd/api/main.go` utilizando o código:

```go
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
```

5. Crie o arquivo docker-compose.yml utilizando o código:

```yaml
version: "3.8"

services:
  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      - POSTGRES_HOST=postgres
  postgres:
    image: postgres:latest
    environment:
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DB=mydb
    ports:
      - "5432:5432"
```

6. Execute com o comando:

```bash
docker-compose up --build
```

7. Teste e observabilidade

- Teste da API:

```bash
curl -X POST http://localhost:8080/produtos -H "Content-Type: application/json" -d '{"nome":"Laptop","preco":999.99}'

curl http://localhost:8080/produtos
```

- Para visualizar as métricas (prometheus), acesse a partir do browser o endereço [http://localhost:8080/metrics](http://localhost:8080/metrics)

- **Tracing**: Verifique a saída do `stdouttrace` nos logs do container.

- Para verificar a cobertura de testes, utilize os comandos:

```bash
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```

- Para verificar o código, utilize o comando:

```bash
golangci-lint run
```

- Para visualizar a documentação gerada, utilize o comando:

```bash
godoc -http=:6060
```

> Para visualizar a documentação, acesse a partir do browser o endereço [http://localhost:6060/pkg/github.com/seu-usuario/lab6/](http://localhost:6060/pkg/github.com/seu-usuario/lab6/)

</TabItem>
<TabItem value="task">

### Tarefa

- Configure o `OpenTelemetry` para exportar traces para `Jaeger` (ex: via Docker)
- Adicione `métricas customizadas` (ex: **contador de requisições** por endpoint)
- Implemente um `health check` endpoint (`/health`) com status do banco

#### Saída esperada

##### Console

###### POST

- /produtos

```json
{ "id": "", "nome": "Laptop", "preco": 999.99 }
```

###### GET

- /produtos

```json
[{ "id": "", "nome": "Laptop", "preco": 999.99 }]
```

##### Logs JSON (exemplo)

```json
{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Migrações aplicadas"}

{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}
```

:::info Caso de uso prático
Este lab simula um microsserviço em produção, com `containerização`, `observabilidade` e `práticas robustas`, pronto para deploy em _Kubernetes ou AWS_.
:::

</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu _build_, _containerização_, _logging estruturado_ (`zap`), _tracing_ (`OpenTelemetry`), _métricas_ (`Prometheus`), `linters`, _cobertura e documentação_ (`godoc`). O lab prático containerizou a API CRUD, integrando observabilidade para produção. Engenheiros Java notarão semelhanças com Spring Boot e ferramentas como Prometheus, mas com a simplicidade e eficiência de Go.

:::tip Próximos passos
Explore tópicos avançados, como `gRPC`, `integração com message brokers` (ex: `RabbitMQ`), ou _escalabilidade com Kubernetes_.
:::
