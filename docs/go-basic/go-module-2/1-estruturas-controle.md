---
sidebar_position: 2
sidebar_label: Estruturas de controle
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Estruturas de controle

Go possui estruturas de controle simples, sem parênteses obrigatórios (diferente de Java) e com foco em clareza.

## `if`

- Suporta inicialização na própria declaração.
- Não há operador ternário em Go.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/controle-if.go" 
    allowExecute={true} 
    allowEdit={false} />

### Comparação com Java

#### Java

```java
if (x % 2 == 0) { ... }
```

#### Go

```go
if x := numero % 2; x == 0 { ... } // inicialização direta
```

## `for`

- Única estrutura de laço em Go (não há `while` ou `do-while`)
- Suporta quatro formas: `tradicional`, condição `simples`, `infinito` e `range`

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/controle-for.go" 
    allowExecute={true} 
    allowEdit={false} />

## `switch`

- Mais simples que em Java: não precisa de break (não há fall-through por padrão)
- Suporta múltiplas condições por caso e inicialização

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/controle-switch.go" 
    allowExecute={true} 
    allowEdit={false} />

## `defer`

- Adia a execução de uma instrução até o fim da função
- Útil para liberar recursos (similar ao `try-with-resources` em Java)

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/controle-defer.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
`defer` é usado para fechar arquivos, conexões de banco ou liberar mutexes
:::
