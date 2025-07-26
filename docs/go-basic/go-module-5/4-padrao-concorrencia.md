---
sidebar_position: 5
sidebar_label: Padrões de concorrência
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Padrões de concorrência em Go

## Pipeline:

- Encadeia `goroutines` com `channels` para processar dados em etapas.
  **Exemplo**: Ler dados → Transformar → Salvar.

## Worker pool

- Um `grupo de goroutines` processa tarefas de uma fila compartilhada.
- Similar a `ThreadPoolExecutor` em Java.

## Fan-out e Fan-in

- **Fan-out**: Distribui tarefas para `múltiplas goroutines`
- **Fan-in**: Combina resultados em `um único channel`

## Cancelamento

- Usa `context.Context` para interromper goroutines (ex: timeout de requisições).

## Exemplo

### Pipeline

<InteractiveCodeSnippet 
    src="code/mod5/pipeline.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
`Worker pools` são comuns em sistemas de alta carga, como servidores web ou processamento de lotes.
:::
