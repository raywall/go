---
sidebar_position: 4
sidebar_label: Module 02
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Control structures and types composed in GO

<div className="row">
<div className="col">

This module deepens the GO fundamentals, addressing `control structures`, compound types (`arrays`,` slices`, `maps`),` structts` and `pointers`. The content is aimed at Java engineers, with practical examples and objective use cases, maintaining depth for future consultation.

Practical Lab implements a CRUD (Create, Read, Update and Delete) in memory, reinforcing the learned concepts.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-10px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## Control structures

Go has simple control structures, without mandatory parentheses (unlike Java) and focused on clarity.

### `if`

- Supports boot in the statement itself.
- There is no ternary operator in Go.

#### Example

```go
package main

import "fmt"

func main() {
	number := 45
	if x := numger % 2; x == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
```

> Click [here](@site/static/code/mod2/controle-if.go) to download this source code!

#### Comparison with Java:

##### Java

```java
if (x % 2 == 0) { ... }
```

##### Go

```go
if x := number % 2; x == 0 { ... } // direct inicialization
```

### `for`

- Only Go loop structure (there is no `While` or` Do-While`)
- Supports three forms: `traditional`, `simple` and `infinite` condition

#### Exemplo

```go
package main

import "fmt"

func main() {
	// Traditional loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// With while-like loop
	counter := 0
	for counter < 3 {
		fmt.Println(counter)
		counter++
	}

	// Infinite (with break)
	for {
		fmt.Println("Infinite")
		break
	}
}
```

> Click [here](@site/static/code/mod2/controle-for.go) to download this source code!

### `switch`

- Simpler than Java: no need of break (no fall-through by default)
- Supports multiple conditions per case and initialization

#### Example

```go
package main

import "fmt"

func main() {
    day := 3
    switch day {
        case 1, 2:
            fmt.Println("Beginning of the week")
        case 3, 4, 5:
            fmt.Println("Midweek")
        default:
            fmt.Println("Weekend")
    }

    // Switch whith expression
    switch x := dia * 2; x {
    case 6:
        fmt.Println("Day 3 doubled")
    default:
        fmt.Println("Another value")
    }
}
```

> Click [here](@site/static/code/mod2/controle-switch.go) to download this source code!

### `defer`

- Defers the execution of a statement until the end of the function
- Useful for freeing resources (similar to `try-with-resources` in Java)

#### Example

```go
package main

import "fmt"

func main() {
	defer fmt.Println("Executed at the end")
	fmt.Println("Executed first")
}
```

> Click [here](@site/static/code/mod2/controle-defer.go) to download this source code!

:::info Use case
`defer` is used to close files, bank connections or release Mutexes
:::

<br />

## Arrays, slices and maps

### Arrays

- Fixed size, defined in the declaration
- Less common in Go, as slices are more flexible

#### Example

```go
var numbers [3]int = [3]int{1, 2, 3}
fmt.Println(numbers[0]) // Output: 1
```

#### Comparison with Java

Arrays in Go are similar to Java arrays (`int[]`), but less used due to slices

### Slices

- Dynamic structure, based on arrays, but with variable size
- Declared with `[]type` or created with make

#### Example

```go
package main

import "fmt"

func main() {
	// Declaration
	slice := []int{1, 2, 3}
	fmt.Println(slice) // [1 2 3]

	// Add elements
	slice = append(slice, 4)
	fmt.Println(slice) // [1 2 3 4]

	// Slice from Array
	array := [5]int{1, 2, 3, 4, 5}
	subSlice := array[1:4] // [2 3 4]
	fmt.Println(subSlice)
}
```

> Click [here](@site/static/code/mod2/slices.go) to download this source code!

:::tip Note
Slices are passed by reference, but the underlying array can be copied if modified
:::

### Maps

- `Key-value` structure, similar to `HashMap` in Java
- Declared with `map[KeyType]ValueType`

#### Example

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["age"] = 45
	m["price"] = 99
	fmt.Println(m) // map[age:42 price:99]

	// Check if exists
	value, exists := m["age"]
	if exists {
		fmt.Println("Age:", value)
	}

	// Delete
	delete(m, "price")
	fmt.Println(m) // map[age:42]
}
```

> Click [here](@site/static/code/mod2/maps.go) to download this source code!

:::info Use case
Slices are ideal for dynamic lists (eg user list), while maps are useful for fast associations (eg: caching IDs to values)
:::

### Structs and struct tags

#### Structs

- Equivalent to classes in Java, but without inheritance
- Defined with `type Name struct`

##### Example

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age int
}

func main() {
	p := Person{Name: "Raywall", Age: 45}
	fmt.Println(p) // {Raywall 45}

	// Access to fields
	fmt.Println(p.Name) // Raywall

	// Anonymous struct
	anon := struct {
		Value string
	}{Value: "Test"}

	fmt.Println(anon) // {Test}
}
```

