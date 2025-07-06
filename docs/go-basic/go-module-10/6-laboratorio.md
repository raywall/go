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

1. Use a estrutura do [Módulo 09](go-module-9/5-laboratorio.md)

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

## Containerizar o serviço e expor métricas, trace e logs

### Objetivo

Containerizar a API CRUD do [Módulo 09](go-module-9/index.md) (`PostgreSQL`), configurar _logging estruturado_ com `zap`, tracing com `OpenTelemetry`, e expor métricas com `Prometheus`. Gerar cobertura de testes e documentação com `godoc`.

### Passo a passo

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

<InteractiveCodeSnippet 
    src="code/mod10/lab/models/produto.go" 
    allowExecute={false} 
    allowEdit={false} />

3. Atualize o arquivo `internal/repo/postgres.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod10/lab/internal/repo/postgres.go" 
    allowExecute={false} 
    allowEdit={false} />

4. Atualize o arquivo `cmd/api/main.go` utilizando o código:

<InteractiveCodeSnippet 
    src="code/mod10/lab/cmd/api/main.go" 
    allowExecute={false} 
    allowEdit={false} />

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

## Tarefa

- Configure o `OpenTelemetry` para exportar traces para `Jaeger` (ex: via Docker)
- Adicione `métricas customizadas` (ex: **contador de requisições** por endpoint)
- Implemente um `health check` endpoint (`/health`) com status do banco

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
{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Migrações aplicadas"}

{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Produto criado","id":"","nome":"Laptop","preco":999.99}

{"level":"info","ts":"2025-06-12T01:25:00Z","msg":"Requisição processada","method":"POST","path":"/produtos","status":201,"duration":"1ms"}
```

:::info Caso de uso prático
Este lab simula um microsserviço em produção, com `containerização`, `observabilidade` e `práticas robustas`, pronto para deploy em _Kubernetes ou AWS_.
:::

</TabItem>
</Tabs>
