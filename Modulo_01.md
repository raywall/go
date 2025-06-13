
# Módulo 01 - Introdução e fundamentos da linguagem Go

Este módulo apresenta os conceitos fundamentais da linguagem Go, incluindo sua história, características, configuração do ambiente e elementos básicos da sintaxe. O objetivo é fornecer uma base sólida para engenheiros Java que desejam aprender Go, com exemplos práticos e casos de uso. O conteúdo é objetivo, mas detalhado, e pode ser refinado posteriormente.


## História e Propósito do Go

### História

Go, também conhecido como Golang, foi criado em 2009 por Robert Griesemer, Rob Pike e Ken Thompson no Google. A linguagem surgiu para atender às necessidades de desenvolvimento em larga escala, com foco em sistemas distribuídos, infraestrutura de servidores e ferramentas modernas. Inspirada por C, Pascal e outras linguagens, Go combina simplicidade com eficiência.

#### Propósito

Go foi projetado para:

- **Simplicidade**: Sintaxe clara e minimalista, reduzindo a complexidade
- **Performance**: Compilação rápida para binários nativos, com desempenho próximo ao de C/C++
- **Concorrência**: Suporte nativo para programação concorrente via goroutines e channels
- **Produtividade**: Ferramentas integradas (formatação, testes, documentação) para agilizar o desenvolvimento

>[!NOTE] Caso de uso
>Go é amplamente usado em projetos como Kubernetes, Docker e sistemas de alta performance no Google, onde a escalabilidade e a manutenção são críticas.


## Características da linguagem

### Simplicidade

- Sintaxe reduzida, com poucas palavras-chave (25, contra 50+ em Java)
- Formatação automática com go fmt
- Ausência de recursos complexos como herança ou sobrecarga de métodos


### Performance

- Compilação para binário nativo, eliminando a necessidade de uma máquina virtual (diferente de Java)
- Garbage collector otimizado para baixa latência
- Execução rápida, ideal para serviços de alta carga


### Concorrência

- **Goroutines**: Funções leves que executam concorrência de forma eficiente
- **Channels**: Mecanismo para comunicação segura entre goroutines
- Diferente do modelo de threads em Java, Go abstrai a complexidade do gerenciamento de concorrência

>[!NOTE] Caso de uso:
>APIs RESTful e microsserviços, onde Go oferece alta performance e concorrência para lidar com múltiplas requisições simultâneas.

  
## Instalação, workspace e `go mod`

### Instalação

