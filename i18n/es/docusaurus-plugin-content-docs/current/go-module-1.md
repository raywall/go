---
sidebar_position: 3
sidebar_label: Módulo 01
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

:::warning Atención
Este contenido aún no está traducido
:::

<div className="row">
<div className="col">

# Introdução e fundamentos da linguagem Go

Este módulo apresenta os conceitos fundamentais da linguagem Go, incluindo sua história, características, configuração do ambiente e elementos básicos da sintaxe. O objetivo é fornecer uma base sólida para engenheiros Java que desejam aprender Go, com exemplos práticos e casos de uso. O conteúdo é objetivo, mas detalhado, e pode ser refinado posteriormente.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-baby.png').default} 
    style={{ marginTop:'80px' }}
    alt="A diaper brown gopher" />
</div>
</div>

<br />

## História e propósito do Go

### História

Go, também conhecido como Golang, foi criado em 2009 por `Robert Griesemer`, `Rob Pike` e `Ken Thompson` no Google. A linguagem surgiu para atender às necessidades de `desenvolvimento em larga escala`, com `foco em sistemas distribuídos`, `infraestrutura de servidores` e `ferramentas modernas`. Inspirada por C, Pascal e outras linguagens, `Go combina simplicidade com eficiência`.

#### Propósito

Go foi projetado para:

- **Simplicidade**: Sintaxe clara e minimalista, reduzindo a complexidade
- **Performance**: Compilação rápida para binários nativos, com desempenho próximo ao de C/C++
- **Concorrência**: Suporte nativo para programação concorrente via `goroutines` e `channels`
- **Produtividade**: Ferramentas integradas (formatação, testes, documentação) para agilizar o desenvolvimento

:::info Caso de uso
Go é amplamente usado em projetos como Kubernetes, Docker e sistemas de alta performance no Google, onde a escalabilidade e a manutenção são críticas.
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
### Simplicidade

- Sintaxe reduzida, com poucas palavras-chave (25, contra 50+ em Java)
- Formatação automática com `go fmt`
- Ausência de recursos complexos como herança ou sobrecarga de métodos
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

- Compilação para binário nativo, eliminando a necessidade de uma máquina virtual (diferente de Java)
- Garbage collector otimizado para baixa latência
- Execução rápida, ideal para serviços de alta carga
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

### Concorrência

- **Goroutines**: Funções leves que executam concorrência de forma eficiente
- **Channels**: Mecanismo para comunicação segura entre goroutines
- Diferente do modelo de threads em Java, Go abstrai a complexidade do gerenciamento de concorrência
</div>
</div>
:::info Caso de uso
APIs RESTful e microsserviços, onde Go oferece alta performance e concorrência para lidar com múltiplas requisições simultâneas.
:::

<br />

## Instalação, workspace e `go mod`

<!-- INSTALAÇÃO -->
<div className="row">
<div className="col">
### Instalação

