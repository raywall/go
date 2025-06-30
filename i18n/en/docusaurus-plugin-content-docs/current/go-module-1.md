---
sidebar_position: 3
sidebar_label: Module 01
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<div className="row">
<div className="col">

# Introduction and fundamentals of Go language

This module introduces the fundamental concepts of the Go language, including its history, features, environment setup, and basic syntax elements. It aims to provide a solid foundation for Java engineers who want to learn Go, with practical examples and use cases. The content is objective but detailed, and can be refined later.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    style={{ marginTop:'80px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## History and purpose of Go

### History

Go, also known as Golang, was created in 2009 by `Robert Griesemer`, `Rob Pike` and `Ken Thompson` at Google. The language emerged to meet the needs of `large-scale development`, with `focus on distributed systems`, `server infrastructure` and `modern tooling`. Inspired by C, Pascal and other languages, `Go combines simplicity with efficiency`. #### Purpose

Go was designed for:

- **Simplicity**: Clear and minimalist syntax, reducing complexity
- **Performance**: Fast compilation to native binaries, with performance close to that of C/C++
- **Concurrency**: Native support for concurrent programming via `goroutines` and `channels`
- **Productivity**: Built-in tools (formatting, testing, documentation) to speed up development

:::info Use case
Go is widely used in projects such as Kubernetes, Docker, and high-performance systems at Google, where scalability and maintainability are critical.
:::

<br />

## Características da linguagem

<!-- SIMPLICIDADE -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher.png').default} 
    style={{ transform:'scalex(-1) scale(0.7)', marginTop:'0px' }}
    alt="A blue Go gopher" />
</div>
<div className="col">

### Simplicity

- Reduced syntax, with few keywords (25, compared to 50+ in Java)
- Automatic formatting with `go fmt`
- Absence of complex features such as inheritance or method overloading
</div>
</div>

<!-- PERFORMANCE -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-flash.png').default} 
    style={{ marginTop:'20px' }}
    alt="A gopher dressed in flash and running" />
</div>
<div className="col">

### Performance

- Compilation to native binary, eliminating the need for a virtual machine (unlike Java)
- Optimized garbage collector for low latency
- Fast execution, ideal for high-load services
</div>
</div>

<!-- CONCORRENCIA -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default} 
    style={{ transform:'scale(1)', marginTop:'5px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />
</div>
<div className="col">

### Concurrency

- **Goroutines**: Lightweight functions that perform concurrency efficiently
- **Channels**: Mechanism for safe communication between goroutines
- Unlike the thread model in Java, Go abstracts the complexity of concurrency management
</div>
</div>

:::info Use case
RESTful APIs and microservices, where Go offers high performance and concurrency to handle multiple simultaneous requests.
:::

<br />

## Instalação, workspace e `go mod`

<!-- INSTALAÇÃO -->
<div className="row">
<div className="col">

### Installation

1. Download Go from [go.dev](https://go.dev/dl/) for your operating system
2. Install following the specific instructions:

- **Linux or MAC**: Unzip the file and add it to `$PATH`
- **Windows**: Run the installer and check with `go version`
</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer.png').default} 
    style={{ marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>

:::tip Note
On `MacOS` operating system it is also possible to install Go using `brew install go`
:::

4. Verify the installation using the command:

```bash
go version
```

<br />

### Workspace

<!-- WORKSPACE -->
<div className="row">
<div className="col">
Go uses a directory structure to organize projects. Starting with Go 1.11, Go Modules replaced the old $GOPATH as the default.

#### Main directories (optional, with Go Modules)

- `src/`: Source code
- `bin/`: Compiled binaries
- `pkg/`: Shared packages

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer-2.png').default} 
    style={{ transform:'scale(1.1)', marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>

:::tip Note
With Go Modules, you can work in any directory
:::

<br />

### Go Modules

Go Modules manage dependencies similarly to Maven in Java.

1. Initialize a module:

```bash
go mod init github.com/my-github-user/my-project
```

2. Add dependencies:

```bash
go get github.com/sample/package
```

3. The `go.mod` file stores dependencies and versions

#### Example of `go.mod`

```bash
module github.com/my-github-user/my-project

go 1.21

require github.com/sample/package v1.0.0
```

:::info Use Case
Go Modules make it easier to maintain large projects, similar to managing dependencies in Java with Maven or Gradle.
:::

<br />

## Basic structure of a Go program

A Go program follows a simple structure, with the main package as the entry point.

### Example

<div className="row">
<div className="col">

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> Click [here](@site/static/code/mod1/lab/hello-world.go) to download this source code!

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-hello-world.png').default} 
    style={{ transform:'scale(1)', marginTop:'-70px' }}
    alt="A blue Go gopher" />
</div>
</div>

- **package main**: Defines the main package, which generates an executable.
- **import "fmt"**: Imports the fmt package for formatting and output.
- **func main()**: Entry function, equivalent to public static void main in Java.

<br />

## Compilation and execution:

```bash
go run hello.go     # Execute directly
go build hello.go   # Compiles to a binary
```

<br />

## Comparison with Java

- In Java, static classes and methods are mandatory. In Go, the main function is sufficient.
- Go does not use semicolons (;) at the end of lines.

<br />

## Primitive Types, Functions, Variables and Constants

### Primitive Types

Go has simple, statically typed primitive types (similar to Java, but more concise).

<div className="row">
<div className="col">

| Tipo              | Description                     | Exemplo                |
| ----------------- | ------------------------------- | ---------------------- |
| int, int32, int64 | Integers with or without signal | 42, -10                |
| float32, float64  | Floating point numbers          | 3.14, -0.001           |
| bool              | Boolean                         | true, false            |
| string            | Character chain UTF-8           | "Hello, Go!"           |
| byte              | Alias ​​for uint8               | 65 (equivalent to ‘A’) |

</div>
<div className="col col--4 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-primitive.png').default} 
    style={{ transform:'scale(1.4)', marginTop:'10px' }}
    alt="A blue Go gopher" />
</div>
</div>
:::tip Note
Go does not support implicit types like generic var in Java. Type inference is done with :=.
:::

### Variables

- Explicit declaration:

```go
    var name string = "Raywall"
    var age int = 45
```

- Short statement (type inference):

```go
    name := "Raywall"
    age := 45
```

- Multiple variables:

```go
    var x, y int = 10, 20
    a, b := "hello", true
```

### Comparison with Java

#### Java

```java
    String name = "Raywall";
```

#### Go

```go
    name := "Raywall" // more concise, but strictly typed
```

### Constants

Constants are defined with const and cannot be changed.

```go
const Pi float64 = 3.14159
const ProjectName = "MyApp"
```

### Functions

Functions in Go are declared with func, can return multiple values ​​and do not support overloading (unlike Java).

#### Example

```go
package main

import "fmt"

func sum(a int, b int) int {
    return a + b
}

func replace(x, y string) (string, string) {
    return y, x
}

func main() {
    fmt.Println(sum(5, 3)) // Output: 8
    a, b := replace("hello", "world")
    fmt.Println(a, b) // Output: world hello
}
```

> Click [here](@site/static/code/mod1/funcoes.go) to download this source code!

:::info Use case
Functions with multiple returns are useful for error handling, such as value, err := function().
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
{label: 'Exercise 1', value: 'hello'},
{label: 'Exercise 2', value: 'vars'},
{label: 'Task', value: 'task'},
]}>
<TabItem value="config">

### Setup

1. Install Go and check with `go version`
2. Create a `lab1` directory:

```bash
mkdir lab1
```

3. Access the directory and initialize a module:

```bash
cd lab1
go mod init github.com/your-github-user/lab1
```

</TabItem>
<TabItem value="hello">

### Hello, World!

#### Goal

Create a simple program in Go to practice setting up the environment, basic structure, types, variables and functions.

#### Step by step

1. Create the `hello.go` file with the following content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> Click [here](@site/static/code/mod1/lab/hello-world.go) to download this source code!

##### Running

```bash
go run hello.go
```

</TabItem>
<TabItem value="vars">

### Variable and type manipulation

#### Goal

Create a simple Go program to practice setting up the environment, basic structure, types, variables and functions.

#### Step by step

1. Create a `vars.go` file with the following content:

```go
package main

import "fmt"

func main() {
    // Variable declaration
    var name string = "Raywall"
    age := 45

    const version = "1.0.0"

    // Types
    var price float64 = 99.99
    active := true

    // Display
    fmt.Printf("Name: %s, Age: %d, Version: %s\n", name, age, version)
    fmt.Printf("Price: %.2f, Active: %t\n", price, active)

    // Function with multiple returns
    x, y := replace("Go", "Java")
    fmt.Printf("Replaced: %s, %s\n", x, y)
}

func replace(a, b string) (string, string) {
    return b, a
}
```

> Click [here](@site/static/code/mod1/lab/vars.go) to download this source code!

##### Running

```bash
    go run vars.go
```

</TabItem>
<TabItem value="task">

### Task

- Modify `vars.go` to include a function that calculates the double of a number
- Add a constant for VAT (Value Added Tax, e.g. 0.21) and calculate the final price of a product

#### Expected output

```bash
Name: Raywall, Age: 45, Version: 1.0.0
Price: 99.99, Active: true
Replaced: Java, Go
```

:::info Pratical use case
This lab simulates data manipulation in a simple system, such as a product registry, where variables and constants are used to store fixed (version) and dynamic (price, status) information.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusion

This module covers the fundamentals of Go, from its history to creating basic programs. Java engineers will notice similarities (static typing, functions) and differences (concise syntax, lack of classes). The hands-on lab reinforces environment setup and the use of types, variables, and functions, preparing the team for more advanced modules, such as control structures and concurrency.

:::tip Next steps
In the next module, we will explore `control structures`, `slices`, `maps`, and `packages`, with more practical examples.
:::
