---
sidebar_position: 2
sidebar_label: Servidor HTTP
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Servidor HTTP com `net/http`

## Pacote `net/http`

- Fornece funcionalidades para `criar servidores HTTP` e `lidar com requisições/respostas`.
- **Simples, mas poderoso**, usado em muitos projetos Go sem frameworks.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod8/api-http.go" 
    allowExecute={false} 
    allowEdit={false} />

### Execução

```bash
go run main.go
```

> Acesse via browser a partir do endereço [http://localhost:8080/hello](http://localhost:8080/hello)

### Comparação com Java

#### Java

- `Spring Boot` ou `Servlets` requerem mais configuração (ex: `@RestController`)

#### Go

- `net/http` é _mais leve, com menos abstrações_.

:::info Caso de uso
Servidores simples, como ferramentas internas ou APIs mínimas.
:::
