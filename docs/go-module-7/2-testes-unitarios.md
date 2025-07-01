---
sidebar_position: 3
sidebar_label: Testes de unidade e integração
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Testes de unidade e de integração

## Testes de unidade

- Testam `unidades isoladas` (ex: uma função ou método)
- Usam tabelas de teste para múltiplos cenários

### Exemplo (tabela de testes)

<InteractiveCodeSnippet 
    src="code/mod7/teste-unitario.go" 
    allowExecute={true} 
    allowEdit={false} />

## Testes de integração

- Testam a `interação entre componentes` (ex: repositório e API)
- Podem usar `bancos de dados em memória` ou `mocks`

:::info Caso de uso
Testes de unidade `verificam lógica de negócio`; já os testes de integração `garantem que módulos funcionem juntos` (ex: repositório e serviço)
:::
