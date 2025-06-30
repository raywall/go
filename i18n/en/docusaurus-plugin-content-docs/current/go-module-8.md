---
sidebar_position: 9
sidebar_label: Módulo 08
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Web APIs with `net/http` and `Gin` in Go

<div className="row">
<div className="col">

This module covers building `RESTful APIs` in Go, starting with the standard `net/http` package and moving on to the `Gin` framework, which simplifies `routing`, `binding`, and `validation`. The content is geared toward Java engineers, with comparisons to Spring Boot, and remains straightforward with practical examples and use cases for future reference.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default} 
    style={{ transform:'scale(1.3)', marginTop:'-3rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

The hands-on lab `implements a RESTful API` for the CRUD of the previous modules, using `Gin` with `input validation`.

<br />

## HTTP Server with `net/http`

### `net/http` package

- Provides functionality for `creating HTTP servers` and `handling requests/responses`.
- **Simple but powerful**, used in many Go projects without frameworks.

#### Example

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

    http.ListenAndServe(":8080", nil)
}
```

> Click [here](@site/static/code/mod8/api-http.go) to download this source code!

#### Running

```bash
go run main.go
```

> Access the address [http://localhost:8080/hello](http://localhost:8080/hello) from your browser

#### Comparison with Java

##### Java

- `Spring Boot` or `Servlets` require more configuration (e.g. `@RestController`)

##### Go

- `net/http` is _lighter, with less abstractions_.

:::info Use case
Simple servers, such as internal tools or minimal APIs.
:::

<br />

## Middlewares and handlers

### Handlers

- Functions that **handle HTTP requests**, with signature `func(w http.ResponseWriter, r *http.Request)`
- Can be registered with `http.HandleFunc`

### Middlewares

- Functions that **intercept requests/responses**, useful for `authentication`, `logging`, etc.

- Chain `handlers`, returning an `http.Handler`

#### Example (middleware with logging)

```go
package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

    mux.Handle("/", loggingMiddleware(mux))
	http.ListenAndServe(":8080", nil)
}
```

> Click [here](@site/static/code/mod8/api-middleware.go) to download this source code!

#### Comparison with Java

##### Java

- `Middlewares` are filters in Spring (`@Filter`) or `interceptors`

##### Go

- `Middlewares` are **more explicit**, using `handlers` composition

:::info Use case
Middlewares for `logging`, `authentication` (e.g. JWT), or `header validation`.
:::

<br />

## Gin Framework: routing, binding, validation

### Gin

- Lightweight framework for RESTful APIs, with `fast routing` and `JSON support`, `binding` and `validation`. - **Installation**: `go get github.com/gin-gonic/gin`

#### Routing

- Uses `route groups` and `HTTP methods` (`GET`, `POST`, `PUT`, `DELETE` etc.)
- Supports `URL parameters` and `query strings`.

##### Example

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name})
	})

    r.Run(":8080")
}
```

> Click [here](@site/static/code/mod8/api-gin-roteamento.go) to download this source code!

#### Binding and validation

- Converts JSON or forms into `structs`, with `validation via tags`.
- Requires `validator` library: `go get github.com/go-playground/validator/v10`

##### Example

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

func main() {
	r := gin.Default()
	r.POST("/products", func(c *gin.Context) {
		var p Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	})

    r.Run(":8080")
}
```

> Click [here](@site/static/code/mod8/api-gin-binding-validacao.go) to download this source code!

##### Comparison with Java:

###### Java

- Spring Boot uses `@RestController`, `@RequestBody` and validation with `@Valid`.

###### Go

- Gin is **lighter**, with `binding` and `validation` **directly via tags**.

:::info Use case
Gin is ideal for `scalable RESTful APIs`, such as e-commerce systems or microservices. :::

<br />

## JSON, status codes and headers

### JSON

- Gin `automatically serializes/deserializes JSON` with `c.JSON`
- Uses `encoding/json` internally

### Status codes

- Uses `net/http constants` (e.g. `http.StatusOK`, `http.StatusBadRequest`)

### Headers

- Handled with `c.Header` or `c.GetHeader`

### Example

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.Header("X-Custom-Header", "value")
		c.JSON(http.StatusOK, gin.H{
			"message": "API working",
		})
	})

    r.Run(":8080")
}
```

> Click [here](@site/static/code/mod8/api-gin-json.go) to download this source code!

### Test

Use the `curl` tool

```bash
curl -X GET http://localhost:8080/api
```

#### Response

```json
{ "message": "API working" }
```

:::info Use Case
`JSON` for API responses, `status codes` to indicate success/error, and `headers` for authentication or metadata.
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

2. Add the required dependencies:

```bash
go get github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
```

3. Update the `go.mod` file:

```bash
module github.com/seu-usuario/lab8

go 1.21

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/uuid v1.6.0
)
```

> Click [here](@site/static/code/mod8/lab/lab8.zip) to download the source code!

</TabItem>
<TabItem value="app">

### Implement a RESTful API with Gin + validation

#### Objective

