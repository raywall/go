---
sidebar_position: 3
sidebar_label: Channels
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Channels (`unbuffered`, `buffered`)

## Channels

- Mecanismo de comunicação entre goroutines, permitindo sincronização e troca de dados.
- Declarados com `chan tipo`

- Dois tipos:
  - **Unbuffered**: Bloqueia até que remetente e receptor estejam prontos.
  - **Buffered**: Permite enviar até a capacidade do buffer sem bloquear.

### Exemplos

#### Unbuffered channel

<InteractiveCodeSnippet 
    src="code/mod5/channel-unbuffered.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Buffered channel

<InteractiveCodeSnippet 
    src="code/mod5/channel-buffered.go" 
    allowExecute={true} 
    allowEdit={false} />

## Comparação com Java

### Java

- Usa `BlockingQueue` ou sincronização com `synchronized/Lock`

### Go

- Channels são mais simples e seguros, evitando condições de corrida.

:::info Caso de uso
Channels são usados para coordenar goroutines, como em pipelines de dados ou comunicação entre trabalhadores.
:::
