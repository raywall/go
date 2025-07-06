---
sidebar_position: 3
sidebar_label: Docker com Go
---

# Docker com Go

## Docker

- Go é **altamente compatível com Docker**, gerando **binários pequenos** e **imagens leves**.
- Usa `multi-stage builds` para **reduzir o tamanho da imagem**.

### Exemplo

```Dockerfile
# Etapa de build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod10/Dockerfile)!

### Construir e executar

```bash
docker build -t meu-app .
docker run -p 8080:8080 meu-app
```

### Comparação com Java

#### Java

- Imagens _Spring Boot são maiores **devido à JVM**_.

#### Go

- **Imagens são menores** (ex: `~10MB com Alpine`), ideais para microsserviços

:::info Caso de uso
**Containerizar APIs** para deploy em `Kubernetes` ou plataformas como `AWS ECS`.
:::
