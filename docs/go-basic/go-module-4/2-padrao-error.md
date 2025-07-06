---
sidebar_position: 3
sidebar_label: Tratamento de erros
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Padrão `error`, `errors.New` e `fmt.Errorf`

## Tipo `error`

- **Uma interface embutida**: `type error interface { Error() string }`
- Qualquer tipo que implemente o método `Error() string` é um error

## `errors.New`

- Cria erros simples com uma mensagem fixa

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/error.go" 
    allowExecute={true} 
    allowEdit={false} />

## `fmt.Errorf`

- Formata mensagens de erro dinamicamente, semelhante a `String.format` em Java.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/errorf.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
`errors.New` é usado para erros fixos, enquanto `fmt.Errorf` é ideal para erros com contexto dinâmico (ex: validação de entrada)
:::
