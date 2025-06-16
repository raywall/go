---
sidebar_position: 4
sidebar_label: Módulo 02
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Estruturas de controle e tipos compostos em Go

Este módulo aprofunda os fundamentos do Go, abordando `estruturas de controle`, tipos compostos (`arrays`, `slices`, `maps`), `structs` e `ponteiros`. O conteúdo é voltado para engenheiros Java, com exemplos práticos e casos de uso objetivos, mantendo profundidade para consulta futura.

O lab prático implementa uma API fake de CRUD, reforçando os conceitos aprendidos.

<br />

## Estruturas de controle

Go possui estruturas de controle simples, sem parênteses obrigatórios (diferente de Java) e com foco em clareza.

### `if`

- Suporta inicialização na própria declaração.
- Não há operador ternário em Go.

#### Exemplo

```go
package main

import "fmt"

func main() {
	numero := 45
	if x := numero % 2; x == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/controle-if.go)!

#### Comparação com Java:

##### Java

```java
if (x % 2 == 0) { ... }
```

##### Go

```go
if x := numero % 2; x == 0 { ... } // inicialização direta
```

### `for`

- Única estrutura de laço em Go (não há `while` ou `do-while`)
- Suporta três formas: tradicional, condição simples e infinito

#### Exemplo

```go
package main

import "fmt"

func main() {
	// Laço tradicional
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Como while
	contador := 0
	for contador < 3 {
		fmt.Println(contador)
		contador++
	}

	// Infinito (com break)
	for {
		fmt.Println("Infinito")
		break
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/controle-for.go)!

### `switch`

- Mais simples que em Java: não precisa de break (não há fall-through por padrão)
- Suporta múltiplas condições por caso e inicialização

#### Exemplo

```go
package main

import "fmt"

func main() {
    dia := 3
    switch dia {
        case 1, 2:
            fmt.Println("Início da semana")
        case 3, 4, 5:
            fmt.Println("Meio da semana")
        default:
            fmt.Println("Fim de semana")
    }

    // Switch com expressão
    switch x := dia \* 2; x {
    case 6:
        fmt.Println("Dia 3 dobrado")
    default:
        fmt.Println("Outro valor")
    }
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/controle-switch.go)!

### `defer`

- Adia a execução de uma instrução até o fim da função
- Útil para liberar recursos (similar ao `try-with-resources` em Java)

#### Exemplo

```go
package main

import "fmt"

func main() {
	defer fmt.Println("Executado no final")
	fmt.Println("Executado primeiro")
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/controle-defer.go)!

:::info Caso de uso
`defer` é usado para fechar arquivos, conexões de banco ou liberar mutexes
:::

<br />

## Arrays, slices e maps

### Arrays

- Tamanho fixo, definido na declaração
- Menos comum em Go, pois slices são mais flexíveis

#### Exemplo

```go
var numeros [3]int = [3]int{1, 2, 3}
fmt.Println(numeros[0]) // Saída: 1
```

#### Comparação com Java

Arrays em Go são semelhantes aos arrays de Java (`int[]`), mas menos usados devido aos slices

### Slices

- Estrutura dinâmica, baseada em arrays, mas com tamanho variável
- Declarados com `[]tipo` ou criados com make

#### Exemplo

```go
package main

import "fmt"

func main() {
	// Declaração
	slice := []int{1, 2, 3}
	fmt.Println(slice) // [1 2 3]

	// Adicionar elementos
	slice = append(slice, 4)
	fmt.Println(slice) // [1 2 3 4]

	// Slice a partir de array
	array := [5]int{1, 2, 3, 4, 5}
	subSlice := array[1:4] // [2 3 4]
	fmt.Println(subSlice)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/slices.go)!

:::info Nota
Slices são passados por referência, mas o array subjacente pode ser copiado se modificado
:::

### Maps

- Estrutura `chave-valor`, semelhante ao `HashMap` em Java
- Declarados com `map[tipoChave]tipoValor`

#### Exemplo

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["idade"] = 45
	m["preco"] = 99
	fmt.Println(m) // map[idade:42 preco:99]

	// Verificar existência
	valor, existe := m["idade"]
	if existe {
		fmt.Println("Idade:", valor)
	}

	// Deletar
	delete(m, "preco")
	fmt.Println(m) // map[idade:42]
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/maps.go)!

:::info Caso de uso
Slices são ideais para listas dinâmicas (ex: lista de usuários), enquanto maps são úteis para associações rápidas (ex: cache de IDs para valores)
:::

### Structs e tags de struct

#### Structs

- Equivalentes a classes em Java, mas sem herança
- Definidas com `type Nome struct`

##### Exemplo

```go
package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p := Pessoa{Nome: "Raywall", Idade: 45}
	fmt.Println(p) // {Raywall 45}

	// Acesso a campos
	fmt.Println(p.Nome) // Raywall

	// Struct anônima
	anon := struct {
		Valor string
	}{Valor: "Teste"}

	fmt.Println(anon) // {Teste}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/structs.go)!

#### Tags de struct

- Usadas para metadados, como serialização JSON
- Declaradas com crases (`\``)

##### Exemplo

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco,omitempty"`
}

