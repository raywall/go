---
sidebar_position: 6
sidebar_label: Módulo 04
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

:::warning Atención
Este contenido aún no está traducido
:::

# Tratamento de erros em Go

<div className="row">
<div className="col">

Este módulo explora a abordagem de tratamento de erros em Go, que difere significativamente do modelo de exceções usado em Java. O foco está na filosofia de erros explícitos, no uso do tipo error, técnicas de `wrapping`/`unwrapping` e `logging estruturado`. O conteúdo é voltado para engenheiros Java, com exemplos práticos e casos de uso objetivos para consulta futura.

O lab prático implementa funções com tratamento de erros e logging estruturado.

</div>
<div className="col col--4 text--center">
<img 
	src={require('@site/static/img/gophers/gopher-coffee.png').default} 
	style={{ transform:'scale(0.8)', marginTop:'-65px' }}
	alt="A diaper brown gopher" />
</div>
</div>

<br />

## Filosofia do Go: erros explícitos

- **Erros como valores**: Em Go, erros são valores do tipo error, retornados explicitamente pelas funções, ao invés de lançados como exceções em Java (`try-catch`).
- **Clareza e previsibilidade**: O programador deve verificar erros explicitamente, reduzindo comportamentos inesperados.
- **Evite `panic`**: Diferente de Java, panic é reservado para falhas irrecuperáveis (ex: falha de inicialização crítica) e não para fluxo normal.

### Comparação com Java

#### Java

```java
try { ... } catch (Exception e) { ... }
```

#### Go

```go
if err != nil { return err }
```

:::info Caso de uso
A abordagem explícita é ideal para sistemas onde a robustez é crítica, como APIs de servidores ou ferramentas CLI.
:::

<br />

## Padrão `error`, `errors.New` e `fmt.Errorf`

### Tipo `error`

- **Uma interface embutida**: `type error interface { Error() string }`
- Qualquer tipo que implemente o método `Error() string` é um error

### `errors.New`

- Cria erros simples com uma mensagem fixa

#### Exemplo

```go
package main

import (
	"errors"
	"fmt"
)

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}

	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 0)

	if err != nil {
		fmt.Println("Erro:", err) // Saída: Erro: divisão por zero
		return
	}

	fmt.Println("Resultado:", resultado)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/error.go)!

### `fmt.Errorf`

- Formata mensagens de erro dinamicamente, semelhante a `String.format` em Java.

#### Exemplo

```go
func acessarIndice(slice []int, indice int) (int, error) {
	if indice >= len(slice) {
		return 0, fmt.Errorf("índice %d fora do limite (%d)", indice, len(slice))
	}

	return slice[indice], nil
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/errorf.go)!

:::info Caso de uso
`errors.New` é usado para erros fixos, enquanto `fmt.Errorf` é ideal para erros com contexto dinâmico (ex: validação de entrada)
:::

<br />

## Wrapping e Unwrapping com `errors.Is` e `errors.As`

### Wrapping

- Permite adicionar contexto a um erro, mantendo o erro original
- Usado com o operador `%w` em `fmt.Errorf`

#### Exemplo

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("não encontrado")

func buscarDado(id int) error {
	return ErrNotFound
}

func processar(id int) error {
	err := buscarDado(id)
	if err != nil {
		return fmt.Errorf("falha ao processar id %d: %w", id, err)
	}
	return nil
}

```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/wrapping.go)!

### `errors.Is`

- Verifica se um erro (ou seus erros encapsulados) corresponde a um erro específico.

#### Exemplo

```go
func main() {
	err := processar(45)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Erro é 'não encontrado'") // Saída: Erro é 'não encontrado'
	}
	fmt.Println(err) // Saída: falha ao processar id 45: não encontrado
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/errors-is.go)!

### `errors.As`

- Extrai um erro de um tipo específico de uma cadeia de erros encapsulados.

#### Exemplo

