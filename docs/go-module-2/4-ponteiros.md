---
sidebar_position: 5
sidebar_label: Ponteiros
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Ponteiros

- Go suporta ponteiros para manipulação direta de memória, mas de forma segura
- Declarados com `*tipo` (endereço) e `&variavel` (referência)

## Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/ponteiros.go" 
    allowExecute={true} 
    allowEdit={false} />

## Comparação com Java

- Java usa referências implícitas para objetos, sem controle direto de ponteiros
- Go exige ponteiros explícitos para alterações em funções (passagem por valor é padrão)

:::info Caso de uso
Ponteiros são úteis para modificar structs grandes sem copiar dados, similar a passar objetos por referência em Java.
:::
