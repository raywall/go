---
sidebar_position: 7
sidebar_label: Módulo 06
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Pacotes, módulos e organização do código em Go

<div className="row">
<div className="col">

Este módulo explora a organização de código em Go, incluindo a estrutura idiomática de pacotes, convenções de projeto, gerenciamento de dependências com go mod, e boas práticas para engenheiros Java que estão aprendendo Go. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura.

O lab prático reorganiza o CRUD dos módulos anteriores em múltiplos pacotes, utilizando go mod para gerenciar dependências.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default} 
    style={{ transform:'scale(0.65)', marginTop:'-65px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Estrutura de pacotes idiomática

### O que são pacotes?

- Pacotes são `unidades de organização de código` em Go, semelhantes a pacotes em Java (`package`).
- Cada arquivo Go pertence a um pacote, declarado com `package` nome.
- O pacote main gera um executável; outros pacotes são bibliotecas.

### Estrutura idiomática

- Go recomenda uma organização clara para projetos:

- Raiz do projeto: Contém go.mod e subdiretórios.
- Pacotes: Organizados em subdiretórios com nomes descritivos (ex.: api, models, repo).
- Nomenclatura: Nomes de pacotes são curtos, em minúsculas, sem sublinhados (ex.: http, db).
- Exportação: Funções, tipos e variáveis com inicial maiúscula são exportados (ex.: func Criar é visível fora do pacote).

### Exemplo

```dirtree
meu-projeto/
├── go.mod
├── main.go
├── models/
│   └── produto.go
└── repo/
    └── memoria.go
```

### Comparação com Java

#### Java

- Pacotes seguem a convenção `com.empresa.projeto`

#### Go

- Pacotes usam caminhos baseados no módulo (ex: `github.com/usuario/projeto/models`)

:::info Caso de uso
Pacotes são usados para modularizar código em microsserviços, separando lógica de negócio, acesso a dados e APIs.
:::

<br />

## Convenções de projeto (`cmd`, `internal` e `pkg`)

### Convenções comuns

#### `cmd/`

- Contém pontos de entrada (`main`) para executáveis.

##### Exemplos

- `cmd/api/main.go` para uma API
- `cmd/cli/main.go` para uma ferramenta CLI

#### `internal/`

- Utilizado para pacotes privados, acessíveis apenas dentro do projeto ou subdiretórios.
- Ideal para lógica sensível ou reutilização interna.

#### `pkg/`

- Pacotes reutilizáveis por outros projetos (opcional, menos comum).

#### Outros diretórios

- Nomes descritivos como: api, db, models, utils.

### Exemplo de estrutura

```dirtree
projeto/
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

### Comparação com Java

#### Java

- Estrutura rígida com `src/main/java` e **pacotes hierárquicos**.

#### Go

- Mais flexível, mas `internal` é semelhante a `package-private` em Java.

:::info Caso de uso
Usar internal para proteger implementações de repositórios, enquanto cmd organiza múltiplos executáveis (ex: servidor e ferramenta de migração).
:::

<br />

## `go mod` e versionamento

### Go Modules

- Introduzidos no Go 1.11, substituem `$GOPATH` para gerenciar dependências.
- O arquivo `go.mod` define o módulo e suas dependências.

#### Exemplo de go.mod

```text
module github.com/seu-usuario/meu-projeto

go 1.21

require (
    github.com/google/uuid v1.6.0
)
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod6/go.mod)!

#### Comandos Principais

- `go mod init`: Inicializa um módulo.
- `go mod tidy`: Adiciona dependências usadas e remove não usadas.
- `go mod vendor` (opcional): Cria um diretório vendor/ com dependências.

### Versionamento

- Go usa `versionamento semântico` (ex: v1.6.0)
- Módulos são referenciados por URLs de repositórios (ex: `github.com/autor/projeto`)
- Tags no repositório (ex: `git tag v1.0.0`) definem versões

#### Comparação com Java

##### Java

- Usa `Maven/Gradle` com repositórios centralizados (ex: `Maven Central`)

##### Go

- Usa URLs de repositórios diretamente, com proxy (ex: `proxy.golang.org`)

:::info Caso de uso
Go Modules simplificam o gerenciamento de dependências em projetos grandes, como APIs que integram bibliotecas externas.
:::

<br />

## Gerenciamento de dependências com `go get` e `replace`

### `go get`

- Adiciona ou atualiza dependências no `go.mod`
  Exemplo: `go get github.com/google/uuid@v1.6.0`

- Atualiza para a versão mais recente:
  `go get -u github.com/google/uuid`

### `replace`

- Permite substituir uma dependência por outra versão ou caminho local.
- Útil para desenvolvimento `local` ou `forks`

### Exemplo de `go.mod` com `replace`

```text
module github.com/seu-usuario/meu-projeto

go 1.21

require github.com/google/uuid v1.6.0
replace github.com/google/uuid => ../uuid-fork
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod6/go-replace.mod)!

### Comparação com Java

#### Java

- Usa no `Maven` ou `implementation` no Gradle

#### Go

- `go get` e `replace` são mais simples, mas menos configuráveis

:::info Caso de uso
`replace` é útil para testar alterações locais em dependências antes de publicar
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

### Organizar o projeto CRUD em múltiplos pacotes com `go mod`

#### Objetivo

Reorganizar o CRUD dos módulos anteriores em uma estrutura de pacotes idiomática, usando go mod para gerenciar dependências e adicionar uma biblioteca externa (github.com/google/uuid) para gerar IDs.

#### Passo a passo

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

#### Execução

```bash
go mod tidy
go run cmd/api/main.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione um pacote api com uma função que simule um endpoint HTTP (ex.: `ListarProdutos` retornando JSON).
- Use replace no `go.mod` para testar uma versão local da biblioteca `github.com/google/uuid`.
- Crie um pacote utils com uma função para validar preços (ex.: `ValidarPreco`).

#### Saída esperada

##### Console:

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

##### Logs JSON (exemplo)

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
<br />

---

## Conclusão

Este módulo cobriu a organização de pacotes, convenções de projeto (`cmd`, `internal`, `pkg`), gerenciamento de dependências com `go mod`, e o uso de `go get` e `replace`. O lab prático reorganizou o CRUD em uma estrutura idiomática, integrando uma dependência externa. Engenheiros Java notarão semelhanças com a organização de pacotes Maven/Gradle, mas com a abordagem mais simples e direta de Go.

:::tip Próximos passos
No próximo módulo, exploraremos `testes unitários`, `benchmarks` e `integração com bibliotecas externas`, consolidando as práticas para aplicações robustas e escaláveis em Go.
:::
