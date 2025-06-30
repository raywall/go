---
sidebar_position: 8
sidebar_label: Module 07
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Testing and code quality in Go

<div className="row">
<div className="col">

This module covers testing and code quality in Go, including `unit testing` and `integration testing` with the `testing` package, using `mocks` with libraries such as `testify`, `benchmarks`, `profiling`, and `static analysis` tools. The content is aimed at Java engineers, comparing it to practices such as `JUnit` and `Mockito`, while remaining objective for future reference.

</div>
<div className="col col--4 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default} 
    style={{ transform:'scale(0.8)', marginTop:'-1.1rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

The practical lab implements unit and integration tests for CRUD from the previous modules, with a focus on `bug coverage`.

<br />

## Testing with the `testing` package

### `testing` package

- The standard testing package provides support for `unit tests` and `benchmarks`.
- Tests are written in files with the `_test.go` suffix and functions prefixed with Test.
- Uses `t.Error` or `t.Fatal` to report failures.

#### Example

```go
package main

import "testing"

func sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Errorf("sum(2, 3) = %d; waiting 5", result)
	}
}
```

> Click [here](@site/static/code/mod7/testing.go) to download this source code!

#### Running

```bash
go test
```

:::tip Tip
To run the test in all directories and subdirectories of the project, use `go test ./...`
:::

#### Comparison with Java

##### Java

- Uses `JUnit/TestNG` with annotations (`@Test`).

##### Go

- Simpler, `no annotations`, and with `naming conventions`.

:::info Use case
Unit tests with `testing` are ideal for validating pure functions or isolated components.
:::

<br />

## Unit and integration tests

### Unit tests

- Tests `isolated units` (e.g. a function or method)
- Uses test tables for multiple scenarios

#### Example (test table)

```go
package main

import "testing"

func TestSum(t *testing.T) {
	testes := []struct {
		name    string
		a, b    int
		waiting int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -1, -2},
		{"zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soma(tt.a, tt.b)
			if result != tt.waiting {
				t.Errorf("sum(%d, %d) = %d; waiting %d", tt.a, tt.b, result, tt.waiting)
			}
		})
	}
}
```

> Click [here](@site/static/code/mod7/teste-unitario.go) to download this source code!

### Integration tests

- Tests the `interaction between components` (e.g. repository and API)
- They can use `in-memory databases` or `mocks`

:::info Use case
Unit tests `verify business logic`; integration tests `ensure that modules work together` (e.g. repository and service)
:::

<br />

## Testing with `mocks` (`testify` and `gomock`)

### `testify` package

- Simplifies `assertions` and `mock creation`
- Includes `assert` and `mock` to make testing easier

#### Installation

```bash
go to github.com/stretchr/testify
```

#### Example with testify/mock

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Buscar(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestService(t *testing.T) {
	repo := &MockRepository{}
	repo.On("Get", 1).Return("Product", nil)

    result, err := repo.Get(1)
	if err != nil {
		t.Fatal(err)
	}

    if result != "Product" {
		t.Errorf("waiting Product, got %s", result)
	}
	repo.AssertExpectations(t)
}
```

> Click [here](@site/static/code/mod7/teste-mock.go) to download this source code!

### Package `gomock`

- **Automatically generates mocks** from `interfaces`
- **More powerful** for complex projects

#### Installation

```bash
go install github.com/golang/mock/gomock@latest
go install github.com/golang/mock/mockgen@latest
```

:::info Use case
`Mocks` are useful for testing _services that depend on repositories_ or `external APIs`, similar to `Mockito` in Java.
:::

<br />

## Benchmarks and profiling

### Benchmarks

- Functions prefixed with `Benchmark` **measure performance**
- Use `b.N` to _run the code multiple times_

#### Example

```go
package main

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(2, 3)
	}
}
```

> Click [here](@site/static/code/mod7/benchmarks.go) to download this source code!

#### Running

```bash
go test -bench=.
```

### Profiling

Go offers built-in tools for performance analysis:

- `pprof`: Profile CPU, memory, and deadlocks
- `go test -cpuprofile cpu.out`: Generate CPU profile
- `go tool pprof cpu.out`: Analyze the profile

#### Example

```bash
go test -bench=. -cpuprofile cpu.out
go tool pprof cpu.out
```

### Comparison with Java

#### Java

- Uses `VisualVM` or `JProfiler` for profiling

#### Go

- Built-in tools (`pprof`) are more lightweight and integrated

:::info Use case
`Benchmarks` help optimize critical functions; while `profiling` **identifies bottlenecks** in _APIs_ or _concurrent systems_. :::

<br />

## Tools: `go vet`, `golint`, `staticcheck`

### `go vet`

- Analyzes code for common errors (e.g. incorrect string formatting)
  **Runs with**: `go vet ./...`

### `golint` (deprecated, but still used)

- Suggests style improvements (e.g. variable names)
  **Modern alternative**: `golangci-lint`

### `staticcheck`

- Advanced static analysis tool, checking for bugs, performance and style. **Installation**: `go install honnef.co/go/tools/cmd/staticcheck@latest`
  **Usage**: `staticcheck ./...`

### Comparison with Java

#### Java

- Uses `Checkstyle`, `PMD` or `SonarQube`

#### Go

- Tools are more `integrated` and `lightweight`, with `focus on simplicity`.

:::info Use case
These tools ensure consistency and quality in large projects, such as microservices.
:::

<br />

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

## Laboratório

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

1. Use the structure of [Module 06](go-module-6.md)

```dirtree
lab6/
├── go.mod
├── cmd/
│   └── api/
│		├── main.go
│       └── main_test.go
├── internal/
│   └── repo/
│		├── memory.go
│       └── memory_test.go
└── models/
    └── product.go
```

2. Add the package `testify`:

```bash
go get github.com/stretchr/testigy
```

> Click [here](@site/static/code/mod7/lab/lab7.zip) to download this source code!

</TabItem>
<TabItem value="app">

### Create unit and integration tests for CRUD with error coverage

#### Objective

Implement unit and integration tests for CRUD in [Module 06](go-module-6.md), using `testing` and `testify`, with a focus on error coverage. Generate coverage reports and use `staticcheck` for analysis.

#### Step by step

1. Create the file `internal/repo/memoria_test.go` using the code:

```go
import (
	"github.com/google/uuid"
	"github.com/your-github-user/labлог6/models"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestInMemoryRepository(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NewInMemoryRepository(logger)

	t.Run("Create product successfully", func(t *testing.T) {
		product, err := repo.Create("Laptop", 999.99)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, product.ID)
		assert.Equal(t, "Laptop", product.Name)
		assert.Equal(t, 999.99, product.Price)
	})

	t.Run("Create product with invalid price", func(t *testing.T) {
		_, err := repo.Criar("Laptop", 1)
		assert.ErrorIs(t, err, ErrInvalidPrice)
	})

	t.Run("Get a existing product", func(t *testing.T) {
		product, err := repo.Create("Mouse", 29.99)
		assert.NoError(t, err)

		founded, err := repo.Get(product.ID)
		assert.NoError(t, err)
		assert.Equal(t, product, founded)
	})

	t.Run("Get a nonexistent product", func(t *testing.T) {
		_, err := repo.Get(uuid.New())
		assert.ErrorIs(t, err, ErrProductNotFound)
	})

	t.Run("List products", func(t *testing.T) {
		repo = NewInMemoryRepository(logger)
		repo.Create("Laptop", 999.99)
		repo.Create("Mouse", 29.99)

		products, err := repo.List()
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})

	t.Run("Update an existing product", func(t *testing.T) {
		product, err := repo.Create("Laptop", 999.99)
		assert.NoError(t, err)

		updated, err := repo.Update(product.ID, "Laptop Pro", 1299.99)
		assert.NoError(t, err)
		assert.Equal(t, "Laptop Pro", updated.Name)
		assert.Equal(t, 1299.99, updated.Price)
	})

	t.Run("Update product with invalid price", func(t *testing.T) {
		product, err := repo.Create("Laptop", 999.99)
		assert.NoError(t, err)

		_, err = repo.Update(produto.ID, "Laptop Pro", 1)
		assert.ErrorIs(t, err, ErrInvalidPrice)
	})

	t.Run("Delete an existing product", func(t *testing.T) {
		product, err := repo.Create("Laptop", 999.99)
		assert.NoError(t, err)

		err = repo.Delete(product.ID)
		assert.NoError(t, err)

		_, err = repo.Get(product.ID)
		assert.ErrorIs(t, err, ErrProductNotFound)
	})

	t.Run("Delete a non-existent product", func(t *testing.T) {
		err := repo.Delete(uuid.New())
		assert.ErrorIs(t, err, ErrProductNotFound)
	})
}
```

2. Create the file `cmd/api/main_test.go` (integration test) using the code:

```go
package main

import (
	"github.com/your-github-user/lab6/internal/repo"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NewInMemoryRepository(logger)

	t.Run("Complete CRUD flow", func(t *testing.T) {

		// Create
		product, err := repo.Create("Laptop", 999.99)
		assert.NoError(t, err)

		// List
		products, err := repo.List()
		assert.NoError(t, err)
		assert.Len(t, products, 1)

		// Get
		founded, err := repo.Get(product.ID)
		assert.NoError(t, err)
		assert.Equal(t, product, founded)

		// Update
		updated, err := repo.Update(product.ID, "Laptop Pro", 1299.99)
		assert.NoError(t,, err)
		assert.Equal(t, "Laptop Pro", updated.Name)

		// Delete
		err = repo.Delete(product.ID)
		assert.NoError(t, err)

		// Verify exclusion
		products, err = repo.List()
		assert.NoError(t, err)
		assert.Len(t, products, 0)
	})
}
```

##### Running

###### Run the tests

- `go test ./... -v`

###### Generate coverage report

- `go test ./... -coverprofile=cover.out`

- `go tool cover -html=cover.out`

###### Analyze with staticcheck

- `staticcheck ./...`

</TabItem>
<TabItem value="task">

### Task

- Add a `benchmark` for the List function of the repository.
- Create a `mock` with `testify/mock` to test a service that uses `RepositorioProdutos`.
- Use `go vet` and `golangci-lint` to verify the code.

#### Expected output (tests)

```bash
=== RUN   TestInMemoryRepository
=== RUN   TestInMemoryRepository/Create_product_successfully
=== RUN   TestInMemoryRepository/Create_product_with_invalid_price
=== RUN   TestInMemoryRepository/Search_for_existing_product
=== RUN   TestInMemoryRepository/Search_for_non-existing_product
=== RUN   TestInMemoryRepository/List_products
=== RUN   TestInMemoryRepository/Update_existing_product
=== RUN   TestInMemoryRepository/Update_product_with_invalid_price
=== RUN   TestInMemoryRepository/Delete_existing_product
=== RUN   TestInMemoryRepository/Delete_non-existing_product
--- PASS: TestInMemoryRepository (0.00s)
--- PASS: TestInMemoryRepository/Create_product_successfully (0.00s)
--- PASS: TestInMemoryRepository/Create_product_with_invalid_price (0.00s)
--- PASS: TestInMemoryRepository/Search_existing_productsignalize
--- PASS: TestInMemoryRepository/SearchSweeten
--- PASS: TestInMemoryRepository/List_products (0.00s)
--- PASS: TestInMemoryRepository/Update_existing_product (0.00s)
--- PASS: TestInMemoryRepository/Update_product_with_invalid_price (0.00s)
--- PASS: TestInMemoryRepository/Delete_existing_product (0.00s)
--- PASS: TestInMemoryRepository/Delete_inexisting_product (0.00s)
=== RUN   TestMainIntegration
=== RUN   TestMainIntegration/Complete_CRUD_flow
--- PASS: TestMainIntegration (0.00s)
--- PASS: TestMainIntegration/Complete_CRUD_flow (0.00s)
PASS
ok      github.com/your-github-user/lab6/... 0.012s
```

:::info Practical use case
**Unit and integration tests ensure the robustness** of CRUD, while `bug coverage prevents failures` in scenarios such as `invalid prices` or `non-existent IDs`.
:::

</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered testing with `testing`, `unit tests` and `integration tests`, `mocks` with `testify`, `benchmarks`, `profiling` and static analysis tools (`go vet` and `staticcheck`). The practical lab implemented complete tests for CRUD, with bug coverage and structured logging. Java engineers will notice similarities to `JUnit` and `Mockito`, but with the simplicity and native integration of the Go ecosystem.

:::tip Next steps
In the next module, we will explore building `RESTful APIs` with libraries like `gin` or `echo`, integrating CRUD into a real web server.
:::
