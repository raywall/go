---
sidebar_position: 5
sidebar_label: Módulo 03
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Funções, métodos e interfaces em Go

Este módulo explora funções avançadas, métodos em structs, interfaces e boas práticas em Go, com foco em engenheiros Java que desejam adotar o estilo idiomático da linguagem. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura.

O lab prático refatora o CRUD do [Módulo 02](go-module-2), introduzindo interfaces para repositórios.

<br />

## Funções com múltiplos retornos

Go permite que funções retornem múltiplos valores, um recurso poderoso para tratamento de erros e resultados compostos, diferente de Java, que usa exceções ou objetos.

### Exemplo

```go
package main

import "fmt"

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado) // Resultado: 5
}
```

:::info Caso de uso
Funções com múltiplos retornos são comuns em Go para retornar um valor e um erro, substituindo o uso de exceções como em Java (`try-catch`).
:::

<br />

## Funções anônimas e closures

### Funções anônimas

- Funções sem nome, definidas inline, semelhantes a lambdas em Java
- Úteis para tarefas temporárias ou callbacks

#### Exemplo

```go
package main

import "fmt"

func main() {
    soma := func(a, b int) int {
        return a + b
    }

    fmt.Println(soma(3, 4)) // Saída: 7
}
```

### Closures

- Funções anônimas que capturam variáveis do escopo externo
- Similar a `closures` em Java (ex: lambdas com acesso a variáveis externas)

#### Exemplo

```go
package main

import "fmt"

func contador() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := contador()
    fmt.Println(c1()) // Saída: 1
    fmt.Println(c1()) // Saída: 2

    c2 := contador()
    fmt.Println(c2()) // Saída: 1 (novo contador)
}
```

:::info Caso de uso
Closures são úteis para manter estado em funções, como contadores ou geradores, sem a necessidade de `structs`.
:::

## Métodos em structs

- Go não possui classes, mas métodos podem ser associados a structs usando receivers
- Receivers podem ser por valor (`T`) ou por ponteiro (`*T`)

### Exemplo

```go
package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

// Método com receiver por valor
func (p Produto) Descricao() string {
	return fmt.Sprintf("%s: R$%.2f", p.Nome, p.Preco)
}

// Método com receiver por ponteiro
func (p *Produto) AplicarDesconto(desconto float64) {
	p.Preco -= p.Preco * desconto
}

func main() {
	p := Produto{Nome: "Laptop", Preco: 1000}

	fmt.Println(p.Descricao()) // Saída: Laptop: R$1000.00
	p.AplicarDesconto(0.1)

	fmt.Println(p.Descricao()) // Saída: Laptop: R$900.00
}
```

### Comparação com Java

- Em Java, métodos são definidos dentro de classes. Em Go, métodos são funções com receivers, associados a tipos (geralmente structs).
- Receiver por ponteiro (`*T`) é semelhante a modificar o estado de um objeto em Java.

:::info Caso de uso
Métodos em structs são usados para encapsular comportamento, como formatar dados ou aplicar regras de negócio.
:::

## Interfaces e duck typing

### Interfaces

- Definidas com type Nome interface, especificando métodos que um tipo deve implementar.
- **Go usa `duck typing`**: se um tipo implementa os métodos de uma interface, ele a satisfaz automaticamente (sem declaração explícita, diferente de Java com implements).

#### Exemplo

```go
package main

import "fmt"

type Descrivivel interface {
	Descricao() string
}

type Produto struct {
	Nome  string
	Preco float64
}

type Servico struct {
	Nome string
	Taxa float64
}

func (p Produto) Descricao() string {
	return fmt.Sprintf("Produto: %s, R$%.2f", p.Nome, p.Preco)
}

func (s Servico) Descricao() string {
	return fmt.Sprintf("Serviço: %s, R$%.2f", s.Nome, s.Taxa)
}

func exibir(d Descrivivel) {
	fmt.Println(d.Descricao())
}

func main() {
	p := Produto{Nome: "Laptop", Preco: 1000}
	s := Servico{Nome: "Consultoria", Taxa: 500}

	exibir(p) // Saída: Produto: Laptop, R$1000.00
	exibir(s) // Saída: Serviço: Consultoria, R$500.00
}
```

#### Comparação com Java

- Em Java, interfaces exigem `implements`. Em Go, a implementação é implícita.
- Go não suporta herança, mas interfaces permitem polimorfismo.

### Interface vazia (`interface{}`)

- Equivalente a `Object` em Java, aceita qualquer tipo.
- Usada com moderação, geralmente com type assertions ou type switches.

#### Exemplo

```go
func inspecionar(v interface{}) {
    switch t := v.(type) {
    case Produto:
        fmt.Println("É um Produto:", t.Nome)
    case string:
        fmt.Println("É uma string:", t)
    default:
        fmt.Println("Tipo desconhecido")
    }
}
```

:::info Caso de uso
Interfaces são amplamente usadas em Go para abstrair repositórios, serviços ou plugins, garantindo flexibilidade e testabilidade.
:::

<br />

## Boas práticas e princípios de design idiomático em Go

### Simplicidade

- Prefira soluções claras e concisas
- Evite complexidade desnecessária (ex: camadas excessivas de abstração)

