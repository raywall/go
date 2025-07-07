---
sidebar_position: 4
sidebar_label: Métodos em structs
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Métodos em structs

- Go não possui classes, mas métodos podem ser associados a structs usando receivers
- Receivers podem ser por valor (`T`) ou por ponteiro (`*T`)

## Exemplo

<InteractiveCodeSnippet 
    src="code/mod3/metodos-em-structs.go" 
    allowExecute={true} 
    allowEdit={false} />

## Comparação com Java

- Em Java, métodos são definidos dentro de classes. Em Go, métodos são funções com receivers, associados a tipos (geralmente structs).
- Receiver por ponteiro (`*T`) é semelhante a modificar o estado de um objeto em Java.

:::info Caso de uso
Métodos em structs são usados para encapsular comportamento, como formatar dados ou aplicar regras de negócio.
:::
