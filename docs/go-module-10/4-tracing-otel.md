---
sidebar_position: 5
sidebar_label: Tracing com OpenTelemetry
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Tracing com OpenTelemetry

## OpenTelemetry

- Framework para _tracing distribuído, métricas e logs_, compatível com `Jaeger`, `Zipkin`, etc.
- Permite rastrear requisições em microsserviços.

### Instalação

```bash
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace
go get go.opentelemetry.io/otel/sdk/trace
```

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod10/opentelemetry.go" 
    allowExecute={false} 
    allowEdit={false} />

### Comparação com Java

#### Java

- Usa `Spring Cloud Sleuth` ou `Micrometer` para tracing.

#### Go

- `OpenTelemetry` é **mais flexível**, com integração nativa.

:::info Caso de uso
`Tracing` é _crucial para diagnosticar latências_ em sistemas distribuídos, como APIs em Kubernetes.
:::
