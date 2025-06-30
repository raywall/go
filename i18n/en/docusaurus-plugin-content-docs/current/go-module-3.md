---
sidebar_position: 5
sidebar_label: Module 03
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Functions, methods and interfaces

<div className="row">
<div className="col">

This module explores advanced functions, struct methods, interfaces, and best practices in Go, with a focus on Java engineers who want to adopt the language's idiomatic style. The content is detailed but objective, with examples and use cases for future reference.

The hands-on lab refactors the CRUD of [Module 02](go-module-2), introducing interfaces for repositories.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ transform:'scale(1.1)', marginTop:'-10px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Functions with multiple returns

Go allows functions to return multiple values, a powerful feature for error handling and compounding results, unlike Java, which uses exceptions or objects.

### Example

```go
package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("zero division")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result) // Result: 5
}
```

> Click [here](@site/static/code/mod3/funcao-multiplos-retornos.go) to download this source code!

:::info Use case
Functions with multiple returns are common in Go to return a value and an error, replacing the use of exceptions as in Java (`try-catch`).
:::

<br />

## Anonymous functions and closures

### Anonymous functions

- Unnamed functions, defined inline, similar to lambdas in Java
- Useful for temporary tasks or callbacks

#### Example

```go
package main

import "fmt"

func main() {
    sum := func(a, b int) int {
        return a + b
    }

    fmt.Println(sum(3, 4)) // Output: 7
}
```

> Click [here](@site/static/code/mod3/funcao-anonima.go) to download this source code!

### Closures

- Anonymous functions that capture variables from the outer scope
- Similar to `closures` in Java (e.g. lambdas with access to outer variables)

#### Example

```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := counter()
    fmt.Println(c1()) // Output: 1
    fmt.Println(c1()) // Output: 2

    c2 := contador()
    fmt.Println(c2()) // Output: 1 (new counter)
}
```

> Click [here](@site/static/code/mod3/closures.go) to download this source code!

:::info Use case
Closures are useful for maintaining state in functions, such as counters or generators, without the need for `structs`.
:::

## Methods in structs

- Go does not have classes, but methods can be associated with structs using receivers
- Receivers can be by value (`T`) or by pointer (`*T`)

### Example

```go
package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

// Método com receiver por valor
func (p Product) Description() string {
	return fmt.Sprintf("%s: R$ %.2f", p.Name, p.Price)
}

// Method with receiver by pointer
func (p *Product) ApplyDiscount(discount float64) {
	p.Price -= p.Price * discount
}

func main() {
	p := Product{Name: "Laptop", Price: 1000}

	fmt.Println(p.Description()) // Output: Laptop: R$ 1000.00
	p.ApplyDiscount(0.1)

	fmt.Println(p.Description()) // Output: Laptop: R$ 900.00
}
```

> Click [here](@site/static/code/mod3/metodos-em-structs.go) to download this source code!

### Comparison with Java

- In Java, methods are defined inside classes. In Go, methods are functions with receivers, associated with types (usually structs).
- Receiver by pointer (`*T`) is similar to modifying the state of an object in Java.

:::info Use case
Methods in structs are used to encapsulate behavior, such as formatting data or applying business rules.
:::

## Interfaces and duck typing

### Interfaces

- Defined with type Interface name, specifying methods that a type must implement.
- **Go uses `duck typing`**: if a type implements the methods of an interface, it satisfies it automatically (without explicit declaration, unlike Java with implements).

#### Example

```go
package main

import "fmt"

type Describable interface {
	Description() string
}

type Product struct {
	Name  string
	Price float64
}

type Service struct {
	Name string
	Tax float64
}

func (p Product) Description() string {
	return fmt.Sprintf("Product: %s, R$ %.2f", p.Name, p.Price)
}

func (s Service) Description() string {
	return fmt.Sprintf("Service: %s, R$ %.2f", s.Name, s.Tax)
}

func show(d Describable) {
	fmt.Println(d.Description())
}

func main() {
	p := Product{Name: "Laptop", Price: 1000}
	s := Service{Name: "Consultoria", Tax: 500}

	show(p) // Output: Product: Laptop, R$ 1000.00
	show(s) // Output: Service: Consultoria, R$ 500.00
}
```

> Click [here](@site/static/code/mod3/interfaces.go) to download this source code!

#### Comparison with Java

- In Java, interfaces require `implements`. In Go, implementation is implicit.
- Go does not support inheritance, but interfaces allow polymorphism.

### Empty interface (`interface{}`)

- Equivalent to `Object` in Java, accepts any type.
- Used sparingly, usually with type assertions or type switches.

#### Example

```go
func inspect(v interface{}) {
    switch t := v.(type) {
    case Product:
        fmt.Println("Is a Product:", t.Name)
    case string:
        fmt.Println("Is a string:", t)
    default:
        fmt.Println("Unknown type")
    }
}
```

> Click [here](@site/static/code/mod3/interface-vazia.go) to download this source code!

