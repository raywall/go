---
sidebar_position: 6
sidebar_label: Module 04
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Error handling

<div className="row">
<div className="col">

This module explores the error handling approach in Go, which differs significantly from the exception model used in Java. The focus is on the philosophy of explicit errors, the use of the error type, `wrapping`/`unwrapping` techniques, and `structured logging`. The content is aimed at Java engineers, with practical examples and objective use cases for future reference.

The hands-on lab implements functions with error handling and structured logging.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scale(0.8)', marginTop:'-65px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Go philosophy: explicit errors

- **Errors as values**: In Go, errors are values ​​of type error, returned explicitly by functions, instead of thrown as exceptions in Java (`try-catch`).
- **Clarity and predictability**: The programmer should explicitly check for errors, reducing unexpected behavior.
- **Avoid `panic`**: Unlike Java, panic is reserved for unrecoverable failures (e.g. critical initialization failure) and not for normal flow.

### Comparison with Java

#### Java

```java
try { ... } catch (Exception e) { ... }
```

#### Go

```go
if err != nil { return err }
```

:::info Use Case
The explicit approach is ideal for systems where robustness is critical, such as server APIs or CLI tools.
:::

<br />

## Pattern `error`, `errors.New`, and `fmt.Errorf`

### Type `error`

- **A built-in interface**: `type error interface { Error() string }`
- Any type that implements the `Error() string` method is an error

### `errors.New`

- Creates simple errors with a fixed message

#### Example

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("zero division")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: zero division
		return
	}

	fmt.Println("Resultado:", result)
}
```

> Click [here](@site/static/code/mod4/error.go) to download this source code!

### `fmt.Errorf`

- Dynamically formats error messages, similar to `String.format` in Java.

#### Example

```go
func getIndex(slice []int, index int) (int, error) {
    if index >= len(slice) {
        return 0, fmt.Errorf("index %d out of bounds (%d)", index, len(slice))
    }

    return slice[index], nil
}
```

> Click [here](@site/static/code/mod4/errorf.go) to download this source code!

:::info Use Case
`errors.New` is used for fixed errors, while `fmt.Errorf` is ideal for errors with dynamic context (e.g. input validation)
:::

<br />

## Wrapping and Unwrapping with `errors.Is` and `errors.As`

### Wrapping

- Allows you to add context to an error while keeping the original error
- Used with the `%w` operator in `fmt.Errorf`

#### Example

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func getData(id int) error {
	return ErrNotFound
}

func process(id int) error {
	err := getData(id)
	if err != nil {
		return fmt.Errorf("failure to process id %d: %w", id, err)
	}
	return nil
}

```

> Click [here](@site/static/code/mod4/wrapping.go) to download this source code!

### `errors.Is`

- Checks whether an error (or its encapsulated errors) matches a specific error.

#### Example

```go
func main() {
	err := process(45)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Error is 'not found'") // Output: Error is 'not found'
	}
	fmt.Println(err) // Output: Failure to process ID 45: Not found
}
```

> Click [here](@site/static/code/mod4/errors-is.go) to download this source code!

### `errors.As`

- Extracts an error of a specific type from a string of encapsulated errors.

#### Example

```go
package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func operation() error {
	return fmt.Errorf("context: %w", &MyError{Message: "specific failure"})
}

func main() {
	err := operation()

	var myError *MyError
	if errors.As(err, &myError) {
		fmt.Println("Captured error:", myError.Message) // Output: specific failure
	}
}
```

> Click [here](@site/static/code/mod4/errors-as.go) to download this source code!

#### Comparison with Java:

- `Wrapping` in Go is similar to `Exception.getCause()` in Java, but more explicit.
- `errors.Is` and `errors.As` replace checks like instanceof or exception type comparisons.

:::info Use case
Wrapping is useful in APIs to add context (e.g. request ID) without losing the original error.
:::

<br />

## `log` and `log/slog` packages

### `log` package

