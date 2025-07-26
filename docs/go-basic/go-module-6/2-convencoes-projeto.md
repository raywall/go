---
sidebar_position: 3
sidebar_label: Convenções de projeto
---

# Convenções de projeto (`cmd`, `internal` e `pkg`)

## Convenções comuns

### `cmd/`

- Contém pontos de entrada (`main`) para executáveis.

#### Exemplos

- `cmd/api/main.go` para uma API
- `cmd/cli/main.go` para uma ferramenta CLI

### `internal/`

- Utilizado para pacotes privados, acessíveis apenas dentro do projeto ou subdiretórios.
- Ideal para lógica sensível ou reutilização interna.

### `pkg/`

- Pacotes reutilizáveis por outros projetos (opcional, menos comum).

### Outros diretórios

- Nomes descritivos como: api, db, models, utils.

## Exemplo de estrutura

```dirtree
projeto/
├── go.mod
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── repo/
│       └── memoria.go
└── models/
    └── produto.go
```

## Comparação com Java

### Java

- Estrutura rígida com `src/main/java` e **pacotes hierárquicos**.

### Go

- Mais flexível, mas `internal` é semelhante a `package-private` em Java.

:::info Caso de uso
Usar internal para proteger implementações de repositórios, enquanto cmd organiza múltiplos executáveis (ex: servidor e ferramenta de migração).
:::
