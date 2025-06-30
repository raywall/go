---
sidebar_position: 7
sidebar_label: Module 06
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Packages, modules and code organization in Go

<div className="row">
<div className="col">

This module explores code organization in Go, including idiomatic package structure, design conventions, dependency management with go mod, and best practices for Java engineers learning Go. The content is detailed but straightforward, with examples and use cases for future reference.

The hands-on lab reorganizes the CRUD from previous modules into multiple packages, using go mod to manage dependencies.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default} 
    style={{ transform:'scale(0.65)', marginTop:'-65px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Idiomatic package structure

### What are packages?

- Packages are `code organization units` in Go, similar to packages in Java (`package`).
- Each Go file belongs to a package, declared with the name `package`.
- The main package generates an executable; other packages are libraries.

### Idiomatic structure

- Go recommends a clear organization for projects:

- Project root: Contains go.mod and subdirectories.
- Packages: Organized into subdirectories with descriptive names (e.g. api, models, repo).
- Naming: Package names are short, in lowercase, without underscores (e.g. http, db).
- Export: Functions, types, and variables with capital letters are exported (e.g. func Create is visible outside the package).

### Example

```dirtree
my-project/
├── go.mod
├── main.go
├── models/
│   └── product.go
└── repo/
    └── memory.go
```

### Comparison with Java

#### Java

- Packages follow the `com.company.project` convention

#### Go

- Packages use module-based paths (e.g. `github.com/user/project/models`)

:::info Use case
Packages are used to modularize code into microservices, separating business logic, data access, and APIs.
:::

<br />

## Design conventions (`cmd`, `internal`, and `pkg`)

### Common conventions

#### `cmd/`

- Contains entry points (`main`) for executables.

##### Examples

- `cmd/api/main.go` for an API
- `cmd/cli/main.go` for a CLI tool

#### `internal/`

- Used for private packages, accessible only within the project or subdirectories.
- Ideal for sensitive logic or internal reuse.

#### `pkg/`

- Packages reusable by other projects (optional, less common).

#### Other directories

- Descriptive names such as: api, db, models, utils.

### Example structure

```dirtree
my-project/
├── go.mod
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── repo/
│       └── memory.go
└── models/
    └── product.go
```

### Comparison with Java

#### Java

- Rigid structure with `src/main/java` and **hierarchical packages**.

#### Go

- More flexible, but `internal` is similar to `package-private` in Java.

:::info Use case
Use internal to protect repository implementations, while cmd organizes multiple executables (e.g. server and migration tool).
:::

<br />

## `go mod` and versioning

### Go Modules

- Introduced in Go 1.11, replaces `$GOPATH` to manage dependencies.
- The `go.mod` file defines the module and its dependencies.

#### Example of go.mod

```bash
module github.com/your-github-user/your-project

go 1.21

require (
    github.com/google/uuid v1.6.0
)
```

> Click [here](@site/static/code/mod6/go.mod) to download this source code!

#### Main Commands

- `go mod init`: Initializes a module.
- `go mod tidy`: Adds used dependencies and removes unused ones.
- `go mod vendor` (optional): Creates a vendor/ directory with dependencies.

### Versioning

- Go uses `semantic versioning` (e.g. v1.6.0)
- Modules are referenced by repository URLs (e.g. `github.com/author/project`)
- Tags in the repository (e.g. `git tag v1.0.0`) define versions

#### Comparison with Java

##### Java

- Uses `Maven/Gradle` with centralized repositories (e.g. `Maven Central`)

##### Go

- Uses repository URLs directly, with proxy (e.g. `proxy.golang.org`)

:::info Use case
Go Modules simplify dependency management in large projects, such as APIs that integrate external libraries.
:::

<br />

## Dependency management with `go get` and `replace`

### `go get`

- Add or update dependencies in `go.mod`
  Example: `go get github.com/google/uuid@v1.6.0`

- Update to the latest version:
  `go get -u github.com/google/uuid`

### `replace`

- Allows you to replace a dependency with another version or local path.
- Useful for `local` development or `forks`

### Example of `go.mod` with `replace`

```bash
module github.com/your-github-user/your-project

go 1.21

require github.com/google/uuid v1.6.0
replace github.com/google/uuid => ../uuid-fork
```

> Click [here](@site/static/code/mod6/go-replace.mod) to download this source code!

### Comparison with Java

#### Java

- Use in `Maven` or `implementation` in Gradle

#### Go

- `go get` and `replace` are simpler, but less configurable

:::info Use case
`replace` is useful for testing local changes to dependencies before pushing
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
2. Create a `lab6` directory:

```bash
mkdir lab6
```

3. Access the directory and initialize a module:

```bash
cd lab6
go mod init github.com/your-github-user/lab6
```

4. Create the following structure for the project:

```dirtree
lab6/
├── go.mod
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── repo/
│       └── memory.go
└── models/
    └── product.go
```

> Click [here](@site/static/code/mod6/lab/lab6.zip) to download this source code!

</TabItem>
<TabItem value="app">

### Organize the CRUD project into multiple packages with `go mod`

#### Objective

Reorganize the CRUD of the previous modules into an idiomatic package structure, using go mod to manage dependencies and adding an external library (github.com/google/uuid) to generate IDs.

#### Step by step

1. Create the file `models/product.go` with the following code:

```go
package models

import "github.com/google/uuid"

// Product represents a product in the system
type Produto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

```

2. Create the file `internal/repo/memoria.go` with the following code:

```go
package repo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/your-github-user/lab6/models"
)

var (
	ErrInvlaidPrice    = errors.New("price cannot be negative")
	ErrProductNotFound = errors.New("product not found")
)

// ProductRepository defines the repository interface
type ProductRepository interface {
	Create(name string, price float64) (models.Product, error)
	Get(id uuid.UUID) (models.Product, error)
	List() ([]models.Product, error)
	Update(id uuid.UUID, name string, price float64) (models.Product, error)
	Delete(id uuid.UUID) error
}

// InMemoryRepository implements the repository in memory
type InMemoryRepository struct {
	products map[uuid.UUID]models.Product
	logger   *slog.Logger
}

// NewInMemoryRepository creates a new repository
func NewInMemoryRepository(logger *slog.Logger) *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[uuid.UUID]models.Product),
		logger:   logger,
	}
}

func (r *InMemoryRepository) Create(name string, price float64) (models.Product, error) {
	if price < 0 {
		r.logger.Error("failed to create a product", "error", ErrInvalidPrice, "name", name)
		return models.Product{}, ErrInvalidPrice
	}

    id := uuid.New()
	product := models.Product{ID: id, Name: name, Price: price}
	r.product[id] = product
	r.logger.Info("Product created", "id", id, "name", name, "price", price)
	return product, nil
}

func (r *InMemoryRepository) Get(id uuid.UUID) (models.Product, error) {
	product, exists := r.products[id]
	if !exists {
		r.logger.Error("failed to get a product", "error", ErrProductNotFound, "id", id)
		return models.Product{}, fmt.Errorf("get product id %s: %w", id, ErrProductNotFound)
	}

    r.logger.Info("Product founded", "id", id)
	return produto, nil
}

func (r *InMemoryRepository) List() ([]models.Product, error) {
	var products []models.Product
	for _, p := range r.products {
		products = append(products, p)
	}

    r.logger.Info("Listing products", "total", len(products))
	return products, nil
}

func (r *InMemoryRepository) Update(id uuid.UUID, name string, price float64) (models.Product, error) {
	if price < 0 {
		r.logger.Error("failed to update a product", "error", ErrInvalidPrice, "id", id)
		return models.Product{}, ErrInvalidPrice
	}

    product, exists := r.products[id]
	if !exists {
		r.logger.Error("failed to update a product", "error", ErrProductNotFound, "id", id)
		return models.Product{}, fmt.Errorf("update product id %s: %w", id, ErrProductNotFound)
	}

    product.Name = name
	product.Price = price

    r.products[id] = product
	r.logger.Info("Product updated", "id", id, "name", name, "price", price)
	return product, nil
}

func (r *InMemoryRepository) Delete(id uuid.UUID) error {
	if _, exists := r.products[id]; !exists {
		r.logger.Error("failed to delete a product", "error", ErrProductNotFound, "id", id)
		return fmt.Errorf("delete product id %s: %w", id, ErrProductNotFound)
	}

    delete(r.products, id)
	r.logger.Info("Product deleted", "id", id)
	return nil
}
```

3. Create the `cmd/api/main.go` file with the following code:

```go
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/your-github-product/lab6/internal/repo"
)

func showProducts(repo repo.ProductRepository) error {
	products, err := repo.List()
	if err != nil {
		return fmt.Errorf("show products: %w", err)
	}

    fmt.Println("List of products:")
	for _, p := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NewInMemoryRepository(logger)

    // Create products
	p1, err := repo.Create("Laptop", 999.99)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	repo.Create("Mouse", 29.99)

    // List products
	if err := showProducts(repo); err != nil {
		fmt.Println("Error:", err)
		return
	}

    // Get product
	if p, err := repo.Get(p1.ID); err == nil {
		fmt.Printf("Product founded: %+v\n", p)
	} else {
		fmt.Println("Error:", err)
	}

    // Update product
	if p, err := repo.Update(p1.ID, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Product updated: %+v\n", p)
	} else {
		fmt.Println("Error:", err)
	}

    // Delete product
	if err := repo.Delete(p1.ID); err == nil {
		fmt.Println("Product deleted successfully")
	} else {
		fmt.Println("Error:", err)
	}

    // List products again
	if err := showProducts(repo); err != nil {
		fmt.Println("Error:", err)
	}
}
```

#### Running

```bash
go mod tidy
go run cmd/api/main.go
```

</TabItem>
<TabItem value="task">

### Task

- Add an api package with a function that simulates an HTTP endpoint (e.g.: `ListProducts` returning JSON).
- Use replace in `go.mod` to test a local version of the `github.com/google/uuid` library.
- Create a utils package with a function to validate prices (e.g.: `ValidatePrice`).

#### Expected output

##### Console:

```bash
List of products:
ID: 1, Name: Laptop, Price: 999.99
ID: 2, Name: Mouse, Price: 29.99

Product founded: {ID:1 Name:Laptop Price:999.99}
Product updated: {ID:1 Name:Laptop Pro Price:1299.99}
Product deleted successfully

List of products:
ID: 2, Name: Mouse, Price: 29.99
```

##### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Product created","id":"1","name":"Laptop","price":999.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Product created","id":"2","name":"Mouse","price":29.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Listing products","total":2}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Product found","id":"1"}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Product updated","id":"1","name":"Laptop Pro","price":1299.99}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Product deleted","id":"1"}
{"time":"2025-06-12T01:09:00Z","level":"INFO","msg":"Listing products","total":1}
```

:::info Practical use case
This structure is typical of real Go projects, such as `RESTful APIs`, where packages separate business logic, data access, and input/output.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered package organization, design conventions (`cmd`, `internal`, `pkg`), dependency management with `go mod`, and the use of `go get` and `replace`. The practical lab reorganized CRUD into an idiomatic structure, integrating an external dependency. Java engineers will notice similarities to Maven/Gradle package organization, but with Go's simpler and more straightforward approach.

:::tip Next steps
In the next module, we will explore `unit testing`, `benchmarks`, and `integration with external libraries`, solidifying the best practices for robust and scalable applications in Go.
:::
