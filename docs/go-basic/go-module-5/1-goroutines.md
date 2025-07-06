---
sidebar_position: 2
sidebar_label: Goroutines
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Goroutines: o que são e como usar

O que são goroutines?

- Goroutines são funções ou métodos executados de forma concorrente, `gerenciados pelo runtime de Go` (não pelo sistema operacional).
- São `extremamente leves` (poucos KB de memória) comparados a `threads` Java (que consomem MB).
- Criadas com a palavra-chave `go` antes da chamada de uma função.

## Exemplo

<InteractiveCodeSnippet 
    src="code/mod5/goroutines.go" 
    allowExecute={true} 
    allowEdit={false} />

## Saída (pode variar devido à concorrência)

```bash
Main: 0
Goroutine 1: 0
Main: 1
Goroutine 1: 1
Main: 2
Goroutine 1: 2
```

## Comparação com Java

### Java

- `Threads` (`Thread`, `Runnable`) ou `ExecutorService` são mais pesados e complexos.

### Go

- `Goroutines` abstraem a complexidade, com o runtime gerenciando escalonamento.

:::info Caso de uso
Goroutines são ideais para tarefas paralelas, como processar requisições HTTP ou executar cálculos independentes.
:::
