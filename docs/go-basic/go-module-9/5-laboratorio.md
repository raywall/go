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

## Pré-requisitos

- `PostgreSQL` instalado e rodando:

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres
```

- Banco `mydb` criado:

```bash
psql -U postgres -c "CREATE DATABASE mydb;"
```

## Configuração

1. Use a estrutura do [Módulo 08](go-basic/go-module-8/5-laboratorio.md)

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
└── migrations/
 	├── 202506120001_create_produtos.up.sql
    └── 202506120001_create_produtos.down.sql
```

2. Adicione as dependências necessárias:

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-migrate/migrate/v4
```

3. Atualize o arquivo `go.mod`:

```bash
module github.com/seu-usuario/lab6

go 1.21

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/uuid v1.6.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
	github.com/golang-migrate/migrate/v4 v4.17.1
)
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod9/lab/lab9.zip)!

</TabItem>
<TabItem value="app">

## Persistir o CRUD em Banco Real (PostgreSQL)

### Objetivo

Refatorar o CRUD do [Módulo 08](go-basic/go-module-8/5-laboratorio.md) para usar `PostgreSQL` com `GORM`, aplicar migrações com `golang-migrate`, e _implementar testes de integração_.

### Passo a passo

1. Crie o arquivo `migrations/202506120001_create_produtos.up.sql` com o conteúdo:

```sql
CREATE TABLE produtos (
    id UUID PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    preco DOUBLE PRECISION NOT NULL
);
```

2. Crie o arquivo `migrations/202506120001_create_produtos.down.sql` com o conteúdo:

```sql
DROP TABLE produtos;
```

3. Atualize o arquivo `models/produto.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod9/lab/models/produto.go" 
    allowExecute={false} 
    allowEdit={false} />

4. Crie o arquivo `internal/repo/postgres.go` com o conteúdo:

<InteractiveCodeSnippet 
    src="code/mod9/lab/internal/repo/postgres.go" 
    allowExecute={false} 
    allowEdit={false} />

5. Atualize o arquivo `cmd/api/main.go` com o conteúdo:

<InteractiveCodeSnippet 
    src="code/mod9/lab/cmd/api/main.go" 
    allowExecute={false} 
    allowEdit={false} />

6. Para realizar os testes de integração, crie o arquivo `internal/repo/postgres_test.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod9/lab/internal/repo/postgres_test.go" 
    allowExecute={false} 
    allowEdit={false} />

### Execução

- Inicie o `PostgreSQL` (ex: via Docker):

```bash
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres
psql -U postgres -c "CREATE DATABASE mydb;"
```

- Aplique migrações e execute a API:

```bash
   go run cmd/api/main.go
```

- Teste com o comando curl:

```bash
   curl -X POST http://localhost:8080/produtos -H "Content-Type: application/json" -d '{"nome":"Laptop","preco":999.99}'

   curl http://localhost:8080/produtos
```

- Execute testes de integração com o comando:

```bash
   go test ./internal/repo -v
```

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione validação no `GORM` para garantir que o nome tenha no _mínimo 3 caracteres_.
- Implemente um _teste de integração para os endpoints da API_ usando `httptest`.
- Use `golang-migrate` para _adicionar uma nova coluna_ (ex: `created_at`) à tabela produtos.

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

#### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Migrações aplicadas"}

{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"time":"2025-06-12T01:20:00Z","level":"INFO","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}
```

:::info Caso de uso prático
Este lab _simula um backend de e-commerce com persistência em `PostgreSQL`_, _migrações automatizadas_ e _testes de integração_, ideal para aplicações em produção.
:::

</TabItem>
</Tabs>