1. Baixe o Go em [go.dev](https://go.dev/dl/) para o seu sistema operacional
2. Instale seguindo as instruções específicas:

- **Linux ou MAC**: Descompacte o arquivo e adicione ao `$PATH`
- **Windows**: Execute o instalador e verifique com `go version`
</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer.png').default} 
    style={{ marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>
:::tip Nota
No sistema operacional `MacOS` também é possível instalar o Go utilizando o `brew install go`
:::

4. Verifique a instalação utilizando o comando:

```bash
go version
```

<br />

### Workspace

<!-- WORKSPACE -->
<div className="row">
<div className="col">
Go utiliza uma estrutura de diretórios para organizar projetos. A partir do Go 1.11, o Go Modules substituiu o antigo $GOPATH como padrão.

#### Diretórios principais (opcional, com Go Modules)

- `src/`: Código-fonte
- `bin/`: Binários compilados
- `pkg/`: Pacotes compartilhados

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer-2.png').default} 
    style={{ transform:'scale(1.1)', marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>

:::tip Nota
Com Go Modules, você pode trabalhar em qualquer diretório
:::

<br />

### Go Modules

Go Modules gerenciam dependências de forma semelhante ao Maven em Java.

1. Inicialize um módulo:

```bash
go mod init github.com/meu-usuario-github/meu-projeto
```

2. Adicione dependências:

```bash
go get github.com/exemplo/pacote
```

3. O arquivo `go.mod` armazena as dependências e versões

#### Exemplo de `go.mod`

```bash
module github.com/meu-usuario-github/meu-projeto

go 1.21

require github.com/exemplo/pacote v1.0.0
```

:::info Caso de uso
Go Modules facilita a manutenção de projetos grandes, similar ao gerenciamento de dependências em Java com Maven ou Gradle.
:::

<br />

## Estrutura básica de um programa

Um programa Go segue uma estrutura simples, com o pacote main como ponto de entrada.

### Exemplo

<div className="row">
<div className="col">

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod1/lab/hello-world.go)!

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-hello-world.png').default} 
    style={{ transform:'scale(1)', marginTop:'-70px' }}
    alt="A blue Go gopher" />
</div>
</div>

- **package main**: Define o pacote principal, que gera um executável.
- **import "fmt"**: Importa o pacote fmt para formatação e saída.
- **func main()**: Função de entrada, equivalente ao public static void main em Java.

<br />

## Compilação e execução:

```bash
go run hello.go     # Executa diretamente
go build hello.go   # Compila para um binário
```

<br />

## Comparação com Java

- Em Java, classes e métodos estáticos são obrigatórios. Em Go, a função main é suficiente.
- Go não usa ponto e vírgula (;) ao final das linhas.

<br />

## Tipos Primitivos, Funções, Variáveis e Constantes

### Tipos primitivos

Go possui tipos primitivos simples, com tipagem estática (semelhante a Java, mas mais concisa).

<div className="row">
<div className="col">

| Tipo              | Descrição                  | Exemplo                |
| ----------------- | -------------------------- | ---------------------- |
| int, int32, int64 | Inteiros com ou sem sinal  | 42, -10                |
| float32, float64  | Números de ponto flutuante | 3.14, -0.001           |
| bool              | Booleano                   | true, false            |
| string            | Cadeia de caracteres UTF-8 | "Hello, Go!"           |
| byte              | Alias para uint8           | 65 (equivalente a ‘A’) |

</div>
<div className="col col--4 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-primitive.png').default} 
    style={{ transform:'scale(1.4)', marginTop:'10px' }}
    alt="A blue Go gopher" />
</div>
</div>
:::tip Nota
Go não suporta tipos implícitos como var genérico em Java. A inferência de tipo é feita com :=.
:::

### Variáveis

- Declaração explícita:

```go
    var nome string = "Raywall"
    var idade int = 45
```

- Declaração curta (inferência de tipo):

```go
    nome := "Raywall"
    idade := 45
```

- Múltiplas variáveis:

```go
    var x, y int = 10, 20
    a, b := "hello", true
```

### Comparação com Java

#### Java

```java
    String nome = "Raywall";
```

#### Go

```go
    nome := "Raywall" // mais conciso, mas estritamente tipado
```

### Constantes

Constantes são definidas com const e não podem ser alteradas.

```go
const Pi float64 = 3.14159
const NomeProjeto = "MeuApp"
```

### Funções

Funções em Go são declaradas com func, podem retornar múltiplos valores e não suportam sobrecarga (diferente de Java).

#### Exemplo

```go
package main

import "fmt"

func soma(a int, b int) int {
    return a + b
}

func troca(x, y string) (string, string) {
    return y, x
}

func main() {
    fmt.Println(soma(5, 3)) // Saída: 8
    a, b := troca("hello", "world")
    fmt.Println(a, b) // Saída: world hello
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod1/funcoes.go)!

:::info Caso de uso
Funções com múltiplos retornos são úteis para tratamento de erros, como value, err := funcao().
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
{label: 'Configuração', value: 'config'},
{label: 'Exercício 1', value: 'hello'},
{label: 'Exercício 2', value: 'vars'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

### Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab1`:

```bash
mkdir lab1
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab1
go mod init github.com/seu-usuario/lab1
```

</TabItem>
<TabItem value="hello">

### Hello, World!

#### Objetivo

Criar um programa simples em Go para praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

#### Passo a passo

1. Crie o arquivo `hello.go` com o seguinte conteúdo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod1/lab/hello-world.go)!

##### Execute

```bash
go run hello.go
```

</TabItem>
<TabItem value="vars">

### Manipulação de variáveis e tipos

#### Objetivo

Criar um programa simples em Go para praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

#### Passo a passo

1. Crie um arquivo `vars.go` com o seguinte conteúdo:

```go
package main

import "fmt"

func main() {
    // Declaração de variáveis
    var nome string = "Raywall"
    idade := 45

    const versao = "1.0.0"

    // Tipos
    var preco float64 = 99.99
    ativo := true

    // Exibição
    fmt.Printf("Nome: %s, Idade: %d, Versão: %s\n", nome, idade, versao)
    fmt.Printf("Preço: %.2f, Ativo: %t\n", preco, ativo)

    // Função com múltiplos retornos
    x, y := troca("Go", "Java")
    fmt.Printf("Troca: %s, %s\n", x, y)
}

func troca(a, b string) (string, string) {
    return b, a
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod1/lab/vars.go)!

##### Execute

```bash
    go run vars.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Modifique `vars.go` para incluir uma função que calcula o dobro de um número
- Adicione uma constante para o IVA (Imposto sobre Valor Agregado, ex: 0.21) e calcule o preço final de um produto

#### Saída esperada

```bash
Nome: Raywall, Idade: 45, Versão: 1.0.0
Preço: 99.99, Ativo: true
Troca: Java, Go
```

:::info Caso de uso prático
Este lab simula a manipulação de dados em um sistema simples, como um cadastro de produtos, onde variáveis e constantes são usadas para armazenar informações fixas (versão) e dinâmicas (preço, status).
:::
</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobre os fundamentos do Go, desde sua história até a criação de programas básicos. Engenheiros Java notarão semelhanças (tipagem estática, funções) e diferenças (sintaxe concisa, ausência de classes). O lab prático reforça a configuração do ambiente e o uso de tipos, variáveis e funções, preparando o time para módulos mais avançados, como estruturas de controle e concorrência.

:::tip Próximos passos
No próximo módulo, exploraremos `estruturas de controle`, `slices`, `maps` e `pacotes`, com mais exemplos práticos.
:::