```go
package main

import (
	"errors"
	"fmt"
)

type MeuErro struct {
	Mensagem string
}

func (e *MeuErro) Error() string {
	return e.Mensagem
}

func operacao() error {
	return fmt.Errorf("contexto: %w", &MeuErro{Mensagem: "falha específica"})
}

func main() {
	err := operacao()

	var meuErro *MeuErro
	if errors.As(err, &meuErro) {
		fmt.Println("Erro capturado:", meuErro.Mensagem) // Saída: falha específica
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/errors-as.go)!

#### Comparação com Java:

- `Wrapping` em Go é semelhante a `Exception.getCause()` em Java, mas mais explícito.
- `errors.Is` e `errors.As` substituem verificações como instanceof ou comparações de tipos de exceção.

:::info Caso de uso
Wrapping é útil em APIs para adicionar contexto (ex: ID da requisição) sem perder o erro original.
:::

<br />

## Pacote `log` e `log/slog`

### Pacote `log`

- Fornece logging básico, com saída para `stderr` por padrão.
- Funções como `log.Print`, `log.Fatal` e `log.Panic`.

#### Exemplo

```go
package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("Iniciando aplicação...")
	err := errors.New("falha crítica")

	if err != nil {
		log.Fatal("Erro:", err) // Sai com código 1
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/log.go)!

### Pacote `log/slog` (Go 1.21+)

- Oferece logging estruturado, com suporte a JSON e campos personalizados.
- Mais adequado para aplicações modernas, como microsserviços.

#### Exemplo

