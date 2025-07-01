---
sidebar_position: 5
sidebar_label: Interfaces e duck typing
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Interfaces e duck typing

## Interfaces

- Definidas com type Nome interface, especificando métodos que um tipo deve implementar.
- **Go usa `duck typing`**: se um tipo implementa os métodos de uma interface, ele a satisfaz automaticamente (sem declaração explícita, diferente de Java com implements).

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/interfaces.go" 
    allowExecute={true} 
    allowEdit={false} />

### Comparação com Java

- Em Java, interfaces exigem `implements`. Em Go, a implementação é implícita.
- Go não suporta herança, mas interfaces permitem polimorfismo.

## Interface vazia (`interface{}`)

- Equivalente a `Object` em Java, aceita qualquer tipo.
- Usada com moderação, geralmente com type assertions ou type switches.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/interface-vazia.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
Interfaces são amplamente usadas em Go para abstrair repositórios, serviços ou plugins, garantindo flexibilidade e testabilidade.
:::
