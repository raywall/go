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

1. Use a estrutura do [Módulo 06](go-basic/go-module-6/5-laboratorio.md)

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

## Implementar uma API RESTful com Gin + validação

### Objetivo

Implementar uma API RESTful para o CRUD do [Módulo 06](go-basic/go-module-6/5-laboratorio.md), usando `Gin` para `roteamento`, `binding` e `validação`, integrando o repositório em memória e `logging estruturado`.

### Passo a passo

1. Atualize o conteúdo do arquivo `models/produto.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod8/lab/models/produto.go" 
    allowExecute={false} 
    allowEdit={false} />

2. Atualize o conteúdo do arquivo `internal/repo/memoria.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod8/lab/internal/repo/memoria.go" 
    allowExecute={false} 
    allowEdit={false} />

3. Atualize o conteúdo do arquivo `cmd/api/main.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod8/lab/cmd/api/main.go" 
    allowExecute={false} 
    allowEdit={false} />

#### Testes da API

##### Execute o servidor

```bash
go run cmd/api/main.go
```

#### Teste com curl

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

## Tarefa

- Adicione validação para garantir que o nome do produto tenha pelo _menos 3 caracteres_.
- Implemente um `middleware` para _autenticação simples_ (ex: verificar um `header Authorization`)
- Adicione _testes unitários para os endpoints_ usando `httptest` e `testify`.

### Saída esperada

#### Console

##### POST

- /produtos

```json
{ "id": "", "nome": "Laptop", "preco": 999.99 }
```

##### GET

- /produtos

```json
[{ "id": "", "nome": "Laptop", "preco": 999.99 }]
```

- /produtos/1

```json
{ "id": "", "nome": "Laptop", "preco": 999.99 }
```

##### PUT

- /produtos/1

```json
{ "id": "", "nome": "Laptop Pro", "preco": 1299.99 }
```

##### DELETE

- /produtos/1

```bash
(status 204, sem corpo)
```

#### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}
```

:::info Caso de uso prático
Esta API simula um backend de e-commerce, com endpoints RESTful, validação de entrada e logging estruturado, ideal para integração com frontends ou outros serviços.
:::

</TabItem>
</Tabs>
