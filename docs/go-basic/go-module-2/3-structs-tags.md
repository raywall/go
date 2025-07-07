---
sidebar_position: 4
sidebar_label: Structs e tags de struct
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

## Structs e tags de struct

### Structs

- Equivalentes a classes em Java, mas sem herança
- Definidas com `type Nome struct`

#### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/structs.go" 
    allowExecute={true} 
    allowEdit={false} />

### Tags de struct

- Usadas para metadados, como serialização JSON
- Declaradas com crases (`\``)

#### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/tags-de-struct.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Comparação com Java

Structs substituem classes, mas tags são semelhantes às anotações (`@JsonProperty`) do Jackson em Java

:::info Caso de uso
Structs com tags são amplamente usadas em APIs REST para `serializar`/`deserializar` JSON
:::
