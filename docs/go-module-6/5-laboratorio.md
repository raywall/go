---
sidebar_position: 6
sidebar_label: Laboratório
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Laboratório

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c', marginTop:'2rem', marginBottom:'5rem' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

## Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab6`

```bash
mkdir lab6
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab6
go mod init github.com/seu-usuario/lab6
```

4. Crie a seguinte estrutura para o projeto:

```dirtree
lab6/
├── go.mod
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── repo/
│       └── memoria.go
└── models/
    └── produto.go
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod6/lab/lab6.zip)!

</TabItem>
<TabItem value="app">

## Organizar o projeto CRUD em múltiplos pacotes com `go mod`

### Objetivo

Reorganizar o CRUD dos módulos anteriores em uma estrutura de pacotes idiomática, usando go mod para gerenciar dependências e adicionar uma biblioteca externa (github.com/google/uuid) para gerar IDs.

### Passo a passo

1. Crie o arquivo `models/produto.go` com o seguinte código:

```go
package models

import "github.com/google/uuid"

// Produto representa um produto no sistema
type Produto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Preco float64   `json:"preco"`
}

```

2. Crie o arquivo `internal/repo/memoria.go` com o seguinte código:

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

// RepositorioProdutos define a interface do repositório
type RepositorioProdutos interface {
	Criar(nome string, preco float64) (models.Produto, error)
	Buscar(id uuid.UUID) (models.Produto, error)
	Listar() ([]models.Produto, error)
	Atualizar(id uuid.UUID, nome string, preco float64) (models.Produto, error)
	Deletar(id uuid.UUID) error
}

// RepositorioEmMemoria implementa o repositório em memória
type RepositorioEmMemoria struct {
	produtos map[uuid.UUID]models.Produto
	logger   *slog.Logger
}

// NovoRepositorioEmMemoria cria um novo repositório
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

3. Crie o arquivo `cmd/api/main.go` com o seguinte código:

```go
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/seu-usuario/lab6/internal/repo"
)

func exibirProdutos(repo repo.RepositorioProdutos) error {
	produtos, err := repo.Listar()
	if err != nil {
		return fmt.Errorf("exibir produtos: %w", err)
	}

    fmt.Println("Lista de produtos:")
	for _, p := range produtos {
		fmt.Printf("ID: %s, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NovoRepositorioEmMemoria(logger)

    // Criar produtos
	p1, err := repo.Criar("Laptop", 999.99)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	repo.Criar("Mouse", 29.99)

    // Listar produtos
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
		return
	}

    // Buscar produto
	if p, err := repo.Buscar(p1.ID); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

    // Atualizar produto
	if p, err := repo.Atualizar(p1.ID, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

    // Deletar produto
	if err := repo.Deletar(p1.ID); err == nil {
		fmt.Println("Produto deletado com sucesso")
	} else {
		fmt.Println("Erro:", err)
	}

    // Listar novamente
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
	}
}
```

### Execução

```bash
go mod tidy
go run cmd/api/main.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione um pacote api com uma função que simule um endpoint HTTP (ex.: `ListarProdutos` retornando JSON).
- Use replace no `go.mod` para testar uma versão local da biblioteca `github.com/google/uuid`.
- Crie um pacote utils com uma função para validar preços (ex.: `ValidarPreco`).

### Saída esperada

#### Console:

```bash
Lista de produtos:
ID: 1, Nome: Laptop, Preço: 999.99
ID: 2, Nome: Mouse, Preço: 29.99

Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
Produto deletado com sucesso

Lista de produtos:
ID:2 , Nome: Mouse, Preço: 29.99
```

#### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Produto criado","id":"1","nome":"Laptop","preco":999.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Produto criado","id":"2","nome":"Mouse","preco":29.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Listando produtos","total":2}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Produto encontrado","id":"1"}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Produto atualizado","id":"1","nome":"Laptop Pro","preco":1299.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Produto deletado","id":"1"}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Listando produtos","total":1}
```

:::info Caso de uso prático
Esta estrutura é típica de projetos Go reais, como `APIs RESTful`, onde pacotes separam lógica de negócio, acesso a dados e entrada/saída.
:::
</TabItem>
</Tabs>