- Provides basic logging, with output to `stderr` by default.
- Functions like `log.Print`, `log.Fatal` and `log.Panic`.

#### Example

```go
package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("Starting application...")
	err := errors.New("critical failure")

	if err != nil {
		log.Fatal("Error:", err) // Comes out with code 1
	}
}
```

> Click [here](@site/static/code/mod4/log.go) to download this source code!

### `log/slog` package (Go 1.21+)

- Provides structured logging, with support for JSON and custom fields.
- Best suited for modern applications such as microservices.

#### Example

```go
package main

import (
	"errors"
	"log/slog"
	"os"
)

func main() {
	// configure logger JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// log with structured fields
	logger.Info("processing request", "id", 45, "method", "GET")
	err := errors.New("operation failure")

	logger.Error("error not found", "error", err, "id", 45)
}
```

> Click [here](@site/static/code/mod4/slog.go) to download this source code!

#### Output (JSON)

```json
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"processing request","id":45,"method":"GET"}
{"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"error not found","error":"operation failed","id":45}
```

#### Comparison with Java

- `log` is similar to `System.out.println` or `java.util.logging`
- `log/slog` is comparable to libraries like `Log4j` or `SLF4J`, with structured logging support.

:::info Use case
`log/slog` is ideal for distributed applications, where structured logs facilitate analysis in tools such as `ELK Stack` or `Prometheus`.
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
2. Create a `lab4` directory:

```bash
mkdir lab4
```

3. Access the directory and initialize a module:

```bash
cd lab4
go mod init github.com/your-github-user/lab4
```

</TabItem>
<TabItem value="app">

### Create functions with error handling and structured logging

#### Goal

Refactor the CRUD from Module 3 to include robust error handling and structured logging with log/slog, while maintaining the repository interface.

#### Step by step

1. Create a `crud.go` file with the following code:

```go
package main

import (
	"errors"
	"fmt"
	"log/slog"
)

// Customized errors
var (
	ErrInvalidPrice    = errors.New("price cannot be negative")
	ErrProductNotFound = errors.New("product not found")
)

// Product structure
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Repository interface
type ProductRepository interface {
	Create(name string, price float64) (Product, error)
	Get(id int) (Product, error)
	List() ([]Product, error)
	Update(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

// In-memory repository
type InMemoryRepository struct {
	products  []Product
	idCounter int
	logger    *slog.Logger
}

// New repository with logger
func NewInMemoryRepository(logger *slog.Logger) *InMemoryRepository {
	return &InMemoryRepository{
		products:  make([]Product, 0), // Initialize an empty slice for products
		idCounter: 0,                  // Initialize the ID counter
		logger:    logger,             // Associates the logger with the repository
	}
}

// Interface methods
func (r *InMemoryRepository) Add(name string, price float64) (Product, error) {
	if price < 0 {
		r.logger.Error("failed to create a new product", "error", ErrInvalidPrice, "name", name)
		return Product{}, ErrInvalidPrice
	}

    r.idCounter++
	product := Product{ID: r.idCounter, Name: name, Price: price}
	r.products = append(r.products, product)
	r.logger.Info("Product created", "id", product.ID, "name", name, "price", price)
	return product, nil
}

func (r *InMemoryRepository) Get(id int) (Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			r.logger.Info("Product founded", "id", id)
			return p, nil
		}
	}

    r.logger.Error("failed to get a product", "error", ErrProductNotFound, "id", id)
	return Product{}, fmt.Errorf("get product id %d: %w", id, ErrProductNotFound)
}

func (r *InMemoryRepository) List() ([]Product, error) {
	r.logger.Info("listing products", "total", len(r.products))
	return r.products, nil
}

func (r *InMemoryRepository) Update(id int, name string, price float64) (Product, error) {
	if price < 0 {
		r.logger.Error("failed to update a product", "error", ErrInvalidPrice, "id", id)
		return Product{}, ErrInvalidPrice
	}

    for i, p := range r.products {
		if p.ID == id {
			r.products[i] = Product{ID: id, Name: name, Price: price}
			r.logger.Info("product updated", "id", id, "name", name, "price", price)
			return r.produts[i], nil
		}
	}

    r.logger.Error("failed to update a product", "error", ErrProductNotFound, "id", id)
	return Produto{}, fmt.Errorf("update product id %d: %w", id, ErrProductNotFound)
}

func (r *InMemoryRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			r.logger.Info("Product deleted", "id", id)
			return nil
		}
	}

    r.logger.Error("failed to delete a product", "error", ErrProductNotFound, "id", id)
	return fmt.Errorf("delete product id %d: %w", id, ErrProductNotFound)
}

// Function to show products
func showProducts(repo ProductRepository) error {
	products, err := repo.List()
	if err != nil {
		return fmt.Errorf("list products: %w", err)
	}
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

    return nil
}

func main() {
	// Initialize repository with logger
	repo := NewInMemoryRepository()

    // Create products
	p1, err := repo.Create("Laptop", 999.99)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	repo.Create("Mouse", 29.99)

    // List products
	fmt.Println("List of products:")
	if err := showProducts(repo); err != nil {
		fmt.Println("Erroe:", err)
		return
	}

    // Get a product
	if p, err := repo.Buscar(p1.ID); err == nil {
		fmt.Printf("Product founded: %+v\n", p)
	} else {
		fmt.Println("Error:", err)
	}

    // Update a product
	if p, err := repo.Update(p1.ID, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Product updated: %+v\n", p)
	} else {
		fmt.Println("Error:", err)
	}

    // Try to update with an invalid price
	if _, err := repo.Update(p1.ID, "Laptop Pro", -100); err != nil {
		fmt.Println("Expected error:", err) // Output: price cannot be negative
	}

    // Delete a product
	if err := repo.Delete(2); err == nil {
		fmt.Println("Product deleted successfully")
	} else {
		fmt.Println("Error:", err)
	}

    // List the products again
	fmt.Println("Final list:")
	if err := showProducts(repo); err != nil {
		fmt.Println("Error:", err)
	}
}
```

