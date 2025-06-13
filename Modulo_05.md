Módulo 5 – Concorrência com Goroutines e Channels em Go

Este módulo aborda a concorrência em Go, um dos pontos fortes da linguagem, com foco em goroutines, channels, e ferramentas de sincronização como select, sync.WaitGroup e sync.Mutex. Para engenheiros Java, a concorrência em Go é mais leve e idiomática do que threads e ExecutorService. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura. O lab prático implementa um worker pool para processamento concorrente de tarefas.

  

1. Goroutines: O que São e Como Usar

O que são Goroutines?

- Goroutines são funções ou métodos executados de forma concorrente, gerenciados pelo runtime de Go (não pelo sistema operacional).
- São extremamente leves (poucos KB de memória) comparados a threads Java (que consomem MB).
- Criadas com a palavra-chave go antes da chamada de uma função.

Exemplo:

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

    go tarefa("Goroutine 1") // Executa concorrente

    tarefa("Main")           // Executa sequencialmente

    time.Sleep(1 * time.Second) // Aguarda para visualizar saída

}

Saída (pode variar devido à concorrência):

Main: 0

Goroutine 1: 0

Main: 1

Goroutine 1: 1

Main: 2

Goroutine 1: 2

Comparação com Java:

- Java: Threads (Thread, Runnable) ou ExecutorService são mais pesados e complexos.
- Go: Goroutines abstraem a complexidade, com o runtime gerenciando escalonamento.

Caso de uso: Goroutines são ideais para tarefas paralelas, como processar requisições HTTP ou executar cálculos independentes.

  

2. Channels (Unbuffered, Buffered)

Channels

- Mecanismo de comunicação entre goroutines, permitindo sincronização e troca de dados.
- Declarados com chan tipo.
- Dois tipos:

- Unbuffered: Bloqueia até que remetente e receptor estejam prontos.
- Buffered: Permite enviar até a capacidade do buffer sem bloquear.

Exemplo (Unbuffered Channel):

package main

  

import "fmt"

  

func main() {

    ch := make(chan string)

    go func() {

        ch <- "Mensagem da goroutine"

    }()

    msg := <-ch // Bloqueia até receber

    fmt.Println(msg) // Saída: Mensagem da goroutine

}

Exemplo (Buffered Channel):

package main

  

import "fmt"

  

func main() {

    ch := make(chan int, 2) // Buffer de tamanho 2

    ch <- 1

    ch <- 2

    fmt.Println(<-ch) // Saída: 1

    fmt.Println(<-ch) // Saída: 2

}

Comparação com Java:

- Java: Usa BlockingQueue ou sincronização com synchronized/Lock.
- Go: Channels são mais simples e seguros, evitando condições de corrida.

Caso de uso: Channels são usados para coordenar goroutines, como em pipelines de dados ou comunicação entre trabalhadores.

  

3. `select`, `sync.WaitGroup` e `sync.Mutex`

`select`

- Permite gerenciar múltiplos channels, escolhendo o primeiro pronto para operação.
- Similar a select em sistemas operacionais ou Selector em Java NIO.

Exemplo:

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

Saída:

Recebido: Canal 1

Recebido: Canal 2

`sync.WaitGroup`

- Sincroniza a espera pelo término de múltiplas goroutines.
- Similar a CountDownLatch em Java.

Exemplo:

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

`sync.Mutex`

- Garante exclusão mútua para proteger dados compartilhados.
- Similar a synchronized ou ReentrantLock em Java.

Exemplo:

package main

  

import (

    "fmt"

    "sync"

)

  

type Contador struct {

    mu    sync.Mutex

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

Caso de uso: select é usado em sistemas reativos, WaitGroup em tarefas paralelas, e Mutex em acesso a recursos compartilhados, como contadores ou caches.

  

4. Padrões de Concorrência em Go

5. Pipeline:

- Encadeia goroutines com channels para processar dados em etapas.
- Ex.: Ler dados → Transformar → Salvar.

3. Worker Pool:

- Um grupo de goroutines processa tarefas de uma fila compartilhada.
- Similar a ThreadPoolExecutor em Java.

5. Fan-out/Fan-in:

- Fan-out: Distribui tarefas para múltiplas goroutines.
- Fan-in: Combina resultados em um único channel.

7. Cancelamento:

- Usa context.Context para interromper goroutines (ex.: timeout de requisições).

Exemplo (Pipeline):

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

Caso de uso: Worker pools são comuns em sistemas de alta carga, como servidores web ou processamento de lotes.

  

📌 Lab: Criar um Worker Pool para Processamento Concorrente de Tarefas

Objetivo

Implementar um worker pool que processa tarefas de cálculo (ex.: calcular o quadrado de números) usando goroutines e channels, com sincronização via sync.WaitGroup e logging estruturado.

Passo a Passo

1. Configuração:

- Crie um diretório lab5:  
    mkdir lab5
- cd lab5
- go mod init github.com/seu-usuario/lab5
-   
    

3. Implementação: Crie um arquivo workerpool.go com o seguinte código:

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

    ID    int

    Valor int

}

  

// Resultado representa o resultado de uma tarefa

type Resultado struct {

    ID    int

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

3. Execução:  
    go run workerpool.go
4.   
    
5. Tarefa:

- Adicione um context.Context para suportar cancelamento do worker pool após um timeout.
- Implemente um mecanismo para lidar com erros em uma tarefa (ex.: valor inválido).
- Use select para processar resultados com um timeout por tarefa.

Saída esperada (console):

Tarefa 1: Quadrado = 1

Tarefa 2: Quadrado = 4

Tarefa 3: Quadrado = 9

Tarefa 4: Quadrado = 16

Tarefa 5: Quadrado = 25

Saída esperada (logs JSON, exemplo):

{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Tarefa enviada","tarefa_id":1,"valor":1}

{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Processando tarefa","worker":1,"tarefa_id":1,"valor":1}

...

{"time":"2025-06-12T01:07:01Z","level":"INFO","msg":"Processamento concluído","total_tarefas":5}

Caso de uso prático: O worker pool simula cenários reais, como processamento de filas de mensagens (ex.: RabbitMQ) ou tarefas em lote (ex.: redimensionamento de imagens).

  

Conclusão

Este módulo cobriu goroutines, channels (unbuffered e buffered), ferramentas de sincronização (select, sync.WaitGroup, sync.Mutex) e padrões de concorrência em Go. O lab prático implementou um worker pool, reforçando a aplicação prática desses conceitos. Engenheiros Java notarão semelhanças com ExecutorService, mas com uma abordagem mais leve e idiomática.

Próximos passos: No próximo módulo, exploraremos testes unitários, benchmarks e integração com bibliotecas externas, preparando o time para desenvolver aplicações robustas e testáveis em Go.
