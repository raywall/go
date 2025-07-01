---
sidebar_position: 6
sidebar_label: Linter, cobertura e documentação
---

# Linter, cobertura e documentação com `godoc`

## Linter

- Ferramentas como `golangci-lint` _verificam estilo, erros e boas práticas_.

### Instalação

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Uso

```bash
golangci-lint run
```

## Cobertura

- Go gera relatórios de cobertura de testes nativamente.

### Exemplo

```bash
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```

## `godoc`

- Gera documentação automática a partir de comentários no código.

### Uso local

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
```

> Acesse a documentação via browser através do endereço [http://localhost:6060/pkg](http://localhost:6060/pkg)

### Exemplo de comentário

```go
// Criar adiciona um novo produto ao repositório.
// Retorna o produto criado ou um erro se o preço for inválido.
func (r *Repositorio) Criar(nome string, preco float64) (models.Produto, error) {
    // ...
}
```

### Comparação com Java

#### Java

- Usa `Javadoc`, `SonarQube` e `Checkstyle`.

#### Go

- **Ferramentas nativas** como `godoc`, `go test` e `golangci-lint` são mais integradas.

:::info Caso de uso
`Linters` _garantem consistência_, cobertura valida testes, e `godoc` facilita manutenção.
:::
