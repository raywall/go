---
sidebar_position: 4
sidebar_label: Testes mocados
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Testes com `mocks` (`testify` e `gomock`)

## Pacote `testify`

- Simplifica `asserções` e `criação de mocks`
- Inclui `assert` e `mock` para facilitar testes

### Instalação

```bash
go get github.com/stretchr/testify
```

### Exemplo com testify/mock

<InteractiveCodeSnippet 
    src="code/mod7/teste-mock.go" 
    allowExecute={false} 
    allowEdit={false} />

## Pacote `gomock`

- **Gera mocks automaticamente** a partir de `interfaces`
- **Mais poderoso** para projetos complexos

### Instalação

```bash
go install github.com/golang/mock/gomock@latest
go install github.com/golang/mock/mockgen@latest
```

:::info Caso de uso
`Mocks` são úteis para testar _serviços que dependem de repositórios_ ou `APIs externas`, similar ao `Mockito` em Java.
:::
