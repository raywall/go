---
sidebar_position: 5
sidebar_label: Repositórios e testes
---

# Repositórios e testes de integração com banco de dados

## Repositórios

- Abstraem o acesso ao banco, usando `database/sql` ou `GORM`.
- Seguem o padrão `Repository`, como no Spring Data.

## Testes de integração

- Testam a interação com o banco, usando containers (ex: `Testcontainers`) ou bancos em memória (ex: `SQLite`).
- Requerem configuração do ambiente de teste.

:::info Caso de uso
Testes de integração garantem que queries e mapeamentos funcionem corretamente em um banco real.
:::