:::info Use case
Interfaces are widely used in Go to abstract repositories, services or plugins, ensuring flexibility and testability.
:::

<br />

## Good practices and idiomatic design principles in Go

### Simplicity

- Prefer clear and concise solutions
- Avoid unnecessary complexity (e.g. excessive layers of abstraction)

### Naming

- Use short and descriptive names (e.g. p for Product, instead of `productInstance`)
- Exported functions and methods start with a capital letter (e.g. `Description`)

### Error handling

- Always check for returned errors (`if err != nil`)
- Avoid panics in production code, except in extreme cases.

### Small interfaces

- Define interfaces with few methods, focused on a single responsibility.
- **Example**: `io.Reader` and `io.Writer` instead of generic interfaces.

### Composition over inheritance

- Use `nested structs` for composition, instead of inheritance like in Java.

### Avoid empty interface (`interface{}`)

- Use type assertions or type switches only when necessary

### Use Go idiomatically

- Avoid Java patterns (e.g. unnecessary `getters`/`setters`)
- Leverage `go fmt` and tools like `golint` for consistency

:::info Use case
Good practices ensure readable and maintainable code, especially in large teams, such as microservices projects.
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

1. Install Go and check with `go version`
2. Create a `lab3` directory:

```bash
mkdir lab3
```

3. Access the directory and initialize a module:

```bash
cd lab3
go mod init github.com/your-github-user/lab3
```

</TabItem>
<TabItem value="app">

### Refactor CRUD using repository interfaces

#### Goal

Refactor CRUD from Module 2 to use a repository interface, abstracting data storage and making testing and maintenance easier.

#### Step by step

1. Create a `crud.go` file with the following code:

```go
package main

import (
	"errors"
	"fmt"
)

// Structure for Product
type Product struct {
	ID 	  int     `json:"id"`
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
}

// Interface methods
func (r *InMemoryRepository) Create(name string, price float64) (Product, error) {
	if price < 0 {
		return Produto{}, errors.New("price cannot be negative")
	}

	r.idCounter++
	product := Produto{ID: r.idCounter, Name: name, Price: price}

	r.products = append(r.products, product)
	return product, nil
}

func (r *InMemoryRepository) Get(id int) (Product, error)
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}

	return Product{}, errors.New("product not found")
}

func (r *InMemoryRepository) List() ([]Product, error) {
	return r.products, nil
}

func (r *InMemoryRepository) Update(id int, name string, price float64) (Product, error) {
	if price < 0 {
		return Product{}, errors.New("price cannot be negative")
	}

	for i, p := range r.products {
		if p.ID == id {
			r.products[i] = Product{ID: id, Name: name, Price: price}
			return r.products[i], nil
		}
	}

	return Product{}, errors.New("product not found")
}

func (r *InMemoryRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}

// Function to display products
func showProducts(repo ProductRepository) {
	products, _ := repo.List()
	fmt.Println("List of products:")

	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

func main() {
	// Initialize repository
	repo := &InMemoryRepository{}

	// Create products
	repo.Create("Laptop", 999.99)
	repo.Create("Mouse", 29.99)

	// List Products
	showProducts(repo)

	// Get a product
	if p, err := repo.Get(1); err == nil {
		fmt.Printf("Product founded: %+v\n", p)
	}

	// Update a product
	if p, err := repo.Update(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Product updated: %+v\n", p)
	}

	// Delete a product
	if err := repo.Delete(2); err == nil {
		fmt.Println("Product deleted successfully")
	}

	// List products again
	exibirProdutos(repo)
}
```

> Click [here](@site/static/code/mod3/lab/crud.go) to download this source code!

##### Running

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Task

- Create a second repository (e.g. `RepositorioMock`) that implements `RepositorioProdutos` for testing, returning fixed data.
- Add an anonymous function in main to filter products with a price above a certain value
- Use a method in Produto to calculate the price with tax (e.g. 21%)

#### Expected output

```bash
List of products:
ID: 1, Name: Laptop, Price: 999.99
ID: 2, Name: Mouse, Price: 29.99

Product founded: {ID:1 Name:Laptop Price:999.99}
Product updated: {ID:1 Name:Laptop Pro Price:1299.99}
Product deleted successfully

List of products::
ID: 1, Name: Laptop Pro, Price: 1299.99
```

:::info Practical use case
The RepositorioProdutos interface abstracts the storage, allowing to change the implementation (e.g. from memory to database) without changing the main code, similar to the Repository pattern in Java.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered functions with multiple returns, anonymous functions, closures, methods in structs, interfaces, and idiomatic best practices in Go. The hands-on lab refactored the CRUD from [Module 02](go-module-2.md), introducing interfaces for greater modularity. Java engineers will notice similarities with interfaces and lambdas, but with the simpler, more implicit approach of Go.

:::tip Next steps
In the next module, we will explore concurrency with `goroutines` and `channels`, as well as `advanced error handling and testing`, preparing the team for scalable and robust applications in Go.
:::
