---
sidebar_position: 2
sidebar_label: Estrutura de pacotes
---

# Estrutura de pacotes idiomática

## O que são pacotes?

- Pacotes são `unidades de organização de código` em Go, semelhantes a pacotes em Java (`package`).
- Cada arquivo Go pertence a um pacote, declarado com `package` nome.
- O pacote main gera um executável; outros pacotes são bibliotecas.

## Estrutura idiomática

- Go recomenda uma organização clara para projetos:

- Raiz do projeto: Contém go.mod e subdiretórios.
- Pacotes: Organizados em subdiretórios com nomes descritivos (ex.: api, models, repo).
- Nomenclatura: Nomes de pacotes são curtos, em minúsculas, sem sublinhados (ex.: http, db).
- Exportação: Funções, tipos e variáveis com inicial maiúscula são exportados (ex.: func Criar é visível fora do pacote).

## Exemplo

```dirtree
meu-projeto/
├── go.mod
├── main.go
├── models/
│   └── produto.go
└── repo/
    └── memoria.go
```

## Comparação com Java

### Java

- Pacotes seguem a convenção `com.empresa.projeto`

### Go

- Pacotes usam caminhos baseados no módulo (ex: `github.com/usuario/projeto/models`)

:::info Caso de uso
Pacotes são usados para modularizar código em microsserviços, separando lógica de negócio, acesso a dados e APIs.
:::
