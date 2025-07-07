---
sidebar_position: 2
sidebar_label: Conteúdo programático
---

# Conteúdo programático

## [Módulo 01](./go-module-1/index.md) – Introdução e fundamentos da linguagem

<div className="row">
<div className="col">

- História e propósito do Go
- Características da linguagem (simplicidade, performance, concorrência)
- Instalação, workspace, go mod
- Estrutura básica de um programa
- Tipos primitivos, funções, variáveis, constantes

##### 📌 Laboratório

- Hello world, manipulação de variáveis e tipos, primeiro programa com go run

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    alt="A diaper brown gopher" />
</div>
</div>
---

## [Módulo 02](./go-module-2/index.md) – Construtos da linguagem e boas práticas

<div className="row">
<div className="col">

- Condicionais, laços, arrays, slices e maps
- Structs, interfaces e orientação a interfaces
- Manipulação de erros e boas práticas de código idiomático
- Pacotes, organização de projetos

#### 📌 Laboratório

- Implementação de uma calculadora modular com boas práticas
- Exercícios com slices, maps e structs

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-30px' }}
    alt="Three gophers in a meeting, where one is explaining graphics on a slate while the others watch sitting on a round table" />
</div>
</div>
---

## [Módulo 03](./go-module-3/index.md) – Concorrência em Go e aplicações serverless

<div className="row">
<div className="col">

- Goroutines, canais e sync.WaitGroup
- Modelos concorrentes no contexto da AWS
- Introdução à AWS Lambda
- Benefícios de usar Go em ambientes serverless

#### 📌 Laboratório

- Criação de Lambda em Go
- Deploy via AWS CLI
- Benchmark simples com goroutines

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ marginTop:'-10px' }}
    alt="A gopher designing a flow diagram" />
</div>
</div>
---

## [Módulo 04](./go-module-4/index.md) – Go com AWS Lambda e API Gateway

<div className="row">
<div className="col">

- Handlers HTTP e integração com o API Gateway
- Eventos e payloads com aws-lambda-go
- Timeout, cold start e boas práticas

### 📌 Laboratório

- API RESTful com Lambda + API Gateway em Go
- Deploy com SAM ou Terraform

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scalex(1)', marginTop:'-60px' }}
    alt="A clumsy gopher, wearing a brown sweatshirt and dark green pants, carrying some roles and having a coffee" />
</div>
</div>
---

## [Módulo 05](./go-module-5/index.md) – Persistência com DynamoDB

<div className="row">
<div className="col">

- Conceitos básicos e modelagem de dados no DynamoDB
- SDK da AWS para Go (aws-sdk-go-v2)
- PutItem, GetItem, Query e Scan

### 📌 Laboratório

- API com persistência em DynamoDB usando Go
- Criação de tabela via IaC
- Testes locais com DynamoDB Local

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default}
    style={{ transform:'scale(1.1)', marginTop:'-30px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />

</div>
</div>
---

## [Módulo 06](./go-module-6/index.md) – Persistência com Amazon RDS (MySQL e PostgreSQL)

<div className="row">
<div className="col">

- Acesso a bancos relacionais com Go (database/sql)
- Drivers recomendados para MySQL e PostgreSQL
- Pool de conexões e práticas para ambientes de nuvem

### 📌 Laboratório

- API simples com leitura/gravação em RDS
- Exemplos com MySQL e PostgreSQL
- Uso de gorm ou sqlx (opcional)

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default}
    style={{ transform:'scale(0.9)', marginTop:'-30px' }}
    alt="A brown gopher with glasses and white shirt, organizing boxes in a warehouse" />

</div>
</div>
---

## [Módulo 07](./go-module-7/index.md) – Go com Amazon Aurora Serverless

<div className="row">
<div className="col">

- Diferenças entre RDS tradicional e Aurora Serverless
- Estratégias para escalabilidade e performance
- Conexão segura com Go e IAM

### 📌 Laboratório

- Criação e integração com Aurora Serverless
- Teste de carga simulando concorrência com goroutines

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default}
    style={{ transform:'scale(1.1)', marginTop:'25px' }}
    alt="A gopher with a box of small pieces and a paper where it fits them with the name Schema.go written at the top" />

</div>
</div>
---

## [Módulo 08](./go-module-8/index.md) – Armazenamento com S3

<div className="row">
<div className="col">

- Upload e download de objetos
- Gerenciamento de buckets e permissões
- Uso eficiente de s3manager no Go SDK

### 📌 Laboratório

- Aplicação Go que realiza upload de imagens para S3
- Geração de URLs assinadas

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default}
    style={{ transform:'scale(1.5)', marginTop:'-10px' }}
    alt="A gopher by programming on a computer, supported by a box full of books" />

</div>
</div>
---

## [Módulo 09](./go-module-9/index.md) – Gerenciamento de filas e notificações com SQS e SNS

<div className="row">
<div className="col">

- Uso de SQS para comunicação assíncrona
- Publicação de eventos com SNS
- Estratégias de retry e dead-letter queues

### 📌 Laboratório

- Aplicação Go que publica e consome mensagens de SQS
- Integração com SNS para envio de notificações

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default}
    style={{ transform:'scale(1.1)' }}
    alt="A Gopher teaching SQL class in a board, with the elephant, symbol of the PostgreSQL next door" />

</div>
</div>
---

## [Módulo 10](./go-module-10/index.md) – Containers e orquestração: ECS, EKS e EC2

<div className="row">
<div className="col">

- Dockerização de aplicações Go
- Implantação em EC2, ECS e EKS
- Quando usar cada serviço

### 📌 Laboratório

- Container Go + Dockerfile
- Deploy em ECS Fargate com Load Balancer
- Simulação de escalabilidade no EKS

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
---

## [Módulo 11](./go-module-11/index.md) – Cache com ElastiCache for Redis

<div className="row">
<div className="col">

- Conceitos de cache e TTL
- Conexão e uso com Go (go-redis)
- Estratégias para cache eficiente

### 📌 Laboratório

- Integração de API Go com Redis
- Simulação de cache hit/miss

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
---

## [Módulo 12](./go-module-12/index.md) – Gerenciamento de configurações e segredos

<div className="row">
<div className="col">

- AWS Secrets Manager vs Parameter Store
- Boas práticas de segurança na AWS com Go
- IAM roles e políticas recomendadas

### 📌 Laboratório

- Aplicação Go que lê segredos seguros de forma dinâmica
- Parâmetros e versionamento no Parameter Store

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
---

## [Módulo 13](./go-module-13/index.md) – Balanceamento de carga com ALB e NLB

<div className="row">
<div className="col">

- Diferenças entre ALB e NLB
- Integração com ECS/EKS
- Estratégias de alta disponibilidade e health check

### 📌 Laboratório

- Deploy de aplicação Go atrás de ALB e NLB
- Simulação de falhas e observação de comportamento

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
---

## [Módulo 14](./go-module-14/index.md) – CObservabilidade com CloudWatch, OpenTelemetry e Datadog

<div className="row">
<div className="col">

- Logs estruturados e métricas customizadas em Go
- Tracing distribuído com OpenTelemetry
- Integração com Datadog APM

### 📌 Laboratório

- Instrumentação de aplicação Go com OpenTelemetry
- Visualização no CloudWatch e Datadog

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
