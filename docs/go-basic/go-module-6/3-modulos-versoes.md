---
sidebar_position: 4
sidebar_label: Modulos e versionamnto
---

# `go mod` e versionamento

## Go Modules

- Introduzidos no Go 1.11, substituem `$GOPATH` para gerenciar dependências.
- O arquivo `go.mod` define o módulo e suas dependências.

### Exemplo de go.mod

```text
module github.com/seu-usuario/meu-projeto

go 1.21

require (
    github.com/google/uuid v1.6.0
)
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod6/go.mod)!

### Comandos Principais

- `go mod init`: Inicializa um módulo.
- `go mod tidy`: Adiciona dependências usadas e remove não usadas.
- `go mod vendor` (opcional): Cria um diretório vendor/ com dependências.

## Versionamento

- Go usa `versionamento semântico` (ex: v1.6.0)
- Módulos são referenciados por URLs de repositórios (ex: `github.com/autor/projeto`)
- Tags no repositório (ex: `git tag v1.0.0`) definem versões

### Comparação com Java

#### Java

- Usa `Maven/Gradle` com repositórios centralizados (ex: `Maven Central`)

#### Go

- Usa URLs de repositórios diretamente, com proxy (ex: `proxy.golang.org`)

:::info Caso de uso
Go Modules simplificam o gerenciamento de dependências em projetos grandes, como APIs que integram bibliotecas externas.
:::
