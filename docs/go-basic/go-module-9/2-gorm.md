---
sidebar_position: 3
sidebar_label: GORM
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# ORM com `GORM`

## `GORM`

- Biblioteca popular que _fornece uma API ORM_, semelhante ao `Hibernate` em Java.
- Suporta `PostgreSQL`, `MySQL`, `SQLite`, etc.
- Possui funcionalidades como: `mapeamento de structs`, `associações` e `migrações automáticas`.

### Instalação

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod9/gorm.go" 
    allowExecute={false} 
    allowEdit={false} />

### Comparação com Java

#### Java

- `Hibernate`/`Spring Data` usa anotações como `@Entity` e `@Column`.

#### Go

- `GORM` usa tags (`gorm:"..."`) e é _mais leve, com menos overhead_.

:::info Caso de uso
`GORM` é ideal para APIs RESTful que _precisam de mapeamento objeto-relacional rápido_, como sistemas de e-commerce.
:::
