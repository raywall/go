---
sidebar_position: 1
sidebar_label: Módulo 09
---

# Persistência com banco de dados em Go

<div className="row">
<div className="col">

Este módulo aborda a persistência de dados em Go, usando o pacote padrão `database/sql`, o ORM `GORM`, migrações com `golang-migrate`, e _testes de integração com bancos de dados_. O conteúdo é voltado para engenheiros Java, com comparações ao Spring Data e Hibernate, e mantém-se objetivo com exemplos práticos e casos de uso.

## Conteúdo

- [Drivers e `database/sql`](./1-drivers-database-sql.md)
- [ORM com `gorm`](./2-gorm.md)
- [Migrations com `golang-migrate`](./3-golang-migrate.md)
- [Repositórios e testes de integração com DB](./4-repo-testes-integracao.md)

📌 [Lab](./5-laboratorio.md): Persistir o CRUD em banco real (PostgreSQL por exemplo).

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-sql.png').default} 
    style={{ transform:'scale(0.85)', marginTop:'-1.5rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

O lab prático implementa o CRUD dos módulos anteriores com persistência em `PostgreSQL`, incluindo _migrações e testes de integração_.
