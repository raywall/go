---
sidebar_position: 9
sidebar_label: Módulo 08
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Web APIs com `net/http` e `Gin` em Go

<div className="row">
<div className="col">

Este módulo aborda a construção de `APIs RESTful` em Go, começando com o pacote padrão `net/http` e avançando para o framework `Gin`, que simplifica `roteamento`, `binding` e `validação`. O conteúdo é voltado para engenheiros Java, com comparações ao Spring Boot, e mantém-se objetivo com exemplos práticos e casos de uso para consulta futura.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default} 
    style={{ transform:'scale(1.3)', marginTop:'-3rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

O lab prático `implementa uma API RESTful` para o CRUD dos módulos anteriores, usando `Gin` com `validação de entrada`.

<br />

## Servidor HTTP com `net/http`

### Pacote `net/http`

- Fornece funcionalidades para `criar servidores HTTP` e `lidar com requisições/respostas`.
- **Simples, mas poderoso**, usado em muitos projetos Go sem frameworks.

#### Exemplo

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

    http.ListenAndServe(":8080", nil)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/api-http.go)!

#### Execução

```bash
go run main.go
```

> Acesse via browser a partir do endereço [http://localhost:8080/hello](http://localhost:8080/hello)

#### Comparação com Java

##### Java

- `Spring Boot` ou `Servlets` requerem mais configuração (ex: `@RestController`)

##### Go

- `net/http` é _mais leve, com menos abstrações_.

:::info Caso de uso
Servidores simples, como ferramentas internas ou APIs mínimas.
:::

<br />

## Middlewares e handlers

### Handlers

- Funções que **lidam com requisições HTTP**, com assinatura `func(w http.ResponseWriter, r *http.Request)`
- Podem ser registrados com `http.HandleFunc`

### Middlewares

- Funções que **interceptam requisições/respostas**, úteis para `autenticação`, `logging`, etc.
- Encadeiam `handlers`, retornando um `http.Handler`

#### Exemplo (middleware com logging)

```go
package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

    mux.Handle("/", loggingMiddleware(mux))
	http.ListenAndServe(":8080", nil)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/api-middleware.go)!

#### Comparação com Java

##### Java

- `Middlewares` são filtros no Spring (`@Filter`) ou `interceptores`

##### Go

- `Middlewares` são **mais explícitos**, usando composição de `handlers`

:::info Caso de uso
Middlewares para `logging`, `autenticação` (ex: JWT), ou `validação de cabeçalhos`.
:::

<br />

## Framework Gin: roteamento, binding, validação

### Gin

- Framework leve para APIs RESTful, com `roteamento rápido` e `suporte a JSON`, `binding` e `validação`.
- **Instalação**: `go get github.com/gin-gonic/gin`

#### Roteamento

- Usa `grupos de rotas` e `métodos HTTP` (`GET`, `POST`, `PUT`, `DELETE` etc.)
- Suporta `parâmetros de URL` e `query strings`.

##### Exemplo

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:nome", func(c *gin.Context) {
		nome := c.Param("nome")
		c.JSON(http.StatusOK, gin.H{"mensagem": "Olá, " + nome})
	})

    r.Run(":8080")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/api-gin-roteamento.go)!

#### Binding e validação

- Converte JSON ou formulários em `structs`, com `validação via tags`.
- Requer biblioteca `validator`: `go get github.com/go-playground/validator/v10`

##### Exemplo

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Produto struct {
	Nome  string  `json:"nome" binding:"required"`
	Preco float64 `json:"preco" binding:"required,gt=0"`
}

func main() {
	r := gin.Default()
	r.POST("/produtos", func(c *gin.Context) {
		var p Produto
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	})

    r.Run(":8080")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/api-gin-binding-validacao.go)!

##### Comparação com Java:

###### Java

- Spring Boot usa `@RestController`, `@RequestBody` e validação com `@Valid`.

###### Go

- Gin é **mais leve**, com `binding` e `validação` **diretamente via tags**.

:::info Caso de uso
Gin é ideal para `APIs RESTful escaláveis`, como sistemas de e-commerce ou microsserviços.
:::

<br />

## JSON, status codes e headers

### JSON

- Gin `serializa/deserializa JSON automaticamente` com `c.JSON`
- Usa `encoding/json` internamente

### Status codes

- Usa `constantes do net/http` (ex: `http.StatusOK`, `http.StatusBadRequest`)

### Headers

- Manipulados com `c.Header` ou `c.GetHeader`

### Exemplo

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.Header("X-Custom-Header", "valor")
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "API funcionando",
		})
	})

    r.Run(":8080")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/api-gin-json.go)!

### Teste

Utilize a ferramenta `curl`

```bash
curl -X GET http://localhost:8080/api
```

#### Resposta

```json
{ "mensagem": "API funcionando" }
```

:::info Caso de uso
`JSON` para respostas de APIs, `status codes` para indicar sucesso/erro, e `headers` para autenticação ou metadados.
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

### Configuração

1. Use a estrutura do [Módulo 06](go-module-6.md)

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
│       └── memoria_test.go
└── models/
    └── produto.go
```

2. Adicione as dependências necessárias:

```bash
go get github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
```

3. Atualize o arquivo `go.mod`:

```bash
module github.com/seu-usuario/lab8

go 1.21

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/uuid v1.6.0
)
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod8/lab/lab8.zip)!

</TabItem>
<TabItem value="app">

### Implementar uma API RESTful com Gin + validação

#### Objetivo

Implementar uma API RESTful para o CRUD do [Módulo 06](go-module-6.md), usando `Gin` para `roteamento`, `binding` e `validação`, integrando o repositório em memória e `logging estruturado`.

#### Passo a passo

1. Atualize o conteúdo do arquivo `models/produto.go` utilizando o código:

```go
import "github.com/google/uuid"

type Produto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome" binding:"required"`
	Preco float64   `json:"preco" binding:"required,gt=0"`
}
```

2. Atualize o conteúdo do arquivo `internal/repo/memoria.go` utilizando o código:

```go
package repo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
)

