---
sidebar_position: 4
sidebar_label: Ferramentas de controle
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# `select`, `sync.WaitGroup` e `sync.Mutex`

## `select`

- Permite gerenciar múltiplos `channels`, escolhendo o primeiro pronto para operação.
- Similar a `select` em sistemas operacionais ou `Selector` em Java NIO.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod5/select.go" 
    allowExecute={true} 
    allowEdit={false} />

### Saída

```bash
Recebido: Canal 1
Recebido: Canal 2
```

## `sync.WaitGroup`

- Sincroniza a espera pelo término de múltiplas `goroutines`.
- Similar a `CountDownLatch` em Java.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod5/sync-waitgroup.go" 
    allowExecute={true} 
    allowEdit={false} />

## `sync.Mutex`

- Garante exclusão mútua para proteger dados compartilhados.
- Similar a `synchronized` ou `ReentrantLock` em Java.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod5/sync-mutex.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
`select` é usado em sistemas reativos, `WaitGroup` em tarefas paralelas, e `Mutex` em acesso a recursos compartilhados, como contadores ou caches.
:::
