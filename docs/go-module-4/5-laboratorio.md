---
sidebar_position: 5
sidebar_label: Laboratório
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Laboratório

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c', marginTop:'2rem', marginBottom:'5rem' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

## Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab4`

```bash
mkdir lab4
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab4
go mod init github.com/seu-usuario/lab4
```

</TabItem>
<TabItem value="app">

## Criar Funções com tratamento de erros e logging estruturado

### Objetivo

Refatorar o CRUD do Módulo 3 para incluir tratamento de erros robusto e logging estruturado com log/slog, mantendo a interface de repositório.

### Passo a passo

1. Crie um arquivo `crud.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod4/lab/crud.go" 
    allowExecute={true} 
    allowEdit={false} />

2. Execução:

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione um erro personalizado para quando o nome do produto estiver vazio.
- Implemente uma função que use errors.As para capturar e tratar erros específicos (ex.: ErrPrecoInvalido).
- Configure o logger para salvar logs em um arquivo, além de exibir no console.

### Saída esperada

#### Console

```bash
Lista de produtos:
ID: 1, Nome: Laptop, Preço: 999.99
ID: 2, Nome: Mouse, Preço: 29.99

Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
Erro esperado: preço não pode ser negativo
Produto deletado com sucesso

Lista final:
ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto criado","id":1,"nome":"Laptop","preco":999.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto criado","id":2,"nome":"Mouse","preco":29.99}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listando produtos","total":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto encontrado","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto atualizado","id":1,"nome":"Laptop Pro","preco":1299.99}
{"time":"2025-06-12T01:05:00Z","level":"ERROR","msg":"Falha ao atualizar produto","error":"preço não pode ser negativo","id":1}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Produto deletado","id":2}
{"time":"2025-06-12T01:05:00Z","level":"INFO","msg":"Listando produtos","total":1}
```

:::info Caso de uso prático
Este lab simula um sistema real com logging estruturado, útil para depuração e monitoramento em aplicações distribuídas, como APIs RESTful.
:::
</TabItem>
</Tabs>
