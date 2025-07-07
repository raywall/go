---
sidebar_position: 2
sidebar_label: Funções com múltiplos retornos
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Funções com múltiplos retornos

Go permite que funções retornem múltiplos valores, um recurso poderoso para tratamento de erros e resultados compostos, diferente de Java, que usa exceções ou objetos.

## Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/funcao-multiplos-retornos.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
Funções com múltiplos retornos são comuns em Go para retornar um valor e um erro, substituindo o uso de exceções como em Java (`try-catch`).
:::

<br />