> Click [here](@site/static/code/mod4/lab/crud.go) to download this source code!

#### Running:

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Task

- Add a custom error for when the product name is empty.
- Implement a function that uses errors.As to capture and handle specific errors (e.g.: ErrInvalidPrice).
- Configure the logger to save logs to a file, in addition to displaying them in the console.

#### Expected output

##### Console

```bash
List of products:
ID: 1, Name: Laptop, Price: 999.99
ID: 2, Name: Mouse, Price: 29.99

Product founded: {ID:1 Name:Laptop Price:999.99}
Product updated: {ID:1 Name:Laptop Pro Price:1299.99}
Expected error: price cannot be negative
Product deleted successfully

Final list:
ID: 1, Name: Laptop Pro, Price: 1299.99
```

#### JSON logs (example)

```json
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Product created","id":1,"name":"Laptop","price":999.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"ProductCreated","id":2,"name":"Mouse","price":29.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listing products","total":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Product founded","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Product founded","id":1,"name":"Laptop Pro","price":1299.99}
{"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"Failed to update product","error":"price cannot be negative","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Product deleted","id":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listing products","total":1}
```

:::info Practical use case
This lab simulates a real system with structured logging, useful for debugging and monitoring distributed applications, such as RESTful APIs.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered Go's explicit error philosophy, the use of `errors.New`, `fmt.Errorf`, `wrapping/unwrapping` with `errors.Is` and `errors.As`, and logging with `log` and `log/slog`. The hands-on lab reinforces the application of these concepts in a CRUD, with structured logs for easier maintenance. Java engineers will notice the difference in relation to the exception model, but will see similarities with structured logging in frameworks such as `SLF4J`.

:::tip Next steps
In the next module, we will explore concurrency with goroutines and channels, as well as testing and benchmarks, preparing the team to develop scalable and robust applications in Go.
:::
