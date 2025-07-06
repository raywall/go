---
sidebar_position: 2
sidebar_label: Pacote testing
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Testes com o pacote `testing`

## Pacote `testing`

- O pacote padrão testing fornece suporte para `testes unitários` e `benchmarks`.
- Testes são escritos em arquivos com sufixo `_test.go` e funções prefixadas com Test.
- Usa `t.Error` ou `t.Fatal` para reportar falhas.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod7/testing.go" 
    allowExecute={true} 
    allowEdit={false} />

### Execução

```bash
go test
```

:::tip Dica
Para executar o teste em todos os diretórios e subdiretórios do projeto utilize `go test ./...`
:::

### Comparação com Java

#### Java

- Usa `JUnit/TestNG` com anotações (`@Test`).

#### Go

- Mais `simples`, `sem anotações`, e com `convenções baseadas em nomes`.

:::info Caso de uso
Testes unitários com `testing` são ideais para validar funções puras ou componentes isolados.
:::
