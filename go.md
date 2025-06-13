

🧭 Estrutura do Treinamento

- Duração sugerida: 5 dias (4h por dia) ou 10 sessões (2h cada)
- Formato: Teórico + Mão na massa (cada módulo tem lab)
- Pré-requisitos: Conhecimento prévio em backend com Java/Python, Git, HTTP/REST, e banco de dados relacional.


📚 Conteúdo Programático


Módulo 1 – Introdução e Fundamentos da Linguagem

- História e propósito do Go
- Características da linguagem (simplicidade, performance, concorrência)
- Instalação, workspace, go mod
- Estrutura básica de um programa
- Tipos primitivos, funções, variáveis, constantes

📌 Lab: Hello world, manipulação de variáveis e tipos, primeiro programa com go run.

---

Módulo 2 – Estruturas de Controle e Tipos Compostos

- if, for, switch, defer
- Arrays, slices, maps
- Structs e tags de struct
- Ponteiros (conceito e aplicação)

📌 Lab: Implementar uma API fake para CRUD em memória usando slices e maps.

---

Módulo 3 – Funções, Métodos e Interfaces

- Funções com múltiplos retornos
- Funções anônimas e closures
- Métodos em structs
- Interfaces e duck typing
- Boas práticas e princípios de design idiomático em Go

📌 Lab: Refatorar o CRUD usando interfaces para repositórios.

---

Módulo 4 – Tratamento de Erros

- Filosofia do Go: erros explícitos
- Padrão error, errors.New, fmt.Errorf
- Wrapping e unwrapping com errors.Is, errors.As
- Pacote log e log/slog

📌 Lab: Criar funções com tratamento de erros e logging estruturado.

---

Módulo 5 – Concorrência com Goroutines e Channels

- Goroutines: o que são e como usar
- Channels (unbuffered, buffered)
- select, sync.WaitGroup, sync.Mutex
- Padrões de concorrência em Go

📌 Lab: Criar um worker pool para processamento concorrente de tarefas.

---

Módulo 6 – Pacotes, Módulos e Organização do Código

- Estrutura de pacotes idiomática
- Convenções de projeto (cmd, internal, pkg)
- go mod e versionamento
- Gerenciamento de dependências com go get, replace

📌 Lab: Organizar o projeto CRUD em múltiplos pacotes com go mod.

---

Módulo 7 – Testes e Qualidade de Código

- Testes com testing
- Testes de unidade e integração
- Testes com mocks (testify, gomock)
- Benchmarks e profiling
- Ferramentas: go vet, golint, staticcheck

📌 Lab: Criar testes unitários e de integração para o CRUD com cobertura de erro.

---

Módulo 8 – Web APIs com net/http e Gin

- Servidor HTTP com net/http
- Middlewares e handlers
- Framework Gin: roteamento, binding, validação
- JSON, status codes e headers

📌 Lab: Implementar uma API RESTful com Gin + validação.

---

Módulo 9 – Persistência com Banco de Dados

- Drivers e database/sql
- ORM com gorm
- Migrations com golang-migrate
- Repositórios e testes de integração com DB

📌 Lab: Persistir o CRUD em banco real (PostgreSQL por exemplo).

---

Módulo 10 – Deploy, Observabilidade e Boas Práticas

- Build com go build, cross-compilation
- Docker com Go
- Logging estruturado (slog, zap)
- Tracing com OpenTelemetry
- Linter, cobertura, documentação automática com godoc

📌 Lab: Containerizar o serviço e expor métricas/trace/logs.

---

🎯 Objetivo Final (Desafio de Encerramento)

Construir em grupo uma API de pedidos de e-commerce com:

- Autenticação
- Validação
- Persistência
- Concorrência para envio de e-mails assíncronos
- Testes
- Docker
- Observabilidade