func main() {
	p := Produto{ID: 1, Nome: "Laptop"}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData)) // {"id":1,"nome":"Laptop"}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/tags-de-struct.go)!

##### Comparação com Java

Structs substituem classes, mas tags são semelhantes às anotações (`@JsonProperty`) do Jackson em Java

:::info Caso de uso
Structs com tags são amplamente usadas em APIs REST para `serializar`/`deserializar` JSON
:::

<br />

## Ponteiros

- Go suporta ponteiros para manipulação direta de memória, mas de forma segura
- Declarados com `\*tipo` (endereço) e `&variavel` (referência)

### Exemplo

```go
package main

import "fmt"

func main() {
	x := 45
	p := &x         // Ponteiro para x
	fmt.Println(*p) // 45 (desreferenciação)

	// Modificar via ponteiro
	*p = 100
	fmt.Println(x) // 100
}

func incrementar(p *int) {
	*p++
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/ponteiros.go)!

### Comparação com Java

- Java usa referências implícitas para objetos, sem controle direto de ponteiros
- Go exige ponteiros explícitos para alterações em funções (passagem por valor é padrão)

:::info Caso de uso
Ponteiros são úteis para modificar structs grandes sem copiar dados, similar a passar objetos por referência em Java.
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
2. Crie um diretório `lab2`:

```bash
mkdir lab2
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab2
go mod init github.com/seu-usuario/lab2
```

</TabItem>
<TabItem value="app">

### Implementar uma API Fake para CRUD em Memória

#### Objetivo

Criar uma API fake em Go que gerencia uma lista de produtos em memória, usando `slices`, `maps`, `structs` e `estruturas de controle`. O lab simula um CRUD (`Create`, `Read`, `Update`, `Delete`).

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

// Banco de dados em memória (slice de produtos)
var produtos []Produto
var idCounter = 1

// Create: Adiciona um produto
func criarProduto(nome string, preco float64) Produto {
	produto := Produto{ID: idCounter, Nome: nome, Preco: preco}
	idCounter++

	produtos = append(produtos, produto)
	return produto
}

// Read: Busca todos os produtos
func listarProdutos() []Produto {
	return produtos
}

// Read: Busca produto por ID
func buscarProduto(id int) (Produto, error) {
	for _, p := range produtos {
		if p.ID == id {
			return p, nil
		}
	}

	return Produto{}, errors.New("Produto não encontrado")
}

// Update: Atualiza um produto
func atualizarProduto(id int, nome string, preco float64) (Produto, error) {
	for i, p := range produtos {
		if p.ID == id {
			produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			return produtos[i], nil
		}
	}

	return Produto{}, errors.New("Produto não encontrado")
}

// Delete: Remove um produto
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
	// Criar produtos
	criarProduto("Laptop", 999.99)
	criarProduto("Mouse", 29.99)

	// Listar produtos
	fmt.Println("Lista de produtos:")
	for _, p := range listarProdutos() {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}

	// Buscar produto
	if p, err := buscarProduto(1); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	}

	// Atualizar produto
	if p, err := atualizarProduto(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	}

	// Deletar produto
	if err := deletarProduto(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	}

	// Listar novamente
	fmt.Println("Lista final:")

	for _, p := range listarProdutos() {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod2/lab/crud.go)!

#### Execute

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione validação para impedir preços negativos em criarProduto e atualizarProduto
- Implemente uma função que liste produtos com preço acima de um valor específico
- Use um map para armazenar produtos por ID, em vez de um slice, e compare o desempenho

#### Saída esperada

```bash
    Lista de produtos:
    ID: 1, Nome: Laptop, Preço: 999.99
    ID: 2, Nome: Mouse, Preço: 29.99

    Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
    Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
    Produto deletado com sucesso

    Lista final:
    ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

:::info Caso de uso prático
Este lab simula uma API simples de gerenciamento de produtos, semelhante a um backend de e-commerce. Slices e maps são usados para manipulação de dados em memória, enquanto structs modelam entidades.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu estruturas de controle (`if`, `for`, `switch`, `defer`), tipos compostos (`arrays`, `slices`, `maps`), `structs com tags e ponteiros`. O lab prático reforça a aplicação desses conceitos em um cenário realista de CRUD. Engenheiros Java notarão semelhanças com coleções (`list`, `map`) e diferenças na ausência de herança ou generics (até Go 1.18).

:::tip Próximos passos
No próximo módulo, exploraremos concorrência com `goroutines` e `channels`, além de `tratamento de erros` e `pacotes`, preparando o time para desenvolver aplicações escaláveis em Go.
:::
