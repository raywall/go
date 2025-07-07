---
sidebar_position: 5
sidebar_label: Gestão de dependências
---

# Gerenciamento de dependências com `go get` e `replace`

## `go get`

- Adiciona ou atualiza dependências no `go.mod`
  Exemplo: `go get github.com/google/uuid@v1.6.0`

- Atualiza para a versão mais recente:
  `go get -u github.com/google/uuid`

## `replace`

- Permite substituir uma dependência por outra versão ou caminho local.
- Útil para desenvolvimento `local` ou `forks`

## Exemplo de `go.mod` com `replace`

```text
module github.com/seu-usuario/meu-projeto

go 1.21

require github.com/google/uuid v1.6.0
replace github.com/google/uuid => ../uuid-fork
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod6/go-replace.mod)!

## Comparação com Java

### Java

- Usa no `Maven` ou `implementation` no Gradle

### Go

- `go get` e `replace` são mais simples, mas menos configuráveis

:::info Caso de uso
`replace` é útil para testar alterações locais em dependências antes de publicar
:::
