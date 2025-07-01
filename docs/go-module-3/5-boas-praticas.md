---
sidebar_position: 6
sidebar_label: Boas práticas e princípios de design
---

# Boas práticas e princípios de design idiomático em Go

## Simplicidade

- Prefira soluções claras e concisas
- Evite complexidade desnecessária (ex: camadas excessivas de abstração)

## Nomenclatura

- Use nomes curtos e descritivos (ex: p para Produto, em vez de `produtoInstance`)
- Funções e métodos exportados começam com maiúscula (ex: `Descricao`)

## Tratamento de erros

- Sempre cheque erros retornados (`if err != nil`)
- Evite panic em código de produção, exceto em casos extremos.

## Interfaces pequenas

- Defina interfaces com poucos métodos, focadas em uma única responsabilidade.
- **Exemplo**: `io.Reader` e `io.Writer` em vez de interfaces genéricas.

## Composição sobre herança

- Use `structs aninhados` para composição, em vez de herança como em Java.

## Evite interface vazia (`interface{}`)

- Use type assertions ou type switches apenas quando necessário

## Use Go idiomaticamente

- Evite padrões de Java (ex: `getters`/`setters` desnecessários)
- Aproveite `go fmt` e ferramentas como `golint` para consistência

:::info Caso de uso
Boas práticas garantem código legível e manutenível, especialmente em equipes grandes, como em projetos de microsserviços.
:::
