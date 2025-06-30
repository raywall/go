---
sidebar_position: 7
sidebar_label: Module 05
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Concurrency with goroutines and channels

<div className="row">
<div className="col">

This module covers concurrency in Go, one of the strengths of the language, focusing on `goroutines`, `channels`, and synchronization tools such as `select`, `sync.WaitGroup`, and `sync.Mutex`. For Java engineers, concurrency in Go is more lightweight and idiomatic than `threads` and `ExecutorService`. The content is detailed but objective, with examples and use cases for future reference.

The hands-on lab implements a `worker pool` for concurrent processing of tasks.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default} 
    style={{ marginTop:'-50px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Goroutines: what they are and how to use them

What are goroutines?

- Goroutines are functions or methods executed concurrently, `managed by the Go runtime` (not by the operating system).
- They are `extremely lightweight` (few KB of memory) compared to Java `threads` (which consume MB).
- Created with the `go` keyword before calling a function.

### Example

```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task("Goroutine 1")    // Concurrently executes
	task("Main")              // Performs sequentially
	time.Sleep(1 * time.Second) // Wait to view output
}
```

> Click [here](@site/static/code/mod5/goroutines.go) to download this source code!

### Output (may vary due to competition)

```bash
Main: 0
Goroutine 1: 0
Main: 1
Goroutine 1: 1
Main: 2
Goroutine 1: 2
```

### Comparison with Java

#### Java

- `Threads` (`Thread`, `Runnable`) or `ExecutorService` are heavier and more complex.

#### Go

- `Goroutines` abstract the complexity, with the runtime managing scheduling.

:::info Use case
Goroutines are ideal for parallel tasks, such as processing HTTP requests or performing independent calculations.
:::

<br />

## Channels (`unbuffered`, `buffered`)

### Channels

- Communication mechanism between goroutines, allowing synchronization and data exchange.
- Declared with `chan` type
- Two types:
  - **Unbuffered**: Blocks until sender and receiver are ready.
  - **Buffered**: Allows sending up to the buffer capacity without blocking.

#### Examples

##### Unbuffered channels

```go
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Goroutine message"
	}()

    msg := <-ch      // Block until receiving
	fmt.Println(msg) // Output: Goroutine message
}
```

> Click [here](@site/static/code/mod5/channel-unbuffered.go) to download this source code!

##### Buffered channel

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffer of size 2
	ch <- 1
	ch <- 2

    fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
```

> Click [here](@site/static/code/mod5/channel-buffered.go) to download this source code!

### Comparison with Java

#### Java

- Uses `BlockingQueue` or synchronization with `synchronized/Lock`

#### Go

- Channels are simpler and safer, avoiding race conditions.

:::info Use case
Channels are used to coordinate goroutines, such as in data pipelines or communication between workers.
:::

<br />

## `select`, `sync.WaitGroup` and `sync.Mutex`

### `select`

- Allows you to manage multiple `channels`, choosing the first one ready for operation.
- Similar to `select` in operating systems or `Selector` in Java NIO.

#### Example

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
		ch1 <- "Channel 1"
	}()

    go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Channel 2"
	}()

    for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
```

> Click [here](@site/static/code/mod5/select.go) to download this source code!

#### Saída

```bash
Received: Channel 1
Received: Channel 2
```

### `sync.WaitGroup`

- Synchronizes the waiting for multiple goroutines to finish.
- Similar to `CountDownLatch` in Java.

#### Example

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

    wg.Wait()
	fmt.Println("All tasks finished")
}
```

> Click [here](@site/static/code/mod5/sync-waitgroup.go) to download this source code!

### `sync.Mutex`

- Ensures mutual exclusion to protect shared data.
- Similar to `synchronized` or `ReentrantLock` in Java.

#### Example

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	c := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

    wg.Wait()
	fmt.Println("Final value:", c.value) // Output: Final value: 1000
}
```

> Click [here](@site/static/code/mod5/sync-mutex.go) to download this source code!

:::info Use case
`select` is used in reactive systems, `WaitGroup` in parallel tasks, and `Mutex` in accessing shared resources, such as counters or caches.
:::

<br />

## Concurrency Patterns in Go

### Pipeline:

- Chains `goroutines` with `channels` to process data in steps.
  **Example**: Read data → Transform → Save.

### Worker pool

- A `group of goroutines` processes tasks from a shared queue.
- Similar to `ThreadPoolExecutor` in Java.

### Fan-out and Fan-in

- **Fan-out**: Distributes tasks to `multiple goroutines`
- **Fan-in**: Combines results into `single channel`

### Cancellation

- Uses `context.Context` to interrupt goroutines (e.g. request timeout).

### Example

#### Pipeline

```go
package main

import "fmt"

func generate(numbers ...int) chan int {
	out := make(chan int)

    go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

    return out
}
func double(in chan int) chan int {
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
	input := generate(1, 2, 3)
	output := double(input)

    for n := range output {
		fmt.Println(n) // Output: 2, 4, 6
	}
}
```

> Click [here](@site/static/code/mod5/pipeline.go) to download this source code!

:::info Use Case
`Worker pools` are common in high-load systems, such as web servers or batch processing.
:::

<br />

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

## Lab

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Setup', value: 'config'},
{label: 'Exercise', value: 'app'},
{label: 'Task', value: 'task'},
]}>
<TabItem value="config">

