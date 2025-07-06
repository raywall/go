---
sidebar_position: 3
sidebar_label: Arrays, slices e maps
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Arrays, slices e maps

## Arrays

- Tamanho fixo, definido na declaração
- Menos comum em Go, pois slices são mais flexíveis

### Exemplo

```go
var numeros [3]int = [3]int{1, 2, 3}
fmt.Println(numeros[0]) // Saída: 1
```

### Comparação com Java

Arrays em Go são semelhantes aos arrays de Java (`int[]`), mas menos usados devido aos slices

## Slices

- Estrutura dinâmica, baseada em arrays, mas com tamanho variável
- Declarados com `[]tipo` ou criados com make

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/slices.go" 
    allowExecute={true} 
    allowEdit={false} />

:::tip Nota
Slices são passados por referência, mas o array subjacente pode ser copiado se modificado
:::

## Maps

- Estrutura `chave-valor`, semelhante ao `HashMap` em Java
- Declarados com `map[tipoChave]tipoValor`

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod2/maps.go" 
    allowExecute={true} 
    allowEdit={false} />

:::info Caso de uso
Slices são ideais para listas dinâmicas (ex: lista de usuários), enquanto maps são úteis para associações rápidas (ex: cache de IDs para valores)
:::
