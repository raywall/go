---
sidebar_position: 4
sidebar_label: Wrapping e unwrapping
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Wrapping e Unwrapping com `errors.Is` e `errors.As`

## Wrapping

- Permite adicionar contexto a um erro, mantendo o erro original
- Usado com o operador `%w` em `fmt.Errorf`

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/wrapping.go" 
    allowExecute={true} 
    allowEdit={false} />

## `errors.Is`

- Verifica se um erro (ou seus erros encapsulados) corresponde a um erro específico.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/errors-is.go" 
    allowExecute={true} 
    allowEdit={false} />

## `errors.As`

- Extrai um erro de um tipo específico de uma cadeia de erros encapsulados.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/errors-as.go" 
    allowExecute={true} 
    allowEdit={false} />

### Comparação com Java

- `Wrapping` em Go é semelhante a `Exception.getCause()` em Java, mas mais explícito.
- `errors.Is` e `errors.As` substituem verificações como instanceof ou comparações de tipos de exceção.

:::info Caso de uso
Wrapping é útil em APIs para adicionar contexto (ex: ID da requisição) sem perder o erro original.
:::
