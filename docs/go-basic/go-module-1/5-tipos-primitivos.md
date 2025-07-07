---
sidebar_position: 6
sidebar_label: Tipos primitivos, variáveis e funções
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Tipos primitivos, funções, variáveis e constantes

## Tipos primitivos

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

## Variáveis

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

## Comparação com Java

### Java

```java
    String nome = "Raywall";
```

### Go

```go
    nome := "Raywall" // mais conciso, mas estritamente tipado
```

## Constantes

Constantes são definidas com const e não podem ser alteradas.

```go
const Pi float64 = 3.14159
const NomeProjeto = "MeuApp"
```

## Funções

Funções em Go são declaradas com func, podem retornar múltiplos valores e não suportam sobrecarga (diferente de Java).

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod1/funcoes.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
Funções com múltiplos retornos são úteis para tratamento de erros, como value, err := funcao().
:::
