---
sidebar_position: 7
sidebar_label: Concorrência
---

# Concorrência com goroutines e channels em Go

<div className="row">
<div className="col">

Este módulo aborda a concorrência em Go, um dos pontos fortes da linguagem, com foco em `goroutines`, `channels`, e ferramentas de sincronização como `select`, `sync.WaitGroup` e `sync.Mutex`. Para engenheiros Java, a concorrência em Go é mais leve e idiomática do que `threads` e `ExecutorService`. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura.

O lab prático implementa um `worker pool` para processamento concorrente de tarefas.

## Conteúdo

- [Goroutines: o que são e como usar](./1-goroutines.md)
- [Channels (unbuffered, buffered)](./2-channels.md)
- [`select`, `sync.WaitGroup`, `sync.Mutex`](./3-goroutines-handle.md)
- [Padrões de concorrência em Go](./4-padrao-concorrencia.md)

📌 [Lab](./5-laboratorio.md): Criar um worker pool para processamento concorrente de tarefas.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default} 
    style={{ marginTop:'-50px' }}
    alt="A diaper brown gopher" />
</div>
</div>
