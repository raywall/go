---
sidebar_position: 4
sidebar_label: Migrations
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Migrations com `golang-migrate`

## `golang-migrate`

- Biblioteca para _gerenciar migrações de banco de dados_, semelhante ao `Flyway` ou `Liquibase` em Java.
- **Suporta SQL ou arquivos Go** para _definir migrações_.

### Instalação

```bash
go get -u github.com/golang-migrate/migrate/v4
```

### Exemplo (migração SQL)

Crie arquivos de migração:

```dirtree
migrations/
├── 202506120001_create_produtos.up.sql
└── 202506120001_create_produtos.down.sql
```

- No arquivo `202506120001_create_produtos.up.sql`, temos:

```sql
CREATE TABLE produtos (
    id UUID PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    preco DOUBLE PRECISION NOT NULL
);
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod9/202506120001_create_produtos.up.sql)!

- No arquivo `202506120001_create_produtos.down.sql`, temos:

```sql
DROP TABLE produtos;
```

> Para download do código-fonte, clique [aqui](@site/static/code/mod9/202506120001_create_produtos.down.sql)!

- Código para aplicar migrações:

<InteractiveCodeSnippet 
    src="code/mod9/golang-migrate.go" 
    allowExecute={false} 
    allowEdit={false} />

### Comparação com Java

#### Java

- `Flyway` e `Liquibase` gerenciam migrações com SQL ou Java.

#### Go

- `golang-migrate` é **mais simples**, focado em SQL ou Go.

:::info Caso de uso
Migrações **garantem consistência do schema** em ambientes de desenvolvimento, teste e produção.
:::