1. Baixe o Go em [go.dev](https://go.dev/dl/) para o seu sistema operacional
2. Instale seguindo as instruções específicas:

- Linux/Mac: Descompacte o arquivo e adicione ao $PATH.
- Windows: Execute o instalador e verifique com go version.

4. Verifique a instalação:
```bash
go version
```

### Workspace

Go utiliza uma estrutura de diretórios para organizar projetos. A partir do Go 1.11, o Go Modules substituiu o antigo $GOPATH como padrão.


#### Diretórios principais (opcional, com Go Modules)

- `src/`: Código-fonte
- `bin/`: Binários compilados
- `pkg/`: Pacotes compartilhados

> Com Go Modules, você pode trabalhar em qualquer diretório


### Go Modules

Go Modules gerenciam dependências de forma semelhante ao Maven em Java.

1. Inicialize um módulo:

```bash
go mod init github.com/seu-usuario/meu-projeto
```


2. Adicione dependências:  

```bash
go get github.com/exemplo/pacote
```

    
3. O arquivo go.mod armazena as dependências e versões


**Exemplo de `go.mod`**:
```bash
module github.com/seu-usuario/meu-projeto

go 1.21

require github.com/exemplo/pacote v1.0.0
```

>[!NOTE] Caso de uso
>Go Modules facilita a manutenção de projetos grandes, similar ao gerenciamento de dependências em Java com Maven ou Gradle.

  

4. Estrutura Básica de um Programa

Um programa Go segue uma estrutura simples, com o pacote main como ponto de entrada.

Exemplo:

package main

  

import "fmt"

  

func main() {

    fmt.Println("Hello, World!")

}

- package main: Define o pacote principal, que gera um executável.
- import "fmt": Importa o pacote fmt para formatação e saída.
- func main(): Função de entrada, equivalente ao public static void main em Java.

Compilação e execução:

go run hello.go  # Executa diretamente

go build hello.go  # Compila para um binário

Comparação com Java:

- Em Java, classes e métodos estáticos são obrigatórios. Em Go, a função main é suficiente.
- Go não usa ponto e vírgula (;) ao final das linhas.

  

5. Tipos Primitivos, Funções, Variáveis e Constantes

Tipos Primitivos

Go possui tipos primitivos simples, com tipagem estática (semelhante a Java, mas mais concisa).

|   |   |   |
|---|---|---|
|Tipo|Descrição|Exemplo|
|int, int32, int64|Inteiros com ou sem sinal|42, -10|
|float32, float64|Números de ponto flutuante|3.14, -0.001|
|bool|Booleano|true, false|
|string|Cadeia de caracteres UTF-8|"Hello, Go!"|
|byte|Alias para uint8|65 (equivalente a ‘A’)|

Nota: Go não suporta tipos implícitos como var genérico em Java. A inferência de tipo é feita com :=.

Variáveis

- Declaração explícita:  
    var nome string = "Grok"
- var idade int = 42
-   
    
- Declaração curta (inferência de tipo):  
    nome := "Grok"
- idade := 42
-   
    
- Múltiplas variáveis:  
    var x, y int = 10, 20
- a, b := "hello", true
-   
    

Comparação com Java:

- Em Java: String nome = "Grok";
- Em Go: nome := "Grok" (mais conciso, mas estritamente tipado).

Constantes

Constantes são definidas com const e não podem ser alteradas.

const Pi float64 = 3.14159

const NomeProjeto = "MeuApp"

Funções

Funções em Go são declaradas com func, podem retornar múltiplos valores e não suportam sobrecarga (diferente de Java).

Exemplo:

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

Caso de uso: Funções com múltiplos retornos são úteis para tratamento de erros, como value, err := funcao().

  

📌 Lab: Hello World, Manipulação de Variáveis e Tipos

Objetivo

Criar um programa simples em Go para praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

Passo a Passo

1. Configuração:

- Instale o Go e verifique com go version.
- Crie um diretório lab1 e inicialize um módulo:mkdir lab1
- cd lab1
- go mod init github.com/seu-usuario/lab1
-   
    

3. Hello World: Crie um arquivo hello.go:  
    package main
4.   
    
5. import "fmt"
6.   
    
7. func main() {
8.     fmt.Println("Hello, World!")
9. }
10.   
    

``. Execute:

go run hello.go

3. Manipulação de Variáveis e Tipos: Crie um arquivo vars.go:  
    package main
4.   
    
5. import "fmt"
6.   
    
7. func main() {
8.     // Declaração de variáveis
9.     var nome string = "Grok"
10.     idade := 42
11.     const versao = "1.0.0"
12.   
    
13.     // Tipos
14.     var preco float64 = 99.99
15.     ativo := true
16.   
    
17.     // Exibição
18.     fmt.Printf("Nome: %s, Idade: %d, Versão: %s\n", nome, idade, versao)
19.     fmt.Printf("Preço: %.2f, Ativo: %t\n", preco, ativo)
20.   
    
21.     // Função com múltiplos retornos
22.     x, y := troca("Go", "Java")
23.     fmt.Printf("Troca: %s, %s\n", x, y)
24. }
25.   
    
26. func troca(a, b string) (string, string) {
27.     return b, a
28. }
29.   
    Execute:  
    go run vars.go
30.   
    
31. Tarefa:

- Modifique vars.go para incluir uma função que calcula o dobro de um número.
- Adicione uma constante para o IVA (Imposto sobre Valor Agregado, ex.: 0.21) e calcule o preço final de um produto.

Saída esperada:

Nome: Grok, Idade: 42, Versão: 1.0.0

Preço: 99.99, Ativo: true

Troca: Java, Go

Caso de uso prático: Este lab simula a manipulação de dados em um sistema simples, como um cadastro de produtos, onde variáveis e constantes são usadas para armazenar informações fixas (versão) e dinâmicas (preço, status).

  

Conclusão

Este módulo cobre os fundamentos do Go, desde sua história até a criação de programas básicos. Engenheiros Java notarão semelhanças (tipagem estática, funções) e diferenças (sintaxe concisa, ausência de classes). O lab prático reforça a configuração do ambiente e o uso de tipos, variáveis e funções, preparando o time para módulos mais avançados, como estruturas de controle e concorrência.

Próximos passos: No próximo módulo, exploraremos estruturas de controle, slices, maps e pacotes, com mais exemplos práticos.