> Click [here](@site/static/code/mod2/structs.go) to download this source code!

#### Struct tags

- Used for metadata, such as JSON serialization
- Declared with phrases (`\``)

##### Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

func main() {
	p := Product{ID: 1, Nome: "Laptop"}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData)) // {"id":1,"name":"Laptop"}
}
```

> Click [here](@site/static/code/mod2/tags-de-struct.go) to download the source code!

##### Comparison with Java

Structs replace classes, but tags are similar to Jackson's annotations (`@JsonProperty`) in Java

:::info Use case
Structs with tags are widely used in REST APIs to `serialize`/`deserialize` JSON
:::

<br />

## Pointers

- Go supports pointers for direct memory manipulation, but in a safe way
- Declared with `*type` (address) and `&variable` (reference)

### Example

```go
package main

import "fmt"

func main() {
	x := 45
	p := &x         // Pointer to x
	fmt.Println(*p) // 45 (dereferencing)

	// Modify via pointer
	*p = 100
	fmt.Println(x) // 100
}

func increase(p *int) {
	*p++
}
```

> Click [here](@site/static/code/mod2/ponteiros.go) to download this source code!

### Comparison with Java

- Java uses implicit references to objects, without direct control of pointers
- Go requires explicit pointers for changes in functions (passing by value is the default)

:::info Use case
Pointers are useful for modifying large structs without copying data, similar to passing objects by reference in Java.
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
2. Create a `lab2` directory:

```bash
mkdir lab2
```

3. Access the directory and initialize a module:

```bash
cd lab2
go mod init github.com/your-github-user/lab2
```

</TabItem>
<TabItem value="app">

### Implement the functions of a CRUD (Create, Read, Update and Delete) in memory

#### Goal

Create a CRUD capable of managing a list of products in memory, using `slices`, `maps`, `structs` and `control structures`. The lab simulates a CRUD (`Create`, `Read`, `Update`, `Delete`).

#### Step by step

1. Create a `crud.go` file with the following code:

```go
package main

import (
	"errors"
	"fmt"
)

// Product struct
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Database in memory (product slice)
var produtos []Product
var idCounter = 1

// Create: Add a new product to the database
func addProduct(name string, prico float64) Product {
	product := Product{ID: idCounter, Name: name, Price: price}
	idCounter++

	products = append(products, product)
	return product
}

// Read: Search all products
func listProducts() []Product {
	return products
}

// Read: Search a product by ID
func getProduct(id int) (Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return Product{}, errors.New("product not found")
}

// Update: Updates a product
func updateProduct(id int, name string, price float64) (Product, error) {
	for i, p := range products {
		if p.ID == id {
			products[i] = Product{ID: id, Name: name, Price: price}
			return products[i], nil
		}
	}

	return Product{}, errors.New("product not found")
}

// Delete: Removes a product
func deleteProduct(id int) error {
	for i, p := range products {
		if p.ID == id {
			produtos = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

func main() {
	// Create a product
	addProduct("Laptop", 999.99)
	addProduct("Mouse", 29.99)

	// List products
	fmt.Println("List of products:")
	for _, p := range listProducts() {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// Get a product
	if p, err := getProduct(1); err == nil {
		fmt.Printf("Product founded: %+v\n", p)
	}

	// Update a product
	if p, err := updateProduct(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto updated: %+v\n", p)
	}

	// Delete a product
	if err := deleteProduct(2); err == nil {
		fmt.Println("Product deleted successfully")
	}

	// Listing products again
	fmt.Println("Final list:")

	for _, p := range listProducts() {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}
```

> Click [here](@site/static/code/mod2/lab/crud.go) to download this source code!

#### Running

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Task

- Add validation to prevent negative prices on createProduct and updateProduct
- Implement a function that lists products with a price above a specific value
- Use a map to store products by ID, instead of a slice, and compare performance

#### Output expected

```bash
    List of products:
    ID: 1, Name: Laptop, Price: 999.99
    ID: 2, Name: Mouse, Price: 29.99

    Product founded: {ID:1 Name:Laptop Price:999.99}
    Product updated: {ID:1 Name:Laptop Pro Price:1299.99}
    Product deleted successfully

    Final list:
    ID: 1, Name: Laptop Pro, Price: 1299.99
```

:::info Practical Use Case
This lab simulates a simple product management API, similar to an e-commerce backend. Slices and maps are used for in-memory data manipulation, while structs model entities.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covered control structures (`if`, `for`, `switch`, `defer`), composite types (`arrays`, `slices`, `maps`), `tagged structs and pointers`. The hands-on lab reinforces the application of these concepts in a realistic CRUD scenario. Java engineers will notice similarities with collections (`list`, `map`) and differences in the absence of inheritance or generics (up to Go 1.18).

:::tip Next steps
In the next module, we will explore concurrency with `goroutines` and `channels`, as well as `error handling` and `packages`, preparing the team to develop scalable applications in Go.
:::
