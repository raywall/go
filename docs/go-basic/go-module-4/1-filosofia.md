---
sidebar_position: 2
sidebar_label: Filosofia de erros explícitos
---

# Filosofia do Go: erros explícitos

- **Erros como valores**: Em Go, erros são valores do tipo error, retornados explicitamente pelas funções, ao invés de lançados como exceções em Java (`try-catch`).
- **Clareza e previsibilidade**: O programador deve verificar erros explicitamente, reduzindo comportamentos inesperados.
- **Evite `panic`**: Diferente de Java, panic é reservado para falhas irrecuperáveis (ex: falha de inicialização crítica) e não para fluxo normal.

## Comparação com Java

### Java

```java
try { ... } catch (Exception e) { ... }
```

### Go

```go
if err != nil { return err }
```

:::info Caso de uso
A abordagem explícita é ideal para sistemas onde a robustez é crítica, como APIs de servidores ou ferramentas CLI.
:::
