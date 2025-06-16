---
sidebar_position: 7
sidebar_label: Módulo 05
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Concorrência com goroutines e channels em Go

Este módulo aborda a concorrência em Go, um dos pontos fortes da linguagem, com foco em `goroutines`, `channels`, e ferramentas de sincronização como `select`, `sync.WaitGroup` e `sync.Mutex`. Para engenheiros Java, a concorrência em Go é mais leve e idiomática do que `threads` e `ExecutorService`. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura.

O lab prático implementa um `worker pool` para processamento concorrente de tarefas.

<br />

## Goroutines: o que são e como usar

O que são goroutines?

- Goroutines são funções ou métodos executados de forma concorrente, `gerenciados pelo runtime de Go` (não pelo sistema operacional).
- São `extremamente leves` (poucos KB de memória) comparados a `threads` Java (que consomem MB).
- Criadas com a palavra-chave `go` antes da chamada de uma função.

### Exemplo

```go
package main

import (
	"fmt"
	"time"
)

func tarefa(nome string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", nome, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go tarefa("Goroutine 1")    // Executa concorrente
	tarefa("Main")              // Executa sequencialmente
	time.Sleep(1 * time.Second) // Aguarda para visualizar saída
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/goroutines.go)!

### Saída (pode variar devido à concorrência)

```bash
Main: 0
Goroutine 1: 0
Main: 1
Goroutine 1: 1
Main: 2
Goroutine 1: 2
```

### Comparação com Java

#### Java

- `Threads` (`Thread`, `Runnable`) ou `ExecutorService` são mais pesados e complexos.

#### Go

- `Goroutines` abstraem a complexidade, com o runtime gerenciando escalonamento.

:::info Caso de uso
Goroutines são ideais para tarefas paralelas, como processar requisições HTTP ou executar cálculos independentes.
:::

<br />

## Channels (`unbuffered`, `buffered`)

### Channels

- Mecanismo de comunicação entre goroutines, permitindo sincronização e troca de dados.
- Declarados com `chan tipo`

- Dois tipos:
  - **Unbuffered**: Bloqueia até que remetente e receptor estejam prontos.
  - **Buffered**: Permite enviar até a capacidade do buffer sem bloquear.

#### Exemplos

##### Unbuffered channel

```go
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Mensagem da goroutine"
	}()

    msg := <-ch      // Bloqueia até receber
	fmt.Println(msg) // Saída: Mensagem da goroutine
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/channel-unbuffered.go)!

##### Buffered channel

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffer de tamanho 2
	ch <- 1
	ch <- 2

    fmt.Println(<-ch) // Saída: 1
	fmt.Println(<-ch) // Saída: 2
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/channel-buffered.go)!

### Comparação com Java

#### Java

- Usa `BlockingQueue` ou sincronização com `synchronized/Lock`

#### Go

- Channels são mais simples e seguros, evitando condições de corrida.

:::info Caso de uso
Channels são usados para coordenar goroutines, como em pipelines de dados ou comunicação entre trabalhadores.
:::

<br />

## `select`, `sync.WaitGroup` e `sync.Mutex`

### `select`

- Permite gerenciar múltiplos `channels`, escolhendo o primeiro pronto para operação.
- Similar a `select` em sistemas operacionais ou `Selector` em Java NIO.

#### Exemplo

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

    go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Canal 1"
	}()

    go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Canal 2"
	}()

    for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido:", msg2)
		}
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/select.go)!

#### Saída

```bash
Recebido: Canal 1
Recebido: Canal 2
```

### `sync.WaitGroup`

- Sincroniza a espera pelo término de múltiplas `goroutines`.
- Similar a `CountDownLatch` em Java.

#### Exemplo

```go
package main

import (
	"fmt"
	"sync"
)

func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tarefa(i, &wg)
	}

    wg.Wait()
	fmt.Println("Todas as tarefas concluídas")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/sync-waitgroup.go)!

### `sync.Mutex`

- Garante exclusão mútua para proteger dados compartilhados.
- Similar a `synchronized` ou `ReentrantLock` em Java.

#### Exemplo

```go
package main

import (
	"fmt"
	"sync"
)

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func main() {
	c := Contador{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Incrementar()
			wg.Done()
		}()
	}

    wg.Wait()
	fmt.Println("Valor final:", c.valor) // Saída: Valor final: 1000
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/sync-mutex.go)!

:::info Caso de uso
`select` é usado em sistemas reativos, `WaitGroup` em tarefas paralelas, e `Mutex` em acesso a recursos compartilhados, como contadores ou caches.
:::

<br />

## Padrões de Concorrência em Go

### Pipeline:

- Encadeia `goroutines` com `channels` para processar dados em etapas.
  **Exemplo**: Ler dados → Transformar → Salvar.

### Worker pool