var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
)

type RepositorioProdutos interface {
	Criar(nome string, preco float64) (models.Produto, error)
	Buscar(id uuid.UUID) (models.Produto, error)
	Listar() ([]models.Produto, error)
	Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error)
	Deletar(id uuid.UUID) error
}

type RepositorioEmMemoria struct {
	produtos map[uuid.UUID]models.Produto
	logger   *slog.Logger
}

func NovoRepositorioEmMemoria(logger *slog.Logger) *RepositorioEmMemoria {
	return &RepositorioEmMemoria{
		produtos: make(map[uuid.UUID]models.Produto),
		logger:   logger,
	}
}

func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)

        return models.Produto{}, ErrPrecoInvalido
	}

    id := uuid.New()
	produto := models.Produto{ID: id, Nome: nome, Preco: preco}

    r.produtos[id] = produto
	r.logger.Info("Produto criado", "id", id, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id uuid.UUID) (models.Produto, error) {
	produto, existe := r.produtos[id]
	if !existe {
		r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)

        return models.Produto{}, fmt.Errorf("buscar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

    r.logger.Info("Produto encontrado", "id", id)
	return produto, nil
}

func (r *RepositorioEmMemoria) Listar() ([]models.Produto, error) {
	var produtos []models.Produto
	for _, p := range r.produtos {
		produtos = append(produtos, p)
	}

    r.logger.Info("Listando produtos", "total", len(produtos))
	return produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)

        return models.Produto{}, ErrPrecoInvalido
	}

    produto, existe := r.produtos[id]
	if !existe {
		r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)

        return models.Produto{}, fmt.Errorf("atualizar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

    produto.Nome = nome
	produto.Preco = preco

    r.produtos[id] = produto
	r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Deletar(id uuid.UUID) error {
	if _, existe := r.produtos[id]; !existe {
		r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)

        return fmt.Errorf("deletar produto id %s: %w", id, ErrProdutoNaoEncontrado)
	}

    delete(r.produtos, id)
	r.logger.Info("Produto deletado", "id", id)

    return nil
}
```

3. Atualize o conteúdo do arquivo `cmd/api/main.go` utilizando o código:

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/lab6/internal/repo"
	"github.com/seu-usuario/lab6/models"

	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NovoRepositorioEmMemoria(logger)

	r := gin.Default()

	// Middleware de logging
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
```

##### Testes da API

###### Execute o servidor

```bash
go run cmd/api/main.go
```

##### Teste com curl

- Criar um produto

```bash
curl -X POST http://localhost:8080/produtos -H "Content-Type: application/json" -d '{"nome":"Laptop","preco":999.99}'
```

- Listar produtos

```bash
curl http://localhost:8080/produtos
```

- Buscar um produto

```bash
curl http://localhost:8080/produtos/1
```

- Atualizar um produto

```bash
curl -X PUT http://localhost:8080/produtos/1 -H "Content-Type: application/json" -d '{"nome":"Laptop Pro","preco":1299.99}'
```

- Deletar um produto

```bash
curl -X DELETE http://localhost:8080/produtos/1
```

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione validação para garantir que o nome do produto tenha pelo _menos 3 caracteres_.
- Implemente um `middleware` para _autenticação simples_ (ex: verificar um `header Authorization`)
- Adicione _testes unitários para os endpoints_ usando `httptest` e `testify`.

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

- /produtos/1

```json
{ "id": "", "nome": "Laptop", "preco": 999.99 }
```

###### PUT

- /produtos/1

```json
{ "id": "", "nome": "Laptop Pro", "preco": 1299.99 }
```

###### DELETE

- /produtos/1

```bash
(status 204, sem corpo)
```

##### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}
```

:::info Caso de uso prático
Esta API simula um backend de e-commerce, com endpoints RESTful, validação de entrada e logging estruturado, ideal para integração com frontends ou outros serviços.
:::

</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu a criação de APIs RESTful com `net/http` e `Gin`, incluindo `roteamento`, `middlewares`, `binding`, `validação`, `JSON`, `status codes` e `headers`. O lab prático implementou uma API completa para o CRUD, integrando `validação` e `logging`. Engenheiros Java notarão semelhanças com Spring Boot, mas _com a abordagem mais leve e direta de Go_.

:::tip Próximos passos
No próximo módulo, _exploraremos integração com bancos de dados_ (ex: `PostgreSQL`) e _boas práticas para deploy de APIs_ em produção.
:::
