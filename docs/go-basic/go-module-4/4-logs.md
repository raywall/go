---
sidebar_position: 5
sidebar_label: Pacote logs e log/slog
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Pacote `log` e `log/slog`

## Pacote `log`

- Fornece logging básico, com saída para `stderr` por padrão.
- Funções como `log.Print`, `log.Fatal` e `log.Panic`.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/log.go" 
    allowExecute={true} 
    allowEdit={false} />

## Pacote `log/slog` (Go 1.21+)

- Oferece logging estruturado, com suporte a JSON e campos personalizados.
- Mais adequado para aplicações modernas, como microsserviços.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod4/slog.go" 
    allowExecute={true} 
    allowEdit={false} />

### Saída (JSON)

```json
    {"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Processando requisição","id":45,"metodo":"GET"}
    {"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"Erro encontrado","error":"falha na operação","id":45}
```

### Comparação com Java

- `log` é semelhante ao `System.out.println` ou `java.util.logging`
- `log/slog` é comparável a bibliotecas como `Log4j` ou `SLF4J`, com suporte a logging estruturado.

:::info Caso de uso
`log/slog` é ideal para aplicações distribuídas, onde logs estruturados facilitam a análise em ferramentas como `ELK Stack` ou `Prometheus`.
:::
