---
sidebar_position: 5
sidebar_label: JSON, status codes e headers
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# JSON, status codes e headers

## JSON

- Gin `serializa/deserializa JSON automaticamente` com `c.JSON`
- Usa `encoding/json` internamente

## Status codes

- Usa `constantes do net/http` (ex: `http.StatusOK`, `http.StatusBadRequest`)

## Headers

- Manipulados com `c.Header` ou `c.GetHeader`

## Exemplo

<InteractiveCodeSnippet 
    src="code/mod8/api-gin-json.go" 
    allowExecute={true} 
    allowEdit={false} />

## Teste

Utilize a ferramenta `curl`

```bash
curl -X GET http://localhost:8080/api
```

### Resposta

```json
{ "mensagem": "API funcionando" }
```

:::info Caso de uso
`JSON` para respostas de APIs, `status codes` para indicar sucesso/erro, e `headers` para autenticação ou metadados.
:::
