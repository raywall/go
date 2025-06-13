# Módulo 02 - Estruturas de Controle e Tipos Compostos em Go

Este módulo aprofunda os fundamentos do Go, com foco em estruturas de controle (`if`, `for`, `switch`, `defer`), tipos compostos (`arrays`, `slices`, `maps`), `structs` com tags e ponteiros. Feito sob medida para engenheiros Java, traz exemplos diretos, comparações e um lab prático simulando uma API CRUD em memória.

---

## ✅ 1. Estruturas de Controle

### `if`

- Não usa parênteses.
- Pode incluir inicialização na própria declaração.
- Não existe operador ternário (`?:`) em Go.

```go
if x := numero % 2; x == 0 {
    fmt.Println("Par")
} else {
    fmt.Println("Ímpar")
}
```

> 🔁 **Java vs Go:**  
> Java: `if (x % 2 == 0) { ... }`  
> Go: `if x := numero % 2; x == 0 { ... }`

---

### `for`

Go tem **apenas uma estrutura de repetição**, o `for`, com três formas:

```go
// Tradicional
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

// Tipo while
contador := 0
for contador < 3 {
    fmt.Println(contador)
    contador++
}

// Loop infinito
for {
    fmt.Println("Executa até break")
    break
}
```

---

### `switch`

- Mais simples que em Java.
- Não há fall-through por padrão.
- Suporta múltiplas condições por caso.

```go
switch dia {
case 1, 2:
    fmt.Println("Início da semana")
case 3, 4, 5:
    fmt.Println("Meio da semana")
default:
    fmt.Println("Fim de semana")
}
```

```go
// Switch com expressão e inicialização
switch x := dia * 2; x {
case 6:
    fmt.Println("Dia 3 dobrado")
default:
    fmt.Println("Outro valor")
}
```

---

### `defer`

- Adia a execução até o final da função.
- Muito usado para fechar arquivos, conexões ou liberar recursos.

```go
defer fmt.Println("Executado no final")
fmt.Println("Executado primeiro")
```

> 🔄 **Semelhante ao `try-with-resources` do Java.**

---

## 🧱 2. Arrays, Slices e Maps

### Arrays

- Tamanho fixo.
- Menos comuns que slices.

```go
var numeros [3]int = [3]int{1, 2, 3}
fmt.Println(numeros[0]) // 1
```

> 🔍 Semelhantes a `int[]` em Java.

---

### Slices

- Tamanho variável.
- Baseados em arrays.
- Criados com `[]tipo` ou `make`.

```go
slice := []int{1, 2, 3}
slice = append(slice, 4) // [1 2 3 4]

array := [5]int{1, 2, 3, 4, 5}
subSlice := array[1:4]   // [2 3 4]
```

> ✅ Slices são passados por referência.

---

### Maps

- Estrutura chave-valor (como `HashMap`).

```go
m := make(map[string]int)
m["idade"] = 42

valor, existe := m["idade"]
if existe {
    fmt.Println("Idade:", valor)
}

delete(m, "idade")
```

> 🧠 Use maps para buscas rápidas e caches.

---

## 🧩 3. Structs e Tags

### Structs

- Substituem classes.
- Não possuem herança.

```go
type Pessoa struct {
    Nome  string
    Idade int
}

p := Pessoa{Nome: "Grok", Idade: 42}
fmt.Println(p.Nome) // Grok
```

### Struct Anônima

```go
anon := struct {
    Valor string
}{Valor: "Teste"}
fmt.Println(anon) // {Teste}
```

---

### Tags de Struct (JSON)

```go
type Produto struct {
    ID    int     `json:"id"`
    Nome  string  `json:"nome"`
    Preco float64 `json:"preco,omitempty"`
}
```

```go
jsonData, _ := json.Marshal(produto)
fmt.Println(string(jsonData))
```

> 📌 Tags funcionam como anotações `@JsonProperty` do Jackson.

---

## 🧠 4. Ponteiros

- Go suporta ponteiros (de forma segura).
- Usados para modificar valores fora do escopo local.

```go
x := 42
p := &x
fmt.Println(*p) // 42

*p = 100
fmt.Println(x) // 100
```

```go
func incrementar(p *int) {
    *p++
}
```

> 📘 Java usa referências implícitas. Em Go, o uso de ponteiros é explícito.

---

## 🧪 Lab 2: API Fake de Produtos (CRUD)

### 🎯 Objetivo

Criar uma API em memória simulando operações de **Create, Read, Update, Delete** com `slices`, `maps`, `structs` e controle de fluxo.

---

### 🚧 Setup

```bash
mkdir lab2
cd lab2
go mod init github.com/seu-usuario/lab2
```

---

### 📁 Código (crud.go)

```go
package main

import (
    "errors"
    "fmt"
)

type Produto struct {
    ID    int     `json:"id"`
    Nome  string  `json:"nome"`
    Preco float64 `json:"preco"`
}

var produtos []Produto
var idCounter = 1

func criarProduto(nome string, preco float64) Produto {
    produto := Produto{ID: idCounter, Nome: nome, Preco: preco}
    idCounter++
    produtos = append(produtos, produto)
    return produto
}

func listarProdutos() []Produto {
    return produtos
}

func buscarProduto(id int) (Produto, error) {
    for _, p := range produtos {
        if p.ID == id {
            return p, nil
        }
    }
    return Produto{}, errors.New("Produto não encontrado")
}

func atualizarProduto(id int, nome string, preco float64) (Produto, error) {
    for i, p := range produtos {
        if p.ID == id {
            produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
            return produtos[i], nil
        }
    }
    return Produto{}, errors.New("Produto não encontrado")
}

func deletarProduto(id int) error {
    for i, p := range produtos {
        if p.ID == id {
            produtos = append(produtos[:i], produtos[i+1:]...)
            return nil
        }
    }
    return errors.New("Produto não encontrado")
}

func main() {
    criarProduto("Laptop", 999.99)
    criarProduto("Mouse", 29.99)

    fmt.Println("Lista de produtos:")
    for _, p := range listarProdutos() {
        fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
    }

    if p, err := buscarProduto(1); err == nil {
        fmt.Printf("Produto encontrado: %+v\n", p)
    }

    if p, err := atualizarProduto(1, "Laptop Pro", 1299.99); err == nil {
        fmt.Printf("Produto atualizado: %+v\n", p)
    }

    if err := deletarProduto(2); err == nil {
        fmt.Println("Produto deletado com sucesso")
    }

    fmt.Println("Lista final:")
    for _, p := range listarProdutos() {
        fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
    }
}
```

---

### 💡 Tarefas adicionais

- ❌ Bloquear preços negativos.
- 🔍 Listar produtos com preço acima de um valor.
- 🔁 Substituir slice por `map[int]Produto` e comparar performance.

---

### 📌 Caso de uso

Simulação realista de backend para e-commerce com dados em memória. Reforça conceitos essenciais de Go aplicados a APIs REST.

---

## 🧭 Conclusão

Neste módulo, vimos:

- Estruturas de controle: `if`, `for`, `switch`, `defer`.
- Tipos compostos: `arrays`, `slices`, `maps`.
- Structs com tags e ponteiros.
- Um CRUD funcional para fixar os conceitos.

> Próximo passo: Módulo 3 – **Concorrência com goroutines e channels, tratamento de erros e pacotes.**
