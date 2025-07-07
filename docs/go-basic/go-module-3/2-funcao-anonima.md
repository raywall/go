---
sidebar_position: 3
sidebar_label: Funções anônimas e closures
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Funções anônimas e closures

## Funções anônimas

- Funções sem nome, definidas inline, semelhantes a lambdas em Java
- Úteis para tarefas temporárias ou callbacks

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/funcao-anonima.go" 
    allowExecute={true} 
    allowEdit={false} />

## Closures

- Funções anônimas que capturam variáveis do escopo externo
- Similar a `closures` em Java (ex: lambdas com acesso a variáveis externas)

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/closures.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
Closures são úteis para manter estado em funções, como contadores ou geradores, sem a necessidade de `structs`.
:::
