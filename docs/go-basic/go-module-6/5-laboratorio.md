---
sidebar_position: 6
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

<InteractiveCodeSnippet 
    src="code/mod10/lab/models/produto.go" 
    allowExecute={false} 
    allowEdit={false} />

2. Crie o arquivo `internal/repo/memoria.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod10/lab/internal/repo/memoria.go" 
    allowExecute={false} 
    allowEdit={false} />

3. Crie o arquivo `cmd/api/main.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod10/lab/cmd/api/main.go" 
    allowExecute={false} 
    allowEdit={false} />

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
