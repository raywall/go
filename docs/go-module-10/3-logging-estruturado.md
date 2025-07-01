---
sidebar_position: 4
sidebar_label: Logging estruturado
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Logging estruturado (`slog` e `zap`)

## `slog`

- Pacote padrão do Go (1.21+) para logging estruturado, com saída JSON.
- Já usado nos módulos anteriores.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod10/logging-estruturado-slog.go" 
    allowExecute={true} 
    allowEdit={false} />

## `zap`

- Biblioteca de logging de alta performance, alternativa a slog.
- **Instalação**: `go get go.uber.org/zap`

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod10/logging-estruturado-zap.go" 
    allowExecute={true} 
    allowEdit={false} />

### Comparação com Java

#### Java

- Usa `SLF4J`/`Logback` para _logging estruturado_.

#### Go

- `slog` **é nativo e simples**; `zap` é _mais rápido para sistemas de alta carga_.

:::info Caso de uso
Logging estruturado é essencial para monitoramento em ferramentas como `ELK Stack` ou `Grafana Loki`.
:::