Implement a RESTful API for the CRUD of [Module 06](go-module-6.md), using `Gin` for `routing`, `binding` and `validation`, integrating the in-memory repository and `structured logging`.

#### Step by step

1. Update the contents of the `models/product.go` file using the code:

```go
import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" binding:"required"`
	Price float64   `json:"price" binding:"required,gt=0"`
}
```

2. Update the contents of the `internal/repo/memoria.go` file using the code:

```go
package repo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seu-usuario/lab6/models"
)

var (
	ErrPriceInvalid    = errors.New("price cannot be negative")
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	Create(name string, price float64) (models.Product, error)
	Get(id uuid.UUID) (models.Product, error)
	List() ([]models.Product, error)
	Update(id uuid.UUID, name string, price float64) (models.Product, error)
	Delete(id uuid.UUID) error
}

type InMemoryRepository struct {
	products map[uuid.UUID]models.Product
	logger   *slog.Logger
}

func NewInMemoryRepository(logger *slog.Logger) *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[uuid.UUID]models.Product),
		logger:   logger,
	}
}

func (r *InMemoryRepository) Create(name string, price float64) (models.Product, error) {
	if price < 0 {
		r.logger.Error("failed to create a new product", "error", ErrInvalidPrice, "name", name)

        return models.Product{}, ErrInvalidPrice
	}

    id := uuid.New()
	product := models.Product{ID: id, Name: name, Price: price}

    r.products[id] = product
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
	return product, nil
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
	if !exsists {
		r.logger.Error("failed to update a product", "error", ErrProductNotFound, "id", id)

        return models.Product{}, fmt.Errorf("update product id %s: %w", id, ErrProductNotfound)
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

3. Update the contents of the `cmd/api/main.go` file using the code:

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/your-github-user/lab6/internal/repo"
	"github.com/your-github-user/lab6/models"

	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := repo.NewInMemoryRepository(logger)

	r := gin.Default()

	// Logging middleware
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

        logger.Info("Processed request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start))
	})

	// Routes
	products := r.Group("/products")
	{
		products.POST("", func(c *gin.Context) {
			var p models.Product
			if err := c.ShouldBindJSON(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

            product, err := repo.Create(p.Name, p.Price)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, product)
		})

		products.GET("", func(c *gin.Context) {
			products. err := repo.List()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, products)
		})

		products.GET("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
				return
			}

            product, err := repo.Get(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, product)
		})

		products.PUT("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
				return
			}

            var p models.Product
			if err := c.ShouldBindJSON(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

            product, err := repo.Update(id, p.Name, p.Price)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
            c.JSON(http.StatusOK, product)
		})

		products.DELETE("/:id", func(c *gin.Context) {
			id, err := uuid.Parse(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

            if err := repo.Delete(id); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
            c.Status(http.StatusNoContent)
		})
	}

	r.Run(":8080")
}
```

##### API Tests

###### Run the Server

```bash
go run cmd/api/main.go
```

##### Test with curl

- Create a product

```bash
curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name":"Laptop","price":999.99}'
```

- List products

```bash
curl http://localhost:8080/products
```

- Get a product

```bash
curl http://localhost:8080/products/1
```

- Update a product

```bash
curl -X PUT http://localhost:8080/products/1 -H "Content-Type: application/json" -d '{"name":"Laptop Pro","price":1299.99}'
```

- Delete a product

```bash
curl -X DELETE http://localhost:8080/products/1
```

</TabItem>
<TabItem value="task">

### Task

- Add validation to ensure that the product name is at least _3 characters long_.
- Implement a `middleware` for _simple authentication_ (e.g. checking an `Authorization` header)
- Add _unit tests for the endpoints_ using `httptest` and `testify`.

#### Expected output

##### Console

###### POST

- /products

```json
{ "id": "", "name": "Laptop", "price": 999.99 }
```

###### GET

- /products

```json
[{ "id": "", "name": "Laptop", "price": 999.99 }]
```

- /products/1

```json
{ "id": "", "name": "Laptop", "price": 999.99 }
```

###### PUT

- /products/1

```json
{ "id": "", "name": "Laptop Pro", "price": 1299.99 }
```

###### DELETE

- /products/1

```bash
(status 204, no body)
```

##### JSON logs (example)

```json
{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Product created","id":"1","name":"Laptop","price":999.99}
{"time":"2025-06-12T01:15:00Z","level":"INFO","msg":"Processed request","method":"POST","path":"/products","status":201,"duration":"1ms"}
```

:::info Practical use case
This API simulates an e-commerce backend, with RESTful endpoints, input validation and structured logging, ideal for integration with frontends or other services.
:::

</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered building RESTful APIs with `net/http` and `Gin`, including `routing`, `middlewares`, `binding`, `validation`, `JSON`, `status codes`, and `headers`. The hands-on lab implemented a full-featured API for CRUD, integrating `validation` and `logging`. Java engineers will notice similarities to Spring Boot, but with the lighter, more straightforward approach of Go.

:::tip Next steps
In the next module, we will explore database integration (e.g. `PostgreSQL`) and best practices for deploying APIs in production.
:::