- Um `grupo de goroutines` processa tarefas de uma fila compartilhada.
- Similar a `ThreadPoolExecutor` em Java.

### Fan-out e Fan-in

- **Fan-out**: Distribui tarefas para `múltiplas goroutines`
- **Fan-in**: Combina resultados em `um único channel`

### Cancelamento

- Usa `context.Context` para interromper goroutines (ex: timeout de requisições).

### Exemplo

#### Pipeline

```go
package main

import "fmt"

func gerar(nums ...int) chan int {
	out := make(chan int)

    go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

    return out
}
func dobrar(in chan int) chan int {
	out := make(chan int)

    go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

    return out
}

func main() {
	entrada := gerar(1, 2, 3)
	saida := dobrar(entrada)

    for n := range saida {
		fmt.Println(n) // Saída: 2, 4, 6
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/pipeline.go)!

:::info Caso de uso
`Worker pools` são comuns em sistemas de alta carga, como servidores web ou processamento de lotes.
:::

<br />

---

## Laboratório

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

### Configuração

1. Instale o Go e verifique com `go version`
2. Crie o diretório `lab5`:

```bash
mkdir lab5
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab5
go mod init github.om/seu-usuario/lab5
```

</TabItem>

<TabItem value="app">

### Criar um worker pool para processamento concorrente de tarefas

#### Objetivo

Implementar um worker pool que processa tarefas de cálculo (ex.: calcular o quadrado de números) usando goroutines e channels, com sincronização via sync.WaitGroup e logging estruturado.

#### Passo a passo

1. Crie um arquivo `workerpool.go` com o seguinte código:

```go
package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

// Tarefa representa uma tarefa a ser processada
type Tarefa struct {
	ID    int
	Valor int
}

// Resultado representa o resultado de uma tarefa
type Resultado struct {
	ID       int
	Quadrado int
}

// Worker processa tarefas de um channel
func worker(id int, tarefas chan Tarefa, resultados chan Resultado, wg *sync.WaitGroup, logger *slog.Logger) {
	defer wg.Done()
	for t := range tarefas {
		logger.Info("Processando tarefa", "worker", id, "tarefa_id", t.ID, "valor", t.Valor)
		time.Sleep(100 * time.Millisecond) // Simula trabalho
		resultados <- Resultado{ID: t.ID, Quadrado: t.Valor * t.Valor}
	}
}

func main() {
	// Configurar logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    // Canais para tarefas e resultados
	tarefas := make(chan Tarefa, 10)
	resultados := make(chan Resultado, 10)
	var wg sync.WaitGroup

    // Iniciar 3 workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tarefas, resultados, &wg, logger)
	}

    // Enviar tarefas
	numTarefas := 5
	go func() {
		for i := 1; i <= numTarefas; i++ {
			tarefas <- Tarefa{ID: i, Valor: i}
			logger.Info("Tarefa enviada", "tarefa_id", i, "valor", i)
		}
		close(tarefas)
	}()

    // Coletar resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

    // Exibir resultados
	for r := range resultados {
		fmt.Printf("Tarefa %d: Quadrado = %d\n", r.ID, r.Quadrado)
	}
	logger.Info("Processamento concluído", "total_tarefas", numTarefas)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod5/lab/workerpool.go)!

##### Execução

```bash
go run workerpool.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione um `context.Context` para suportar cancelamento do worker pool após um timeout.
- Implemente um mecanismo para lidar com erros em uma tarefa (ex: valor inválido).
- Use `select` para processar resultados com um timeout por tarefa.

#### Saída esperada

##### Console

```bash
Tarefa 1: Quadrado = 1
Tarefa 2: Quadrado = 4
Tarefa 3: Quadrado = 9
Tarefa 4: Quadrado = 16
Tarefa 5: Quadrado = 25
```

##### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Tarefa enviada","tarefa_id":1,"valor":1}
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Processando tarefa","worker":1,"tarefa_id":1,"valor":1}
...
{"time":"2025-06-12T01:07:01Z","level":"INFO","msg":"Processamento concluído","total_tarefas":5}
```

:::info Caso de uso prático
O `worker pool` simula cenários reais, como processamento de filas de mensagens (ex: `RabbitMQ`) ou tarefas em lote (ex: redimensionamento de imagens).
:::

</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu `goroutines`, `channels` (`unbuffered` e `buffered`), ferramentas de sincronização (`select`, `sync.WaitGroup`, `sync.Mutex`) e padrões de concorrência em Go. O lab prático implementou um `worker pool`, reforçando a aplicação prática desses conceitos. Engenheiros Java notarão semelhanças com ExecutorService, mas com uma abordagem mais leve e idiomática.

:::tip Próximos passos
No próximo módulo, exploraremos `testes unitários`, `benchmarks` e `integração com bibliotecas externas`, preparando o time para desenvolver aplicações robustas e testáveis em Go.
:::
