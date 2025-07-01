---
sidebar_position: 4
sidebar_label: Framework Gin
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Framework Gin: roteamento, binding, validação

## Gin

- Framework leve para APIs RESTful, com `roteamento rápido` e `suporte a JSON`, `binding` e `validação`.
- **Instalação**: `go get github.com/gin-gonic/gin`

### Roteamento

- Usa `grupos de rotas` e `métodos HTTP` (`GET`, `POST`, `PUT`, `DELETE` etc.)
- Suporta `parâmetros de URL` e `query strings`.

#### Exemplo

<InteractiveCodeSnippet 
    src="code/mod8/api-gin-roteamento.go" 
    allowExecute={true} 
    allowEdit={false} />

### Binding e validação

- Converte JSON ou formulários em `structs`, com `validação via tags`.
- Requer biblioteca `validator`: `go get github.com/go-playground/validator/v10`

#### Exemplo

<InteractiveCodeSnippet 
    src="code/mod8/api-gin-binding-validacao.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Comparação com Java:

##### Java

- Spring Boot usa `@RestController`, `@RequestBody` e validação com `@Valid`.

##### Go

- Gin é **mais leve**, com `binding` e `validação` **diretamente via tags**.

:::info Caso de uso
Gin é ideal para `APIs RESTful escaláveis`, como sistemas de e-commerce ou microsserviços.
:::
