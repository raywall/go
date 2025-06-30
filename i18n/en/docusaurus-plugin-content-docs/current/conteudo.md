---
sidebar_position: 2
sidebar_label: Program
---

# Program

## [Module 01](go-module-1.md) – Introduction and fundamentals of Go language

<div className="row">
<div className="col">
- History and purpose of Go
- Language features (simplicity, performance, concurrency)
- Installation, workspace, `go mod`
- Basic structure of a program
- Primitive types, functions, variables, constants

📌 Lab: Hello world, manipulation of variables and types, first program with `go run`.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    alt="A diaper brown gopher" />
</div>
</div>
---

## [Module 02](go-module-2.md) – Control structures and composite types

<div className="row">
<div className="col">
- `if`, `for`, `switch`, `defer`
- Arrays, slices, maps
- Structs and struct tags
- Pointers (concept and application)

📌 Lab: Implement a CRUD (Create, Read, Update and Delete) in memory using slices and maps.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-meeting.png').default} 
    style={{ transform:'scalex(-1)', marginTop:'-30px' }}
    alt="Three gophers in a meeting, where one is explaining graphics on a slate while the others watch sitting on a round table" />
</div>
</div>
---

## [Module 03](go-module-3.md) – Functions, methods and interfaces

<div className="row">
<div className="col">
- Returns with multiple functions
- Anonymous and closure functions
- Structs methods
- Interfaces and duck typing
- Best practices and principles of idiomatic design in GO

📌 Lab: Refactor CRUD using interfaces for repositories.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-flow.png').default} 
    style={{ marginTop:'-10px' }}
    alt="A gopher designing a flow diagram" />
</div>
</div>
---

## [Module 04](go-module-4.md) – Error handling

<div className="row">
<div className="col">

- GO philosophy: explicit errors
- Default error, `errors.New`, `fmt.Errorf`
- Wrapping and unwrapping with `errors.Is`, `errors.As`
- Packages `log` and `log/slog`

📌 Lab: Create functions with structured errors and logging.

</div>
<div className="col col--3 text--left" style={{ paddingTop: '8px' }}>
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scalex(1)', marginTop:'-60px' }}
    alt="A clumsy gopher, wearing a brown sweatshirt and dark green pants, carrying some roles and having a coffee" />
</div>
</div>
---

## [Module 05](go-module-5.md) – Concurrency with goroutines and channels

<div className="row">
<div className="col">

- Goroutines: what they are and how to use them
- Channels (`unbuffered`, `buffered`)
- `select`, `sync.WaitGroup`, `sync.Mutex`
- Competition patterns in GO

📌 Lab: Create a pool worker for competing task processing.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default}
    style={{ transform:'scale(1.1)', marginTop:'-30px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />

</div>
</div>
---

## [Module 06](go-module-6.md) – Packages, modules and code organization

<div className="row">
<div className="col">

- Idiomatic package structure
- Design conventions (`cmd`, `internal`, `pkg`)
- `go mod` and versioning
- Dependency management with `go get`, `replace`

📌 Lab: Organize CRUD project into multiple packages with `go mod`.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default}
    style={{ transform:'scale(0.9)', marginTop:'-30px' }}
    alt="A brown gopher with glasses and white shirt, organizing boxes in a warehouse" />

</div>
</div>
---

## [Module 07](go-module-7.md) – Testing and code quality

<div className="row">
<div className="col">

- Testing with testing
- Unit and integration tests
- Testing with mocks (`testify`, `gomock`)
- Benchmarks and profiling
- Tools: `go vet`, `golint` and `staticcheck`

📌 Lab: Create unit and integration tests for CRUD with error coverage.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default}
    style={{ transform:'scale(1.1)', marginTop:'25px' }}
    alt="A gopher with a box of small pieces and a paper where it fits them with the name Schema.go written at the top" />

</div>
</div>
---

## [Module 08](go-module-8.md) – Web APIs with net/http and Gin

<div className="row">
<div className="col">

- HTTP server with `net/http`
- Middlewares and handlers
- Gin framework: `routing`, `binding`, `validation`
- JSON, status codes and headers

📌 Lab: Implement a RESTful API with Gin + validation.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default}
    style={{ transform:'scale(1.5)', marginTop:'-10px' }}
    alt="A gopher by programming on a computer, supported by a box full of books" />

</div>
</div>
---

## [Module 09](go-module-9.md) – Database persistence

<div className="row">
<div className="col">

- Drivers and `database/sql`
- ORM with `gorm`
- Migrations with `golang-migrate`
- Repositories and integration tests with DB

📌 Lab: Persist CRUD in a real database (PostgreSQL for example).

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default}
    style={{ transform:'scale(1.1)' }}
    alt="A Gopher teaching SQL class in a board, with the elephant, symbol of the PostgreSQL next door" />

</div>
</div>
---

## [Module 10](go-module-10.md) – Deployment, observability and best practices

<div className="row">
<div className="col">

- Build with `go build`, cross-compilation
- Docker with Go
- Structured logging (`slog`, `zap`)
- Tracing with OpenTelemetry
- `Linter`, coverage, automatic documentation with `godoc`

📌 Lab: Containerize the service and expose metrics/trace/logs.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default}
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-15px' }}
    alt="A blue gopher with detective hat by analyzing a note with a magnifying glass" />
</div>
</div>
