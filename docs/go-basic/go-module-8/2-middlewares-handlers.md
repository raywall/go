---
sidebar_position: 3
sidebar_label: Middlewares e handlers
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Middlewares e handlers

## Handlers

- Funções que **lidam com requisições HTTP**, com assinatura `func(w http.ResponseWriter, r *http.Request)`
- Podem ser registrados com `http.HandleFunc`

## Middlewares

- Funções que **interceptam requisições/respostas**, úteis para `autenticação`, `logging`, etc.
- Encadeiam `handlers`, retornando um `http.Handler`

### Exemplo (middleware com logging)

<InteractiveCodeSnippet 
    src="code/mod8/api-middleware.go" 
    allowExecute={false} 
    allowEdit={false} />

### Comparação com Java

#### Java

- `Middlewares` são filtros no Spring (`@Filter`) ou `interceptores`

#### Go

- `Middlewares` são **mais explícitos**, usando composição de `handlers`

:::info Caso de uso
Middlewares para `logging`, `autenticação` (ex: JWT), ou `validação de cabeçalhos`.
:::
