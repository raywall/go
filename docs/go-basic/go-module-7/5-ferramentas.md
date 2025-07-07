---
sidebar_position: 6
sidebar_label: Ferramentas
---

# Ferramentas: `go vet`, `golint`, `staticcheck`

## `go vet`

- Analisa código para erros comuns (ex: formatação de strings incorreta)
  **Executado com**: `go vet ./...`

## `golint` (obsoleto, mas ainda usado)

- Sugere melhorias de estilo (ex: nomes de variáveis)
  **Alternativa moderna**: `golangci-lint`

## `staticcheck`

- Ferramenta avançada de análise estática, verificando bugs, desempenho e estilo.
  **Instalação**: `go install honnef.co/go/tools/cmd/staticcheck@latest`
  **Uso**: `staticcheck ./...`

## Comparação com Java

### Java

- Usa `Checkstyle`, `PMD` ou `SonarQube`

### Go

- Ferramentas são mais `integradas` e `leves`, com `foco em simplicidade`.

:::info Caso de uso
Essas ferramentas garantem consistência e qualidade em projetos grandes, como microsserviços.
:::