### Nomenclatura

- Use nomes curtos e descritivos (ex: p para Produto, em vez de `produtoInstance`)
- Funções e métodos exportados começam com maiúscula (ex: `Descricao`)

### Tratamento de erros

- Sempre cheque erros retornados (`if err != nil`)
- Evite panic em código de produção, exceto em casos extremos.

### Interfaces pequenas

- Defina interfaces com poucos métodos, focadas em uma única responsabilidade.
- **Exemplo**: `io.Reader` e `io.Writer` em vez de interfaces genéricas.

### Composição sobre herança

- Use `structs aninhados` para composição, em vez de herança como em Java.

### Evite interface vazia (`interface{}`)

- Use type assertions ou type switches apenas quando necessário

### Use Go idiomaticamente

- Evite padrões de Java (ex: `getters`/`setters` desnecessários)
- Aproveite `go fmt` e ferramentas como `golint` para consistência

:::info Caso de uso
Boas práticas garantem código legível e manutenível, especialmente em equipes grandes, como em projetos de microsserviços.
:::

<br />

---

## Laboratório

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

### Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab3`

```bash
mkdir lab3
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab3
go mod init github.com/seu-usuario/lab3
```

</TabItem>
<TabItem value="app">

### Refatorar o CRUD Usando Interfaces para Repositórios

#### Objetivo

Refatorar o CRUD do Módulo 2 para usar uma interface de repositório, abstraindo o armazenamento de dados e facilitando testes e manutenção.

#### Passo a passo

1. Crie um arquivo `crud.go` com o seguinte código:

```go
package main

import (
	"errors"
	"fmt"
)

// Struct para Produto
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// Interface do repositório
type RepositorioProdutos interface {
	Criar(nome string, preco float64) (Produto, error)
	Buscar(id int) (Produto, error)
	Listar() ([]Produto, error)
	Atualizar(id int, nome string, preco float64) (Produto, error)
	Deletar(id int) error
}

// Repositório em memória
type RepositorioEmMemoria struct {
	produtos  []Produto
	idCounter int
}

// Implementação da interface
func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (Produto, error) {
	if preco < 0 {
		return Produto{}, errors.New("preço não pode ser negativo")
	}

	r.idCounter++
	produto := Produto{ID: r.idCounter, Nome: nome, Preco: preco}

	r.produtos = append(r.produtos, produto)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id int) (Produto, error) {
	for _, p := range r.produtos {
		if p.ID == id {
			return p, nil
		}
	}

	return Produto{}, errors.New("produto não encontrado")
}

func (r *RepositorioEmMemoria) Listar() ([]Produto, error) {
	return r.produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id int, nome string, preco float64) (Produto, error) {
	if preco < 0 {
		return Produto{}, errors.New("preço não pode ser negativo")
	}

	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			return r.produtos[i], nil
		}
	}

	return Produto{}, errors.New("produto não encontrado")
}

func (r *RepositorioEmMemoria) Deletar(id int) error {
	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos = append(r.produtos[:i], r.produtos[i+1:]...)
			return nil
		}
	}

	return errors.New("produto não encontrado")
}

// Função para exibir produtos
func exibirProdutos(repo RepositorioProdutos) {
	produtos, _ := repo.Listar()
	fmt.Println("Lista de produtos:")

	for _, p := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
}

func main() {
	// Inicializar repositório
	repo := &RepositorioEmMemoria{}

	// Criar produtos
	repo.Criar("Laptop", 999.99)
	repo.Criar("Mouse", 29.99)

	// Listar produtos
	exibirProdutos(repo)

	// Buscar produto
	if p, err := repo.Buscar(1); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	}

	// Atualizar produto
	if p, err := repo.Atualizar(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	}

	// Deletar produto
	if err := repo.Deletar(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	}

	// Listar novamente
	exibirProdutos(repo)
}
```

##### Execução

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Crie um segundo repositório (ex: `RepositorioMock`) que implemente `RepositorioProdutos` para testes, retornando dados fixos.
- Adicione uma função anônima no main para filtrar produtos com preço acima de um valor
- Use um método em Produto para calcular o preço com imposto (ex: 21%)

#### Saída esperada

```bash
Lista de produtos:
ID: 1, Nome: Laptop, Preço: 999.99
ID: 2, Nome: Mouse, Preço: 29.99

Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
Produto deletado com sucesso

Lista de produtos:
ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

:::info Caso de uso prático
A interface RepositorioProdutos abstrai o armazenamento, permitindo trocar a implementação (ex: de memória para banco de dados) sem alterar o código principal, similar ao padrão Repository em Java.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu funções com múltiplos retornos, funções anônimas, closures, métodos em structs, interfaces e boas práticas idiomáticas em Go. O lab prático refatorou o CRUD do [Módulo 02](go-module-2.md), introduzindo interfaces para maior modularidade. Engenheiros Java notarão semelhanças com interfaces e lambdas, mas com a abordagem mais simples e implícita de Go.

:::tip Próximos passos
No próximo módulo, exploraremos concorrência com `goroutines` e `channels`, além de `tratamento avançado de erros e testes`, preparando o time para aplicações escaláveis e robustas em Go.
:::
