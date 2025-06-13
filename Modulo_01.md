# Módulo 01 - Introdução e Fundamentos da Linguagem Go

Este módulo apresenta os conceitos fundamentais da linguagem Go, incluindo sua história, características, configuração do ambiente e elementos básicos da sintaxe. O objetivo é fornecer uma base sólida para engenheiros Java que desejam aprender Go, com exemplos práticos e casos de uso.

---

## 1. História e Propósito do Go

### História

Go (Golang) foi criado em 2009 por Robert Griesemer, Rob Pike e Ken Thompson no Google, com foco em sistemas distribuídos e infraestrutura de servidores. É inspirada em C, Pascal e outras linguagens.

### Propósito

Go foi projetado para:

- **Simplicidade**: Sintaxe clara e minimalista.
- **Performance**: Compilação rápida para binários nativos.
- **Concorrência**: Suporte nativo com goroutines e channels.
- **Produtividade**: Ferramentas integradas (formatação, testes, documentação).

**Caso de uso:** Kubernetes, Docker, sistemas de alta performance no Google.

---

## 2. Características da Linguagem

### Simplicidade

- Poucas palavras-chave (25, contra 50+ em Java).
- Formatação automática com `go fmt`.
- Sem herança ou sobrecarga de métodos.

### Performance

- Compilação para binário nativo.
- Garbage collector otimizado.
- Execução rápida e eficiente.

### Concorrência

- **Goroutines**: Execução leve e eficiente.
- **Channels**: Comunicação segura entre goroutines.

**Caso de uso:** APIs RESTful e microsserviços concorrentes.

---

## 3. Instalação, Workspace e `go mod`

### Instalação

1. Baixe o Go em [go.dev](https://go.dev/dl/)
2. Siga as instruções por sistema operacional.
3. Verifique a instalação:

```bash
go version
```

### Workspace

Com Go Modules (a partir do Go 1.11), a estrutura tradicional do `$GOPATH` tornou-se opcional. Pode-se trabalhar em qualquer diretório.

### Go Modules

Gerenciador de dependências (similar ao Maven).

```bash
go mod init github.com/seu-usuario/meu-projeto
go get github.com/exemplo/pacote
```

Exemplo de `go.mod`:

```go
module github.com/seu-usuario/meu-projeto

go 1.21

require github.com/exemplo/pacote v1.0.0
```

---

## 4. Estrutura Básica de um Programa

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Execução

```bash
go run hello.go     # Executa
go build hello.go   # Compila para binário
```

**Comparação com Java:**

- Em Go, `func main()` é suficiente.
- Sem ponto e vírgula no final das linhas.

---

## 5. Tipos Primitivos, Funções, Variáveis e Constantes

### Tipos Primitivos

| Tipo              | Descrição                  | Exemplo         |
|-------------------|----------------------------|-----------------|
| int, int32, int64 | Inteiros com ou sem sinal  | `42`, `-10`     |
| float32, float64  | Ponto flutuante            | `3.14`, `-0.001`|
| bool              | Booleano                   | `true`, `false` |
| string            | Cadeia de caracteres UTF-8 | `"Hello, Go!"`  |
| byte              | Alias para uint8           | `65` (→ 'A')    |

### Variáveis

```go
var nome string = "Grok"
idade := 42
var x, y int = 10, 20
a, b := "hello", true
```

**Comparação com Java:**  
Em Java: `String nome = "Grok";`  
Em Go: `nome := "Grok"`

### Constantes

```go
const Pi float64 = 3.14159
const NomeProjeto = "MeuApp"
```

### Funções

```go
func soma(a int, b int) int {
    return a + b
}

func troca(x, y string) (string, string) {
    return y, x
}
```

---

## 📌 Lab: Hello World, Manipulação de Variáveis e Tipos

### Objetivo

Praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

### Passo a Passo

#### 1. Configuração

```bash
mkdir lab1
cd lab1
go mod init github.com/seu-usuario/lab1
```

#### 2. Hello World (arquivo `hello.go`)

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Execute:

```bash
go run hello.go
```

#### 3. Manipulação de Variáveis e Tipos (arquivo `vars.go`)

```go
package main

import "fmt"

func main() {
    var nome string = "Grok"
    idade := 42
    const versao = "1.0.0"

    var preco float64 = 99.99
    ativo := true

    fmt.Printf("Nome: %s, Idade: %d, Versão: %s\n", nome, idade, versao)
    fmt.Printf("Preço: %.2f, Ativo: %t\n", preco, ativo)

    x, y := troca("Go", "Java")
    fmt.Printf("Troca: %s, %s\n", x, y)
}

func troca(a, b string) (string, string) {
    return b, a
}
```

Execute:

```bash
go run vars.go
```

### 🧪 Tarefa:

- Crie uma função que calcule o dobro de um número.
- Adicione uma constante de IVA (`0.21`) e calcule o preço final de um produto.

**Saída esperada:**

```
Nome: Grok, Idade: 42, Versão: 1.0.0
Preço: 99.99, Ativo: true
Troca: Java, Go
```

---

## Conclusão

Este módulo cobre os fundamentos da linguagem Go, comparando com Java sempre que útil. O laboratório prático reforça o aprendizado com foco em variáveis, funções e estrutura de código.

### 🚀 Próximos passos:

No próximo módulo: estruturas de controle, slices, maps e pacotes, com exemplos práticos.
