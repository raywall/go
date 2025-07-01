---
sidebar_position: 2
sidebar_label: Drivers e database/sql
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Drivers e `database/sql`

## Pacote `database/sql`

- Parte da biblioteca padrão de Go, _fornece uma interface genérica para bancos de dados SQL_.
- Requer um driver específico (ex: `github.com/lib/pq` para `PostgreSQL`)
- Usa conexão `pooling` e é `thread-safe`.

### Exemplo (`PostgreSQL`)

<InteractiveCodeSnippet 
    src="code/mod9/postgres-sql.go" 
    allowExecute={false} 
    allowEdit={false} />

### Instalação do driver `PostgreSQL`

```bash
go get github.com/lib/pq
```

### Comparação com Java

#### Java

- Usa `JDBC` ou `Spring Data` para abstrair acesso a bancos.

#### Go

- `database/sql` é mais explícito, com menos abstrações que `Hibernate`.

:::info Caso de uso
Ideal para aplicações que precisam de _controle fino sobre consultas SQL_, como relatórios ou sistemas legados.
:::