### Setup

1. Install Go and check with `go version`
2. Create a `lab5` directory:

```bash
mkdir lab5
```

3. Access the directory and initialize a module:

```bash
cd lab5
go mod init github.om/your-github-user/lab5
```

</TabItem>

<TabItem value="app">

### Create a worker pool for concurrent processing of tasks

#### Objective

Implement a worker pool that processes calculation tasks (e.g., calculating the square of numbers) using goroutines and channels, with synchronization via `sync.WaitGroup` and structured `logging`.

#### Step by step

1. Create a `workerpool.go` file with the following code:

```go
package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

// Task structure represents a task to be processed
type Task struct {
	ID    int
	Value int
}

// Result structure represents the result of a task
type Result struct {
	ID     int
	Double int
}

// Worker processes a channel tasks
func worker(id int, tasks chan Task, results chan Resultado, wg *sync.WaitGroup, logger *slog.Logger) {
	defer wg.Done()
	for t := range tasks {
		logger.Info("Processing task", "worker", id, "task_id", t.ID, "value", t.Value))
		time.Sleep(100 * time.Millisecond) // Simulates work
		results <- Result{ID: t.ID, Double: t.Value * t.Value}
	}
}

func main() {
	// Logger configuration
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    // Channels for tasks and results
	tasks := make(chan Task, 10)
	results := make(chan Result, 10)
	var wg sync.WaitGroup

    // Start 3 workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg, logger)
	}

    // Send tasks
	numTasks := 5
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{ID: i, Value: i}
			logger.Info("Task submitted", "task_id", i, "value", i)
		}
		close(tasks)
	}()

    // Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

    // Display results
	for r := range results {
		fmt.Printf("Task %d: Double = %d\n", r.ID, r.Double)
	}
	logger.Info("Processing completed", "total_tasks", numTasks)
}
```

> Click [here](@site/static/code/mod5/lab/workerpool.go) to download this source code!

##### Running

```bash
go run workerpool.go
```

</TabItem>
<TabItem value="task">

### Task

- Add a `context.Context` to support worker pool cancellation after a timeout.
- Implement a mechanism to handle errors in a task (e.g. invalid value).
- Use `select` to process results with a timeout per task.

#### Output expected

##### Console

```bash
Task 1: Double = 1
Task 2: Double = 4
Task 3: Double = 9
Task 4: Double = 16
Task 5: Double = 25
```

##### JSON logs (example)

```json
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Task submitted","task_id":1,"value":1}
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Processing task","worker":1,"task_id":1,"value":1}
...
{"time":"2025-06-12T01:07:01Z","level":"INFO","msg":"Processing completed","total_tasks":5}
```

:::info Practical use case
The `worker pool` simulates real-world scenarios, such as processing message queues (e.g. `RabbitMQ`) or batch tasks (e.g. resizing images).
:::

</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered `goroutines`, `channels` (`unbuffered` and `buffered`), synchronization tools (`select`, `sync.WaitGroup`, `sync.Mutex`) and concurrency patterns in Go. The hands-on lab implemented a `worker pool`, reinforcing the practical application of these concepts. Java engineers will notice similarities to ExecutorService, but with a more lightweight and idiomatic approach.

:::tip Next steps
In the next module, we will explore `unit testing`, `benchmarks` and `integration with external libraries`, preparing the team to develop robust and testable applications in Go.
:::
