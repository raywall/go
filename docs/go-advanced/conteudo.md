---
sidebar_position: 2
sidebar_label: Conteúdo programático
---

# Conteúdo programático

## [Módulo 01](./go-module-1/index.md) – Anatomia da performance em Go

<div className="row">
<div className="col">

- Como o Go compila e executa
- Benchmarking com `testing.B`
- Perfilamento com `pprof`
- Ferramentas: `go tool trace`, `benchstat`, `race`

### 📌 Laboratório

- Benchmark de funções com diferentes abordagens

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    alt="A diaper brown gopher" />
</div>
</div>
---

## [Módulo 02](./go-module-2/index.md) – Otimização de alocação e memória

<div className="row">
<div className="col">

- Passagem por valor vs por referência
- Minimização de alocações com `sync.Pool`
- Zero-allocation APIs
- Profiling de heap e GC tuning

### 📌 Laboratório

- Reescrever código para reduzir GC pressure e melhorar throughput

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-30px' }}
    alt="Three gophers in a meeting, where one is explaining graphics on a slate while the others watch sitting on a round table" />
</div>
</div>
---

## [Módulo 03](./go-module-3/index.md) – Concorrência avançada com goroutines e channels

<div className="row">
<div className="col">

- Design de pipelines com canais
- Fan-in, fan-out, backpressure
- Evitando leaks com `select` e `default`
- `context.Context` para controle de cancelamento

### 📌 Laboratório

- Pipeline concorrente de scraping com controle de tempo e timeout

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ marginTop:'-10px' }}
    alt="A gopher designing a flow diagram" />
</div>
</div>
---

## [Módulo 04](./go-module-4/index.md) – Tratamento resiliente de falhas

<div className="row">
<div className="col">

- Retry com `exponential backoff`
- Implementação de **circuit breaker** simples
- `context.WithTimeout`, `WithCancel`, `WithValue`
- Tolerância a falhas com fallback e timeout

### 📌 Laboratório

- Serviço com retry inteligente e breaker entre dependências

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scalex(1)', marginTop:'-60px' }}
    alt="A clumsy gopher, wearing a brown sweatshirt and dark green pants, carrying some roles and having a coffee" />
</div>
</div>
---

## [Módulo 05](./go-module-5/index.md) – Paralelismo controlado com workers e semáforos

<div className="row">
<div className="col">

- Limitação de concorrência
- `sync.WaitGroup`, `sync.Mutex`, `sync.Cond`
- Implementação de rate limiters

### 📌 Laboratório

- Serviço que executa milhares de requisições com limitação de paralelismo

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default}
    style={{ transform:'scale(1.1)', marginTop:'-30px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />

</div>
</div>
---

## [Módulo 06](./go-module-6/index.md) – Estruturas de dados otimizadas

<div className="row">
<div className="col">

- Comparação entre maps, slices, arrays
- Algoritmos customizados
- Utilização de pacotes como `container/heap`, `ring`, `list`
- Quando usar `unsafe`

### 📌 Laboratório

- Criar estruturas otimizadas para cache com TTL e LRU

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default}
    style={{ transform:'scale(0.9)', marginTop:'-30px' }}
    alt="A brown gopher with glasses and white shirt, organizing boxes in a warehouse" />

</div>
</div>
---

## [Módulo 07](./go-module-7/index.md) – High performance networking

<div className="row">
<div className="col">

- Servidores HTTP low-level com `net/http` + `http.Server`
- Pooling de conexões
- Performance tuning de HTTP clients
- Uso de gRPC com Go (introdução)

### 📌 Laboratório

- Serviço com benchmark de throughput entre HTTP e gRPC

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default}
    style={{ transform:'scale(1.1)', marginTop:'25px' }}
    alt="A gopher with a box of small pieces and a paper where it fits them with the name Schema.go written at the top" />

</div>
</div>
---

## [Módulo 08](./go-module-8/index.md) – Streams e processamento de dados em tempo real

<div className="row">
<div className="col">

- Processamento contínuo com goroutines
- Padrão de buffer e janela
- Otimização com canais e lock-free design

### 📌 Laboratório

- Criar um processador de stream em tempo real com múltiplos consumidores

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default}
    style={{ transform:'scale(1.5)', marginTop:'-10px' }}
    alt="A gopher by programming on a computer, supported by a box full of books" />

</div>
</div>
---

## [Módulo 09](./go-module-9/index.md) – Integração com C/C++ e system calls

<div className="row">
<div className="col">

- Usando `cgo`
- Chamada de bibliotecas nativas (opcional, para bibliotecas críticas)
- Avaliação de custo-benefício

### 📌 Laboratório

- Criar binding simples com biblioteca C

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default}
    style={{ transform:'scale(1.1)' }}
    alt="A Gopher teaching SQL class in a board, with the elephant, symbol of the PostgreSQL next door" />

</div>
</div>
---

## [Módulo 10](./go-module-10/index.md) – Design de sistemas de alta performance

<div className="row">
<div className="col">

- Estratégias de particionamento e cache
- Microserviços com Go de alta performance
- Deployment otimizado (builds tiny, alpine, etc)
- Feature flags, toggles e circuit breaker via middleware

### 📌 Laboratório

- Serviço com rotas de alta disponibilidade e circuit breaker centralizado

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />

</div>
</div>
