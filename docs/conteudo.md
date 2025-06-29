---
sidebar_position: 2
sidebar_label: Conteúdo programático
---

# Conteúdo programático

## [Módulo 01](go-module-1.md) – Introdução e Fundamentos da Linguagem

<div className="row">
<div className="col">
- História e propósito do Go
- Características da linguagem (simplicidade, performance, concorrência)
- Instalação, workspace, go mod
- Estrutura básica de um programa
- Tipos primitivos, funções, variáveis, constantes

📌 Lab: Hello world, manipulação de variáveis e tipos, primeiro programa com go run.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    alt="A diaper brown gopher" />
</div>
</div>
---

## [Módulo 02](go-module-2.md) – Estruturas de Controle e Tipos Compostos

<div className="row">
<div className="col">
- if, for, switch, defer
- Arrays, slices, maps
- Structs e tags de struct
- Ponteiros (conceito e aplicação)

📌 Lab: Implementar uma API fake para CRUD em memória usando slices e maps.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-30px' }}
    alt="Three gophers in a meeting, where one is explaining graphics on a slate while the others watch sitting on a round table" />
</div>
</div>
---

## [Módulo 03](go-module-3.md) – Funções, Métodos e Interfaces

<div className="row">
<div className="col">
- Funções com múltiplos retornos
- Funções anônimas e closures
- Métodos em structs
- Interfaces e duck typing
- Boas práticas e princípios de design idiomático em Go

📌 Lab: Refatorar o CRUD usando interfaces para repositórios.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ marginTop:'-10px' }}
    alt="A gopher designing a flow diagram" />
</div>
</div>
---

## [Módulo 04](go-module-4.md) – Tratamento de Erros

<div className="row">
<div className="col">

- Filosofia do Go: erros explícitos
- Padrão error, errors.New, fmt.Errorf
- Wrapping e unwrapping com errors.Is, errors.As
- Pacote log e log/slog

📌 Lab: Criar funções com tratamento de erros e logging estruturado.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scalex(1)', marginTop:'-60px' }}
    alt="A clumsy gopher, wearing a brown sweatshirt and dark green pants, carrying some roles and having a coffee" />
</div>
</div>
---

## [Módulo 05](go-module-5.md) – Concorrência com Goroutines e Channels

<div className="row">
<div className="col">

- Goroutines: o que são e como usar
- Channels (unbuffered, buffered)
- select, sync.WaitGroup, sync.Mutex
- Padrões de concorrência em Go

📌 Lab: Criar um worker pool para processamento concorrente de tarefas.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default}
    style={{ transform:'scale(1.1)', marginTop:'-30px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />

</div>
</div>
---

## [Módulo 06](go-module-6.md) – Pacotes, Módulos e Organização do Código

<div className="row">
<div className="col">

- Estrutura de pacotes idiomática
- Convenções de projeto (cmd, internal, pkg)
- go mod e versionamento
- Gerenciamento de dependências com go get, replace

📌 Lab: Organizar o projeto CRUD em múltiplos pacotes com go mod.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default}
    style={{ transform:'scale(0.9)', marginTop:'-30px' }}
    alt="A brown gopher with glasses and white shirt, organizing boxes in a warehouse" />

</div>
</div>
---

## [Módulo 07](go-module-7.md) – Testes e Qualidade de Código

<div className="row">
<div className="col">

- Testes com testing
- Testes de unidade e integração
- Testes com mocks (testify, gomock)
- Benchmarks e profiling
- Ferramentas: go vet, golint, staticcheck

📌 Lab: Criar testes unitários e de integração para o CRUD com cobertura de erro.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default}
    style={{ transform:'scale(1.1)', marginTop:'25px' }}
    alt="A gopher with a box of small pieces and a paper where it fits them with the name Schema.go written at the top" />

</div>
</div>
---

## [Módulo 08](go-module-8.md) – Web APIs com net/http e Gin

<div className="row">
<div className="col">

- Servidor HTTP com net/http
- Middlewares e handlers
- Framework Gin: roteamento, binding, validação
- JSON, status codes e headers

📌 Lab: Implementar uma API RESTful com Gin + validação.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default}
    style={{ transform:'scale(1.5)', marginTop:'-10px' }}
    alt="A gopher by programming on a computer, supported by a box full of books" />

</div>
</div>
---

## [Módulo 09](go-module-9.md) – Persistência com Banco de Dados

<div className="row">
<div className="col">

- Drivers e database/sql
- ORM com gorm
- Migrations com golang-migrate
- Repositórios e testes de integração com DB

📌 Lab: Persistir o CRUD em banco real (PostgreSQL por exemplo).

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default}
    style={{ transform:'scale(1.1)' }}
    alt="A Gopher teaching SQL class in a board, with the elephant, symbol of the PostgreSQL next door" />

</div>
</div>
---

## [Módulo 10](go-module-10.md) – Deploy, Observabilidade e Boas Práticas

<div className="row">
<div className="col">

- Build com go build, cross-compilation
- Docker com Go
- Logging estruturado (slog, zap)
- Tracing com OpenTelemetry
- Linter, cobertura, documentação automática com godoc

📌 Lab: Containerizar o serviço e expor métricas/trace/logs.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
