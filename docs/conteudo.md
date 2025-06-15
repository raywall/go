---
sidebar_position: 2
---

# Conteúdo programático

## [Módulo 01](go-module-1.md) – Introdução e Fundamentos da Linguagem

- História e propósito do Go
- Características da linguagem (simplicidade, performance, concorrência)
- Instalação, workspace, go mod
- Estrutura básica de um programa
- Tipos primitivos, funções, variáveis, constantes

📌 Lab: Hello world, manipulação de variáveis e tipos, primeiro programa com go run.

---

## [Módulo 02](go-module-2.md) – Estruturas de Controle e Tipos Compostos

- if, for, switch, defer
- Arrays, slices, maps
- Structs e tags de struct
- Ponteiros (conceito e aplicação)

📌 Lab: Implementar uma API fake para CRUD em memória usando slices e maps.

---

## [Módulo 03](go-module-3.md) – Funções, Métodos e Interfaces

- Funções com múltiplos retornos
- Funções anônimas e closures
- Métodos em structs
- Interfaces e duck typing
- Boas práticas e princípios de design idiomático em Go

📌 Lab: Refatorar o CRUD usando interfaces para repositórios.

---

## [Módulo 04](go-module-4.md) – Tratamento de Erros

- Filosofia do Go: erros explícitos
- Padrão error, errors.New, fmt.Errorf
- Wrapping e unwrapping com errors.Is, errors.As
- Pacote log e log/slog

📌 Lab: Criar funções com tratamento de erros e logging estruturado.

---

## [Módulo 05](go-module-5.md) – Concorrência com Goroutines e Channels

- Goroutines: o que são e como usar
- Channels (unbuffered, buffered)
- select, sync.WaitGroup, sync.Mutex
- Padrões de concorrência em Go

📌 Lab: Criar um worker pool para processamento concorrente de tarefas.

---

## [Módulo 06](go-module-6.md) – Pacotes, Módulos e Organização do Código

- Estrutura de pacotes idiomática
- Convenções de projeto (cmd, internal, pkg)
- go mod e versionamento
- Gerenciamento de dependências com go get, replace

📌 Lab: Organizar o projeto CRUD em múltiplos pacotes com go mod.

---

## [Módulo 07](go-module-7.md) – Testes e Qualidade de Código

- Testes com testing
- Testes de unidade e integração
- Testes com mocks (testify, gomock)
- Benchmarks e profiling
- Ferramentas: go vet, golint, staticcheck

📌 Lab: Criar testes unitários e de integração para o CRUD com cobertura de erro.

---

## [Módulo 08](go-module-8.md) – Web APIs com net/http e Gin

- Servidor HTTP com net/http
- Middlewares e handlers
- Framework Gin: roteamento, binding, validação
- JSON, status codes e headers

📌 Lab: Implementar uma API RESTful com Gin + validação.

---

## [Módulo 09](go-module-9.md) – Persistência com Banco de Dados

- Drivers e database/sql
- ORM com gorm
- Migrations com golang-migrate
- Repositórios e testes de integração com DB

📌 Lab: Persistir o CRUD em banco real (PostgreSQL por exemplo).

---

## [Módulo 10](go-module-10.md) – Deploy, Observabilidade e Boas Práticas

- Build com go build, cross-compilation
- Docker com Go
- Logging estruturado (slog, zap)
- Tracing com OpenTelemetry
- Linter, cobertura, documentação automática com godoc

📌 Lab: Containerizar o serviço e expor métricas/trace/logs.
