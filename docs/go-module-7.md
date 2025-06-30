---
sidebar_position: 8
sidebar_label: Módulo 07
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Testes e Qualidade de Código em Go

<div className="row">
<div className="col">

Este módulo aborda testes e qualidade de código em Go, incluindo `testes unitários` e `testes de integração` com o pacote `testing`, uso de `mocks` com bibliotecas como `testify`, `benchmarks`, `profiling` e ferramentas de `análise estática`. O conteúdo é voltado para engenheiros Java, comparando com práticas como `JUnit` e `Mockito`, e mantendo-se objetivo para consulta futura.

</div>
<div className="col col--4 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default} 
    style={{ transform:'scale(0.8)', marginTop:'-1.1rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

O lab prático implementa testes unitários e de integração para o CRUD dos módulos anteriores, com foco em `cobertura de erros`.

<br />

## Testes com o pacote `testing`

### Pacote `testing`

- O pacote padrão testing fornece suporte para `testes unitários` e `benchmarks`.
- Testes são escritos em arquivos com sufixo `_test.go` e funções prefixadas com Test.
- Usa `t.Error` ou `t.Fatal` para reportar falhas.

#### Exemplo

```go
package main

import "testing"

func soma(a, b int) int {
	return a + b
}

func TestSoma(t *testing.T) {
	resultado := soma(2, 3)
	if resultado != 5 {
		t.Errorf("soma(2, 3) = %d; esperado 5", resultado)
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod7/testing.go)!

#### Execução

```bash
go test
```

:::tip Dica
Para executar o teste em todos os diretórios e subdiretórios do projeto utilize `go test ./...`
:::

#### Comparação com Java

##### Java

- Usa `JUnit/TestNG` com anotações (`@Test`).

##### Go

- Mais `simples`, `sem anotações`, e com `convenções baseadas em nomes`.

:::info Caso de uso
Testes unitários com `testing` são ideais para validar funções puras ou componentes isolados.
:::

<br />

## Testes de unidade e de integração

### Testes de unidade

- Testam `unidades isoladas` (ex: uma função ou método)
- Usam tabelas de teste para múltiplos cenários

#### Exemplo (tabela de testes)

```go
package main

import "testing"

func TestSoma(t *testing.T) {
	testes := []struct {
		nome     string
		a, b     int
		esperado int
	}{
		{"positivo", 2, 3, 5},
		{"negativo", -1, -1, -2},
		{"zero", 0, 5, 5},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := soma(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("soma(%d, %d) = %d; esperado %d", tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod7/teste-unitario.go)!

### Testes de integração

- Testam a `interação entre componentes` (ex: repositório e API)
- Podem usar `bancos de dados em memória` ou `mocks`

:::info Caso de uso
Testes de unidade `verificam lógica de negócio`; já os testes de integração `garantem que módulos funcionem juntos` (ex: repositório e serviço)
:::

<br />

## Testes com `mocks` (`testify` e `gomock`)

### Pacote `testify`

- Simplifica `asserções` e `criação de mocks`
- Inclui `assert` e `mock` para facilitar testes

#### Instalação

```bash
go get github.com/stretchr/testify
```

#### Exemplo com testify/mock

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

type RepositorioMock struct {
	mock.Mock
}

func (m *RepositorioMock) Buscar(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestServico(t *testing.T) {
	repo := &RepositorioMock{}
	repo.On("Buscar", 1).Return("Produto", nil)

    resultado, err := repo.Buscar(1)
	if err != nil {
		t.Fatal(err)
	}

    if resultado != "Produto" {
		t.Errorf("esperado Produto, recebido %s", resultado)
	}
	repo.AssertExpectations(t)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod7/teste-mock.go)!

### Pacote `gomock`

- **Gera mocks automaticamente** a partir de `interfaces`
- **Mais poderoso** para projetos complexos

#### Instalação

```bash
go install github.com/golang/mock/gomock@latest
go install github.com/golang/mock/mockgen@latest
```

:::info Caso de uso
`Mocks` são úteis para testar _serviços que dependem de repositórios_ ou `APIs externas`, similar ao `Mockito` em Java.
:::

<br />

## Benchmarks e profiling

### Benchmarks

- Funções prefixadas com `Benchmark` **medem desempenho**
- Usam `b.N` para _executar o código várias vezes_

#### Exemplo

```go
package main

import "testing"

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(2, 3)
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod7/benchmarks.go)!

#### Execução

```bash
go test -bench=.
```

### Profiling

Go oferece ferramentas integradas para análise de desempenho:

- `pprof`: Perfil de CPU, memória e bloqueios
- `go test -cpuprofile cpu.out`: Gera perfil de CPU
- `go tool pprof cpu.out`: Analisa o perfil

#### Exemplo

```bash
go test -bench=. -cpuprofile cpu.out
go tool pprof cpu.out
```

### Comparação com Java

#### Java

- Usa `VisualVM` ou `JProfiler` para profiling

#### Go

- Ferramentas embutidas (`pprof`) são mais leves e integradas

:::info Caso de uso
`Benchmarks` ajudam a otimizar funções críticas; enquanto o `profiling` **identifica gargalos** em _APIs_ ou _sistemas concorrentes_.
:::

<br />

## Ferramentas: `go vet`, `golint`, `staticcheck`

### `go vet`

- Analisa código para erros comuns (ex: formatação de strings incorreta)
  **Executado com**: `go vet ./...`

### `golint` (obsoleto, mas ainda usado)

- Sugere melhorias de estilo (ex: nomes de variáveis)
  **Alternativa moderna**: `golangci-lint`

### `staticcheck`

- Ferramenta avançada de análise estática, verificando bugs, desempenho e estilo.
  **Instalação**: `go install honnef.co/go/tools/cmd/staticcheck@latest`
  **Uso**: `staticcheck ./...`

### Comparação com Java

#### Java

- Usa `Checkstyle`, `PMD` ou `SonarQube`

#### Go

- Ferramentas são mais `integradas` e `leves`, com `foco em simplicidade`.

:::info Caso de uso
Essas ferramentas garantem consistência e qualidade em projetos grandes, como microsserviços.
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

2. Adicione o pacote `testify`:

```bash
go get github.com/stretchr/testigy
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod7/lab/lab7.zip)!

</TabItem>
<TabItem value="app">

### Criar testes unitários e de integração para o CRUD com cobertura de erro

#### Objetivo

Implementar testes unitários e de integração para o CRUD do [Módulo 06](go-module-6.md), usando `testing` e `testify`, com foco em cobertura de erros. Gerar relatórios de cobertura e usar `staticcheck` para análise.

#### Passo a passo

1. Crie o arquivo `internal/repo/memoria_test.go` utilizando o código:

```go
import (
	"github.com/google/uuid"
	"github.com/meu-usuario/labлог6/models"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestRepositorioEmMemoria(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NovoRepositorioEmMemoria(logger)

	t.Run("Criar produto com sucesso", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, produto.ID)
		assert.Equal(t, "Laptop", produto.Nome)
		assert.Equal(t, 999.99, produto.Preco)
	})

	t.Run("Criar produto com preço inválido", func(t *testing.T) {
		_, err := repo.Criar("Laptop", 1)
		assert.ErrorIs(t, err, ErrPrecoInvalido)
	})

	t.Run("Buscar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Mouse", 29.99)
		assert.NoError(t, err)

		encontrado, err := repo.Buscar(produto.ID)
		assert.NoError(t, err)
		assert.Equal(t, produto, encontrado)
	})

	t.Run("Buscar produto inexistente", func(t *testing.T) {
		_, err := repo.Buscar(uuid.New())
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})

	t.Run("Listar produtos", func(t *testing.T) {
		repo = NovoRepositorioEmMemoria(logger)
		repo.Criar("Laptop", 999.99)
		repo.Criar("Mouse", 29.99)

		produtos, err := repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 2)
	})

	t.Run("Atualizar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		atualizado, err := repo.Atualizar(produto.ID, "Laptop Pro", 1299.99)
		assert.NoError(t, err)
		assert.Equal(t, "Laptop Pro", atualizado.Nome)
		assert.Equal(t, 1299.99, atualizado.Preco)
	})

	t.Run("Atualizar produto com preço inválido", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		_, err = repo.Atualizar(produto.ID, "Laptop Pro", 1)
		assert.ErrorIs(t, err, ErrPrecoInvalido)
	})

	t.Run("Deletar produto existente", func(t *testing.T) {
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		err = repo.Deletar(produto.ID)
		assert.NoError(t, err)

		_, err = repo.Buscar(produto.ID)
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})

	t.Run("Deletar produto inexistente", func(t *testing.T) {
		err := repo.Deletar(uuid.New())
		assert.ErrorIs(t, err, ErrProdutoNaoEncontrado)
	})
}
```

2. Crie o arquivo `cmd/api/main_test.go` (teste de integração) utilizando o código:

```go
package main

import (
	"github.com/meu-usuario/lab6/internal/repo"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NovoRepositorioEmMemoria(logger)

	t.Run("Fluxo completo do CRUD", func(t *testing.T) {

		// Criar
		produto, err := repo.Criar("Laptop", 999.99)
		assert.NoError(t, err)

		// Listar
		produtos, err := repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 1)

		// Buscar
		encontrado, err := repo.Buscar(produto.ID)
		assert.NoError(t, err)
		assert.Equal(t, produto, encontrado)

		// Atualizar
		atualizado, err := repo.Atualizar(produto.ID, "Laptop Pro", 1299.99)
		assert.NoError(t,, err)
		assert.Equal(t, "Laptop Pro", atualizado.Nome)

		// Deletar
		err = repo.Deletar(produto.ID)
		assert.NoError(t, err)

		// Verificar exclusão
		produtos, err = repo.Listar()
		assert.NoError(t, err)
		assert.Len(t, produtos, 0)
	})
}
```

##### Execução

###### Execute os testes

- `go test ./... -v`

###### Gere relatório de cobertura

- `go test ./... -coverprofile=cover.out`
- `go tool cover -html=cover.out`

###### Analise com staticcheck

- `staticcheck ./...`

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione um `benchmark` para a função Listar do repositório.
- Crie um `mock` com `testify/mock` para testar um serviço que usa `RepositorioProdutos`.
- Use `go vet` e `golangci-lint` para verificar o código.

#### Saída esperada (testes)

```bash
=== RUN   TestRepositorioEmMemoria
=== RUN   TestRepositorioEmMemoria/Criar_produto_com_sucesso
=== RUN   TestRepositorioEmMemoria/Criar_produto_com_preço_inválido
=== RUN   TestRepositorioEmMemoria/Buscar_produto_existente
=== RUN   TestRepositorioEmMemoria/Buscar_produto_inexistente
=== RUN   TestRepositorioEmMemoria/Listar_produtos
=== RUN   TestRepositorioEmMemoria/Atualizar_produto_existente
=== RUN   TestRepositorioEmMemoria/Atualizar_produto_com_preço_inválido
=== RUN   TestRepositorioEmMemoria/Deletar_produto_existente
=== RUN   TestRepositorioEmMemoria/Deletar_produto_inexistente
--- PASS: TestRepositorioEmMemoria (0.00s)
--- PASS: TestRepositorioEmMemoria/Criar_produto_com_sucesso (0.00s)
--- PASS: TestRepositorioEmMemoria/Criar_produto_com_preço_inválido (0.00s)
--- PASS: TestRepositorioEmMemoria/Buscar_produto_existentesignalize
--- PASS: TestRepositorioEmMemoria/BuscarSweeten
--- PASS: TestRepositorioEmMemoria/Listar_produtos (0.00s)
--- PASS: TestRepositorioEmMemoria/Atualizar_produto_existente (0.00s)
--- PASS: TestRepositorioEmMemoria/Atualizar_produto_com_preço_inválido (0.00s)
--- PASS: TestRepositorioEmMemoria/Deletar_produto_existente (0.00s)
--- PASS: TestRepositorioEmMemoria/Deletar_produto_inexistente (0.00s)
=== RUN   TestMainIntegration
=== RUN   TestMainIntegration/Fluxo_completo_do_CRUD
--- PASS: TestMainIntegration (0.00s)
--- PASS: TestMainIntegration/Fluxo_completo_do_CRUD (0.00s)
PASS
ok      github.com/seu-usuario/lab6/... 0.012s
```

:::info Caso de uso prático
**Testes unitários e de integração garantem a robustez** do CRUD, enquanto a `cobertura de erros previne falhas` em cenários como `preços inválidos` ou `IDs inexistentes`.
:::

</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu testes com `testing`, `testes de unidade` e `testes de integração`, `mocks` com `testify`, `benchmarks`, `profiling` e ferramentas de análise estática (`go vet` e `staticcheck`). O lab prático implementou testes completos para o CRUD, com cobertura de erros e logging estruturado. Engenheiros Java notarão semelhanças com `JUnit` e `Mockito`, mas com a simplicidade e integração nativa do ecossistema Go.

:::tip Próximos passos
No próximo módulo, exploraremos a construção de `APIs RESTful` com bibliotecas como `gin` ou `echo`, integrando o CRUD em um servidor web real.
:::
