---
sidebar_position: 5
sidebar_label: Benchmarks e profiling
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Benchmarks e profiling

## Benchmarks

- Funções prefixadas com `Benchmark` **medem desempenho**
- Usam `b.N` para _executar o código várias vezes_

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod7/benchmarks.go" 
    allowExecute={false} 
    allowEdit={false} />

### Execução

```bash
go test -bench=.
```

## Profiling

Go oferece ferramentas integradas para análise de desempenho:

- `pprof`: Perfil de CPU, memória e bloqueios
- `go test -cpuprofile cpu.out`: Gera perfil de CPU
- `go tool pprof cpu.out`: Analisa o perfil

### Exemplo

```bash
go test -bench=. -cpuprofile cpu.out
go tool pprof cpu.out
```

## Comparação com Java

### Java

- Usa `VisualVM` ou `JProfiler` para profiling

### Go

- Ferramentas embutidas (`pprof`) são mais leves e integradas

:::info Caso de uso
`Benchmarks` ajudam a otimizar funções críticas; enquanto o `profiling` **identifica gargalos** em _APIs_ ou _sistemas concorrentes_.
:::