```go
package main

import (
	"errors"
	"log/slog"
	"os"
)

func main() {
	// Configurar logger JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Log com campos estruturados
	logger.Info("Processando requisição", "id", 45, "metodo", "GET")
	err := errors.New("falha na operação")

	logger.Error("Erro encontrado", "error", err, "id", 45)
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/slog.go)!

#### Saída (JSON)

```json
	{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Processando requisição","id":45,"metodo":"GET"}
	{"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"Erro encontrado","error":"falha na operação","id":45}
```

#### Comparação com Java

- `log` é semelhante ao `System.out.println` ou `java.util.logging`
- `log/slog` é comparável a bibliotecas como `Log4j` ou `SLF4J`, com suporte a logging estruturado.

:::info Caso de uso
`log/slog` é ideal para aplicações distribuídas, onde logs estruturados facilitam a análise em ferramentas como `ELK Stack` ou `Prometheus`.
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
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

### Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab4`

```bash
mkdir lab4
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab4
go mod init github.com/seu-usuario/lab4
```

</TabItem>
<TabItem value="app">

### Criar Funções com tratamento de erros e logging estruturado

#### Objetivo

Refatorar o CRUD do Módulo 3 para incluir tratamento de erros robusto e logging estruturado com log/slog, mantendo a interface de repositório.

#### Passo a passo

1. Crie um arquivo `crud.go` com o seguinte código:

```go
package main

import (
	"errors"
	"fmt"
	"log/slog"
)

// Erros personalizados
var (
	ErrPrecoInvalido        = errors.New("preço não pode ser negativo")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")
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
	logger    *slog.Logger
}

// Novo repositório com logger
func NovoRepositorioEmMemoria(logger *slog.Logger) *RepositorioEmMemoria {
	return &RepositorioEmMemoria{
		produtos:  make([]Produto, 0), // Inicializa um slice vazio para produtos
		idCounter: 0,                  // Inicializa o contador de IDs
		logger:    logger,             //Associa o logger ao repositório
	}
}

// Implementação da interface
func (r *RepositorioEmMemoria) Criar(nome string, preco float64) (Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao criar produto", "error", ErrPrecoInvalido, "nome", nome)
		return Produto{}, ErrPrecoInvalido
	}

	r.idCounter++
	produto := Produto{ID: r.idCounter, Nome: nome, Preco: preco}
	r.produtos = append(r.produtos, produto)
	r.logger.Info("Produto criado", "id", produto.ID, "nome", nome, "preco", preco)
	return produto, nil
}

func (r *RepositorioEmMemoria) Buscar(id int) (Produto, error) {
	for _, p := range r.produtos {
		if p.ID == id {
			r.logger.Info("Produto encontrado", "id", id)
			return p, nil
		}
	}

	r.logger.Error("Falha ao buscar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return Produto{}, fmt.Errorf("buscar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

func (r *RepositorioEmMemoria) Listar() ([]Produto, error) {
	r.logger.Info("Listando produtos", "total", len(r.produtos))
	return r.produtos, nil
}

func (r *RepositorioEmMemoria) Atualizar(id int, nome string, preco float64) (Produto, error) {
	if preco < 0 {
		r.logger.Error("Falha ao atualizar produto", "error", ErrPrecoInvalido, "id", id)
		return Produto{}, ErrPrecoInvalido
	}

	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos[i] = Produto{ID: id, Nome: nome, Preco: preco}
			r.logger.Info("Produto atualizado", "id", id, "nome", nome, "preco", preco)
			return r.produtos[i], nil
		}
	}

	r.logger.Error("Falha ao atualizar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return Produto{}, fmt.Errorf("atualizar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

func (r *RepositorioEmMemoria) Deletar(id int) error {
	for i, p := range r.produtos {
		if p.ID == id {
			r.produtos = append(r.produtos[:i], r.produtos[i+1:]...)
			r.logger.Info("Produto deletado", "id", id)
			return nil
		}
	}

	r.logger.Error("Falha ao deletar produto", "error", ErrProdutoNaoEncontrado, "id", id)
	return fmt.Errorf("deletar produto id %d: %w", id, ErrProdutoNaoEncontrado)
}

// Função para exibir produtos
func exibirProdutos(repo RepositorioProdutos) error {
	produtos, err := repo.Listar()
	if err != nil {
		return fmt.Errorf("exibir produtos: %w", err)
	}
	for _, p := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", p.ID, p.Nome, p.Preco)
	}

	return nil
}

func main() {
	// Inicializar repositório com logger
	repo := NovoRepositorioEmMemoria()

	// Criar produtos
	p1, err := repo.Criar("Laptop", 999.99)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	repo.Criar("Mouse", 29.99)

	// Listar produtos
	fmt.Println("Lista de produtos:")
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Buscar produto
	if p, err := repo.Buscar(p1.ID); err == nil {
		fmt.Printf("Produto encontrado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

	// Atualizar produto
	if p, err := repo.Atualizar(p1.ID, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Produto atualizado: %+v\n", p)
	} else {
		fmt.Println("Erro:", err)
	}

	// Tentar atualizar com preço inválido
	if _, err := repo.Atualizar(p1.ID, "Laptop Pro", -100); err != nil {
		fmt.Println("Erro esperado:", err) // Saída: preço não pode ser negativo
	}

	// Deletar produto
	if err := repo.Deletar(2); err == nil {
		fmt.Println("Produto deletado com sucesso")
	} else {
		fmt.Println("Erro:", err)
	}

	// Listar novamente
	fmt.Println("Lista final:")
	if err := exibirProdutos(repo); err != nil {
		fmt.Println("Erro:", err)
	}
}
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod4/lab/crud.go)!

2. Execução:

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

### Tarefa

- Adicione um erro personalizado para quando o nome do produto estiver vazio.
- Implemente uma função que use errors.As para capturar e tratar erros específicos (ex.: ErrPrecoInvalido).
- Configure o logger para salvar logs em um arquivo, além de exibir no console.

#### Saída esperada

##### Console

```bash
Lista de produtos:
ID: 1, Nome: Laptop, Preço: 999.99
ID: 2, Nome: Mouse, Preço: 29.99

Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
Erro esperado: preço não pode ser negativo
Produto deletado com sucesso

Lista final:
ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

#### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto criado","id":1,"nome":"Laptop","preco":999.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto criado","id":2,"nome":"Mouse","preco":29.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listando produtos","total":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto encontrado","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto atualizado","id":1,"nome":"Laptop Pro","preco":1299.99}
{"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"Falha ao atualizar produto","error":"preço não pode ser negativo","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto deletado","id":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listando produtos","total":1}
```

:::info Caso de uso prático
Este lab simula um sistema real com logging estruturado, útil para depuração e monitoramento em aplicações distribuídas, como APIs RESTful.
:::
</TabItem>
</Tabs>
<br />

---

## Conclusão

Este módulo cobriu a filosofia de erros explícitos do Go, o uso de `errors.New`, `fmt.Errorf`, `wrapping/unwrapping` com `errors.Is` e `errors.As`, e logging com `log` e `log/slog`. O lab prático reforça a aplicação desses conceitos em um CRUD, com logs estruturados para facilitar a manutenção. Engenheiros Java notarão a diferença em relação ao modelo de exceções, mas verão semelhanças com logging estruturado em frameworks como `SLF4J`.

:::tip Próximos passos
No próximo módulo, exploraremos concorrência com goroutines e channels, além de testes e benchmarks, preparando o time para desenvolver aplicações escaláveis e robustas em Go.
:::
