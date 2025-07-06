---
sidebar_position: 2
sidebar_label: Conteúdo programático
---

# Conteúdo programático

## [Módulo 01](./go-module-1/index.md) – Revisão rápida Go + setup Cloud

<div className="row">
<div className="col">

- Reforço prático da base Go
- Instalação AWS CLI / SDK Go
- Autenticação com AWS IAM (perfis e roles)
- Estrutura de projeto Go para cloud

### 📌 Laboratório

- Setup local + execução de chamadas autenticadas via SDK AWS (`STS/GetCallerIdentity`)

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    alt="A diaper brown gopher" />
</div>
</div>
---

## [Módulo 02](./go-module-2/index.md) – Armazenamento com S3 e DynamoDB

<div className="row">
<div className="col">

- Upload/download com `s3manager`
- Uso de presigned URLs
- Operações CRUD com `DynamoDB`
- Tipagem forte com structs + marshaling
- Padrão de repositório para acesso ao Dynamo

### 📌 Laboratório

- Criar um serviço de upload de arquivos com metadados persistidos em DynamoDB

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-30px' }}
    alt="Three gophers in a meeting, where one is explaining graphics on a slate while the others watch sitting on a round table" />
</div>
</div>
---

## [Módulo 03](./go-module-3/index.md) – Fila de Mensagens com SQS

<div className="row">
<div className="col">

- Produção e consumo de mensagens
- Dead Letter Queues
- Retry e visibilidade de mensagens
- Processamento assíncrono com goroutines

### 📌 Laboratório

- Criar uma fila de notificação e um worker para processá-la

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ marginTop:'-10px' }}
    alt="A gopher designing a flow diagram" />
</div>
</div>
---

## [Módulo 04](./go-module-4/index.md) – Eventos e Lambda com Go

<div className="row">
<div className="col">

- Escrever funções Lambda com Go
- Eventos de S3, Dynamo e SQS
- Empacotamento com `zip` e deploy com `SAM CLI`
- Timeout e limites de memória

### 📌 Laboratório

- Criar Lambda que reage a upload no S3, processa e persiste dados

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scalex(1)', marginTop:'-60px' }}
    alt="A clumsy gopher, wearing a brown sweatshirt and dark green pants, carrying some roles and having a coffee" />
</div>
</div>
---

## [Módulo 05](./go-module-5/index.md) – API Gateway + Lambda Proxy

<div className="row">
<div className="col">

- Criação de endpoints com API Gateway
- Integração direta com Lambda (Proxy)
- Controle de métodos e autenticação
- Padrão API Gateway + Lambda + Dynamo

### 📌 Laboratório

- Criar uma API Serverless completa em Go (CRUD de produtos)

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default}
    style={{ transform:'scale(1.1)', marginTop:'-30px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />

</div>
</div>
---

## [Módulo 06](./go-module-6/index.md) – Observabilidade: CloudWatch, Datadog e X-Ray

<div className="row">
<div className="col">

- **Boas práticas de logging estruturado com `slog`**
- Criação e envio de logs customizados para **CloudWatch Logs**
- Coleta de métricas customizadas com `PutMetricData` (namespace próprio)
- Integração com **AWS X-Ray** para rastreamento de chamadas (tracing)
- Uso do middleware X-Ray para instrumentar handlers HTTP

- **Integração com Datadog**:

  - Configuração do agente Datadog para ECS ou Lambda
  - Envio de métricas via `statsd` e `dd-trace-go`
  - Criação de spans manuais e automáticos (Datadog APM)
  - Logging estruturado com correlação de trace ID (Datadog Logging)

- Comparativo entre X-Ray e Datadog para tracing e métricas

### 📌 Laboratório

- Instrumentar uma aplicação Go com:

  - Logs estruturados no CloudWatch
  - Métricas customizadas para CloudWatch e Datadog
  - Traces distribuídos com X-Ray e Datadog APM
  - Correlação entre log, trace e métricas

- Gerar carga e validar os painéis em ambos os sistemas

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default}
    style={{ transform:'scale(0.9)', marginTop:'-30px' }}
    alt="A brown gopher with glasses and white shirt, organizing boxes in a warehouse" />

</div>
</div>
---

## [Módulo 07](./go-module-7/index.md) – Deploy e CI/CD com AWS

<div className="row">
<div className="col">

- Build cross-platform com `GOOS`/`GOARCH`
- Dockerização com multi-stage builds
- Deploy com ECS Fargate + Load Balancer
- Integração com CodePipeline

### 📌 Laboratório

- Dockerizar aplicação e publicar via ECS com configuração mínima

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default}
    style={{ transform:'scale(1.1)', marginTop:'25px' }}
    alt="A gopher with a box of small pieces and a paper where it fits them with the name Schema.go written at the top" />

</div>
</div>
---

## [Módulo 08](./go-module-8/index.md) – Secrets e configuração

<div className="row">
<div className="col">

- Uso do Secrets Manager com SDK Go
- Parâmetros no SSM Parameter Store
- Boas práticas para leitura segura de configurações

### 📌 Laboratório

- Configurar leitura dinâmica de secrets e variáveis de ambiente

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default}
    style={{ transform:'scale(1.5)', marginTop:'-10px' }}
    alt="A gopher by programming on a computer, supported by a box full of books" />

</div>
</div>
---

## [Módulo 09](./go-module-9/index.md) – Integração com EventBridge e SNS

<div className="row">
<div className="col">

- Disparo e escuta de eventos customizados
- Fanout de mensagens com SNS
- Padrão de eventos com payloads tipados

### 📌 Laboratório

- Criar um sistema que publica eventos e envia notificações

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default}
    style={{ transform:'scale(1.1)' }}
    alt="A Gopher teaching SQL class in a board, with the elephant, symbol of the PostgreSQL next door" />

</div>
</div>
---

## [Módulo 10](./go-module-10/index.md) – Infraestrutura como código com terraform

<div className="row">
<div className="col">

- Criação de recursos AWS com Terraform
- Integração entre Terraform e Go
- Separação de ambientes (dev/stage/prod)

### 📌 Laboratório

- Provisionar toda a stack Go + AWS com Terraform

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
