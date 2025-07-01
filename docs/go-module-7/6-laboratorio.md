---
sidebar_position: 7
sidebar_label: Laboratório
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

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

1. Use a estrutura do [Módulo 06](go-module-6/5-laboratorio.md)

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

## Criar testes unitários e de integração para o CRUD com cobertura de erro

### Objetivo

Implementar testes unitários e de integração para o CRUD do [Módulo 06](go-module-6/5-laboratorio.md), usando `testing` e `testify`, com foco em cobertura de erros. Gerar relatórios de cobertura e usar `staticcheck` para análise.

### Passo a passo

1. Crie o arquivo `internal/repo/memoria_test.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod7/lab/internal/repo/memoria_test.go" 
    allowExecute={false} 
    allowEdit={false} />

2. Crie o arquivo `cmd/api/main_test.go` (teste de integração) utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod7/lab/internal/repo/memoria.go" 
    allowExecute={false} 
    allowEdit={false} />

#### Execução

##### Execute os testes

- `go test ./... -v`

##### Gere relatório de cobertura

- `go test ./... -coverprofile=cover.out`
- `go tool cover -html=cover.out`

##### Analise com staticcheck

- `staticcheck ./...`

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione um `benchmark` para a função Listar do repositório.
- Crie um `mock` com `testify/mock` para testar um serviço que usa `RepositorioProdutos`.
- Use `go vet` e `golangci-lint` para verificar o código.

### Saída esperada (testes)

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
